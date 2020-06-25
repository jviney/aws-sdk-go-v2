// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package machinelearning

import (
	"context"
	"time"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
)

type GetEvaluationInput struct {
	_ struct{} `type:"structure"`

	// The ID of the Evaluation to retrieve. The evaluation of each MLModel is recorded
	// and cataloged. The ID provides the means to access the information.
	//
	// EvaluationId is a required field
	EvaluationId *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s GetEvaluationInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetEvaluationInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetEvaluationInput"}

	if s.EvaluationId == nil {
		invalidParams.Add(aws.NewErrParamRequired("EvaluationId"))
	}
	if s.EvaluationId != nil && len(*s.EvaluationId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("EvaluationId", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Represents the output of a GetEvaluation operation and describes an Evaluation.
type GetEvaluationOutput struct {
	_ struct{} `type:"structure"`

	// The approximate CPU time in milliseconds that Amazon Machine Learning spent
	// processing the Evaluation, normalized and scaled on computation resources.
	// ComputeTime is only available if the Evaluation is in the COMPLETED state.
	ComputeTime *int64 `type:"long"`

	// The time that the Evaluation was created. The time is expressed in epoch
	// time.
	CreatedAt *time.Time `type:"timestamp"`

	// The AWS user account that invoked the evaluation. The account type can be
	// either an AWS root account or an AWS Identity and Access Management (IAM)
	// user account.
	CreatedByIamUser *string `type:"string"`

	// The DataSource used for this evaluation.
	EvaluationDataSourceId *string `min:"1" type:"string"`

	// The evaluation ID which is same as the EvaluationId in the request.
	EvaluationId *string `min:"1" type:"string"`

	// The epoch time when Amazon Machine Learning marked the Evaluation as COMPLETED
	// or FAILED. FinishedAt is only available when the Evaluation is in the COMPLETED
	// or FAILED state.
	FinishedAt *time.Time `type:"timestamp"`

	// The location of the data file or directory in Amazon Simple Storage Service
	// (Amazon S3).
	InputDataLocationS3 *string `type:"string"`

	// The time of the most recent edit to the Evaluation. The time is expressed
	// in epoch time.
	LastUpdatedAt *time.Time `type:"timestamp"`

	// A link to the file that contains logs of the CreateEvaluation operation.
	LogUri *string `type:"string"`

	// The ID of the MLModel that was the focus of the evaluation.
	MLModelId *string `min:"1" type:"string"`

	// A description of the most recent details about evaluating the MLModel.
	Message *string `type:"string"`

	// A user-supplied name or description of the Evaluation.
	Name *string `type:"string"`

	// Measurements of how well the MLModel performed using observations referenced
	// by the DataSource. One of the following metric is returned based on the type
	// of the MLModel:
	//
	//    * BinaryAUC: A binary MLModel uses the Area Under the Curve (AUC) technique
	//    to measure performance.
	//
	//    * RegressionRMSE: A regression MLModel uses the Root Mean Square Error
	//    (RMSE) technique to measure performance. RMSE measures the difference
	//    between predicted and actual values for a single variable.
	//
	//    * MulticlassAvgFScore: A multiclass MLModel uses the F1 score technique
	//    to measure performance.
	//
	// For more information about performance metrics, please see the Amazon Machine
	// Learning Developer Guide (http://docs.aws.amazon.com/machine-learning/latest/dg).
	PerformanceMetrics *PerformanceMetrics `type:"structure"`

	// The epoch time when Amazon Machine Learning marked the Evaluation as INPROGRESS.
	// StartedAt isn't available if the Evaluation is in the PENDING state.
	StartedAt *time.Time `type:"timestamp"`

	// The status of the evaluation. This element can have one of the following
	// values:
	//
	//    * PENDING - Amazon Machine Language (Amazon ML) submitted a request to
	//    evaluate an MLModel.
	//
	//    * INPROGRESS - The evaluation is underway.
	//
	//    * FAILED - The request to evaluate an MLModel did not run to completion.
	//    It is not usable.
	//
	//    * COMPLETED - The evaluation process completed successfully.
	//
	//    * DELETED - The Evaluation is marked as deleted. It is not usable.
	Status EntityStatus `type:"string" enum:"true"`
}

// String returns the string representation
func (s GetEvaluationOutput) String() string {
	return awsutil.Prettify(s)
}

const opGetEvaluation = "GetEvaluation"

// GetEvaluationRequest returns a request value for making API operation for
// Amazon Machine Learning.
//
// Returns an Evaluation that includes metadata as well as the current status
// of the Evaluation.
//
//    // Example sending a request using GetEvaluationRequest.
//    req := client.GetEvaluationRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
func (c *Client) GetEvaluationRequest(input *GetEvaluationInput) GetEvaluationRequest {
	op := &aws.Operation{
		Name:       opGetEvaluation,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &GetEvaluationInput{}
	}

	req := c.newRequest(op, input, &GetEvaluationOutput{})

	return GetEvaluationRequest{Request: req, Input: input, Copy: c.GetEvaluationRequest}
}

// GetEvaluationRequest is the request type for the
// GetEvaluation API operation.
type GetEvaluationRequest struct {
	*aws.Request
	Input *GetEvaluationInput
	Copy  func(*GetEvaluationInput) GetEvaluationRequest
}

// Send marshals and sends the GetEvaluation API request.
func (r GetEvaluationRequest) Send(ctx context.Context) (*GetEvaluationResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetEvaluationResponse{
		GetEvaluationOutput: r.Request.Data.(*GetEvaluationOutput),
		response:            &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetEvaluationResponse is the response type for the
// GetEvaluation API operation.
type GetEvaluationResponse struct {
	*GetEvaluationOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetEvaluation request.
func (r *GetEvaluationResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
