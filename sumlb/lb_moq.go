// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package sumlb

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/service/elasticloadbalancingv2"
	"sync"
)

// Ensure, that Elbv2interfaceMock does implement Elbv2interface.
// If this is not the case, regenerate this file with moq.
var _ Elbv2interface = &Elbv2interfaceMock{}

// Elbv2interfaceMock is a mock implementation of Elbv2interface.
//
//     func TestSomethingThatUsesElbv2interface(t *testing.T) {
//
//         // make and configure a mocked Elbv2interface
//         mockedElbv2interface := &Elbv2interfaceMock{
//             DescribeLoadBalancersFunc: func(ctx context.Context, params *elasticloadbalancingv2.DescribeLoadBalancersInput, optFns ...func(*elasticloadbalancingv2.Options)) (*elasticloadbalancingv2.DescribeLoadBalancersOutput, error) {
// 	               panic("mock out the DescribeLoadBalancers method")
//             },
//         }
//
//         // use mockedElbv2interface in code that requires Elbv2interface
//         // and then make assertions.
//
//     }
type Elbv2interfaceMock struct {
	// DescribeLoadBalancersFunc mocks the DescribeLoadBalancers method.
	DescribeLoadBalancersFunc func(ctx context.Context, params *elasticloadbalancingv2.DescribeLoadBalancersInput, optFns ...func(*elasticloadbalancingv2.Options)) (*elasticloadbalancingv2.DescribeLoadBalancersOutput, error)

	// calls tracks calls to the methods.
	calls struct {
		// DescribeLoadBalancers holds details about calls to the DescribeLoadBalancers method.
		DescribeLoadBalancers []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Params is the params argument value.
			Params *elasticloadbalancingv2.DescribeLoadBalancersInput
			// OptFns is the optFns argument value.
			OptFns []func(*elasticloadbalancingv2.Options)
		}
	}
	lockDescribeLoadBalancers sync.RWMutex
}

// DescribeLoadBalancers calls DescribeLoadBalancersFunc.
func (mock *Elbv2interfaceMock) DescribeLoadBalancers(ctx context.Context, params *elasticloadbalancingv2.DescribeLoadBalancersInput, optFns ...func(*elasticloadbalancingv2.Options)) (*elasticloadbalancingv2.DescribeLoadBalancersOutput, error) {
	if mock.DescribeLoadBalancersFunc == nil {
		panic("Elbv2interfaceMock.DescribeLoadBalancersFunc: method is nil but Elbv2interface.DescribeLoadBalancers was just called")
	}
	callInfo := struct {
		Ctx    context.Context
		Params *elasticloadbalancingv2.DescribeLoadBalancersInput
		OptFns []func(*elasticloadbalancingv2.Options)
	}{
		Ctx:    ctx,
		Params: params,
		OptFns: optFns,
	}
	mock.lockDescribeLoadBalancers.Lock()
	mock.calls.DescribeLoadBalancers = append(mock.calls.DescribeLoadBalancers, callInfo)
	mock.lockDescribeLoadBalancers.Unlock()
	return mock.DescribeLoadBalancersFunc(ctx, params, optFns...)
}

// DescribeLoadBalancersCalls gets all the calls that were made to DescribeLoadBalancers.
// Check the length with:
//     len(mockedElbv2interface.DescribeLoadBalancersCalls())
func (mock *Elbv2interfaceMock) DescribeLoadBalancersCalls() []struct {
	Ctx    context.Context
	Params *elasticloadbalancingv2.DescribeLoadBalancersInput
	OptFns []func(*elasticloadbalancingv2.Options)
} {
	var calls []struct {
		Ctx    context.Context
		Params *elasticloadbalancingv2.DescribeLoadBalancersInput
		OptFns []func(*elasticloadbalancingv2.Options)
	}
	mock.lockDescribeLoadBalancers.RLock()
	calls = mock.calls.DescribeLoadBalancers
	mock.lockDescribeLoadBalancers.RUnlock()
	return calls
}
