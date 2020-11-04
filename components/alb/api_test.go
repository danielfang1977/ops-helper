package alb

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/aws/aws-sdk-go/service/elbv2"
	"github.com/aws/aws-sdk-go/service/elbv2/elbv2iface"
)

type mockedReceiveMsgs struct {
	elbv2iface.ELBV2API
	DeregisterResp                 elbv2.DeregisterTargetsOutput
	DeregisterErr                  error
	WaitUntilTargetDeregisteredErr error
	RegisterResp                   elbv2.RegisterTargetsOutput
	RegisterErr                    error
	WaitUntilTargetInServiceErr    error
}

func (m mockedReceiveMsgs) DeregisterTargets(in *elbv2.DeregisterTargetsInput) (*elbv2.DeregisterTargetsOutput, error) {
	return &m.DeregisterResp, m.DeregisterErr
}

func (m mockedReceiveMsgs) WaitUntilTargetDeregistered(in *elbv2.DescribeTargetHealthInput) error {
	return m.WaitUntilTargetDeregisteredErr
}

func (m mockedReceiveMsgs) RegisterTargets(in *elbv2.RegisterTargetsInput) (*elbv2.RegisterTargetsOutput, error) {
	return &m.RegisterResp, m.RegisterErr
}

func (m mockedReceiveMsgs) WaitUntilTargetInService(in *elbv2.DescribeTargetHealthInput) error {
	return m.WaitUntilTargetInServiceErr
}

func TestDeregister(t *testing.T) {
	cases := []struct {
		name                           string
		instanceID                     string
		targetGroupArn                 string
		DeregisterTargetsResp          elbv2.DeregisterTargetsOutput
		DeregisterErr                  error
		WaitUntilTargetDeregisteredErr error
		Expected                       error
	}{
		{
			"happy test case",
			"instanceId-foo",
			"targetgroup-foo",
			elbv2.DeregisterTargetsOutput{},
			nil,
			nil,
			nil,
		},
		{
			"Deregister error test",
			"instanceId-foo",
			"targetgroup-foo",
			elbv2.DeregisterTargetsOutput{},
			fmt.Errorf("Deregister error"),
			nil,
			fmt.Errorf("Deregister error"),
		},
		{
			"WaitUntilTargetDeregistered error test",
			"instanceId-foo",
			"targetgroup-foo",
			elbv2.DeregisterTargetsOutput{},
			nil,
			fmt.Errorf("WaitUntilTargetDeregistered error"),
			fmt.Errorf("WaitUntilTargetDeregistered error"),
		},
	}
	for _, c := range cases {
		svc := Client{
			mockedReceiveMsgs{
				DeregisterResp:                 c.DeregisterTargetsResp,
				DeregisterErr:                  c.DeregisterErr,
				WaitUntilTargetDeregisteredErr: c.WaitUntilTargetDeregisteredErr,
			},
		}
		err := svc.Deregister(&c.instanceID, &c.targetGroupArn)
		if !reflect.DeepEqual(err, c.Expected) {
			t.Errorf("expected: %v, got: %v", c.Expected, err)
		}
	}
}

func TestRegister(t *testing.T) {
	cases := []struct {
		name                        string
		instanceID                  string
		targetGroupArn              string
		RegisterResp                elbv2.RegisterTargetsOutput
		RegisterErr                 error
		WaitUntilTargetInServiceErr error
		Expected                    error
	}{
		{
			"happy test",
			"instanceId-foo",
			"targetgroup-foo",
			elbv2.RegisterTargetsOutput{},
			nil,
			nil,
			nil,
		},
		{
			"Register error test",
			"instanceId-foo",
			"targetgroup-foo",
			elbv2.RegisterTargetsOutput{},
			fmt.Errorf("Register error"),
			nil,
			fmt.Errorf("Register error"),
		},
		{
			"WaitUntilTargetInService error test",
			"instanceId-foo",
			"targetgroup-foo",
			elbv2.RegisterTargetsOutput{},
			fmt.Errorf("WaitUntilTargetInService error"),
			nil,
			fmt.Errorf("WaitUntilTargetInService error"),
		},
	}
	for _, c := range cases {
		svc := Client{
			mockedReceiveMsgs{
				RegisterResp:                c.RegisterResp,
				RegisterErr:                 c.RegisterErr,
				WaitUntilTargetInServiceErr: c.WaitUntilTargetInServiceErr,
			},
		}
		err := svc.Register(&c.instanceID, &c.targetGroupArn)
		if !reflect.DeepEqual(err, c.Expected) {
			t.Errorf("expected: %v, got: %v", c.Expected, err)
		}
	}
}
