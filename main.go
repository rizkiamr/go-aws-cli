package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/config"

	"github.com/aws/aws-sdk-go-v2/service/sts"
)

type GetCallerIdentityResponse struct {
	UserId  string `json:"UserId"`
	Account string `json:"Account"`
	Arn     string `json:"Arn"`
}

func main() {
	// todo: detect and parse Web Identity Token / IRSA

	// using the SDK's default configuration, load additional config
	// and credentials values from the environment variables, shared
	// credentials, and shared configuration files
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithDefaultRegion("us-east-1"))
	if err != nil {
		log.Fatalf("unable to load SDK config, %v", err)
	}

	svc := sts.NewFromConfig(cfg)

	// call `aws sts get-caller-identity` equivalent method, and parse the result
	awsResp, err := svc.GetCallerIdentity(context.TODO(), &sts.GetCallerIdentityInput{})
	if err != nil {
		log.Fatalf("failed to call get-caller-identity: %v", err)
	}

	cliResp := &GetCallerIdentityResponse{
		UserId:  *awsResp.UserId,
		Account: *awsResp.Account,
		Arn:     *awsResp.Arn,
	}

	jsonData, err := json.Marshal(cliResp)
	if err != nil {
		log.Fatalf("failed to marshal json: %v", err)
	}

	fmt.Printf("%s", string(jsonData))
}
