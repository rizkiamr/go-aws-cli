package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/sts"
	"github.com/aws/smithy-go/logging"
	"github.com/rizkiamr/go-aws-cli/internal/model"
)

func main() {
	// todo: detect and parse Web Identity Token / IRSA

	// using the SDK's default configuration, load additional config
	// and credentials values from the environment variables, shared
	// credentials, and shared configuration files

	debugMode := flag.Bool("debug", false, "Enable debug mode")
	flag.Parse()

	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithDefaultRegion("us-east-1"))
	if err != nil {
		log.Fatalf("unable to load SDK config, %v", err)
	}

	if *debugMode {
		cfg.Logger = logging.NewStandardLogger(os.Stdout) // Sends logs to stdout
		cfg.ClientLogMode = aws.LogRequestWithBody | aws.LogResponseWithBody
	}

	svc := sts.NewFromConfig(cfg)

	// call `aws sts get-caller-identity` equivalent method, and parse the result
	awsResp, err := svc.GetCallerIdentity(context.TODO(), &sts.GetCallerIdentityInput{})
	if err != nil {
		log.Fatalf("failed to call get-caller-identity: %v", err)
	}

	cliResp := &model.GetCallerIdentityResponse{
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
