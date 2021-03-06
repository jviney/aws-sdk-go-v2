package sqs_test

import (
	"bytes"
	"context"
	"errors"
	"io/ioutil"
	"net/http"
	"strings"
	"testing"

	"github.com/jviney/aws-sdk-go-v2/aws"
	request "github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/aws/defaults"
	"github.com/jviney/aws-sdk-go-v2/internal/awstesting/unit"
	"github.com/jviney/aws-sdk-go-v2/service/sqs"
)

var svc = func() *sqs.Client {
	cfg := unit.Config()
	cfg.Handlers.Validate.Remove(defaults.ValidateParametersHandler)

	s := sqs.New(cfg)
	s.Handlers.Send.Clear()
	return s
}()

func TestSendMessageChecksum(t *testing.T) {
	req := svc.SendMessageRequest(&sqs.SendMessageInput{
		MessageBody: aws.String("test"),
	})
	req.Handlers.Send.PushBack(func(r *request.Request) {
		body := ioutil.NopCloser(bytes.NewReader([]byte("")))
		r.HTTPResponse = &http.Response{StatusCode: 200, Body: body}
		r.Data = &sqs.SendMessageOutput{
			MD5OfMessageBody: aws.String("098f6bcd4621d373cade4e832627b4f6"),
			MessageId:        aws.String("12345"),
		}
	})
	_, err := req.Send(context.Background())
	if err != nil {
		t.Errorf("expect no error, got %v", err)
	}
}

func TestSendMessageChecksumInvalid(t *testing.T) {
	req := svc.SendMessageRequest(&sqs.SendMessageInput{
		MessageBody: aws.String("test"),
	})
	req.Handlers.Send.PushBack(func(r *request.Request) {
		body := ioutil.NopCloser(bytes.NewReader([]byte("")))
		r.HTTPResponse = &http.Response{StatusCode: 200, Body: body}
		r.Data = &sqs.SendMessageOutput{
			MD5OfMessageBody: aws.String("000"),
			MessageId:        aws.String("12345"),
		}
	})
	_, err := req.Send(context.Background())
	if err == nil {
		t.Fatalf("expect error, got nil")
	}

	var ce *sqs.InvalidChecksumError
	if !errors.As(err, &ce) {
		t.Fatalf("expect %T error, got %v", ce, err)
	}
	expectReason := "expected MD5 checksum '000', got '098f6bcd4621d373cade4e832627b4f6'"
	if e, a := expectReason, ce.Error(); !strings.Contains(a, e) {
		t.Errorf("expect %v to be in %v, was not", e, a)
	}
}

func TestSendMessageChecksumInvalidNoValidation(t *testing.T) {
	cfg := unit.Config()
	cfg.Handlers.Validate.Remove(defaults.ValidateParametersHandler)

	s := sqs.New(cfg)
	s.DisableComputeChecksums = true
	s.Handlers.Send.Clear()

	req := s.SendMessageRequest(&sqs.SendMessageInput{
		MessageBody: aws.String("test"),
	})
	req.Handlers.Send.PushBack(func(r *request.Request) {
		body := ioutil.NopCloser(bytes.NewReader([]byte("")))
		r.HTTPResponse = &http.Response{StatusCode: 200, Body: body}
		r.Data = &sqs.SendMessageOutput{
			MD5OfMessageBody: aws.String("000"),
			MessageId:        aws.String("12345"),
		}
	})
	_, err := req.Send(context.Background())
	if err != nil {
		t.Errorf("expect no error, got %v", err)
	}
}

func TestSendMessageChecksumNoInput(t *testing.T) {
	req := svc.SendMessageRequest(&sqs.SendMessageInput{})
	req.Handlers.Send.PushBack(func(r *request.Request) {
		body := ioutil.NopCloser(bytes.NewReader([]byte("")))
		r.HTTPResponse = &http.Response{StatusCode: 200, Body: body}
		r.Data = &sqs.SendMessageOutput{}
	})
	_, err := req.Send(context.Background())
	if err == nil {
		t.Fatalf("expect error, got nil")
	}

	var ce *sqs.InvalidChecksumError
	if !errors.As(err, &ce) {
		t.Fatalf("expect %T error, got %v", ce, err)
	}
	expectReason := "cannot compute checksum. missing body"
	if e, a := expectReason, ce.Error(); !strings.Contains(a, e) {
		t.Errorf("expect %v to be in %v, was not", e, a)
	}
}

