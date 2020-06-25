// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ec2

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
)

type CreateFpgaImageInput struct {
	_ struct{} `type:"structure"`

	// Unique, case-sensitive identifier that you provide to ensure the idempotency
	// of the request. For more information, see Ensuring Idempotency (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/Run_Instance_Idempotency.html).
	ClientToken *string `type:"string"`

	// A description for the AFI.
	Description *string `type:"string"`

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have
	// the required permissions, the error response is DryRunOperation. Otherwise,
	// it is UnauthorizedOperation.
	DryRun *bool `type:"boolean"`

	// The location of the encrypted design checkpoint in Amazon S3. The input must
	// be a tarball.
	//
	// InputStorageLocation is a required field
	InputStorageLocation *StorageLocation `type:"structure" required:"true"`

	// The location in Amazon S3 for the output logs.
	LogsStorageLocation *StorageLocation `type:"structure"`

	// A name for the AFI.
	Name *string `type:"string"`

	// The tags to apply to the FPGA image during creation.
	TagSpecifications []TagSpecification `locationName:"TagSpecification" locationNameList:"item" type:"list"`
}

// String returns the string representation
func (s CreateFpgaImageInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateFpgaImageInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateFpgaImageInput"}

	if s.InputStorageLocation == nil {
		invalidParams.Add(aws.NewErrParamRequired("InputStorageLocation"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type CreateFpgaImageOutput struct {
	_ struct{} `type:"structure"`

	// The global FPGA image identifier (AGFI ID).
	FpgaImageGlobalId *string `locationName:"fpgaImageGlobalId" type:"string"`

	// The FPGA image identifier (AFI ID).
	FpgaImageId *string `locationName:"fpgaImageId" type:"string"`
}

// String returns the string representation
func (s CreateFpgaImageOutput) String() string {
	return awsutil.Prettify(s)
}

const opCreateFpgaImage = "CreateFpgaImage"

// CreateFpgaImageRequest returns a request value for making API operation for
// Amazon Elastic Compute Cloud.
//
// Creates an Amazon FPGA Image (AFI) from the specified design checkpoint (DCP).
//
// The create operation is asynchronous. To verify that the AFI is ready for
// use, check the output logs.
//
// An AFI contains the FPGA bitstream that is ready to download to an FPGA.
// You can securely deploy an AFI on multiple FPGA-accelerated instances. For
// more information, see the AWS FPGA Hardware Development Kit (https://github.com/aws/aws-fpga/).
//
//    // Example sending a request using CreateFpgaImageRequest.
//    req := client.CreateFpgaImageRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/CreateFpgaImage
func (c *Client) CreateFpgaImageRequest(input *CreateFpgaImageInput) CreateFpgaImageRequest {
	op := &aws.Operation{
		Name:       opCreateFpgaImage,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &CreateFpgaImageInput{}
	}

	req := c.newRequest(op, input, &CreateFpgaImageOutput{})

	return CreateFpgaImageRequest{Request: req, Input: input, Copy: c.CreateFpgaImageRequest}
}

// CreateFpgaImageRequest is the request type for the
// CreateFpgaImage API operation.
type CreateFpgaImageRequest struct {
	*aws.Request
	Input *CreateFpgaImageInput
	Copy  func(*CreateFpgaImageInput) CreateFpgaImageRequest
}

// Send marshals and sends the CreateFpgaImage API request.
func (r CreateFpgaImageRequest) Send(ctx context.Context) (*CreateFpgaImageResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateFpgaImageResponse{
		CreateFpgaImageOutput: r.Request.Data.(*CreateFpgaImageOutput),
		response:              &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateFpgaImageResponse is the response type for the
// CreateFpgaImage API operation.
type CreateFpgaImageResponse struct {
	*CreateFpgaImageOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateFpgaImage request.
func (r *CreateFpgaImageResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
