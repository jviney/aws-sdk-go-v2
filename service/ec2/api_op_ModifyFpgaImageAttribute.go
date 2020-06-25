// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ec2

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
)

type ModifyFpgaImageAttributeInput struct {
	_ struct{} `type:"structure"`

	// The name of the attribute.
	Attribute FpgaImageAttributeName `type:"string" enum:"true"`

	// A description for the AFI.
	Description *string `type:"string"`

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have
	// the required permissions, the error response is DryRunOperation. Otherwise,
	// it is UnauthorizedOperation.
	DryRun *bool `type:"boolean"`

	// The ID of the AFI.
	//
	// FpgaImageId is a required field
	FpgaImageId *string `type:"string" required:"true"`

	// The load permission for the AFI.
	LoadPermission *LoadPermissionModifications `type:"structure"`

	// A name for the AFI.
	Name *string `type:"string"`

	// The operation type.
	OperationType OperationType `type:"string" enum:"true"`

	// The product codes. After you add a product code to an AFI, it can't be removed.
	// This parameter is valid only when modifying the productCodes attribute.
	ProductCodes []string `locationName:"ProductCode" locationNameList:"ProductCode" type:"list"`

	// The user groups. This parameter is valid only when modifying the loadPermission
	// attribute.
	UserGroups []string `locationName:"UserGroup" locationNameList:"UserGroup" type:"list"`

	// The AWS account IDs. This parameter is valid only when modifying the loadPermission
	// attribute.
	UserIds []string `locationName:"UserId" locationNameList:"UserId" type:"list"`
}

// String returns the string representation
func (s ModifyFpgaImageAttributeInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ModifyFpgaImageAttributeInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ModifyFpgaImageAttributeInput"}

	if s.FpgaImageId == nil {
		invalidParams.Add(aws.NewErrParamRequired("FpgaImageId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type ModifyFpgaImageAttributeOutput struct {
	_ struct{} `type:"structure"`

	// Information about the attribute.
	FpgaImageAttribute *FpgaImageAttribute `locationName:"fpgaImageAttribute" type:"structure"`
}

// String returns the string representation
func (s ModifyFpgaImageAttributeOutput) String() string {
	return awsutil.Prettify(s)
}

const opModifyFpgaImageAttribute = "ModifyFpgaImageAttribute"

// ModifyFpgaImageAttributeRequest returns a request value for making API operation for
// Amazon Elastic Compute Cloud.
//
// Modifies the specified attribute of the specified Amazon FPGA Image (AFI).
//
//    // Example sending a request using ModifyFpgaImageAttributeRequest.
//    req := client.ModifyFpgaImageAttributeRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/ModifyFpgaImageAttribute
func (c *Client) ModifyFpgaImageAttributeRequest(input *ModifyFpgaImageAttributeInput) ModifyFpgaImageAttributeRequest {
	op := &aws.Operation{
		Name:       opModifyFpgaImageAttribute,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ModifyFpgaImageAttributeInput{}
	}

	req := c.newRequest(op, input, &ModifyFpgaImageAttributeOutput{})

	return ModifyFpgaImageAttributeRequest{Request: req, Input: input, Copy: c.ModifyFpgaImageAttributeRequest}
}

// ModifyFpgaImageAttributeRequest is the request type for the
// ModifyFpgaImageAttribute API operation.
type ModifyFpgaImageAttributeRequest struct {
	*aws.Request
	Input *ModifyFpgaImageAttributeInput
	Copy  func(*ModifyFpgaImageAttributeInput) ModifyFpgaImageAttributeRequest
}

// Send marshals and sends the ModifyFpgaImageAttribute API request.
func (r ModifyFpgaImageAttributeRequest) Send(ctx context.Context) (*ModifyFpgaImageAttributeResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ModifyFpgaImageAttributeResponse{
		ModifyFpgaImageAttributeOutput: r.Request.Data.(*ModifyFpgaImageAttributeOutput),
		response:                       &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// ModifyFpgaImageAttributeResponse is the response type for the
// ModifyFpgaImageAttribute API operation.
type ModifyFpgaImageAttributeResponse struct {
	*ModifyFpgaImageAttributeOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ModifyFpgaImageAttribute request.
func (r *ModifyFpgaImageAttributeResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
