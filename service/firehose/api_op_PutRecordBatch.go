// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package firehose

import (
	"context"
	"fmt"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
)

type PutRecordBatchInput struct {
	_ struct{} `type:"structure"`

	// The name of the delivery stream.
	//
	// DeliveryStreamName is a required field
	DeliveryStreamName *string `min:"1" type:"string" required:"true"`

	// One or more records.
	//
	// Records is a required field
	Records []Record `min:"1" type:"list" required:"true"`
}

// String returns the string representation
func (s PutRecordBatchInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *PutRecordBatchInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "PutRecordBatchInput"}

	if s.DeliveryStreamName == nil {
		invalidParams.Add(aws.NewErrParamRequired("DeliveryStreamName"))
	}
	if s.DeliveryStreamName != nil && len(*s.DeliveryStreamName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("DeliveryStreamName", 1))
	}

	if s.Records == nil {
		invalidParams.Add(aws.NewErrParamRequired("Records"))
	}
	if s.Records != nil && len(s.Records) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Records", 1))
	}
	if s.Records != nil {
		for i, v := range s.Records {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "Records", i), err.(aws.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type PutRecordBatchOutput struct {
	_ struct{} `type:"structure"`

	// Indicates whether server-side encryption (SSE) was enabled during this operation.
	Encrypted *bool `type:"boolean"`

	// The number of records that might have failed processing. This number might
	// be greater than 0 even if the PutRecordBatch call succeeds. Check FailedPutCount
	// to determine whether there are records that you need to resend.
	//
	// FailedPutCount is a required field
	FailedPutCount *int64 `type:"integer" required:"true"`

	// The results array. For each record, the index of the response element is
	// the same as the index used in the request array.
	//
	// RequestResponses is a required field
	RequestResponses []PutRecordBatchResponseEntry `min:"1" type:"list" required:"true"`
}

// String returns the string representation
func (s PutRecordBatchOutput) String() string {
	return awsutil.Prettify(s)
}

const opPutRecordBatch = "PutRecordBatch"

// PutRecordBatchRequest returns a request value for making API operation for
// Amazon Kinesis Firehose.
//
// Writes multiple data records into a delivery stream in a single call, which
// can achieve higher throughput per producer than when writing single records.
// To write single data records into a delivery stream, use PutRecord. Applications
// using these operations are referred to as producers.
//
// By default, each delivery stream can take in up to 2,000 transactions per
// second, 5,000 records per second, or 5 MB per second. If you use PutRecord
// and PutRecordBatch, the limits are an aggregate across these two operations
// for each delivery stream. For more information about limits, see Amazon Kinesis
// Data Firehose Limits (https://docs.aws.amazon.com/firehose/latest/dev/limits.html).
//
// Each PutRecordBatch request supports up to 500 records. Each record in the
// request can be as large as 1,000 KB (before 64-bit encoding), up to a limit
// of 4 MB for the entire request. These limits cannot be changed.
//
// You must specify the name of the delivery stream and the data record when
// using PutRecord. The data record consists of a data blob that can be up to
// 1,000 KB in size, and any kind of data. For example, it could be a segment
// from a log file, geographic location data, website clickstream data, and
// so on.
//
// Kinesis Data Firehose buffers records before delivering them to the destination.
// To disambiguate the data blobs at the destination, a common solution is to
// use delimiters in the data, such as a newline (\n) or some other character
// unique within the data. This allows the consumer application to parse individual
// data items when reading the data from the destination.
//
// The PutRecordBatch response includes a count of failed records, FailedPutCount,
// and an array of responses, RequestResponses. Even if the PutRecordBatch call
// succeeds, the value of FailedPutCount may be greater than 0, indicating that
// there are records for which the operation didn't succeed. Each entry in the
// RequestResponses array provides additional information about the processed
// record. It directly correlates with a record in the request array using the
// same ordering, from the top to the bottom. The response array always includes
// the same number of records as the request array. RequestResponses includes
// both successfully and unsuccessfully processed records. Kinesis Data Firehose
// tries to process all records in each PutRecordBatch request. A single record
// failure does not stop the processing of subsequent records.
//
// A successfully processed record includes a RecordId value, which is unique
// for the record. An unsuccessfully processed record includes ErrorCode and
// ErrorMessage values. ErrorCode reflects the type of error, and is one of
// the following values: ServiceUnavailableException or InternalFailure. ErrorMessage
// provides more detailed information about the error.
//
// If there is an internal server error or a timeout, the write might have completed
// or it might have failed. If FailedPutCount is greater than 0, retry the request,
// resending only those records that might have failed processing. This minimizes
// the possible duplicate records and also reduces the total bytes sent (and
// corresponding charges). We recommend that you handle any duplicates at the
// destination.
//
// If PutRecordBatch throws ServiceUnavailableException, back off and retry.
// If the exception persists, it is possible that the throughput limits have
// been exceeded for the delivery stream.
//
// Data records sent to Kinesis Data Firehose are stored for 24 hours from the
// time they are added to a delivery stream as it attempts to send the records
// to the destination. If the destination is unreachable for more than 24 hours,
// the data is no longer available.
//
// Don't concatenate two or more base64 strings to form the data fields of your
// records. Instead, concatenate the raw data, then perform base64 encoding.
//
//    // Example sending a request using PutRecordBatchRequest.
//    req := client.PutRecordBatchRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/firehose-2015-08-04/PutRecordBatch
func (c *Client) PutRecordBatchRequest(input *PutRecordBatchInput) PutRecordBatchRequest {
	op := &aws.Operation{
		Name:       opPutRecordBatch,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &PutRecordBatchInput{}
	}

	req := c.newRequest(op, input, &PutRecordBatchOutput{})

	return PutRecordBatchRequest{Request: req, Input: input, Copy: c.PutRecordBatchRequest}
}

// PutRecordBatchRequest is the request type for the
// PutRecordBatch API operation.
type PutRecordBatchRequest struct {
	*aws.Request
	Input *PutRecordBatchInput
	Copy  func(*PutRecordBatchInput) PutRecordBatchRequest
}

// Send marshals and sends the PutRecordBatch API request.
func (r PutRecordBatchRequest) Send(ctx context.Context) (*PutRecordBatchResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &PutRecordBatchResponse{
		PutRecordBatchOutput: r.Request.Data.(*PutRecordBatchOutput),
		response:             &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// PutRecordBatchResponse is the response type for the
// PutRecordBatch API operation.
type PutRecordBatchResponse struct {
	*PutRecordBatchOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// PutRecordBatch request.
func (r *PutRecordBatchResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
