// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package codedeploy

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
)

// Represents the input of a GetApplicationRevision operation.
type GetApplicationRevisionInput struct {
	_ struct{} `type:"structure"`

	// The name of the application that corresponds to the revision.
	//
	// ApplicationName is a required field
	ApplicationName *string `locationName:"applicationName" min:"1" type:"string" required:"true"`

	// Information about the application revision to get, including type and location.
	//
	// Revision is a required field
	Revision *RevisionLocation `locationName:"revision" type:"structure" required:"true"`
}

// String returns the string representation
func (s GetApplicationRevisionInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetApplicationRevisionInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetApplicationRevisionInput"}

	if s.ApplicationName == nil {
		invalidParams.Add(aws.NewErrParamRequired("ApplicationName"))
	}
	if s.ApplicationName != nil && len(*s.ApplicationName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ApplicationName", 1))
	}

	if s.Revision == nil {
		invalidParams.Add(aws.NewErrParamRequired("Revision"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Represents the output of a GetApplicationRevision operation.
type GetApplicationRevisionOutput struct {
	_ struct{} `type:"structure"`

	// The name of the application that corresponds to the revision.
	ApplicationName *string `locationName:"applicationName" min:"1" type:"string"`

	// Additional information about the revision, including type and location.
	Revision *RevisionLocation `locationName:"revision" type:"structure"`

	// General information about the revision.
	RevisionInfo *GenericRevisionInfo `locationName:"revisionInfo" type:"structure"`
}

// String returns the string representation
func (s GetApplicationRevisionOutput) String() string {
	return awsutil.Prettify(s)
}

const opGetApplicationRevision = "GetApplicationRevision"

// GetApplicationRevisionRequest returns a request value for making API operation for
// AWS CodeDeploy.
//
// Gets information about an application revision.
//
//    // Example sending a request using GetApplicationRevisionRequest.
//    req := client.GetApplicationRevisionRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/codedeploy-2014-10-06/GetApplicationRevision
func (c *Client) GetApplicationRevisionRequest(input *GetApplicationRevisionInput) GetApplicationRevisionRequest {
	op := &aws.Operation{
		Name:       opGetApplicationRevision,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &GetApplicationRevisionInput{}
	}

	req := c.newRequest(op, input, &GetApplicationRevisionOutput{})

	return GetApplicationRevisionRequest{Request: req, Input: input, Copy: c.GetApplicationRevisionRequest}
}

// GetApplicationRevisionRequest is the request type for the
// GetApplicationRevision API operation.
type GetApplicationRevisionRequest struct {
	*aws.Request
	Input *GetApplicationRevisionInput
	Copy  func(*GetApplicationRevisionInput) GetApplicationRevisionRequest
}

// Send marshals and sends the GetApplicationRevision API request.
func (r GetApplicationRevisionRequest) Send(ctx context.Context) (*GetApplicationRevisionResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetApplicationRevisionResponse{
		GetApplicationRevisionOutput: r.Request.Data.(*GetApplicationRevisionOutput),
		response:                     &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetApplicationRevisionResponse is the response type for the
// GetApplicationRevision API operation.
type GetApplicationRevisionResponse struct {
	*GetApplicationRevisionOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetApplicationRevision request.
func (r *GetApplicationRevisionResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
