// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package groundstationiface provides an interface to enable mocking the AWS Ground Station service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package groundstationiface

import (
	"github.com/jviney/aws-sdk-go-v2/service/groundstation"
)

// ClientAPI provides an interface to enable mocking the
// groundstation.Client methods. This make unit testing your code that
// calls out to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the SDK's request pipeline.
//
//    // myFunc uses an SDK service client to make a request to
//    // AWS Ground Station.
//    func myFunc(svc groundstationiface.ClientAPI) bool {
//        // Make svc.CancelContact request
//    }
//
//    func main() {
//        cfg, err := external.LoadDefaultAWSConfig()
//        if err != nil {
//            panic("failed to load config, " + err.Error())
//        }
//
//        svc := groundstation.New(cfg)
//
//        myFunc(svc)
//    }
//
// In your _test.go file:
//
//    // Define a mock struct to be used in your unit tests of myFunc.
//    type mockClientClient struct {
//        groundstationiface.ClientPI
//    }
//    func (m *mockClientClient) CancelContact(input *groundstation.CancelContactInput) (*groundstation.CancelContactOutput, error) {
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
	CancelContactRequest(*groundstation.CancelContactInput) groundstation.CancelContactRequest

	CreateConfigRequest(*groundstation.CreateConfigInput) groundstation.CreateConfigRequest

	CreateDataflowEndpointGroupRequest(*groundstation.CreateDataflowEndpointGroupInput) groundstation.CreateDataflowEndpointGroupRequest

	CreateMissionProfileRequest(*groundstation.CreateMissionProfileInput) groundstation.CreateMissionProfileRequest

	DeleteConfigRequest(*groundstation.DeleteConfigInput) groundstation.DeleteConfigRequest

	DeleteDataflowEndpointGroupRequest(*groundstation.DeleteDataflowEndpointGroupInput) groundstation.DeleteDataflowEndpointGroupRequest

	DeleteMissionProfileRequest(*groundstation.DeleteMissionProfileInput) groundstation.DeleteMissionProfileRequest

	DescribeContactRequest(*groundstation.DescribeContactInput) groundstation.DescribeContactRequest

	GetConfigRequest(*groundstation.GetConfigInput) groundstation.GetConfigRequest

	GetDataflowEndpointGroupRequest(*groundstation.GetDataflowEndpointGroupInput) groundstation.GetDataflowEndpointGroupRequest

	GetMinuteUsageRequest(*groundstation.GetMinuteUsageInput) groundstation.GetMinuteUsageRequest

	GetMissionProfileRequest(*groundstation.GetMissionProfileInput) groundstation.GetMissionProfileRequest

	GetSatelliteRequest(*groundstation.GetSatelliteInput) groundstation.GetSatelliteRequest

	ListConfigsRequest(*groundstation.ListConfigsInput) groundstation.ListConfigsRequest

	ListContactsRequest(*groundstation.ListContactsInput) groundstation.ListContactsRequest

	ListDataflowEndpointGroupsRequest(*groundstation.ListDataflowEndpointGroupsInput) groundstation.ListDataflowEndpointGroupsRequest

	ListGroundStationsRequest(*groundstation.ListGroundStationsInput) groundstation.ListGroundStationsRequest

	ListMissionProfilesRequest(*groundstation.ListMissionProfilesInput) groundstation.ListMissionProfilesRequest

	ListSatellitesRequest(*groundstation.ListSatellitesInput) groundstation.ListSatellitesRequest

	ListTagsForResourceRequest(*groundstation.ListTagsForResourceInput) groundstation.ListTagsForResourceRequest

	ReserveContactRequest(*groundstation.ReserveContactInput) groundstation.ReserveContactRequest

	TagResourceRequest(*groundstation.TagResourceInput) groundstation.TagResourceRequest

	UntagResourceRequest(*groundstation.UntagResourceInput) groundstation.UntagResourceRequest

	UpdateConfigRequest(*groundstation.UpdateConfigInput) groundstation.UpdateConfigRequest

	UpdateMissionProfileRequest(*groundstation.UpdateMissionProfileInput) groundstation.UpdateMissionProfileRequest
}

var _ ClientAPI = (*groundstation.Client)(nil)
