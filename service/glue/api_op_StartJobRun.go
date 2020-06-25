// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package glue

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
)

type StartJobRunInput struct {
	_ struct{} `type:"structure"`

	// This field is deprecated. Use MaxCapacity instead.
	//
	// The number of AWS Glue data processing units (DPUs) to allocate to this JobRun.
	// From 2 to 100 DPUs can be allocated; the default is 10. A DPU is a relative
	// measure of processing power that consists of 4 vCPUs of compute capacity
	// and 16 GB of memory. For more information, see the AWS Glue pricing page
	// (https://docs.aws.amazon.com/https:/aws.amazon.com/glue/pricing/).
	AllocatedCapacity *int64 `deprecated:"true" type:"integer"`

	// The job arguments specifically for this run. For this job run, they replace
	// the default arguments set in the job definition itself.
	//
	// You can specify arguments here that your own job-execution script consumes,
	// as well as arguments that AWS Glue itself consumes.
	//
	// For information about how to specify and consume your own Job arguments,
	// see the Calling AWS Glue APIs in Python (https://docs.aws.amazon.com/glue/latest/dg/aws-glue-programming-python-calling.html)
	// topic in the developer guide.
	//
	// For information about the key-value pairs that AWS Glue consumes to set up
	// your job, see the Special Parameters Used by AWS Glue (https://docs.aws.amazon.com/glue/latest/dg/aws-glue-programming-etl-glue-arguments.html)
	// topic in the developer guide.
	Arguments map[string]string `type:"map"`

	// The name of the job definition to use.
	//
	// JobName is a required field
	JobName *string `min:"1" type:"string" required:"true"`

	// The ID of a previous JobRun to retry.
	JobRunId *string `min:"1" type:"string"`

	// The number of AWS Glue data processing units (DPUs) that can be allocated
	// when this job runs. A DPU is a relative measure of processing power that
	// consists of 4 vCPUs of compute capacity and 16 GB of memory. For more information,
	// see the AWS Glue pricing page (https://docs.aws.amazon.com/https:/aws.amazon.com/glue/pricing/).
	//
	// Do not set Max Capacity if using WorkerType and NumberOfWorkers.
	//
	// The value that can be allocated for MaxCapacity depends on whether you are
	// running a Python shell job, or an Apache Spark ETL job:
	//
	//    * When you specify a Python shell job (JobCommand.Name="pythonshell"),
	//    you can allocate either 0.0625 or 1 DPU. The default is 0.0625 DPU.
	//
	//    * When you specify an Apache Spark ETL job (JobCommand.Name="glueetl"),
	//    you can allocate from 2 to 100 DPUs. The default is 10 DPUs. This job
	//    type cannot have a fractional DPU allocation.
	MaxCapacity *float64 `type:"double"`

	// Specifies configuration properties of a job run notification.
	NotificationProperty *NotificationProperty `type:"structure"`

	// The number of workers of a defined workerType that are allocated when a job
	// runs.
	//
	// The maximum number of workers you can define are 299 for G.1X, and 149 for
	// G.2X.
	NumberOfWorkers *int64 `type:"integer"`

	// The name of the SecurityConfiguration structure to be used with this job
	// run.
	SecurityConfiguration *string `min:"1" type:"string"`

	// The JobRun timeout in minutes. This is the maximum time that a job run can
	// consume resources before it is terminated and enters TIMEOUT status. The
	// default is 2,880 minutes (48 hours). This overrides the timeout value set
	// in the parent job.
	Timeout *int64 `min:"1" type:"integer"`

	// The type of predefined worker that is allocated when a job runs. Accepts
	// a value of Standard, G.1X, or G.2X.
	//
	//    * For the Standard worker type, each worker provides 4 vCPU, 16 GB of
	//    memory and a 50GB disk, and 2 executors per worker.
	//
	//    * For the G.1X worker type, each worker provides 4 vCPU, 16 GB of memory
	//    and a 64GB disk, and 1 executor per worker.
	//
	//    * For the G.2X worker type, each worker provides 8 vCPU, 32 GB of memory
	//    and a 128GB disk, and 1 executor per worker.
	WorkerType WorkerType `type:"string" enum:"true"`
}

// String returns the string representation
func (s StartJobRunInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *StartJobRunInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "StartJobRunInput"}

	if s.JobName == nil {
		invalidParams.Add(aws.NewErrParamRequired("JobName"))
	}
	if s.JobName != nil && len(*s.JobName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("JobName", 1))
	}
	if s.JobRunId != nil && len(*s.JobRunId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("JobRunId", 1))
	}
	if s.SecurityConfiguration != nil && len(*s.SecurityConfiguration) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("SecurityConfiguration", 1))
	}
	if s.Timeout != nil && *s.Timeout < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("Timeout", 1))
	}
	if s.NotificationProperty != nil {
		if err := s.NotificationProperty.Validate(); err != nil {
			invalidParams.AddNested("NotificationProperty", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type StartJobRunOutput struct {
	_ struct{} `type:"structure"`

	// The ID assigned to this job run.
	JobRunId *string `min:"1" type:"string"`
}

// String returns the string representation
func (s StartJobRunOutput) String() string {
	return awsutil.Prettify(s)
}

const opStartJobRun = "StartJobRun"

// StartJobRunRequest returns a request value for making API operation for
// AWS Glue.
//
// Starts a job run using a job definition.
//
//    // Example sending a request using StartJobRunRequest.
//    req := client.StartJobRunRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/glue-2017-03-31/StartJobRun
func (c *Client) StartJobRunRequest(input *StartJobRunInput) StartJobRunRequest {
	op := &aws.Operation{
		Name:       opStartJobRun,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &StartJobRunInput{}
	}

	req := c.newRequest(op, input, &StartJobRunOutput{})

	return StartJobRunRequest{Request: req, Input: input, Copy: c.StartJobRunRequest}
}

// StartJobRunRequest is the request type for the
// StartJobRun API operation.
type StartJobRunRequest struct {
	*aws.Request
	Input *StartJobRunInput
	Copy  func(*StartJobRunInput) StartJobRunRequest
}

// Send marshals and sends the StartJobRun API request.
func (r StartJobRunRequest) Send(ctx context.Context) (*StartJobRunResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &StartJobRunResponse{
		StartJobRunOutput: r.Request.Data.(*StartJobRunOutput),
		response:          &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// StartJobRunResponse is the response type for the
// StartJobRun API operation.
type StartJobRunResponse struct {
	*StartJobRunOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// StartJobRun request.
func (r *StartJobRunResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
