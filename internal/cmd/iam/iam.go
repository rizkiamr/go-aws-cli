package iam

import (
	"github.com/spf13/cobra"
)

// IAMCmd represents the iam command
var IAMCmd = &cobra.Command{
	Use:   "iam",
	Short: "Operations about AWS Identity and Access Management (IAM)",
	Long: `Identity and Access Management (IAM) is a web service for securely controlling access to Amazon Web Services services.
With IAM, you can centrally manage users, security credentials such as access keys, and permissions that control which Amazon Web Services resources users and applications can access.
For more information about IAM, see Identity and Access Management (IAM) (http://aws.amazon.com/iam/) and the Identity and Access Management User Guide (https://docs.aws.amazon.com/IAM/latest/UserGuide/).`,
}

func init() {
	IAMCmd.AddCommand(simulatePrincipalPolicyCmd)
}
