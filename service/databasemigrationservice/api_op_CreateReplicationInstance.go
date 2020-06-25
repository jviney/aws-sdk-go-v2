// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package databasemigrationservice

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
)

type CreateReplicationInstanceInput struct {
	_ struct{} `type:"structure"`

	// The amount of storage (in gigabytes) to be initially allocated for the replication
	// instance.
	AllocatedStorage *int64 `type:"integer"`

	// A value that indicates whether minor engine upgrades are applied automatically
	// to the replication instance during the maintenance window. This parameter
	// defaults to true.
	//
	// Default: true
	AutoMinorVersionUpgrade *bool `type:"boolean"`

	// The Availability Zone where the replication instance will be created. The
	// default value is a random, system-chosen Availability Zone in the endpoint's
	// AWS Region, for example: us-east-1d
	AvailabilityZone *string `type:"string"`

	// A list of DNS name servers supported for the replication instance.
	DnsNameServers *string `type:"string"`

	// The engine version number of the replication instance.
	EngineVersion *string `type:"string"`

	// An AWS KMS key identifier that is used to encrypt the data on the replication
	// instance.
	//
	// If you don't specify a value for the KmsKeyId parameter, then AWS DMS uses
	// your default encryption key.
	//
	// AWS KMS creates the default encryption key for your AWS account. Your AWS
	// account has a different default encryption key for each AWS Region.
	KmsKeyId *string `type:"string"`

	// Specifies whether the replication instance is a Multi-AZ deployment. You
	// can't set the AvailabilityZone parameter if the Multi-AZ parameter is set
	// to true.
	MultiAZ *bool `type:"boolean"`

	// The weekly time range during which system maintenance can occur, in Universal
	// Coordinated Time (UTC).
	//
	// Format: ddd:hh24:mi-ddd:hh24:mi
	//
	// Default: A 30-minute window selected at random from an 8-hour block of time
	// per AWS Region, occurring on a random day of the week.
	//
	// Valid Days: Mon, Tue, Wed, Thu, Fri, Sat, Sun
	//
	// Constraints: Minimum 30-minute window.
	PreferredMaintenanceWindow *string `type:"string"`

	// Specifies the accessibility options for the replication instance. A value
	// of true represents an instance with a public IP address. A value of false
	// represents an instance with a private IP address. The default value is true.
	PubliclyAccessible *bool `type:"boolean"`

	// The compute and memory capacity of the replication instance as specified
	// by the replication instance class.
	//
	// Valid Values: dms.t2.micro | dms.t2.small | dms.t2.medium | dms.t2.large
	// | dms.c4.large | dms.c4.xlarge | dms.c4.2xlarge | dms.c4.4xlarge
	//
	// ReplicationInstanceClass is a required field
	ReplicationInstanceClass *string `type:"string" required:"true"`

	// The replication instance identifier. This parameter is stored as a lowercase
	// string.
	//
	// Constraints:
	//
	//    * Must contain from 1 to 63 alphanumeric characters or hyphens.
	//
	//    * First character must be a letter.
	//
	//    * Can't end with a hyphen or contain two consecutive hyphens.
	//
	// Example: myrepinstance
	//
	// ReplicationInstanceIdentifier is a required field
	ReplicationInstanceIdentifier *string `type:"string" required:"true"`

	// A subnet group to associate with the replication instance.
	ReplicationSubnetGroupIdentifier *string `type:"string"`

	// One or more tags to be assigned to the replication instance.
	Tags []Tag `type:"list"`

	// Specifies the VPC security group to be used with the replication instance.
	// The VPC security group must work with the VPC containing the replication
	// instance.
	VpcSecurityGroupIds []string `type:"list"`
}

// String returns the string representation
func (s CreateReplicationInstanceInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateReplicationInstanceInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateReplicationInstanceInput"}

	if s.ReplicationInstanceClass == nil {
		invalidParams.Add(aws.NewErrParamRequired("ReplicationInstanceClass"))
	}

	if s.ReplicationInstanceIdentifier == nil {
		invalidParams.Add(aws.NewErrParamRequired("ReplicationInstanceIdentifier"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type CreateReplicationInstanceOutput struct {
	_ struct{} `type:"structure"`

	// The replication instance that was created.
	ReplicationInstance *ReplicationInstance `type:"structure"`
}

// String returns the string representation
func (s CreateReplicationInstanceOutput) String() string {
	return awsutil.Prettify(s)
}

const opCreateReplicationInstance = "CreateReplicationInstance"

// CreateReplicationInstanceRequest returns a request value for making API operation for
// AWS Database Migration Service.
//
// Creates the replication instance using the specified parameters.
//
// AWS DMS requires that your account have certain roles with appropriate permissions
// before you can create a replication instance. For information on the required
// roles, see Creating the IAM Roles to Use With the AWS CLI and AWS DMS API
// (https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Security.APIRole.html).
// For information on the required permissions, see IAM Permissions Needed to
// Use AWS DMS (https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Security.IAMPermissions.html).
//
//    // Example sending a request using CreateReplicationInstanceRequest.
//    req := client.CreateReplicationInstanceRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/dms-2016-01-01/CreateReplicationInstance
func (c *Client) CreateReplicationInstanceRequest(input *CreateReplicationInstanceInput) CreateReplicationInstanceRequest {
	op := &aws.Operation{
		Name:       opCreateReplicationInstance,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &CreateReplicationInstanceInput{}
	}

	req := c.newRequest(op, input, &CreateReplicationInstanceOutput{})

	return CreateReplicationInstanceRequest{Request: req, Input: input, Copy: c.CreateReplicationInstanceRequest}
}

// CreateReplicationInstanceRequest is the request type for the
// CreateReplicationInstance API operation.
type CreateReplicationInstanceRequest struct {
	*aws.Request
	Input *CreateReplicationInstanceInput
	Copy  func(*CreateReplicationInstanceInput) CreateReplicationInstanceRequest
}

// Send marshals and sends the CreateReplicationInstance API request.
func (r CreateReplicationInstanceRequest) Send(ctx context.Context) (*CreateReplicationInstanceResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateReplicationInstanceResponse{
		CreateReplicationInstanceOutput: r.Request.Data.(*CreateReplicationInstanceOutput),
		response:                        &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateReplicationInstanceResponse is the response type for the
// CreateReplicationInstance API operation.
type CreateReplicationInstanceResponse struct {
	*CreateReplicationInstanceOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateReplicationInstance request.
func (r *CreateReplicationInstanceResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
