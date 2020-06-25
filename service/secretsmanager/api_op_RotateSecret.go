// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package secretsmanager

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
)

type RotateSecretInput struct {
	_ struct{} `type:"structure"`

	// (Optional) Specifies a unique identifier for the new version of the secret
	// that helps ensure idempotency.
	//
	// If you use the AWS CLI or one of the AWS SDK to call this operation, then
	// you can leave this parameter empty. The CLI or SDK generates a random UUID
	// for you and includes that in the request for this parameter. If you don't
	// use the SDK and instead generate a raw HTTP request to the Secrets Manager
	// service endpoint, then you must generate a ClientRequestToken yourself for
	// new versions and include that value in the request.
	//
	// You only need to specify your own value if you are implementing your own
	// retry logic and want to ensure that a given secret is not created twice.
	// We recommend that you generate a UUID-type (https://wikipedia.org/wiki/Universally_unique_identifier)
	// value to ensure uniqueness within the specified secret.
	//
	// Secrets Manager uses this value to prevent the accidental creation of duplicate
	// versions if there are failures and retries during the function's processing.
	// This value becomes the VersionId of the new version.
	ClientRequestToken *string `min:"32" type:"string" idempotencyToken:"true"`

	// (Optional) Specifies the ARN of the Lambda function that can rotate the secret.
	RotationLambdaARN *string `type:"string"`

	// A structure that defines the rotation configuration for this secret.
	RotationRules *RotationRulesType `type:"structure"`

	// Specifies the secret that you want to rotate. You can specify either the
	// Amazon Resource Name (ARN) or the friendly name of the secret.
	//
	// If you specify an ARN, we generally recommend that you specify a complete
	// ARN. You can specify a partial ARN too—for example, if you don’t include
	// the final hyphen and six random characters that Secrets Manager adds at the
	// end of the ARN when you created the secret. A partial ARN match can work
	// as long as it uniquely matches only one secret. However, if your secret has
	// a name that ends in a hyphen followed by six characters (before Secrets Manager
	// adds the hyphen and six characters to the ARN) and you try to use that as
	// a partial ARN, then those characters cause Secrets Manager to assume that
	// you’re specifying a complete ARN. This confusion can cause unexpected results.
	// To avoid this situation, we recommend that you don’t create secret names
	// that end with a hyphen followed by six characters.
	//
	// SecretId is a required field
	SecretId *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s RotateSecretInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *RotateSecretInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "RotateSecretInput"}
	if s.ClientRequestToken != nil && len(*s.ClientRequestToken) < 32 {
		invalidParams.Add(aws.NewErrParamMinLen("ClientRequestToken", 32))
	}

	if s.SecretId == nil {
		invalidParams.Add(aws.NewErrParamRequired("SecretId"))
	}
	if s.SecretId != nil && len(*s.SecretId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("SecretId", 1))
	}
	if s.RotationRules != nil {
		if err := s.RotationRules.Validate(); err != nil {
			invalidParams.AddNested("RotationRules", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type RotateSecretOutput struct {
	_ struct{} `type:"structure"`

	// The ARN of the secret.
	ARN *string `min:"20" type:"string"`

	// The friendly name of the secret.
	Name *string `min:"1" type:"string"`

	// The ID of the new version of the secret created by the rotation started by
	// this request.
	VersionId *string `min:"32" type:"string"`
}

// String returns the string representation
func (s RotateSecretOutput) String() string {
	return awsutil.Prettify(s)
}

const opRotateSecret = "RotateSecret"

// RotateSecretRequest returns a request value for making API operation for
// AWS Secrets Manager.
//
// Configures and starts the asynchronous process of rotating this secret. If
// you include the configuration parameters, the operation sets those values
// for the secret and then immediately starts a rotation. If you do not include
// the configuration parameters, the operation starts a rotation with the values
// already stored in the secret. After the rotation completes, the protected
// service and its clients all use the new version of the secret.
//
// This required configuration information includes the ARN of an AWS Lambda
// function and the time between scheduled rotations. The Lambda rotation function
// creates a new version of the secret and creates or updates the credentials
// on the protected service to match. After testing the new credentials, the
// function marks the new secret with the staging label AWSCURRENT so that your
// clients all immediately begin to use the new version. For more information
// about rotating secrets and how to configure a Lambda function to rotate the
// secrets for your protected service, see Rotating Secrets in AWS Secrets Manager
// (https://docs.aws.amazon.com/secretsmanager/latest/userguide/rotating-secrets.html)
// in the AWS Secrets Manager User Guide.
//
// Secrets Manager schedules the next rotation when the previous one is complete.
// Secrets Manager schedules the date by adding the rotation interval (number
// of days) to the actual date of the last rotation. The service chooses the
// hour within that 24-hour date window randomly. The minute is also chosen
// somewhat randomly, but weighted towards the top of the hour and influenced
// by a variety of factors that help distribute load.
//
// The rotation function must end with the versions of the secret in one of
// two states:
//
//    * The AWSPENDING and AWSCURRENT staging labels are attached to the same
//    version of the secret, or
//
//    * The AWSPENDING staging label is not attached to any version of the secret.
//
// If instead the AWSPENDING staging label is present but is not attached to
// the same version as AWSCURRENT then any later invocation of RotateSecret
// assumes that a previous rotation request is still in progress and returns
// an error.
//
// Minimum permissions
//
// To run this command, you must have the following permissions:
//
//    * secretsmanager:RotateSecret
//
//    * lambda:InvokeFunction (on the function specified in the secret's metadata)
//
// Related operations
//
//    * To list the secrets in your account, use ListSecrets.
//
//    * To get the details for a version of a secret, use DescribeSecret.
//
//    * To create a new version of a secret, use CreateSecret.
//
//    * To attach staging labels to or remove staging labels from a version
//    of a secret, use UpdateSecretVersionStage.
//
//    // Example sending a request using RotateSecretRequest.
//    req := client.RotateSecretRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/secretsmanager-2017-10-17/RotateSecret
func (c *Client) RotateSecretRequest(input *RotateSecretInput) RotateSecretRequest {
	op := &aws.Operation{
		Name:       opRotateSecret,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &RotateSecretInput{}
	}

	req := c.newRequest(op, input, &RotateSecretOutput{})

	return RotateSecretRequest{Request: req, Input: input, Copy: c.RotateSecretRequest}
}

// RotateSecretRequest is the request type for the
// RotateSecret API operation.
type RotateSecretRequest struct {
	*aws.Request
	Input *RotateSecretInput
	Copy  func(*RotateSecretInput) RotateSecretRequest
}

// Send marshals and sends the RotateSecret API request.
func (r RotateSecretRequest) Send(ctx context.Context) (*RotateSecretResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &RotateSecretResponse{
		RotateSecretOutput: r.Request.Data.(*RotateSecretOutput),
		response:           &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// RotateSecretResponse is the response type for the
// RotateSecret API operation.
type RotateSecretResponse struct {
	*RotateSecretOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// RotateSecret request.
func (r *RotateSecretResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
