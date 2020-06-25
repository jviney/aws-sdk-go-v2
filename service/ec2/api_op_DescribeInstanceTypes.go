// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ec2

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
)

type DescribeInstanceTypesInput struct {
	_ struct{} `type:"structure"`

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have
	// the required permissions, the error response is DryRunOperation. Otherwise,
	// it is UnauthorizedOperation.
	DryRun *bool `type:"boolean"`

	// One or more filters. Filter names and values are case-sensitive.
	//
	//    * auto-recovery-supported - Indicates whether auto recovery is supported.
	//    (true | false)
	//
	//    * bare-metal - Indicates whether it is a bare metal instance type. (true
	//    | false)
	//
	//    * burstable-performance-supported - Indicates whether it is a burstable
	//    performance instance type. (true | false)
	//
	//    * current-generation - Indicates whether this instance type is the latest
	//    generation instance type of an instance family. (true | false)
	//
	//    * ebs-info.ebs-optimized-info.baseline-bandwidth-in-mbps - The baseline
	//    bandwidth performance for an EBS-optimized instance type, in Mbps.
	//
	//    * ebs-info.ebs-optimized-info.baseline-throughput-in-mbps - The baseline
	//    throughput performance for an EBS-optimized instance type, in MBps.
	//
	//    * ebs-info.ebs-optimized-info.baseline-iops - The baseline input/output
	//    storage operations per second for an EBS-optimized instance type.
	//
	//    * ebs-info.ebs-optimized-info.maximum-bandwidth-in-mbps - The maximum
	//    bandwidth performance for an EBS-optimized instance type, in Mbps.
	//
	//    * ebs-info.ebs-optimized-info.maximum-throughput-in-mbps - The maximum
	//    throughput performance for an EBS-optimized instance type, in MBps.
	//
	//    * ebs-info.ebs-optimized-info.maximum-iops - The maximum input/output
	//    storage operations per second for an EBS-optimized instance type.
	//
	//    * ebs-info.ebs-optimized-support - Indicates whether the instance type
	//    is EBS-optimized. (supported | unsupported | default)
	//
	//    * ebs-info.encryption-support - Indicates whether EBS encryption is supported.
	//    (supported | unsupported)
	//
	//    * free-tier-eligible - Indicates whether the instance type is eligible
	//    to use in the free tier. (true | false)
	//
	//    * hibernation-supported - Indicates whether On-Demand hibernation is supported.
	//    (true | false)
	//
	//    * hypervisor - The hypervisor used. (nitro | xen)
	//
	//    * instance-storage-info.disk.count - The number of local disks.
	//
	//    * instance-storage-info.disk.size-in-gb - The storage size of each instance
	//    storage disk, in GB.
	//
	//    * instance-storage-info.disk.type - The storage technology for the local
	//    instance storage disks. (hdd | ssd)
	//
	//    * instance-storage-info.total-size-in-gb - The total amount of storage
	//    available from all local instance storage, in GB.
	//
	//    * instance-storage-supported - Indicates whether the instance type has
	//    local instance storage. (true | false)
	//
	//    * memory-info.size-in-mib - The memory size.
	//
	//    * network-info.ena-support - Indicates whether Elastic Network Adapter
	//    (ENA) is supported or required. (required | supported | unsupported)
	//
	//    * network-info.efa-supported - Indicates whether the instance type supports
	//    Elastic Fabric Adapter (EFA). (true | false)
	//
	//    * network-info.ipv4-addresses-per-interface - The maximum number of private
	//    IPv4 addresses per network interface.
	//
	//    * network-info.ipv6-addresses-per-interface - The maximum number of private
	//    IPv6 addresses per network interface.
	//
	//    * network-info.ipv6-supported - Indicates whether the instance type supports
	//    IPv6. (true | false)
	//
	//    * network-info.maximum-network-interfaces - The maximum number of network
	//    interfaces per instance.
	//
	//    * network-info.network-performance - Describes the network performance.
	//
	//    * processor-info.sustained-clock-speed-in-ghz - The CPU clock speed, in
	//    GHz.
	//
	//    * vcpu-info.default-cores - The default number of cores for the instance
	//    type.
	//
	//    * vcpu-info.default-threads-per-core - The default number of threads per
	//    core for the instance type.
	//
	//    * vcpu-info.default-vcpus - The default number of vCPUs for the instance
	//    type.
	Filters []Filter `locationName:"Filter" locationNameList:"Filter" type:"list"`

	// The instance types. For more information, see Instance Types (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/instance-types.html)
	// in the Amazon Elastic Compute Cloud User Guide.
	InstanceTypes []InstanceType `locationName:"InstanceType" type:"list"`

	// The maximum number of results to return for the request in a single page.
	// The remaining results can be seen by sending another request with the next
	// token value.
	MaxResults *int64 `min:"5" type:"integer"`

	// The token to retrieve the next page of results.
	NextToken *string `type:"string"`
}

