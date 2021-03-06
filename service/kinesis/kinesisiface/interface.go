// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package kinesisiface provides an interface to enable mocking the Amazon Kinesis service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package kinesisiface

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/service/kinesis"
)

// ClientAPI provides an interface to enable mocking the
// kinesis.Client methods. This make unit testing your code that
// calls out to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the SDK's request pipeline.
//
//    // myFunc uses an SDK service client to make a request to
//    // Kinesis.
//    func myFunc(svc kinesisiface.ClientAPI) bool {
//        // Make svc.AddTagsToStream request
//    }
//
//    func main() {
//        cfg, err := external.LoadDefaultAWSConfig()
//        if err != nil {
//            panic("failed to load config, " + err.Error())
//        }
//
//        svc := kinesis.New(cfg)
//
//        myFunc(svc)
//    }
//
// In your _test.go file:
//
//    // Define a mock struct to be used in your unit tests of myFunc.
//    type mockClientClient struct {
//        kinesisiface.ClientPI
//    }
//    func (m *mockClientClient) AddTagsToStream(input *kinesis.AddTagsToStreamInput) (*kinesis.AddTagsToStreamOutput, error) {
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
	AddTagsToStreamRequest(*kinesis.AddTagsToStreamInput) kinesis.AddTagsToStreamRequest

	CreateStreamRequest(*kinesis.CreateStreamInput) kinesis.CreateStreamRequest

	DecreaseStreamRetentionPeriodRequest(*kinesis.DecreaseStreamRetentionPeriodInput) kinesis.DecreaseStreamRetentionPeriodRequest

	DeleteStreamRequest(*kinesis.DeleteStreamInput) kinesis.DeleteStreamRequest

	DeregisterStreamConsumerRequest(*kinesis.DeregisterStreamConsumerInput) kinesis.DeregisterStreamConsumerRequest

	DescribeLimitsRequest(*kinesis.DescribeLimitsInput) kinesis.DescribeLimitsRequest

	DescribeStreamRequest(*kinesis.DescribeStreamInput) kinesis.DescribeStreamRequest

	DescribeStreamConsumerRequest(*kinesis.DescribeStreamConsumerInput) kinesis.DescribeStreamConsumerRequest

	DescribeStreamSummaryRequest(*kinesis.DescribeStreamSummaryInput) kinesis.DescribeStreamSummaryRequest

	DisableEnhancedMonitoringRequest(*kinesis.DisableEnhancedMonitoringInput) kinesis.DisableEnhancedMonitoringRequest

	EnableEnhancedMonitoringRequest(*kinesis.EnableEnhancedMonitoringInput) kinesis.EnableEnhancedMonitoringRequest

	GetRecordsRequest(*kinesis.GetRecordsInput) kinesis.GetRecordsRequest

	GetShardIteratorRequest(*kinesis.GetShardIteratorInput) kinesis.GetShardIteratorRequest

	IncreaseStreamRetentionPeriodRequest(*kinesis.IncreaseStreamRetentionPeriodInput) kinesis.IncreaseStreamRetentionPeriodRequest

	ListShardsRequest(*kinesis.ListShardsInput) kinesis.ListShardsRequest

	ListStreamConsumersRequest(*kinesis.ListStreamConsumersInput) kinesis.ListStreamConsumersRequest

	ListStreamsRequest(*kinesis.ListStreamsInput) kinesis.ListStreamsRequest

	ListTagsForStreamRequest(*kinesis.ListTagsForStreamInput) kinesis.ListTagsForStreamRequest

	MergeShardsRequest(*kinesis.MergeShardsInput) kinesis.MergeShardsRequest

	PutRecordRequest(*kinesis.PutRecordInput) kinesis.PutRecordRequest

	PutRecordsRequest(*kinesis.PutRecordsInput) kinesis.PutRecordsRequest

	RegisterStreamConsumerRequest(*kinesis.RegisterStreamConsumerInput) kinesis.RegisterStreamConsumerRequest

	RemoveTagsFromStreamRequest(*kinesis.RemoveTagsFromStreamInput) kinesis.RemoveTagsFromStreamRequest

	SplitShardRequest(*kinesis.SplitShardInput) kinesis.SplitShardRequest

	StartStreamEncryptionRequest(*kinesis.StartStreamEncryptionInput) kinesis.StartStreamEncryptionRequest

	StopStreamEncryptionRequest(*kinesis.StopStreamEncryptionInput) kinesis.StopStreamEncryptionRequest

	UpdateShardCountRequest(*kinesis.UpdateShardCountInput) kinesis.UpdateShardCountRequest

	WaitUntilStreamExists(context.Context, *kinesis.DescribeStreamInput, ...aws.WaiterOption) error

	WaitUntilStreamNotExists(context.Context, *kinesis.DescribeStreamInput, ...aws.WaiterOption) error
}

var _ ClientAPI = (*kinesis.Client)(nil)
