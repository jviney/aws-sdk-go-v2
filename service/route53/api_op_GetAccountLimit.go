// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package route53

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
	"github.com/jviney/aws-sdk-go-v2/private/protocol"
)

// A complex type that contains information about the request to create a hosted
// zone.
type GetAccountLimitInput struct {
	_ struct{} `type:"structure"`

	// The limit that you want to get. Valid values include the following:
	//
	//    * MAX_HEALTH_CHECKS_BY_OWNER: The maximum number of health checks that
	//    you can create using the current account.
	//
	//    * MAX_HOSTED_ZONES_BY_OWNER: The maximum number of hosted zones that you
	//    can create using the current account.
	//
	//    * MAX_REUSABLE_DELEGATION_SETS_BY_OWNER: The maximum number of reusable
	//    delegation sets that you can create using the current account.
	//
	//    * MAX_TRAFFIC_POLICIES_BY_OWNER: The maximum number of traffic policies
	//    that you can create using the current account.
	//
	//    * MAX_TRAFFIC_POLICY_INSTANCES_BY_OWNER: The maximum number of traffic
	//    policy instances that you can create using the current account. (Traffic
	//    policy instances are referred to as traffic flow policy records in the
	//    Amazon Route 53 console.)
	//
	// Type is a required field
	Type AccountLimitType `location:"uri" locationName:"Type" type:"string" required:"true" enum:"true"`
}

// String returns the string representation
func (s GetAccountLimitInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetAccountLimitInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetAccountLimitInput"}
	if len(s.Type) == 0 {
		invalidParams.Add(aws.NewErrParamRequired("Type"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetAccountLimitInput) MarshalFields(e protocol.FieldEncoder) error {

	if len(s.Type) > 0 {
		v := s.Type

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "Type", v, metadata)
	}
	return nil
}

// A complex type that contains the requested limit.
type GetAccountLimitOutput struct {
	_ struct{} `type:"structure"`

	// The current number of entities that you have created of the specified type.
	// For example, if you specified MAX_HEALTH_CHECKS_BY_OWNER for the value of
	// Type in the request, the value of Count is the current number of health checks
	// that you have created using the current account.
	//
	// Count is a required field
	Count *int64 `type:"long" required:"true"`

	// The current setting for the specified limit. For example, if you specified
	// MAX_HEALTH_CHECKS_BY_OWNER for the value of Type in the request, the value
	// of Limit is the maximum number of health checks that you can create using
	// the current account.
	//
	// Limit is a required field
	Limit *AccountLimit `type:"structure" required:"true"`
}

// String returns the string representation
func (s GetAccountLimitOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetAccountLimitOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.Count != nil {
		v := *s.Count

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "Count", protocol.Int64Value(v), metadata)
	}
	if s.Limit != nil {
		v := s.Limit

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "Limit", v, metadata)
	}
	return nil
}

const opGetAccountLimit = "GetAccountLimit"

// GetAccountLimitRequest returns a request value for making API operation for
// Amazon Route 53.
//
// Gets the specified limit for the current account, for example, the maximum
// number of health checks that you can create using the account.
//
// For the default limit, see Limits (https://docs.aws.amazon.com/Route53/latest/DeveloperGuide/DNSLimitations.html)
// in the Amazon Route 53 Developer Guide. To request a higher limit, open a
// case (https://console.aws.amazon.com/support/home#/case/create?issueType=service-limit-increase&limitType=service-code-route53).
//
// You can also view account limits in AWS Trusted Advisor. Sign in to the AWS
// Management Console and open the Trusted Advisor console at https://console.aws.amazon.com/trustedadvisor/
// (https://console.aws.amazon.com/trustedadvisor). Then choose Service limits
// in the navigation pane.
//
//    // Example sending a request using GetAccountLimitRequest.
//    req := client.GetAccountLimitRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/route53-2013-04-01/GetAccountLimit
func (c *Client) GetAccountLimitRequest(input *GetAccountLimitInput) GetAccountLimitRequest {
	op := &aws.Operation{
		Name:       opGetAccountLimit,
		HTTPMethod: "GET",
		HTTPPath:   "/2013-04-01/accountlimit/{Type}",
	}

	if input == nil {
		input = &GetAccountLimitInput{}
	}

	req := c.newRequest(op, input, &GetAccountLimitOutput{})

	return GetAccountLimitRequest{Request: req, Input: input, Copy: c.GetAccountLimitRequest}
}

// GetAccountLimitRequest is the request type for the
// GetAccountLimit API operation.
type GetAccountLimitRequest struct {
	*aws.Request
	Input *GetAccountLimitInput
	Copy  func(*GetAccountLimitInput) GetAccountLimitRequest
}

// Send marshals and sends the GetAccountLimit API request.
func (r GetAccountLimitRequest) Send(ctx context.Context) (*GetAccountLimitResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetAccountLimitResponse{
		GetAccountLimitOutput: r.Request.Data.(*GetAccountLimitOutput),
		response:              &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetAccountLimitResponse is the response type for the
// GetAccountLimit API operation.
type GetAccountLimitResponse struct {
	*GetAccountLimitOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetAccountLimit request.
func (r *GetAccountLimitResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
