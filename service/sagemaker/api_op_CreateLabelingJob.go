// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package sagemaker

import (
	"context"
	"fmt"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
)

type CreateLabelingJobInput struct {
	_ struct{} `type:"structure"`

	// Configures the labeling task and how it is presented to workers; including,
	// but not limited to price, keywords, and batch size (task count).
	//
	// HumanTaskConfig is a required field
	HumanTaskConfig *HumanTaskConfig `type:"structure" required:"true"`

	// Input data for the labeling job, such as the Amazon S3 location of the data
	// objects and the location of the manifest file that describes the data objects.
	//
	// InputConfig is a required field
	InputConfig *LabelingJobInputConfig `type:"structure" required:"true"`

	// The attribute name to use for the label in the output manifest file. This
	// is the key for the key/value pair formed with the label that a worker assigns
	// to the object. The name can't end with "-metadata". If you are running a
	// semantic segmentation labeling job, the attribute name must end with "-ref".
	// If you are running any other kind of labeling job, the attribute name must
	// not end with "-ref".
	//
	// LabelAttributeName is a required field
	LabelAttributeName *string `min:"1" type:"string" required:"true"`

	// The S3 URL of the file that defines the categories used to label the data
	// objects.
	//
	// The file is a JSON structure in the following format:
	//
	// {
	//
	// "document-version": "2018-11-28"
	//
	// "labels": [
	//
	// {
	//
	// "label": "label 1"
	//
	// },
	//
	// {
	//
	// "label": "label 2"
	//
	// },
	//
	// ...
	//
	// {
	//
	// "label": "label n"
	//
	// }
	//
	// ]
	//
	// }
	LabelCategoryConfigS3Uri *string `type:"string"`

	// Configures the information required to perform automated data labeling.
	LabelingJobAlgorithmsConfig *LabelingJobAlgorithmsConfig `type:"structure"`

	// The name of the labeling job. This name is used to identify the job in a
	// list of labeling jobs.
	//
	// LabelingJobName is a required field
	LabelingJobName *string `min:"1" type:"string" required:"true"`

	// The location of the output data and the AWS Key Management Service key ID
	// for the key used to encrypt the output data, if any.
	//
	// OutputConfig is a required field
	OutputConfig *LabelingJobOutputConfig `type:"structure" required:"true"`

	// The Amazon Resource Number (ARN) that Amazon SageMaker assumes to perform
	// tasks on your behalf during data labeling. You must grant this role the necessary
	// permissions so that Amazon SageMaker can successfully complete data labeling.
	//
	// RoleArn is a required field
	RoleArn *string `min:"20" type:"string" required:"true"`

	// A set of conditions for stopping the labeling job. If any of the conditions
	// are met, the job is automatically stopped. You can use these conditions to
	// control the cost of data labeling.
	StoppingConditions *LabelingJobStoppingConditions `type:"structure"`

	// An array of key/value pairs. For more information, see Using Cost Allocation
	// Tags (https://docs.aws.amazon.com/awsaccountbilling/latest/aboutv2/cost-alloc-tags.html#allocation-what)
	// in the AWS Billing and Cost Management User Guide.
	Tags []Tag `type:"list"`
}

