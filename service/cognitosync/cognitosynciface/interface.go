// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package cognitosynciface provides an interface to enable mocking the Amazon Cognito Sync service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package cognitosynciface

import (
	"github.com/jviney/aws-sdk-go-v2/service/cognitosync"
)

// ClientAPI provides an interface to enable mocking the
// cognitosync.Client methods. This make unit testing your code that
// calls out to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the SDK's request pipeline.
//
//    // myFunc uses an SDK service client to make a request to
//    // Amazon Cognito Sync.
//    func myFunc(svc cognitosynciface.ClientAPI) bool {
//        // Make svc.BulkPublish request
//    }
//
//    func main() {
//        cfg, err := external.LoadDefaultAWSConfig()
//        if err != nil {
//            panic("failed to load config, " + err.Error())
//        }
//
//        svc := cognitosync.New(cfg)
//
//        myFunc(svc)
//    }
//
// In your _test.go file:
//
//    // Define a mock struct to be used in your unit tests of myFunc.
//    type mockClientClient struct {
//        cognitosynciface.ClientPI
//    }
//    func (m *mockClientClient) BulkPublish(input *cognitosync.BulkPublishInput) (*cognitosync.BulkPublishOutput, error) {
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
	BulkPublishRequest(*cognitosync.BulkPublishInput) cognitosync.BulkPublishRequest

	DeleteDatasetRequest(*cognitosync.DeleteDatasetInput) cognitosync.DeleteDatasetRequest

	DescribeDatasetRequest(*cognitosync.DescribeDatasetInput) cognitosync.DescribeDatasetRequest

	DescribeIdentityPoolUsageRequest(*cognitosync.DescribeIdentityPoolUsageInput) cognitosync.DescribeIdentityPoolUsageRequest

	DescribeIdentityUsageRequest(*cognitosync.DescribeIdentityUsageInput) cognitosync.DescribeIdentityUsageRequest

	GetBulkPublishDetailsRequest(*cognitosync.GetBulkPublishDetailsInput) cognitosync.GetBulkPublishDetailsRequest

	GetCognitoEventsRequest(*cognitosync.GetCognitoEventsInput) cognitosync.GetCognitoEventsRequest

	GetIdentityPoolConfigurationRequest(*cognitosync.GetIdentityPoolConfigurationInput) cognitosync.GetIdentityPoolConfigurationRequest

	ListDatasetsRequest(*cognitosync.ListDatasetsInput) cognitosync.ListDatasetsRequest

	ListIdentityPoolUsageRequest(*cognitosync.ListIdentityPoolUsageInput) cognitosync.ListIdentityPoolUsageRequest

	ListRecordsRequest(*cognitosync.ListRecordsInput) cognitosync.ListRecordsRequest

	RegisterDeviceRequest(*cognitosync.RegisterDeviceInput) cognitosync.RegisterDeviceRequest

	SetCognitoEventsRequest(*cognitosync.SetCognitoEventsInput) cognitosync.SetCognitoEventsRequest

	SetIdentityPoolConfigurationRequest(*cognitosync.SetIdentityPoolConfigurationInput) cognitosync.SetIdentityPoolConfigurationRequest

	SubscribeToDatasetRequest(*cognitosync.SubscribeToDatasetInput) cognitosync.SubscribeToDatasetRequest

	UnsubscribeFromDatasetRequest(*cognitosync.UnsubscribeFromDatasetInput) cognitosync.UnsubscribeFromDatasetRequest

	UpdateRecordsRequest(*cognitosync.UpdateRecordsInput) cognitosync.UpdateRecordsRequest
}

var _ ClientAPI = (*cognitosync.Client)(nil)
