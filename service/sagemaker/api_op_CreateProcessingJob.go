// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package sagemaker

import (
	"context"
	"fmt"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
)

type CreateProcessingJobInput struct {
	_ struct{} `type:"structure"`

	// Configures the processing job to run a specified Docker container image.
	//
	// AppSpecification is a required field
	AppSpecification *AppSpecification `type:"structure" required:"true"`

	// Sets the environment variables in the Docker container.
	Environment map[string]string `type:"map"`

	// Configuration for the experiment.
	ExperimentConfig *ExperimentConfig `type:"structure"`

	// Networking options for a processing job.
	NetworkConfig *NetworkConfig `type:"structure"`

	// For each input, data is downloaded from S3 into the processing container
	// before the processing job begins running if "S3InputMode" is set to File.
	ProcessingInputs []ProcessingInput `type:"list"`

	// The name of the processing job. The name must be unique within an AWS Region
	// in the AWS account.
	//
	// ProcessingJobName is a required field
	ProcessingJobName *string `min:"1" type:"string" required:"true"`

	// Output configuration for the processing job.
	ProcessingOutputConfig *ProcessingOutputConfig `type:"structure"`

	// Identifies the resources, ML compute instances, and ML storage volumes to
	// deploy for a processing job. In distributed training, you specify more than
	// one instance.
	//
	// ProcessingResources is a required field
	ProcessingResources *ProcessingResources `type:"structure" required:"true"`

	// The Amazon Resource Name (ARN) of an IAM role that Amazon SageMaker can assume
	// to perform tasks on your behalf.
	//
	// RoleArn is a required field
	RoleArn *string `min:"20" type:"string" required:"true"`

	// The time limit for how long the processing job is allowed to run.
	StoppingCondition *ProcessingStoppingCondition `type:"structure"`

	// (Optional) An array of key-value pairs. For more information, see Using Cost
	// Allocation Tags (https://docs.aws.amazon.com/awsaccountbilling/latest/aboutv2/cost-alloc-tags.html#allocation-whatURL)
	// in the AWS Billing and Cost Management User Guide.
	Tags []Tag `type:"list"`
}

// String returns the string representation
func (s CreateProcessingJobInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateProcessingJobInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateProcessingJobInput"}

	if s.AppSpecification == nil {
		invalidParams.Add(aws.NewErrParamRequired("AppSpecification"))
	}

	if s.ProcessingJobName == nil {
		invalidParams.Add(aws.NewErrParamRequired("ProcessingJobName"))
	}
	if s.ProcessingJobName != nil && len(*s.ProcessingJobName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ProcessingJobName", 1))
	}

	if s.ProcessingResources == nil {
		invalidParams.Add(aws.NewErrParamRequired("ProcessingResources"))
	}

	if s.RoleArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("RoleArn"))
	}
	if s.RoleArn != nil && len(*s.RoleArn) < 20 {
		invalidParams.Add(aws.NewErrParamMinLen("RoleArn", 20))
	}
	if s.AppSpecification != nil {
		if err := s.AppSpecification.Validate(); err != nil {
			invalidParams.AddNested("AppSpecification", err.(aws.ErrInvalidParams))
		}
	}
	if s.ExperimentConfig != nil {
		if err := s.ExperimentConfig.Validate(); err != nil {
			invalidParams.AddNested("ExperimentConfig", err.(aws.ErrInvalidParams))
		}
	}
	if s.NetworkConfig != nil {
		if err := s.NetworkConfig.Validate(); err != nil {
			invalidParams.AddNested("NetworkConfig", err.(aws.ErrInvalidParams))
		}
	}
	if s.ProcessingInputs != nil {
		for i, v := range s.ProcessingInputs {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "ProcessingInputs", i), err.(aws.ErrInvalidParams))
			}
		}
	}
	if s.ProcessingOutputConfig != nil {
		if err := s.ProcessingOutputConfig.Validate(); err != nil {
			invalidParams.AddNested("ProcessingOutputConfig", err.(aws.ErrInvalidParams))
		}
	}
	if s.ProcessingResources != nil {
		if err := s.ProcessingResources.Validate(); err != nil {
			invalidParams.AddNested("ProcessingResources", err.(aws.ErrInvalidParams))
		}
	}
	if s.StoppingCondition != nil {
		if err := s.StoppingCondition.Validate(); err != nil {
			invalidParams.AddNested("StoppingCondition", err.(aws.ErrInvalidParams))
		}
	}
	if s.Tags != nil {
		for i, v := range s.Tags {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "Tags", i), err.(aws.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type CreateProcessingJobOutput struct {
	_ struct{} `type:"structure"`

	// The Amazon Resource Name (ARN) of the processing job.
	//
	// ProcessingJobArn is a required field
	ProcessingJobArn *string `type:"string" required:"true"`
}

// String returns the string representation
func (s CreateProcessingJobOutput) String() string {
	return awsutil.Prettify(s)
}

const opCreateProcessingJob = "CreateProcessingJob"

// CreateProcessingJobRequest returns a request value for making API operation for
// Amazon SageMaker Service.
//
// Creates a processing job.
//
//    // Example sending a request using CreateProcessingJobRequest.
//    req := client.CreateProcessingJobRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/sagemaker-2017-07-24/CreateProcessingJob
func (c *Client) CreateProcessingJobRequest(input *CreateProcessingJobInput) CreateProcessingJobRequest {
	op := &aws.Operation{
		Name:       opCreateProcessingJob,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &CreateProcessingJobInput{}
	}

	req := c.newRequest(op, input, &CreateProcessingJobOutput{})

	return CreateProcessingJobRequest{Request: req, Input: input, Copy: c.CreateProcessingJobRequest}
}

// CreateProcessingJobRequest is the request type for the
// CreateProcessingJob API operation.
type CreateProcessingJobRequest struct {
	*aws.Request
	Input *CreateProcessingJobInput
	Copy  func(*CreateProcessingJobInput) CreateProcessingJobRequest
}

// Send marshals and sends the CreateProcessingJob API request.
func (r CreateProcessingJobRequest) Send(ctx context.Context) (*CreateProcessingJobResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateProcessingJobResponse{
		CreateProcessingJobOutput: r.Request.Data.(*CreateProcessingJobOutput),
		response:                  &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateProcessingJobResponse is the response type for the
// CreateProcessingJob API operation.
type CreateProcessingJobResponse struct {
	*CreateProcessingJobOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateProcessingJob request.
func (r *CreateProcessingJobResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
