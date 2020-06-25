// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package awsendpointdiscoverytest

import (
	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
)

var _ aws.Config
var _ = awsutil.Prettify

type Endpoint struct {
	_ struct{} `type:"structure"`

	// Address is a required field
	Address *string `type:"string" required:"true"`

	// CachePeriodInMinutes is a required field
	CachePeriodInMinutes *int64 `type:"long" required:"true"`
}

// String returns the string representation
func (s Endpoint) String() string {
	return awsutil.Prettify(s)
}
