// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// +build integration

package marketplacecommerceanalytics_test

import (
	"context"
	"testing"
	"time"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/aws/awserr"
	"github.com/jviney/aws-sdk-go-v2/aws/defaults"
	"github.com/jviney/aws-sdk-go-v2/internal/awstesting/integration"
	"github.com/jviney/aws-sdk-go-v2/service/marketplacecommerceanalytics"
)

var _ aws.Config
var _ awserr.Error

func TestInteg_00_GenerateDataSet(t *testing.T) {
	ctx, cancelFn := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelFn()

	cfg := integration.ConfigWithDefaultRegion("us-east-1")
	svc := marketplacecommerceanalytics.New(cfg)
	params := &marketplacecommerceanalytics.GenerateDataSetInput{
		DataSetPublicationDate:  aws.Time(time.Unix(0, 0)),
		DataSetType:             marketplacecommerceanalytics.DataSetType("fake-type"),
		DestinationS3BucketName: aws.String("fake-bucket"),
		RoleNameArn:             aws.String("fake-arn"),
		SnsTopicArn:             aws.String("fake-arn"),
	}

	req := svc.GenerateDataSetRequest(params)
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
