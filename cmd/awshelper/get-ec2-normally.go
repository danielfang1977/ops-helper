package awshelper

import (
	"log"
	"ops-helper/components/ec2"

	"github.com/spf13/cobra"
)

func getEc2Normally() *cobra.Command {
	command := &cobra.Command{
		Use:   "get-ec2",
		Short: "get ec2 data of normally query",
		Long:  `example: ops-helper aws get-ec2 -p dev`,
		Run: func(cmd *cobra.Command, args []string) {
			log.Println("Start get instance data...")

			sess, err := getSession(inputProfile)
			if err != nil {
				log.Fatalln("Make aws session error", err)
			}

			ec2Svc, err := ec2.New(sess)
			if err != nil {
				log.Fatalln("Make ec2 service error", err)
			}

			err = ec2Svc.GetEc2Normally()
			if err != nil {
				log.Fatalln("Query ec2 tag error", err)
			}
		},
	}
	return command
}
