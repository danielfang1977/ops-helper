package alb

import (
	"github.com/aws/aws-sdk-go/service/elbv2"
)

// Deregister a instance from target group
func (c *Client) Deregister(instanceID *string, targetGroupArn *string) error {
	_, err := c.svc.DeregisterTargets(&elbv2.DeregisterTargetsInput{
		TargetGroupArn: targetGroupArn,
		Targets: []*elbv2.TargetDescription{
			{
				Id: instanceID,
			},
		},
	})
	if err != nil {
		return err
	}

	err = c.svc.WaitUntilTargetDeregistered(&elbv2.DescribeTargetHealthInput{
		TargetGroupArn: targetGroupArn,
		Targets: []*elbv2.TargetDescription{
			{
				Id: instanceID,
			},
		},
	})

	return err
}

// Register a instance to target group
func (c *Client) Register(instanceID *string, targetGroupArn *string) error {
	_, err := c.svc.RegisterTargets(&elbv2.RegisterTargetsInput{
		TargetGroupArn: targetGroupArn,
		Targets: []*elbv2.TargetDescription{
			{
				Id: instanceID,
			},
		},
	})
	if err != nil {
		return err
	}

	err = c.svc.WaitUntilTargetInService(&elbv2.DescribeTargetHealthInput{
		TargetGroupArn: targetGroupArn,
		Targets: []*elbv2.TargetDescription{
			{
				Id: instanceID,
			},
		},
	})

	return err
}
