// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package forecastquery

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
)

type QueryForecastInput struct {
	_ struct{} `type:"structure"`

	// The end date for the forecast. Specify the date using this format: yyyy-MM-dd'T'HH:mm:ss
	// (ISO 8601 format). For example, 2015-01-01T20:00:00.
	EndDate *string `type:"string"`

	// The filtering criteria to apply when retrieving the forecast. For example,
	// to get the forecast for client_21 in the electricity usage dataset, specify
	// the following:
	//
	// {"item_id" : "client_21"}
	//
	// To get the full forecast, use the CreateForecastExportJob (https://docs.aws.amazon.com/en_us/forecast/latest/dg/API_CreateForecastExportJob.html)
	// operation.
	//
	// Filters is a required field
	Filters map[string]string `min:"1" type:"map" required:"true"`

	// The Amazon Resource Name (ARN) of the forecast to query.
	//
	// ForecastArn is a required field
	ForecastArn *string `type:"string" required:"true"`

	// If the result of the previous request was truncated, the response includes
	// a NextToken. To retrieve the next set of results, use the token in the next
	// request. Tokens expire after 24 hours.
	NextToken *string `min:"1" type:"string"`

	// The start date for the forecast. Specify the date using this format: yyyy-MM-dd'T'HH:mm:ss
	// (ISO 8601 format). For example, 2015-01-01T08:00:00.
	StartDate *string `type:"string"`
}

// String returns the string representation
func (s QueryForecastInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *QueryForecastInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "QueryForecastInput"}

	if s.Filters == nil {
		invalidParams.Add(aws.NewErrParamRequired("Filters"))
	}
	if s.Filters != nil && len(s.Filters) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Filters", 1))
	}

	if s.ForecastArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("ForecastArn"))
	}
	if s.NextToken != nil && len(*s.NextToken) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("NextToken", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type QueryForecastOutput struct {
	_ struct{} `type:"structure"`

	// The forecast.
	Forecast *Forecast `type:"structure"`
}

// String returns the string representation
func (s QueryForecastOutput) String() string {
	return awsutil.Prettify(s)
}

const opQueryForecast = "QueryForecast"

// QueryForecastRequest returns a request value for making API operation for
// Amazon Forecast Query Service.
//
// Retrieves a forecast for a single item, filtered by the supplied criteria.
//
// The criteria is a key-value pair. The key is either item_id (or the equivalent
// non-timestamp, non-target field) from the TARGET_TIME_SERIES dataset, or
// one of the forecast dimensions specified as part of the FeaturizationConfig
// object.
//
// By default, QueryForecast returns the complete date range for the filtered
// forecast. You can request a specific date range.
//
// To get the full forecast, use the CreateForecastExportJob (https://docs.aws.amazon.com/en_us/forecast/latest/dg/API_CreateForecastExportJob.html)
// operation.
//
// The forecasts generated by Amazon Forecast are in the same timezone as the
// dataset that was used to create the predictor.
//
//    // Example sending a request using QueryForecastRequest.
//    req := client.QueryForecastRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/forecastquery-2018-06-26/QueryForecast
func (c *Client) QueryForecastRequest(input *QueryForecastInput) QueryForecastRequest {
	op := &aws.Operation{
		Name:       opQueryForecast,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &QueryForecastInput{}
	}

	req := c.newRequest(op, input, &QueryForecastOutput{})

	return QueryForecastRequest{Request: req, Input: input, Copy: c.QueryForecastRequest}
}

// QueryForecastRequest is the request type for the
// QueryForecast API operation.
type QueryForecastRequest struct {
	*aws.Request
	Input *QueryForecastInput
	Copy  func(*QueryForecastInput) QueryForecastRequest
}

// Send marshals and sends the QueryForecast API request.
func (r QueryForecastRequest) Send(ctx context.Context) (*QueryForecastResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &QueryForecastResponse{
		QueryForecastOutput: r.Request.Data.(*QueryForecastOutput),
		response:            &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// QueryForecastResponse is the response type for the
// QueryForecast API operation.
type QueryForecastResponse struct {
	*QueryForecastOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// QueryForecast request.
func (r *QueryForecastResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
