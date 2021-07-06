package sumlb

import (
	"context"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	elbv2 "github.com/aws/aws-sdk-go-v2/service/elasticloadbalancingv2"
	"github.com/aws/aws-sdk-go-v2/service/elasticloadbalancingv2/types"
	"github.com/megaproaktiv/awsummary"
)

var Client Elbv2interface

func init() {
	autoinit := awsummary.Autoinit()
	if autoinit {
		cfg, err := config.LoadDefaultConfig(context.TODO())
		if err != nil {
			panic("configuration error, " + err.Error())
		}

		Client = elbv2.NewFromConfig(cfg)
	}
}

type LBV2Totals struct {
	Application int
	Network int
	Gateway int
}

//go:generate moq -out lb_moq.go -pkg sumlb . Elbv2interface

type Elbv2interface interface {
	DescribeLoadBalancers( ctx context.Context,
		params *elbv2.DescribeLoadBalancersInput, 
		optFns ...func(*elbv2.Options)) (*elbv2.DescribeLoadBalancersOutput, error)
}

func SetClient(c Elbv2interface) {
	Client = c
}

func List(region string, verbose bool) (LBV2Totals) {

	params := &elbv2.DescribeLoadBalancersInput{
		PageSize: aws.Int32(100),
	}

	resp, err := Client.DescribeLoadBalancers(context.TODO(), params)
	if err != nil {
		log.Fatal("Cannot get ALB data: ", err)
	}

	var stats LBV2Totals
	for _, playerOne := range resp.LoadBalancers{
		if playerOne.Type == types.LoadBalancerTypeEnumApplication {
			stats.Application ++
		}
		if playerOne.Type == types.LoadBalancerTypeEnumGateway {
			stats.Gateway ++
		}
		if playerOne.Type == types.LoadBalancerTypeEnumNetwork {
			stats.Network ++
		}
	}
	return stats
}
