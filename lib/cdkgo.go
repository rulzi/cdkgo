package lib

import (
	"cdkgo/lib/config"
	"cdkgo/lib/stack"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/jsii-runtime-go"
	"github.com/posener/goaction"
)

// RunStack Run Stack AWS
func RunStack() {
	defer jsii.Close()

	app := awscdk.NewApp(nil)

	if goaction.Ref == "refs/heads/master" {
		stack.RunStack(app, "CDKGoStackBeta", config.BetaProps())
	}

	if goaction.Ref == "refs/tags/release" {
		stack.RunStack(app, "CDKGoStackProd", config.ProdProps())
	}

	app.Synth(nil)
}
