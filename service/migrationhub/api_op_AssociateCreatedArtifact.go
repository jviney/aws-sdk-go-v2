// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package migrationhub

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
)

type AssociateCreatedArtifactInput struct {
	_ struct{} `type:"structure"`

	// An ARN of the AWS resource related to the migration (e.g., AMI, EC2 instance,
	// RDS instance, etc.)
	//
	// CreatedArtifact is a required field
	CreatedArtifact *CreatedArtifact `type:"structure" required:"true"`

	// Optional boolean flag to indicate whether any effect should take place. Used
	// to test if the caller has permission to make the call.
	DryRun *bool `type:"boolean"`

	// Unique identifier that references the migration task. Do not store personal
	// data in this field.
	//
	// MigrationTaskName is a required field
	MigrationTaskName *string `min:"1" type:"string" required:"true"`

	// The name of the ProgressUpdateStream.
	//
	// ProgressUpdateStream is a required field
	ProgressUpdateStream *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s AssociateCreatedArtifactInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *AssociateCreatedArtifactInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "AssociateCreatedArtifactInput"}

	if s.CreatedArtifact == nil {
		invalidParams.Add(aws.NewErrParamRequired("CreatedArtifact"))
	}

	if s.MigrationTaskName == nil {
		invalidParams.Add(aws.NewErrParamRequired("MigrationTaskName"))
	}
	if s.MigrationTaskName != nil && len(*s.MigrationTaskName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("MigrationTaskName", 1))
	}

	if s.ProgressUpdateStream == nil {
		invalidParams.Add(aws.NewErrParamRequired("ProgressUpdateStream"))
	}
	if s.ProgressUpdateStream != nil && len(*s.ProgressUpdateStream) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ProgressUpdateStream", 1))
	}
	if s.CreatedArtifact != nil {
		if err := s.CreatedArtifact.Validate(); err != nil {
			invalidParams.AddNested("CreatedArtifact", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type AssociateCreatedArtifactOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s AssociateCreatedArtifactOutput) String() string {
	return awsutil.Prettify(s)
}

const opAssociateCreatedArtifact = "AssociateCreatedArtifact"

// AssociateCreatedArtifactRequest returns a request value for making API operation for
// AWS Migration Hub.
//
// Associates a created artifact of an AWS cloud resource, the target receiving
// the migration, with the migration task performed by a migration tool. This
// API has the following traits:
//
//    * Migration tools can call the AssociateCreatedArtifact operation to indicate
//    which AWS artifact is associated with a migration task.
//
//    * The created artifact name must be provided in ARN (Amazon Resource Name)
//    format which will contain information about type and region; for example:
//    arn:aws:ec2:us-east-1:488216288981:image/ami-6d0ba87b.
//
//    * Examples of the AWS resource behind the created artifact are, AMI's,
//    EC2 instance, or DMS endpoint, etc.
//
//    // Example sending a request using AssociateCreatedArtifactRequest.
//    req := client.AssociateCreatedArtifactRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/AWSMigrationHub-2017-05-31/AssociateCreatedArtifact
func (c *Client) AssociateCreatedArtifactRequest(input *AssociateCreatedArtifactInput) AssociateCreatedArtifactRequest {
	op := &aws.Operation{
		Name:       opAssociateCreatedArtifact,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &AssociateCreatedArtifactInput{}
	}

	req := c.newRequest(op, input, &AssociateCreatedArtifactOutput{})

	return AssociateCreatedArtifactRequest{Request: req, Input: input, Copy: c.AssociateCreatedArtifactRequest}
}

// AssociateCreatedArtifactRequest is the request type for the
// AssociateCreatedArtifact API operation.
type AssociateCreatedArtifactRequest struct {
	*aws.Request
	Input *AssociateCreatedArtifactInput
	Copy  func(*AssociateCreatedArtifactInput) AssociateCreatedArtifactRequest
}

// Send marshals and sends the AssociateCreatedArtifact API request.
func (r AssociateCreatedArtifactRequest) Send(ctx context.Context) (*AssociateCreatedArtifactResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &AssociateCreatedArtifactResponse{
		AssociateCreatedArtifactOutput: r.Request.Data.(*AssociateCreatedArtifactOutput),
		response:                       &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// AssociateCreatedArtifactResponse is the response type for the
// AssociateCreatedArtifact API operation.
type AssociateCreatedArtifactResponse struct {
	*AssociateCreatedArtifactOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// AssociateCreatedArtifact request.
func (r *AssociateCreatedArtifactResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
