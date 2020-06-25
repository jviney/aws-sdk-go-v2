// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package route53

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
	"github.com/jviney/aws-sdk-go-v2/private/protocol"
)

// A complex type that contains information about the tags that you want to
// add, edit, or delete.
type ChangeTagsForResourceInput struct {
	_ struct{} `locationName:"ChangeTagsForResourceRequest" type:"structure" xmlURI:"https://route53.amazonaws.com/doc/2013-04-01/"`

	// A complex type that contains a list of the tags that you want to add to the
	// specified health check or hosted zone and/or the tags that you want to edit
	// Value for.
	//
	// You can add a maximum of 10 tags to a health check or a hosted zone.
	AddTags []Tag `locationNameList:"Tag" min:"1" type:"list"`

	// A complex type that contains a list of the tags that you want to delete from
	// the specified health check or hosted zone. You can specify up to 10 keys.
	RemoveTagKeys []string `locationNameList:"Key" min:"1" type:"list"`

	// The ID of the resource for which you want to add, change, or delete tags.
	//
	// ResourceId is a required field
	ResourceId *string `location:"uri" locationName:"ResourceId" type:"string" required:"true"`

	// The type of the resource.
	//
	//    * The resource type for health checks is healthcheck.
	//
	//    * The resource type for hosted zones is hostedzone.
	//
	// ResourceType is a required field
	ResourceType TagResourceType `location:"uri" locationName:"ResourceType" type:"string" required:"true" enum:"true"`
}

// String returns the string representation
func (s ChangeTagsForResourceInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ChangeTagsForResourceInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ChangeTagsForResourceInput"}
	if s.AddTags != nil && len(s.AddTags) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("AddTags", 1))
	}
	if s.RemoveTagKeys != nil && len(s.RemoveTagKeys) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("RemoveTagKeys", 1))
	}

	if s.ResourceId == nil {
		invalidParams.Add(aws.NewErrParamRequired("ResourceId"))
	}
	if len(s.ResourceType) == 0 {
		invalidParams.Add(aws.NewErrParamRequired("ResourceType"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ChangeTagsForResourceInput) MarshalFields(e protocol.FieldEncoder) error {

	e.SetFields(protocol.BodyTarget, "ChangeTagsForResourceRequest", protocol.FieldMarshalerFunc(func(e protocol.FieldEncoder) error {
		if s.AddTags != nil {
			v := s.AddTags

			metadata := protocol.Metadata{ListLocationName: "Tag"}
			ls0 := e.List(protocol.BodyTarget, "AddTags", metadata)
			ls0.Start()
			for _, v1 := range v {
				ls0.ListAddFields(v1)
			}
			ls0.End()

		}
		if s.RemoveTagKeys != nil {
			v := s.RemoveTagKeys

			metadata := protocol.Metadata{ListLocationName: "Key"}
			ls0 := e.List(protocol.BodyTarget, "RemoveTagKeys", metadata)
			ls0.Start()
			for _, v1 := range v {
				ls0.ListAddValue(protocol.StringValue(v1))
			}
			ls0.End()

		}
		return nil
	}), protocol.Metadata{XMLNamespaceURI: "https://route53.amazonaws.com/doc/2013-04-01/"})
	if s.ResourceId != nil {
		v := *s.ResourceId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "ResourceId", protocol.StringValue(v), metadata)
	}
	if len(s.ResourceType) > 0 {
		v := s.ResourceType

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "ResourceType", v, metadata)
	}
	return nil
}

// Empty response for the request.
type ChangeTagsForResourceOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s ChangeTagsForResourceOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ChangeTagsForResourceOutput) MarshalFields(e protocol.FieldEncoder) error {
	return nil
}

const opChangeTagsForResource = "ChangeTagsForResource"

// ChangeTagsForResourceRequest returns a request value for making API operation for
// Amazon Route 53.
//
// Adds, edits, or deletes tags for a health check or a hosted zone.
//
// For information about using tags for cost allocation, see Using Cost Allocation
// Tags (https://docs.aws.amazon.com/awsaccountbilling/latest/aboutv2/cost-alloc-tags.html)
// in the AWS Billing and Cost Management User Guide.
//
//    // Example sending a request using ChangeTagsForResourceRequest.
//    req := client.ChangeTagsForResourceRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/route53-2013-04-01/ChangeTagsForResource
func (c *Client) ChangeTagsForResourceRequest(input *ChangeTagsForResourceInput) ChangeTagsForResourceRequest {
	op := &aws.Operation{
		Name:       opChangeTagsForResource,
		HTTPMethod: "POST",
		HTTPPath:   "/2013-04-01/tags/{ResourceType}/{ResourceId}",
	}

	if input == nil {
		input = &ChangeTagsForResourceInput{}
	}

	req := c.newRequest(op, input, &ChangeTagsForResourceOutput{})

	return ChangeTagsForResourceRequest{Request: req, Input: input, Copy: c.ChangeTagsForResourceRequest}
}

// ChangeTagsForResourceRequest is the request type for the
// ChangeTagsForResource API operation.
type ChangeTagsForResourceRequest struct {
	*aws.Request
	Input *ChangeTagsForResourceInput
	Copy  func(*ChangeTagsForResourceInput) ChangeTagsForResourceRequest
}

// Send marshals and sends the ChangeTagsForResource API request.
func (r ChangeTagsForResourceRequest) Send(ctx context.Context) (*ChangeTagsForResourceResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ChangeTagsForResourceResponse{
		ChangeTagsForResourceOutput: r.Request.Data.(*ChangeTagsForResourceOutput),
		response:                    &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// ChangeTagsForResourceResponse is the response type for the
// ChangeTagsForResource API operation.
type ChangeTagsForResourceResponse struct {
	*ChangeTagsForResourceOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ChangeTagsForResource request.
func (r *ChangeTagsForResourceResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
