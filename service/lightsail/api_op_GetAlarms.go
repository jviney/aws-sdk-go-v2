// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package lightsail

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
)

type GetAlarmsInput struct {
	_ struct{} `type:"structure"`

	// The name of the alarm.
	//
	// Specify an alarm name to return information about a specific alarm.
	AlarmName *string `locationName:"alarmName" type:"string"`

	// The name of the Lightsail resource being monitored by the alarm.
	//
	// Specify a monitored resource name to return information about all alarms
	// for a specific resource.
	MonitoredResourceName *string `locationName:"monitoredResourceName" type:"string"`

	// The token to advance to the next page of results from your request.
	//
	// To get a page token, perform an initial GetAlarms request. If your results
	// are paginated, the response will return a next page token that you can specify
	// as the page token in a subsequent request.
	PageToken *string `locationName:"pageToken" type:"string"`
}

// String returns the string representation
func (s GetAlarmsInput) String() string {
	return awsutil.Prettify(s)
}

type GetAlarmsOutput struct {
	_ struct{} `type:"structure"`

	// An array of objects that describe the alarms.
	Alarms []Alarm `locationName:"alarms" type:"list"`

	// The token to advance to the next page of resutls from your request.
	//
	// A next page token is not returned if there are no more results to display.
	//
	// To get the next page of results, perform another GetAlarms request and specify
	// the next page token using the pageToken parameter.
	NextPageToken *string `locationName:"nextPageToken" type:"string"`
}

// String returns the string representation
func (s GetAlarmsOutput) String() string {
	return awsutil.Prettify(s)
}

const opGetAlarms = "GetAlarms"

// GetAlarmsRequest returns a request value for making API operation for
// Amazon Lightsail.
//
// Returns information about the configured alarms. Specify an alarm name in
// your request to return information about a specific alarm, or specify a monitored
// resource name to return information about all alarms for a specific resource.
//
// An alarm is used to monitor a single metric for one of your resources. When
// a metric condition is met, the alarm can notify you by email, SMS text message,
// and a banner displayed on the Amazon Lightsail console. For more information,
// see Alarms in Amazon Lightsail (https://lightsail.aws.amazon.com/ls/docs/en_us/articles/amazon-lightsail-alarms).
//
//    // Example sending a request using GetAlarmsRequest.
//    req := client.GetAlarmsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/lightsail-2016-11-28/GetAlarms
func (c *Client) GetAlarmsRequest(input *GetAlarmsInput) GetAlarmsRequest {
	op := &aws.Operation{
		Name:       opGetAlarms,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &GetAlarmsInput{}
	}

	req := c.newRequest(op, input, &GetAlarmsOutput{})

	return GetAlarmsRequest{Request: req, Input: input, Copy: c.GetAlarmsRequest}
}

// GetAlarmsRequest is the request type for the
// GetAlarms API operation.
type GetAlarmsRequest struct {
	*aws.Request
	Input *GetAlarmsInput
	Copy  func(*GetAlarmsInput) GetAlarmsRequest
}

// Send marshals and sends the GetAlarms API request.
func (r GetAlarmsRequest) Send(ctx context.Context) (*GetAlarmsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetAlarmsResponse{
		GetAlarmsOutput: r.Request.Data.(*GetAlarmsOutput),
		response:        &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetAlarmsResponse is the response type for the
// GetAlarms API operation.
type GetAlarmsResponse struct {
	*GetAlarmsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetAlarms request.
func (r *GetAlarmsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
