// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package kinesis

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
)

// Represents the input for DescribeStream.
type DescribeStreamInput struct {
	_ struct{} `type:"structure"`

	// The shard ID of the shard to start with.
	ExclusiveStartShardId *string `min:"1" type:"string"`

	// The maximum number of shards to return in a single call. The default value
	// is 100. If you specify a value greater than 100, at most 100 shards are returned.
	Limit *int64 `min:"1" type:"integer"`

	// The name of the stream to describe.
	//
	// StreamName is a required field
	StreamName *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s DescribeStreamInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeStreamInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeStreamInput"}
	if s.ExclusiveStartShardId != nil && len(*s.ExclusiveStartShardId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ExclusiveStartShardId", 1))
	}
	if s.Limit != nil && *s.Limit < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("Limit", 1))
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

// Represents the output for DescribeStream.
type DescribeStreamOutput struct {
	_ struct{} `type:"structure"`

	// The current status of the stream, the stream Amazon Resource Name (ARN),
	// an array of shard objects that comprise the stream, and whether there are
	// more shards available.
	//
	// StreamDescription is a required field
	StreamDescription *StreamDescription `type:"structure" required:"true"`
}

// String returns the string representation
func (s DescribeStreamOutput) String() string {
	return awsutil.Prettify(s)
}

const opDescribeStream = "DescribeStream"

// DescribeStreamRequest returns a request value for making API operation for
// Amazon Kinesis.
//
// Describes the specified Kinesis data stream.
//
// The information returned includes the stream name, Amazon Resource Name (ARN),
// creation time, enhanced metric configuration, and shard map. The shard map
// is an array of shard objects. For each shard object, there is the hash key
// and sequence number ranges that the shard spans, and the IDs of any earlier
// shards that played in a role in creating the shard. Every record ingested
// in the stream is identified by a sequence number, which is assigned when
// the record is put into the stream.
//
// You can limit the number of shards returned by each call. For more information,
// see Retrieving Shards from a Stream (http://docs.aws.amazon.com/kinesis/latest/dev/kinesis-using-sdk-java-retrieve-shards.html)
// in the Amazon Kinesis Data Streams Developer Guide.
//
// There are no guarantees about the chronological order shards returned. To
// process shards in chronological order, use the ID of the parent shard to
// track the lineage to the oldest shard.
//
// This operation has a limit of 10 transactions per second per account.
//
//    // Example sending a request using DescribeStreamRequest.
//    req := client.DescribeStreamRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/kinesis-2013-12-02/DescribeStream
func (c *Client) DescribeStreamRequest(input *DescribeStreamInput) DescribeStreamRequest {
	op := &aws.Operation{
		Name:       opDescribeStream,
		HTTPMethod: "POST",
		HTTPPath:   "/",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"ExclusiveStartShardId"},
			OutputTokens:    []string{"StreamDescription.Shards[-1].ShardId"},
			LimitToken:      "Limit",
			TruncationToken: "StreamDescription.HasMoreShards",
		},
	}

	if input == nil {
		input = &DescribeStreamInput{}
	}

	req := c.newRequest(op, input, &DescribeStreamOutput{})

	return DescribeStreamRequest{Request: req, Input: input, Copy: c.DescribeStreamRequest}
}

// DescribeStreamRequest is the request type for the
// DescribeStream API operation.
type DescribeStreamRequest struct {
	*aws.Request
	Input *DescribeStreamInput
	Copy  func(*DescribeStreamInput) DescribeStreamRequest
}

// Send marshals and sends the DescribeStream API request.
func (r DescribeStreamRequest) Send(ctx context.Context) (*DescribeStreamResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeStreamResponse{
		DescribeStreamOutput: r.Request.Data.(*DescribeStreamOutput),
		response:             &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewDescribeStreamRequestPaginator returns a paginator for DescribeStream.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.DescribeStreamRequest(input)
//   p := kinesis.NewDescribeStreamRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewDescribeStreamPaginator(req DescribeStreamRequest) DescribeStreamPaginator {
	return DescribeStreamPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *DescribeStreamInput
				if req.Input != nil {
					tmp := *req.Input
					inCpy = &tmp
				}

				newReq := req.Copy(inCpy)
				newReq.SetContext(ctx)
				return newReq.Request, nil
			},
		},
	}
}

// DescribeStreamPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type DescribeStreamPaginator struct {
	aws.Pager
}

func (p *DescribeStreamPaginator) CurrentPage() *DescribeStreamOutput {
	return p.Pager.CurrentPage().(*DescribeStreamOutput)
}

// DescribeStreamResponse is the response type for the
// DescribeStream API operation.
type DescribeStreamResponse struct {
	*DescribeStreamOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeStream request.
func (r *DescribeStreamResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
