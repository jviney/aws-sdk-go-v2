// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package kinesis

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
)

// Represents the input for PutRecord.
type PutRecordInput struct {
	_ struct{} `type:"structure"`

	// The data blob to put into the record, which is base64-encoded when the blob
	// is serialized. When the data blob (the payload before base64-encoding) is
	// added to the partition key size, the total size must not exceed the maximum
	// record size (1 MB).
	//
	// Data is automatically base64 encoded/decoded by the SDK.
	//
	// Data is a required field
	Data []byte `type:"blob" required:"true"`

	// The hash value used to explicitly determine the shard the data record is
	// assigned to by overriding the partition key hash.
	ExplicitHashKey *string `type:"string"`

	// Determines which shard in the stream the data record is assigned to. Partition
	// keys are Unicode strings with a maximum length limit of 256 characters for
	// each key. Amazon Kinesis Data Streams uses the partition key as input to
	// a hash function that maps the partition key and associated data to a specific
	// shard. Specifically, an MD5 hash function is used to map partition keys to
	// 128-bit integer values and to map associated data records to shards. As a
	// result of this hashing mechanism, all data records with the same partition
	// key map to the same shard within the stream.
	//
	// PartitionKey is a required field
	PartitionKey *string `min:"1" type:"string" required:"true"`

	// Guarantees strictly increasing sequence numbers, for puts from the same client
	// and to the same partition key. Usage: set the SequenceNumberForOrdering of
	// record n to the sequence number of record n-1 (as returned in the result
	// when putting record n-1). If this parameter is not set, records are coarsely
	// ordered based on arrival time.
	SequenceNumberForOrdering *string `type:"string"`

	// The name of the stream to put the data record into.
	//
	// StreamName is a required field
	StreamName *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s PutRecordInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *PutRecordInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "PutRecordInput"}

	if s.Data == nil {
		invalidParams.Add(aws.NewErrParamRequired("Data"))
	}

	if s.PartitionKey == nil {
		invalidParams.Add(aws.NewErrParamRequired("PartitionKey"))
	}
	if s.PartitionKey != nil && len(*s.PartitionKey) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("PartitionKey", 1))
	}

	if s.StreamName == nil {
		invalidParams.Add(aws.NewErrParamRequired("StreamName"))
	}
	if s.StreamName != nil && len(*s.StreamName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("StreamName", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Represents the output for PutRecord.
type PutRecordOutput struct {
	_ struct{} `type:"structure"`

	// The encryption type to use on the record. This parameter can be one of the
	// following values:
	//
	//    * NONE: Do not encrypt the records in the stream.
	//
	//    * KMS: Use server-side encryption on the records in the stream using a
	//    customer-managed AWS KMS key.
	EncryptionType EncryptionType `type:"string" enum:"true"`

	// The sequence number identifier that was assigned to the put data record.
	// The sequence number for the record is unique across all records in the stream.
	// A sequence number is the identifier associated with every record put into
	// the stream.
	//
	// SequenceNumber is a required field
	SequenceNumber *string `type:"string" required:"true"`

	// The shard ID of the shard where the data record was placed.
	//
	// ShardId is a required field
	ShardId *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s PutRecordOutput) String() string {
	return awsutil.Prettify(s)
}

const opPutRecord = "PutRecord"

// PutRecordRequest returns a request value for making API operation for
// Amazon Kinesis.
//
// Writes a single data record into an Amazon Kinesis data stream. Call PutRecord
// to send data into the stream for real-time ingestion and subsequent processing,
// one record at a time. Each shard can support writes up to 1,000 records per
// second, up to a maximum data write total of 1 MB per second.
//
// You must specify the name of the stream that captures, stores, and transports
// the data; a partition key; and the data blob itself.
//
// The data blob can be any type of data; for example, a segment from a log
// file, geographic/location data, website clickstream data, and so on.
//
// The partition key is used by Kinesis Data Streams to distribute data across
// shards. Kinesis Data Streams segregates the data records that belong to a
// stream into multiple shards, using the partition key associated with each
// data record to determine the shard to which a given data record belongs.
//
// Partition keys are Unicode strings, with a maximum length limit of 256 characters
// for each key. An MD5 hash function is used to map partition keys to 128-bit
// integer values and to map associated data records to shards using the hash
// key ranges of the shards. You can override hashing the partition key to determine
// the shard by explicitly specifying a hash value using the ExplicitHashKey
// parameter. For more information, see Adding Data to a Stream (http://docs.aws.amazon.com/kinesis/latest/dev/developing-producers-with-sdk.html#kinesis-using-sdk-java-add-data-to-stream)
// in the Amazon Kinesis Data Streams Developer Guide.
//
// PutRecord returns the shard ID of where the data record was placed and the
// sequence number that was assigned to the data record.
//
// Sequence numbers increase over time and are specific to a shard within a
// stream, not across all shards within a stream. To guarantee strictly increasing
// ordering, write serially to a shard and use the SequenceNumberForOrdering
// parameter. For more information, see Adding Data to a Stream (http://docs.aws.amazon.com/kinesis/latest/dev/developing-producers-with-sdk.html#kinesis-using-sdk-java-add-data-to-stream)
// in the Amazon Kinesis Data Streams Developer Guide.
//
// If a PutRecord request cannot be processed because of insufficient provisioned
// throughput on the shard involved in the request, PutRecord throws ProvisionedThroughputExceededException.
//
// By default, data records are accessible for 24 hours from the time that they
// are added to a stream. You can use IncreaseStreamRetentionPeriod or DecreaseStreamRetentionPeriod
// to modify this retention period.
//
//    // Example sending a request using PutRecordRequest.
//    req := client.PutRecordRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/kinesis-2013-12-02/PutRecord
func (c *Client) PutRecordRequest(input *PutRecordInput) PutRecordRequest {
	op := &aws.Operation{
		Name:       opPutRecord,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &PutRecordInput{}
	}

	req := c.newRequest(op, input, &PutRecordOutput{})

	return PutRecordRequest{Request: req, Input: input, Copy: c.PutRecordRequest}
}

// PutRecordRequest is the request type for the
// PutRecord API operation.
type PutRecordRequest struct {
	*aws.Request
	Input *PutRecordInput
	Copy  func(*PutRecordInput) PutRecordRequest
}

// Send marshals and sends the PutRecord API request.
func (r PutRecordRequest) Send(ctx context.Context) (*PutRecordResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &PutRecordResponse{
		PutRecordOutput: r.Request.Data.(*PutRecordOutput),
		response:        &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// PutRecordResponse is the response type for the
// PutRecord API operation.
type PutRecordResponse struct {
	*PutRecordOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// PutRecord request.
func (r *PutRecordResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
