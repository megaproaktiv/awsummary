package sumdynamodb

import (
	"context"
	"encoding/json"
	"os"
	"reflect"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

func TestList(t *testing.T) {
	os.Setenv("AUTO_INIT", "false")

	mockedDynamoDBinterfacePerRequest := &DynamoDBinterfaceMock{
		            DescribeTableFunc: func(ctx context.Context, params *dynamodb.DescribeTableInput, optFns ...func(*dynamodb.Options)) (*dynamodb.DescribeTableOutput, error) {
						   var output dynamodb.DescribeTableOutput
						   data, err := os.ReadFile("testdata/table-payperrequest.json")
						   if err != nil {
							   t.Error("Cant read input testdata")
							   t.Error(err)
						   }
						   json.Unmarshal(data, &output);
						   return &output,nil		
		            },
		            ListTablesFunc: func(ctx context.Context, params *dynamodb.ListTablesInput, optFns ...func(*dynamodb.Options)) (*dynamodb.ListTablesOutput, error) {
			         
						   var output dynamodb.ListTablesOutput
						   data, err := os.ReadFile("testdata/table-list-payperrequest.json")
						   if err != nil {
							   t.Error("Cant read input testdata")
							   t.Error(err)
						   }
						   json.Unmarshal(data, &output);
						   return &output,nil		
		            },
		        }
	mockedDynamoDBinterfaceReserved := &DynamoDBinterfaceMock{
		            DescribeTableFunc: func(ctx context.Context, params *dynamodb.DescribeTableInput, optFns ...func(*dynamodb.Options)) (*dynamodb.DescribeTableOutput, error) {

						   var output dynamodb.DescribeTableOutput
						   data, err := os.ReadFile("testdata/table-payreserved.json")
						   if err != nil {
							   t.Error("Cant read input testdata")
							   t.Error(err)
						   }
						   json.Unmarshal(data, &output);
						   return &output,nil	
		            },
		            ListTablesFunc: func(ctx context.Context, params *dynamodb.ListTablesInput, optFns ...func(*dynamodb.Options)) (*dynamodb.ListTablesOutput, error) {
						   var output dynamodb.ListTablesOutput
						   data, err := os.ReadFile("testdata/table-list-payperrequest.json")
						   if err != nil {
							   t.Error("Cant read input testdata")
							   t.Error(err)
						   }
						   json.Unmarshal(data, &output);
						   return &output,nil	

		            },
		        }
	type args struct {
		region  string
		verbose bool
	}
	tests := []struct {
		name string
		args args
		want DynamoDBTotals
	}{
		{
			name: "Per Request",
			args: args{
				region:  "eu-central-1",
				verbose: false,
			},
			want: DynamoDBTotals{
				Total:       1,
				OnDemand:    1,
				Provisioned: 0,
			},
		},
		{
			name: "Reserved",
			args: args{
				region:  "eu-central-1",
				verbose: false,
			},
			want: DynamoDBTotals{
				Total:       1,
				OnDemand:    0,
				Provisioned: 1,
			},
		},
	
	}
	for i, tt := range tests {
		if i == 0 {
			SetClient(mockedDynamoDBinterfacePerRequest)
		}
		
		if i == 1 {
			SetClient(mockedDynamoDBinterfaceReserved)
		}
		
		t.Run(tt.name, func(t *testing.T) {
			if got := List(tt.args.region, tt.args.verbose); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("List() = %v, want %v", got, tt.want)
			}
		})
	}
}
