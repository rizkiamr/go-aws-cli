package sts

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/sts"
	"github.com/aws/smithy-go/logging"
	"github.com/rizkiamr/go-aws-cli/internal/model"
	"github.com/spf13/cobra"
)

var isDebugModeEnabled bool

// getCallerIdentityCmd represents the getCallerIdentity command
var getCallerIdentityCmd = &cobra.Command{
	Use:   "get-caller-identity",
	Short: "Returns details about the IAM user or role whose credentials are used to call the operation.",
	Long: `Returns details about the IAM user or role whose credentials are used to call the operation.
No permissions are required to perform this operation.
If an administrator attaches a policy to your identity that explicitly denies access to the sts:GetCallerIdentity action, you can still perform this operation.
Permissions are not required because the same information is returned when access is denied.
See also: AWS API Documentation (https://docs.aws.amazon.com/goto/WebAPI/sts-2011-06-15/GetCallerIdentity)`,
	Run: func(cmd *cobra.Command, args []string) {
		// todo: detect and parse Web Identity Token / IRSA

		cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithDefaultRegion("us-east-1"))
		if err != nil {
			log.Fatalf("unable to load SDK config, %v", err)
		}

		if isDebugModeEnabled {
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
	},
}

func init() {
	getCallerIdentityCmd.Flags().BoolVarP(&isDebugModeEnabled, "debug", "", false, "Enable debug mode")
}
