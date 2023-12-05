package routes

import (
	"cdkgo/lib/config"
	"cdkgo/lib/environment"
	"cdkgo/lib/lamda/index"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdkapigatewayv2alpha/v2"
	"github.com/aws/jsii-runtime-go"
)

// NewRouteGateway create Route Gateway
func NewRouteGateway(stack awscdk.Stack, apiName *string, props *config.StackProps) awscdkapigatewayv2alpha.HttpApi {

	defaultEnv := environment.DefaultEnv(props)

	apiIntegrationIndex := index.GetIndexLamda(stack, jsii.String("LamdaIndex"), defaultEnv)

	// Create an HTTP API
	api := awscdkapigatewayv2alpha.NewHttpApi(stack, apiName, &awscdkapigatewayv2alpha.HttpApiProps{
		ApiName: apiName,
	})

	// Add the Lambda integration to the default route
	api.AddRoutes(&awscdkapigatewayv2alpha.AddRoutesOptions{
		Path:        jsii.String("/"),
		Integration: apiIntegrationIndex,
		Methods:     &[]awscdkapigatewayv2alpha.HttpMethod{awscdkapigatewayv2alpha.HttpMethod_ANY},
	})

	APINameURL := *apiName + "URL"

	awscdk.NewCfnOutput(stack, jsii.String(APINameURL), &awscdk.CfnOutputProps{
		Value: api.Url(),
	})

	return api
}
