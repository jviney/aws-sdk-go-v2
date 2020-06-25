// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package inspector

import (
	"context"
	"fmt"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
)

type CreateAssessmentTemplateInput struct {
	_ struct{} `type:"structure"`

	// The ARN that specifies the assessment target for which you want to create
	// the assessment template.
	//
	// AssessmentTargetArn is a required field
	AssessmentTargetArn *string `locationName:"assessmentTargetArn" min:"1" type:"string" required:"true"`

	// The user-defined name that identifies the assessment template that you want
	// to create. You can create several assessment templates for an assessment
	// target. The names of the assessment templates that correspond to a particular
	// assessment target must be unique.
	//
	// AssessmentTemplateName is a required field
	AssessmentTemplateName *string `locationName:"assessmentTemplateName" min:"1" type:"string" required:"true"`

	// The duration of the assessment run in seconds.
	//
	// DurationInSeconds is a required field
	DurationInSeconds *int64 `locationName:"durationInSeconds" min:"180" type:"integer" required:"true"`

	// The ARNs that specify the rules packages that you want to attach to the assessment
	// template.
	//
	// RulesPackageArns is a required field
	RulesPackageArns []string `locationName:"rulesPackageArns" type:"list" required:"true"`

	// The user-defined attributes that are assigned to every finding that is generated
	// by the assessment run that uses this assessment template. An attribute is
	// a key and value pair (an Attribute object). Within an assessment template,
	// each key must be unique.
	UserAttributesForFindings []Attribute `locationName:"userAttributesForFindings" type:"list"`
}

// String returns the string representation
func (s CreateAssessmentTemplateInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateAssessmentTemplateInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateAssessmentTemplateInput"}

	if s.AssessmentTargetArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("AssessmentTargetArn"))
	}
	if s.AssessmentTargetArn != nil && len(*s.AssessmentTargetArn) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("AssessmentTargetArn", 1))
	}

	if s.AssessmentTemplateName == nil {
		invalidParams.Add(aws.NewErrParamRequired("AssessmentTemplateName"))
	}
	if s.AssessmentTemplateName != nil && len(*s.AssessmentTemplateName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("AssessmentTemplateName", 1))
	}

	if s.DurationInSeconds == nil {
		invalidParams.Add(aws.NewErrParamRequired("DurationInSeconds"))
	}
	if s.DurationInSeconds != nil && *s.DurationInSeconds < 180 {
		invalidParams.Add(aws.NewErrParamMinValue("DurationInSeconds", 180))
	}

	if s.RulesPackageArns == nil {
		invalidParams.Add(aws.NewErrParamRequired("RulesPackageArns"))
	}
	if s.UserAttributesForFindings != nil {
		for i, v := range s.UserAttributesForFindings {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "UserAttributesForFindings", i), err.(aws.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type CreateAssessmentTemplateOutput struct {
	_ struct{} `type:"structure"`

	// The ARN that specifies the assessment template that is created.
	//
	// AssessmentTemplateArn is a required field
	AssessmentTemplateArn *string `locationName:"assessmentTemplateArn" min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s CreateAssessmentTemplateOutput) String() string {
	return awsutil.Prettify(s)
}

const opCreateAssessmentTemplate = "CreateAssessmentTemplate"

// CreateAssessmentTemplateRequest returns a request value for making API operation for
// Amazon Inspector.
//
// Creates an assessment template for the assessment target that is specified
// by the ARN of the assessment target. If the service-linked role (https://docs.aws.amazon.com/inspector/latest/userguide/inspector_slr.html)
// isn’t already registered, this action also creates and registers a service-linked
// role to grant Amazon Inspector access to AWS Services needed to perform security
// assessments.
//
//    // Example sending a request using CreateAssessmentTemplateRequest.
//    req := client.CreateAssessmentTemplateRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/inspector-2016-02-16/CreateAssessmentTemplate
func (c *Client) CreateAssessmentTemplateRequest(input *CreateAssessmentTemplateInput) CreateAssessmentTemplateRequest {
	op := &aws.Operation{
		Name:       opCreateAssessmentTemplate,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &CreateAssessmentTemplateInput{}
	}

	req := c.newRequest(op, input, &CreateAssessmentTemplateOutput{})

	return CreateAssessmentTemplateRequest{Request: req, Input: input, Copy: c.CreateAssessmentTemplateRequest}
}

// CreateAssessmentTemplateRequest is the request type for the
// CreateAssessmentTemplate API operation.
type CreateAssessmentTemplateRequest struct {
	*aws.Request
	Input *CreateAssessmentTemplateInput
	Copy  func(*CreateAssessmentTemplateInput) CreateAssessmentTemplateRequest
}

// Send marshals and sends the CreateAssessmentTemplate API request.
func (r CreateAssessmentTemplateRequest) Send(ctx context.Context) (*CreateAssessmentTemplateResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateAssessmentTemplateResponse{
		CreateAssessmentTemplateOutput: r.Request.Data.(*CreateAssessmentTemplateOutput),
		response:                       &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateAssessmentTemplateResponse is the response type for the
// CreateAssessmentTemplate API operation.
type CreateAssessmentTemplateResponse struct {
	*CreateAssessmentTemplateOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateAssessmentTemplate request.
func (r *CreateAssessmentTemplateResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
