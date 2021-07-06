package sumcloudformation

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/cloudformation"
	"github.com/aws/aws-sdk-go-v2/service/cloudformation/types"
	"github.com/megaproaktiv/awsummary"
	"log"
)

var Client CloudFormationInterface

const DefaultRegion = "eu-central-1"

func init() {
	autoinit := awsummary.Autoinit()
	if autoinit {
		cfg, err := config.LoadDefaultConfig(context.TODO())
		if err != nil {
			panic("configuration error, " + err.Error())
		}
		
		Client = cloudformation.NewFromConfig(cfg)
	}
}


//go:generate moq -out cfn_moq.go -pkg sumcloudformation . CloudFormationInterface

type CloudFormationInterface interface {
	ListStacks(ctx context.Context, params *cloudformation.ListStacksInput, optFns ...func(*cloudformation.Options)) (*cloudformation.ListStacksOutput, error)
}

func SetClient(c CloudFormationInterface) {
	Client = c
}

func List(region string, verbose bool) (cfnNumber int) {
	
	params := &cloudformation.ListStacksInput{
		StackStatusFilter: []types.StackStatus{
			types.StackStatusCreateComplete,
			types.StackStatusCreateFailed,
			types.StackStatusCreateInProgress,
			types.StackStatusDeleteFailed,
			types.StackStatusDeleteInProgress,
			types.StackStatusRollbackFailed,
			types.StackStatusRollbackInProgress,
			types.StackStatusRollbackComplete,
			types.StackStatusUpdateComplete,
			types.StackStatusUpdateCompleteCleanupInProgress,
			types.StackStatusUpdateInProgress,
			types.StackStatusUpdateRollbackComplete,
			types.StackStatusUpdateRollbackCompleteCleanupInProgress,
			types.StackStatusUpdateRollbackFailed,
			types.StackStatusUpdateRollbackInProgress,
		},
	}
	resp, err := Client.ListStacks(context.TODO(), params, func(o *cloudformation.Options) {
		o.Region = region
	} )
	if err != nil {
		log.Fatal("Cannot get CFN data: ", err)
	}

	if verbose {
		for _, name := range resp.StackSummaries {
			log.Println("CFN Stack Name: ", *name.StackName, " Status: ", name.StackStatus)
		}
	}
	cfnNumber = len(resp.StackSummaries)
	return cfnNumber
}
