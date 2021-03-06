// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package quicksight

import (
	"context"
	"fmt"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
	"github.com/jviney/aws-sdk-go-v2/private/protocol"
)

type CreateDataSourceInput struct {
	_ struct{} `type:"structure"`

	// The AWS account ID.
	//
	// AwsAccountId is a required field
	AwsAccountId *string `location:"uri" locationName:"AwsAccountId" min:"12" type:"string" required:"true"`

	// The credentials QuickSight that uses to connect to your underlying source.
	// Currently, only credentials based on user name and password are supported.
	Credentials *DataSourceCredentials `type:"structure" sensitive:"true"`

	// An ID for the data source. This ID is unique per AWS Region for each AWS
	// account.
	//
	// DataSourceId is a required field
	DataSourceId *string `type:"string" required:"true"`

	// The parameters that QuickSight uses to connect to your underlying source.
	DataSourceParameters *DataSourceParameters `type:"structure"`

	// A display name for the data source.
	//
	// Name is a required field
	Name *string `min:"1" type:"string" required:"true"`

	// A list of resource permissions on the data source.
	Permissions []ResourcePermission `min:"1" type:"list"`

	// Secure Socket Layer (SSL) properties that apply when QuickSight connects
	// to your underlying source.
	SslProperties *SslProperties `type:"structure"`

	// Contains a map of the key-value pairs for the resource tag or tags assigned
	// to the data source.
	Tags []Tag `min:"1" type:"list"`

	// The type of the data source. Currently, the supported types for this operation
	// are: ATHENA, AURORA, AURORA_POSTGRESQL, MARIADB, MYSQL, POSTGRESQL, PRESTO,
	// REDSHIFT, S3, SNOWFLAKE, SPARK, SQLSERVER, TERADATA. Use ListDataSources
	// to return a list of all data sources.
	//
	// Type is a required field
	Type DataSourceType `type:"string" required:"true" enum:"true"`

	// Use this parameter only when you want QuickSight to use a VPC connection
	// when connecting to your underlying source.
	VpcConnectionProperties *VpcConnectionProperties `type:"structure"`
}

