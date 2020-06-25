// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package fmsiface provides an interface to enable mocking the Firewall Management Service service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package fmsiface

import (
	"github.com/jviney/aws-sdk-go-v2/service/fms"
)

// ClientAPI provides an interface to enable mocking the
// fms.Client methods. This make unit testing your code that
// calls out to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the SDK's request pipeline.
//
//    // myFunc uses an SDK service client to make a request to
//    // FMS.
//    func myFunc(svc fmsiface.ClientAPI) bool {
//        // Make svc.AssociateAdminAccount request
//    }
//
//    func main() {
//        cfg, err := external.LoadDefaultAWSConfig()
//        if err != nil {
//            panic("failed to load config, " + err.Error())
//        }
//
//        svc := fms.New(cfg)
//
//        myFunc(svc)
//    }
//
// In your _test.go file:
//
//    // Define a mock struct to be used in your unit tests of myFunc.
//    type mockClientClient struct {
//        fmsiface.ClientPI
//    }
//    func (m *mockClientClient) AssociateAdminAccount(input *fms.AssociateAdminAccountInput) (*fms.AssociateAdminAccountOutput, error) {
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
	AssociateAdminAccountRequest(*fms.AssociateAdminAccountInput) fms.AssociateAdminAccountRequest

	DeleteNotificationChannelRequest(*fms.DeleteNotificationChannelInput) fms.DeleteNotificationChannelRequest

	DeletePolicyRequest(*fms.DeletePolicyInput) fms.DeletePolicyRequest

	DisassociateAdminAccountRequest(*fms.DisassociateAdminAccountInput) fms.DisassociateAdminAccountRequest

	GetAdminAccountRequest(*fms.GetAdminAccountInput) fms.GetAdminAccountRequest

	GetComplianceDetailRequest(*fms.GetComplianceDetailInput) fms.GetComplianceDetailRequest

	GetNotificationChannelRequest(*fms.GetNotificationChannelInput) fms.GetNotificationChannelRequest

	GetPolicyRequest(*fms.GetPolicyInput) fms.GetPolicyRequest

	GetProtectionStatusRequest(*fms.GetProtectionStatusInput) fms.GetProtectionStatusRequest

	ListComplianceStatusRequest(*fms.ListComplianceStatusInput) fms.ListComplianceStatusRequest

	ListMemberAccountsRequest(*fms.ListMemberAccountsInput) fms.ListMemberAccountsRequest

	ListPoliciesRequest(*fms.ListPoliciesInput) fms.ListPoliciesRequest

	ListTagsForResourceRequest(*fms.ListTagsForResourceInput) fms.ListTagsForResourceRequest

	PutNotificationChannelRequest(*fms.PutNotificationChannelInput) fms.PutNotificationChannelRequest

	PutPolicyRequest(*fms.PutPolicyInput) fms.PutPolicyRequest

	TagResourceRequest(*fms.TagResourceInput) fms.TagResourceRequest

	UntagResourceRequest(*fms.UntagResourceInput) fms.UntagResourceRequest
}

var _ ClientAPI = (*fms.Client)(nil)
