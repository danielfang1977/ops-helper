package autoscaling

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/autoscaling"
	"github.com/aws/aws-sdk-go/service/autoscaling/autoscalingiface"
)

// Client struct
type Client struct {
	svc autoscalingiface.AutoScalingAPI
}

// New getClient
func New(sess *session.Session) (*Client, error) {
	svc := &Client{
		autoscaling.New(sess),
	}
	return svc, nil
}