func TestSendMessageChecksumNoOutput(t *testing.T) {
	req := svc.SendMessageRequest(&sqs.SendMessageInput{
		MessageBody: aws.String("test"),
	})
	req.Handlers.Send.PushBack(func(r *request.Request) {
		body := ioutil.NopCloser(bytes.NewReader([]byte("")))
		r.HTTPResponse = &http.Response{StatusCode: 200, Body: body}
		r.Data = &sqs.SendMessageOutput{}
	})
	_, err := req.Send(context.Background())
	if err == nil {
		t.Fatalf("expect error, got nil")
	}

	var ce *sqs.InvalidChecksumError
	if !errors.As(err, &ce) {
		t.Fatalf("expect %T error, got %v", ce, err)
	}
	expectReason := "cannot verify checksum. missing response MD5"
	if e, a := expectReason, ce.Error(); !strings.Contains(a, e) {
		t.Errorf("expect %v to be in %v, was not", e, a)
	}
}

func TestRecieveMessageChecksum(t *testing.T) {
	req := svc.ReceiveMessageRequest(&sqs.ReceiveMessageInput{})
	req.Handlers.Send.PushBack(func(r *request.Request) {
		md5 := "098f6bcd4621d373cade4e832627b4f6"
		body := ioutil.NopCloser(bytes.NewReader([]byte("")))
		r.HTTPResponse = &http.Response{StatusCode: 200, Body: body}
		r.Data = &sqs.ReceiveMessageOutput{
			Messages: []sqs.Message{
				{Body: aws.String("test"), MD5OfBody: &md5},
				{Body: aws.String("test"), MD5OfBody: &md5},
				{Body: aws.String("test"), MD5OfBody: &md5},
				{Body: aws.String("test"), MD5OfBody: &md5},
			},
		}
	})
	_, err := req.Send(context.Background())
	if err != nil {
		t.Errorf("expect no error, got %v", err)
	}
}

func TestRecieveMessageChecksumInvalid(t *testing.T) {
	req := svc.ReceiveMessageRequest(&sqs.ReceiveMessageInput{})
	req.Handlers.Send.PushBack(func(r *request.Request) {
		md5 := "098f6bcd4621d373cade4e832627b4f6"
		body := ioutil.NopCloser(bytes.NewReader([]byte("")))
		r.HTTPResponse = &http.Response{StatusCode: 200, Body: body}
		r.Data = &sqs.ReceiveMessageOutput{
			Messages: []sqs.Message{
				{Body: aws.String("test"), MD5OfBody: &md5},
				{Body: aws.String("test"), MD5OfBody: aws.String("000"), MessageId: aws.String("123")},
				{Body: aws.String("test"), MD5OfBody: aws.String("000"), MessageId: aws.String("456")},
				{Body: aws.String("test"), MD5OfBody: aws.String("000")},
				{Body: aws.String("test"), MD5OfBody: &md5},
			},
		}
	})
	_, err := req.Send(context.Background())
	if err == nil {
		t.Fatalf("expect error, got nil")
	}

	var ce *sqs.InvalidChecksumError
	if !errors.As(err, &ce) {
		t.Fatalf("expect %T error, got %v", ce, err)
	}
	expectReason := "invalid messages: 123, 456"
	if e, a := expectReason, ce.Error(); !strings.Contains(a, e) {
		t.Errorf("expect %v to be in %v, was not", e, a)
	}
}

