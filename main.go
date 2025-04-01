package main

import "fmt"

func main() {
	// todo: parse aws creds from aws-default environment variables
	// - AWS_ACCESS_KEY_ID
	// - AWS_SECRET_ACCESS_KEY
	// - AWS_SESSION_TOKEN (optional)
	// - AWS_DEFAULT_REGION (optional, default: us-east-1)

	// good to have: detect Web Identity Token / IRSA 

	// todo: call `aws sts get-caller-identity` equivalent method, and parse the result
	fmt.Println("hello world")
}
