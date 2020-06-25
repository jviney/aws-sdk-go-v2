// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package codedeploy

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
)

// Represents the input of a ListDeploymentInstances operation.
type ListDeploymentInstancesInput struct {
	_ struct{} `type:"structure"`

	// The unique ID of a deployment.
	//
	// DeploymentId is a required field
	DeploymentId *string `locationName:"deploymentId" type:"string" required:"true"`

	// A subset of instances to list by status:
	//
	//    * Pending: Include those instances with pending deployments.
	//
	//    * InProgress: Include those instances where deployments are still in progress.
	//
	//    * Succeeded: Include those instances with successful deployments.
	//
	//    * Failed: Include those instances with failed deployments.
	//
	//    * Skipped: Include those instances with skipped deployments.
	//
	//    * Unknown: Include those instances with deployments in an unknown state.
	InstanceStatusFilter []InstanceStatus `locationName:"instanceStatusFilter" type:"list"`

	// The set of instances in a blue/green deployment, either those in the original
	// environment ("BLUE") or those in the replacement environment ("GREEN"), for
	// which you want to view instance information.
	InstanceTypeFilter []InstanceType `locationName:"instanceTypeFilter" type:"list"`

	// An identifier returned from the previous list deployment instances call.
	// It can be used to return the next set of deployment instances in the list.
	NextToken *string `locationName:"nextToken" type:"string"`
}

// String returns the string representation
func (s ListDeploymentInstancesInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListDeploymentInstancesInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListDeploymentInstancesInput"}

	if s.DeploymentId == nil {
		invalidParams.Add(aws.NewErrParamRequired("DeploymentId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Represents the output of a ListDeploymentInstances operation.
type ListDeploymentInstancesOutput struct {
	_ struct{} `type:"structure"`

	// A list of instance IDs.
	InstancesList []string `locationName:"instancesList" type:"list"`

	// If a large amount of information is returned, an identifier is also returned.
	// It can be used in a subsequent list deployment instances call to return the
	// next set of deployment instances in the list.
	NextToken *string `locationName:"nextToken" type:"string"`
}

// String returns the string representation
func (s ListDeploymentInstancesOutput) String() string {
	return awsutil.Prettify(s)
}

const opListDeploymentInstances = "ListDeploymentInstances"

// ListDeploymentInstancesRequest returns a request value for making API operation for
// AWS CodeDeploy.
//
//
// The newer BatchGetDeploymentTargets should be used instead because it works
// with all compute types. ListDeploymentInstances throws an exception if it
// is used with a compute platform other than EC2/On-premises or AWS Lambda.
//
// Lists the instance for a deployment associated with the IAM user or AWS account.
//
//    // Example sending a request using ListDeploymentInstancesRequest.
//    req := client.ListDeploymentInstancesRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/codedeploy-2014-10-06/ListDeploymentInstances
func (c *Client) ListDeploymentInstancesRequest(input *ListDeploymentInstancesInput) ListDeploymentInstancesRequest {
	if c.Client.Config.Logger != nil {
		c.Client.Config.Logger.Log("This operation, ListDeploymentInstances, has been deprecated")
	}
	op := &aws.Operation{
		Name:       opListDeploymentInstances,
		HTTPMethod: "POST",
		HTTPPath:   "/",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"nextToken"},
			OutputTokens:    []string{"nextToken"},
			LimitToken:      "",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &ListDeploymentInstancesInput{}
	}

	req := c.newRequest(op, input, &ListDeploymentInstancesOutput{})

	return ListDeploymentInstancesRequest{Request: req, Input: input, Copy: c.ListDeploymentInstancesRequest}
}

// ListDeploymentInstancesRequest is the request type for the
// ListDeploymentInstances API operation.
type ListDeploymentInstancesRequest struct {
	*aws.Request
	Input *ListDeploymentInstancesInput
	Copy  func(*ListDeploymentInstancesInput) ListDeploymentInstancesRequest
}

// Send marshals and sends the ListDeploymentInstances API request.
func (r ListDeploymentInstancesRequest) Send(ctx context.Context) (*ListDeploymentInstancesResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListDeploymentInstancesResponse{
		ListDeploymentInstancesOutput: r.Request.Data.(*ListDeploymentInstancesOutput),
		response:                      &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewListDeploymentInstancesRequestPaginator returns a paginator for ListDeploymentInstances.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.ListDeploymentInstancesRequest(input)
//   p := codedeploy.NewListDeploymentInstancesRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewListDeploymentInstancesPaginator(req ListDeploymentInstancesRequest) ListDeploymentInstancesPaginator {
	return ListDeploymentInstancesPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *ListDeploymentInstancesInput
				if req.Input != nil {
					tmp := *req.Input
					inCpy = &tmp
				}

				newReq := req.Copy(inCpy)
				newReq.SetContext(ctx)
				return newReq.Request, nil
			},
		},
	}
}

// ListDeploymentInstancesPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type ListDeploymentInstancesPaginator struct {
	aws.Pager
}

func (p *ListDeploymentInstancesPaginator) CurrentPage() *ListDeploymentInstancesOutput {
	return p.Pager.CurrentPage().(*ListDeploymentInstancesOutput)
}

// ListDeploymentInstancesResponse is the response type for the
// ListDeploymentInstances API operation.
type ListDeploymentInstancesResponse struct {
	*ListDeploymentInstancesOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListDeploymentInstances request.
func (r *ListDeploymentInstancesResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
