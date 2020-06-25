// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package translateiface provides an interface to enable mocking the Amazon Translate service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package translateiface

import (
	"github.com/jviney/aws-sdk-go-v2/service/translate"
)

// ClientAPI provides an interface to enable mocking the
// translate.Client methods. This make unit testing your code that
// calls out to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the SDK's request pipeline.
//
//    // myFunc uses an SDK service client to make a request to
//    // Amazon Translate.
//    func myFunc(svc translateiface.ClientAPI) bool {
//        // Make svc.DeleteTerminology request
//    }
//
//    func main() {
//        cfg, err := external.LoadDefaultAWSConfig()
//        if err != nil {
//            panic("failed to load config, " + err.Error())
//        }
//
//        svc := translate.New(cfg)
//
//        myFunc(svc)
//    }
//
// In your _test.go file:
//
//    // Define a mock struct to be used in your unit tests of myFunc.
//    type mockClientClient struct {
//        translateiface.ClientPI
//    }
//    func (m *mockClientClient) DeleteTerminology(input *translate.DeleteTerminologyInput) (*translate.DeleteTerminologyOutput, error) {
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
	DeleteTerminologyRequest(*translate.DeleteTerminologyInput) translate.DeleteTerminologyRequest

	DescribeTextTranslationJobRequest(*translate.DescribeTextTranslationJobInput) translate.DescribeTextTranslationJobRequest

	GetTerminologyRequest(*translate.GetTerminologyInput) translate.GetTerminologyRequest

	ImportTerminologyRequest(*translate.ImportTerminologyInput) translate.ImportTerminologyRequest

	ListTerminologiesRequest(*translate.ListTerminologiesInput) translate.ListTerminologiesRequest

	ListTextTranslationJobsRequest(*translate.ListTextTranslationJobsInput) translate.ListTextTranslationJobsRequest

	StartTextTranslationJobRequest(*translate.StartTextTranslationJobInput) translate.StartTextTranslationJobRequest

	StopTextTranslationJobRequest(*translate.StopTextTranslationJobInput) translate.StopTextTranslationJobRequest

	TranslateTextRequest(*translate.TranslateTextInput) translate.TranslateTextRequest
}

var _ ClientAPI = (*translate.Client)(nil)
