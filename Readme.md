# Go API to send/receive message from AWS SQS
This app is developed in golang with 2 API endpoints. For sending(post) and receiving(get) messages from AWS SQS. API endpoint is protected using a Bearer token

## Architecture

## Features
* API Token is stored in AWS Secret
* AWS SQS url is passed as env variable from AWS SSM parameter store
* API app is hosted in AWS Fargate
* Terraform based deployment via Github Action

## Running in AWS Fargate (Build using GitAction)
1. Create a S3 bucket for Terraform backend (provider.tf)
```terraform
terraform {
  backend "s3" {
    bucket = "gitaction-go-api-sqs"
    key    = "gitaction"
    region  = "ap-southeast-1"
  }
}
```

2. Create Programmatic IAM user and store it's credential as Github secrets
    * AWS_ACCESS_KEY_ID
    * AWS_SECRET_ACCESS_KEY

3. Create a ECR repo and copy the url
    * This url needs to be updated in the github/workflow/build_deploy.yml (#40 to 42) and Terraform.tfvars

4. Fill the terraform.tfvars and update the API token name in app/main.go (#38)

## Running Locally 
* Create a environment variable file as env.list
* Run the below command
``` bash
docker run -d --name my-go-app-container -p 9020:8080 --env-file env.list my-go-app
```

## Blog 
https://scripting4ever.wordpress.com/2022/12/31/go-api-to-send-and-receive-message-from-aws-sqs/

## Developer
K.Janarthanan