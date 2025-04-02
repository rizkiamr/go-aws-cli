package cmd

import (
	"os"

	"github.com/rizkiamr/go-aws-cli/internal/cmd/sts"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "go-aws-cli",
	Short: "go-aws-cli - aws-cli clone in Go",
	Long: `aws-cli clone in Go.
THIS IS A TOY PROJECT, DO NOT USE IN PRODUCTION`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(sts.STSCmd)
}