// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package codecommit

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
)

type UpdatePullRequestApprovalRuleContentInput struct {
	_ struct{} `type:"structure"`

	// The name of the approval rule you want to update.
	//
	// ApprovalRuleName is a required field
	ApprovalRuleName *string `locationName:"approvalRuleName" min:"1" type:"string" required:"true"`

	// The SHA-256 hash signature for the content of the approval rule. You can
	// retrieve this information by using GetPullRequest.
	ExistingRuleContentSha256 *string `locationName:"existingRuleContentSha256" type:"string"`

	// The updated content for the approval rule.
	//
	// When you update the content of the approval rule, you can specify approvers
	// in an approval pool in one of two ways:
	//
	//    * CodeCommitApprovers: This option only requires an AWS account and a
	//    resource. It can be used for both IAM users and federated access users
	//    whose name matches the provided resource name. This is a very powerful
	//    option that offers a great deal of flexibility. For example, if you specify
	//    the AWS account 123456789012 and Mary_Major, all of the following are
	//    counted as approvals coming from that user: An IAM user in the account
	//    (arn:aws:iam::123456789012:user/Mary_Major) A federated user identified
	//    in IAM as Mary_Major (arn:aws:sts::123456789012:federated-user/Mary_Major)
	//    This option does not recognize an active session of someone assuming the
	//    role of CodeCommitReview with a role session name of Mary_Major (arn:aws:sts::123456789012:assumed-role/CodeCommitReview/Mary_Major)
	//    unless you include a wildcard (*Mary_Major).
	//
	//    * Fully qualified ARN: This option allows you to specify the fully qualified
	//    Amazon Resource Name (ARN) of the IAM user or role.
	//
	// For more information about IAM ARNs, wildcards, and formats, see IAM Identifiers
	// (https://docs.aws.amazon.com/iam/latest/UserGuide/reference_identifiers.html)
	// in the IAM User Guide.
	//
	// NewRuleContent is a required field
	NewRuleContent *string `locationName:"newRuleContent" min:"1" type:"string" required:"true"`

	// The system-generated ID of the pull request.
	//
	// PullRequestId is a required field
	PullRequestId *string `locationName:"pullRequestId" type:"string" required:"true"`
}

// String returns the string representation
func (s UpdatePullRequestApprovalRuleContentInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UpdatePullRequestApprovalRuleContentInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "UpdatePullRequestApprovalRuleContentInput"}

	if s.ApprovalRuleName == nil {
		invalidParams.Add(aws.NewErrParamRequired("ApprovalRuleName"))
	}
	if s.ApprovalRuleName != nil && len(*s.ApprovalRuleName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ApprovalRuleName", 1))
	}

	if s.NewRuleContent == nil {
		invalidParams.Add(aws.NewErrParamRequired("NewRuleContent"))
	}
	if s.NewRuleContent != nil && len(*s.NewRuleContent) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("NewRuleContent", 1))
	}

	if s.PullRequestId == nil {
		invalidParams.Add(aws.NewErrParamRequired("PullRequestId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type UpdatePullRequestApprovalRuleContentOutput struct {
	_ struct{} `type:"structure"`

	// Information about the updated approval rule.
	//
	// ApprovalRule is a required field
	ApprovalRule *ApprovalRule `locationName:"approvalRule" type:"structure" required:"true"`
}

// String returns the string representation
func (s UpdatePullRequestApprovalRuleContentOutput) String() string {
	return awsutil.Prettify(s)
}

const opUpdatePullRequestApprovalRuleContent = "UpdatePullRequestApprovalRuleContent"

// UpdatePullRequestApprovalRuleContentRequest returns a request value for making API operation for
// AWS CodeCommit.
//
// Updates the structure of an approval rule created specifically for a pull
// request. For example, you can change the number of required approvers and
// the approval pool for approvers.
//
//    // Example sending a request using UpdatePullRequestApprovalRuleContentRequest.
//    req := client.UpdatePullRequestApprovalRuleContentRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/codecommit-2015-04-13/UpdatePullRequestApprovalRuleContent
func (c *Client) UpdatePullRequestApprovalRuleContentRequest(input *UpdatePullRequestApprovalRuleContentInput) UpdatePullRequestApprovalRuleContentRequest {
	op := &aws.Operation{
		Name:       opUpdatePullRequestApprovalRuleContent,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &UpdatePullRequestApprovalRuleContentInput{}
	}

	req := c.newRequest(op, input, &UpdatePullRequestApprovalRuleContentOutput{})

	return UpdatePullRequestApprovalRuleContentRequest{Request: req, Input: input, Copy: c.UpdatePullRequestApprovalRuleContentRequest}
}

// UpdatePullRequestApprovalRuleContentRequest is the request type for the
// UpdatePullRequestApprovalRuleContent API operation.
type UpdatePullRequestApprovalRuleContentRequest struct {
	*aws.Request
	Input *UpdatePullRequestApprovalRuleContentInput
	Copy  func(*UpdatePullRequestApprovalRuleContentInput) UpdatePullRequestApprovalRuleContentRequest
}

// Send marshals and sends the UpdatePullRequestApprovalRuleContent API request.
func (r UpdatePullRequestApprovalRuleContentRequest) Send(ctx context.Context) (*UpdatePullRequestApprovalRuleContentResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &UpdatePullRequestApprovalRuleContentResponse{
		UpdatePullRequestApprovalRuleContentOutput: r.Request.Data.(*UpdatePullRequestApprovalRuleContentOutput),
		response: &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// UpdatePullRequestApprovalRuleContentResponse is the response type for the
// UpdatePullRequestApprovalRuleContent API operation.
type UpdatePullRequestApprovalRuleContentResponse struct {
	*UpdatePullRequestApprovalRuleContentOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// UpdatePullRequestApprovalRuleContent request.
func (r *UpdatePullRequestApprovalRuleContentResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
