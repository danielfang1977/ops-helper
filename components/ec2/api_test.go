package ec2

import (
	"reflect"
	"testing"

	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/aws/aws-sdk-go/service/ec2/ec2iface"
)

type mockedReceiveMsgs struct {
	ec2iface.EC2API
	resp ec2.DescribeInstancesOutput
	err  error
}

func (m mockedReceiveMsgs) DescribeInstances(in *ec2.DescribeInstancesInput) (*ec2.DescribeInstancesOutput, error) {
	return &m.resp, m.err
}

func TestGetTagValue(t *testing.T) {
	cases := []struct {
		name                    string
		instanceID              string
		tagName                 string
		DescribeInstancesOutput ec2.DescribeInstancesOutput
		DescribeInstancesErr    error
		expResp                 string
		expErr                  error
	}{
		{
			"happy test case",
			"instance-foo",
			"tagName-foo",
			ec2.DescribeInstancesOutput{},
			nil,
			"",
			nil,
		},
	}

	for _, c := range cases {
		svc := Client{
			mockedReceiveMsgs{
				resp: c.DescribeInstancesOutput,
				err:  c.DescribeInstancesErr,
			},
		}
		resp, err := svc.GetTagValue(c.instanceID, c.tagName)
		if !reflect.DeepEqual(err, c.expErr) {
			t.Errorf("expected: %v, got: %v", c.expErr, err)
		}
		if !reflect.DeepEqual(resp, c.expResp) {
			t.Errorf("expected: %v, got: %v", c.expResp, err)
		}
	}
}
