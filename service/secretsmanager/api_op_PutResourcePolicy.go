// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package secretsmanager

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
)

type PutResourcePolicyInput struct {
	_ struct{} `type:"structure"`

	// A JSON-formatted string that's constructed according to the grammar and syntax
	// for an AWS resource-based policy. The policy in the string identifies who
	// can access or manage this secret and its versions. For information on how
	// to format a JSON parameter for the various command line tool environments,
	// see Using JSON for Parameters (http://docs.aws.amazon.com/cli/latest/userguide/cli-using-param.html#cli-using-param-json)
	// in the AWS CLI User Guide.
	//
	// ResourcePolicy is a required field
	ResourcePolicy *string `min:"1" type:"string" required:"true"`

	// Specifies the secret that you want to attach the resource-based policy to.
	// You can specify either the ARN or the friendly name of the secret.
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
func (s PutResourcePolicyInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *PutResourcePolicyInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "PutResourcePolicyInput"}

	if s.ResourcePolicy == nil {
		invalidParams.Add(aws.NewErrParamRequired("ResourcePolicy"))
	}
	if s.ResourcePolicy != nil && len(*s.ResourcePolicy) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ResourcePolicy", 1))
	}

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

type PutResourcePolicyOutput struct {
	_ struct{} `type:"structure"`

	// The ARN of the secret that the resource-based policy was retrieved for.
	ARN *string `min:"20" type:"string"`

	// The friendly name of the secret that the resource-based policy was retrieved
	// for.
	Name *string `min:"1" type:"string"`
}

// String returns the string representation
func (s PutResourcePolicyOutput) String() string {
	return awsutil.Prettify(s)
}

const opPutResourcePolicy = "PutResourcePolicy"

// PutResourcePolicyRequest returns a request value for making API operation for
// AWS Secrets Manager.
//
// Attaches the contents of the specified resource-based permission policy to
// a secret. A resource-based policy is optional. Alternatively, you can use
// IAM identity-based policies that specify the secret's Amazon Resource Name
// (ARN) in the policy statement's Resources element. You can also use a combination
// of both identity-based and resource-based policies. The affected users and
// roles receive the permissions that are permitted by all of the relevant policies.
// For more information, see Using Resource-Based Policies for AWS Secrets Manager
// (http://docs.aws.amazon.com/secretsmanager/latest/userguide/auth-and-access_resource-based-policies.html).
// For the complete description of the AWS policy syntax and grammar, see IAM
// JSON Policy Reference (https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_policies.html)
// in the IAM User Guide.
//
// Minimum permissions
//
// To run this command, you must have the following permissions:
//
//    * secretsmanager:PutResourcePolicy
//
// Related operations
//
//    * To retrieve the resource policy that's attached to a secret, use GetResourcePolicy.
//
//    * To delete the resource-based policy that's attached to a secret, use
//    DeleteResourcePolicy.
//
//    * To list all of the currently available secrets, use ListSecrets.
//
//    // Example sending a request using PutResourcePolicyRequest.
//    req := client.PutResourcePolicyRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/secretsmanager-2017-10-17/PutResourcePolicy
func (c *Client) PutResourcePolicyRequest(input *PutResourcePolicyInput) PutResourcePolicyRequest {
	op := &aws.Operation{
		Name:       opPutResourcePolicy,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &PutResourcePolicyInput{}
	}

	req := c.newRequest(op, input, &PutResourcePolicyOutput{})

	return PutResourcePolicyRequest{Request: req, Input: input, Copy: c.PutResourcePolicyRequest}
}

// PutResourcePolicyRequest is the request type for the
// PutResourcePolicy API operation.
type PutResourcePolicyRequest struct {
	*aws.Request
	Input *PutResourcePolicyInput
	Copy  func(*PutResourcePolicyInput) PutResourcePolicyRequest
}

// Send marshals and sends the PutResourcePolicy API request.
func (r PutResourcePolicyRequest) Send(ctx context.Context) (*PutResourcePolicyResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &PutResourcePolicyResponse{
		PutResourcePolicyOutput: r.Request.Data.(*PutResourcePolicyOutput),
		response:                &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// PutResourcePolicyResponse is the response type for the
// PutResourcePolicy API operation.
type PutResourcePolicyResponse struct {
	*PutResourcePolicyOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// PutResourcePolicy request.
func (r *PutResourcePolicyResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
