// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ses

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
)

// Represents a request to create a configuration set. Configuration sets enable
// you to publish email sending events. For information about using configuration
// sets, see the Amazon SES Developer Guide (https://docs.aws.amazon.com/ses/latest/DeveloperGuide/monitor-sending-activity.html).
type CreateConfigurationSetInput struct {
	_ struct{} `type:"structure"`

	// A data structure that contains the name of the configuration set.
	//
	// ConfigurationSet is a required field
	ConfigurationSet *ConfigurationSet `type:"structure" required:"true"`
}

// String returns the string representation
func (s CreateConfigurationSetInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateConfigurationSetInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateConfigurationSetInput"}

	if s.ConfigurationSet == nil {
		invalidParams.Add(aws.NewErrParamRequired("ConfigurationSet"))
	}
	if s.ConfigurationSet != nil {
		if err := s.ConfigurationSet.Validate(); err != nil {
			invalidParams.AddNested("ConfigurationSet", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// An empty element returned on a successful request.
type CreateConfigurationSetOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s CreateConfigurationSetOutput) String() string {
	return awsutil.Prettify(s)
}

const opCreateConfigurationSet = "CreateConfigurationSet"

// CreateConfigurationSetRequest returns a request value for making API operation for
// Amazon Simple Email Service.
//
// Creates a configuration set.
//
// Configuration sets enable you to publish email sending events. For information
// about using configuration sets, see the Amazon SES Developer Guide (https://docs.aws.amazon.com/ses/latest/DeveloperGuide/monitor-sending-activity.html).
//
// You can execute this operation no more than once per second.
//
//    // Example sending a request using CreateConfigurationSetRequest.
//    req := client.CreateConfigurationSetRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/email-2010-12-01/CreateConfigurationSet
func (c *Client) CreateConfigurationSetRequest(input *CreateConfigurationSetInput) CreateConfigurationSetRequest {
	op := &aws.Operation{
		Name:       opCreateConfigurationSet,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &CreateConfigurationSetInput{}
	}

	req := c.newRequest(op, input, &CreateConfigurationSetOutput{})

	return CreateConfigurationSetRequest{Request: req, Input: input, Copy: c.CreateConfigurationSetRequest}
}

// CreateConfigurationSetRequest is the request type for the
// CreateConfigurationSet API operation.
type CreateConfigurationSetRequest struct {
	*aws.Request
	Input *CreateConfigurationSetInput
	Copy  func(*CreateConfigurationSetInput) CreateConfigurationSetRequest
}

// Send marshals and sends the CreateConfigurationSet API request.
func (r CreateConfigurationSetRequest) Send(ctx context.Context) (*CreateConfigurationSetResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateConfigurationSetResponse{
		CreateConfigurationSetOutput: r.Request.Data.(*CreateConfigurationSetOutput),
		response:                     &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateConfigurationSetResponse is the response type for the
// CreateConfigurationSet API operation.
type CreateConfigurationSetResponse struct {
	*CreateConfigurationSetOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateConfigurationSet request.
func (r *CreateConfigurationSetResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
