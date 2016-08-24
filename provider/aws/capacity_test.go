package aws_test

import "github.com/convox/rack/api/awsutil"

var cycleCapacityDescribeContainerInstances = awsutil.Cycle{
	awsutil.Request{
		RequestURI: "/",
		Operation:  "AmazonEC2ContainerServiceV20141113.DescribeContainerInstances",
		Body: `{
			"cluster":"cluster-test",
			"containerInstances": [
				"arn:aws:ecs:us-east-1:901416387788:container-instance/0ac4bb1c-be98-4202-a9c1-03153e91c05e",
				"arn:aws:ecs:us-east-1:901416387788:container-instance/38a59629-6f5d-4d02-8733-fdb49500ae45",
				"arn:aws:ecs:us-east-1:901416387788:container-instance/e7c311ae-968f-4125-8886-f9b724860d4c"
			]
		}`},
	awsutil.Response{
		StatusCode: 200,
		Body: `{
			"containerInstances": [
			  {
					"agentConnected": true,
					"containerInstanceArn": "arn:aws:ecs:us-east-1:901416387788:container-instance/0ac4bb1c-be98-4202-a9c1-03153e91c05e",
					"ec2InstanceId": "i-4a5513f4",
					"pendingTasksCount": 0,
					"registeredResources": [
						{ "doubleValue":0.0, "integerValue":1024, "longValue":0, "name":"CPU", "type":"INTEGER" },
						{ "doubleValue":0.0, "integerValue":2004, "longValue":0, "name":"MEMORY", "type":"INTEGER" },
						{ "doubleValue":0.0, "integerValue":0, "longValue":0, "name":"PORTS", "stringSetValue":["22","2376","2375","51678"], "type":"STRINGSET"},
						{ "doubleValue":0.0, "integerValue":0, "longValue":0, "name":"PORTS_UDP", "stringSetValue":[], "type":"STRINGSET"}
					],
					"remainingResources": [
						{"doubleValue":0.0, "integerValue":1024, "longValue":0, "name":"CPU", "type":"INTEGER" },
						{"doubleValue":0.0, "integerValue":2004, "longValue":0, "name":"MEMORY", "type":"INTEGER"},
						{"doubleValue":0.0,"integerValue":0,"longValue":0,"name":"PORTS","stringSetValue":["22","2376","2375","51678"],"type":"STRINGSET"},
						{"doubleValue":0.0,"integerValue":0,"longValue":0,"name":"PORTS_UDP","stringSetValue":[],"type":"STRINGSET"}
					],
					"runningTasksCount":0,
					"status":"ACTIVE",
					"versionInfo":{"agentHash":"4ab1051","agentVersion":"1.4.0","dockerVersion":"DockerVersion: 1.7.1"}
				},
			  {
					"agentConnected": true,
					"containerInstanceArn": "arn:aws:ecs:us-east-1:901416387788:container-instance/38a59629-6f5d-4d02-8733-fdb49500ae45",
					"ec2InstanceId": "i-4a5513f4",
					"pendingTasksCount": 0,
					"registeredResources": [
						{ "doubleValue":0.0, "integerValue":1024, "longValue":0, "name":"CPU", "type":"INTEGER" },
						{ "doubleValue":0.0, "integerValue":2004, "longValue":0, "name":"MEMORY", "type":"INTEGER" },
						{ "doubleValue":0.0, "integerValue":0, "longValue":0, "name":"PORTS", "stringSetValue":["22","2376","2375","51678"], "type":"STRINGSET"},
						{ "doubleValue":0.0, "integerValue":0, "longValue":0, "name":"PORTS_UDP", "stringSetValue":[], "type":"STRINGSET"}
					],
					"remainingResources": [
						{"doubleValue":0.0, "integerValue":1024, "longValue":0, "name":"CPU", "type":"INTEGER" },
						{"doubleValue":0.0, "integerValue":2004, "longValue":0, "name":"MEMORY", "type":"INTEGER"},
						{"doubleValue":0.0,"integerValue":0,"longValue":0,"name":"PORTS","stringSetValue":["22","2376","2375","51678"],"type":"STRINGSET"},
						{"doubleValue":0.0,"integerValue":0,"longValue":0,"name":"PORTS_UDP","stringSetValue":[],"type":"STRINGSET"}
					],
					"runningTasksCount":0,
					"status":"ACTIVE",
					"versionInfo":{"agentHash":"4ab1051","agentVersion":"1.4.0","dockerVersion":"DockerVersion: 1.7.1"}
				},
			  {
					"agentConnected": true,
					"containerInstanceArn": "arn:aws:ecs:us-east-1:901416387788:container-instance/e7c311ae-968f-4125-8886-f9b724860d4c",
					"ec2InstanceId": "i-4a5513f4",
					"pendingTasksCount": 0,
					"registeredResources": [
						{ "doubleValue":0.0, "integerValue":1024, "longValue":0, "name":"CPU", "type":"INTEGER" },
						{ "doubleValue":0.0, "integerValue":2004, "longValue":0, "name":"MEMORY", "type":"INTEGER" },
						{ "doubleValue":0.0, "integerValue":0, "longValue":0, "name":"PORTS", "stringSetValue":["22","2376","2375","51678"], "type":"STRINGSET"},
						{ "doubleValue":0.0, "integerValue":0, "longValue":0, "name":"PORTS_UDP", "stringSetValue":[], "type":"STRINGSET"}
					],
					"remainingResources": [
						{"doubleValue":0.0, "integerValue":1024, "longValue":0, "name":"CPU", "type":"INTEGER" },
						{"doubleValue":0.0, "integerValue":2004, "longValue":0, "name":"MEMORY", "type":"INTEGER"},
						{"doubleValue":0.0,"integerValue":0,"longValue":0,"name":"PORTS","stringSetValue":["22","2376","2375","51678"],"type":"STRINGSET"},
						{"doubleValue":0.0,"integerValue":0,"longValue":0,"name":"PORTS_UDP","stringSetValue":[],"type":"STRINGSET"}
					],
					"runningTasksCount":0,
					"status":"ACTIVE",
					"versionInfo":{"agentHash":"4ab1051","agentVersion":"1.4.0","dockerVersion":"DockerVersion: 1.7.1"}
				}
			]
		}`,
	},
}

