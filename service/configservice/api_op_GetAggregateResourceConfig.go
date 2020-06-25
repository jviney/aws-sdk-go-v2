// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package configservice

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
)

type GetAggregateResourceConfigInput struct {
	_ struct{} `type:"structure"`

	// The name of the configuration aggregator.
	//
	// ConfigurationAggregatorName is a required field
	ConfigurationAggregatorName *string `min:"1" type:"string" required:"true"`

	// An object that identifies aggregate resource.
	//
	// ResourceIdentifier is a required field
	ResourceIdentifier *AggregateResourceIdentifier `type:"structure" required:"true"`
}

// String returns the string representation
func (s GetAggregateResourceConfigInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetAggregateResourceConfigInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetAggregateResourceConfigInput"}

	if s.ConfigurationAggregatorName == nil {
		invalidParams.Add(aws.NewErrParamRequired("ConfigurationAggregatorName"))
	}
	if s.ConfigurationAggregatorName != nil && len(*s.ConfigurationAggregatorName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ConfigurationAggregatorName", 1))
	}

	if s.ResourceIdentifier == nil {
		invalidParams.Add(aws.NewErrParamRequired("ResourceIdentifier"))
	}
	if s.ResourceIdentifier != nil {
		if err := s.ResourceIdentifier.Validate(); err != nil {
			invalidParams.AddNested("ResourceIdentifier", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type GetAggregateResourceConfigOutput struct {
	_ struct{} `type:"structure"`

	// Returns a ConfigurationItem object.
	ConfigurationItem *ConfigurationItem `type:"structure"`
}

// String returns the string representation
func (s GetAggregateResourceConfigOutput) String() string {
	return awsutil.Prettify(s)
}

const opGetAggregateResourceConfig = "GetAggregateResourceConfig"

// GetAggregateResourceConfigRequest returns a request value for making API operation for
// AWS Config.
//
// Returns configuration item that is aggregated for your specific resource
// in a specific source account and region.
//
//    // Example sending a request using GetAggregateResourceConfigRequest.
//    req := client.GetAggregateResourceConfigRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/config-2014-11-12/GetAggregateResourceConfig
func (c *Client) GetAggregateResourceConfigRequest(input *GetAggregateResourceConfigInput) GetAggregateResourceConfigRequest {
	op := &aws.Operation{
		Name:       opGetAggregateResourceConfig,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &GetAggregateResourceConfigInput{}
	}

	req := c.newRequest(op, input, &GetAggregateResourceConfigOutput{})

	return GetAggregateResourceConfigRequest{Request: req, Input: input, Copy: c.GetAggregateResourceConfigRequest}
}

// GetAggregateResourceConfigRequest is the request type for the
// GetAggregateResourceConfig API operation.
type GetAggregateResourceConfigRequest struct {
	*aws.Request
	Input *GetAggregateResourceConfigInput
	Copy  func(*GetAggregateResourceConfigInput) GetAggregateResourceConfigRequest
}

// Send marshals and sends the GetAggregateResourceConfig API request.
func (r GetAggregateResourceConfigRequest) Send(ctx context.Context) (*GetAggregateResourceConfigResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetAggregateResourceConfigResponse{
		GetAggregateResourceConfigOutput: r.Request.Data.(*GetAggregateResourceConfigOutput),
		response:                         &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetAggregateResourceConfigResponse is the response type for the
// GetAggregateResourceConfig API operation.
type GetAggregateResourceConfigResponse struct {
	*GetAggregateResourceConfigOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetAggregateResourceConfig request.
func (r *GetAggregateResourceConfigResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
