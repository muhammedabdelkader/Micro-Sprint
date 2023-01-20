package main

import (
    "fmt"
    "github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/aws/session"
    "github.com/aws/aws-sdk-go/service/securityhub"
)
/*
This code uses the AWS SDK for Go to create a new instance of the securityhub client. 
it specifies the Dome9 (Check Point CloudGuard Dome9) rule set ID you want to use, and then performs a scan on all resources in your AWS account. 
The number of findings imported will be printed out after the scan is finished.
*/
func main() {
    // Create an AWS session
    sess := session.Must(session.NewSessionWithOptions(session.Options{
        SharedConfigState: session.SharedConfigEnable,
    }))

    // Create a new instance of the Security Hub client
    sh := securityhub.New(sess)

    // Define the Dome9 (Check Point CloudGuard Dome9) rule set ID you want to use
    ruleSetId := "dome9-rule-set-id"

    // Define the list of all AWS resources to scan
    resources := []string{"all"}

    // Scan all resources in your AWS account
    result, err := sh.BatchImportFindings(&securityhub.BatchImportFindingsInput{
        Product: aws.String("Check Point CloudGuard Dome9"),
        ImportFindings: []*securityhub.ImportFinding{
            {
                ProductFields: map[string]*string{
                    "ruleSetId": aws.String(ruleSetId),
                    "resources": aws.String(strings.Join(resources, ",")),
                },
            },
        },
    })
    if err != nil {
        fmt.Println("Error scanning resources:", err)
        return
    }

    // Print the number of findings that were imported
    fmt.Println("Number of findings imported:", len(result.ProcessedFindings))
}
