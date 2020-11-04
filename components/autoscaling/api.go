package autoscaling

import (
	"errors"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/autoscaling"
)

type targetGroup struct {
	Arn *string
}

// GetTargetGroups get autoscaling target group
func (c *Client) GetTargetGroups(autoscalingName string) ([]targetGroup, error) {
	resp, err := c.svc.DescribeLoadBalancerTargetGroups(&autoscaling.DescribeLoadBalancerTargetGroupsInput{
		AutoScalingGroupName: aws.String(autoscalingName),
	})
	if err != nil {
		return nil, err
	}
	list := []targetGroup{}
	for _, tg := range resp.LoadBalancerTargetGroups {
		t := targetGroup{tg.LoadBalancerTargetGroupARN}
		list = append(list, t)
	}
	return list, nil
}

// StartInstanceRefresh refersh instance by autoscaling group
func (c *Client) StartInstanceRefresh(autoscalingName string) (string, error) {
	resp, err := c.svc.StartInstanceRefresh(&autoscaling.StartInstanceRefreshInput{
		AutoScalingGroupName: aws.String(autoscalingName),
	})
	if err != nil {
		return "", err
	}

	return *resp.InstanceRefreshId, nil
}

// UpdateAutoscalingSize change size
func (c *Client) UpdateAutoscalingSize(autoscalingName string, min int64, max int64, desired int64) error {
	if min < 1 && max < 1 && desired < 1 {
		return errors.New("Input fail")
	}

	in := &autoscaling.UpdateAutoScalingGroupInput{
		AutoScalingGroupName: aws.String(autoscalingName),
	}

	if min > 0 {
		in.MinSize = aws.Int64(min)
	}

	if max > 0 {
		in.MaxSize = aws.Int64(max)
	}

	if desired > 0 {
		in.DesiredCapacity = aws.Int64(desired)
	}

	_, err := c.svc.UpdateAutoScalingGroup(in)

	return err
}
