// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package codeguruprofiler

import (
	"time"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
	"github.com/jviney/aws-sdk-go-v2/private/protocol"
)

var _ aws.Config
var _ = awsutil.Prettify

type AgentConfiguration struct {
	_ struct{} `type:"structure"`

	// PeriodInSeconds is a required field
	PeriodInSeconds *int64 `locationName:"periodInSeconds" type:"integer" required:"true"`

	// ShouldProfile is a required field
	ShouldProfile *bool `locationName:"shouldProfile" type:"boolean" required:"true"`
}

// String returns the string representation
func (s AgentConfiguration) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s AgentConfiguration) MarshalFields(e protocol.FieldEncoder) error {
	if s.PeriodInSeconds != nil {
		v := *s.PeriodInSeconds

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "periodInSeconds", protocol.Int64Value(v), metadata)
	}
	if s.ShouldProfile != nil {
		v := *s.ShouldProfile

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "shouldProfile", protocol.BoolValue(v), metadata)
	}
	return nil
}

type AgentOrchestrationConfig struct {
	_ struct{} `type:"structure"`

	// ProfilingEnabled is a required field
	ProfilingEnabled *bool `locationName:"profilingEnabled" type:"boolean" required:"true"`
}

// String returns the string representation
func (s AgentOrchestrationConfig) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *AgentOrchestrationConfig) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "AgentOrchestrationConfig"}

	if s.ProfilingEnabled == nil {
		invalidParams.Add(aws.NewErrParamRequired("ProfilingEnabled"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s AgentOrchestrationConfig) MarshalFields(e protocol.FieldEncoder) error {
	if s.ProfilingEnabled != nil {
		v := *s.ProfilingEnabled

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "profilingEnabled", protocol.BoolValue(v), metadata)
	}
	return nil
}

// Information about the time range of the latest available aggregated profile.
type AggregatedProfileTime struct {
	_ struct{} `type:"structure"`

	// The time period.
	Period AggregationPeriod `locationName:"period" type:"string" enum:"true"`

	// The start time.
	Start *time.Time `locationName:"start" type:"timestamp" timestampFormat:"iso8601"`
}

// String returns the string representation
func (s AggregatedProfileTime) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s AggregatedProfileTime) MarshalFields(e protocol.FieldEncoder) error {
	if len(s.Period) > 0 {
		v := s.Period

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "period", protocol.QuotedValue{ValueMarshaler: v}, metadata)
	}
	if s.Start != nil {
		v := *s.Start

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "start",
			protocol.TimeValue{V: v, Format: "iso8601", QuotedFormatTime: true}, metadata)
	}
	return nil
}

// Information about the profile time.
type ProfileTime struct {
	_ struct{} `type:"structure"`

	// The start time of the profile.
	Start *time.Time `locationName:"start" type:"timestamp" timestampFormat:"iso8601"`
}

// String returns the string representation
func (s ProfileTime) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ProfileTime) MarshalFields(e protocol.FieldEncoder) error {
	if s.Start != nil {
		v := *s.Start

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "start",
			protocol.TimeValue{V: v, Format: "iso8601", QuotedFormatTime: true}, metadata)
	}
	return nil
}

// The description of a profiling group.
type ProfilingGroupDescription struct {
	_ struct{} `type:"structure"`

	AgentOrchestrationConfig *AgentOrchestrationConfig `locationName:"agentOrchestrationConfig" type:"structure"`

	// The Amazon Resource Name (ARN) identifying the profiling group.
	Arn *string `locationName:"arn" type:"string"`

	// The time, in milliseconds since the epoch, when the profiling group was created.
	CreatedAt *time.Time `locationName:"createdAt" type:"timestamp" timestampFormat:"iso8601"`

	// The name of the profiling group.
	Name *string `locationName:"name" min:"1" type:"string"`

	// The status of the profiling group.
	ProfilingStatus *ProfilingStatus `locationName:"profilingStatus" type:"structure"`

	// The time, in milliseconds since the epoch, when the profiling group was last
	// updated.
	UpdatedAt *time.Time `locationName:"updatedAt" type:"timestamp" timestampFormat:"iso8601"`
}

// String returns the string representation
func (s ProfilingGroupDescription) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ProfilingGroupDescription) MarshalFields(e protocol.FieldEncoder) error {
	if s.AgentOrchestrationConfig != nil {
		v := s.AgentOrchestrationConfig

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "agentOrchestrationConfig", v, metadata)
	}
	if s.Arn != nil {
		v := *s.Arn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "arn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.CreatedAt != nil {
		v := *s.CreatedAt

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "createdAt",
			protocol.TimeValue{V: v, Format: "iso8601", QuotedFormatTime: true}, metadata)
	}
	if s.Name != nil {
		v := *s.Name

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "name", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.ProfilingStatus != nil {
		v := s.ProfilingStatus

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "profilingStatus", v, metadata)
	}
	if s.UpdatedAt != nil {
		v := *s.UpdatedAt

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "updatedAt",
			protocol.TimeValue{V: v, Format: "iso8601", QuotedFormatTime: true}, metadata)
	}
	return nil
}

// Information about the profiling status.
type ProfilingStatus struct {
	_ struct{} `type:"structure"`

	// The time, in milliseconds since the epoch, when the latest agent was orchestrated.
	LatestAgentOrchestratedAt *time.Time `locationName:"latestAgentOrchestratedAt" type:"timestamp" timestampFormat:"iso8601"`

	// The time, in milliseconds since the epoch, when the latest agent was reported..
	LatestAgentProfileReportedAt *time.Time `locationName:"latestAgentProfileReportedAt" type:"timestamp" timestampFormat:"iso8601"`

	// The latest aggregated profile
	LatestAggregatedProfile *AggregatedProfileTime `locationName:"latestAggregatedProfile" type:"structure"`
}

// String returns the string representation
func (s ProfilingStatus) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ProfilingStatus) MarshalFields(e protocol.FieldEncoder) error {
	if s.LatestAgentOrchestratedAt != nil {
		v := *s.LatestAgentOrchestratedAt

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "latestAgentOrchestratedAt",
			protocol.TimeValue{V: v, Format: "iso8601", QuotedFormatTime: true}, metadata)
	}
	if s.LatestAgentProfileReportedAt != nil {
		v := *s.LatestAgentProfileReportedAt

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "latestAgentProfileReportedAt",
			protocol.TimeValue{V: v, Format: "iso8601", QuotedFormatTime: true}, metadata)
	}
	if s.LatestAggregatedProfile != nil {
		v := s.LatestAggregatedProfile

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "latestAggregatedProfile", v, metadata)
	}
	return nil
}
