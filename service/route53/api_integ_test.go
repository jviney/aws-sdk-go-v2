// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// +build integration

package route53_test

import (
	"context"
	"testing"
	"time"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/aws/awserr"
	"github.com/jviney/aws-sdk-go-v2/aws/defaults"
	"github.com/jviney/aws-sdk-go-v2/internal/awstesting/integration"
	"github.com/jviney/aws-sdk-go-v2/service/route53"
)

var _ aws.Config
var _ awserr.Error

func TestInteg_00_ListHostedZones(t *testing.T) {
	ctx, cancelFn := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelFn()

	cfg := integration.ConfigWithDefaultRegion("us-east-1")
	svc := route53.New(cfg)
	params := &route53.ListHostedZonesInput{}

	req := svc.ListHostedZonesRequest(params)
	req.Handlers.Validate.Remove(defaults.ValidateParametersHandler)
	_, err := req.Send(ctx)
	if err != nil {
		t.Errorf("expect no error, got %v", err)
	}
}
func TestInteg_01_GetHostedZone(t *testing.T) {
	ctx, cancelFn := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelFn()

	cfg := integration.ConfigWithDefaultRegion("us-east-1")
	svc := route53.New(cfg)
	params := &route53.GetHostedZoneInput{
		Id: aws.String("fake-zone"),
	}

	req := svc.GetHostedZoneRequest(params)
	req.Handlers.Validate.Remove(defaults.ValidateParametersHandler)
	_, err := req.Send(ctx)
	if err == nil {
		t.Fatalf("expect request to fail")
	}
	aerr, ok := err.(awserr.RequestFailure)
	if !ok {
		t.Fatalf("expect awserr, was %T", err)
	}
	if len(aerr.Code()) == 0 {
		t.Errorf("expect non-empty error code")
	}
	if v := aerr.Code(); v == aws.ErrCodeSerialization {
		t.Errorf("expect API error code got serialization failure")
	}
}
