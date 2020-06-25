// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package schemas

import (
	"time"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
	"github.com/jviney/aws-sdk-go-v2/private/protocol"
)

var _ aws.Config
var _ = awsutil.Prettify

type DiscovererSummary struct {
	_ struct{} `type:"structure"`

	// The ARN of the discoverer.
	DiscovererArn *string `type:"string"`

	// The ID of the discoverer.
	DiscovererId *string `type:"string"`

	// The ARN of the event bus.
	SourceArn *string `type:"string"`

	// The state of the discoverer.
	State DiscovererState `type:"string" enum:"true"`

	// Tags associated with the resource.
	Tags map[string]string `locationName:"tags" type:"map"`
}

// String returns the string representation
func (s DiscovererSummary) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DiscovererSummary) MarshalFields(e protocol.FieldEncoder) error {
	if s.DiscovererArn != nil {
		v := *s.DiscovererArn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "DiscovererArn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.DiscovererId != nil {
		v := *s.DiscovererId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "DiscovererId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.SourceArn != nil {
		v := *s.SourceArn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "SourceArn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if len(s.State) > 0 {
		v := s.State

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "State", protocol.QuotedValue{ValueMarshaler: v}, metadata)
	}
	if s.Tags != nil {
		v := s.Tags

		metadata := protocol.Metadata{}
		ms0 := e.Map(protocol.BodyTarget, "tags", metadata)
		ms0.Start()
		for k1, v1 := range v {
			ms0.MapSetValue(k1, protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v1)})
		}
		ms0.End()

	}
	return nil
}

type RegistrySummary struct {
	_ struct{} `type:"structure"`

	// The ARN of the registry.
	RegistryArn *string `type:"string"`

	// The name of the registry.
	RegistryName *string `type:"string"`

	// Tags associated with the registry.
	Tags map[string]string `locationName:"tags" type:"map"`
}

// String returns the string representation
func (s RegistrySummary) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s RegistrySummary) MarshalFields(e protocol.FieldEncoder) error {
	if s.RegistryArn != nil {
		v := *s.RegistryArn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "RegistryArn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.RegistryName != nil {
		v := *s.RegistryName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "RegistryName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Tags != nil {
		v := s.Tags

		metadata := protocol.Metadata{}
		ms0 := e.Map(protocol.BodyTarget, "tags", metadata)
		ms0.Start()
		for k1, v1 := range v {
			ms0.MapSetValue(k1, protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v1)})
		}
		ms0.End()

	}
	return nil
}

// A summary of schema details.
type SchemaSummary struct {
	_ struct{} `type:"structure"`

	// The date and time that schema was modified.
	LastModified *time.Time `type:"timestamp" timestampFormat:"iso8601"`

	// The ARN of the schema.
	SchemaArn *string `type:"string"`

	// The name of the schema.
	SchemaName *string `type:"string"`

	// Tags associated with the schema.
	Tags map[string]string `locationName:"tags" type:"map"`

	// The number of versions available for the schema.
	VersionCount *int64 `type:"long"`
}

// String returns the string representation
func (s SchemaSummary) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s SchemaSummary) MarshalFields(e protocol.FieldEncoder) error {
	if s.LastModified != nil {
		v := *s.LastModified

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "LastModified",
			protocol.TimeValue{V: v, Format: "iso8601", QuotedFormatTime: true}, metadata)
	}
	if s.SchemaArn != nil {
		v := *s.SchemaArn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "SchemaArn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.SchemaName != nil {
		v := *s.SchemaName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "SchemaName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Tags != nil {
		v := s.Tags

		metadata := protocol.Metadata{}
		ms0 := e.Map(protocol.BodyTarget, "tags", metadata)
		ms0.Start()
		for k1, v1 := range v {
			ms0.MapSetValue(k1, protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v1)})
		}
		ms0.End()

	}
	if s.VersionCount != nil {
		v := *s.VersionCount

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "VersionCount", protocol.Int64Value(v), metadata)
	}
	return nil
}

type SchemaVersionSummary struct {
	_ struct{} `type:"structure"`

	// The ARN of the schema version.
	SchemaArn *string `type:"string"`

	// The name of the schema.
	SchemaName *string `type:"string"`

	// The version number of the schema.
	SchemaVersion *string `type:"string"`
}

// String returns the string representation
func (s SchemaVersionSummary) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s SchemaVersionSummary) MarshalFields(e protocol.FieldEncoder) error {
	if s.SchemaArn != nil {
		v := *s.SchemaArn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "SchemaArn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.SchemaName != nil {
		v := *s.SchemaName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "SchemaName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.SchemaVersion != nil {
		v := *s.SchemaVersion

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "SchemaVersion", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type SearchSchemaSummary struct {
	_ struct{} `type:"structure"`

	// The name of the registry.
	RegistryName *string `type:"string"`

	// The ARN of the schema.
	SchemaArn *string `type:"string"`

	// The name of the schema.
	SchemaName *string `type:"string"`

	// An array of schema version summaries.
	SchemaVersions []SearchSchemaVersionSummary `type:"list"`
}

// String returns the string representation
func (s SearchSchemaSummary) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s SearchSchemaSummary) MarshalFields(e protocol.FieldEncoder) error {
	if s.RegistryName != nil {
		v := *s.RegistryName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "RegistryName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.SchemaArn != nil {
		v := *s.SchemaArn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "SchemaArn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.SchemaName != nil {
		v := *s.SchemaName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "SchemaName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.SchemaVersions != nil {
		v := s.SchemaVersions

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "SchemaVersions", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	return nil
}

type SearchSchemaVersionSummary struct {
	_ struct{} `type:"structure"`

	// The date the schema version was created.
	CreatedDate *time.Time `type:"timestamp" timestampFormat:"iso8601"`

	// The version number of the schema
	SchemaVersion *string `type:"string"`
}

// String returns the string representation
func (s SearchSchemaVersionSummary) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s SearchSchemaVersionSummary) MarshalFields(e protocol.FieldEncoder) error {
	if s.CreatedDate != nil {
		v := *s.CreatedDate

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "CreatedDate",
			protocol.TimeValue{V: v, Format: "iso8601", QuotedFormatTime: true}, metadata)
	}
	if s.SchemaVersion != nil {
		v := *s.SchemaVersion

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "SchemaVersion", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}
