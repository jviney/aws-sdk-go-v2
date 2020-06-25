// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package storagegateway

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
)

// CancelRetrievalInput
type CancelRetrievalInput struct {
	_ struct{} `type:"structure"`

	// The Amazon Resource Name (ARN) of the gateway. Use the ListGateways operation
	// to return a list of gateways for your account and AWS Region.
	//
	// GatewayARN is a required field
	GatewayARN *string `min:"50" type:"string" required:"true"`

	// The Amazon Resource Name (ARN) of the virtual tape you want to cancel retrieval
	// for.
	//
	// TapeARN is a required field
	TapeARN *string `min:"50" type:"string" required:"true"`
}

// String returns the string representation
func (s CancelRetrievalInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CancelRetrievalInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CancelRetrievalInput"}

	if s.GatewayARN == nil {
		invalidParams.Add(aws.NewErrParamRequired("GatewayARN"))
	}
	if s.GatewayARN != nil && len(*s.GatewayARN) < 50 {
		invalidParams.Add(aws.NewErrParamMinLen("GatewayARN", 50))
	}

	if s.TapeARN == nil {
		invalidParams.Add(aws.NewErrParamRequired("TapeARN"))
	}
	if s.TapeARN != nil && len(*s.TapeARN) < 50 {
		invalidParams.Add(aws.NewErrParamMinLen("TapeARN", 50))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// CancelRetrievalOutput
type CancelRetrievalOutput struct {
	_ struct{} `type:"structure"`

	// The Amazon Resource Name (ARN) of the virtual tape for which retrieval was
	// canceled.
	TapeARN *string `min:"50" type:"string"`
}

// String returns the string representation
func (s CancelRetrievalOutput) String() string {
	return awsutil.Prettify(s)
}

const opCancelRetrieval = "CancelRetrieval"

// CancelRetrievalRequest returns a request value for making API operation for
// AWS Storage Gateway.
//
// Cancels retrieval of a virtual tape from the virtual tape shelf (VTS) to
// a gateway after the retrieval process is initiated. The virtual tape is returned
// to the VTS. This operation is only supported in the tape gateway type.
//
//    // Example sending a request using CancelRetrievalRequest.
//    req := client.CancelRetrievalRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/storagegateway-2013-06-30/CancelRetrieval
func (c *Client) CancelRetrievalRequest(input *CancelRetrievalInput) CancelRetrievalRequest {
	op := &aws.Operation{
		Name:       opCancelRetrieval,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &CancelRetrievalInput{}
	}

	req := c.newRequest(op, input, &CancelRetrievalOutput{})

	return CancelRetrievalRequest{Request: req, Input: input, Copy: c.CancelRetrievalRequest}
}

// CancelRetrievalRequest is the request type for the
// CancelRetrieval API operation.
type CancelRetrievalRequest struct {
	*aws.Request
	Input *CancelRetrievalInput
	Copy  func(*CancelRetrievalInput) CancelRetrievalRequest
}

// Send marshals and sends the CancelRetrieval API request.
func (r CancelRetrievalRequest) Send(ctx context.Context) (*CancelRetrievalResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CancelRetrievalResponse{
		CancelRetrievalOutput: r.Request.Data.(*CancelRetrievalOutput),
		response:              &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CancelRetrievalResponse is the response type for the
// CancelRetrieval API operation.
type CancelRetrievalResponse struct {
	*CancelRetrievalOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CancelRetrieval request.
func (r *CancelRetrievalResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
