package config

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/jsii-runtime-go"
)

// BetaProps config props Beta
func BetaProps() *StackProps {
	props := &StackProps{
		StackProps: awscdk.StackProps{
			Env: &awscdk.Environment{
				// Account: jsii.String("123456789012"),
				Region: jsii.String("ap-southeast-1"),
			},
		},
		Stage: jsii.String("Beta"),
	}

	return props
}
