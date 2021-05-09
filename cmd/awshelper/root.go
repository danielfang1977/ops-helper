package awshelper

import (
	"github.com/spf13/cobra"
)

var (
	inputProfile          string
	inputInstanceID       string
	inputAutoscalingGroup string
	inputMinSize          int64
	inputMaxSize          int64
	inputDesiredCapacity  int64
)

// NewCmd return newCommand
func NewCmd() *cobra.Command {
	command := &cobra.Command{
		Use:   "aws",
		Short: "aws helper function",
	}

	command.PersistentFlags().StringVarP(&inputProfile, "profile", "p", "", "aws profiles")

	command.AddCommand(newAllowTrafficCmd())
	command.AddCommand(newBlockTrafficCmd())
	command.AddCommand(newInstanceRefreshCmd())
	command.AddCommand(newUpdateAsgCmd())
	command.AddCommand(getEc2Normally())
	return command
}
