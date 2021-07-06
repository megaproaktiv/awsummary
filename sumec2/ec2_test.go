package sumec2

import (
	"context"
	"encoding/json"
	"os"
	"testing"
//	"gotest.tools/v3/assert"
	"github.com/aws/aws-sdk-go-v2/service/ec2"

)

func TestListEC2(t *testing.T) {
	os.Setenv("AUTO_INIT", "false")

	type args struct {
		region  string
		verbose bool
	}
	tests := []struct {
		name                  string
		args                  args
		wantEc2Number         int
		wantEc2RunningNumber  int
		wantEc2RunningWindows int
	}{
		{
		name: "Alle Instances",
		args: args{
			region: "eu-central-1",
			verbose: false,
		},
		wantEc2Number: 1,
		wantEc2RunningNumber: 1,
		wantEc2RunningWindows: 0,
		},
	}
	mockedEc2Interface := &Ec2InterfaceMock{
		DescribeInstancesFunc: func(ctx context.Context, params *ec2.DescribeInstancesInput, optFns ...func(*ec2.Options)) (*ec2.DescribeInstancesOutput, error) {
			var output ec2.DescribeInstancesOutput
			data, err := os.ReadFile("testdata/describe-instances.json")
			if err != nil {
				t.Error("Cant read input testdata")
				t.Error(err)
			}
			json.Unmarshal(data, &output);
			return &output,nil		
		},
	}
	SetClient(mockedEc2Interface)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			totals := List(tt.args.region, tt.args.verbose)
			gotEc2Number := totals.Total
			gotEc2RunningNumber := totals.Running
			gotEc2RunningWindows := totals.Windows
			if gotEc2Number != tt.wantEc2Number {
				t.Errorf("ListEC2() gotEc2Number = %v, want %v", gotEc2Number, tt.wantEc2Number)
			}
			if gotEc2RunningNumber != tt.wantEc2RunningNumber {
				t.Errorf("ListEC2() gotEc2RunningNumber = %v, want %v", gotEc2RunningNumber, tt.wantEc2RunningNumber)
			}
			if gotEc2RunningWindows != tt.wantEc2RunningWindows {
				t.Errorf("ListEC2() gotEc2RunningWindows = %v, want %v", gotEc2RunningWindows, tt.wantEc2RunningWindows)
			}
		})
	}
}
