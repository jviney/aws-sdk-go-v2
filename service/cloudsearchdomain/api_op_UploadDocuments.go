// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package cloudsearchdomain

import (
	"context"
	"io"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
	"github.com/jviney/aws-sdk-go-v2/private/protocol"
)

// Container for the parameters to the UploadDocuments request.
type UploadDocumentsInput struct {
	_ struct{} `type:"structure" payload:"Documents"`

	// The format of the batch you are uploading. Amazon CloudSearch supports two
	// document batch formats:
	//
	//    * application/json
	//
	//    * application/xml
	//
	// ContentType is a required field
	ContentType ContentType `location:"header" locationName:"Content-Type" type:"string" required:"true" enum:"true"`

	// A batch of documents formatted in JSON or HTML.
	//
	// Documents is a required field
	Documents io.ReadSeeker `locationName:"documents" type:"blob" required:"true"`
}

// String returns the string representation
func (s UploadDocumentsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UploadDocumentsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "UploadDocumentsInput"}
	if len(s.ContentType) == 0 {
		invalidParams.Add(aws.NewErrParamRequired("ContentType"))
	}

	if s.Documents == nil {
		invalidParams.Add(aws.NewErrParamRequired("Documents"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s UploadDocumentsInput) MarshalFields(e protocol.FieldEncoder) error {

	if len(s.ContentType) > 0 {
		v := s.ContentType

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.QuotedValue{ValueMarshaler: v}, metadata)
	}
	if s.Documents != nil {
		v := s.Documents

		metadata := protocol.Metadata{}
		e.SetStream(protocol.PayloadTarget, "documents", protocol.ReadSeekerStream{V: v}, metadata)
	}
	return nil
}

// Contains the response to an UploadDocuments request.
type UploadDocumentsOutput struct {
	_ struct{} `type:"structure"`

	// The number of documents that were added to the search domain.
	Adds *int64 `locationName:"adds" type:"long"`

	// The number of documents that were deleted from the search domain.
	Deletes *int64 `locationName:"deletes" type:"long"`

	// The status of an UploadDocumentsRequest.
	Status *string `locationName:"status" type:"string"`

	// Any warnings returned by the document service about the documents being uploaded.
	Warnings []DocumentServiceWarning `locationName:"warnings" type:"list"`
}

// String returns the string representation
func (s UploadDocumentsOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s UploadDocumentsOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.Adds != nil {
		v := *s.Adds

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "adds", protocol.Int64Value(v), metadata)
	}
	if s.Deletes != nil {
		v := *s.Deletes

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "deletes", protocol.Int64Value(v), metadata)
	}
	if s.Status != nil {
		v := *s.Status

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "status", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Warnings != nil {
		v := s.Warnings

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "warnings", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	return nil
}

const opUploadDocuments = "UploadDocuments"

// UploadDocumentsRequest returns a request value for making API operation for
// Amazon CloudSearch Domain.
//
// Posts a batch of documents to a search domain for indexing. A document batch
// is a collection of add and delete operations that represent the documents
// you want to add, update, or delete from your domain. Batches can be described
// in either JSON or XML. Each item that you want Amazon CloudSearch to return
// as a search result (such as a product) is represented as a document. Every
// document has a unique ID and one or more fields that contain the data that
// you want to search and return in results. Individual documents cannot contain
// more than 1 MB of data. The entire batch cannot exceed 5 MB. To get the best
// possible upload performance, group add and delete operations in batches that
// are close the 5 MB limit. Submitting a large volume of single-document batches
// can overload a domain's document service.
//
// The endpoint for submitting UploadDocuments requests is domain-specific.
// To get the document endpoint for your domain, use the Amazon CloudSearch
// configuration service DescribeDomains action. A domain's endpoints are also
// displayed on the domain dashboard in the Amazon CloudSearch console.
//
// For more information about formatting your data for Amazon CloudSearch, see
// Preparing Your Data (http://docs.aws.amazon.com/cloudsearch/latest/developerguide/preparing-data.html)
// in the Amazon CloudSearch Developer Guide. For more information about uploading
// data for indexing, see Uploading Data (http://docs.aws.amazon.com/cloudsearch/latest/developerguide/uploading-data.html)
// in the Amazon CloudSearch Developer Guide.
//
//    // Example sending a request using UploadDocumentsRequest.
//    req := client.UploadDocumentsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
func (c *Client) UploadDocumentsRequest(input *UploadDocumentsInput) UploadDocumentsRequest {
	op := &aws.Operation{
		Name:       opUploadDocuments,
		HTTPMethod: "POST",
		HTTPPath:   "/2013-01-01/documents/batch?format=sdk",
	}

	if input == nil {
		input = &UploadDocumentsInput{}
	}

	req := c.newRequest(op, input, &UploadDocumentsOutput{})

	return UploadDocumentsRequest{Request: req, Input: input, Copy: c.UploadDocumentsRequest}
}

// UploadDocumentsRequest is the request type for the
// UploadDocuments API operation.
type UploadDocumentsRequest struct {
	*aws.Request
	Input *UploadDocumentsInput
	Copy  func(*UploadDocumentsInput) UploadDocumentsRequest
}

// Send marshals and sends the UploadDocuments API request.
func (r UploadDocumentsRequest) Send(ctx context.Context) (*UploadDocumentsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &UploadDocumentsResponse{
		UploadDocumentsOutput: r.Request.Data.(*UploadDocumentsOutput),
		response:              &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// UploadDocumentsResponse is the response type for the
// UploadDocuments API operation.
type UploadDocumentsResponse struct {
	*UploadDocumentsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// UploadDocuments request.
func (r *UploadDocumentsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
