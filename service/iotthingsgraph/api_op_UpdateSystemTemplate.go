// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package iotthingsgraph

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
)

type UpdateSystemTemplateInput struct {
	_ struct{} `type:"structure"`

	// The version of the user's namespace. Defaults to the latest version of the
	// user's namespace.
	//
	// If no value is specified, the latest version is used by default.
	CompatibleNamespaceVersion *int64 `locationName:"compatibleNamespaceVersion" type:"long"`

	// The DefinitionDocument that contains the updated system definition.
	//
	// Definition is a required field
	Definition *DefinitionDocument `locationName:"definition" type:"structure" required:"true"`

	// The ID of the system to be updated.
	//
	// The ID should be in the following format.
	//
	// urn:tdm:REGION/ACCOUNT ID/default:system:SYSTEMNAME
	//
	// Id is a required field
	Id *string `locationName:"id" type:"string" required:"true"`
}

// String returns the string representation
func (s UpdateSystemTemplateInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UpdateSystemTemplateInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "UpdateSystemTemplateInput"}

	if s.Definition == nil {
		invalidParams.Add(aws.NewErrParamRequired("Definition"))
	}

	if s.Id == nil {
		invalidParams.Add(aws.NewErrParamRequired("Id"))
	}
	if s.Definition != nil {
		if err := s.Definition.Validate(); err != nil {
			invalidParams.AddNested("Definition", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type UpdateSystemTemplateOutput struct {
	_ struct{} `type:"structure"`

	// An object containing summary information about the updated system.
	Summary *SystemTemplateSummary `locationName:"summary" type:"structure"`
}

// String returns the string representation
func (s UpdateSystemTemplateOutput) String() string {
	return awsutil.Prettify(s)
}

const opUpdateSystemTemplate = "UpdateSystemTemplate"

// UpdateSystemTemplateRequest returns a request value for making API operation for
// AWS IoT Things Graph.
//
// Updates the specified system. You don't need to run this action after updating
// a workflow. Any deployment that uses the system will see the changes in the
// system when it is redeployed.
//
//    // Example sending a request using UpdateSystemTemplateRequest.
//    req := client.UpdateSystemTemplateRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/iotthingsgraph-2018-09-06/UpdateSystemTemplate
func (c *Client) UpdateSystemTemplateRequest(input *UpdateSystemTemplateInput) UpdateSystemTemplateRequest {
	op := &aws.Operation{
		Name:       opUpdateSystemTemplate,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &UpdateSystemTemplateInput{}
	}

	req := c.newRequest(op, input, &UpdateSystemTemplateOutput{})

	return UpdateSystemTemplateRequest{Request: req, Input: input, Copy: c.UpdateSystemTemplateRequest}
}

// UpdateSystemTemplateRequest is the request type for the
// UpdateSystemTemplate API operation.
type UpdateSystemTemplateRequest struct {
	*aws.Request
	Input *UpdateSystemTemplateInput
	Copy  func(*UpdateSystemTemplateInput) UpdateSystemTemplateRequest
}

// Send marshals and sends the UpdateSystemTemplate API request.
func (r UpdateSystemTemplateRequest) Send(ctx context.Context) (*UpdateSystemTemplateResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &UpdateSystemTemplateResponse{
		UpdateSystemTemplateOutput: r.Request.Data.(*UpdateSystemTemplateOutput),
		response:                   &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// UpdateSystemTemplateResponse is the response type for the
// UpdateSystemTemplate API operation.
type UpdateSystemTemplateResponse struct {
	*UpdateSystemTemplateOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// UpdateSystemTemplate request.
func (r *UpdateSystemTemplateResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
