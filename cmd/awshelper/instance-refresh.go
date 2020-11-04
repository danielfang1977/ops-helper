package awshelper

import (
	"log"
	"ops-helper/components/autoscaling"

	"github.com/spf13/cobra"
)

func newInstanceRefreshCmd() *cobra.Command {
	command := &cobra.Command{
		Use:   "instance-refresh",
		Short: "refresh instance by autoscaling group",
		Long:  `example: ops-helper aws instance-refresh -p dev -n rest-api`,
		Run: func(cmd *cobra.Command, args []string) {
			if inputAutoscalingGroup == "" {
				log.Fatalln("autoscaling group name is required")
			}
			sess, err := getSession(inputProfile)
			if err != nil {
				log.Fatalln("Make aws session error", err)
			}

			svc, err := autoscaling.New(sess)
			if err != nil {
				log.Fatalln("Make autoscaling service error", err)
			}

			id, err := svc.StartInstanceRefresh(inputAutoscalingGroup)
			if err != nil {
				log.Fatalln(err)
			}

			log.Printf("Refresh id[%s]", id)
		},
	}

	command.Flags().StringVarP(&inputAutoscalingGroup, "name", "n", "", "autoscaling group name(required)")
	command.MarkFlagRequired("name")
	return command
}