// String returns the string representation
func (s CreateLabelingJobInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateLabelingJobInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateLabelingJobInput"}

	if s.HumanTaskConfig == nil {
		invalidParams.Add(aws.NewErrParamRequired("HumanTaskConfig"))
	}

	if s.InputConfig == nil {
		invalidParams.Add(aws.NewErrParamRequired("InputConfig"))
	}

	if s.LabelAttributeName == nil {
		invalidParams.Add(aws.NewErrParamRequired("LabelAttributeName"))
	}
	if s.LabelAttributeName != nil && len(*s.LabelAttributeName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("LabelAttributeName", 1))
	}

	if s.LabelingJobName == nil {
		invalidParams.Add(aws.NewErrParamRequired("LabelingJobName"))
	}
	if s.LabelingJobName != nil && len(*s.LabelingJobName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("LabelingJobName", 1))
	}

	if s.OutputConfig == nil {
		invalidParams.Add(aws.NewErrParamRequired("OutputConfig"))
	}

	if s.RoleArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("RoleArn"))
	}
	if s.RoleArn != nil && len(*s.RoleArn) < 20 {
		invalidParams.Add(aws.NewErrParamMinLen("RoleArn", 20))
	}
	if s.HumanTaskConfig != nil {
		if err := s.HumanTaskConfig.Validate(); err != nil {
			invalidParams.AddNested("HumanTaskConfig", err.(aws.ErrInvalidParams))
		}
	}
	if s.InputConfig != nil {
		if err := s.InputConfig.Validate(); err != nil {
			invalidParams.AddNested("InputConfig", err.(aws.ErrInvalidParams))
		}
	}
	if s.LabelingJobAlgorithmsConfig != nil {
		if err := s.LabelingJobAlgorithmsConfig.Validate(); err != nil {
			invalidParams.AddNested("LabelingJobAlgorithmsConfig", err.(aws.ErrInvalidParams))
		}
	}
	if s.OutputConfig != nil {
		if err := s.OutputConfig.Validate(); err != nil {
			invalidParams.AddNested("OutputConfig", err.(aws.ErrInvalidParams))
		}
	}
	if s.StoppingConditions != nil {
		if err := s.StoppingConditions.Validate(); err != nil {
			invalidParams.AddNested("StoppingConditions", err.(aws.ErrInvalidParams))
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

type CreateLabelingJobOutput struct {
	_ struct{} `type:"structure"`

	// The Amazon Resource Name (ARN) of the labeling job. You use this ARN to identify
	// the labeling job.
	//
	// LabelingJobArn is a required field
	LabelingJobArn *string `type:"string" required:"true"`
}

// String returns the string representation
func (s CreateLabelingJobOutput) String() string {
	return awsutil.Prettify(s)
}

const opCreateLabelingJob = "CreateLabelingJob"

// CreateLabelingJobRequest returns a request value for making API operation for
// Amazon SageMaker Service.
//
// Creates a job that uses workers to label the data objects in your input dataset.
// You can use the labeled data to train machine learning models.
//
// You can select your workforce from one of three providers:
//
//    * A private workforce that you create. It can include employees, contractors,
//    and outside experts. Use a private workforce when want the data to stay
//    within your organization or when a specific set of skills is required.
//
//    * One or more vendors that you select from the AWS Marketplace. Vendors
//    provide expertise in specific areas.
//
//    * The Amazon Mechanical Turk workforce. This is the largest workforce,
//    but it should only be used for public data or data that has been stripped
//    of any personally identifiable information.
//
// You can also use automated data labeling to reduce the number of data objects
// that need to be labeled by a human. Automated data labeling uses active learning
// to determine if a data object can be labeled by machine or if it needs to
// be sent to a human worker. For more information, see Using Automated Data
// Labeling (https://docs.aws.amazon.com/sagemaker/latest/dg/sms-automated-labeling.html).
//
// The data objects to be labeled are contained in an Amazon S3 bucket. You
// create a manifest file that describes the location of each object. For more
// information, see Using Input and Output Data (https://docs.aws.amazon.com/sagemaker/latest/dg/sms-data.html).
//
// The output can be used as the manifest file for another labeling job or as
// training data for your machine learning models.
//
//    // Example sending a request using CreateLabelingJobRequest.
//    req := client.CreateLabelingJobRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/sagemaker-2017-07-24/CreateLabelingJob
func (c *Client) CreateLabelingJobRequest(input *CreateLabelingJobInput) CreateLabelingJobRequest {
	op := &aws.Operation{
		Name:       opCreateLabelingJob,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &CreateLabelingJobInput{}
	}

	req := c.newRequest(op, input, &CreateLabelingJobOutput{})

	return CreateLabelingJobRequest{Request: req, Input: input, Copy: c.CreateLabelingJobRequest}
}

// CreateLabelingJobRequest is the request type for the
// CreateLabelingJob API operation.
type CreateLabelingJobRequest struct {
	*aws.Request
	Input *CreateLabelingJobInput
	Copy  func(*CreateLabelingJobInput) CreateLabelingJobRequest
}

// Send marshals and sends the CreateLabelingJob API request.
func (r CreateLabelingJobRequest) Send(ctx context.Context) (*CreateLabelingJobResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateLabelingJobResponse{
		CreateLabelingJobOutput: r.Request.Data.(*CreateLabelingJobOutput),
		response:                &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateLabelingJobResponse is the response type for the
// CreateLabelingJob API operation.
type CreateLabelingJobResponse struct {
	*CreateLabelingJobOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateLabelingJob request.
func (r *CreateLabelingJobResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
