// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package codestar

import (
	"context"
	"fmt"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
)

type CreateProjectInput struct {
	_ struct{} `type:"structure"`

	// A user- or system-generated token that identifies the entity that requested
	// project creation. This token can be used to repeat the request.
	ClientRequestToken *string `locationName:"clientRequestToken" min:"1" type:"string"`

	// The description of the project, if any.
	Description *string `locationName:"description" type:"string" sensitive:"true"`

	// The ID of the project to be created in AWS CodeStar.
	//
	// Id is a required field
	Id *string `locationName:"id" min:"2" type:"string" required:"true"`

	// The display name for the project to be created in AWS CodeStar.
	//
	// Name is a required field
	Name *string `locationName:"name" min:"1" type:"string" required:"true" sensitive:"true"`

	// A list of the Code objects submitted with the project request. If this parameter
	// is specified, the request must also include the toolchain parameter.
	SourceCode []Code `locationName:"sourceCode" type:"list"`

	// The tags created for the project.
	Tags map[string]string `locationName:"tags" type:"map"`

	// The name of the toolchain template file submitted with the project request.
	// If this parameter is specified, the request must also include the sourceCode
	// parameter.
	Toolchain *Toolchain `locationName:"toolchain" type:"structure"`
}

// String returns the string representation
func (s CreateProjectInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateProjectInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateProjectInput"}
	if s.ClientRequestToken != nil && len(*s.ClientRequestToken) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ClientRequestToken", 1))
	}

	if s.Id == nil {
		invalidParams.Add(aws.NewErrParamRequired("Id"))
	}
	if s.Id != nil && len(*s.Id) < 2 {
		invalidParams.Add(aws.NewErrParamMinLen("Id", 2))
	}

	if s.Name == nil {
		invalidParams.Add(aws.NewErrParamRequired("Name"))
	}
	if s.Name != nil && len(*s.Name) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Name", 1))
	}
	if s.SourceCode != nil {
		for i, v := range s.SourceCode {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "SourceCode", i), err.(aws.ErrInvalidParams))
			}
		}
	}
	if s.Toolchain != nil {
		if err := s.Toolchain.Validate(); err != nil {
			invalidParams.AddNested("Toolchain", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type CreateProjectOutput struct {
	_ struct{} `type:"structure"`

	// The Amazon Resource Name (ARN) of the created project.
	//
	// Arn is a required field
	Arn *string `locationName:"arn" type:"string" required:"true"`

	// A user- or system-generated token that identifies the entity that requested
	// project creation.
	ClientRequestToken *string `locationName:"clientRequestToken" min:"1" type:"string"`

	// The ID of the project.
	//
	// Id is a required field
	Id *string `locationName:"id" min:"2" type:"string" required:"true"`

	// Reserved for future use.
	ProjectTemplateId *string `locationName:"projectTemplateId" min:"1" type:"string"`
}

// String returns the string representation
func (s CreateProjectOutput) String() string {
	return awsutil.Prettify(s)
}

const opCreateProject = "CreateProject"

// CreateProjectRequest returns a request value for making API operation for
// AWS CodeStar.
//
// Creates a project, including project resources. This action creates a project
// based on a submitted project request. A set of source code files and a toolchain
// template file can be included with the project request. If these are not
// provided, an empty project is created.
//
//    // Example sending a request using CreateProjectRequest.
//    req := client.CreateProjectRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/codestar-2017-04-19/CreateProject
func (c *Client) CreateProjectRequest(input *CreateProjectInput) CreateProjectRequest {
	op := &aws.Operation{
		Name:       opCreateProject,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &CreateProjectInput{}
	}

	req := c.newRequest(op, input, &CreateProjectOutput{})

	return CreateProjectRequest{Request: req, Input: input, Copy: c.CreateProjectRequest}
}

// CreateProjectRequest is the request type for the
// CreateProject API operation.
type CreateProjectRequest struct {
	*aws.Request
	Input *CreateProjectInput
	Copy  func(*CreateProjectInput) CreateProjectRequest
}

// Send marshals and sends the CreateProject API request.
func (r CreateProjectRequest) Send(ctx context.Context) (*CreateProjectResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateProjectResponse{
		CreateProjectOutput: r.Request.Data.(*CreateProjectOutput),
		response:            &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateProjectResponse is the response type for the
// CreateProject API operation.
type CreateProjectResponse struct {
	*CreateProjectOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateProject request.
func (r *CreateProjectResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
