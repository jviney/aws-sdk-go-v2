// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package amplifyiface provides an interface to enable mocking the AWS Amplify service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package amplifyiface

import (
	"github.com/jviney/aws-sdk-go-v2/service/amplify"
)

// ClientAPI provides an interface to enable mocking the
// amplify.Client methods. This make unit testing your code that
// calls out to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the SDK's request pipeline.
//
//    // myFunc uses an SDK service client to make a request to
//    // Amplify.
//    func myFunc(svc amplifyiface.ClientAPI) bool {
//        // Make svc.CreateApp request
//    }
//
//    func main() {
//        cfg, err := external.LoadDefaultAWSConfig()
//        if err != nil {
//            panic("failed to load config, " + err.Error())
//        }
//
//        svc := amplify.New(cfg)
//
//        myFunc(svc)
//    }
//
// In your _test.go file:
//
//    // Define a mock struct to be used in your unit tests of myFunc.
//    type mockClientClient struct {
//        amplifyiface.ClientPI
//    }
//    func (m *mockClientClient) CreateApp(input *amplify.CreateAppInput) (*amplify.CreateAppOutput, error) {
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
	CreateAppRequest(*amplify.CreateAppInput) amplify.CreateAppRequest

	CreateBackendEnvironmentRequest(*amplify.CreateBackendEnvironmentInput) amplify.CreateBackendEnvironmentRequest

	CreateBranchRequest(*amplify.CreateBranchInput) amplify.CreateBranchRequest

	CreateDeploymentRequest(*amplify.CreateDeploymentInput) amplify.CreateDeploymentRequest

	CreateDomainAssociationRequest(*amplify.CreateDomainAssociationInput) amplify.CreateDomainAssociationRequest

	CreateWebhookRequest(*amplify.CreateWebhookInput) amplify.CreateWebhookRequest

	DeleteAppRequest(*amplify.DeleteAppInput) amplify.DeleteAppRequest

	DeleteBackendEnvironmentRequest(*amplify.DeleteBackendEnvironmentInput) amplify.DeleteBackendEnvironmentRequest

	DeleteBranchRequest(*amplify.DeleteBranchInput) amplify.DeleteBranchRequest

	DeleteDomainAssociationRequest(*amplify.DeleteDomainAssociationInput) amplify.DeleteDomainAssociationRequest

	DeleteJobRequest(*amplify.DeleteJobInput) amplify.DeleteJobRequest

	DeleteWebhookRequest(*amplify.DeleteWebhookInput) amplify.DeleteWebhookRequest

	GenerateAccessLogsRequest(*amplify.GenerateAccessLogsInput) amplify.GenerateAccessLogsRequest

	GetAppRequest(*amplify.GetAppInput) amplify.GetAppRequest

	GetArtifactUrlRequest(*amplify.GetArtifactUrlInput) amplify.GetArtifactUrlRequest

	GetBackendEnvironmentRequest(*amplify.GetBackendEnvironmentInput) amplify.GetBackendEnvironmentRequest

	GetBranchRequest(*amplify.GetBranchInput) amplify.GetBranchRequest

	GetDomainAssociationRequest(*amplify.GetDomainAssociationInput) amplify.GetDomainAssociationRequest

	GetJobRequest(*amplify.GetJobInput) amplify.GetJobRequest

	GetWebhookRequest(*amplify.GetWebhookInput) amplify.GetWebhookRequest

	ListAppsRequest(*amplify.ListAppsInput) amplify.ListAppsRequest

	ListArtifactsRequest(*amplify.ListArtifactsInput) amplify.ListArtifactsRequest

	ListBackendEnvironmentsRequest(*amplify.ListBackendEnvironmentsInput) amplify.ListBackendEnvironmentsRequest

	ListBranchesRequest(*amplify.ListBranchesInput) amplify.ListBranchesRequest

	ListDomainAssociationsRequest(*amplify.ListDomainAssociationsInput) amplify.ListDomainAssociationsRequest

	ListJobsRequest(*amplify.ListJobsInput) amplify.ListJobsRequest

	ListTagsForResourceRequest(*amplify.ListTagsForResourceInput) amplify.ListTagsForResourceRequest

	ListWebhooksRequest(*amplify.ListWebhooksInput) amplify.ListWebhooksRequest

	StartDeploymentRequest(*amplify.StartDeploymentInput) amplify.StartDeploymentRequest

	StartJobRequest(*amplify.StartJobInput) amplify.StartJobRequest

	StopJobRequest(*amplify.StopJobInput) amplify.StopJobRequest

	TagResourceRequest(*amplify.TagResourceInput) amplify.TagResourceRequest

	UntagResourceRequest(*amplify.UntagResourceInput) amplify.UntagResourceRequest

	UpdateAppRequest(*amplify.UpdateAppInput) amplify.UpdateAppRequest

	UpdateBranchRequest(*amplify.UpdateBranchInput) amplify.UpdateBranchRequest

	UpdateDomainAssociationRequest(*amplify.UpdateDomainAssociationInput) amplify.UpdateDomainAssociationRequest

	UpdateWebhookRequest(*amplify.UpdateWebhookInput) amplify.UpdateWebhookRequest
}

var _ ClientAPI = (*amplify.Client)(nil)
