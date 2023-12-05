package config

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/jsii-runtime-go"
)

// ProdProps config props Beta
func ProdProps() *StackProps {
	props := &StackProps{
		StackProps: awscdk.StackProps{
			Env: &awscdk.Environment{
				// Account: jsii.String("123456789012"),
				Region: jsii.String("ap-southeast-1"),
			},
		},
		Stage: jsii.String("Production"),
	}

	return props
}
