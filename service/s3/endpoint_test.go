// +build go1.7

package s3

import (
	"bytes"
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"testing"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/aws/endpoints"
	"github.com/jviney/aws-sdk-go-v2/internal/awstesting/unit"
)

func TestEndpointARN(t *testing.T) {
	type config struct {
		region        string
		resolver      aws.EndpointResolver
		useARNRegion  bool
		useDualStack  bool
		useAccelerate bool
	}
	cases := map[string]struct {
		bucket                string
		config                config
		expectedEndpoint      string
		expectedSigningName   string
		expectedSigningRegion string
		expectedErr           string
	}{
		"AccessPoint": {
			bucket:                "arn:aws:s3:us-west-2:123456789012:accesspoint:myendpoint",
			config:                config{region: "us-west-2"},
			expectedEndpoint:      "https://myendpoint-123456789012.s3-accesspoint.us-west-2.amazonaws.com",
			expectedSigningName:   "s3",
			expectedSigningRegion: "us-west-2",
		},
		"AccessPoint slash delimiter": {
			bucket:                "arn:aws:s3:us-west-2:123456789012:accesspoint/myendpoint",
			config:                config{region: "us-west-2"},
			expectedEndpoint:      "https://myendpoint-123456789012.s3-accesspoint.us-west-2.amazonaws.com",
			expectedSigningName:   "s3",
			expectedSigningRegion: "us-west-2",
		},
		"AccessPoint other partition": {
			bucket:                "arn:aws-cn:s3:cn-north-1:123456789012:accesspoint:myendpoint",
			config:                config{region: "cn-north-1"},
			expectedEndpoint:      "https://myendpoint-123456789012.s3-accesspoint.cn-north-1.amazonaws.com.cn",
			expectedSigningName:   "s3",
			expectedSigningRegion: "cn-north-1",
		},
		"AccessPoint Cross-Region Disabled": {
			bucket:      "arn:aws:s3:ap-south-1:123456789012:accesspoint:myendpoint",
			config:      config{region: "us-west-2"},
			expectedErr: "client region does not match provided ARN region",
		},
		"AccessPoint Cross-Region Enabled": {
			bucket: "arn:aws:s3:ap-south-1:123456789012:accesspoint:myendpoint",
			config: config{
				region:       "us-west-2",
				useARNRegion: true,
			},
			expectedEndpoint:      "https://myendpoint-123456789012.s3-accesspoint.ap-south-1.amazonaws.com",
			expectedSigningName:   "s3",
			expectedSigningRegion: "ap-south-1",
		},
		"AccessPoint us-east-1": {
			bucket: "arn:aws:s3:us-east-1:123456789012:accesspoint:myendpoint",
			config: config{
				region:       "us-east-1",
				useARNRegion: true,
			},
			expectedEndpoint:      "https://myendpoint-123456789012.s3-accesspoint.us-east-1.amazonaws.com",
			expectedSigningName:   "s3",
			expectedSigningRegion: "us-east-1",
		},
		"AccessPoint us-east-1 cross region": {
			bucket: "arn:aws:s3:us-east-1:123456789012:accesspoint:myendpoint",
			config: config{
				region:       "us-west-2",
				useARNRegion: true,
			},
			expectedEndpoint:      "https://myendpoint-123456789012.s3-accesspoint.us-east-1.amazonaws.com",
			expectedSigningName:   "s3",
			expectedSigningRegion: "us-east-1",
		},
		"AccessPoint Cross-Partition not supported": {
			bucket: "arn:aws-cn:s3:cn-north-1:123456789012:accesspoint:myendpoint",
			config: config{
				region:       "us-west-2",
				useARNRegion: true,
				useDualStack: true,
			},
			expectedErr: "client partition does not match provided ARN partition",
		},
		"AccessPoint DualStack": {
			bucket: "arn:aws:s3:us-west-2:123456789012:accesspoint:myendpoint",
			config: config{
				region:       "us-west-2",
				useDualStack: true,
			},
			expectedEndpoint:      "https://myendpoint-123456789012.s3-accesspoint.dualstack.us-west-2.amazonaws.com",
			expectedSigningName:   "s3",
			expectedSigningRegion: "us-west-2",
		},
		"AccessPoint FIPS same region with cross region disabled": {
			bucket: "arn:aws-us-gov:s3:us-gov-west-1:123456789012:accesspoint:myendpoint",
			config: config{
				region: "fips-us-gov-west-1",
			},
			expectedEndpoint:      "https://myendpoint-123456789012.s3-accesspoint-fips-us-gov-west-1.amazonaws.com",
			expectedSigningName:   "s3",
			expectedSigningRegion: "us-gov-west-1",
		},
		"AccessPoint FIPS same region with cross region enabled": {
			bucket: "arn:aws-us-gov:s3:us-gov-west-1:123456789012:accesspoint:myendpoint",
			config: config{
				region:       "fips-us-gov-west-1",
				useARNRegion: true,
			},
			expectedEndpoint:      "https://myendpoint-123456789012.s3-accesspoint-fips-us-gov-west-1.amazonaws.com",
			expectedSigningName:   "s3",
			expectedSigningRegion: "us-gov-west-1",
		},
		"AccessPoint FIPS cross region not supported": {
			bucket: "arn:aws-us-gov:s3:us-gov-east-1:123456789012:accesspoint:myendpoint",
			config: config{
				region:       "fips-us-gov-west-1",
				useARNRegion: true,
			},
			expectedErr: "client configured for FIPS",
		},
		"AccessPoint Accelerate not supported": {
			bucket: "arn:aws:s3:us-west-2:123456789012:accesspoint:myendpoint",
			config: config{
				region:        "us-west-2",
				useARNRegion:  true,
				useAccelerate: true,
			},
			expectedErr: "client configured for S3 Accelerate",
		},
		"Custom Resolver Without PartitionID in ClientInfo": {
			bucket: "arn:aws:s3:us-west-2:123456789012:accesspoint:myendpoint",
			config: config{
				region: "us-west-2",
				resolver: aws.EndpointResolverFunc(
					func(service, region string) (aws.Endpoint, error) {
						switch region {
						case "us-west-2":
							return aws.Endpoint{
								URL:           "s3.us-west-2.amazonaws.com",
								SigningRegion: "us-west-2",
								SigningName:   service,
								SigningMethod: "s3v4",
							}, nil
						}
						return aws.Endpoint{}, nil
					}),
			},
			expectedErr: "client partition does not match provided ARN partition",
		},
		"Custom Resolver Without PartitionID in Cross-Region Target": {
			bucket: "arn:aws:s3:us-west-2:123456789012:accesspoint:myendpoint",
			config: config{
				region: "us-east-1",
				resolver: aws.EndpointResolverFunc(
					func(service, region string) (aws.Endpoint, error) {
						switch region {
						case "us-west-2":
							return aws.Endpoint{
								URL:           "s3.us-west-2.amazonaws.com",
								PartitionID:   "aws",
								SigningRegion: "us-west-2",
								SigningName:   service,
								SigningMethod: "s3v4",
							}, nil
						case "us-east-1":
							return aws.Endpoint{
								URL:           "s3.us-east-1.amazonaws.com",
								SigningRegion: "us-east-1",
								SigningName:   service,
								SigningMethod: "s3v4",
							}, nil
						}

						return aws.Endpoint{}, nil
					}),
			},
			expectedErr: "client partition does not match provided ARN partition",
		},
	}

	for name, c := range cases {
		t.Run(name, func(t *testing.T) {
			cfg := unit.Config()
			cfg.Region = c.config.region
			if c.config.resolver != nil {
				cfg.EndpointResolver = c.config.resolver
			} else {
				cfg.EndpointResolver = endpoints.NewDefaultResolver()
				cfg.EndpointResolver.(*endpoints.Resolver).UseDualStack = c.config.useDualStack
			}

			svc := New(cfg)
			svc.UseAccelerate = c.config.useAccelerate
			svc.UseARNRegion = c.config.useARNRegion

			req := svc.GetObjectRequest(&GetObjectInput{
				Bucket: &c.bucket,
				Key:    aws.String("testkey"),
			})
			req.Handlers.Send.Clear()
			req.Handlers.Send.PushBack(func(r *aws.Request) {
				defer func() {
					r.HTTPResponse = &http.Response{
						StatusCode:    200,
						ContentLength: 0,
						Body:          ioutil.NopCloser(bytes.NewReader(nil)),
					}
				}()
				if len(c.expectedErr) != 0 {
					return
				}

				endpoint := fmt.Sprintf("%s://%s", r.HTTPRequest.URL.Scheme, r.HTTPRequest.URL.Host)
				if e, a := c.expectedEndpoint, endpoint; e != a {
					t.Errorf("expected %v, got %v", e, a)
				}
				if e, a := c.expectedSigningName, r.Endpoint.SigningName; e != a {
					t.Errorf("expected %v, got %v", e, a)
				}
				if e, a := c.expectedSigningRegion, r.Endpoint.SigningRegion; e != a {
					t.Errorf("expected %v, got %v", e, a)
				}
			})
			_, err := req.Send(context.Background())
			if len(c.expectedErr) == 0 && err != nil {
				t.Errorf("expected no error but got: %v", err)
			} else if len(c.expectedErr) != 0 && err == nil {
				t.Errorf("expected err %q, but got nil", c.expectedErr)
			} else if len(c.expectedErr) != 0 && err != nil && !strings.Contains(err.Error(), c.expectedErr) {
				t.Errorf("expected %v, got %v", c.expectedErr, err.Error())
			}
		})
	}
}
