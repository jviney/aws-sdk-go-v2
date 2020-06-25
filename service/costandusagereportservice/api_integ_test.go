// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// +build integration

package costandusagereportservice_test

import (
	"context"
	"testing"
	"time"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/aws/awserr"
	"github.com/jviney/aws-sdk-go-v2/aws/defaults"
	"github.com/jviney/aws-sdk-go-v2/internal/awstesting/integration"
	"github.com/jviney/aws-sdk-go-v2/service/costandusagereportservice"
)

var _ aws.Config
var _ awserr.Error

func TestInteg_00_DescribeReportDefinitions(t *testing.T) {
	ctx, cancelFn := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelFn()

	cfg := integration.ConfigWithDefaultRegion("us-east-1")
	svc := costandusagereportservice.New(cfg)
	params := &costandusagereportservice.DescribeReportDefinitionsInput{}

	req := svc.DescribeReportDefinitionsRequest(params)
	req.Handlers.Validate.Remove(defaults.ValidateParametersHandler)
	_, err := req.Send(ctx)
	if err != nil {
		t.Errorf("expect no error, got %v", err)
	}
}
