// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package applicationdiscoveryservice

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
)

type AssociateConfigurationItemsToApplicationInput struct {
	_ struct{} `type:"structure"`

	// The configuration ID of an application with which items are to be associated.
	//
	// ApplicationConfigurationId is a required field
	ApplicationConfigurationId *string `locationName:"applicationConfigurationId" type:"string" required:"true"`

	// The ID of each configuration item to be associated with an application.
	//
	// ConfigurationIds is a required field
	ConfigurationIds []string `locationName:"configurationIds" type:"list" required:"true"`
}

// String returns the string representation
func (s AssociateConfigurationItemsToApplicationInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *AssociateConfigurationItemsToApplicationInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "AssociateConfigurationItemsToApplicationInput"}

	if s.ApplicationConfigurationId == nil {
		invalidParams.Add(aws.NewErrParamRequired("ApplicationConfigurationId"))
	}

	if s.ConfigurationIds == nil {
		invalidParams.Add(aws.NewErrParamRequired("ConfigurationIds"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type AssociateConfigurationItemsToApplicationOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s AssociateConfigurationItemsToApplicationOutput) String() string {
	return awsutil.Prettify(s)
}

const opAssociateConfigurationItemsToApplication = "AssociateConfigurationItemsToApplication"

// AssociateConfigurationItemsToApplicationRequest returns a request value for making API operation for
// AWS Application Discovery Service.
//
// Associates one or more configuration items with an application.
//
//    // Example sending a request using AssociateConfigurationItemsToApplicationRequest.
//    req := client.AssociateConfigurationItemsToApplicationRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/discovery-2015-11-01/AssociateConfigurationItemsToApplication
func (c *Client) AssociateConfigurationItemsToApplicationRequest(input *AssociateConfigurationItemsToApplicationInput) AssociateConfigurationItemsToApplicationRequest {
	op := &aws.Operation{
		Name:       opAssociateConfigurationItemsToApplication,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &AssociateConfigurationItemsToApplicationInput{}
	}

	req := c.newRequest(op, input, &AssociateConfigurationItemsToApplicationOutput{})

	return AssociateConfigurationItemsToApplicationRequest{Request: req, Input: input, Copy: c.AssociateConfigurationItemsToApplicationRequest}
}

// AssociateConfigurationItemsToApplicationRequest is the request type for the
// AssociateConfigurationItemsToApplication API operation.
type AssociateConfigurationItemsToApplicationRequest struct {
	*aws.Request
	Input *AssociateConfigurationItemsToApplicationInput
	Copy  func(*AssociateConfigurationItemsToApplicationInput) AssociateConfigurationItemsToApplicationRequest
}

// Send marshals and sends the AssociateConfigurationItemsToApplication API request.
func (r AssociateConfigurationItemsToApplicationRequest) Send(ctx context.Context) (*AssociateConfigurationItemsToApplicationResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &AssociateConfigurationItemsToApplicationResponse{
		AssociateConfigurationItemsToApplicationOutput: r.Request.Data.(*AssociateConfigurationItemsToApplicationOutput),
		response: &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// AssociateConfigurationItemsToApplicationResponse is the response type for the
// AssociateConfigurationItemsToApplication API operation.
type AssociateConfigurationItemsToApplicationResponse struct {
	*AssociateConfigurationItemsToApplicationOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// AssociateConfigurationItemsToApplication request.
func (r *AssociateConfigurationItemsToApplicationResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