// String returns the string representation
func (s CreateDataSourceInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateDataSourceInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateDataSourceInput"}

	if s.AwsAccountId == nil {
		invalidParams.Add(aws.NewErrParamRequired("AwsAccountId"))
	}
	if s.AwsAccountId != nil && len(*s.AwsAccountId) < 12 {
		invalidParams.Add(aws.NewErrParamMinLen("AwsAccountId", 12))
	}

	if s.DataSourceId == nil {
		invalidParams.Add(aws.NewErrParamRequired("DataSourceId"))
	}

	if s.Name == nil {
		invalidParams.Add(aws.NewErrParamRequired("Name"))
	}
	if s.Name != nil && len(*s.Name) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Name", 1))
	}
	if s.Permissions != nil && len(s.Permissions) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Permissions", 1))
	}
	if s.Tags != nil && len(s.Tags) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Tags", 1))
	}
	if len(s.Type) == 0 {
		invalidParams.Add(aws.NewErrParamRequired("Type"))
	}
	if s.Credentials != nil {
		if err := s.Credentials.Validate(); err != nil {
			invalidParams.AddNested("Credentials", err.(aws.ErrInvalidParams))
		}
	}
	if s.DataSourceParameters != nil {
		if err := s.DataSourceParameters.Validate(); err != nil {
			invalidParams.AddNested("DataSourceParameters", err.(aws.ErrInvalidParams))
		}
	}
	if s.Permissions != nil {
		for i, v := range s.Permissions {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "Permissions", i), err.(aws.ErrInvalidParams))
			}
		}
	}
	if s.Tags != nil {
		for i, v := range s.Tags {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "Tags", i), err.(aws.ErrInvalidParams))
			}
		}
	}
	if s.VpcConnectionProperties != nil {
		if err := s.VpcConnectionProperties.Validate(); err != nil {
			invalidParams.AddNested("VpcConnectionProperties", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s CreateDataSourceInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.Credentials != nil {
		v := s.Credentials

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "Credentials", v, metadata)
	}
	if s.DataSourceId != nil {
		v := *s.DataSourceId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "DataSourceId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.DataSourceParameters != nil {
		v := s.DataSourceParameters

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "DataSourceParameters", v, metadata)
	}
	if s.Name != nil {
		v := *s.Name

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "Name", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Permissions != nil {
		v := s.Permissions

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "Permissions", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	if s.SslProperties != nil {
		v := s.SslProperties

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "SslProperties", v, metadata)
	}
	if s.Tags != nil {
		v := s.Tags

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "Tags", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	if len(s.Type) > 0 {
		v := s.Type

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "Type", protocol.QuotedValue{ValueMarshaler: v}, metadata)
	}
	if s.VpcConnectionProperties != nil {
		v := s.VpcConnectionProperties

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "VpcConnectionProperties", v, metadata)
	}
	if s.AwsAccountId != nil {
		v := *s.AwsAccountId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "AwsAccountId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type CreateDataSourceOutput struct {
	_ struct{} `type:"structure"`

	// The Amazon Resource Name (ARN) of the data source.
	Arn *string `type:"string"`

	// The status of creating the data source.
	CreationStatus ResourceStatus `type:"string" enum:"true"`

	// The ID of the data source. This ID is unique per AWS Region for each AWS
	// account.
	DataSourceId *string `type:"string"`

	// The AWS request ID for this operation.
	RequestId *string `type:"string"`

	// The HTTP status of the request.
	Status *int64 `location:"statusCode" type:"integer"`
}

// String returns the string representation
func (s CreateDataSourceOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s CreateDataSourceOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.Arn != nil {
		v := *s.Arn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "Arn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if len(s.CreationStatus) > 0 {
		v := s.CreationStatus

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "CreationStatus", protocol.QuotedValue{ValueMarshaler: v}, metadata)
	}
	if s.DataSourceId != nil {
		v := *s.DataSourceId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "DataSourceId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.RequestId != nil {
		v := *s.RequestId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "RequestId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	// ignoring invalid encode state, StatusCode. Status
	return nil
}

const opCreateDataSource = "CreateDataSource"

// CreateDataSourceRequest returns a request value for making API operation for
// Amazon QuickSight.
//
// Creates a data source.
//
//    // Example sending a request using CreateDataSourceRequest.
//    req := client.CreateDataSourceRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/quicksight-2018-04-01/CreateDataSource
func (c *Client) CreateDataSourceRequest(input *CreateDataSourceInput) CreateDataSourceRequest {
	op := &aws.Operation{
		Name:       opCreateDataSource,
		HTTPMethod: "POST",
		HTTPPath:   "/accounts/{AwsAccountId}/data-sources",
	}

	if input == nil {
		input = &CreateDataSourceInput{}
	}

	req := c.newRequest(op, input, &CreateDataSourceOutput{})

	return CreateDataSourceRequest{Request: req, Input: input, Copy: c.CreateDataSourceRequest}
}

// CreateDataSourceRequest is the request type for the
// CreateDataSource API operation.
type CreateDataSourceRequest struct {
	*aws.Request
	Input *CreateDataSourceInput
	Copy  func(*CreateDataSourceInput) CreateDataSourceRequest
}

// Send marshals and sends the CreateDataSource API request.
func (r CreateDataSourceRequest) Send(ctx context.Context) (*CreateDataSourceResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateDataSourceResponse{
		CreateDataSourceOutput: r.Request.Data.(*CreateDataSourceOutput),
		response:               &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateDataSourceResponse is the response type for the
// CreateDataSource API operation.
type CreateDataSourceResponse struct {
	*CreateDataSourceOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateDataSource request.
func (r *CreateDataSourceResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
