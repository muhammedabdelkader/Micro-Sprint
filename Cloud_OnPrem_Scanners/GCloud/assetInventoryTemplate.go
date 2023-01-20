package main

import (
    "context"
    "fmt"

    "google.golang.org/api/compute/v1"
    "google.golang.org/api/iam/v1"
    "google.golang.org/api/storage/v1"
    "google.golang.org/api/option"
)
/*
Uses the Google Cloud Go client library to create new instances of the `compute`, `iam`, and `storage` services, which you can then use to retrieve information about your GCP resources, such as Compute Engine instances, IAM users, and Storage Buckets. 
It then prints out the information obtained for each resource type.
*/

func main() {
    // create a context
    ctx := context.Background()

    // Create a new Compute client
    computeService, err := compute.NewService(ctx, option.WithCredentialsFile("path/to/credentials.json"))
    if err != nil {
        fmt.Printf("Failed to create compute client: %v", err)
        return
    }

    // Get all Compute Engine instances
    instances, err := computeService.Instances.List("YOUR_PROJECT_ID", "YOUR_ZONE").Do()
    if err != nil {
        fmt.Printf("Failed to list instances: %v", err)
        return
    }

    // Print information about Compute Engine instances
    fmt.Println("Compute Engine Instances:")
    for _, instance := range instances.Items {
        fmt.Printf("- Instance name: %s, CPU Platform: %s, Status: %s\n",
            instance.Name, instance.CpuPlatform, instance.Status)
    }

    // Create a new IAM client
    iamService, err := iam.NewService(ctx, option.WithCredentialsFile("path/to/credentials.json"))
    if err != nil {
        fmt.Printf("Failed to create IAM client: %v", err)
        return
    }

    // Get all IAM users
    users, err := iamService.Projects.ServiceAccounts.List("YOUR_PROJECT_ID").Do()
    if err != nil {
        fmt.Printf("Failed to list IAM users: %v", err)
        return
    }

    // Print information about IAM users
    fmt.Println("IAM Users:")
    for _, user := range users.Accounts {
        fmt.Printf("- User name: %s, Email: %s\n", user.Name, user.Email)
    }

    // Create a new Storage client
    storageService, err := storage.NewService(ctx, option.WithCredentialsFile("path/to/credentials.json"))
    if err != nil {
        fmt.Printf("Failed to create storage client: %v", err)
        return
    }

    // Get all Storage Buckets
    buckets, err := storageService.Buckets.List("YOUR_PROJECT_ID").Do()
    if err != nil {
        fmt.Printf("Failed to list Storage Buckets: %v", err)
        return
    }

    // Print information about Storage Buckets
    fmt.Println("Storage Buckets:")
    for _, bucket := range buckets.Items {
        fmt.Printf("- Bucket name: %s, Location: %s, Time Created: %s\n",
        bucket.Name, bucket.Location, bucket.TimeCreated)
}
}