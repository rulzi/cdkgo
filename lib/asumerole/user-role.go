package asumerole

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/jsii-runtime-go"
)

// UserRole config asume role user
func UserRole(stack awscdk.Stack) awsiam.Role {
	lambdaRole := awsiam.NewRole(stack, jsii.String("MyLambdaRole"), &awsiam.RoleProps{
		AssumedBy: awsiam.NewServicePrincipal(jsii.String("lambda.amazonaws.com"), nil),
	})

	return lambdaRole
}
