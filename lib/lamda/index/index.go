package index

import (
	"cdkgo/lib/asumerole"
	"cdkgo/lib/environment"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdkapigatewayv2integrationsalpha/v2"
	"github.com/aws/aws-cdk-go/awscdklambdagoalpha/v2"
	"github.com/aws/jsii-runtime-go"
)

// GetIndexLamda function lamda Index
func GetIndexLamda(stack awscdk.Stack, lamdaName *string, defaultEnv *map[string]*string) awscdkapigatewayv2integrationsalpha.HttpLambdaIntegration {

	lamdaEnv := &map[string]*string{
		"ENV": jsii.String("ENV"),
		"REG": jsii.String("Reg"),
	}

	mergeEnv := environment.MergeEnv(lamdaEnv, defaultEnv)

	lamdaRole := asumerole.UserRole(stack)

	// Define the Lambda function
	lambdaIndexFunction := awscdklambdagoalpha.NewGoFunction(stack, lamdaName, &awscdklambdagoalpha.GoFunctionProps{
		FunctionName: jsii.String("Index"),
		Entry:        jsii.String("lamda-go/domain/app"),
		Environment:  mergeEnv,
		Role:         lamdaRole,
	})

	// Create an integration between the API and the Lambda function
	apiIntegrationIndex := awscdkapigatewayv2integrationsalpha.NewHttpLambdaIntegration(
		lamdaName,
		lambdaIndexFunction,
		&awscdkapigatewayv2integrationsalpha.HttpLambdaIntegrationProps{},
	)

	return apiIntegrationIndex
}
