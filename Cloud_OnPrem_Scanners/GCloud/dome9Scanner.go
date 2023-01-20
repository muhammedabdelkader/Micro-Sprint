package main

import (
    "context"
    "fmt"
    "log"

    cloudsecurity "cloud.google.com/go/securitycenter/apiv1"
    securitycenter "google.golang.org/genproto/googleapis/cloud/securitycenter/v1"
)
/*
This code uses the Google Cloud SDK for Go to create a new instance of the securitycenter client, 
which you can then use to scan your GCP resources using Dome9 (Check Point CloudGuard Dome9) rules. 
It specifies the Dome9 (Check Point CloudGuard Dome9) rule set ID you want to use, and then performs a scan on all resources 
in your GCP account. The number of findings imported will be printed out after the scan is finished.
*/
func main() {
    // Create a context
    ctx := context.Background()

    // Create a new instance of the Security Center client
    client, err := cloudsecurity.NewClient(ctx)
    if err != nil {
        log.Fatalf("Failed to create client: %v", err)
    }

    // Define the Dome9 (Check Point CloudGuard Dome9) rule set ID you want to use
    ruleSetId := "dome9-rule-set-id"

    // Define the list of all GCP services to scan
    services := []string{"all"}

    // Scan all resources in your GCP account
    req := &securitycenter.RunAssetDiscoveryRequest{
        Parent: "organizations/YOUR_ORGANIZATION_ID",
        RunAssetDiscoveryRequest: &securitycenter.RunAssetDiscoveryRequest_InlineSource{
            InlineSource: &securitycenter.InlineSource{
                RulesetName: ruleSetId,
                Services:    services,
            },
        },
    }

    resp, err := client.RunAssetDiscovery(ctx, req)
    if err != nil {
        log.Fatalf("Failed to run asset discovery: %v", err)
    }

    // Print the number of findings that were imported
    fmt.Println("Number of findings imported:", resp.GetTotalSize())
}
