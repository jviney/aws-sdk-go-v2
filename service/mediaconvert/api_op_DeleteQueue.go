// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package mediaconvert

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
	"github.com/jviney/aws-sdk-go-v2/private/protocol"
)

// Delete a queue by sending a request with the queue name. You can't delete
// a queue with an active pricing plan or one that has unprocessed jobs in it.
type DeleteQueueInput struct {
	_ struct{} `type:"structure"`

	// The name of the queue that you want to delete.
	//
	// Name is a required field
	Name *string `location:"uri" locationName:"name" type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteQueueInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteQueueInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteQueueInput"}

	if s.Name == nil {
		invalidParams.Add(aws.NewErrParamRequired("Name"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DeleteQueueInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.Name != nil {
		v := *s.Name

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "name", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// Delete queue requests return an OK message or error message with an empty
// body.
type DeleteQueueOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DeleteQueueOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DeleteQueueOutput) MarshalFields(e protocol.FieldEncoder) error {
	return nil
}

const opDeleteQueue = "DeleteQueue"

// DeleteQueueRequest returns a request value for making API operation for
// AWS Elemental MediaConvert.
//
// Permanently delete a queue you have created.
//
//    // Example sending a request using DeleteQueueRequest.
//    req := client.DeleteQueueRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/mediaconvert-2017-08-29/DeleteQueue
func (c *Client) DeleteQueueRequest(input *DeleteQueueInput) DeleteQueueRequest {
	op := &aws.Operation{
		Name:       opDeleteQueue,
		HTTPMethod: "DELETE",
		HTTPPath:   "/2017-08-29/queues/{name}",
	}

	if input == nil {
		input = &DeleteQueueInput{}
	}

	req := c.newRequest(op, input, &DeleteQueueOutput{})

	return DeleteQueueRequest{Request: req, Input: input, Copy: c.DeleteQueueRequest}
}

// DeleteQueueRequest is the request type for the
// DeleteQueue API operation.
type DeleteQueueRequest struct {
	*aws.Request
	Input *DeleteQueueInput
	Copy  func(*DeleteQueueInput) DeleteQueueRequest
}

// Send marshals and sends the DeleteQueue API request.
func (r DeleteQueueRequest) Send(ctx context.Context) (*DeleteQueueResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteQueueResponse{
		DeleteQueueOutput: r.Request.Data.(*DeleteQueueOutput),
		response:          &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteQueueResponse is the response type for the
// DeleteQueue API operation.
type DeleteQueueResponse struct {
	*DeleteQueueOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteQueue request.
func (r *DeleteQueueResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
