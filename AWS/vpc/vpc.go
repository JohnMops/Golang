package main

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
	_ "strconv"
)

var (
	regionSelect string = "us-east-1"
	cidrBlock    string = "172.16.0.0/20"
)

func main() {
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(regionSelect)},
	)

	if err != nil {
		fmt.Println("Error establishing session")
		return
	}

	svc := ec2.New(sess)

	createVpc, vpcerr := svc.CreateVpc(&ec2.CreateVpcInput{
		AmazonProvidedIpv6CidrBlock:     nil,
		CidrBlock:                       aws.String(cidrBlock),
		DryRun:                          nil,
		InstanceTenancy:                 nil,
		Ipv6CidrBlock:                   nil,
		Ipv6CidrBlockNetworkBorderGroup: nil,
		Ipv6Pool:                        nil,
		TagSpecifications: []*ec2.TagSpecification{
			{
				ResourceType: aws.String("vpc"),
				Tags: []*ec2.Tag{
					{
						Key:   aws.String("Name"),
						Value: aws.String("Golang"),
					},
				},
			},
		},
	})

	if vpcerr != nil {
		fmt.Println("Error creating VPC")
		return
	}
	fmt.Printf(*createVpc.Vpc.VpcId + "\n")

}
