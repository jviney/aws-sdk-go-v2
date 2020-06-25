// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package marketplacecatalogiface provides an interface to enable mocking the AWS Marketplace Catalog Service service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package marketplacecatalogiface

import (
	"github.com/jviney/aws-sdk-go-v2/service/marketplacecatalog"
)

// ClientAPI provides an interface to enable mocking the
// marketplacecatalog.Client methods. This make unit testing your code that
// calls out to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the SDK's request pipeline.
//
//    // myFunc uses an SDK service client to make a request to
//    // AWS Marketplace Catalog.
//    func myFunc(svc marketplacecatalogiface.ClientAPI) bool {
//        // Make svc.CancelChangeSet request
//    }
//
//    func main() {
//        cfg, err := external.LoadDefaultAWSConfig()
//        if err != nil {
//            panic("failed to load config, " + err.Error())
//        }
//
//        svc := marketplacecatalog.New(cfg)
//
//        myFunc(svc)
//    }
//
// In your _test.go file:
//
//    // Define a mock struct to be used in your unit tests of myFunc.
//    type mockClientClient struct {
//        marketplacecatalogiface.ClientPI
//    }
//    func (m *mockClientClient) CancelChangeSet(input *marketplacecatalog.CancelChangeSetInput) (*marketplacecatalog.CancelChangeSetOutput, error) {
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
	CancelChangeSetRequest(*marketplacecatalog.CancelChangeSetInput) marketplacecatalog.CancelChangeSetRequest

	DescribeChangeSetRequest(*marketplacecatalog.DescribeChangeSetInput) marketplacecatalog.DescribeChangeSetRequest

	DescribeEntityRequest(*marketplacecatalog.DescribeEntityInput) marketplacecatalog.DescribeEntityRequest

	ListChangeSetsRequest(*marketplacecatalog.ListChangeSetsInput) marketplacecatalog.ListChangeSetsRequest

	ListEntitiesRequest(*marketplacecatalog.ListEntitiesInput) marketplacecatalog.ListEntitiesRequest

	StartChangeSetRequest(*marketplacecatalog.StartChangeSetInput) marketplacecatalog.StartChangeSetRequest
}

var _ ClientAPI = (*marketplacecatalog.Client)(nil)
