// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package configservice

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
	"github.com/jviney/aws-sdk-go-v2/private/protocol"
	"github.com/jviney/aws-sdk-go-v2/private/protocol/jsonrpc"
)

type DeleteOrganizationConformancePackInput struct {
	_ struct{} `type:"structure"`

	// The name of organization conformance pack that you want to delete.
	//
	// OrganizationConformancePackName is a required field
	OrganizationConformancePackName *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteOrganizationConformancePackInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteOrganizationConformancePackInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteOrganizationConformancePackInput"}

	if s.OrganizationConformancePackName == nil {
		invalidParams.Add(aws.NewErrParamRequired("OrganizationConformancePackName"))
	}
	if s.OrganizationConformancePackName != nil && len(*s.OrganizationConformancePackName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("OrganizationConformancePackName", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DeleteOrganizationConformancePackOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DeleteOrganizationConformancePackOutput) String() string {
	return awsutil.Prettify(s)
}

const opDeleteOrganizationConformancePack = "DeleteOrganizationConformancePack"

// DeleteOrganizationConformancePackRequest returns a request value for making API operation for
// AWS Config.
//
// Deletes the specified organization conformance pack and all of the config
// rules and remediation actions from all member accounts in that organization.
// Only a master account can delete an organization conformance pack.
//
// AWS Config sets the state of a conformance pack to DELETE_IN_PROGRESS until
// the deletion is complete. You cannot update a conformance pack while it is
// in this state.
//
//    // Example sending a request using DeleteOrganizationConformancePackRequest.
//    req := client.DeleteOrganizationConformancePackRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/config-2014-11-12/DeleteOrganizationConformancePack
func (c *Client) DeleteOrganizationConformancePackRequest(input *DeleteOrganizationConformancePackInput) DeleteOrganizationConformancePackRequest {
	op := &aws.Operation{
		Name:       opDeleteOrganizationConformancePack,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DeleteOrganizationConformancePackInput{}
	}

	req := c.newRequest(op, input, &DeleteOrganizationConformancePackOutput{})
	req.Handlers.Unmarshal.Remove(jsonrpc.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)

	return DeleteOrganizationConformancePackRequest{Request: req, Input: input, Copy: c.DeleteOrganizationConformancePackRequest}
}

// DeleteOrganizationConformancePackRequest is the request type for the
// DeleteOrganizationConformancePack API operation.
type DeleteOrganizationConformancePackRequest struct {
	*aws.Request
	Input *DeleteOrganizationConformancePackInput
	Copy  func(*DeleteOrganizationConformancePackInput) DeleteOrganizationConformancePackRequest
}

// Send marshals and sends the DeleteOrganizationConformancePack API request.
func (r DeleteOrganizationConformancePackRequest) Send(ctx context.Context) (*DeleteOrganizationConformancePackResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteOrganizationConformancePackResponse{
		DeleteOrganizationConformancePackOutput: r.Request.Data.(*DeleteOrganizationConformancePackOutput),
		response:                                &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteOrganizationConformancePackResponse is the response type for the
// DeleteOrganizationConformancePack API operation.
type DeleteOrganizationConformancePackResponse struct {
	*DeleteOrganizationConformancePackOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteOrganizationConformancePack request.
func (r *DeleteOrganizationConformancePackResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
