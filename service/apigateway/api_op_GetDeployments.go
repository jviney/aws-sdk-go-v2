// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package apigateway

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
	"github.com/jviney/aws-sdk-go-v2/private/protocol"
)

// Requests API Gateway to get information about a Deployments collection.
type GetDeploymentsInput struct {
	_ struct{} `type:"structure"`

	// The maximum number of returned results per page. The default value is 25
	// and the maximum value is 500.
	Limit *int64 `location:"querystring" locationName:"limit" type:"integer"`

	// The current pagination position in the paged result set.
	Position *string `location:"querystring" locationName:"position" type:"string"`

	// [Required] The string identifier of the associated RestApi.
	//
	// RestApiId is a required field
	RestApiId *string `location:"uri" locationName:"restapi_id" type:"string" required:"true"`
}

// String returns the string representation
func (s GetDeploymentsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetDeploymentsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetDeploymentsInput"}

	if s.RestApiId == nil {
		invalidParams.Add(aws.NewErrParamRequired("RestApiId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetDeploymentsInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.RestApiId != nil {
		v := *s.RestApiId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "restapi_id", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Limit != nil {
		v := *s.Limit

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "limit", protocol.Int64Value(v), metadata)
	}
	if s.Position != nil {
		v := *s.Position

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "position", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// Represents a collection resource that contains zero or more references to
// your existing deployments, and links that guide you on how to interact with
// your collection. The collection offers a paginated view of the contained
// deployments.
//
// To create a new deployment of a RestApi, make a POST request against this
// resource. To view, update, or delete an existing deployment, make a GET,
// PATCH, or DELETE request, respectively, on a specified Deployment resource.
//
// Deploying an API (https://docs.aws.amazon.com/apigateway/latest/developerguide/how-to-deploy-api.html),
// AWS CLI (https://docs.aws.amazon.com/cli/latest/reference/apigateway/get-deployment.html),
// AWS SDKs (https://aws.amazon.com/tools/)
type GetDeploymentsOutput struct {
	_ struct{} `type:"structure"`

	// The current page of elements from this collection.
	Items []Deployment `locationName:"item" type:"list"`

	Position *string `locationName:"position" type:"string"`
}

// String returns the string representation
func (s GetDeploymentsOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetDeploymentsOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.Items != nil {
		v := s.Items

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "item", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	if s.Position != nil {
		v := *s.Position

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "position", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

const opGetDeployments = "GetDeployments"

// GetDeploymentsRequest returns a request value for making API operation for
// Amazon API Gateway.
//
// Gets information about a Deployments collection.
//
//    // Example sending a request using GetDeploymentsRequest.
//    req := client.GetDeploymentsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
func (c *Client) GetDeploymentsRequest(input *GetDeploymentsInput) GetDeploymentsRequest {
	op := &aws.Operation{
		Name:       opGetDeployments,
		HTTPMethod: "GET",
		HTTPPath:   "/restapis/{restapi_id}/deployments",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"position"},
			OutputTokens:    []string{"position"},
			LimitToken:      "limit",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &GetDeploymentsInput{}
	}

	req := c.newRequest(op, input, &GetDeploymentsOutput{})

	return GetDeploymentsRequest{Request: req, Input: input, Copy: c.GetDeploymentsRequest}
}

// GetDeploymentsRequest is the request type for the
// GetDeployments API operation.
type GetDeploymentsRequest struct {
	*aws.Request
	Input *GetDeploymentsInput
	Copy  func(*GetDeploymentsInput) GetDeploymentsRequest
}

// Send marshals and sends the GetDeployments API request.
func (r GetDeploymentsRequest) Send(ctx context.Context) (*GetDeploymentsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetDeploymentsResponse{
		GetDeploymentsOutput: r.Request.Data.(*GetDeploymentsOutput),
		response:             &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewGetDeploymentsRequestPaginator returns a paginator for GetDeployments.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.GetDeploymentsRequest(input)
//   p := apigateway.NewGetDeploymentsRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewGetDeploymentsPaginator(req GetDeploymentsRequest) GetDeploymentsPaginator {
	return GetDeploymentsPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *GetDeploymentsInput
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

// GetDeploymentsPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type GetDeploymentsPaginator struct {
	aws.Pager
}

func (p *GetDeploymentsPaginator) CurrentPage() *GetDeploymentsOutput {
	return p.Pager.CurrentPage().(*GetDeploymentsOutput)
}

// GetDeploymentsResponse is the response type for the
// GetDeployments API operation.
type GetDeploymentsResponse struct {
	*GetDeploymentsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetDeployments request.
func (r *GetDeploymentsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
