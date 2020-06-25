// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package cloudsearch

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
)

// Container for the parameters to the UpdateScalingParameters operation. Specifies
// the name of the domain you want to update and the scaling parameters you
// want to configure.
type UpdateScalingParametersInput struct {
	_ struct{} `type:"structure"`

	// A string that represents the name of a domain. Domain names are unique across
	// the domains owned by an account within an AWS region. Domain names start
	// with a letter or number and can contain the following characters: a-z (lowercase),
	// 0-9, and - (hyphen).
	//
	// DomainName is a required field
	DomainName *string `min:"3" type:"string" required:"true"`

	// The desired instance type and desired number of replicas of each index partition.
	//
	// ScalingParameters is a required field
	ScalingParameters *ScalingParameters `type:"structure" required:"true"`
}

// String returns the string representation
func (s UpdateScalingParametersInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UpdateScalingParametersInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "UpdateScalingParametersInput"}

	if s.DomainName == nil {
		invalidParams.Add(aws.NewErrParamRequired("DomainName"))
	}
	if s.DomainName != nil && len(*s.DomainName) < 3 {
		invalidParams.Add(aws.NewErrParamMinLen("DomainName", 3))
	}

	if s.ScalingParameters == nil {
		invalidParams.Add(aws.NewErrParamRequired("ScalingParameters"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// The result of a UpdateScalingParameters request. Contains the status of the
// newly-configured scaling parameters.
type UpdateScalingParametersOutput struct {
	_ struct{} `type:"structure"`

	// The status and configuration of a search domain's scaling parameters.
	//
	// ScalingParameters is a required field
	ScalingParameters *ScalingParametersStatus `type:"structure" required:"true"`
}

// String returns the string representation
func (s UpdateScalingParametersOutput) String() string {
	return awsutil.Prettify(s)
}

const opUpdateScalingParameters = "UpdateScalingParameters"

// UpdateScalingParametersRequest returns a request value for making API operation for
// Amazon CloudSearch.
//
// Configures scaling parameters for a domain. A domain's scaling parameters
// specify the desired search instance type and replication count. Amazon CloudSearch
// will still automatically scale your domain based on the volume of data and
// traffic, but not below the desired instance type and replication count. If
// the Multi-AZ option is enabled, these values control the resources used per
// Availability Zone. For more information, see Configuring Scaling Options
// (http://docs.aws.amazon.com/cloudsearch/latest/developerguide/configuring-scaling-options.html)
// in the Amazon CloudSearch Developer Guide.
//
//    // Example sending a request using UpdateScalingParametersRequest.
//    req := client.UpdateScalingParametersRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
func (c *Client) UpdateScalingParametersRequest(input *UpdateScalingParametersInput) UpdateScalingParametersRequest {
	op := &aws.Operation{
		Name:       opUpdateScalingParameters,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &UpdateScalingParametersInput{}
	}

	req := c.newRequest(op, input, &UpdateScalingParametersOutput{})

	return UpdateScalingParametersRequest{Request: req, Input: input, Copy: c.UpdateScalingParametersRequest}
}

// UpdateScalingParametersRequest is the request type for the
// UpdateScalingParameters API operation.
type UpdateScalingParametersRequest struct {
	*aws.Request
	Input *UpdateScalingParametersInput
	Copy  func(*UpdateScalingParametersInput) UpdateScalingParametersRequest
}

// Send marshals and sends the UpdateScalingParameters API request.
func (r UpdateScalingParametersRequest) Send(ctx context.Context) (*UpdateScalingParametersResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &UpdateScalingParametersResponse{
		UpdateScalingParametersOutput: r.Request.Data.(*UpdateScalingParametersOutput),
		response:                      &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// UpdateScalingParametersResponse is the response type for the
// UpdateScalingParameters API operation.
type UpdateScalingParametersResponse struct {
	*UpdateScalingParametersOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// UpdateScalingParameters request.
func (r *UpdateScalingParametersResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
