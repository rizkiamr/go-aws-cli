package iam

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/iam"
	"github.com/aws/smithy-go/logging"
	"github.com/rizkiamr/go-aws-cli/internal/model"
	"github.com/spf13/cobra"
)

var isDebugModeEnabled bool
var policySourceARN string
var actionNames []string

// simulatePrincipalPolicyCmd represents the simulatePrincipalPolicy command
var simulatePrincipalPolicyCmd = &cobra.Command{
	Use:   "simulate-principal-policy",
	Short: "Simulate how a set of IAM policies attached to an IAM entity works with a list of API operations and Amazon Web Services resources to determine the policies' effective permissions.",
	Long: `Simulate how a set of IAM policies attached to an IAM entity works with a list of API operations and Amazon Web Services resources to determine the policies' effective permissions. The entity can be an IAM user, group, or role. If you specify a user, then the simulation also includes all of the policies that are attached to groups that the user belongs to. You can simulate resources that don't exist in your account.
You can optionally include a list of one or more additional policies specified as strings to include in the simulation. If you want to simulate only policies specified as strings, use SimulateCustomPolicy instead.

You can also optionally include one resource-based policy to be evaluated with each of the resources included in the simulation for IAM users only.

The simulation does not perform the API operations; it only checks the authorization to determine if the simulated policies allow or deny the operations.

Note: This operation discloses information about the permissions granted to other users. If you do not want users to see other user's permissions, then consider allowing them to use SimulateCustomPolicy instead.
Context keys are variables maintained by Amazon Web Services and its services that provide details about the context of an API query request. You can use the Condition element of an IAM policy to evaluate context keys. To get the list of context keys that the policies require for correct simulation, use GetContextKeysForPrincipalPolicy .

If the output is long, you can use the MaxItems and Marker parameters to paginate the results.

Note: The IAM policy simulator evaluates statements in the identity-based policy and the inputs that you provide during simulation. The policy simulator results can differ from your live Amazon Web Services environment. We recommend that you check your policies against your live Amazon Web Services environment after testing using the policy simulator to confirm that you have the desired results. For more information about using the policy simulator, see Testing IAM policies with the IAM policy simulator (https://docs.aws.amazon.com/IAM/latest/UserGuide/access_policies_testing-policies.html) in the IAM User Guide.

See also: AWS API Documentation

simulate-principal-policy is a paginated operation. Multiple API calls may be issued in order to retrieve the entire data set of results. You can disable pagination by providing the --no-paginate argument. When using --output text and the --query argument on a paginated response, the --query argument must extract data from the results of the following query expressions: EvaluationResults
`,
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

		svc := iam.NewFromConfig(cfg)

		// call `aws iam simulate-principal-policy` equivalent method, and parse the result
		// todo: test this
		awsResp, err := svc.SimulatePrincipalPolicy(context.TODO(), &iam.SimulatePrincipalPolicyInput{
			PolicySourceArn: &policySourceARN,
			ActionNames:     actionNames,
		})
		if err != nil {
			log.Fatalf("failed to call simulate-principal-policy: %v", err)
		}

		cliResp := &model.SimulatePrincipalPolicyResponse{
			EvaluationResults: awsResp.EvaluationResults,
		}

		jsonData, err := json.Marshal(cliResp)
		if err != nil {
			log.Fatalf("failed to marshal json: %v", err)
		}

		fmt.Printf("%s", string(jsonData))
	},
}

func init() {
	simulatePrincipalPolicyCmd.Flags().BoolVarP(&isDebugModeEnabled, "debug", "", false, "Enable debug mode")
	simulatePrincipalPolicyCmd.Flags().StringVar(&policySourceARN, "policy-source-arn", "", "The Amazon Resource Name (ARN) of a user, group, or role whose policies you want to include in the simulation. If you specify a user, group, or role, the simulation includes all policies that are associated with that entity. If you specify a user, the simulation also includes all policies that are attached to any groups the user belongs to.")
	simulatePrincipalPolicyCmd.Flags().StringArrayVar(&actionNames, "action-names", actionNames, "A list of names of API operations to evaluate in the simulation. Each operation is evaluated for each resource. Each operation must include the service identifier, such as iam:CreateUser.")
	simulatePrincipalPolicyCmd.MarkFlagsOneRequired("policy-source-arn", "action-names")
	simulatePrincipalPolicyCmd.MarkFlagsRequiredTogether("policy-source-arn", "action-names")
}
