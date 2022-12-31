# Go API to send/receive message from AWS SQS
This app is developed in golang with 2 API endpoints. For sending(post) and receiving(get) messages from AWS SQS. API endpoint is protected using a Bearer token

## Architecture

## Features
* API Token is stored in AWS Secret
* AWS SQS url is passed as env variable from AWS SSM parameter store
* API app is hosted in AWS Fargate

## Running in AWS Fargate (Build using GitAction)
1. Create a S3 bucket for Terraform backend
2. Create Programmatic IAM user and store it as Github secrets
3. Create a ECR repo and copy the url
4. Fill the terraform.tfvars and update the API token name is app/main.go (#38)

## Running Locally 
* Create a environment variable file 
* Run the below command
``` bash
docker run -d --name my-go-app-container -p 9020:8080 --env-file env.list my-go-app
```

## Developer
K.Janarthanan