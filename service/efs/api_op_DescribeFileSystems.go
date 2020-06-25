// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package efs

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
	"github.com/jviney/aws-sdk-go-v2/private/protocol"
)

type DescribeFileSystemsInput struct {
	_ struct{} `type:"structure"`

	// (Optional) Restricts the list to the file system with this creation token
	// (String). You specify a creation token when you create an Amazon EFS file
	// system.
	CreationToken *string `location:"querystring" locationName:"CreationToken" min:"1" type:"string"`

	// (Optional) ID of the file system whose description you want to retrieve (String).
	FileSystemId *string `location:"querystring" locationName:"FileSystemId" type:"string"`

	// (Optional) Opaque pagination token returned from a previous DescribeFileSystems
	// operation (String). If present, specifies to continue the list from where
	// the returning call had left off.
	Marker *string `location:"querystring" locationName:"Marker" type:"string"`

	// (Optional) Specifies the maximum number of file systems to return in the
	// response (integer). This number is automatically set to 100. The response
	// is paginated at 100 per page if you have more than 100 file systems.
	MaxItems *int64 `location:"querystring" locationName:"MaxItems" min:"1" type:"integer"`
}

// String returns the string representation
func (s DescribeFileSystemsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeFileSystemsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeFileSystemsInput"}
	if s.CreationToken != nil && len(*s.CreationToken) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("CreationToken", 1))
	}
	if s.MaxItems != nil && *s.MaxItems < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxItems", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DescribeFileSystemsInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.CreationToken != nil {
		v := *s.CreationToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "CreationToken", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.FileSystemId != nil {
		v := *s.FileSystemId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "FileSystemId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Marker != nil {
		v := *s.Marker

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "Marker", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.MaxItems != nil {
		v := *s.MaxItems

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "MaxItems", protocol.Int64Value(v), metadata)
	}
	return nil
}

type DescribeFileSystemsOutput struct {
	_ struct{} `type:"structure"`

	// An array of file system descriptions.
	FileSystems []FileSystemDescription `type:"list"`

	// Present if provided by caller in the request (String).
	Marker *string `type:"string"`

	// Present if there are more file systems than returned in the response (String).
	// You can use the NextMarker in the subsequent request to fetch the descriptions.
	NextMarker *string `type:"string"`
}

// String returns the string representation
func (s DescribeFileSystemsOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DescribeFileSystemsOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.FileSystems != nil {
		v := s.FileSystems

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "FileSystems", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	if s.Marker != nil {
		v := *s.Marker

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "Marker", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.NextMarker != nil {
		v := *s.NextMarker

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "NextMarker", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

const opDescribeFileSystems = "DescribeFileSystems"

// DescribeFileSystemsRequest returns a request value for making API operation for
// Amazon Elastic File System.
//
// Returns the description of a specific Amazon EFS file system if either the
// file system CreationToken or the FileSystemId is provided. Otherwise, it
// returns descriptions of all file systems owned by the caller's AWS account
// in the AWS Region of the endpoint that you're calling.
//
// When retrieving all file system descriptions, you can optionally specify
// the MaxItems parameter to limit the number of descriptions in a response.
// Currently, this number is automatically set to 10. If more file system descriptions
// remain, Amazon EFS returns a NextMarker, an opaque token, in the response.
// In this case, you should send a subsequent request with the Marker request
// parameter set to the value of NextMarker.
//
// To retrieve a list of your file system descriptions, this operation is used
// in an iterative process, where DescribeFileSystems is called first without
// the Marker and then the operation continues to call it with the Marker parameter
// set to the value of the NextMarker from the previous response until the response
// has no NextMarker.
//
// The order of file systems returned in the response of one DescribeFileSystems
// call and the order of file systems returned across the responses of a multi-call
// iteration is unspecified.
//
// This operation requires permissions for the elasticfilesystem:DescribeFileSystems
// action.
//
//    // Example sending a request using DescribeFileSystemsRequest.
//    req := client.DescribeFileSystemsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/elasticfilesystem-2015-02-01/DescribeFileSystems
func (c *Client) DescribeFileSystemsRequest(input *DescribeFileSystemsInput) DescribeFileSystemsRequest {
	op := &aws.Operation{
		Name:       opDescribeFileSystems,
		HTTPMethod: "GET",
		HTTPPath:   "/2015-02-01/file-systems",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"Marker"},
			OutputTokens:    []string{"NextMarker"},
			LimitToken:      "MaxItems",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &DescribeFileSystemsInput{}
	}

	req := c.newRequest(op, input, &DescribeFileSystemsOutput{})

	return DescribeFileSystemsRequest{Request: req, Input: input, Copy: c.DescribeFileSystemsRequest}
}

// DescribeFileSystemsRequest is the request type for the
// DescribeFileSystems API operation.
type DescribeFileSystemsRequest struct {
	*aws.Request
	Input *DescribeFileSystemsInput
	Copy  func(*DescribeFileSystemsInput) DescribeFileSystemsRequest
}

// Send marshals and sends the DescribeFileSystems API request.
func (r DescribeFileSystemsRequest) Send(ctx context.Context) (*DescribeFileSystemsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeFileSystemsResponse{
		DescribeFileSystemsOutput: r.Request.Data.(*DescribeFileSystemsOutput),
		response:                  &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewDescribeFileSystemsRequestPaginator returns a paginator for DescribeFileSystems.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.DescribeFileSystemsRequest(input)
//   p := efs.NewDescribeFileSystemsRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewDescribeFileSystemsPaginator(req DescribeFileSystemsRequest) DescribeFileSystemsPaginator {
	return DescribeFileSystemsPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *DescribeFileSystemsInput
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

// DescribeFileSystemsPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type DescribeFileSystemsPaginator struct {
	aws.Pager
}

func (p *DescribeFileSystemsPaginator) CurrentPage() *DescribeFileSystemsOutput {
	return p.Pager.CurrentPage().(*DescribeFileSystemsOutput)
}

// DescribeFileSystemsResponse is the response type for the
// DescribeFileSystems API operation.
type DescribeFileSystemsResponse struct {
	*DescribeFileSystemsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeFileSystems request.
func (r *DescribeFileSystemsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
