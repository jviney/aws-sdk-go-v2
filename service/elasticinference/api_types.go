// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package elasticinference

import (
	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
	"github.com/jviney/aws-sdk-go-v2/private/protocol"
)

var _ aws.Config
var _ = awsutil.Prettify

// The details of an Elastic Inference Accelerator type.
type AcceleratorType struct {
	_ struct{} `type:"structure"`

	// The name of the Elastic Inference Accelerator type.
	AcceleratorTypeName *string `locationName:"acceleratorTypeName" min:"1" type:"string"`

	// The memory information of the Elastic Inference Accelerator type.
	MemoryInfo *MemoryInfo `locationName:"memoryInfo" type:"structure"`

	// The throughput information of the Elastic Inference Accelerator type.
	ThroughputInfo []KeyValuePair `locationName:"throughputInfo" type:"list"`
}

// String returns the string representation
func (s AcceleratorType) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s AcceleratorType) MarshalFields(e protocol.FieldEncoder) error {
	if s.AcceleratorTypeName != nil {
		v := *s.AcceleratorTypeName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "acceleratorTypeName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.MemoryInfo != nil {
		v := s.MemoryInfo

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "memoryInfo", v, metadata)
	}
	if s.ThroughputInfo != nil {
		v := s.ThroughputInfo

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "throughputInfo", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	return nil
}

// The offering for an Elastic Inference Accelerator type.
type AcceleratorTypeOffering struct {
	_ struct{} `type:"structure"`

	// The name of the Elastic Inference Accelerator type.
	AcceleratorType *string `locationName:"acceleratorType" min:"1" type:"string"`

	// The location for the offering. It will return either the region, availability
	// zone or availability zone id for the offering depending on the locationType
	// value.
	Location *string `locationName:"location" min:"1" type:"string"`

	// The location type for the offering. It can assume the following values: region:
	// defines that the offering is at the regional level. availability-zone: defines
	// that the offering is at the availability zone level. availability-zone-id:
	// defines that the offering is at the availability zone level, defined by the
	// availability zone id.
	LocationType LocationType `locationName:"locationType" min:"1" type:"string" enum:"true"`
}

// String returns the string representation
func (s AcceleratorTypeOffering) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s AcceleratorTypeOffering) MarshalFields(e protocol.FieldEncoder) error {
	if s.AcceleratorType != nil {
		v := *s.AcceleratorType

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "acceleratorType", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Location != nil {
		v := *s.Location

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "location", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if len(s.LocationType) > 0 {
		v := s.LocationType

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "locationType", protocol.QuotedValue{ValueMarshaler: v}, metadata)
	}
	return nil
}

// The details of an Elastic Inference Accelerator.
type ElasticInferenceAccelerator struct {
	_ struct{} `type:"structure"`

	// The health of the Elastic Inference Accelerator.
	AcceleratorHealth *ElasticInferenceAcceleratorHealth `locationName:"acceleratorHealth" type:"structure"`

	// The ID of the Elastic Inference Accelerator.
	AcceleratorId *string `locationName:"acceleratorId" min:"1" type:"string"`

	// The type of the Elastic Inference Accelerator.
	AcceleratorType *string `locationName:"acceleratorType" min:"1" type:"string"`

	// The ARN of the resource that the Elastic Inference Accelerator is attached
	// to.
	AttachedResource *string `locationName:"attachedResource" min:"1" type:"string"`

	// The availability zone where the Elastic Inference Accelerator is present.
	AvailabilityZone *string `locationName:"availabilityZone" min:"1" type:"string"`
}

// String returns the string representation
func (s ElasticInferenceAccelerator) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ElasticInferenceAccelerator) MarshalFields(e protocol.FieldEncoder) error {
	if s.AcceleratorHealth != nil {
		v := s.AcceleratorHealth

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "acceleratorHealth", v, metadata)
	}
	if s.AcceleratorId != nil {
		v := *s.AcceleratorId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "acceleratorId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.AcceleratorType != nil {
		v := *s.AcceleratorType

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "acceleratorType", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.AttachedResource != nil {
		v := *s.AttachedResource

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "attachedResource", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.AvailabilityZone != nil {
		v := *s.AvailabilityZone

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "availabilityZone", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// The health details of an Elastic Inference Accelerator.
type ElasticInferenceAcceleratorHealth struct {
	_ struct{} `type:"structure"`

	// The health status of the Elastic Inference Accelerator.
	Status *string `locationName:"status" min:"1" type:"string"`
}

// String returns the string representation
func (s ElasticInferenceAcceleratorHealth) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ElasticInferenceAcceleratorHealth) MarshalFields(e protocol.FieldEncoder) error {
	if s.Status != nil {
		v := *s.Status

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "status", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// A filter expression for the Elastic Inference Accelerator list.
type Filter struct {
	_ struct{} `type:"structure"`

	// The filter name for the Elastic Inference Accelerator list. It can assume
	// the following values: accelerator-type: the type of Elastic Inference Accelerator
	// to filter for. instance-id: an EC2 instance id to filter for.
	Name *string `locationName:"name" min:"1" type:"string"`

	// The values for the filter of the Elastic Inference Accelerator list.
	Values []string `locationName:"values" type:"list"`
}

// String returns the string representation
func (s Filter) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *Filter) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "Filter"}
	if s.Name != nil && len(*s.Name) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Name", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s Filter) MarshalFields(e protocol.FieldEncoder) error {
	if s.Name != nil {
		v := *s.Name

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "name", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Values != nil {
		v := s.Values

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "values", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddValue(protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v1)})
		}
		ls0.End()

	}
	return nil
}

// A throughput entry for an Elastic Inference Accelerator type.
type KeyValuePair struct {
	_ struct{} `type:"structure"`

	// The throughput value of the Elastic Inference Accelerator type. It can assume
	// the following values: TFLOPS16bit: the throughput expressed in 16bit TeraFLOPS.
	// TFLOPS32bit: the throughput expressed in 32bit TeraFLOPS.
	Key *string `locationName:"key" min:"1" type:"string"`

	// The throughput value of the Elastic Inference Accelerator type.
	Value *int64 `locationName:"value" type:"integer"`
}

// String returns the string representation
func (s KeyValuePair) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s KeyValuePair) MarshalFields(e protocol.FieldEncoder) error {
	if s.Key != nil {
		v := *s.Key

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "key", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Value != nil {
		v := *s.Value

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "value", protocol.Int64Value(v), metadata)
	}
	return nil
}

// The memory information of an Elastic Inference Accelerator type.
type MemoryInfo struct {
	_ struct{} `type:"structure"`

	// The size in mebibytes of the Elastic Inference Accelerator type.
	SizeInMiB *int64 `locationName:"sizeInMiB" type:"integer"`
}

// String returns the string representation
func (s MemoryInfo) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s MemoryInfo) MarshalFields(e protocol.FieldEncoder) error {
	if s.SizeInMiB != nil {
		v := *s.SizeInMiB

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "sizeInMiB", protocol.Int64Value(v), metadata)
	}
	return nil
}
