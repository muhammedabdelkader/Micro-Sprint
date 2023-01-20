package main

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/aws/aws-sdk-go/service/iam"
	"github.com/aws/aws-sdk-go/service/rds"
	"github.com/aws/aws-sdk-go/service/s3"
)
/*
This code uses the AWS SDK for Go to create new instances of the `ec2`, `iam`, `rds` and `s3` clients, 
which you can then use to retrieve information about your AWS resources, such as EC2 instances, IAM users, RDS DB instances, and S3 buckets.
*/
func main() {
	// Create an AWS session
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	// Create a new instance of the EC2 client
	ec2Client := ec2.New(sess)

	// Get all EC2 instances
	instances, err := ec2Client.DescribeInstances(nil)
	if err != nil {
		fmt.Println("Error getting EC2 instances:", err)
		return
	}

	// Print information about EC2 instances
	fmt.Println("EC2 Instances:")
	for _, reservation := range instances.Reservations {
		for _, instance := range reservation.Instances {
			fmt.Printf("- Instance ID: %s, Type: %s, Private IP: %s, Public IP: %s\n",
				*instance.InstanceId, *instance.InstanceType, *instance.PrivateIpAddress, *instance.PublicIpAddress)
		}
	}

	// Create a new instance of the IAM client
	iamClient := iam.New(sess)

	// Get all IAM users
	users, err := iamClient.ListUsers(nil)
	if err != nil {
		fmt.Println("Error getting IAM users:", err)
		return
	}

	// Print information about IAM users
	fmt.Println("IAM Users:")
	for _, user := range users.Users {
		fmt.Printf("- User name: %s, User ID: %s\n", *user.UserName, *user.UserId)
	}

	// Create a new instance of the RDS client
	rdsClient := rds.New(sess)

	// Get all RDS DB instances
	dbs, err := rdsClient.DescribeDBInstances(nil)
	if err != nil {
		fmt.Println("Error getting RDS DB instances:", err)
		return
	}
// Print information about RDS DB instances
fmt.Println("RDS DB Instances:")
for _, db := range dbs.DBInstances {
fmt.Printf("- DB Instance ID: %s, Engine: %s, Storage: %dGB\n",
*db.DBInstanceIdentifier, *db.Engine, *db.AllocatedStorage)
}

// Create a new instance of the S3 client
s3Client := s3.New(sess)

// Get all S3 buckets
buckets, err := s3Client.ListBuckets(&s3.ListBucketsInput{})
if err != nil {
	fmt.Println("Error getting S3 buckets:", err)
	return
}

// Print information about S3 buckets
fmt.Println("S3 Buckets:")
for _, bucket := range buckets.Buckets {
	fmt.Printf("- Bucket name: %s, Creation date: %s\n",
		*bucket.Name, bucket.CreationDate)
}
}