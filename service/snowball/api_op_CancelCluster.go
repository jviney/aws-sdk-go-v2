// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package snowball

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
)

type CancelClusterInput struct {
	_ struct{} `type:"structure"`

	// The 39-character ID for the cluster that you want to cancel, for example
	// CID123e4567-e89b-12d3-a456-426655440000.
	//
	// ClusterId is a required field
	ClusterId *string `min:"39" type:"string" required:"true"`
}

// String returns the string representation
func (s CancelClusterInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CancelClusterInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CancelClusterInput"}

	if s.ClusterId == nil {
		invalidParams.Add(aws.NewErrParamRequired("ClusterId"))
	}
	if s.ClusterId != nil && len(*s.ClusterId) < 39 {
		invalidParams.Add(aws.NewErrParamMinLen("ClusterId", 39))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type CancelClusterOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s CancelClusterOutput) String() string {
	return awsutil.Prettify(s)
}

const opCancelCluster = "CancelCluster"

// CancelClusterRequest returns a request value for making API operation for
// Amazon Import/Export Snowball.
//
// Cancels a cluster job. You can only cancel a cluster job while it's in the
// AwaitingQuorum status. You'll have at least an hour after creating a cluster
// job to cancel it.
//
//    // Example sending a request using CancelClusterRequest.
//    req := client.CancelClusterRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/snowball-2016-06-30/CancelCluster
func (c *Client) CancelClusterRequest(input *CancelClusterInput) CancelClusterRequest {
	op := &aws.Operation{
		Name:       opCancelCluster,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &CancelClusterInput{}
	}

	req := c.newRequest(op, input, &CancelClusterOutput{})

	return CancelClusterRequest{Request: req, Input: input, Copy: c.CancelClusterRequest}
}

// CancelClusterRequest is the request type for the
// CancelCluster API operation.
type CancelClusterRequest struct {
	*aws.Request
	Input *CancelClusterInput
	Copy  func(*CancelClusterInput) CancelClusterRequest
}

// Send marshals and sends the CancelCluster API request.
func (r CancelClusterRequest) Send(ctx context.Context) (*CancelClusterResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CancelClusterResponse{
		CancelClusterOutput: r.Request.Data.(*CancelClusterOutput),
		response:            &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CancelClusterResponse is the response type for the
// CancelCluster API operation.
type CancelClusterResponse struct {
	*CancelClusterOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CancelCluster request.
func (r *CancelClusterResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
