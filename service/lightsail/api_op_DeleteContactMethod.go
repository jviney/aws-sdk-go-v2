// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package lightsail

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
)

type DeleteContactMethodInput struct {
	_ struct{} `type:"structure"`

	// The protocol that will be deleted, such as Email or SMS (text messaging).
	//
	// To delete an Email and an SMS contact method if you added both, you must
	// run separate DeleteContactMethod actions to delete each protocol.
	//
	// Protocol is a required field
	Protocol ContactProtocol `locationName:"protocol" type:"string" required:"true" enum:"true"`
}

// String returns the string representation
func (s DeleteContactMethodInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteContactMethodInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteContactMethodInput"}
	if len(s.Protocol) == 0 {
		invalidParams.Add(aws.NewErrParamRequired("Protocol"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DeleteContactMethodOutput struct {
	_ struct{} `type:"structure"`

	// An array of objects that describe the result of the action, such as the status
	// of the request, the timestamp of the request, and the resources affected
	// by the request.
	Operations []Operation `locationName:"operations" type:"list"`
}

// String returns the string representation
func (s DeleteContactMethodOutput) String() string {
	return awsutil.Prettify(s)
}

const opDeleteContactMethod = "DeleteContactMethod"

// DeleteContactMethodRequest returns a request value for making API operation for
// Amazon Lightsail.
//
// Deletes a contact method.
//
// A contact method is used to send you notifications about your Amazon Lightsail
// resources. You can add one email address and one mobile phone number contact
// method in each AWS Region. However, SMS text messaging is not supported in
// some AWS Regions, and SMS text messages cannot be sent to some countries/regions.
// For more information, see Notifications in Amazon Lightsail (https://lightsail.aws.amazon.com/ls/docs/en_us/articles/amazon-lightsail-notifications).
//
//    // Example sending a request using DeleteContactMethodRequest.
//    req := client.DeleteContactMethodRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/lightsail-2016-11-28/DeleteContactMethod
func (c *Client) DeleteContactMethodRequest(input *DeleteContactMethodInput) DeleteContactMethodRequest {
	op := &aws.Operation{
		Name:       opDeleteContactMethod,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DeleteContactMethodInput{}
	}

	req := c.newRequest(op, input, &DeleteContactMethodOutput{})

	return DeleteContactMethodRequest{Request: req, Input: input, Copy: c.DeleteContactMethodRequest}
}

// DeleteContactMethodRequest is the request type for the
// DeleteContactMethod API operation.
type DeleteContactMethodRequest struct {
	*aws.Request
	Input *DeleteContactMethodInput
	Copy  func(*DeleteContactMethodInput) DeleteContactMethodRequest
}

// Send marshals and sends the DeleteContactMethod API request.
func (r DeleteContactMethodRequest) Send(ctx context.Context) (*DeleteContactMethodResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteContactMethodResponse{
		DeleteContactMethodOutput: r.Request.Data.(*DeleteContactMethodOutput),
		response:                  &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteContactMethodResponse is the response type for the
// DeleteContactMethod API operation.
type DeleteContactMethodResponse struct {
	*DeleteContactMethodOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteContactMethod request.
func (r *DeleteContactMethodResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
