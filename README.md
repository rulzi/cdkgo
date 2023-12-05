# Welcome to your CDK Go project!

This is a blank project for CDK development with Go.

The `cdk.json` file tells the CDK toolkit how to execute your app.

## Useful commands

 * `cdk deploy`      deploy this stack to your default AWS account/region
 * `cdk diff`        compare deployed stack with current state
 * `cdk synth`       emits the synthesized CloudFormation template
 * `go test`         run unit tests

## Folder Description

* `lambda-go`           lambda golang project
* `lib`                 cdk lib (stage, stack)
* `lib/asumerole`       aws asumerole
* `lib/config`          cdk props for beta & prod env
* `lib/environment`     setting default environment
* `lib/lamda`           funtion lamda
* `lib/routes`          routes lamda
* `lib/secretmanager`   aws secret manager
* `lib/ssm`             aws ssm parameter store
* `lib/stack`           aws create stack