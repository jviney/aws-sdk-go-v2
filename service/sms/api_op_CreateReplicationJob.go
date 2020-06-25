// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package sms

import (
	"context"
	"time"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
)

type CreateReplicationJobInput struct {
	_ struct{} `type:"structure"`

	// The description of the replication job.
	Description *string `locationName:"description" type:"string"`

	// When true, the replication job produces encrypted AMIs. See also KmsKeyId
	// below.
	Encrypted *bool `locationName:"encrypted" type:"boolean"`

	// The time between consecutive replication runs, in hours.
	Frequency *int64 `locationName:"frequency" type:"integer"`

	// KMS key ID for replication jobs that produce encrypted AMIs. Can be any of
	// the following:
	//
	//    * KMS key ID
	//
	//    * KMS key alias
	//
	//    * ARN referring to KMS key ID
	//
	//    * ARN referring to KMS key alias
	//
	// If encrypted is true but a KMS key id is not specified, the customer's default
	// KMS key for EBS is used.
	KmsKeyId *string `locationName:"kmsKeyId" type:"string"`

	// The license type to be used for the AMI created by a successful replication
	// run.
	LicenseType LicenseType `locationName:"licenseType" type:"string" enum:"true"`

	// The maximum number of SMS-created AMIs to retain. The oldest will be deleted
	// once the maximum number is reached and a new AMI is created.
	NumberOfRecentAmisToKeep *int64 `locationName:"numberOfRecentAmisToKeep" type:"integer"`

	// The name of the IAM role to be used by the AWS SMS.
	RoleName *string `locationName:"roleName" type:"string"`

	RunOnce *bool `locationName:"runOnce" type:"boolean"`

	// The seed replication time.
	//
	// SeedReplicationTime is a required field
	SeedReplicationTime *time.Time `locationName:"seedReplicationTime" type:"timestamp" required:"true"`

	// The identifier of the server.
	//
	// ServerId is a required field
	ServerId *string `locationName:"serverId" type:"string" required:"true"`
}

// String returns the string representation
func (s CreateReplicationJobInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateReplicationJobInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateReplicationJobInput"}

	if s.SeedReplicationTime == nil {
		invalidParams.Add(aws.NewErrParamRequired("SeedReplicationTime"))
	}

	if s.ServerId == nil {
		invalidParams.Add(aws.NewErrParamRequired("ServerId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type CreateReplicationJobOutput struct {
	_ struct{} `type:"structure"`

	// The unique identifier of the replication job.
	ReplicationJobId *string `locationName:"replicationJobId" type:"string"`
}

// String returns the string representation
func (s CreateReplicationJobOutput) String() string {
	return awsutil.Prettify(s)
}

const opCreateReplicationJob = "CreateReplicationJob"

// CreateReplicationJobRequest returns a request value for making API operation for
// AWS Server Migration Service.
//
// Creates a replication job. The replication job schedules periodic replication
// runs to replicate your server to AWS. Each replication run creates an Amazon
// Machine Image (AMI).
//
//    // Example sending a request using CreateReplicationJobRequest.
//    req := client.CreateReplicationJobRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/sms-2016-10-24/CreateReplicationJob
func (c *Client) CreateReplicationJobRequest(input *CreateReplicationJobInput) CreateReplicationJobRequest {
	op := &aws.Operation{
		Name:       opCreateReplicationJob,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &CreateReplicationJobInput{}
	}

	req := c.newRequest(op, input, &CreateReplicationJobOutput{})

	return CreateReplicationJobRequest{Request: req, Input: input, Copy: c.CreateReplicationJobRequest}
}

// CreateReplicationJobRequest is the request type for the
// CreateReplicationJob API operation.
type CreateReplicationJobRequest struct {
	*aws.Request
	Input *CreateReplicationJobInput
	Copy  func(*CreateReplicationJobInput) CreateReplicationJobRequest
}

// Send marshals and sends the CreateReplicationJob API request.
func (r CreateReplicationJobRequest) Send(ctx context.Context) (*CreateReplicationJobResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateReplicationJobResponse{
		CreateReplicationJobOutput: r.Request.Data.(*CreateReplicationJobOutput),
		response:                   &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateReplicationJobResponse is the response type for the
// CreateReplicationJob API operation.
type CreateReplicationJobResponse struct {
	*CreateReplicationJobOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateReplicationJob request.
func (r *CreateReplicationJobResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
