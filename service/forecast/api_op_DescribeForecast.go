// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package forecast

import (
	"context"
	"time"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
)

type DescribeForecastInput struct {
	_ struct{} `type:"structure"`

	// The Amazon Resource Name (ARN) of the forecast.
	//
	// ForecastArn is a required field
	ForecastArn *string `type:"string" required:"true"`
}

// String returns the string representation
func (s DescribeForecastInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeForecastInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeForecastInput"}

	if s.ForecastArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("ForecastArn"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DescribeForecastOutput struct {
	_ struct{} `type:"structure"`

	// When the forecast creation task was created.
	CreationTime *time.Time `type:"timestamp"`

	// The ARN of the dataset group that provided the data used to train the predictor.
	DatasetGroupArn *string `type:"string"`

	// The forecast ARN as specified in the request.
	ForecastArn *string `type:"string"`

	// The name of the forecast.
	ForecastName *string `min:"1" type:"string"`

	// The quantiles at which proababilistic forecasts were generated.
	ForecastTypes []string `min:"1" type:"list"`

	// Initially, the same as CreationTime (status is CREATE_PENDING). Updated when
	// inference (creating the forecast) starts (status changed to CREATE_IN_PROGRESS),
	// and when inference is complete (status changed to ACTIVE) or fails (status
	// changed to CREATE_FAILED).
	LastModificationTime *time.Time `type:"timestamp"`

	// If an error occurred, an informational message about the error.
	Message *string `type:"string"`

	// The ARN of the predictor used to generate the forecast.
	PredictorArn *string `type:"string"`

	// The status of the forecast. States include:
	//
	//    * ACTIVE
	//
	//    * CREATE_PENDING, CREATE_IN_PROGRESS, CREATE_FAILED
	//
	//    * DELETE_PENDING, DELETE_IN_PROGRESS, DELETE_FAILED
	//
	// The Status of the forecast must be ACTIVE before you can query or export
	// the forecast.
	Status *string `type:"string"`
}

// String returns the string representation
func (s DescribeForecastOutput) String() string {
	return awsutil.Prettify(s)
}

const opDescribeForecast = "DescribeForecast"

// DescribeForecastRequest returns a request value for making API operation for
// Amazon Forecast Service.
//
// Describes a forecast created using the CreateForecast operation.
//
// In addition to listing the properties provided in the CreateForecast request,
// this operation lists the following properties:
//
//    * DatasetGroupArn - The dataset group that provided the training data.
//
//    * CreationTime
//
//    * LastModificationTime
//
//    * Status
//
//    * Message - If an error occurred, information about the error.
//
//    // Example sending a request using DescribeForecastRequest.
//    req := client.DescribeForecastRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/forecast-2018-06-26/DescribeForecast
func (c *Client) DescribeForecastRequest(input *DescribeForecastInput) DescribeForecastRequest {
	op := &aws.Operation{
		Name:       opDescribeForecast,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeForecastInput{}
	}

	req := c.newRequest(op, input, &DescribeForecastOutput{})

	return DescribeForecastRequest{Request: req, Input: input, Copy: c.DescribeForecastRequest}
}

// DescribeForecastRequest is the request type for the
// DescribeForecast API operation.
type DescribeForecastRequest struct {
	*aws.Request
	Input *DescribeForecastInput
	Copy  func(*DescribeForecastInput) DescribeForecastRequest
}

// Send marshals and sends the DescribeForecast API request.
func (r DescribeForecastRequest) Send(ctx context.Context) (*DescribeForecastResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeForecastResponse{
		DescribeForecastOutput: r.Request.Data.(*DescribeForecastOutput),
		response:               &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeForecastResponse is the response type for the
// DescribeForecast API operation.
type DescribeForecastResponse struct {
	*DescribeForecastOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeForecast request.
func (r *DescribeForecastResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
