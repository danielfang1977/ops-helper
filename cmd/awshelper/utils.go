package awshelper

import (
	"github.com/aws/aws-sdk-go/aws/session"
)

func getSession(profile string) (*session.Session, error) {
	if profile == "" {
		profile = "default"
	}

	sess, err := session.NewSessionWithOptions(session.Options{
		Profile:           profile,
		SharedConfigState: session.SharedConfigEnable,
	})
	if err != nil {
		return nil, err
	}
	return sess, nil
}
