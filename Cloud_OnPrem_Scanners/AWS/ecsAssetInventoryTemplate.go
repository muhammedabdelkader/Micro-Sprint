package main

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ecs"
)
/*
Use the AWS SDK for Go to connect to your AWS account and use the ECS service to list all the clusters, tasks and services. This can be used to get an inventory of all the assets running in AWS Fargate.
*/

func main() {
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))
	svc := ecs.New(sess)

	// Perform an inventory of clusters
	clustersInput := &ecs.ListClustersInput{}
	clustersOutput, err := svc.ListClusters(clustersInput)
	if err != nil {
		fmt.Println("Error", err)
		os.Exit(1)
	}
	for _, cluster := range clustersOutput.ClusterArns {
		fmt.Println("Cluster:", *cluster)
	}

	// Perform an inventory of tasks
	tasksInput := &ecs.ListTasksInput{}
	tasksOutput, err := svc.ListTasks(tasksInput)
	if err != nil {
		fmt.Println("Error", err)
		os.Exit(1)
	}
	for _, task := range tasksOutput.TaskArns {
		fmt.Println("Task:", *task)
	}

	// Perform an inventory of services
	servicesInput := &ecs.ListServicesInput{}
	servicesOutput, err := svc.ListServices(servicesInput)
	if err != nil {
		fmt.Println("Error", err)
		os.Exit(1)
	}
	for _, service := range servicesOutput.ServiceArns {
		fmt.Println("Service:", *service)
	}
}
