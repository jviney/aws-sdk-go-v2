// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package applicationautoscaling

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
)

type DeleteScheduledActionInput struct {
	_ struct{} `type:"structure"`

	// The identifier of the resource associated with the scheduled action. This
	// string consists of the resource type and unique identifier.
	//
	//    * ECS service - The resource type is service and the unique identifier
	//    is the cluster name and service name. Example: service/default/sample-webapp.
	//
	//    * Spot Fleet request - The resource type is spot-fleet-request and the
	//    unique identifier is the Spot Fleet request ID. Example: spot-fleet-request/sfr-73fbd2ce-aa30-494c-8788-1cee4EXAMPLE.
	//
	//    * EMR cluster - The resource type is instancegroup and the unique identifier
	//    is the cluster ID and instance group ID. Example: instancegroup/j-2EEZNYKUA1NTV/ig-1791Y4E1L8YI0.
	//
	//    * AppStream 2.0 fleet - The resource type is fleet and the unique identifier
	//    is the fleet name. Example: fleet/sample-fleet.
	//
	//    * DynamoDB table - The resource type is table and the unique identifier
	//    is the table name. Example: table/my-table.
	//
	//    * DynamoDB global secondary index - The resource type is index and the
	//    unique identifier is the index name. Example: table/my-table/index/my-table-index.
	//
	//    * Aurora DB cluster - The resource type is cluster and the unique identifier
	//    is the cluster name. Example: cluster:my-db-cluster.
	//
	//    * Amazon SageMaker endpoint variant - The resource type is variant and
	//    the unique identifier is the resource ID. Example: endpoint/my-end-point/variant/KMeansClustering.
	//
	//    * Custom resources are not supported with a resource type. This parameter
	//    must specify the OutputValue from the CloudFormation template stack used
	//    to access the resources. The unique identifier is defined by the service
	//    provider. More information is available in our GitHub repository (https://github.com/aws/aws-auto-scaling-custom-resource).
	//
	//    * Amazon Comprehend document classification endpoint - The resource type
	//    and unique identifier are specified using the endpoint ARN. Example: arn:aws:comprehend:us-west-2:123456789012:document-classifier-endpoint/EXAMPLE.
	//
	//    * Lambda provisioned concurrency - The resource type is function and the
	//    unique identifier is the function name with a function version or alias
	//    name suffix that is not $LATEST. Example: function:my-function:prod or
	//    function:my-function:1.
	//
	//    * Amazon Keyspaces table - The resource type is table and the unique identifier
	//    is the table name. Example: keyspace/mykeyspace/table/mytable.
	//
	// ResourceId is a required field
	ResourceId *string `min:"1" type:"string" required:"true"`

	// The scalable dimension. This string consists of the service namespace, resource
	// type, and scaling property.
	//
	//    * ecs:service:DesiredCount - The desired task count of an ECS service.
	//
	//    * ec2:spot-fleet-request:TargetCapacity - The target capacity of a Spot
	//    Fleet request.
	//
	//    * elasticmapreduce:instancegroup:InstanceCount - The instance count of
	//    an EMR Instance Group.
	//
	//    * appstream:fleet:DesiredCapacity - The desired capacity of an AppStream
	//    2.0 fleet.
	//
	//    * dynamodb:table:ReadCapacityUnits - The provisioned read capacity for
	//    a DynamoDB table.
	//
	//    * dynamodb:table:WriteCapacityUnits - The provisioned write capacity for
	//    a DynamoDB table.
	//
	//    * dynamodb:index:ReadCapacityUnits - The provisioned read capacity for
	//    a DynamoDB global secondary index.
	//
	//    * dynamodb:index:WriteCapacityUnits - The provisioned write capacity for
	//    a DynamoDB global secondary index.
	//
	//    * rds:cluster:ReadReplicaCount - The count of Aurora Replicas in an Aurora
	//    DB cluster. Available for Aurora MySQL-compatible edition and Aurora PostgreSQL-compatible
	//    edition.
	//
	//    * sagemaker:variant:DesiredInstanceCount - The number of EC2 instances
	//    for an Amazon SageMaker model endpoint variant.
	//
	//    * custom-resource:ResourceType:Property - The scalable dimension for a
	//    custom resource provided by your own application or service.
	//
	//    * comprehend:document-classifier-endpoint:DesiredInferenceUnits - The
	//    number of inference units for an Amazon Comprehend document classification
	//    endpoint.
	//
	//    * lambda:function:ProvisionedConcurrency - The provisioned concurrency
	//    for a Lambda function.
	//
	//    * cassandra:table:ReadCapacityUnits - The provisioned read capacity for
	//    an Amazon Keyspaces table.
	//
	//    * cassandra:table:WriteCapacityUnits - The provisioned write capacity
	//    for an Amazon Keyspaces table.
	//
	// ScalableDimension is a required field
	ScalableDimension ScalableDimension `type:"string" required:"true" enum:"true"`

	// The name of the scheduled action.
	//
	// ScheduledActionName is a required field
	ScheduledActionName *string `min:"1" type:"string" required:"true"`

	// The namespace of the AWS service that provides the resource. For a resource
	// provided by your own application or service, use custom-resource instead.
	//
	// ServiceNamespace is a required field
	ServiceNamespace ServiceNamespace `type:"string" required:"true" enum:"true"`
}

