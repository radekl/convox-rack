package k8s

import (
	"archive/tar"
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"strings"
	"sync"
	"time"

	"github.com/convox/rack/helpers"
	cc "github.com/convox/rack/provider/k8s/pkg/client/clientset/versioned/typed/convox/v1"
	"github.com/convox/rack/structs"
	ac "k8s.io/api/core/v1"
	am "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type processLister func() (structs.Processes, error)

func (p *Provider) convoxClient() (cc.ConvoxV1Interface, error) {
	return cc.NewForConfig(p.Config)
}

func (p *Provider) podLogs(namespace, name string, opts structs.LogsOptions) (io.ReadCloser, error) {
	if err := p.podWaitForRunning(namespace, name); err != nil {
		return nil, err
	}

	lopts := &ac.PodLogOptions{
		// Container:  "main",
		Follow:     helpers.DefaultBool(opts.Follow, true),
		Timestamps: helpers.DefaultBool(opts.Prefix, false),
	}

	if opts.Since != nil {
		t := am.NewTime(time.Now().UTC().Add(-1 * *opts.Since))
		lopts.SinceTime = &t
	}

	r, err := p.Cluster.CoreV1().Pods(namespace).GetLogs(name, lopts).Stream()
	if err != nil {
		return nil, err
	}

	return r, nil
}

func (p *Provider) podWaitForRunning(namespace, name string) error {
	for {
		pd, err := p.Cluster.CoreV1().Pods(namespace).Get(name, am.GetOptions{})
		if err != nil {
			return err
		}

		for _, c := range pd.Status.ContainerStatuses {
			if c.State.Waiting == nil {
				return nil
			}
		}

		time.Sleep(1 * time.Second)
	}
}

func (p *Provider) streamProcessListLogs(w io.WriteCloser, pl processLister, opts structs.LogsOptions, ch chan error) {
	defer w.Close()
	defer close(ch)

	done := false
	var wg sync.WaitGroup

	go func() {
		time.Sleep(5 * time.Second)
		wg.Wait()
		done = true
	}()

	current := map[string]bool{}

	for {
		if _, err := w.Write([]byte{}); err != nil {
			ch <- err
			return
		}

		if done {
			return
		}

		time.Sleep(1 * time.Second)

		pss, err := pl()
		if err != nil {
			ch <- err
			continue
		}

		for _, ps := range pss {
			if !current[ps.Id] {
				wg.Add(1)
				go p.streamProcessLogsWait(w, ps, opts, &wg)
				current[ps.Id] = true
			}
		}
	}
}

func (p *Provider) streamProcessLogsWait(w io.WriteCloser, ps structs.Process, opts structs.LogsOptions, wg *sync.WaitGroup) {
	defer wg.Done()

	ri, wi := io.Pipe()
	go p.streamProcessLogs(wi, ps, opts)
	io.Copy(w, ri)
}

type imageManifest []struct {
	RepoTags []string
}

func extractImageManifest(r io.Reader) (imageManifest, error) {
	mtr := tar.NewReader(r)

	var manifest imageManifest

	for {
		mh, err := mtr.Next()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}

		if mh.Name == "manifest.json" {
			var mdata bytes.Buffer

			if _, err := io.Copy(&mdata, mtr); err != nil {
				return nil, err
			}

			if err := json.Unmarshal(mdata.Bytes(), &manifest); err != nil {
				return nil, err
			}

			return manifest, nil
		}
	}

	return nil, fmt.Errorf("unable to locate manifest")
}

func processFilter(in structs.Processes, fn func(structs.Process) bool) structs.Processes {
	out := structs.Processes{}

	for _, ps := range in {
		if fn(ps) {
			out = append(out, ps)
		}
	}

	return out
}

func streamLogsWithPrefix(w io.WriteCloser, r io.Reader, prefix string) {
	defer w.Close()

	ls := bufio.NewScanner(r)

	for ls.Scan() {
		parts := strings.SplitN(ls.Text(), " ", 2)

		ts, err := time.Parse(time.RFC3339Nano, parts[0])
		if err != nil {
			fmt.Printf("err = %+v\n", err)
			continue
		}

		fmt.Fprintf(w, "%s %s %s\n", ts.Format(helpers.PrintableTime), prefix, parts[1])
	}
}

func systemVolume(v string) bool {
	switch v {
	case "/var/run/docker.sock":
		return true
	}
	return false
}