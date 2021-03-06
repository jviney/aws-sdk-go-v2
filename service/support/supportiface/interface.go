// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package supportiface provides an interface to enable mocking the AWS Support service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package supportiface

import (
	"github.com/jviney/aws-sdk-go-v2/service/support"
)

// ClientAPI provides an interface to enable mocking the
// support.Client methods. This make unit testing your code that
// calls out to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the SDK's request pipeline.
//
//    // myFunc uses an SDK service client to make a request to
//    // AWS Support.
//    func myFunc(svc supportiface.ClientAPI) bool {
//        // Make svc.AddAttachmentsToSet request
//    }
//
//    func main() {
//        cfg, err := external.LoadDefaultAWSConfig()
//        if err != nil {
//            panic("failed to load config, " + err.Error())
//        }
//
//        svc := support.New(cfg)
//
//        myFunc(svc)
//    }
//
// In your _test.go file:
//
//    // Define a mock struct to be used in your unit tests of myFunc.
//    type mockClientClient struct {
//        supportiface.ClientPI
//    }
//    func (m *mockClientClient) AddAttachmentsToSet(input *support.AddAttachmentsToSetInput) (*support.AddAttachmentsToSetOutput, error) {
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
	AddAttachmentsToSetRequest(*support.AddAttachmentsToSetInput) support.AddAttachmentsToSetRequest

	AddCommunicationToCaseRequest(*support.AddCommunicationToCaseInput) support.AddCommunicationToCaseRequest

	CreateCaseRequest(*support.CreateCaseInput) support.CreateCaseRequest

	DescribeAttachmentRequest(*support.DescribeAttachmentInput) support.DescribeAttachmentRequest

	DescribeCasesRequest(*support.DescribeCasesInput) support.DescribeCasesRequest

	DescribeCommunicationsRequest(*support.DescribeCommunicationsInput) support.DescribeCommunicationsRequest

	DescribeServicesRequest(*support.DescribeServicesInput) support.DescribeServicesRequest

	DescribeSeverityLevelsRequest(*support.DescribeSeverityLevelsInput) support.DescribeSeverityLevelsRequest

	DescribeTrustedAdvisorCheckRefreshStatusesRequest(*support.DescribeTrustedAdvisorCheckRefreshStatusesInput) support.DescribeTrustedAdvisorCheckRefreshStatusesRequest

	DescribeTrustedAdvisorCheckResultRequest(*support.DescribeTrustedAdvisorCheckResultInput) support.DescribeTrustedAdvisorCheckResultRequest

	DescribeTrustedAdvisorCheckSummariesRequest(*support.DescribeTrustedAdvisorCheckSummariesInput) support.DescribeTrustedAdvisorCheckSummariesRequest

	DescribeTrustedAdvisorChecksRequest(*support.DescribeTrustedAdvisorChecksInput) support.DescribeTrustedAdvisorChecksRequest

	RefreshTrustedAdvisorCheckRequest(*support.RefreshTrustedAdvisorCheckInput) support.RefreshTrustedAdvisorCheckRequest

	ResolveCaseRequest(*support.ResolveCaseInput) support.ResolveCaseRequest
}

var _ ClientAPI = (*support.Client)(nil)