// String returns the string representation
func (s DeleteScheduledActionInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteScheduledActionInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteScheduledActionInput"}

	if s.ResourceId == nil {
		invalidParams.Add(aws.NewErrParamRequired("ResourceId"))
	}
	if s.ResourceId != nil && len(*s.ResourceId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ResourceId", 1))
	}
	if len(s.ScalableDimension) == 0 {
		invalidParams.Add(aws.NewErrParamRequired("ScalableDimension"))
	}

	if s.ScheduledActionName == nil {
		invalidParams.Add(aws.NewErrParamRequired("ScheduledActionName"))
	}
	if s.ScheduledActionName != nil && len(*s.ScheduledActionName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ScheduledActionName", 1))
	}
	if len(s.ServiceNamespace) == 0 {
		invalidParams.Add(aws.NewErrParamRequired("ServiceNamespace"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DeleteScheduledActionOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DeleteScheduledActionOutput) String() string {
	return awsutil.Prettify(s)
}

const opDeleteScheduledAction = "DeleteScheduledAction"

// DeleteScheduledActionRequest returns a request value for making API operation for
// Application Auto Scaling.
//
// Deletes the specified scheduled action for an Application Auto Scaling scalable
// target.
//
// For more information, see Delete a Scheduled Action (https://docs.aws.amazon.com/autoscaling/application/userguide/application-auto-scaling-scheduled-scaling.html#delete-scheduled-action)
// in the Application Auto Scaling User Guide.
//
//    // Example sending a request using DeleteScheduledActionRequest.
//    req := client.DeleteScheduledActionRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/application-autoscaling-2016-02-06/DeleteScheduledAction
func (c *Client) DeleteScheduledActionRequest(input *DeleteScheduledActionInput) DeleteScheduledActionRequest {
	op := &aws.Operation{
		Name:       opDeleteScheduledAction,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DeleteScheduledActionInput{}
	}

	req := c.newRequest(op, input, &DeleteScheduledActionOutput{})

	return DeleteScheduledActionRequest{Request: req, Input: input, Copy: c.DeleteScheduledActionRequest}
}

// DeleteScheduledActionRequest is the request type for the
// DeleteScheduledAction API operation.
type DeleteScheduledActionRequest struct {
	*aws.Request
	Input *DeleteScheduledActionInput
	Copy  func(*DeleteScheduledActionInput) DeleteScheduledActionRequest
}

// Send marshals and sends the DeleteScheduledAction API request.
func (r DeleteScheduledActionRequest) Send(ctx context.Context) (*DeleteScheduledActionResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteScheduledActionResponse{
		DeleteScheduledActionOutput: r.Request.Data.(*DeleteScheduledActionOutput),
		response:                    &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteScheduledActionResponse is the response type for the
// DeleteScheduledAction API operation.
type DeleteScheduledActionResponse struct {
	*DeleteScheduledActionOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteScheduledAction request.
func (r *DeleteScheduledActionResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
