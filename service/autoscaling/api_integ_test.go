// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// +build integration

package autoscaling_test

import (
	"context"
	"testing"
	"time"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/aws/awserr"
	"github.com/jviney/aws-sdk-go-v2/aws/defaults"
	"github.com/jviney/aws-sdk-go-v2/internal/awstesting/integration"
	"github.com/jviney/aws-sdk-go-v2/service/autoscaling"
)

var _ aws.Config
var _ awserr.Error

func TestInteg_00_DescribeScalingProcessTypes(t *testing.T) {
	ctx, cancelFn := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelFn()

	cfg := integration.ConfigWithDefaultRegion("us-west-2")
	svc := autoscaling.New(cfg)
	params := &autoscaling.DescribeScalingProcessTypesInput{}

	req := svc.DescribeScalingProcessTypesRequest(params)
	req.Handlers.Validate.Remove(defaults.ValidateParametersHandler)
	_, err := req.Send(ctx)
	if err != nil {
		t.Errorf("expect no error, got %v", err)
	}
}
func TestInteg_01_CreateLaunchConfiguration(t *testing.T) {
	ctx, cancelFn := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelFn()

	cfg := integration.ConfigWithDefaultRegion("us-west-2")
	svc := autoscaling.New(cfg)
	params := &autoscaling.CreateLaunchConfigurationInput{
		ImageId:                 aws.String("ami-12345678"),
		InstanceType:            aws.String("m1.small"),
		LaunchConfigurationName: aws.String("hello, world"),
	}

	req := svc.CreateLaunchConfigurationRequest(params)
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
