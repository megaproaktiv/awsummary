package sumcloudformation

import (
	"context"
	"encoding/json"
	"os"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/cloudformation"
	"gotest.tools/v3/assert"
)

func TestList(t *testing.T) {
	os.Setenv("AUTO_INIT", "false")
	// make and configure a mocked CloudFormationInterface
	mockedCloudFormationInterface := &CloudFormationInterfaceMock{
		ListStacksFunc: func(ctx context.Context, params *cloudformation.ListStacksInput, optFns ...func(*cloudformation.Options)) (*cloudformation.ListStacksOutput, error) {
			var output cloudformation.ListStacksOutput
			data, err := os.ReadFile("testdata/list-stacks.json")
			if err != nil {
				t.Error("Cant read input testdata")
				t.Error(err)
			}
			json.Unmarshal(data, &output);
			return &output,nil		
		},
	}

	SetClient(mockedCloudFormationInterface)

	expected := 71
	actual := List("eu-central-1", false)
	assert.Equal(t,expected,actual, "Stack count should be 8 ")
}
