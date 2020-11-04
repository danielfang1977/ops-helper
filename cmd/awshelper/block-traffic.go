package awshelper

import (
	"log"
	"ops-helper/components/alb"
	"ops-helper/components/autoscaling"
	"ops-helper/components/ec2"
	"sync"

	"github.com/spf13/cobra"
)

func newBlockTrafficCmd() *cobra.Command {
	command := &cobra.Command{
		Use:   "block-traffic",
		Short: "deregister instance of target group",
		Long:  `example: ops-helper aws block-traffic -p dev -i instanceID`,
		Run: func(cmd *cobra.Command, args []string) {
			if inputInstanceID == "" {
				log.Fatalln("instanceID is required")
			}

			log.Println("Start deregister instance...")

			sess, err := getSession(inputProfile)
			if err != nil {
				log.Fatalln("Make aws session error", err)
			}

			ec2Svc, err := ec2.New(sess)
			if err != nil {
				log.Fatalln("Make ec2 service error", err)
			}

			autoscalingName, err := ec2Svc.GetTagValue(inputInstanceID, "aws:autoscaling:groupName")
			if err != nil {
				log.Fatalln("Query ec2 tag error", err)
			}

			autoscalingSvc, err := autoscaling.New(sess)
			if err != nil {
				log.Fatalln("Make autoscaling service error", err)
			}

			targetGroups, err := autoscalingSvc.GetTargetGroups(autoscalingName)
			if err != nil {
				log.Printf("Query autoscaling[%s] target groups error", autoscalingName)
				log.Fatal(err)
			}

			albSvc, err := alb.New(sess)
			if err != nil {
				log.Fatalln("Make alb service error", err)
			}

			size := len(targetGroups)
			if size > 0 {
				wg := new(sync.WaitGroup)
				wg.Add(size)
				for _, t := range targetGroups {
					go func(targetGroupArn *string) {
						defer wg.Done()
						log.Printf("start deregister instanceID[%s], arn[%s]\n", inputInstanceID, *targetGroupArn)
						err := albSvc.Deregister(&inputInstanceID, targetGroupArn)
						if err != nil {
							panic(err)
						}
						log.Printf("end deregister instanceID[%s], arn[%s]\n", inputInstanceID, *targetGroupArn)
					}(t.Arn)
				}
				wg.Wait()
			}
		},
	}

	command.Flags().StringVarP(&inputInstanceID, "instanceid", "i", "", "ec2 instance id(required)")
	command.MarkFlagRequired("instanceid")
	return command
}
