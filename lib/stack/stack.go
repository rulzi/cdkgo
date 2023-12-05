package stack

import (
	"cdkgo/lib/config"
	"cdkgo/lib/routes"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/jsii-runtime-go"

	"github.com/aws/constructs-go/constructs/v10"
)

// NewStack create New Stack
func NewStack(scope constructs.Construct, id string, props *config.StackProps) awscdk.Stack {
	var sprops awscdk.StackProps
	if props != nil {
		sprops = props.StackProps
	}
	stack := awscdk.NewStack(scope, &id, &sprops)

	routes.NewRouteGateway(stack, jsii.String("ApiGateway"), props)

	return stack
}

// RunStack Run Stack
func RunStack(app awscdk.App, StackName string, props *config.StackProps) {

	NewStack(app, StackName, props)
}
