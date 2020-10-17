package main

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"

	"fmt"
	"log"
)

var (
	region       string = "us-west-2"
	instanceType string = "t2.micro"
	amiID        string = "ami-e7527ed7"
)

func main() {
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(region)},
	)

	// Create EC2 service client
	svc := ec2.New(sess)

	// Instance details
	runResult, errInstance := svc.RunInstances(&ec2.RunInstancesInput{
		ImageId:      aws.String(amiID),
		InstanceType: aws.String(instanceType),
		MinCount:     aws.Int64(1),
		MaxCount:     aws.Int64(1),
	})

	if errInstance != nil {
		fmt.Println("Error creating instance", err)
		return
	}

	fmt.Println("Created instance", *runResult.Instances[0].InstanceId)

	// Tag instance
	_, errtag := svc.CreateTags(&ec2.CreateTagsInput{
		Resources: []*string{runResult.Instances[0].InstanceId},
		Tags: []*ec2.Tag{
			{
				Key:   aws.String("Name"),
				Value: aws.String("MyFirstInstance"),
			},
		},
	})
	if errtag != nil {
		log.Println("Could not create tags for instance", runResult.Instances[0].InstanceId, errtag)
		return
	}

	fmt.Println("Successfully tagged instance")
}
