// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package directoryservice

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
)

type UpdateNumberOfDomainControllersInput struct {
	_ struct{} `type:"structure"`

	// The number of domain controllers desired in the directory.
	//
	// DesiredNumber is a required field
	DesiredNumber *int64 `min:"2" type:"integer" required:"true"`

	// Identifier of the directory to which the domain controllers will be added
	// or removed.
	//
	// DirectoryId is a required field
	DirectoryId *string `type:"string" required:"true"`
}

// String returns the string representation
func (s UpdateNumberOfDomainControllersInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UpdateNumberOfDomainControllersInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "UpdateNumberOfDomainControllersInput"}

	if s.DesiredNumber == nil {
		invalidParams.Add(aws.NewErrParamRequired("DesiredNumber"))
	}
	if s.DesiredNumber != nil && *s.DesiredNumber < 2 {
		invalidParams.Add(aws.NewErrParamMinValue("DesiredNumber", 2))
	}

	if s.DirectoryId == nil {
		invalidParams.Add(aws.NewErrParamRequired("DirectoryId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type UpdateNumberOfDomainControllersOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s UpdateNumberOfDomainControllersOutput) String() string {
	return awsutil.Prettify(s)
}

const opUpdateNumberOfDomainControllers = "UpdateNumberOfDomainControllers"

// UpdateNumberOfDomainControllersRequest returns a request value for making API operation for
// AWS Directory Service.
//
// Adds or removes domain controllers to or from the directory. Based on the
// difference between current value and new value (provided through this API
// call), domain controllers will be added or removed. It may take up to 45
// minutes for any new domain controllers to become fully active once the requested
// number of domain controllers is updated. During this time, you cannot make
// another update request.
//
//    // Example sending a request using UpdateNumberOfDomainControllersRequest.
//    req := client.UpdateNumberOfDomainControllersRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ds-2015-04-16/UpdateNumberOfDomainControllers
func (c *Client) UpdateNumberOfDomainControllersRequest(input *UpdateNumberOfDomainControllersInput) UpdateNumberOfDomainControllersRequest {
	op := &aws.Operation{
		Name:       opUpdateNumberOfDomainControllers,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &UpdateNumberOfDomainControllersInput{}
	}

	req := c.newRequest(op, input, &UpdateNumberOfDomainControllersOutput{})

	return UpdateNumberOfDomainControllersRequest{Request: req, Input: input, Copy: c.UpdateNumberOfDomainControllersRequest}
}

// UpdateNumberOfDomainControllersRequest is the request type for the
// UpdateNumberOfDomainControllers API operation.
type UpdateNumberOfDomainControllersRequest struct {
	*aws.Request
	Input *UpdateNumberOfDomainControllersInput
	Copy  func(*UpdateNumberOfDomainControllersInput) UpdateNumberOfDomainControllersRequest
}

// Send marshals and sends the UpdateNumberOfDomainControllers API request.
func (r UpdateNumberOfDomainControllersRequest) Send(ctx context.Context) (*UpdateNumberOfDomainControllersResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &UpdateNumberOfDomainControllersResponse{
		UpdateNumberOfDomainControllersOutput: r.Request.Data.(*UpdateNumberOfDomainControllersOutput),
		response:                              &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// UpdateNumberOfDomainControllersResponse is the response type for the
// UpdateNumberOfDomainControllers API operation.
type UpdateNumberOfDomainControllersResponse struct {
	*UpdateNumberOfDomainControllersOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// UpdateNumberOfDomainControllers request.
func (r *UpdateNumberOfDomainControllersResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
