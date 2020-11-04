package cmd

import (
	"fmt"
	"os"

	"ops-helper/cmd/awshelper"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use: "ops-helper",
}

// Execute command
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(newVersionCmd())
	rootCmd.AddCommand(awshelper.NewCmd())
}
