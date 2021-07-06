package sumlb

import (
	"context"
	"encoding/json"
	"os"
	"reflect"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/elasticloadbalancingv2"
)

func TestList(t *testing.T) {
	type args struct {
		region  string
		verbose bool
	}
	tests := []struct {
		name string
		args args
		want LBV2Totals
	}{
		{
			name: "One ALB",
			args: args{
				region: "eu-central-1",
				verbose: false,
			},
			want: LBV2Totals{
				Application: 1,
				Network:     0,
				Gateway:     0,
			},
		},
		{
			name: "ALB and NetworkLB",
			args: args{
				region: "eu-central-1",
				verbose: false,
			},
			want: LBV2Totals{
				Application: 2,
				Network:     2,
				Gateway:     0,
			},

		},
		
	}

	mockedElbv2interfaceOne := &Elbv2interfaceMock{
		DescribeLoadBalancersFunc: func(ctx context.Context, params *elasticloadbalancingv2.DescribeLoadBalancersInput, optFns ...func(*elasticloadbalancingv2.Options)) (*elasticloadbalancingv2.DescribeLoadBalancersOutput, error) {
			var output elasticloadbalancingv2.DescribeLoadBalancersOutput
			data, err := os.ReadFile("testdata/alb-one.json")
			if err != nil {
				t.Error("Cant read input testdata")
				t.Error(err)
			}
			json.Unmarshal(data, &output);
			return &output,nil		
		},
	}
	mockedElbv2interfaceTwo := &Elbv2interfaceMock{
		DescribeLoadBalancersFunc: func(ctx context.Context, params *elasticloadbalancingv2.DescribeLoadBalancersInput, optFns ...func(*elasticloadbalancingv2.Options)) (*elasticloadbalancingv2.DescribeLoadBalancersOutput, error) {
		var output elasticloadbalancingv2.DescribeLoadBalancersOutput
		data, err := os.ReadFile("testdata/alb-several.json")
		if err != nil {
			t.Error("Cant read input testdata")
			t.Error(err)
		}
		json.Unmarshal(data, &output);
		return &output,nil						
		},
	}
	for i, tt := range tests {
		if i == 0 {
			SetClient(mockedElbv2interfaceOne)
		}
		if i == 1 {
			SetClient(mockedElbv2interfaceTwo)
		}
		t.Run(tt.name, func(t *testing.T) {
			if got := List(tt.args.region, tt.args.verbose); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("List() = %v, want %v", got, tt.want)
			}
		})
	}
}
