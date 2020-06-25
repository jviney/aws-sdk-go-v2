// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package route53

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
	"github.com/jviney/aws-sdk-go-v2/private/protocol"
)

// A complex type that contains information about the traffic policy that you
// want to create a new version for.
type CreateTrafficPolicyVersionInput struct {
	_ struct{} `locationName:"CreateTrafficPolicyVersionRequest" type:"structure" xmlURI:"https://route53.amazonaws.com/doc/2013-04-01/"`

	// The comment that you specified in the CreateTrafficPolicyVersion request,
	// if any.
	Comment *string `type:"string"`

	// The definition of this version of the traffic policy, in JSON format. You
	// specified the JSON in the CreateTrafficPolicyVersion request. For more information
	// about the JSON format, see CreateTrafficPolicy (https://docs.aws.amazon.com/Route53/latest/APIReference/API_CreateTrafficPolicy.html).
	//
	// Document is a required field
	Document *string `type:"string" required:"true"`

	// The ID of the traffic policy for which you want to create a new version.
	//
	// Id is a required field
	Id *string `location:"uri" locationName:"Id" min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s CreateTrafficPolicyVersionInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateTrafficPolicyVersionInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateTrafficPolicyVersionInput"}

	if s.Document == nil {
		invalidParams.Add(aws.NewErrParamRequired("Document"))
	}

	if s.Id == nil {
		invalidParams.Add(aws.NewErrParamRequired("Id"))
	}
	if s.Id != nil && len(*s.Id) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Id", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s CreateTrafficPolicyVersionInput) MarshalFields(e protocol.FieldEncoder) error {

	e.SetFields(protocol.BodyTarget, "CreateTrafficPolicyVersionRequest", protocol.FieldMarshalerFunc(func(e protocol.FieldEncoder) error {
		if s.Comment != nil {
			v := *s.Comment

			metadata := protocol.Metadata{}
			e.SetValue(protocol.BodyTarget, "Comment", protocol.StringValue(v), metadata)
		}
		if s.Document != nil {
			v := *s.Document

			metadata := protocol.Metadata{}
			e.SetValue(protocol.BodyTarget, "Document", protocol.StringValue(v), metadata)
		}
		return nil
	}), protocol.Metadata{XMLNamespaceURI: "https://route53.amazonaws.com/doc/2013-04-01/"})
	if s.Id != nil {
		v := *s.Id

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "Id", protocol.StringValue(v), metadata)
	}
	return nil
}

// A complex type that contains the response information for the CreateTrafficPolicyVersion
// request.
type CreateTrafficPolicyVersionOutput struct {
	_ struct{} `type:"structure"`

	// A unique URL that represents a new traffic policy version.
	//
	// Location is a required field
	Location *string `location:"header" locationName:"Location" type:"string" required:"true"`

	// A complex type that contains settings for the new version of the traffic
	// policy.
	//
	// TrafficPolicy is a required field
	TrafficPolicy *TrafficPolicy `type:"structure" required:"true"`
}

// String returns the string representation
func (s CreateTrafficPolicyVersionOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s CreateTrafficPolicyVersionOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.TrafficPolicy != nil {
		v := s.TrafficPolicy

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "TrafficPolicy", v, metadata)
	}
	if s.Location != nil {
		v := *s.Location

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "Location", protocol.StringValue(v), metadata)
	}
	return nil
}

const opCreateTrafficPolicyVersion = "CreateTrafficPolicyVersion"

// CreateTrafficPolicyVersionRequest returns a request value for making API operation for
// Amazon Route 53.
//
// Creates a new version of an existing traffic policy. When you create a new
// version of a traffic policy, you specify the ID of the traffic policy that
// you want to update and a JSON-formatted document that describes the new version.
// You use traffic policies to create multiple DNS resource record sets for
// one domain name (such as example.com) or one subdomain name (such as www.example.com).
// You can create a maximum of 1000 versions of a traffic policy. If you reach
// the limit and need to create another version, you'll need to start a new
// traffic policy.
//
//    // Example sending a request using CreateTrafficPolicyVersionRequest.
//    req := client.CreateTrafficPolicyVersionRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/route53-2013-04-01/CreateTrafficPolicyVersion
func (c *Client) CreateTrafficPolicyVersionRequest(input *CreateTrafficPolicyVersionInput) CreateTrafficPolicyVersionRequest {
	op := &aws.Operation{
		Name:       opCreateTrafficPolicyVersion,
		HTTPMethod: "POST",
		HTTPPath:   "/2013-04-01/trafficpolicy/{Id}",
	}

	if input == nil {
		input = &CreateTrafficPolicyVersionInput{}
	}

	req := c.newRequest(op, input, &CreateTrafficPolicyVersionOutput{})

	return CreateTrafficPolicyVersionRequest{Request: req, Input: input, Copy: c.CreateTrafficPolicyVersionRequest}
}

// CreateTrafficPolicyVersionRequest is the request type for the
// CreateTrafficPolicyVersion API operation.
type CreateTrafficPolicyVersionRequest struct {
	*aws.Request
	Input *CreateTrafficPolicyVersionInput
	Copy  func(*CreateTrafficPolicyVersionInput) CreateTrafficPolicyVersionRequest
}

// Send marshals and sends the CreateTrafficPolicyVersion API request.
func (r CreateTrafficPolicyVersionRequest) Send(ctx context.Context) (*CreateTrafficPolicyVersionResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateTrafficPolicyVersionResponse{
		CreateTrafficPolicyVersionOutput: r.Request.Data.(*CreateTrafficPolicyVersionOutput),
		response:                         &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateTrafficPolicyVersionResponse is the response type for the
// CreateTrafficPolicyVersion API operation.
type CreateTrafficPolicyVersionResponse struct {
	*CreateTrafficPolicyVersionOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateTrafficPolicyVersion request.
func (r *CreateTrafficPolicyVersionResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
