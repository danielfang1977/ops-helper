package alb

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/elbv2"
	"github.com/aws/aws-sdk-go/service/elbv2/elbv2iface"
)

// Client struct
type Client struct {
	svc elbv2iface.ELBV2API
}

// New getClient
func New(sess *session.Session) (*Client, error) {
	svc := &Client{
		elbv2.New(sess),
	}
	return svc, nil
}
