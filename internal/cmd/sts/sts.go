package sts

import (
	"github.com/spf13/cobra"
)

// STSCmd represents the sts command
var STSCmd = &cobra.Command{
	Use:   "sts",
	Short: "Operations about AWS Security Token Service",
	Long: `Security Token Service (STS) enables you to request temporary, limited-privilege credentials for users.
This guide provides descriptions of the STS API.
For more information about using this service, see Temporary Security Credentials (https://docs.aws.amazon.com/IAM/latest/UserGuide/id_credentials_temp.html)`,
}

func init() {
	STSCmd.AddCommand(getCallerIdentityCmd)
}
