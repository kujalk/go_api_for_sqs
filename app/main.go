package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/secretsmanager"
	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/gorilla/mux"
)

var (
	sqsClient *sqs.SQS
	token     string
	queueURL  string
)

func fromJSON(jsonStr string) (map[string]string, error) {
	var data map[string]string
	err := json.Unmarshal([]byte(jsonStr), &data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func init() {
	sess := session.Must(session.NewSession())

	// Fetch the token from AWS Secret Manager
	smClient := secretsmanager.New(sess)
	input := &secretsmanager.GetSecretValueInput{
		SecretId: aws.String("api_token"),
	}
	result, err := smClient.GetSecretValue(input)
	if err != nil {
		log.Fatal(err)
	}

	jsonStr := *result.SecretString
	dict, err := fromJSON(jsonStr)

	if err != nil {
		fmt.Println(err)
		return
	}

	token = dict["token"]

	// Initialize the SQS client
	sqsClient = sqs.New(sess)

	// Get the queue URL from the environment variable
	queueURL = os.Getenv("SQS_QUEUE_URL")
	if queueURL == "" {
		log.Fatal("SQS_QUEUE_URL environment variable is not set")
	}
}
func main() {
	router := mux.NewRouter()
	fmt.Println("Server started on port 8080")

	// API endpoint for reading messages from the SQS queue
	router.HandleFunc("/read", func(w http.ResponseWriter, r *http.Request) {
		// Validate the token in the headers
		if r.Header.Get("Authorization") != "Bearer "+token {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		// Read messages from the SQS queue
		input := &sqs.ReceiveMessageInput{
			QueueUrl:            aws.String(queueURL),
			MaxNumberOfMessages: aws.Int64(10),
			VisibilityTimeout:   aws.Int64(10),
			WaitTimeSeconds:     aws.Int64(10),
		}
		result, err := sqsClient.ReceiveMessage(input)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		// Write the messages to the response
		for _, message := range result.Messages {
			fmt.Fprintln(w, *message.Body)
		}
	}).Methods("GET")

	// API endpoint for sending messages to the SQS queue
	router.HandleFunc("/send", func(w http.ResponseWriter, r *http.Request) {
		// Validate the token in the headers

		if r.Header.Get("Authorization") != "Bearer "+token {

			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		// Read the JSON request body
		var body struct {
			Message string `json:"message"`
		}

		if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		if body.Message == "" {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		message := body.Message

		// Send the message to the SQS queue
		input := &sqs.SendMessageInput{
			QueueUrl:    aws.String(queueURL),
			MessageBody: aws.String(message),
		}
		_, err := sqsClient.SendMessage(input)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
	}).Methods("POST")

	// Start the server
	log.Fatal(http.ListenAndServe(":8080", router))
}
