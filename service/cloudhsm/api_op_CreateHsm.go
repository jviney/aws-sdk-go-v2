// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package cloudhsm

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
)

// Contains the inputs for the CreateHsm operation.
type CreateHsmInput struct {
	_ struct{} `locationName:"CreateHsmRequest" type:"structure"`

	// A user-defined token to ensure idempotence. Subsequent calls to this operation
	// with the same token will be ignored.
	ClientToken *string `locationName:"ClientToken" type:"string"`

	// The IP address to assign to the HSM's ENI.
	//
	// If an IP address is not specified, an IP address will be randomly chosen
	// from the CIDR range of the subnet.
	EniIp *string `locationName:"EniIp" type:"string"`

	// The external ID from IamRoleArn, if present.
	ExternalId *string `locationName:"ExternalId" type:"string"`

	// The ARN of an IAM role to enable the AWS CloudHSM service to allocate an
	// ENI on your behalf.
	//
	// IamRoleArn is a required field
	IamRoleArn *string `locationName:"IamRoleArn" type:"string" required:"true"`

	// The SSH public key to install on the HSM.
	//
	// SshKey is a required field
	SshKey *string `locationName:"SshKey" type:"string" required:"true"`

	// The identifier of the subnet in your VPC in which to place the HSM.
	//
	// SubnetId is a required field
	SubnetId *string `locationName:"SubnetId" type:"string" required:"true"`

	// Specifies the type of subscription for the HSM.
	//
	//    * PRODUCTION - The HSM is being used in a production environment.
	//
	//    * TRIAL - The HSM is being used in a product trial.
	//
	// SubscriptionType is a required field
	SubscriptionType SubscriptionType `locationName:"SubscriptionType" type:"string" required:"true" enum:"true"`

	// The IP address for the syslog monitoring server. The AWS CloudHSM service
	// only supports one syslog monitoring server.
	SyslogIp *string `locationName:"SyslogIp" type:"string"`
}

// String returns the string representation
func (s CreateHsmInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateHsmInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateHsmInput"}

	if s.IamRoleArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("IamRoleArn"))
	}

	if s.SshKey == nil {
		invalidParams.Add(aws.NewErrParamRequired("SshKey"))
	}

	if s.SubnetId == nil {
		invalidParams.Add(aws.NewErrParamRequired("SubnetId"))
	}
	if len(s.SubscriptionType) == 0 {
		invalidParams.Add(aws.NewErrParamRequired("SubscriptionType"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Contains the output of the CreateHsm operation.
type CreateHsmOutput struct {
	_ struct{} `type:"structure"`

	// The ARN of the HSM.
	HsmArn *string `type:"string"`
}

// String returns the string representation
func (s CreateHsmOutput) String() string {
	return awsutil.Prettify(s)
}

const opCreateHsm = "CreateHsm"

// CreateHsmRequest returns a request value for making API operation for
// Amazon CloudHSM.
//
// This is documentation for AWS CloudHSM Classic. For more information, see
// AWS CloudHSM Classic FAQs (http://aws.amazon.com/cloudhsm/faqs-classic/),
// the AWS CloudHSM Classic User Guide (http://docs.aws.amazon.com/cloudhsm/classic/userguide/),
// and the AWS CloudHSM Classic API Reference (http://docs.aws.amazon.com/cloudhsm/classic/APIReference/).
//
// For information about the current version of AWS CloudHSM, see AWS CloudHSM
// (http://aws.amazon.com/cloudhsm/), the AWS CloudHSM User Guide (http://docs.aws.amazon.com/cloudhsm/latest/userguide/),
// and the AWS CloudHSM API Reference (http://docs.aws.amazon.com/cloudhsm/latest/APIReference/).
//
// Creates an uninitialized HSM instance.
//
// There is an upfront fee charged for each HSM instance that you create with
// the CreateHsm operation. If you accidentally provision an HSM and want to
// request a refund, delete the instance using the DeleteHsm operation, go to
// the AWS Support Center (https://console.aws.amazon.com/support/home), create
// a new case, and select Account and Billing Support.
//
// It can take up to 20 minutes to create and provision an HSM. You can monitor
// the status of the HSM with the DescribeHsm operation. The HSM is ready to
// be initialized when the status changes to RUNNING.
//
//    // Example sending a request using CreateHsmRequest.
//    req := client.CreateHsmRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/cloudhsm-2014-05-30/CreateHsm
func (c *Client) CreateHsmRequest(input *CreateHsmInput) CreateHsmRequest {
	op := &aws.Operation{
		Name:       opCreateHsm,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &CreateHsmInput{}
	}

	req := c.newRequest(op, input, &CreateHsmOutput{})

	return CreateHsmRequest{Request: req, Input: input, Copy: c.CreateHsmRequest}
}

// CreateHsmRequest is the request type for the
// CreateHsm API operation.
type CreateHsmRequest struct {
	*aws.Request
	Input *CreateHsmInput
	Copy  func(*CreateHsmInput) CreateHsmRequest
}

// Send marshals and sends the CreateHsm API request.
func (r CreateHsmRequest) Send(ctx context.Context) (*CreateHsmResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateHsmResponse{
		CreateHsmOutput: r.Request.Data.(*CreateHsmOutput),
		response:        &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateHsmResponse is the response type for the
// CreateHsm API operation.
type CreateHsmResponse struct {
	*CreateHsmOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateHsm request.
func (r *CreateHsmResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
