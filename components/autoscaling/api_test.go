package autoscaling

import (
	"errors"
	"fmt"
	"reflect"
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/autoscaling"
	"github.com/aws/aws-sdk-go/service/autoscaling/autoscalingiface"
)

type mockedReceiveMsgs struct {
	autoscalingiface.AutoScalingAPI
	DescribeLoadBalancerTargetGroupsResp autoscaling.DescribeLoadBalancerTargetGroupsOutput
	DescribeLoadBalancerTargetGroupsErr  error
	StartInstanceRefreshResp             autoscaling.StartInstanceRefreshOutput
	StartInstanceRefreshErr              error
	UpdateAutoScalingGroupResp           autoscaling.UpdateAutoScalingGroupOutput
	UpdateAutoScalingGroupErr            error
}

func (m mockedReceiveMsgs) DescribeLoadBalancerTargetGroups(in *autoscaling.DescribeLoadBalancerTargetGroupsInput) (*autoscaling.DescribeLoadBalancerTargetGroupsOutput, error) {
	return &m.DescribeLoadBalancerTargetGroupsResp, m.DescribeLoadBalancerTargetGroupsErr
}

func (m mockedReceiveMsgs) StartInstanceRefresh(in *autoscaling.StartInstanceRefreshInput) (*autoscaling.StartInstanceRefreshOutput, error) {
	return &m.StartInstanceRefreshResp, m.StartInstanceRefreshErr
}

func (m mockedReceiveMsgs) UpdateAutoScalingGroup(in *autoscaling.UpdateAutoScalingGroupInput) (*autoscaling.UpdateAutoScalingGroupOutput, error) {
	return &m.UpdateAutoScalingGroupResp, m.UpdateAutoScalingGroupErr
}

func TestGetTargetGroups(t *testing.T) {
	cases := []struct {
		name                                 string
		autoscalingName                      string
		DescribeLoadBalancerTargetGroupsResp autoscaling.DescribeLoadBalancerTargetGroupsOutput
		DescribeLoadBalancerTargetGroupsErr  error
		expResp                              []targetGroup
		expErr                               error
	}{
		{
			"happy test case",
			"autoscaling-foo",
			autoscaling.DescribeLoadBalancerTargetGroupsOutput{},
			nil,
			[]targetGroup{},
			nil,
		},
	}

	for _, c := range cases {
		svc := Client{
			mockedReceiveMsgs{
				DescribeLoadBalancerTargetGroupsResp: c.DescribeLoadBalancerTargetGroupsResp,
				DescribeLoadBalancerTargetGroupsErr:  c.DescribeLoadBalancerTargetGroupsErr,
			},
		}
		resp, err := svc.GetTargetGroups(c.autoscalingName)
		if !reflect.DeepEqual(err, c.expErr) {
			t.Errorf("expected: %v, got: %v", c.expErr, err)
		}
		if !reflect.DeepEqual(resp, c.expResp) {
			t.Errorf("expected: %v, got: %v", c.expResp, err)
		}
	}
}

func TestStartInstanceRefresh(t *testing.T) {
	cases := []struct {
		name                     string
		autoscalingName          string
		StartInstanceRefreshResp autoscaling.StartInstanceRefreshOutput
		StartInstanceRefreshErr  error
		expResp                  string
		expErr                   error
	}{
		{
			"happy test case",
			"autoscaling-foo",
			autoscaling.StartInstanceRefreshOutput{
				InstanceRefreshId: aws.String("foo"),
			},
			nil,
			"foo",
			nil,
		},
		{
			"test error",
			"autoscaling-foo",
			autoscaling.StartInstanceRefreshOutput{},
			fmt.Errorf("foo"),
			"",
			fmt.Errorf("foo"),
		},
	}

	for _, c := range cases {
		svc := Client{
			mockedReceiveMsgs{
				StartInstanceRefreshResp: c.StartInstanceRefreshResp,
				StartInstanceRefreshErr:  c.StartInstanceRefreshErr,
			},
		}
		resp, err := svc.StartInstanceRefresh(c.autoscalingName)
		if !reflect.DeepEqual(err, c.expErr) {
			t.Errorf("expected: %v, got: %v", c.expErr, err)
		}
		if !reflect.DeepEqual(resp, c.expResp) {
			t.Errorf("expected: %v, got: %v", c.expResp, err)
		}
	}
}

func TestUpdateAutoscalingSize(t *testing.T) {
	cases := []struct {
		name                       string
		autoscalingName            string
		min                        int64
		max                        int64
		desired                    int64
		UpdateAutoScalingGroupResp autoscaling.UpdateAutoScalingGroupOutput
		UpdateAutoScalingGroupErr  error
		expErr                     error
	}{
		{
			"test all input 0",
			"autoscaling-foo",
			0,
			0,
			0,
			autoscaling.UpdateAutoScalingGroupOutput{},
			errors.New("Input fail"),
			fmt.Errorf("Input fail"),
		},
		{
			"happy test",
			"autoscaling-foo",
			1,
			1,
			1,
			autoscaling.UpdateAutoScalingGroupOutput{},
			nil,
			nil,
		},
	}

	for _, c := range cases {
		svc := Client{
			mockedReceiveMsgs{
				UpdateAutoScalingGroupResp: c.UpdateAutoScalingGroupResp,
				UpdateAutoScalingGroupErr:  c.UpdateAutoScalingGroupErr,
			},
		}
		err := svc.UpdateAutoscalingSize(c.autoscalingName, c.min, c.max, c.desired)
		if !reflect.DeepEqual(err, c.expErr) {
			t.Errorf("expected: %v, got: %v", c.expErr, err)
		}
	}
}
