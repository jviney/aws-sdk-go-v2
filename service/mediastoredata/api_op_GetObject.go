// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package mediastoredata

import (
	"context"
	"io"
	"time"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
	"github.com/jviney/aws-sdk-go-v2/private/protocol"
)

type GetObjectInput struct {
	_ struct{} `type:"structure"`

	// The path (including the file name) where the object is stored in the container.
	// Format: <folder name>/<folder name>/<file name>
	//
	// For example, to upload the file mlaw.avi to the folder path premium\canada
	// in the container movies, enter the path premium/canada/mlaw.avi.
	//
	// Do not include the container name in this path.
	//
	// If the path includes any folders that don't exist yet, the service creates
	// them. For example, suppose you have an existing premium/usa subfolder. If
	// you specify premium/canada, the service creates a canada subfolder in the
	// premium folder. You then have two subfolders, usa and canada, in the premium
	// folder.
	//
	// There is no correlation between the path to the source and the path (folders)
	// in the container in AWS Elemental MediaStore.
	//
	// For more information about folders and how they exist in a container, see
	// the AWS Elemental MediaStore User Guide (http://docs.aws.amazon.com/mediastore/latest/ug/).
	//
	// The file name is the name that is assigned to the file that you upload. The
	// file can have the same name inside and outside of AWS Elemental MediaStore,
	// or it can have the same name. The file name can include or omit an extension.
	//
	// Path is a required field
	Path *string `location:"uri" locationName:"Path" min:"1" type:"string" required:"true"`

	// The range bytes of an object to retrieve. For more information about the
	// Range header, see http://www.w3.org/Protocols/rfc2616/rfc2616-sec14.html#sec14.35
	// (http://www.w3.org/Protocols/rfc2616/rfc2616-sec14.html#sec14.35). AWS Elemental
	// MediaStore ignores this header for partially uploaded objects that have streaming
	// upload availability.
	Range *string `location:"header" locationName:"Range" type:"string"`
}

// String returns the string representation
func (s GetObjectInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetObjectInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetObjectInput"}

	if s.Path == nil {
		invalidParams.Add(aws.NewErrParamRequired("Path"))
	}
	if s.Path != nil && len(*s.Path) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Path", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetObjectInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.Range != nil {
		v := *s.Range

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "Range", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Path != nil {
		v := *s.Path

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "Path", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type GetObjectOutput struct {
	_ struct{} `type:"structure" payload:"Body"`

	// The bytes of the object.
	Body io.ReadCloser `type:"blob"`

	// An optional CacheControl header that allows the caller to control the object's
	// cache behavior. Headers can be passed in as specified in the HTTP spec at
	// https://www.w3.org/Protocols/rfc2616/rfc2616-sec14.html#sec14.9 (https://www.w3.org/Protocols/rfc2616/rfc2616-sec14.html#sec14.9).
	//
	// Headers with a custom user-defined value are also accepted.
	CacheControl *string `location:"header" locationName:"Cache-Control" type:"string"`

	// The length of the object in bytes.
	ContentLength *int64 `location:"header" locationName:"Content-Length" type:"long"`

	// The range of bytes to retrieve.
	ContentRange *string `location:"header" locationName:"Content-Range" type:"string"`

	// The content type of the object.
	ContentType *string `location:"header" locationName:"Content-Type" type:"string"`

	// The ETag that represents a unique instance of the object.
	ETag *string `location:"header" locationName:"ETag" min:"1" type:"string"`

	// The date and time that the object was last modified.
	LastModified *time.Time `location:"header" locationName:"Last-Modified" type:"timestamp"`

	// The HTML status code of the request. Status codes ranging from 200 to 299
	// indicate success. All other status codes indicate the type of error that
	// occurred.
	//
	// StatusCode is a required field
	StatusCode *int64 `location:"statusCode" type:"integer" required:"true"`
}

// String returns the string representation
func (s GetObjectOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetObjectOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.CacheControl != nil {
		v := *s.CacheControl

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "Cache-Control", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.ContentLength != nil {
		v := *s.ContentLength

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "Content-Length", protocol.Int64Value(v), metadata)
	}
	if s.ContentRange != nil {
		v := *s.ContentRange

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "Content-Range", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.ContentType != nil {
		v := *s.ContentType

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.ETag != nil {
		v := *s.ETag

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "ETag", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.LastModified != nil {
		v := *s.LastModified

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "Last-Modified",
			protocol.TimeValue{V: v, Format: protocol.RFC822TimeFormatName, QuotedFormatTime: false}, metadata)
	}
	// Skipping Body Output type's body not valid.
	// ignoring invalid encode state, StatusCode. StatusCode
	return nil
}

const opGetObject = "GetObject"

// GetObjectRequest returns a request value for making API operation for
// AWS Elemental MediaStore Data Plane.
//
// Downloads the object at the specified path. If the object’s upload availability
// is set to streaming, AWS Elemental MediaStore downloads the object even if
// it’s still uploading the object.
//
//    // Example sending a request using GetObjectRequest.
//    req := client.GetObjectRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/mediastore-data-2017-09-01/GetObject
func (c *Client) GetObjectRequest(input *GetObjectInput) GetObjectRequest {
	op := &aws.Operation{
		Name:       opGetObject,
		HTTPMethod: "GET",
		HTTPPath:   "/{Path+}",
	}

	if input == nil {
		input = &GetObjectInput{}
	}

	req := c.newRequest(op, input, &GetObjectOutput{})

	return GetObjectRequest{Request: req, Input: input, Copy: c.GetObjectRequest}
}

// GetObjectRequest is the request type for the
// GetObject API operation.
type GetObjectRequest struct {
	*aws.Request
	Input *GetObjectInput
	Copy  func(*GetObjectInput) GetObjectRequest
}

// Send marshals and sends the GetObject API request.
func (r GetObjectRequest) Send(ctx context.Context) (*GetObjectResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetObjectResponse{
		GetObjectOutput: r.Request.Data.(*GetObjectOutput),
		response:        &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetObjectResponse is the response type for the
// GetObject API operation.
type GetObjectResponse struct {
	*GetObjectOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetObject request.
func (r *GetObjectResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
