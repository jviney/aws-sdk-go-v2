// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ec2

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
)

// Contains the parameters for CancelSpotInstanceRequests.
type CancelSpotInstanceRequestsInput struct {
	_ struct{} `type:"structure"`

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have
	// the required permissions, the error response is DryRunOperation. Otherwise,
	// it is UnauthorizedOperation.
	DryRun *bool `locationName:"dryRun" type:"boolean"`

	// One or more Spot Instance request IDs.
	//
	// SpotInstanceRequestIds is a required field
	SpotInstanceRequestIds []string `locationName:"SpotInstanceRequestId" locationNameList:"SpotInstanceRequestId" type:"list" required:"true"`
}

// String returns the string representation
func (s CancelSpotInstanceRequestsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CancelSpotInstanceRequestsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CancelSpotInstanceRequestsInput"}

	if s.SpotInstanceRequestIds == nil {
		invalidParams.Add(aws.NewErrParamRequired("SpotInstanceRequestIds"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Contains the output of CancelSpotInstanceRequests.
type CancelSpotInstanceRequestsOutput struct {
	_ struct{} `type:"structure"`

	// One or more Spot Instance requests.
	CancelledSpotInstanceRequests []CancelledSpotInstanceRequest `locationName:"spotInstanceRequestSet" locationNameList:"item" type:"list"`
}

// String returns the string representation
func (s CancelSpotInstanceRequestsOutput) String() string {
	return awsutil.Prettify(s)
}

const opCancelSpotInstanceRequests = "CancelSpotInstanceRequests"

// CancelSpotInstanceRequestsRequest returns a request value for making API operation for
// Amazon Elastic Compute Cloud.
//
// Cancels one or more Spot Instance requests.
//
// Canceling a Spot Instance request does not terminate running Spot Instances
// associated with the request.
//
//    // Example sending a request using CancelSpotInstanceRequestsRequest.
//    req := client.CancelSpotInstanceRequestsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/CancelSpotInstanceRequests
func (c *Client) CancelSpotInstanceRequestsRequest(input *CancelSpotInstanceRequestsInput) CancelSpotInstanceRequestsRequest {
	op := &aws.Operation{
		Name:       opCancelSpotInstanceRequests,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &CancelSpotInstanceRequestsInput{}
	}

	req := c.newRequest(op, input, &CancelSpotInstanceRequestsOutput{})

	return CancelSpotInstanceRequestsRequest{Request: req, Input: input, Copy: c.CancelSpotInstanceRequestsRequest}
}

// CancelSpotInstanceRequestsRequest is the request type for the
// CancelSpotInstanceRequests API operation.
type CancelSpotInstanceRequestsRequest struct {
	*aws.Request
	Input *CancelSpotInstanceRequestsInput
	Copy  func(*CancelSpotInstanceRequestsInput) CancelSpotInstanceRequestsRequest
}

// Send marshals and sends the CancelSpotInstanceRequests API request.
func (r CancelSpotInstanceRequestsRequest) Send(ctx context.Context) (*CancelSpotInstanceRequestsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CancelSpotInstanceRequestsResponse{
		CancelSpotInstanceRequestsOutput: r.Request.Data.(*CancelSpotInstanceRequestsOutput),
		response:                         &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CancelSpotInstanceRequestsResponse is the response type for the
// CancelSpotInstanceRequests API operation.
type CancelSpotInstanceRequestsResponse struct {
	*CancelSpotInstanceRequestsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CancelSpotInstanceRequests request.
func (r *CancelSpotInstanceRequestsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
