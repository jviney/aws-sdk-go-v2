// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package organizationsiface provides an interface to enable mocking the AWS Organizations service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package organizationsiface

import (
	"github.com/jviney/aws-sdk-go-v2/service/organizations"
)

// ClientAPI provides an interface to enable mocking the
// organizations.Client methods. This make unit testing your code that
// calls out to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the SDK's request pipeline.
//
//    // myFunc uses an SDK service client to make a request to
//    // Organizations.
//    func myFunc(svc organizationsiface.ClientAPI) bool {
//        // Make svc.AcceptHandshake request
//    }
//
//    func main() {
//        cfg, err := external.LoadDefaultAWSConfig()
//        if err != nil {
//            panic("failed to load config, " + err.Error())
//        }
//
//        svc := organizations.New(cfg)
//
//        myFunc(svc)
//    }
//
// In your _test.go file:
//
//    // Define a mock struct to be used in your unit tests of myFunc.
//    type mockClientClient struct {
//        organizationsiface.ClientPI
//    }
//    func (m *mockClientClient) AcceptHandshake(input *organizations.AcceptHandshakeInput) (*organizations.AcceptHandshakeOutput, error) {
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
	AcceptHandshakeRequest(*organizations.AcceptHandshakeInput) organizations.AcceptHandshakeRequest

	AttachPolicyRequest(*organizations.AttachPolicyInput) organizations.AttachPolicyRequest

	CancelHandshakeRequest(*organizations.CancelHandshakeInput) organizations.CancelHandshakeRequest

	CreateAccountRequest(*organizations.CreateAccountInput) organizations.CreateAccountRequest

	CreateGovCloudAccountRequest(*organizations.CreateGovCloudAccountInput) organizations.CreateGovCloudAccountRequest

	CreateOrganizationRequest(*organizations.CreateOrganizationInput) organizations.CreateOrganizationRequest

	CreateOrganizationalUnitRequest(*organizations.CreateOrganizationalUnitInput) organizations.CreateOrganizationalUnitRequest

	CreatePolicyRequest(*organizations.CreatePolicyInput) organizations.CreatePolicyRequest

	DeclineHandshakeRequest(*organizations.DeclineHandshakeInput) organizations.DeclineHandshakeRequest

	DeleteOrganizationRequest(*organizations.DeleteOrganizationInput) organizations.DeleteOrganizationRequest

	DeleteOrganizationalUnitRequest(*organizations.DeleteOrganizationalUnitInput) organizations.DeleteOrganizationalUnitRequest

	DeletePolicyRequest(*organizations.DeletePolicyInput) organizations.DeletePolicyRequest

	DeregisterDelegatedAdministratorRequest(*organizations.DeregisterDelegatedAdministratorInput) organizations.DeregisterDelegatedAdministratorRequest

	DescribeAccountRequest(*organizations.DescribeAccountInput) organizations.DescribeAccountRequest

	DescribeCreateAccountStatusRequest(*organizations.DescribeCreateAccountStatusInput) organizations.DescribeCreateAccountStatusRequest

	DescribeEffectivePolicyRequest(*organizations.DescribeEffectivePolicyInput) organizations.DescribeEffectivePolicyRequest

	DescribeHandshakeRequest(*organizations.DescribeHandshakeInput) organizations.DescribeHandshakeRequest

	DescribeOrganizationRequest(*organizations.DescribeOrganizationInput) organizations.DescribeOrganizationRequest

	DescribeOrganizationalUnitRequest(*organizations.DescribeOrganizationalUnitInput) organizations.DescribeOrganizationalUnitRequest

	DescribePolicyRequest(*organizations.DescribePolicyInput) organizations.DescribePolicyRequest

	DetachPolicyRequest(*organizations.DetachPolicyInput) organizations.DetachPolicyRequest

	DisableAWSServiceAccessRequest(*organizations.DisableAWSServiceAccessInput) organizations.DisableAWSServiceAccessRequest

	DisablePolicyTypeRequest(*organizations.DisablePolicyTypeInput) organizations.DisablePolicyTypeRequest

	EnableAWSServiceAccessRequest(*organizations.EnableAWSServiceAccessInput) organizations.EnableAWSServiceAccessRequest

	EnableAllFeaturesRequest(*organizations.EnableAllFeaturesInput) organizations.EnableAllFeaturesRequest

	EnablePolicyTypeRequest(*organizations.EnablePolicyTypeInput) organizations.EnablePolicyTypeRequest

	InviteAccountToOrganizationRequest(*organizations.InviteAccountToOrganizationInput) organizations.InviteAccountToOrganizationRequest

	LeaveOrganizationRequest(*organizations.LeaveOrganizationInput) organizations.LeaveOrganizationRequest

	ListAWSServiceAccessForOrganizationRequest(*organizations.ListAWSServiceAccessForOrganizationInput) organizations.ListAWSServiceAccessForOrganizationRequest

	ListAccountsRequest(*organizations.ListAccountsInput) organizations.ListAccountsRequest

	ListAccountsForParentRequest(*organizations.ListAccountsForParentInput) organizations.ListAccountsForParentRequest

	ListChildrenRequest(*organizations.ListChildrenInput) organizations.ListChildrenRequest

	ListCreateAccountStatusRequest(*organizations.ListCreateAccountStatusInput) organizations.ListCreateAccountStatusRequest

	ListDelegatedAdministratorsRequest(*organizations.ListDelegatedAdministratorsInput) organizations.ListDelegatedAdministratorsRequest

	ListDelegatedServicesForAccountRequest(*organizations.ListDelegatedServicesForAccountInput) organizations.ListDelegatedServicesForAccountRequest

	ListHandshakesForAccountRequest(*organizations.ListHandshakesForAccountInput) organizations.ListHandshakesForAccountRequest

	ListHandshakesForOrganizationRequest(*organizations.ListHandshakesForOrganizationInput) organizations.ListHandshakesForOrganizationRequest

	ListOrganizationalUnitsForParentRequest(*organizations.ListOrganizationalUnitsForParentInput) organizations.ListOrganizationalUnitsForParentRequest

	ListParentsRequest(*organizations.ListParentsInput) organizations.ListParentsRequest

	ListPoliciesRequest(*organizations.ListPoliciesInput) organizations.ListPoliciesRequest

	ListPoliciesForTargetRequest(*organizations.ListPoliciesForTargetInput) organizations.ListPoliciesForTargetRequest

	ListRootsRequest(*organizations.ListRootsInput) organizations.ListRootsRequest

	ListTagsForResourceRequest(*organizations.ListTagsForResourceInput) organizations.ListTagsForResourceRequest

	ListTargetsForPolicyRequest(*organizations.ListTargetsForPolicyInput) organizations.ListTargetsForPolicyRequest

	MoveAccountRequest(*organizations.MoveAccountInput) organizations.MoveAccountRequest

	RegisterDelegatedAdministratorRequest(*organizations.RegisterDelegatedAdministratorInput) organizations.RegisterDelegatedAdministratorRequest

	RemoveAccountFromOrganizationRequest(*organizations.RemoveAccountFromOrganizationInput) organizations.RemoveAccountFromOrganizationRequest

	TagResourceRequest(*organizations.TagResourceInput) organizations.TagResourceRequest

	UntagResourceRequest(*organizations.UntagResourceInput) organizations.UntagResourceRequest

	UpdateOrganizationalUnitRequest(*organizations.UpdateOrganizationalUnitInput) organizations.UpdateOrganizationalUnitRequest

	UpdatePolicyRequest(*organizations.UpdatePolicyInput) organizations.UpdatePolicyRequest
}

var _ ClientAPI = (*organizations.Client)(nil)
