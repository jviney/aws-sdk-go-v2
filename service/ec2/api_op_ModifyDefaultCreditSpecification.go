// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ec2

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
)

type ModifyDefaultCreditSpecificationInput struct {
	_ struct{} `type:"structure"`

	// The credit option for CPU usage of the instance family.
	//
	// Valid Values: standard | unlimited
	//
	// CpuCredits is a required field
	CpuCredits *string `type:"string" required:"true"`

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have
	// the required permissions, the error response is DryRunOperation. Otherwise,
	// it is UnauthorizedOperation.
	DryRun *bool `type:"boolean"`

	// The instance family.
	//
	// InstanceFamily is a required field
	InstanceFamily UnlimitedSupportedInstanceFamily `type:"string" required:"true" enum:"true"`
}

// String returns the string representation
func (s ModifyDefaultCreditSpecificationInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ModifyDefaultCreditSpecificationInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ModifyDefaultCreditSpecificationInput"}

	if s.CpuCredits == nil {
		invalidParams.Add(aws.NewErrParamRequired("CpuCredits"))
	}
	if len(s.InstanceFamily) == 0 {
		invalidParams.Add(aws.NewErrParamRequired("InstanceFamily"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type ModifyDefaultCreditSpecificationOutput struct {
	_ struct{} `type:"structure"`

	// The default credit option for CPU usage of the instance family.
	InstanceFamilyCreditSpecification *InstanceFamilyCreditSpecification `locationName:"instanceFamilyCreditSpecification" type:"structure"`
}

// String returns the string representation
func (s ModifyDefaultCreditSpecificationOutput) String() string {
	return awsutil.Prettify(s)
}

const opModifyDefaultCreditSpecification = "ModifyDefaultCreditSpecification"

// ModifyDefaultCreditSpecificationRequest returns a request value for making API operation for
// Amazon Elastic Compute Cloud.
//
// Modifies the default credit option for CPU usage of burstable performance
// instances. The default credit option is set at the account level per AWS
// Region, and is specified per instance family. All new burstable performance
// instances in the account launch using the default credit option.
//
// ModifyDefaultCreditSpecification is an asynchronous operation, which works
// at an AWS Region level and modifies the credit option for each Availability
// Zone. All zones in a Region are updated within five minutes. But if instances
// are launched during this operation, they might not get the new credit option
// until the zone is updated. To verify whether the update has occurred, you
// can call GetDefaultCreditSpecification and check DefaultCreditSpecification
// for updates.
//
// For more information, see Burstable Performance Instances (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/burstable-performance-instances.html)
// in the Amazon Elastic Compute Cloud User Guide.
//
//    // Example sending a request using ModifyDefaultCreditSpecificationRequest.
//    req := client.ModifyDefaultCreditSpecificationRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/ModifyDefaultCreditSpecification
func (c *Client) ModifyDefaultCreditSpecificationRequest(input *ModifyDefaultCreditSpecificationInput) ModifyDefaultCreditSpecificationRequest {
	op := &aws.Operation{
		Name:       opModifyDefaultCreditSpecification,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ModifyDefaultCreditSpecificationInput{}
	}

	req := c.newRequest(op, input, &ModifyDefaultCreditSpecificationOutput{})

	return ModifyDefaultCreditSpecificationRequest{Request: req, Input: input, Copy: c.ModifyDefaultCreditSpecificationRequest}
}

// ModifyDefaultCreditSpecificationRequest is the request type for the
// ModifyDefaultCreditSpecification API operation.
type ModifyDefaultCreditSpecificationRequest struct {
	*aws.Request
	Input *ModifyDefaultCreditSpecificationInput
	Copy  func(*ModifyDefaultCreditSpecificationInput) ModifyDefaultCreditSpecificationRequest
}

// Send marshals and sends the ModifyDefaultCreditSpecification API request.
func (r ModifyDefaultCreditSpecificationRequest) Send(ctx context.Context) (*ModifyDefaultCreditSpecificationResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ModifyDefaultCreditSpecificationResponse{
		ModifyDefaultCreditSpecificationOutput: r.Request.Data.(*ModifyDefaultCreditSpecificationOutput),
		response:                               &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// ModifyDefaultCreditSpecificationResponse is the response type for the
// ModifyDefaultCreditSpecification API operation.
type ModifyDefaultCreditSpecificationResponse struct {
	*ModifyDefaultCreditSpecificationOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ModifyDefaultCreditSpecification request.
func (r *ModifyDefaultCreditSpecificationResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
