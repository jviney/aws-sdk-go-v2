// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package sfniface provides an interface to enable mocking the AWS Step Functions service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package sfniface

import (
	"github.com/jviney/aws-sdk-go-v2/service/sfn"
)

// ClientAPI provides an interface to enable mocking the
// sfn.Client methods. This make unit testing your code that
// calls out to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the SDK's request pipeline.
//
//    // myFunc uses an SDK service client to make a request to
//    // AWS SFN.
//    func myFunc(svc sfniface.ClientAPI) bool {
//        // Make svc.CreateActivity request
//    }
//
//    func main() {
//        cfg, err := external.LoadDefaultAWSConfig()
//        if err != nil {
//            panic("failed to load config, " + err.Error())
//        }
//
//        svc := sfn.New(cfg)
//
//        myFunc(svc)
//    }
//
// In your _test.go file:
//
//    // Define a mock struct to be used in your unit tests of myFunc.
//    type mockClientClient struct {
//        sfniface.ClientPI
//    }
//    func (m *mockClientClient) CreateActivity(input *sfn.CreateActivityInput) (*sfn.CreateActivityOutput, error) {
//        // mock response/functionality
//    }
//
//    func TestMyFunc(t *testing.T) {
//        // Setup Test
//        mockSvc := &mockClientClient{}
//
//        myfunc(mockSvc)
//
//        // Verify myFunc's functionality
//    }
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters. Its suggested to use the pattern above for testing, or using
// tooling to generate mocks to satisfy the interfaces.
type ClientAPI interface {
	CreateActivityRequest(*sfn.CreateActivityInput) sfn.CreateActivityRequest

	CreateStateMachineRequest(*sfn.CreateStateMachineInput) sfn.CreateStateMachineRequest

	DeleteActivityRequest(*sfn.DeleteActivityInput) sfn.DeleteActivityRequest

	DeleteStateMachineRequest(*sfn.DeleteStateMachineInput) sfn.DeleteStateMachineRequest

	DescribeActivityRequest(*sfn.DescribeActivityInput) sfn.DescribeActivityRequest

	DescribeExecutionRequest(*sfn.DescribeExecutionInput) sfn.DescribeExecutionRequest

	DescribeStateMachineRequest(*sfn.DescribeStateMachineInput) sfn.DescribeStateMachineRequest

	DescribeStateMachineForExecutionRequest(*sfn.DescribeStateMachineForExecutionInput) sfn.DescribeStateMachineForExecutionRequest

	GetActivityTaskRequest(*sfn.GetActivityTaskInput) sfn.GetActivityTaskRequest

	GetExecutionHistoryRequest(*sfn.GetExecutionHistoryInput) sfn.GetExecutionHistoryRequest

	ListActivitiesRequest(*sfn.ListActivitiesInput) sfn.ListActivitiesRequest

	ListExecutionsRequest(*sfn.ListExecutionsInput) sfn.ListExecutionsRequest

	ListStateMachinesRequest(*sfn.ListStateMachinesInput) sfn.ListStateMachinesRequest

	ListTagsForResourceRequest(*sfn.ListTagsForResourceInput) sfn.ListTagsForResourceRequest

	SendTaskFailureRequest(*sfn.SendTaskFailureInput) sfn.SendTaskFailureRequest

	SendTaskHeartbeatRequest(*sfn.SendTaskHeartbeatInput) sfn.SendTaskHeartbeatRequest

	SendTaskSuccessRequest(*sfn.SendTaskSuccessInput) sfn.SendTaskSuccessRequest

	StartExecutionRequest(*sfn.StartExecutionInput) sfn.StartExecutionRequest

	StopExecutionRequest(*sfn.StopExecutionInput) sfn.StopExecutionRequest

	TagResourceRequest(*sfn.TagResourceInput) sfn.TagResourceRequest

	UntagResourceRequest(*sfn.UntagResourceInput) sfn.UntagResourceRequest

	UpdateStateMachineRequest(*sfn.UpdateStateMachineInput) sfn.UpdateStateMachineRequest
}

var _ ClientAPI = (*sfn.Client)(nil)
