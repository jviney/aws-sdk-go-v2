// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package amplify

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
	"github.com/jviney/aws-sdk-go-v2/private/protocol"
)

// Request structure for create a new deployment.
type CreateDeploymentInput struct {
	_ struct{} `type:"structure"`

	// Unique Id for an Amplify App.
	//
	// AppId is a required field
	AppId *string `location:"uri" locationName:"appId" min:"1" type:"string" required:"true"`

	// Name for the branch, for the Job.
	//
	// BranchName is a required field
	BranchName *string `location:"uri" locationName:"branchName" min:"1" type:"string" required:"true"`

	// Optional file map that contains file name as the key and file content md5
	// hash as the value. If this argument is provided, the service will generate
	// different upload url per file. Otherwise, the service will only generate
	// a single upload url for the zipped files.
	FileMap map[string]string `locationName:"fileMap" type:"map"`
}

// String returns the string representation
func (s CreateDeploymentInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateDeploymentInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateDeploymentInput"}

	if s.AppId == nil {
		invalidParams.Add(aws.NewErrParamRequired("AppId"))
	}
	if s.AppId != nil && len(*s.AppId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("AppId", 1))
	}

	if s.BranchName == nil {
		invalidParams.Add(aws.NewErrParamRequired("BranchName"))
	}
	if s.BranchName != nil && len(*s.BranchName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("BranchName", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s CreateDeploymentInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.FileMap != nil {
		v := s.FileMap

		metadata := protocol.Metadata{}
		ms0 := e.Map(protocol.BodyTarget, "fileMap", metadata)
		ms0.Start()
		for k1, v1 := range v {
			ms0.MapSetValue(k1, protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v1)})
		}
		ms0.End()

	}
	if s.AppId != nil {
		v := *s.AppId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "appId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.BranchName != nil {
		v := *s.BranchName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "branchName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// Result structure for create a new deployment.
type CreateDeploymentOutput struct {
	_ struct{} `type:"structure"`

	// When the fileMap argument is provided in the request, the fileUploadUrls
	// will contain a map of file names to upload url.
	//
	// FileUploadUrls is a required field
	FileUploadUrls map[string]string `locationName:"fileUploadUrls" type:"map" required:"true"`

	// The jobId for this deployment, will supply to start deployment api.
	JobId *string `locationName:"jobId" type:"string"`

	// When the fileMap argument is NOT provided. This zipUploadUrl will be returned.
	//
	// ZipUploadUrl is a required field
	ZipUploadUrl *string `locationName:"zipUploadUrl" type:"string" required:"true"`
}

// String returns the string representation
func (s CreateDeploymentOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s CreateDeploymentOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.FileUploadUrls != nil {
		v := s.FileUploadUrls

		metadata := protocol.Metadata{}
		ms0 := e.Map(protocol.BodyTarget, "fileUploadUrls", metadata)
		ms0.Start()
		for k1, v1 := range v {
			ms0.MapSetValue(k1, protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v1)})
		}
		ms0.End()

	}
	if s.JobId != nil {
		v := *s.JobId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "jobId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.ZipUploadUrl != nil {
		v := *s.ZipUploadUrl

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "zipUploadUrl", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

const opCreateDeployment = "CreateDeployment"

// CreateDeploymentRequest returns a request value for making API operation for
// AWS Amplify.
//
// Create a deployment for manual deploy apps. (Apps are not connected to repository)
//
//    // Example sending a request using CreateDeploymentRequest.
//    req := client.CreateDeploymentRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/amplify-2017-07-25/CreateDeployment
func (c *Client) CreateDeploymentRequest(input *CreateDeploymentInput) CreateDeploymentRequest {
	op := &aws.Operation{
		Name:       opCreateDeployment,
		HTTPMethod: "POST",
		HTTPPath:   "/apps/{appId}/branches/{branchName}/deployments",
	}

	if input == nil {
		input = &CreateDeploymentInput{}
	}

	req := c.newRequest(op, input, &CreateDeploymentOutput{})

	return CreateDeploymentRequest{Request: req, Input: input, Copy: c.CreateDeploymentRequest}
}

// CreateDeploymentRequest is the request type for the
// CreateDeployment API operation.
type CreateDeploymentRequest struct {
	*aws.Request
	Input *CreateDeploymentInput
	Copy  func(*CreateDeploymentInput) CreateDeploymentRequest
}

// Send marshals and sends the CreateDeployment API request.
func (r CreateDeploymentRequest) Send(ctx context.Context) (*CreateDeploymentResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateDeploymentResponse{
		CreateDeploymentOutput: r.Request.Data.(*CreateDeploymentOutput),
		response:               &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateDeploymentResponse is the response type for the
// CreateDeployment API operation.
type CreateDeploymentResponse struct {
	*CreateDeploymentOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateDeployment request.
func (r *CreateDeploymentResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