func TestSendMessageBatchChecksum(t *testing.T) {
	req := svc.SendMessageBatchRequest(&sqs.SendMessageBatchInput{
		Entries: []sqs.SendMessageBatchRequestEntry{
			{Id: aws.String("1"), MessageBody: aws.String("test")},
			{Id: aws.String("2"), MessageBody: aws.String("test")},
			{Id: aws.String("3"), MessageBody: aws.String("test")},
			{Id: aws.String("4"), MessageBody: aws.String("test")},
		},
	})
	req.Handlers.Send.PushBack(func(r *request.Request) {
		md5 := "098f6bcd4621d373cade4e832627b4f6"
		body := ioutil.NopCloser(bytes.NewReader([]byte("")))
		r.HTTPResponse = &http.Response{StatusCode: 200, Body: body}
		r.Data = &sqs.SendMessageBatchOutput{
			Successful: []sqs.SendMessageBatchResultEntry{
				{MD5OfMessageBody: &md5, MessageId: aws.String("123"), Id: aws.String("1")},
				{MD5OfMessageBody: &md5, MessageId: aws.String("456"), Id: aws.String("2")},
				{MD5OfMessageBody: &md5, MessageId: aws.String("789"), Id: aws.String("3")},
				{MD5OfMessageBody: &md5, MessageId: aws.String("012"), Id: aws.String("4")},
			},
		}
	})
	_, err := req.Send(context.Background())
	if err != nil {
		t.Errorf("expect no error, got %v", err)
	}
}

func TestSendMessageBatchChecksumFailed(t *testing.T) {
	req := svc.SendMessageBatchRequest(&sqs.SendMessageBatchInput{
		Entries: []sqs.SendMessageBatchRequestEntry{
			{Id: aws.String("1"), MessageBody: aws.String("test")},
			{Id: aws.String("2"), MessageBody: aws.String("test")},
			{Id: aws.String("3"), MessageBody: aws.String("test")},
			{Id: aws.String("4"), MessageBody: aws.String("test")},
		},
	})
	req.Handlers.Send.PushBack(func(r *request.Request) {
		body := ioutil.NopCloser(bytes.NewReader([]byte("")))
		r.HTTPResponse = &http.Response{StatusCode: 200, Body: body}
		r.Data = &sqs.SendMessageBatchOutput{
			Failed: []sqs.BatchResultErrorEntry{
				{Id: aws.String("1"), Code: aws.String("test"), Message: aws.String("test"), SenderFault: aws.Bool(false)},
				{Id: aws.String("2"), Code: aws.String("test"), Message: aws.String("test"), SenderFault: aws.Bool(false)},
				{Id: aws.String("3"), Code: aws.String("test"), Message: aws.String("test"), SenderFault: aws.Bool(false)},
				{Id: aws.String("4"), Code: aws.String("test"), Message: aws.String("test"), SenderFault: aws.Bool(false)},
			},
		}
	})
	_, err := req.Send(context.Background())
	if err != nil {
		t.Errorf("expect no error, got %v", err)
	}
}

func TestSendMessageBatchChecksumInvalid(t *testing.T) {
	req := svc.SendMessageBatchRequest(&sqs.SendMessageBatchInput{
		Entries: []sqs.SendMessageBatchRequestEntry{
			{Id: aws.String("1"), MessageBody: aws.String("test")},
			{Id: aws.String("2"), MessageBody: aws.String("test")},
			{Id: aws.String("3"), MessageBody: aws.String("test")},
			{Id: aws.String("4"), MessageBody: aws.String("test")},
		},
	})
	req.Handlers.Send.PushBack(func(r *request.Request) {
		md5 := "098f6bcd4621d373cade4e832627b4f6"
		body := ioutil.NopCloser(bytes.NewReader([]byte("")))
		r.HTTPResponse = &http.Response{StatusCode: 200, Body: body}
		r.Data = &sqs.SendMessageBatchOutput{
			Successful: []sqs.SendMessageBatchResultEntry{
				{MD5OfMessageBody: &md5, MessageId: aws.String("123"), Id: aws.String("1")},
				{MD5OfMessageBody: aws.String("000"), MessageId: aws.String("456"), Id: aws.String("2")},
				{MD5OfMessageBody: aws.String("000"), MessageId: aws.String("789"), Id: aws.String("3")},
				{MD5OfMessageBody: &md5, MessageId: aws.String("012"), Id: aws.String("4")},
			},
		}
	})
	_, err := req.Send(context.Background())
	if err == nil {
		t.Fatalf("expect error, got nil")
	}

	var ce *sqs.InvalidChecksumError
	if !errors.As(err, &ce) {
		t.Fatalf("expect %T error, got %v", ce, err)
	}
	expectReason := "invalid messages: 456, 789"
	if e, a := expectReason, ce.Error(); !strings.Contains(a, e) {
		t.Errorf("expect %v to be in %v, was not", e, a)
	}
}
