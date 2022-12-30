provider "aws" {
  region  = "ap-southeast-1"
}

terraform {
  backend "s3" {
    bucket = "gitaction-go-api-sqs"
    key    = "gitaction"
    region  = "ap-southeast-1"
  }
}