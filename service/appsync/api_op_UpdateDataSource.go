// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package appsync

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
	"github.com/jviney/aws-sdk-go-v2/private/protocol"
)

type UpdateDataSourceInput struct {
	_ struct{} `type:"structure"`

	// The API ID.
	//
	// ApiId is a required field
	ApiId *string `location:"uri" locationName:"apiId" type:"string" required:"true"`

	// The new description for the data source.
	Description *string `locationName:"description" type:"string"`

	// The new Amazon DynamoDB configuration.
	DynamodbConfig *DynamodbDataSourceConfig `locationName:"dynamodbConfig" type:"structure"`

	// The new Elasticsearch Service configuration.
	ElasticsearchConfig *ElasticsearchDataSourceConfig `locationName:"elasticsearchConfig" type:"structure"`

	// The new HTTP endpoint configuration.
	HttpConfig *HttpDataSourceConfig `locationName:"httpConfig" type:"structure"`

	// The new AWS Lambda configuration.
	LambdaConfig *LambdaDataSourceConfig `locationName:"lambdaConfig" type:"structure"`

	// The new name for the data source.
	//
	// Name is a required field
	Name *string `location:"uri" locationName:"name" min:"1" type:"string" required:"true"`

	// The new relational database configuration.
	RelationalDatabaseConfig *RelationalDatabaseDataSourceConfig `locationName:"relationalDatabaseConfig" type:"structure"`

	// The new service role ARN for the data source.
	ServiceRoleArn *string `locationName:"serviceRoleArn" type:"string"`

	// The new data source type.
	//
	// Type is a required field
	Type DataSourceType `locationName:"type" type:"string" required:"true" enum:"true"`
}

// String returns the string representation
func (s UpdateDataSourceInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UpdateDataSourceInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "UpdateDataSourceInput"}

	if s.ApiId == nil {
		invalidParams.Add(aws.NewErrParamRequired("ApiId"))
	}

	if s.Name == nil {
		invalidParams.Add(aws.NewErrParamRequired("Name"))
	}
	if s.Name != nil && len(*s.Name) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Name", 1))
	}
	if len(s.Type) == 0 {
		invalidParams.Add(aws.NewErrParamRequired("Type"))
	}
	if s.DynamodbConfig != nil {
		if err := s.DynamodbConfig.Validate(); err != nil {
			invalidParams.AddNested("DynamodbConfig", err.(aws.ErrInvalidParams))
		}
	}
	if s.ElasticsearchConfig != nil {
		if err := s.ElasticsearchConfig.Validate(); err != nil {
			invalidParams.AddNested("ElasticsearchConfig", err.(aws.ErrInvalidParams))
		}
	}
	if s.HttpConfig != nil {
		if err := s.HttpConfig.Validate(); err != nil {
			invalidParams.AddNested("HttpConfig", err.(aws.ErrInvalidParams))
		}
	}
	if s.LambdaConfig != nil {
		if err := s.LambdaConfig.Validate(); err != nil {
			invalidParams.AddNested("LambdaConfig", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s UpdateDataSourceInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.Description != nil {
		v := *s.Description

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "description", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.DynamodbConfig != nil {
		v := s.DynamodbConfig

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "dynamodbConfig", v, metadata)
	}
	if s.ElasticsearchConfig != nil {
		v := s.ElasticsearchConfig

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "elasticsearchConfig", v, metadata)
	}
	if s.HttpConfig != nil {
		v := s.HttpConfig

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "httpConfig", v, metadata)
	}
	if s.LambdaConfig != nil {
		v := s.LambdaConfig

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "lambdaConfig", v, metadata)
	}
	if s.RelationalDatabaseConfig != nil {
		v := s.RelationalDatabaseConfig

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "relationalDatabaseConfig", v, metadata)
	}
	if s.ServiceRoleArn != nil {
		v := *s.ServiceRoleArn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "serviceRoleArn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if len(s.Type) > 0 {
		v := s.Type

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "type", protocol.QuotedValue{ValueMarshaler: v}, metadata)
	}
	if s.ApiId != nil {
		v := *s.ApiId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "apiId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Name != nil {
		v := *s.Name

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "name", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type UpdateDataSourceOutput struct {
	_ struct{} `type:"structure"`

	// The updated DataSource object.
	DataSource *DataSource `locationName:"dataSource" type:"structure"`
}

// String returns the string representation
func (s UpdateDataSourceOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s UpdateDataSourceOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.DataSource != nil {
		v := s.DataSource

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "dataSource", v, metadata)
	}
	return nil
}

const opUpdateDataSource = "UpdateDataSource"

// UpdateDataSourceRequest returns a request value for making API operation for
// AWS AppSync.
//
// Updates a DataSource object.
//
//    // Example sending a request using UpdateDataSourceRequest.
//    req := client.UpdateDataSourceRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/appsync-2017-07-25/UpdateDataSource
func (c *Client) UpdateDataSourceRequest(input *UpdateDataSourceInput) UpdateDataSourceRequest {
	op := &aws.Operation{
		Name:       opUpdateDataSource,
		HTTPMethod: "POST",
		HTTPPath:   "/v1/apis/{apiId}/datasources/{name}",
	}

	if input == nil {
		input = &UpdateDataSourceInput{}
	}

	req := c.newRequest(op, input, &UpdateDataSourceOutput{})

	return UpdateDataSourceRequest{Request: req, Input: input, Copy: c.UpdateDataSourceRequest}
}

// UpdateDataSourceRequest is the request type for the
// UpdateDataSource API operation.
type UpdateDataSourceRequest struct {
	*aws.Request
	Input *UpdateDataSourceInput
	Copy  func(*UpdateDataSourceInput) UpdateDataSourceRequest
}

// Send marshals and sends the UpdateDataSource API request.
func (r UpdateDataSourceRequest) Send(ctx context.Context) (*UpdateDataSourceResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &UpdateDataSourceResponse{
		UpdateDataSourceOutput: r.Request.Data.(*UpdateDataSourceOutput),
		response:               &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// UpdateDataSourceResponse is the response type for the
// UpdateDataSource API operation.
type UpdateDataSourceResponse struct {
	*UpdateDataSourceOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// UpdateDataSource request.
func (r *UpdateDataSourceResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
