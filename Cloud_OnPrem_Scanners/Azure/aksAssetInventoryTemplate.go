package main

import (
    "context"
    "fmt"
    "github.com/Azure/azure-sdk-for-go/services/containerservice/mgmt/2020-09-01/containerservice"
    "github.com/Azure/go-autorest/autorest"
    "github.com/Azure/go-autorest/autorest/azure/auth"
)
/*
Uses the Azure Go SDK to create a new instance of the `containerservice` client, which you can then use to retrieve information about your Azure Kubernetes Service (AKS) resources, such as AKS Clusters and Nodes. 
It then prints out the information obtained for each resource type.
*/
func main() {
    // create an authorizer
    authorizer, err := auth.NewAuthorizerFromEnvironment()
    if err != nil {
        fmt.Printf("Failed to get an Authorizer: %v", err)
        return
    }

    // Create a new AKS client
    aksClient := containerservice.NewManagedClustersClient("YOUR_SUBSCRIPTION_ID")
    aksClient.Authorizer = authorizer

    // Get all AKS clusters
    result, err := aksClient.List(context.Background())
    if err != nil {
        fmt.Printf("Failed to list AKS clusters: %v", err)
        return
    }

    // Print information about AKS clusters
    fmt.Println("AKS Clusters:")
    for _, cluster := range result.Values() {
        fmt.Printf("- AKS Cluster Name: %s, Location: %s, Kubernetes Version: %s, Agent Pool Count: %d\n",*cluster.Name, *cluster.Location, *cluster.KubernetesVersion, len(cluster.AgentPoolProfiles))
    }
    
    // Get all AKS nodes
for _, cluster := range result.Values() {
    nodes, err := aksClient.ListNodes(context.Background(), *cluster.ResourceGroup, *cluster.Name)
    if err != nil {
        fmt.Printf("Failed to list AKS nodes: %v", err)
        return
    }
    fmt.Println("AKS Nodes:")
    for _, node := range nodes.Values() {
        fmt.Printf("- AKS Node Name: %s, Agent Pool Name: %s, OS Type: %s, Pool Type: %s\n",
            *node.Name, *node.AgentPoolName, *node.OsType, *node.PoolType)
    }
}
}
