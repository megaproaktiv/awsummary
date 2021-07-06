package awsummary

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/elb"
	"log"
)

func ListElb(region string, verbose bool) (elbNumber, elbWithoutEC2 int) {
	svc := elb.New(session.New(&aws.Config{Region: aws.String(region)}))

	params := &elb.DescribeLoadBalancersInput{
		PageSize: aws.Int64(100),
	}

	resp, err := svc.DescribeLoadBalancers(params)
	if err != nil {
		log.Fatal("Cannot get ELB data: ", err)
	}

	for _, name := range resp.LoadBalancerDescriptions {
		var instances string
		if len(name.Instances) == 0 {
			elbWithoutEC2++
		}
		if verbose {
			log.Println("ELB Name: ", *name.LoadBalancerName)
			if len(name.Instances) == 0 {
				log.Println("--> NO Instances are associated with this ELB")
			} else {
				for _, id := range name.Instances {
					instances += *id.InstanceId
					instances += " "
				}
				log.Printf("-->ELB Instances: %s\n", instances)
			}
		}
	}
	elbNumber = len(resp.LoadBalancerDescriptions)
	return elbNumber, elbWithoutEC2
}
