// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package qldb

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
	"github.com/jviney/aws-sdk-go-v2/private/protocol"
)

type DescribeJournalS3ExportInput struct {
	_ struct{} `type:"structure"`

	// The unique ID of the journal export job that you want to describe.
	//
	// ExportId is a required field
	ExportId *string `location:"uri" locationName:"exportId" min:"22" type:"string" required:"true"`

	// The name of the ledger.
	//
	// Name is a required field
	Name *string `location:"uri" locationName:"name" min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s DescribeJournalS3ExportInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeJournalS3ExportInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeJournalS3ExportInput"}

	if s.ExportId == nil {
		invalidParams.Add(aws.NewErrParamRequired("ExportId"))
	}
	if s.ExportId != nil && len(*s.ExportId) < 22 {
		invalidParams.Add(aws.NewErrParamMinLen("ExportId", 22))
	}

	if s.Name == nil {
		invalidParams.Add(aws.NewErrParamRequired("Name"))
	}
	if s.Name != nil && len(*s.Name) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Name", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DescribeJournalS3ExportInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.ExportId != nil {
		v := *s.ExportId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "exportId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Name != nil {
		v := *s.Name

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "name", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type DescribeJournalS3ExportOutput struct {
	_ struct{} `type:"structure"`

	// Information about the journal export job returned by a DescribeJournalS3Export
	// request.
	//
	// ExportDescription is a required field
	ExportDescription *JournalS3ExportDescription `type:"structure" required:"true"`
}

// String returns the string representation
func (s DescribeJournalS3ExportOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DescribeJournalS3ExportOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.ExportDescription != nil {
		v := s.ExportDescription

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "ExportDescription", v, metadata)
	}
	return nil
}

const opDescribeJournalS3Export = "DescribeJournalS3Export"

// DescribeJournalS3ExportRequest returns a request value for making API operation for
// Amazon QLDB.
//
// Returns information about a journal export job, including the ledger name,
// export ID, when it was created, current status, and its start and end time
// export parameters.
//
// This action does not return any expired export jobs. For more information,
// see Export Job Expiration (https://docs.aws.amazon.com/qldb/latest/developerguide/export-journal.request.html#export-journal.request.expiration)
// in the Amazon QLDB Developer Guide.
//
// If the export job with the given ExportId doesn't exist, then throws ResourceNotFoundException.
//
// If the ledger with the given Name doesn't exist, then throws ResourceNotFoundException.
//
//    // Example sending a request using DescribeJournalS3ExportRequest.
//    req := client.DescribeJournalS3ExportRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/qldb-2019-01-02/DescribeJournalS3Export
func (c *Client) DescribeJournalS3ExportRequest(input *DescribeJournalS3ExportInput) DescribeJournalS3ExportRequest {
	op := &aws.Operation{
		Name:       opDescribeJournalS3Export,
		HTTPMethod: "GET",
		HTTPPath:   "/ledgers/{name}/journal-s3-exports/{exportId}",
	}

	if input == nil {
		input = &DescribeJournalS3ExportInput{}
	}

	req := c.newRequest(op, input, &DescribeJournalS3ExportOutput{})

	return DescribeJournalS3ExportRequest{Request: req, Input: input, Copy: c.DescribeJournalS3ExportRequest}
}

// DescribeJournalS3ExportRequest is the request type for the
// DescribeJournalS3Export API operation.
type DescribeJournalS3ExportRequest struct {
	*aws.Request
	Input *DescribeJournalS3ExportInput
	Copy  func(*DescribeJournalS3ExportInput) DescribeJournalS3ExportRequest
}

// Send marshals and sends the DescribeJournalS3Export API request.
func (r DescribeJournalS3ExportRequest) Send(ctx context.Context) (*DescribeJournalS3ExportResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeJournalS3ExportResponse{
		DescribeJournalS3ExportOutput: r.Request.Data.(*DescribeJournalS3ExportOutput),
		response:                      &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeJournalS3ExportResponse is the response type for the
// DescribeJournalS3Export API operation.
type DescribeJournalS3ExportResponse struct {
	*DescribeJournalS3ExportOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeJournalS3Export request.
func (r *DescribeJournalS3ExportResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
