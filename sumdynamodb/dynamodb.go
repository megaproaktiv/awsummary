package sumdynamodb

import (
	"context"
	"log"

	
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/megaproaktiv/awsummary"
)

var Client DynamoDBinterface


func init() {
	autoinit := awsummary.Autoinit()
	if autoinit {
		cfg, err := config.LoadDefaultConfig(context.TODO())
		if err != nil {
			panic("configuration error, " + err.Error())
		}

		Client = dynamodb.NewFromConfig(cfg)
	}
}

type DynamoDBTotals struct {
	Total int
	OnDemand int
	Provisioned int
}


//go:generate moq -out ddb_moq.go -pkg sumdynamodb . DynamoDBinterface

type DynamoDBinterface interface {
	ListTables( ctx context.Context,
		params *dynamodb.ListTablesInput,
		optFns ...func(*dynamodb.Options)) (*dynamodb.ListTablesOutput, error)
	DescribeTable( ctx context.Context,
		params *dynamodb.DescribeTableInput,
		optFns ...func(*dynamodb.Options)) (*dynamodb.DescribeTableOutput, error)
}


func SetClient(c DynamoDBinterface) {
	Client = c
}

func List(region string, verbose bool) (DynamoDBTotals){
	respTableList, err := Client.ListTables(context.TODO(), nil,
		func( o *dynamodb.Options){
			o.Region = region
		})
	if err != nil {
		log.Fatal("Cannot get DynamoDB data: ", err)
	}
	var stats DynamoDBTotals
	// All Tables
	for _, table := range respTableList.TableNames {
		params := &dynamodb.DescribeTableInput{
			TableName: &table,
		}
		resp, err := Client.DescribeTable(context.TODO(), params, 
		func( o *dynamodb.Options){
			o.Region = region
		})
		if err != nil {
			log.Fatal("Cannot get DynamoDB detail data: ", err)
		}
		stats.Total++
		provisionedRWCap :=  *resp.Table.ProvisionedThroughput.ReadCapacityUnits + *resp.Table.ProvisionedThroughput.WriteCapacityUnits
		if provisionedRWCap > 0 {
			stats.Provisioned++
		}else{
			stats.OnDemand++
		}


	}
	return stats
}