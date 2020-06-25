// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package lexmodelbuildingservice

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
	"github.com/jviney/aws-sdk-go-v2/private/protocol"
)

type GetSlotTypeVersionsInput struct {
	_ struct{} `type:"structure"`

	// The maximum number of slot type versions to return in the response. The default
	// is 10.
	MaxResults *int64 `location:"querystring" locationName:"maxResults" min:"1" type:"integer"`

	// The name of the slot type for which versions should be returned.
	//
	// Name is a required field
	Name *string `location:"uri" locationName:"name" min:"1" type:"string" required:"true"`

	// A pagination token for fetching the next page of slot type versions. If the
	// response to this call is truncated, Amazon Lex returns a pagination token
	// in the response. To fetch the next page of versions, specify the pagination
	// token in the next request.
	NextToken *string `location:"querystring" locationName:"nextToken" type:"string"`
}

// String returns the string representation
func (s GetSlotTypeVersionsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetSlotTypeVersionsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetSlotTypeVersionsInput"}
	if s.MaxResults != nil && *s.MaxResults < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxResults", 1))
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
func (s GetSlotTypeVersionsInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.Name != nil {
		v := *s.Name

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "name", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.MaxResults != nil {
		v := *s.MaxResults

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "maxResults", protocol.Int64Value(v), metadata)
	}
	if s.NextToken != nil {
		v := *s.NextToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "nextToken", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type GetSlotTypeVersionsOutput struct {
	_ struct{} `type:"structure"`

	// A pagination token for fetching the next page of slot type versions. If the
	// response to this call is truncated, Amazon Lex returns a pagination token
	// in the response. To fetch the next page of versions, specify the pagination
	// token in the next request.
	NextToken *string `locationName:"nextToken" type:"string"`

	// An array of SlotTypeMetadata objects, one for each numbered version of the
	// slot type plus one for the $LATEST version.
	SlotTypes []SlotTypeMetadata `locationName:"slotTypes" type:"list"`
}

// String returns the string representation
func (s GetSlotTypeVersionsOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetSlotTypeVersionsOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.NextToken != nil {
		v := *s.NextToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "nextToken", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.SlotTypes != nil {
		v := s.SlotTypes

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "slotTypes", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	return nil
}

const opGetSlotTypeVersions = "GetSlotTypeVersions"

// GetSlotTypeVersionsRequest returns a request value for making API operation for
// Amazon Lex Model Building Service.
//
// Gets information about all versions of a slot type.
//
// The GetSlotTypeVersions operation returns a SlotTypeMetadata object for each
// version of a slot type. For example, if a slot type has three numbered versions,
// the GetSlotTypeVersions operation returns four SlotTypeMetadata objects in
// the response, one for each numbered version and one for the $LATEST version.
//
// The GetSlotTypeVersions operation always returns at least one version, the
// $LATEST version.
//
// This operation requires permissions for the lex:GetSlotTypeVersions action.
//
//    // Example sending a request using GetSlotTypeVersionsRequest.
//    req := client.GetSlotTypeVersionsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/lex-models-2017-04-19/GetSlotTypeVersions
func (c *Client) GetSlotTypeVersionsRequest(input *GetSlotTypeVersionsInput) GetSlotTypeVersionsRequest {
	op := &aws.Operation{
		Name:       opGetSlotTypeVersions,
		HTTPMethod: "GET",
		HTTPPath:   "/slottypes/{name}/versions/",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"nextToken"},
			OutputTokens:    []string{"nextToken"},
			LimitToken:      "maxResults",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &GetSlotTypeVersionsInput{}
	}

	req := c.newRequest(op, input, &GetSlotTypeVersionsOutput{})

	return GetSlotTypeVersionsRequest{Request: req, Input: input, Copy: c.GetSlotTypeVersionsRequest}
}

// GetSlotTypeVersionsRequest is the request type for the
// GetSlotTypeVersions API operation.
type GetSlotTypeVersionsRequest struct {
	*aws.Request
	Input *GetSlotTypeVersionsInput
	Copy  func(*GetSlotTypeVersionsInput) GetSlotTypeVersionsRequest
}

// Send marshals and sends the GetSlotTypeVersions API request.
func (r GetSlotTypeVersionsRequest) Send(ctx context.Context) (*GetSlotTypeVersionsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetSlotTypeVersionsResponse{
		GetSlotTypeVersionsOutput: r.Request.Data.(*GetSlotTypeVersionsOutput),
		response:                  &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewGetSlotTypeVersionsRequestPaginator returns a paginator for GetSlotTypeVersions.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.GetSlotTypeVersionsRequest(input)
//   p := lexmodelbuildingservice.NewGetSlotTypeVersionsRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewGetSlotTypeVersionsPaginator(req GetSlotTypeVersionsRequest) GetSlotTypeVersionsPaginator {
	return GetSlotTypeVersionsPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *GetSlotTypeVersionsInput
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

// GetSlotTypeVersionsPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type GetSlotTypeVersionsPaginator struct {
	aws.Pager
}

func (p *GetSlotTypeVersionsPaginator) CurrentPage() *GetSlotTypeVersionsOutput {
	return p.Pager.CurrentPage().(*GetSlotTypeVersionsOutput)
}

// GetSlotTypeVersionsResponse is the response type for the
// GetSlotTypeVersions API operation.
type GetSlotTypeVersionsResponse struct {
	*GetSlotTypeVersionsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetSlotTypeVersions request.
func (r *GetSlotTypeVersionsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
