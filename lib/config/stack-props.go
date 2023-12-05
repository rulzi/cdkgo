package config

import "github.com/aws/aws-cdk-go/awscdk/v2"

// StackProps create struct StackProps
type StackProps struct {
	StackProps    awscdk.StackProps
	Stage         *string
	ZoneName      *string
	HostedZoneID  *string
	DomainName    *string
	subDomainName *string
	networkCidr   *string
}
