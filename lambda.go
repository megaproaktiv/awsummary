package awsummary

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/lambda"
	"log"
)

func ListLambda(region string, verbose bool) (lambdaNumber int) {
	var regions = map[string]bool{
		"us-east-1":      true,
		"us-west-2":      true,
		"eu-west-1":      true,
		"ap-northeast-1": true,
	}
	if regions[region] {
		svc := lambda.New(session.New(&aws.Config{Region: aws.String(region)}))
		params := &lambda.ListFunctionsInput{
			MaxItems: aws.Int64(100),
		}
		resp, err := svc.ListFunctions(params)
		if err != nil {
			log.Fatal("Cannot get Lambda data: ", err)
		}

		if verbose {
			for _, name := range resp.Functions {
				log.Println("Lambda Name: ", *name.FunctionName)
			}
		}

		lambdaNumber = len(resp.Functions)
	} else {
		lambdaNumber = 0
	}
	return lambdaNumber
}
