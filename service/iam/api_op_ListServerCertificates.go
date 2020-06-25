// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package iam

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
)

type ListServerCertificatesInput struct {
	_ struct{} `type:"structure"`

	// Use this parameter only when paginating results and only after you receive
	// a response indicating that the results are truncated. Set it to the value
	// of the Marker element in the response that you received to indicate where
	// the next call should start.
	Marker *string `min:"1" type:"string"`

	// Use this only when paginating results to indicate the maximum number of items
	// you want in the response. If additional items exist beyond the maximum you
	// specify, the IsTruncated response element is true.
	//
	// If you do not include this parameter, the number of items defaults to 100.
	// Note that IAM might return fewer results, even when there are more results
	// available. In that case, the IsTruncated response element returns true, and
	// Marker contains a value to include in the subsequent call that tells the
	// service where to continue from.
	MaxItems *int64 `min:"1" type:"integer"`

	// The path prefix for filtering the results. For example: /company/servercerts
	// would get all server certificates for which the path starts with /company/servercerts.
	//
	// This parameter is optional. If it is not included, it defaults to a slash
	// (/), listing all server certificates. This parameter allows (through its
	// regex pattern (http://wikipedia.org/wiki/regex)) a string of characters consisting
	// of either a forward slash (/) by itself or a string that must begin and end
	// with forward slashes. In addition, it can contain any ASCII character from
	// the ! (\u0021) through the DEL character (\u007F), including most punctuation
	// characters, digits, and upper and lowercased letters.
	PathPrefix *string `min:"1" type:"string"`
}

// String returns the string representation
func (s ListServerCertificatesInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListServerCertificatesInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListServerCertificatesInput"}
	if s.Marker != nil && len(*s.Marker) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Marker", 1))
	}
	if s.MaxItems != nil && *s.MaxItems < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxItems", 1))
	}
	if s.PathPrefix != nil && len(*s.PathPrefix) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("PathPrefix", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Contains the response to a successful ListServerCertificates request.
type ListServerCertificatesOutput struct {
	_ struct{} `type:"structure"`

	// A flag that indicates whether there are more items to return. If your results
	// were truncated, you can make a subsequent pagination request using the Marker
	// request parameter to retrieve more items. Note that IAM might return fewer
	// than the MaxItems number of results even when there are more results available.
	// We recommend that you check IsTruncated after every call to ensure that you
	// receive all your results.
	IsTruncated *bool `type:"boolean"`

	// When IsTruncated is true, this element is present and contains the value
	// to use for the Marker parameter in a subsequent pagination request.
	Marker *string `type:"string"`

	// A list of server certificates.
	//
	// ServerCertificateMetadataList is a required field
	ServerCertificateMetadataList []ServerCertificateMetadata `type:"list" required:"true"`
}

// String returns the string representation
func (s ListServerCertificatesOutput) String() string {
	return awsutil.Prettify(s)
}

const opListServerCertificates = "ListServerCertificates"

// ListServerCertificatesRequest returns a request value for making API operation for
// AWS Identity and Access Management.
//
// Lists the server certificates stored in IAM that have the specified path
// prefix. If none exist, the operation returns an empty list.
//
// You can paginate the results using the MaxItems and Marker parameters.
//
// For more information about working with server certificates, see Working
// with Server Certificates (https://docs.aws.amazon.com/IAM/latest/UserGuide/id_credentials_server-certs.html)
// in the IAM User Guide. This topic also includes a list of AWS services that
// can use the server certificates that you manage with IAM.
//
//    // Example sending a request using ListServerCertificatesRequest.
//    req := client.ListServerCertificatesRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/iam-2010-05-08/ListServerCertificates
func (c *Client) ListServerCertificatesRequest(input *ListServerCertificatesInput) ListServerCertificatesRequest {
	op := &aws.Operation{
		Name:       opListServerCertificates,
		HTTPMethod: "POST",
		HTTPPath:   "/",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"Marker"},
			OutputTokens:    []string{"Marker"},
			LimitToken:      "MaxItems",
			TruncationToken: "IsTruncated",
		},
	}

	if input == nil {
		input = &ListServerCertificatesInput{}
	}

	req := c.newRequest(op, input, &ListServerCertificatesOutput{})

	return ListServerCertificatesRequest{Request: req, Input: input, Copy: c.ListServerCertificatesRequest}
}

// ListServerCertificatesRequest is the request type for the
// ListServerCertificates API operation.
type ListServerCertificatesRequest struct {
	*aws.Request
	Input *ListServerCertificatesInput
	Copy  func(*ListServerCertificatesInput) ListServerCertificatesRequest
}

// Send marshals and sends the ListServerCertificates API request.
func (r ListServerCertificatesRequest) Send(ctx context.Context) (*ListServerCertificatesResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListServerCertificatesResponse{
		ListServerCertificatesOutput: r.Request.Data.(*ListServerCertificatesOutput),
		response:                     &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewListServerCertificatesRequestPaginator returns a paginator for ListServerCertificates.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.ListServerCertificatesRequest(input)
//   p := iam.NewListServerCertificatesRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewListServerCertificatesPaginator(req ListServerCertificatesRequest) ListServerCertificatesPaginator {
	return ListServerCertificatesPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *ListServerCertificatesInput
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

// ListServerCertificatesPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type ListServerCertificatesPaginator struct {
	aws.Pager
}

func (p *ListServerCertificatesPaginator) CurrentPage() *ListServerCertificatesOutput {
	return p.Pager.CurrentPage().(*ListServerCertificatesOutput)
}

// ListServerCertificatesResponse is the response type for the
// ListServerCertificates API operation.
type ListServerCertificatesResponse struct {
	*ListServerCertificatesOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListServerCertificates request.
func (r *ListServerCertificatesResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
