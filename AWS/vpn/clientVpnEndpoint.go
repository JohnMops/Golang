package main

//as

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
)

var (
	region string = "us-east-1"
)

func main() {
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(region)},
	)

	svc := ec2.New(sess)

	runResult, errResult := svc.CreateClientVpnEndpoint(&ec2.CreateClientVpnEndpointInput{
		AuthenticationOptions: aws.[]("MutualAuthentication"),
		ClientCidrBlock:       aws.String("172.16.0.0/20"),
	})
}
