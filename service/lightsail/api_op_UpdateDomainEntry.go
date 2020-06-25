// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package lightsail

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
)

type UpdateDomainEntryInput struct {
	_ struct{} `type:"structure"`

	// An array of key-value pairs containing information about the domain entry.
	//
	// DomainEntry is a required field
	DomainEntry *DomainEntry `locationName:"domainEntry" type:"structure" required:"true"`

	// The name of the domain recordset to update.
	//
	// DomainName is a required field
	DomainName *string `locationName:"domainName" type:"string" required:"true"`
}

// String returns the string representation
func (s UpdateDomainEntryInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UpdateDomainEntryInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "UpdateDomainEntryInput"}

	if s.DomainEntry == nil {
		invalidParams.Add(aws.NewErrParamRequired("DomainEntry"))
	}

	if s.DomainName == nil {
		invalidParams.Add(aws.NewErrParamRequired("DomainName"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type UpdateDomainEntryOutput struct {
	_ struct{} `type:"structure"`

	// An array of objects that describe the result of the action, such as the status
	// of the request, the timestamp of the request, and the resources affected
	// by the request.
	Operations []Operation `locationName:"operations" type:"list"`
}

// String returns the string representation
func (s UpdateDomainEntryOutput) String() string {
	return awsutil.Prettify(s)
}

const opUpdateDomainEntry = "UpdateDomainEntry"

// UpdateDomainEntryRequest returns a request value for making API operation for
// Amazon Lightsail.
//
// Updates a domain recordset after it is created.
//
// The update domain entry operation supports tag-based access control via resource
// tags applied to the resource identified by domain name. For more information,
// see the Lightsail Dev Guide (https://lightsail.aws.amazon.com/ls/docs/en/articles/amazon-lightsail-controlling-access-using-tags).
//
//    // Example sending a request using UpdateDomainEntryRequest.
//    req := client.UpdateDomainEntryRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/lightsail-2016-11-28/UpdateDomainEntry
func (c *Client) UpdateDomainEntryRequest(input *UpdateDomainEntryInput) UpdateDomainEntryRequest {
	op := &aws.Operation{
		Name:       opUpdateDomainEntry,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &UpdateDomainEntryInput{}
	}

	req := c.newRequest(op, input, &UpdateDomainEntryOutput{})

	return UpdateDomainEntryRequest{Request: req, Input: input, Copy: c.UpdateDomainEntryRequest}
}

// UpdateDomainEntryRequest is the request type for the
// UpdateDomainEntry API operation.
type UpdateDomainEntryRequest struct {
	*aws.Request
	Input *UpdateDomainEntryInput
	Copy  func(*UpdateDomainEntryInput) UpdateDomainEntryRequest
}

// Send marshals and sends the UpdateDomainEntry API request.
func (r UpdateDomainEntryRequest) Send(ctx context.Context) (*UpdateDomainEntryResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &UpdateDomainEntryResponse{
		UpdateDomainEntryOutput: r.Request.Data.(*UpdateDomainEntryOutput),
		response:                &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// UpdateDomainEntryResponse is the response type for the
// UpdateDomainEntry API operation.
type UpdateDomainEntryResponse struct {
	*UpdateDomainEntryOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// UpdateDomainEntry request.
func (r *UpdateDomainEntryResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
