package ec2

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ec2"
)

// GetTagValue return value of tag
func (c *Client) GetTagValue(instanceID string, tagName string) (string, error) {
	resp, err := c.svc.DescribeInstances(&ec2.DescribeInstancesInput{
		InstanceIds: []*string{
			aws.String(instanceID),
		},
	})
	if err != nil {
		return "", err
	}
	if len(resp.Reservations) == 0 {
		return "", err
	}
	for idx := range resp.Reservations {
		for _, inst := range resp.Reservations[idx].Instances {
			for _, tag := range inst.Tags {
				if *tag.Key == tagName {
					return *tag.Value, nil
				}
			}
		}
	}
	return "", nil
}
