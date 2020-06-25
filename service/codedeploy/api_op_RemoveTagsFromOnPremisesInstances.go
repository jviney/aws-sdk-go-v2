// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package codedeploy

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
	"github.com/jviney/aws-sdk-go-v2/private/protocol"
	"github.com/jviney/aws-sdk-go-v2/private/protocol/jsonrpc"
)

// Represents the input of a RemoveTagsFromOnPremisesInstances operation.
type RemoveTagsFromOnPremisesInstancesInput struct {
	_ struct{} `type:"structure"`

	// The names of the on-premises instances from which to remove tags.
	//
	// InstanceNames is a required field
	InstanceNames []string `locationName:"instanceNames" type:"list" required:"true"`

	// The tag key-value pairs to remove from the on-premises instances.
	//
	// Tags is a required field
	Tags []Tag `locationName:"tags" type:"list" required:"true"`
}

// String returns the string representation
func (s RemoveTagsFromOnPremisesInstancesInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *RemoveTagsFromOnPremisesInstancesInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "RemoveTagsFromOnPremisesInstancesInput"}

	if s.InstanceNames == nil {
		invalidParams.Add(aws.NewErrParamRequired("InstanceNames"))
	}

	if s.Tags == nil {
		invalidParams.Add(aws.NewErrParamRequired("Tags"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type RemoveTagsFromOnPremisesInstancesOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s RemoveTagsFromOnPremisesInstancesOutput) String() string {
	return awsutil.Prettify(s)
}

const opRemoveTagsFromOnPremisesInstances = "RemoveTagsFromOnPremisesInstances"

// RemoveTagsFromOnPremisesInstancesRequest returns a request value for making API operation for
// AWS CodeDeploy.
//
// Removes one or more tags from one or more on-premises instances.
//
//    // Example sending a request using RemoveTagsFromOnPremisesInstancesRequest.
//    req := client.RemoveTagsFromOnPremisesInstancesRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/codedeploy-2014-10-06/RemoveTagsFromOnPremisesInstances
func (c *Client) RemoveTagsFromOnPremisesInstancesRequest(input *RemoveTagsFromOnPremisesInstancesInput) RemoveTagsFromOnPremisesInstancesRequest {
	op := &aws.Operation{
		Name:       opRemoveTagsFromOnPremisesInstances,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &RemoveTagsFromOnPremisesInstancesInput{}
	}

	req := c.newRequest(op, input, &RemoveTagsFromOnPremisesInstancesOutput{})
	req.Handlers.Unmarshal.Remove(jsonrpc.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)

	return RemoveTagsFromOnPremisesInstancesRequest{Request: req, Input: input, Copy: c.RemoveTagsFromOnPremisesInstancesRequest}
}

// RemoveTagsFromOnPremisesInstancesRequest is the request type for the
// RemoveTagsFromOnPremisesInstances API operation.
type RemoveTagsFromOnPremisesInstancesRequest struct {
	*aws.Request
	Input *RemoveTagsFromOnPremisesInstancesInput
	Copy  func(*RemoveTagsFromOnPremisesInstancesInput) RemoveTagsFromOnPremisesInstancesRequest
}

// Send marshals and sends the RemoveTagsFromOnPremisesInstances API request.
func (r RemoveTagsFromOnPremisesInstancesRequest) Send(ctx context.Context) (*RemoveTagsFromOnPremisesInstancesResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &RemoveTagsFromOnPremisesInstancesResponse{
		RemoveTagsFromOnPremisesInstancesOutput: r.Request.Data.(*RemoveTagsFromOnPremisesInstancesOutput),
		response:                                &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// RemoveTagsFromOnPremisesInstancesResponse is the response type for the
// RemoveTagsFromOnPremisesInstances API operation.
type RemoveTagsFromOnPremisesInstancesResponse struct {
	*RemoveTagsFromOnPremisesInstancesOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// RemoveTagsFromOnPremisesInstances request.
func (r *RemoveTagsFromOnPremisesInstancesResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
