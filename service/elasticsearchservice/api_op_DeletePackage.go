// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package elasticsearchservice

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
	"github.com/jviney/aws-sdk-go-v2/private/protocol"
)

// Container for request parameters to DeletePackage operation.
type DeletePackageInput struct {
	_ struct{} `type:"structure"`

	// Internal ID of the package that you want to delete. Use DescribePackages
	// to find this value.
	//
	// PackageID is a required field
	PackageID *string `location:"uri" locationName:"PackageID" type:"string" required:"true"`
}

// String returns the string representation
func (s DeletePackageInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeletePackageInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeletePackageInput"}

	if s.PackageID == nil {
		invalidParams.Add(aws.NewErrParamRequired("PackageID"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DeletePackageInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.PackageID != nil {
		v := *s.PackageID

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "PackageID", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// Container for response parameters to DeletePackage operation.
type DeletePackageOutput struct {
	_ struct{} `type:"structure"`

	// PackageDetails
	PackageDetails *PackageDetails `type:"structure"`
}

// String returns the string representation
func (s DeletePackageOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DeletePackageOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.PackageDetails != nil {
		v := s.PackageDetails

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "PackageDetails", v, metadata)
	}
	return nil
}

const opDeletePackage = "DeletePackage"

// DeletePackageRequest returns a request value for making API operation for
// Amazon Elasticsearch Service.
//
// Delete the package.
//
//    // Example sending a request using DeletePackageRequest.
//    req := client.DeletePackageRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
func (c *Client) DeletePackageRequest(input *DeletePackageInput) DeletePackageRequest {
	op := &aws.Operation{
		Name:       opDeletePackage,
		HTTPMethod: "DELETE",
		HTTPPath:   "/2015-01-01/packages/{PackageID}",
	}

	if input == nil {
		input = &DeletePackageInput{}
	}

	req := c.newRequest(op, input, &DeletePackageOutput{})

	return DeletePackageRequest{Request: req, Input: input, Copy: c.DeletePackageRequest}
}

// DeletePackageRequest is the request type for the
// DeletePackage API operation.
type DeletePackageRequest struct {
	*aws.Request
	Input *DeletePackageInput
	Copy  func(*DeletePackageInput) DeletePackageRequest
}

// Send marshals and sends the DeletePackage API request.
func (r DeletePackageRequest) Send(ctx context.Context) (*DeletePackageResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeletePackageResponse{
		DeletePackageOutput: r.Request.Data.(*DeletePackageOutput),
		response:            &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeletePackageResponse is the response type for the
// DeletePackage API operation.
type DeletePackageResponse struct {
	*DeletePackageOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeletePackage request.
func (r *DeletePackageResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
