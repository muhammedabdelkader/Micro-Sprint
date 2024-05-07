#!/bin/bash
################################################################################################################################################
## Purpose: Migrating AWS Lambda functions from the Go1.x runtime to the custom runtime on Amazon Linux 2
## https://aws.amazon.com/blogs/compute/migrating-aws-lambda-functions-from-the-go1-x-runtime-to-the-custom-runtime-on-amazon-linux-2/
################################################################################################################################################
## You Need to have AWS credentials valid
## Get the list of AWS regions
regions=$(aws ec2 describe-regions --query "Regions[].RegionName" --output text)
accountId=$(aws sts get-caller-identity | jq '.Account')
# Loop through each region and list Lambda functions
base="aws/${accountId}"
cd "$HOME/Desktop"
mkdir -p $base
cd $base 
for region in $regions; do
    echo "Listing Lambda functions in region: $region"
    aws lambda list-functions --region $region > ".lambda_${accountId}_${region}"
        done
        
## Filter Example 
##
# for j in $(ls ); do cd $j; echo $j; for i in $(ls -ah "." | grep ".lambda" );do cat "${i}" | jq ".Functions[]| select (.Runtime==\"go1.x\") | .FunctionName,.FunctionArn";done;cd .. ;  done
##
