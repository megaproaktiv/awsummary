package sumec2

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/megaproaktiv/awsummary"
	"log"
)

var Client Ec2Interface

const DefaultRegion = "eu-central-1"


func init() {
	autoinit := awsummary.Autoinit()
	if autoinit {
		cfg, err := config.LoadDefaultConfig(context.TODO())
		if err != nil {
			panic("configuration error, " + err.Error())
		}

		Client = ec2.NewFromConfig(cfg)
	}
}

type InstancesTotals struct{
	Total int
	Running int
	Windows  int
}

//go:generate moq -out ec2_moq.go -pkg sumec2 . Ec2Interface

type Ec2Interface interface {
	DescribeInstances(ctx context.Context,
		params *ec2.DescribeInstancesInput, optFns ...func(*ec2.Options)) (*ec2.DescribeInstancesOutput,
		error)
}



func SetClient(c Ec2Interface) {
	Client = c
}

func List(region string, verbose bool) (InstancesTotals) {

	resp, err := Client.DescribeInstances(context.TODO(), nil,  
	func(o *ec2.Options) {
		o.Region = region
	} )
	if err != nil {
		log.Fatal("Cannot get EC2 data: ", err)
	}
	var totals = InstancesTotals{}
	// resp has all of the response data, pull out instance IDs:
	for idx := range resp.Reservations {
		for _, inst := range resp.Reservations[idx].Instances {
			totals.Total++
			if inst.State.Name == types.InstanceStateNameRunning {
				totals.Running++
				if inst.Platform == types.PlatformValuesWindows {
					totals.Windows++
				}
			}
			if verbose {
				log.Println("Instance ID: ", *inst.InstanceId)
				for _, tag := range inst.Tags {
					if *tag.Key == "Name" {
						log.Println("-->Instance Name: ", *tag.Value)
					}
				}
				log.Println("-->Instance State: ", inst.State.Name)
				if inst.Platform  == "" {
					log.Println("-->Instance Platform: Linux like")
				} else {
					log.Println("-->Instance Platform: ", inst.Platform)
				}
			}
		}
	}
	return totals
}