var cycleCapacityDescribeServices = awsutil.Cycle{
	Request: awsutil.Request{
		RequestURI: "/",
		Operation:  "AmazonEC2ContainerServiceV20141113.DescribeServices",
		Body:       `{"cluster":"cluster-test", "services":["arn:aws:ecs:us-west-2:901416387788:service/convox-test-myapp-staging-worker-SCELGCIYSKF"]}`,
	},
	Response: awsutil.Response{
		StatusCode: 200,
		Body: `{
			"services": [
				{
					"status": "ACTIVE",
					"taskDefinition": "arn:aws:ecs:us-west-2:901416387788:task-definition/convox-test-myapp-staging-worker:1",
					"pendingCount": 0,
					"loadBalancers": [
							{
									"containerName": "worker",
									"containerPort": 80,
									"loadBalancerName": "convox-test-myapp-staging"
							}
					],
					"roleArn": "arn:aws:iam::901416387788:role/convox-test-myapp-staging-ServiceRole-1HNRHXNKGNLT9",
					"desiredCount": 1,
					"serviceName": "convox-test-myapp-staging-worker-SCELGCIYSKF",
					"clusterArn": "cluster-test",
					"serviceArn": "arn:aws:ecs:us-west-2:901416387788:service/convox-test-myapp-staging-worker-SCELGCIYSKF",
					"deployments": [
						{
							"status": "ACTIVE",
							"pendingCount": 0,
							"createdAt": 1449511658.683,
							"desiredCount": 1,
							"taskDefinition": "arn:aws:ecs:us-east-1:901416387788:task-definition/convox-test-myapp-staging-worker:1",
							"updatedAt": 1449511869.412,
							"id": "ecs-svc/9223370587343117124",
							"runningCount": 1
						}
					],
					"events": [
						{
							"message": "(service convox-test-myapp-staging-worker-SCELGCIYSKF) has started 1 tasks: (task f120ddee-5aa5-434e-b765-30503080078b).",
							"id": "d84b8245-9653-453f-a449-27d7c7cfdc0a",
							"createdAt": 1449003339.092
						}
					],
					"runningCount": 1
				}
			],
			"failures": []
		}`,
	},
}

var cycleCapacityListContainerInstances = awsutil.Cycle{
	awsutil.Request{
		RequestURI: "/",
		Operation:  "AmazonEC2ContainerServiceV20141113.ListContainerInstances",
		Body:       `{"cluster":"cluster-test", "nextToken":""}`,
	},
	awsutil.Response{
		StatusCode: 200,
		Body: `{
			"containerInstanceArns":[
				"arn:aws:ecs:us-east-1:901416387788:container-instance/0ac4bb1c-be98-4202-a9c1-03153e91c05e",
				"arn:aws:ecs:us-east-1:901416387788:container-instance/38a59629-6f5d-4d02-8733-fdb49500ae45",
				"arn:aws:ecs:us-east-1:901416387788:container-instance/e7c311ae-968f-4125-8886-f9b724860d4c"
			]
		}`,
	},
}

var cycleCapacityListContainerInstancesBadCluster = awsutil.Cycle{
	awsutil.Request{
		RequestURI: "/",
		Operation:  "AmazonEC2ContainerServiceV20141113.ListContainerInstances",
		Body:       `{"cluster":"cluster-test","nextToken":""}`},
	awsutil.Response{
		StatusCode: 400,
		Body:       `{"__type":"ClusterNotFoundException","message":"Cluster not found."}`},
}

var cycleCapacityListServices = awsutil.Cycle{
	Request: awsutil.Request{
		RequestURI: "/",
		Operation:  "AmazonEC2ContainerServiceV20141113.ListServices",
		Body:       `{"cluster":"cluster-test"}`,
	},
	Response: awsutil.Response{
		StatusCode: 200,
		Body:       `{"serviceArns":["arn:aws:ecs:us-west-2:901416387788:service/convox-test-myapp-staging-worker-SCELGCIYSKF"]}`,
	},
}
