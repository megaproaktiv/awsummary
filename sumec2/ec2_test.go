package sumec2

import (
	"context"
	"encoding/json"
	"os"
	"testing"

	//	"gotest.tools/v3/assert"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"gotest.tools/v3/assert"
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
		name: "One Instance",
		args: args{
			verbose: false,
			region: "eu-central-1",
		},
		wantEc2Number: 1,
		wantEc2RunningNumber: 1,
		wantEc2RunningWindows: 0,
		},
		{
		name: "Windows and Linux",
		args: args{
			region: "eu-central-1",
			verbose: false,
		},
		wantEc2Number: 7,
		wantEc2RunningNumber: 5,
		wantEc2RunningWindows: 3,
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
		DescribeNatGatewaysFunc: func(ctx context.Context, params *ec2.DescribeNatGatewaysInput, optFns ...func(*ec2.Options)) (*ec2.DescribeNatGatewaysOutput, error) {
			panic("mock out the DescribeNatGateways method")
		},
	}
	mockedEc2InterfaceWindows := &Ec2InterfaceMock{
		DescribeInstancesFunc: func(ctx context.Context, params *ec2.DescribeInstancesInput, optFns ...func(*ec2.Options)) (*ec2.DescribeInstancesOutput, error) {
			var output ec2.DescribeInstancesOutput
			data, err := os.ReadFile("testdata/somewindowsec2.json")
			if err != nil {
				t.Error("Cant read input testdata")
				t.Error(err)
			}
			json.Unmarshal(data, &output);
			return &output,nil		
		},
		DescribeNatGatewaysFunc: func(ctx context.Context, params *ec2.DescribeNatGatewaysInput, optFns ...func(*ec2.Options)) (*ec2.DescribeNatGatewaysOutput, error) {
		panic("mock out the DescribeNatGateways method")
	},
	}
	for i, tt := range tests {
		if i == 0 {
			SetClient(mockedEc2Interface)
		}
		
		if i == 1 {
			SetClient(mockedEc2InterfaceWindows)
		}
		
		t.Run(tt.name, func(t *testing.T) {
			totals := ListInstances(tt.args.region, tt.args.verbose)
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

func TestListNatGW(t *testing.T){
	os.Setenv("AUTO_INIT", "false")
	var region = "eu-central-1"
		mockedEc2InterfaceTwo := &Ec2InterfaceMock{
            DescribeInstancesFunc: func(ctx context.Context, params *ec2.DescribeInstancesInput, optFns ...func(*ec2.Options)) (*ec2.DescribeInstancesOutput, error) {
	               panic("mock out the DescribeInstances method")
            },
            DescribeNatGatewaysFunc: func(ctx context.Context, params *ec2.DescribeNatGatewaysInput, optFns ...func(*ec2.Options)) (*ec2.DescribeNatGatewaysOutput, error) {
				var output ec2.DescribeNatGatewaysOutput
				data, err := os.ReadFile("testdata/nat-gateway-two.json")
				if err != nil {
					t.Error("Cant read input testdata")
					t.Error(err)
				}
				json.Unmarshal(data, &output);
				return &output,nil	
            },
        }
		mockedEc2InterfaceNone := &Ec2InterfaceMock{
            DescribeInstancesFunc: func(ctx context.Context, params *ec2.DescribeInstancesInput, optFns ...func(*ec2.Options)) (*ec2.DescribeInstancesOutput, error) {
	               panic("mock out the DescribeInstances method")
            },
            DescribeNatGatewaysFunc: func(ctx context.Context, params *ec2.DescribeNatGatewaysInput, optFns ...func(*ec2.Options)) (*ec2.DescribeNatGatewaysOutput, error) {
				var output ec2.DescribeNatGatewaysOutput
				data, err := os.ReadFile("testdata/nat-gateway-none.json")
				if err != nil {
					t.Error("Cant read input testdata")
					t.Error(err)
				}
				json.Unmarshal(data, &output);
				return &output,nil	
            },
        }
		SetClient(mockedEc2InterfaceTwo)
		
		expect := 2
		got := ListNatGW(region, false)
		assert.Equal(t, expect, got.Total, "Number should be 2")

		SetClient(mockedEc2InterfaceNone)
		expect = 0
		got = ListNatGW(region, false)
		assert.Equal(t, expect, got.Total, "Number should be 0")
}