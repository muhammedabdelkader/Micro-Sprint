package main

import (
    "context"
    "fmt"

    "github.com/Azure/azure-sdk-for-go/services/compute/mgmt/2019-07-01/compute"
    "github.com/Azure/azure-sdk-for-go/services/resources/mgmt/2019-05-01/resources"
    "github.com/Azure/azure-sdk-for-go/services/storage/mgmt/2019-04-01/storage"
    "github.com/Azure/go-autorest/autorest"
    "github.com/Azure/go-autorest/autorest/azure/auth"
)
/*
Uses the Azure Go SDK to create new instances of the `compute`, `resources`, and `storage` clients, which you can then use to retrieve information about your Azure resources, such as Virtual Machines, Resource Groups, and Storage Accounts. It then prints out the information obtained for each resource type.
*/

func main() {
    // create an authorizer
    authorizer, err := auth.NewAuthorizerFromEnvironment()
    if err != nil {
        fmt.Printf("Failed to get an Authorizer: %v", err)
        return
    }

    // Create a new Compute client
    computeClient := compute.NewVirtualMachinesClient("YOUR_SUBSCRIPTION_ID")
    computeClient.Authorizer = authorizer

    // Get all Virtual Machines
    result, err := computeClient.List(context.Background())
    if err != nil {
        fmt.Printf("Failed to list Virtual Machines: %v", err)
        return
    }

    // Print information about Virtual Machines
    fmt.Println("Virtual Machines:")
    for _, vm := range result.Values() {
        fmt.Printf("- VM Name: %s, Size: %s, Power State: %s\n",
            *vm.Name, *vm.HardwareProfile.VMSize, *vm.InstanceView.Statuses[1].Code)
    }

    // Createa new Resource Management client
resourcesClient := resources.NewGroupsClient("YOUR_SUBSCRIPTION_ID")
resourcesClient.Authorizer = authorizer
// Get all resource groups
result, err := resourcesClient.List(context.Background())
if err != nil {
    fmt.Printf("Failed to list resource groups: %v", err)
    return
}

// Print information about resource groups
fmt.Println("Resource Groups:")
for _, group := range result.Values() {
    fmt.Printf("- Resource Group Name: %s, Location: %s\n",
        *group.Name, *group.Location)
}

// Create a new Storage client
storageClient := storage.NewAccountsClient("YOUR_SUBSCRIPTION_ID")
storageClient.Authorizer = authorizer

// Get all storage accounts
result, err = storageClient.List(context.Background())
if err != nil {
    fmt.Printf("Failed to list storage accounts: %v", err)
    return
}

// Print information about storage accounts
fmt.Println("Storage Accounts:")
for _, account := range result.Values() {
    fmt.Printf("- Storage Account Name: %s, Location: %s, Account Type: %s\n",
        *account.Name, *account.Location, string(account.AccountProperties.AccountType))
}
}

