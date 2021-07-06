// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package sumdynamodb

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"sync"
)

// Ensure, that DynamoDBinterfaceMock does implement DynamoDBinterface.
// If this is not the case, regenerate this file with moq.
var _ DynamoDBinterface = &DynamoDBinterfaceMock{}

// DynamoDBinterfaceMock is a mock implementation of DynamoDBinterface.
//
//     func TestSomethingThatUsesDynamoDBinterface(t *testing.T) {
//
//         // make and configure a mocked DynamoDBinterface
//         mockedDynamoDBinterface := &DynamoDBinterfaceMock{
//             DescribeTableFunc: func(ctx context.Context, params *dynamodb.DescribeTableInput, optFns ...func(*dynamodb.Options)) (*dynamodb.DescribeTableOutput, error) {
// 	               panic("mock out the DescribeTable method")
//             },
//             ListTablesFunc: func(ctx context.Context, params *dynamodb.ListTablesInput, optFns ...func(*dynamodb.Options)) (*dynamodb.ListTablesOutput, error) {
// 	               panic("mock out the ListTables method")
//             },
//         }
//
//         // use mockedDynamoDBinterface in code that requires DynamoDBinterface
//         // and then make assertions.
//
//     }
type DynamoDBinterfaceMock struct {
	// DescribeTableFunc mocks the DescribeTable method.
	DescribeTableFunc func(ctx context.Context, params *dynamodb.DescribeTableInput, optFns ...func(*dynamodb.Options)) (*dynamodb.DescribeTableOutput, error)

	// ListTablesFunc mocks the ListTables method.
	ListTablesFunc func(ctx context.Context, params *dynamodb.ListTablesInput, optFns ...func(*dynamodb.Options)) (*dynamodb.ListTablesOutput, error)

	// calls tracks calls to the methods.
	calls struct {
		// DescribeTable holds details about calls to the DescribeTable method.
		DescribeTable []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Params is the params argument value.
			Params *dynamodb.DescribeTableInput
			// OptFns is the optFns argument value.
			OptFns []func(*dynamodb.Options)
		}
		// ListTables holds details about calls to the ListTables method.
		ListTables []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Params is the params argument value.
			Params *dynamodb.ListTablesInput
			// OptFns is the optFns argument value.
			OptFns []func(*dynamodb.Options)
		}
	}
	lockDescribeTable sync.RWMutex
	lockListTables    sync.RWMutex
}

// DescribeTable calls DescribeTableFunc.
func (mock *DynamoDBinterfaceMock) DescribeTable(ctx context.Context, params *dynamodb.DescribeTableInput, optFns ...func(*dynamodb.Options)) (*dynamodb.DescribeTableOutput, error) {
	if mock.DescribeTableFunc == nil {
		panic("DynamoDBinterfaceMock.DescribeTableFunc: method is nil but DynamoDBinterface.DescribeTable was just called")
	}
	callInfo := struct {
		Ctx    context.Context
		Params *dynamodb.DescribeTableInput
		OptFns []func(*dynamodb.Options)
	}{
		Ctx:    ctx,
		Params: params,
		OptFns: optFns,
	}
	mock.lockDescribeTable.Lock()
	mock.calls.DescribeTable = append(mock.calls.DescribeTable, callInfo)
	mock.lockDescribeTable.Unlock()
	return mock.DescribeTableFunc(ctx, params, optFns...)
}

// DescribeTableCalls gets all the calls that were made to DescribeTable.
// Check the length with:
//     len(mockedDynamoDBinterface.DescribeTableCalls())
func (mock *DynamoDBinterfaceMock) DescribeTableCalls() []struct {
	Ctx    context.Context
	Params *dynamodb.DescribeTableInput
	OptFns []func(*dynamodb.Options)
} {
	var calls []struct {
		Ctx    context.Context
		Params *dynamodb.DescribeTableInput
		OptFns []func(*dynamodb.Options)
	}
	mock.lockDescribeTable.RLock()
	calls = mock.calls.DescribeTable
	mock.lockDescribeTable.RUnlock()
	return calls
}

// ListTables calls ListTablesFunc.
func (mock *DynamoDBinterfaceMock) ListTables(ctx context.Context, params *dynamodb.ListTablesInput, optFns ...func(*dynamodb.Options)) (*dynamodb.ListTablesOutput, error) {
	if mock.ListTablesFunc == nil {
		panic("DynamoDBinterfaceMock.ListTablesFunc: method is nil but DynamoDBinterface.ListTables was just called")
	}
	callInfo := struct {
		Ctx    context.Context
		Params *dynamodb.ListTablesInput
		OptFns []func(*dynamodb.Options)
	}{
		Ctx:    ctx,
		Params: params,
		OptFns: optFns,
	}
	mock.lockListTables.Lock()
	mock.calls.ListTables = append(mock.calls.ListTables, callInfo)
	mock.lockListTables.Unlock()
	return mock.ListTablesFunc(ctx, params, optFns...)
}

// ListTablesCalls gets all the calls that were made to ListTables.
// Check the length with:
//     len(mockedDynamoDBinterface.ListTablesCalls())
func (mock *DynamoDBinterfaceMock) ListTablesCalls() []struct {
	Ctx    context.Context
	Params *dynamodb.ListTablesInput
	OptFns []func(*dynamodb.Options)
} {
	var calls []struct {
		Ctx    context.Context
		Params *dynamodb.ListTablesInput
		OptFns []func(*dynamodb.Options)
	}
	mock.lockListTables.RLock()
	calls = mock.calls.ListTables
	mock.lockListTables.RUnlock()
	return calls
}