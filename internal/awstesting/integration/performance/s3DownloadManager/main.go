// +build integration,perftest

package main

import (
	"context"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"time"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/aws/external"
	"github.com/jviney/aws-sdk-go-v2/internal/awstesting"
	"github.com/jviney/aws-sdk-go-v2/internal/awstesting/integration"
	"github.com/jviney/aws-sdk-go-v2/internal/sdkio"
	"github.com/jviney/aws-sdk-go-v2/service/s3"
	"github.com/jviney/aws-sdk-go-v2/service/s3/s3manager"
)

var config Config

func main() {
	parseCommandLine()

	log.SetOutput(os.Stderr)

	log.Printf("uploading %s file to s3://%s\n", integration.SizeToName(int(config.Size)), config.Bucket)
	key, err := setupDownloadTest(config.Bucket, config.Size)
	if err != nil {
		log.Fatalf("failed to setup download testing: %v", err)
	}

	traces := make(chan *RequestTrace, config.SDK.Concurrency)
	requestTracer := downloadRequestTracer(traces)
	downloader := newDownloader(config.Client, config.SDK, requestTracer)

	metricReportDone := startTraceReceiver(traces)

	log.Println("starting download...")
	start := time.Now()
	_, err = downloader.Download(&awstesting.DiscardAt{}, &s3.GetObjectInput{
		Bucket: &config.Bucket,
		Key:    &key,
	})
	if err != nil {
		log.Fatalf("failed to download object, %v", err)
	}
	close(traces)

	dur := time.Since(start)
	log.Printf("Download finished, Size: %d, Dur: %s, Throughput: %.5f GB/s",
		config.Size, dur, (float64(config.Size)/(float64(dur)/float64(time.Second)))/float64(1e9),
	)

	<-metricReportDone

	log.Printf("cleaning up s3://%s/%s\n", config.Bucket, key)
	if err = teardownDownloadTest(config.Bucket, key); err != nil {
		log.Fatalf("failed to teardwn test artifacts: %v", err)
	}
}

func parseCommandLine() {
	config.SetupFlags("", flag.CommandLine)

	if err := flag.CommandLine.Parse(os.Args[1:]); err != nil {
		flag.CommandLine.PrintDefaults()
		log.Fatalf("failed to parse CLI commands")
	}
	if err := config.Validate(); err != nil {
		flag.CommandLine.PrintDefaults()
		log.Fatalf("invalid arguments: %v", err)
	}
}

func setupDownloadTest(bucket string, size int64) (key string, err error) {
	er := &awstesting.EndlessReader{}
	lr := io.LimitReader(er, size)

	key = integration.UniqueID()

	cfg, err := external.LoadDefaultAWSConfig()
	if err != nil {
		return "", fmt.Errorf("failed to load config: %v", err)
	}

	client := s3.New(cfg)
	client.Disable100Continue = true

	uploader := s3manager.NewUploaderWithClient(client, func(u *s3manager.Uploader) {
		u.PartSize = 100 * sdkio.MebiByte
		u.RequestOptions = append(u.RequestOptions, func(r *aws.Request) {
			if r.Operation.Name != "UploadPart" && r.Operation.Name != "PutObject" {
				return
			}

			r.HTTPRequest.Header.Set("X-Amz-Content-Sha256", "UNSIGNED-PAYLOAD")
		})
	})

	_, err = uploader.Upload(&s3manager.UploadInput{
		Bucket: &bucket,
		Body:   lr,
		Key:    &key,
	})
	if err != nil {
		err = fmt.Errorf("failed to upload test object to s3: %v", err)
	}

	return
}

func teardownDownloadTest(bucket, key string) error {
	cfg, err := external.LoadDefaultAWSConfig()
	if err != nil {
		return fmt.Errorf("failed to load config: %v", err)
	}
	svc := s3.New(cfg)

	resp := svc.DeleteObjectRequest(&s3.DeleteObjectInput{Bucket: &bucket, Key: &key})
	_, err = resp.Send(context.Background())
	return err
}

func startTraceReceiver(traces <-chan *RequestTrace) <-chan struct{} {
	metricReportDone := make(chan struct{})

	go func() {
		defer close(metricReportDone)
		metrics := map[string]*RequestTrace{}
		for trace := range traces {
			curTrace, ok := metrics[trace.Operation]
			if !ok {
				curTrace = trace
			} else {
				curTrace.attempts = append(curTrace.attempts, trace.attempts...)
				if len(trace.errs) != 0 {
					curTrace.errs = append(curTrace.errs, trace.errs...)
				}
				curTrace.finish = trace.finish
			}

			metrics[trace.Operation] = curTrace
		}

		for _, name := range []string{
			"GetObject",
		} {
			if trace, ok := metrics[name]; ok {
				printAttempts(name, trace, config.LogVerbose)
			}
		}
	}()

	return metricReportDone
}

func printAttempts(op string, trace *RequestTrace, verbose bool) {
	if !verbose {
		return
	}

	log.Printf("%s: latency:%s requests:%d errors:%d",
		op,
		trace.finish.Sub(trace.start),
		len(trace.attempts),
		len(trace.errs),
	)

	for _, a := range trace.attempts {
		log.Printf("  * %s", a)
	}
	if err := trace.Err(); err != nil {
		log.Printf("Operation Errors: %v", err)
	}
	log.Println()
}

func downloadRequestTracer(traces chan<- *RequestTrace) aws.Option {
	tracerOption := func(r *aws.Request) {
		id := "op"
		if v, ok := r.Params.(*s3.GetObjectInput); ok {
			if v.Range != nil {
				id = *v.Range
			}
		}
		tracer := NewRequestTrace(r.Context(), r.Operation.Name, id)
		r.SetContext(tracer)

		r.Handlers.Send.PushFront(tracer.OnSendAttempt)
		r.Handlers.CompleteAttempt.PushBack(tracer.OnCompleteAttempt)
		r.Handlers.Complete.PushBack(tracer.OnComplete)
		r.Handlers.Complete.PushBack(func(rr *aws.Request) {
			traces <- tracer
		})
	}

	return tracerOption
}

func newDownloader(clientConfig ClientConfig, sdkConfig SDKConfig, options ...aws.Option) *s3manager.Downloader {
	client := NewClient(clientConfig)

	cfg, err := external.LoadDefaultAWSConfig(aws.Config{HTTPClient: client})
	if err != nil {
		log.Fatalf("failed to load session, %v", err)
	}

	downloader := s3manager.NewDownloader(cfg, func(d *s3manager.Downloader) {
		d.PartSize = sdkConfig.PartSize
		d.Concurrency = sdkConfig.Concurrency
		d.BufferProvider = sdkConfig.BufferProvider

		d.RequestOptions = append(d.RequestOptions, options...)
	})

	return downloader
}
