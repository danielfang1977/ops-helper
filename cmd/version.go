package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var version string
var githash string

func newVersionCmd() *cobra.Command {
	command := &cobra.Command{
		Use:   "version",
		Short: "Print the version number of Ops Helper",
		Long:  `All software has versions. This is Ops Helper's`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("Ops Helper \nversion: %s\ngithash: %s\n", version, githash)
		},
	}
	return command
}
