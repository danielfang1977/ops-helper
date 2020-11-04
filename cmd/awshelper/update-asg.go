package awshelper

import (
	"log"
	"ops-helper/components/autoscaling"

	"github.com/spf13/cobra"
)

func newUpdateAsgCmd() *cobra.Command {
	command := &cobra.Command{
		Use:   "update-asg",
		Short: "Update autoscaling group",
		Long: `Quick update autoscaling group size
example: aws update-asg -p dev -n rest-api --MaxSize 10 --MinSize 2 --DesiredCapacity 4`,
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

			err = svc.UpdateAutoscalingSize(inputAutoscalingGroup, inputMinSize, inputMaxSize, inputDesiredCapacity)
			if err != nil {
				log.Fatalln(err)
			}

			log.Printf("successfully")
		},
	}

	command.Flags().StringVarP(&inputAutoscalingGroup, "name", "n", "", "autoscaling group name(required)")
	command.Flags().Int64VarP(&inputMinSize, "MinSize", "", 0, "min size")
	command.Flags().Int64VarP(&inputMaxSize, "MaxSize", "", 0, "max size")
	command.Flags().Int64VarP(&inputDesiredCapacity, "DesiredCapacity", "", 0, "Desired Capacity")
	command.MarkFlagRequired("name")
	return command
}
