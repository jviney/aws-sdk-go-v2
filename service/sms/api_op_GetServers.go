// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package sms

import (
	"context"
	"time"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
)

type GetServersInput struct {
	_ struct{} `type:"structure"`

	// The maximum number of results to return in a single call. The default value
	// is 50. To retrieve the remaining results, make another call with the returned
	// NextToken value.
	MaxResults *int64 `locationName:"maxResults" type:"integer"`

	// The token for the next set of results.
	NextToken *string `locationName:"nextToken" type:"string"`

	// List of VmServerAddress objects
	VmServerAddressList []VmServerAddress `locationName:"vmServerAddressList" type:"list"`
}

// String returns the string representation
func (s GetServersInput) String() string {
	return awsutil.Prettify(s)
}

type GetServersOutput struct {
	_ struct{} `type:"structure"`

	// The time when the server was last modified.
	LastModifiedOn *time.Time `locationName:"lastModifiedOn" type:"timestamp"`

	// The token required to retrieve the next set of results. This value is null
	// when there are no more results to return.
	NextToken *string `locationName:"nextToken" type:"string"`

	// The status of the server catalog.
	ServerCatalogStatus ServerCatalogStatus `locationName:"serverCatalogStatus" type:"string" enum:"true"`

	// Information about the servers.
	ServerList []Server `locationName:"serverList" type:"list"`
}

// String returns the string representation
func (s GetServersOutput) String() string {
	return awsutil.Prettify(s)
}

const opGetServers = "GetServers"

// GetServersRequest returns a request value for making API operation for
// AWS Server Migration Service.
//
// Describes the servers in your server catalog.
//
// Before you can describe your servers, you must import them using ImportServerCatalog.
//
//    // Example sending a request using GetServersRequest.
//    req := client.GetServersRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/sms-2016-10-24/GetServers
func (c *Client) GetServersRequest(input *GetServersInput) GetServersRequest {
	op := &aws.Operation{
		Name:       opGetServers,
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
		input = &GetServersInput{}
	}

	req := c.newRequest(op, input, &GetServersOutput{})

	return GetServersRequest{Request: req, Input: input, Copy: c.GetServersRequest}
}

// GetServersRequest is the request type for the
// GetServers API operation.
type GetServersRequest struct {
	*aws.Request
	Input *GetServersInput
	Copy  func(*GetServersInput) GetServersRequest
}

// Send marshals and sends the GetServers API request.
func (r GetServersRequest) Send(ctx context.Context) (*GetServersResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetServersResponse{
		GetServersOutput: r.Request.Data.(*GetServersOutput),
		response:         &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewGetServersRequestPaginator returns a paginator for GetServers.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.GetServersRequest(input)
//   p := sms.NewGetServersRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewGetServersPaginator(req GetServersRequest) GetServersPaginator {
	return GetServersPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *GetServersInput
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

// GetServersPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type GetServersPaginator struct {
	aws.Pager
}

func (p *GetServersPaginator) CurrentPage() *GetServersOutput {
	return p.Pager.CurrentPage().(*GetServersOutput)
}

// GetServersResponse is the response type for the
// GetServers API operation.
type GetServersResponse struct {
	*GetServersOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetServers request.
func (r *GetServersResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
