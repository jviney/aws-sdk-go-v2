// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package guardduty

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
	"github.com/jviney/aws-sdk-go-v2/private/protocol"
)

type ArchiveFindingsInput struct {
	_ struct{} `type:"structure"`

	// The ID of the detector that specifies the GuardDuty service whose findings
	// you want to archive.
	//
	// DetectorId is a required field
	DetectorId *string `location:"uri" locationName:"detectorId" min:"1" type:"string" required:"true"`

	// The IDs of the findings that you want to archive.
	//
	// FindingIds is a required field
	FindingIds []string `locationName:"findingIds" type:"list" required:"true"`
}

// String returns the string representation
func (s ArchiveFindingsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ArchiveFindingsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ArchiveFindingsInput"}

	if s.DetectorId == nil {
		invalidParams.Add(aws.NewErrParamRequired("DetectorId"))
	}
	if s.DetectorId != nil && len(*s.DetectorId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("DetectorId", 1))
	}

	if s.FindingIds == nil {
		invalidParams.Add(aws.NewErrParamRequired("FindingIds"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ArchiveFindingsInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.FindingIds != nil {
		v := s.FindingIds

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "findingIds", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddValue(protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v1)})
		}
		ls0.End()

	}
	if s.DetectorId != nil {
		v := *s.DetectorId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "detectorId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type ArchiveFindingsOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s ArchiveFindingsOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ArchiveFindingsOutput) MarshalFields(e protocol.FieldEncoder) error {
	return nil
}

const opArchiveFindings = "ArchiveFindings"

// ArchiveFindingsRequest returns a request value for making API operation for
// Amazon GuardDuty.
//
// Archives GuardDuty findings that are specified by the list of finding IDs.
//
// Only the master account can archive findings. Member accounts don't have
// permission to archive findings from their accounts.
//
//    // Example sending a request using ArchiveFindingsRequest.
//    req := client.ArchiveFindingsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/guardduty-2017-11-28/ArchiveFindings
func (c *Client) ArchiveFindingsRequest(input *ArchiveFindingsInput) ArchiveFindingsRequest {
	op := &aws.Operation{
		Name:       opArchiveFindings,
		HTTPMethod: "POST",
		HTTPPath:   "/detector/{detectorId}/findings/archive",
	}

	if input == nil {
		input = &ArchiveFindingsInput{}
	}

	req := c.newRequest(op, input, &ArchiveFindingsOutput{})

	return ArchiveFindingsRequest{Request: req, Input: input, Copy: c.ArchiveFindingsRequest}
}

// ArchiveFindingsRequest is the request type for the
// ArchiveFindings API operation.
type ArchiveFindingsRequest struct {
	*aws.Request
	Input *ArchiveFindingsInput
	Copy  func(*ArchiveFindingsInput) ArchiveFindingsRequest
}

// Send marshals and sends the ArchiveFindings API request.
func (r ArchiveFindingsRequest) Send(ctx context.Context) (*ArchiveFindingsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ArchiveFindingsResponse{
		ArchiveFindingsOutput: r.Request.Data.(*ArchiveFindingsOutput),
		response:              &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// ArchiveFindingsResponse is the response type for the
// ArchiveFindings API operation.
type ArchiveFindingsResponse struct {
	*ArchiveFindingsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ArchiveFindings request.
func (r *ArchiveFindingsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
