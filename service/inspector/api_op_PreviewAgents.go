// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package inspector

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
)

type PreviewAgentsInput struct {
	_ struct{} `type:"structure"`

	// You can use this parameter to indicate the maximum number of items you want
	// in the response. The default value is 10. The maximum value is 500.
	MaxResults *int64 `locationName:"maxResults" type:"integer"`

	// You can use this parameter when paginating results. Set the value of this
	// parameter to null on your first call to the PreviewAgents action. Subsequent
	// calls to the action fill nextToken in the request with the value of NextToken
	// from the previous response to continue listing data.
	NextToken *string `locationName:"nextToken" min:"1" type:"string"`

	// The ARN of the assessment target whose agents you want to preview.
	//
	// PreviewAgentsArn is a required field
	PreviewAgentsArn *string `locationName:"previewAgentsArn" min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s PreviewAgentsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *PreviewAgentsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "PreviewAgentsInput"}
	if s.NextToken != nil && len(*s.NextToken) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("NextToken", 1))
	}

	if s.PreviewAgentsArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("PreviewAgentsArn"))
	}
	if s.PreviewAgentsArn != nil && len(*s.PreviewAgentsArn) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("PreviewAgentsArn", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type PreviewAgentsOutput struct {
	_ struct{} `type:"structure"`

	// The resulting list of agents.
	//
	// AgentPreviews is a required field
	AgentPreviews []AgentPreview `locationName:"agentPreviews" type:"list" required:"true"`

	// When a response is generated, if there is more data to be listed, this parameter
	// is present in the response and contains the value to use for the nextToken
	// parameter in a subsequent pagination request. If there is no more data to
	// be listed, this parameter is set to null.
	NextToken *string `locationName:"nextToken" min:"1" type:"string"`
}

// String returns the string representation
func (s PreviewAgentsOutput) String() string {
	return awsutil.Prettify(s)
}

const opPreviewAgents = "PreviewAgents"

// PreviewAgentsRequest returns a request value for making API operation for
// Amazon Inspector.
//
// Previews the agents installed on the EC2 instances that are part of the specified
// assessment target.
//
//    // Example sending a request using PreviewAgentsRequest.
//    req := client.PreviewAgentsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/inspector-2016-02-16/PreviewAgents
func (c *Client) PreviewAgentsRequest(input *PreviewAgentsInput) PreviewAgentsRequest {
	op := &aws.Operation{
		Name:       opPreviewAgents,
		HTTPMethod: "POST",
		HTTPPath:   "/",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"nextToken"},
			OutputTokens:    []string{"nextToken"},
			LimitToken:      "maxResults",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &PreviewAgentsInput{}
	}

	req := c.newRequest(op, input, &PreviewAgentsOutput{})

	return PreviewAgentsRequest{Request: req, Input: input, Copy: c.PreviewAgentsRequest}
}

// PreviewAgentsRequest is the request type for the
// PreviewAgents API operation.
type PreviewAgentsRequest struct {
	*aws.Request
	Input *PreviewAgentsInput
	Copy  func(*PreviewAgentsInput) PreviewAgentsRequest
}

// Send marshals and sends the PreviewAgents API request.
func (r PreviewAgentsRequest) Send(ctx context.Context) (*PreviewAgentsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &PreviewAgentsResponse{
		PreviewAgentsOutput: r.Request.Data.(*PreviewAgentsOutput),
		response:            &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewPreviewAgentsRequestPaginator returns a paginator for PreviewAgents.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.PreviewAgentsRequest(input)
//   p := inspector.NewPreviewAgentsRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewPreviewAgentsPaginator(req PreviewAgentsRequest) PreviewAgentsPaginator {
	return PreviewAgentsPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *PreviewAgentsInput
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

// PreviewAgentsPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type PreviewAgentsPaginator struct {
	aws.Pager
}

func (p *PreviewAgentsPaginator) CurrentPage() *PreviewAgentsOutput {
	return p.Pager.CurrentPage().(*PreviewAgentsOutput)
}

// PreviewAgentsResponse is the response type for the
// PreviewAgents API operation.
type PreviewAgentsResponse struct {
	*PreviewAgentsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// PreviewAgents request.
func (r *PreviewAgentsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
