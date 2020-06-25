// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package secretsmanager

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
)

type CancelRotateSecretInput struct {
	_ struct{} `type:"structure"`

	// Specifies the secret for which you want to cancel a rotation request. You
	// can specify either the Amazon Resource Name (ARN) or the friendly name of
	// the secret.
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
func (s CancelRotateSecretInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CancelRotateSecretInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CancelRotateSecretInput"}

	if s.SecretId == nil {
		invalidParams.Add(aws.NewErrParamRequired("SecretId"))
	}
	if s.SecretId != nil && len(*s.SecretId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("SecretId", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type CancelRotateSecretOutput struct {
	_ struct{} `type:"structure"`

	// The ARN of the secret for which rotation was canceled.
	ARN *string `min:"20" type:"string"`

	// The friendly name of the secret for which rotation was canceled.
	Name *string `min:"1" type:"string"`

	// The unique identifier of the version of the secret that was created during
	// the rotation. This version might not be complete, and should be evaluated
	// for possible deletion. At the very least, you should remove the VersionStage
	// value AWSPENDING to enable this version to be deleted. Failing to clean up
	// a cancelled rotation can block you from successfully starting future rotations.
	VersionId *string `min:"32" type:"string"`
}

// String returns the string representation
func (s CancelRotateSecretOutput) String() string {
	return awsutil.Prettify(s)
}

const opCancelRotateSecret = "CancelRotateSecret"

// CancelRotateSecretRequest returns a request value for making API operation for
// AWS Secrets Manager.
//
// Disables automatic scheduled rotation and cancels the rotation of a secret
// if one is currently in progress.
//
// To re-enable scheduled rotation, call RotateSecret with AutomaticallyRotateAfterDays
// set to a value greater than 0. This will immediately rotate your secret and
// then enable the automatic schedule.
//
// If you cancel a rotation that is in progress, it can leave the VersionStage
// labels in an unexpected state. Depending on what step of the rotation was
// in progress, you might need to remove the staging label AWSPENDING from the
// partially created version, specified by the VersionId response value. You
// should also evaluate the partially rotated new version to see if it should
// be deleted, which you can do by removing all staging labels from the new
// version's VersionStage field.
//
// To successfully start a rotation, the staging label AWSPENDING must be in
// one of the following states:
//
//    * Not be attached to any version at all
//
//    * Attached to the same version as the staging label AWSCURRENT
//
// If the staging label AWSPENDING is attached to a different version than the
// version with AWSCURRENT then the attempt to rotate fails.
//
// Minimum permissions
//
// To run this command, you must have the following permissions:
//
//    * secretsmanager:CancelRotateSecret
//
// Related operations
//
//    * To configure rotation for a secret or to manually trigger a rotation,
//    use RotateSecret.
//
//    * To get the rotation configuration details for a secret, use DescribeSecret.
//
//    * To list all of the currently available secrets, use ListSecrets.
//
//    * To list all of the versions currently associated with a secret, use
//    ListSecretVersionIds.
//
//    // Example sending a request using CancelRotateSecretRequest.
//    req := client.CancelRotateSecretRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/secretsmanager-2017-10-17/CancelRotateSecret
func (c *Client) CancelRotateSecretRequest(input *CancelRotateSecretInput) CancelRotateSecretRequest {
	op := &aws.Operation{
		Name:       opCancelRotateSecret,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &CancelRotateSecretInput{}
	}

	req := c.newRequest(op, input, &CancelRotateSecretOutput{})

	return CancelRotateSecretRequest{Request: req, Input: input, Copy: c.CancelRotateSecretRequest}
}

// CancelRotateSecretRequest is the request type for the
// CancelRotateSecret API operation.
type CancelRotateSecretRequest struct {
	*aws.Request
	Input *CancelRotateSecretInput
	Copy  func(*CancelRotateSecretInput) CancelRotateSecretRequest
}

// Send marshals and sends the CancelRotateSecret API request.
func (r CancelRotateSecretRequest) Send(ctx context.Context) (*CancelRotateSecretResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CancelRotateSecretResponse{
		CancelRotateSecretOutput: r.Request.Data.(*CancelRotateSecretOutput),
		response:                 &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CancelRotateSecretResponse is the response type for the
// CancelRotateSecret API operation.
type CancelRotateSecretResponse struct {
	*CancelRotateSecretOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CancelRotateSecret request.
func (r *CancelRotateSecretResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
