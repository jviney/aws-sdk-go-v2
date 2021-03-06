// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package quicksightiface provides an interface to enable mocking the Amazon QuickSight service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package quicksightiface

import (
	"github.com/jviney/aws-sdk-go-v2/service/quicksight"
)

// ClientAPI provides an interface to enable mocking the
// quicksight.Client methods. This make unit testing your code that
// calls out to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the SDK's request pipeline.
//
//    // myFunc uses an SDK service client to make a request to
//    // Amazon QuickSight.
//    func myFunc(svc quicksightiface.ClientAPI) bool {
//        // Make svc.CancelIngestion request
//    }
//
//    func main() {
//        cfg, err := external.LoadDefaultAWSConfig()
//        if err != nil {
//            panic("failed to load config, " + err.Error())
//        }
//
//        svc := quicksight.New(cfg)
//
//        myFunc(svc)
//    }
//
// In your _test.go file:
//
//    // Define a mock struct to be used in your unit tests of myFunc.
//    type mockClientClient struct {
//        quicksightiface.ClientPI
//    }
//    func (m *mockClientClient) CancelIngestion(input *quicksight.CancelIngestionInput) (*quicksight.CancelIngestionOutput, error) {
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
	CancelIngestionRequest(*quicksight.CancelIngestionInput) quicksight.CancelIngestionRequest

	CreateDashboardRequest(*quicksight.CreateDashboardInput) quicksight.CreateDashboardRequest

	CreateDataSetRequest(*quicksight.CreateDataSetInput) quicksight.CreateDataSetRequest

	CreateDataSourceRequest(*quicksight.CreateDataSourceInput) quicksight.CreateDataSourceRequest

	CreateGroupRequest(*quicksight.CreateGroupInput) quicksight.CreateGroupRequest

	CreateGroupMembershipRequest(*quicksight.CreateGroupMembershipInput) quicksight.CreateGroupMembershipRequest

	CreateIAMPolicyAssignmentRequest(*quicksight.CreateIAMPolicyAssignmentInput) quicksight.CreateIAMPolicyAssignmentRequest

	CreateIngestionRequest(*quicksight.CreateIngestionInput) quicksight.CreateIngestionRequest

	CreateTemplateRequest(*quicksight.CreateTemplateInput) quicksight.CreateTemplateRequest

	CreateTemplateAliasRequest(*quicksight.CreateTemplateAliasInput) quicksight.CreateTemplateAliasRequest

	DeleteDashboardRequest(*quicksight.DeleteDashboardInput) quicksight.DeleteDashboardRequest

	DeleteDataSetRequest(*quicksight.DeleteDataSetInput) quicksight.DeleteDataSetRequest

	DeleteDataSourceRequest(*quicksight.DeleteDataSourceInput) quicksight.DeleteDataSourceRequest

	DeleteGroupRequest(*quicksight.DeleteGroupInput) quicksight.DeleteGroupRequest

	DeleteGroupMembershipRequest(*quicksight.DeleteGroupMembershipInput) quicksight.DeleteGroupMembershipRequest

	DeleteIAMPolicyAssignmentRequest(*quicksight.DeleteIAMPolicyAssignmentInput) quicksight.DeleteIAMPolicyAssignmentRequest

	DeleteTemplateRequest(*quicksight.DeleteTemplateInput) quicksight.DeleteTemplateRequest

	DeleteTemplateAliasRequest(*quicksight.DeleteTemplateAliasInput) quicksight.DeleteTemplateAliasRequest

	DeleteUserRequest(*quicksight.DeleteUserInput) quicksight.DeleteUserRequest

	DeleteUserByPrincipalIdRequest(*quicksight.DeleteUserByPrincipalIdInput) quicksight.DeleteUserByPrincipalIdRequest

	DescribeDashboardRequest(*quicksight.DescribeDashboardInput) quicksight.DescribeDashboardRequest

	DescribeDashboardPermissionsRequest(*quicksight.DescribeDashboardPermissionsInput) quicksight.DescribeDashboardPermissionsRequest

	DescribeDataSetRequest(*quicksight.DescribeDataSetInput) quicksight.DescribeDataSetRequest

	DescribeDataSetPermissionsRequest(*quicksight.DescribeDataSetPermissionsInput) quicksight.DescribeDataSetPermissionsRequest

	DescribeDataSourceRequest(*quicksight.DescribeDataSourceInput) quicksight.DescribeDataSourceRequest

	DescribeDataSourcePermissionsRequest(*quicksight.DescribeDataSourcePermissionsInput) quicksight.DescribeDataSourcePermissionsRequest

	DescribeGroupRequest(*quicksight.DescribeGroupInput) quicksight.DescribeGroupRequest

	DescribeIAMPolicyAssignmentRequest(*quicksight.DescribeIAMPolicyAssignmentInput) quicksight.DescribeIAMPolicyAssignmentRequest

	DescribeIngestionRequest(*quicksight.DescribeIngestionInput) quicksight.DescribeIngestionRequest

	DescribeTemplateRequest(*quicksight.DescribeTemplateInput) quicksight.DescribeTemplateRequest

	DescribeTemplateAliasRequest(*quicksight.DescribeTemplateAliasInput) quicksight.DescribeTemplateAliasRequest

	DescribeTemplatePermissionsRequest(*quicksight.DescribeTemplatePermissionsInput) quicksight.DescribeTemplatePermissionsRequest

	DescribeUserRequest(*quicksight.DescribeUserInput) quicksight.DescribeUserRequest

	GetDashboardEmbedUrlRequest(*quicksight.GetDashboardEmbedUrlInput) quicksight.GetDashboardEmbedUrlRequest

	ListDashboardVersionsRequest(*quicksight.ListDashboardVersionsInput) quicksight.ListDashboardVersionsRequest

	ListDashboardsRequest(*quicksight.ListDashboardsInput) quicksight.ListDashboardsRequest

	ListDataSetsRequest(*quicksight.ListDataSetsInput) quicksight.ListDataSetsRequest

	ListDataSourcesRequest(*quicksight.ListDataSourcesInput) quicksight.ListDataSourcesRequest

	ListGroupMembershipsRequest(*quicksight.ListGroupMembershipsInput) quicksight.ListGroupMembershipsRequest

	ListGroupsRequest(*quicksight.ListGroupsInput) quicksight.ListGroupsRequest

	ListIAMPolicyAssignmentsRequest(*quicksight.ListIAMPolicyAssignmentsInput) quicksight.ListIAMPolicyAssignmentsRequest

	ListIAMPolicyAssignmentsForUserRequest(*quicksight.ListIAMPolicyAssignmentsForUserInput) quicksight.ListIAMPolicyAssignmentsForUserRequest

	ListIngestionsRequest(*quicksight.ListIngestionsInput) quicksight.ListIngestionsRequest

	ListTagsForResourceRequest(*quicksight.ListTagsForResourceInput) quicksight.ListTagsForResourceRequest

	ListTemplateAliasesRequest(*quicksight.ListTemplateAliasesInput) quicksight.ListTemplateAliasesRequest

	ListTemplateVersionsRequest(*quicksight.ListTemplateVersionsInput) quicksight.ListTemplateVersionsRequest

	ListTemplatesRequest(*quicksight.ListTemplatesInput) quicksight.ListTemplatesRequest

	ListUserGroupsRequest(*quicksight.ListUserGroupsInput) quicksight.ListUserGroupsRequest

	ListUsersRequest(*quicksight.ListUsersInput) quicksight.ListUsersRequest

	RegisterUserRequest(*quicksight.RegisterUserInput) quicksight.RegisterUserRequest

	SearchDashboardsRequest(*quicksight.SearchDashboardsInput) quicksight.SearchDashboardsRequest

	TagResourceRequest(*quicksight.TagResourceInput) quicksight.TagResourceRequest

	UntagResourceRequest(*quicksight.UntagResourceInput) quicksight.UntagResourceRequest

	UpdateDashboardRequest(*quicksight.UpdateDashboardInput) quicksight.UpdateDashboardRequest

	UpdateDashboardPermissionsRequest(*quicksight.UpdateDashboardPermissionsInput) quicksight.UpdateDashboardPermissionsRequest

	UpdateDashboardPublishedVersionRequest(*quicksight.UpdateDashboardPublishedVersionInput) quicksight.UpdateDashboardPublishedVersionRequest

	UpdateDataSetRequest(*quicksight.UpdateDataSetInput) quicksight.UpdateDataSetRequest

	UpdateDataSetPermissionsRequest(*quicksight.UpdateDataSetPermissionsInput) quicksight.UpdateDataSetPermissionsRequest

	UpdateDataSourceRequest(*quicksight.UpdateDataSourceInput) quicksight.UpdateDataSourceRequest

	UpdateDataSourcePermissionsRequest(*quicksight.UpdateDataSourcePermissionsInput) quicksight.UpdateDataSourcePermissionsRequest

	UpdateGroupRequest(*quicksight.UpdateGroupInput) quicksight.UpdateGroupRequest

	UpdateIAMPolicyAssignmentRequest(*quicksight.UpdateIAMPolicyAssignmentInput) quicksight.UpdateIAMPolicyAssignmentRequest

	UpdateTemplateRequest(*quicksight.UpdateTemplateInput) quicksight.UpdateTemplateRequest

	UpdateTemplateAliasRequest(*quicksight.UpdateTemplateAliasInput) quicksight.UpdateTemplateAliasRequest

	UpdateTemplatePermissionsRequest(*quicksight.UpdateTemplatePermissionsInput) quicksight.UpdateTemplatePermissionsRequest

	UpdateUserRequest(*quicksight.UpdateUserInput) quicksight.UpdateUserRequest
}

var _ ClientAPI = (*quicksight.Client)(nil)
