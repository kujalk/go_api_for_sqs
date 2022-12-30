
## Running in AWS Fargate (Build using GitAction)


## Running Locally 
* Create a environment variable file 
``` bash
docker run -d --name my-go-app-container -p 9020:8080 --env-file env.list my-go-app
```