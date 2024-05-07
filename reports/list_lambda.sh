#!/bin/bash
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