// String returns the string representation
func (s DescribeInstanceTypesInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeInstanceTypesInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeInstanceTypesInput"}
	if s.MaxResults != nil && *s.MaxResults < 5 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxResults", 5))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DescribeInstanceTypesOutput struct {
	_ struct{} `type:"structure"`

	// The instance type. For more information, see Instance Types (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/instance-types.html)
	// in the Amazon Elastic Compute Cloud User Guide.
	InstanceTypes []InstanceTypeInfo `locationName:"instanceTypeSet" locationNameList:"item" type:"list"`

	// The token to use to retrieve the next page of results. This value is null
	// when there are no more results to return.
	NextToken *string `locationName:"nextToken" type:"string"`
}

// String returns the string representation
func (s DescribeInstanceTypesOutput) String() string {
	return awsutil.Prettify(s)
}

const opDescribeInstanceTypes = "DescribeInstanceTypes"

// DescribeInstanceTypesRequest returns a request value for making API operation for
// Amazon Elastic Compute Cloud.
//
// Describes the details of the instance types that are offered in a location.
// The results can be filtered by the attributes of the instance types.
//
//    // Example sending a request using DescribeInstanceTypesRequest.
//    req := client.DescribeInstanceTypesRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/DescribeInstanceTypes
func (c *Client) DescribeInstanceTypesRequest(input *DescribeInstanceTypesInput) DescribeInstanceTypesRequest {
	op := &aws.Operation{
		Name:       opDescribeInstanceTypes,
		HTTPMethod: "POST",
		HTTPPath:   "/",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"NextToken"},
			OutputTokens:    []string{"NextToken"},
			LimitToken:      "MaxResults",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &DescribeInstanceTypesInput{}
	}

	req := c.newRequest(op, input, &DescribeInstanceTypesOutput{})

	return DescribeInstanceTypesRequest{Request: req, Input: input, Copy: c.DescribeInstanceTypesRequest}
}

// DescribeInstanceTypesRequest is the request type for the
// DescribeInstanceTypes API operation.
type DescribeInstanceTypesRequest struct {
	*aws.Request
	Input *DescribeInstanceTypesInput
	Copy  func(*DescribeInstanceTypesInput) DescribeInstanceTypesRequest
}

// Send marshals and sends the DescribeInstanceTypes API request.
func (r DescribeInstanceTypesRequest) Send(ctx context.Context) (*DescribeInstanceTypesResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeInstanceTypesResponse{
		DescribeInstanceTypesOutput: r.Request.Data.(*DescribeInstanceTypesOutput),
		response:                    &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewDescribeInstanceTypesRequestPaginator returns a paginator for DescribeInstanceTypes.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.DescribeInstanceTypesRequest(input)
//   p := ec2.NewDescribeInstanceTypesRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewDescribeInstanceTypesPaginator(req DescribeInstanceTypesRequest) DescribeInstanceTypesPaginator {
	return DescribeInstanceTypesPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *DescribeInstanceTypesInput
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

// DescribeInstanceTypesPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type DescribeInstanceTypesPaginator struct {
	aws.Pager
}

func (p *DescribeInstanceTypesPaginator) CurrentPage() *DescribeInstanceTypesOutput {
	return p.Pager.CurrentPage().(*DescribeInstanceTypesOutput)
}

// DescribeInstanceTypesResponse is the response type for the
// DescribeInstanceTypes API operation.
type DescribeInstanceTypesResponse struct {
	*DescribeInstanceTypesOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeInstanceTypes request.
func (r *DescribeInstanceTypesResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
