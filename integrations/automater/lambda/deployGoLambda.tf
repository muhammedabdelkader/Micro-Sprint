/*
 This Terraform configuration does the following:

 1. It declares the AWS provider and sets the region to "us-west-2".
 2. It creates an AWS Lambda function resource called "example", specifying the filename of the deployment package (main.zip), the function name, the IAM role that the function will assume, the handler, the runtime, and the source code hash of the deployment package.
 3. It creates an IAM role resource called "example", with an assume role policy that allows the Lambda service to assume this role.
 4. It creates an IAM role policy resource called "example" that allows the function to write logs to CloudWatch.
 */
provider "aws" {
    region = "us-west-2"
    
}

resource "aws_lambda_function" "example" {
    filename         = "main.zip"
      function_name    = "example"
        role             = aws_iam_role.example.arn
          handler          = "main"
            runtime          = "go1.x"
              source_code_hash = filebase64sha256("main.zip")
              
}

resource "aws_iam_role" "example" {
    name = "example"

      assume_role_policy = <<EOF
      {
          "Version": "2012-10-17",
          "Statement": [
            {
                    "Effect": "Allow",
                    "Principal": {
                              "Service": "lambda.amazonaws.com"
                                    
                    },
                          "Action": "sts:AssumeRole"
                              
            }
              
          ]
          
      }
      EOF
      
}

resource "aws_iam_role_policy" "example" {
    name = "example"
      role = aws_iam_role.example.id

        policy = <<EOF
        {
            "Version": "2012-10-17",
            "Statement": [
              {
                      "Effect": "Allow",
                      "Action": [
                                "logs:*"
                                      
                      ],
                            "Resource": "arn:aws:logs:*:*:*"
                                
              }
                
            ]
            
        }
        EOF
        
}

