// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package appmeshiface provides an interface to enable mocking the AWS App Mesh service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package appmeshiface

import (
	"github.com/jviney/aws-sdk-go-v2/service/appmesh"
)

// ClientAPI provides an interface to enable mocking the
// appmesh.Client methods. This make unit testing your code that
// calls out to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the SDK's request pipeline.
//
//    // myFunc uses an SDK service client to make a request to
//    // AWS App Mesh.
//    func myFunc(svc appmeshiface.ClientAPI) bool {
//        // Make svc.CreateMesh request
//    }
//
//    func main() {
//        cfg, err := external.LoadDefaultAWSConfig()
//        if err != nil {
//            panic("failed to load config, " + err.Error())
//        }
//
//        svc := appmesh.New(cfg)
//
//        myFunc(svc)
//    }
//
// In your _test.go file:
//
//    // Define a mock struct to be used in your unit tests of myFunc.
//    type mockClientClient struct {
//        appmeshiface.ClientPI
//    }
//    func (m *mockClientClient) CreateMesh(input *appmesh.CreateMeshInput) (*appmesh.CreateMeshOutput, error) {
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
	CreateMeshRequest(*appmesh.CreateMeshInput) appmesh.CreateMeshRequest

	CreateRouteRequest(*appmesh.CreateRouteInput) appmesh.CreateRouteRequest

	CreateVirtualNodeRequest(*appmesh.CreateVirtualNodeInput) appmesh.CreateVirtualNodeRequest

	CreateVirtualRouterRequest(*appmesh.CreateVirtualRouterInput) appmesh.CreateVirtualRouterRequest

	CreateVirtualServiceRequest(*appmesh.CreateVirtualServiceInput) appmesh.CreateVirtualServiceRequest

	DeleteMeshRequest(*appmesh.DeleteMeshInput) appmesh.DeleteMeshRequest

	DeleteRouteRequest(*appmesh.DeleteRouteInput) appmesh.DeleteRouteRequest

	DeleteVirtualNodeRequest(*appmesh.DeleteVirtualNodeInput) appmesh.DeleteVirtualNodeRequest

	DeleteVirtualRouterRequest(*appmesh.DeleteVirtualRouterInput) appmesh.DeleteVirtualRouterRequest

	DeleteVirtualServiceRequest(*appmesh.DeleteVirtualServiceInput) appmesh.DeleteVirtualServiceRequest

	DescribeMeshRequest(*appmesh.DescribeMeshInput) appmesh.DescribeMeshRequest

	DescribeRouteRequest(*appmesh.DescribeRouteInput) appmesh.DescribeRouteRequest

	DescribeVirtualNodeRequest(*appmesh.DescribeVirtualNodeInput) appmesh.DescribeVirtualNodeRequest

	DescribeVirtualRouterRequest(*appmesh.DescribeVirtualRouterInput) appmesh.DescribeVirtualRouterRequest

	DescribeVirtualServiceRequest(*appmesh.DescribeVirtualServiceInput) appmesh.DescribeVirtualServiceRequest

	ListMeshesRequest(*appmesh.ListMeshesInput) appmesh.ListMeshesRequest

	ListRoutesRequest(*appmesh.ListRoutesInput) appmesh.ListRoutesRequest

	ListTagsForResourceRequest(*appmesh.ListTagsForResourceInput) appmesh.ListTagsForResourceRequest

	ListVirtualNodesRequest(*appmesh.ListVirtualNodesInput) appmesh.ListVirtualNodesRequest

	ListVirtualRoutersRequest(*appmesh.ListVirtualRoutersInput) appmesh.ListVirtualRoutersRequest

	ListVirtualServicesRequest(*appmesh.ListVirtualServicesInput) appmesh.ListVirtualServicesRequest

	TagResourceRequest(*appmesh.TagResourceInput) appmesh.TagResourceRequest

	UntagResourceRequest(*appmesh.UntagResourceInput) appmesh.UntagResourceRequest

	UpdateMeshRequest(*appmesh.UpdateMeshInput) appmesh.UpdateMeshRequest

	UpdateRouteRequest(*appmesh.UpdateRouteInput) appmesh.UpdateRouteRequest

	UpdateVirtualNodeRequest(*appmesh.UpdateVirtualNodeInput) appmesh.UpdateVirtualNodeRequest

	UpdateVirtualRouterRequest(*appmesh.UpdateVirtualRouterInput) appmesh.UpdateVirtualRouterRequest

	UpdateVirtualServiceRequest(*appmesh.UpdateVirtualServiceInput) appmesh.UpdateVirtualServiceRequest
}

var _ ClientAPI = (*appmesh.Client)(nil)
