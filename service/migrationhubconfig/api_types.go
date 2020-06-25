// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package migrationhubconfig

import (
	"time"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
)

var _ aws.Config
var _ = awsutil.Prettify

// A home region control is an object that specifies the home region for an
// account, with some additional information. It contains a target (always of
// type ACCOUNT), an ID, and a time at which the home region was set.
type HomeRegionControl struct {
	_ struct{} `type:"structure"`

	// A unique identifier that's generated for each home region control. It's always
	// a string that begins with "hrc-" followed by 12 lowercase letters and numbers.
	ControlId *string `min:"1" type:"string"`

	// The AWS Region that's been set as home region. For example, "us-west-2" or
	// "eu-central-1" are valid home regions.
	HomeRegion *string `min:"1" type:"string"`

	// A timestamp representing the time when the customer called CreateHomeregionControl
	// and set the home region for the account.
	RequestedTime *time.Time `type:"timestamp"`

	// The target parameter specifies the identifier to which the home region is
	// applied, which is always an ACCOUNT. It applies the home region to the current
	// ACCOUNT.
	Target *Target `type:"structure"`
}

// String returns the string representation
func (s HomeRegionControl) String() string {
	return awsutil.Prettify(s)
}

// The target parameter specifies the identifier to which the home region is
// applied, which is always an ACCOUNT. It applies the home region to the current
// ACCOUNT.
type Target struct {
	_ struct{} `type:"structure"`

	// The TargetID is a 12-character identifier of the ACCOUNT for which the control
	// was created. (This must be the current account.)
	Id *string `min:"12" type:"string"`

	// The target type is always an ACCOUNT.
	//
	// Type is a required field
	Type TargetType `type:"string" required:"true" enum:"true"`
}

// String returns the string representation
func (s Target) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *Target) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "Target"}
	if s.Id != nil && len(*s.Id) < 12 {
		invalidParams.Add(aws.NewErrParamMinLen("Id", 12))
	}
	if len(s.Type) == 0 {
		invalidParams.Add(aws.NewErrParamRequired("Type"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}
