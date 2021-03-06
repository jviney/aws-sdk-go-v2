package aws_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/aws/defaults"
	"github.com/jviney/aws-sdk-go-v2/internal/awstesting"
	"github.com/jviney/aws-sdk-go-v2/internal/awstesting/unit"
)

func TestRequest_SetContext(t *testing.T) {
	svc := awstesting.NewClient(unit.Config())

	svc.Handlers.Clear()
	svc.Handlers.Send.PushBackNamed(defaults.SendHandler)

	r := svc.NewRequest(&aws.Operation{Name: "Operation"}, nil, nil)
	ctx := &awstesting.FakeContext{DoneCh: make(chan struct{})}
	r.SetContext(ctx)

	ctx.Error = fmt.Errorf("context canceled")
	close(ctx.DoneCh)

	err := r.Send()
	if err == nil {
		t.Fatalf("expected error, got none")
	}

	if e, a := ctx.Error.Error(), err.Error(); !strings.Contains(a, e) {
		t.Errorf("expect %q to be in %q, but was not", e, a)
	}
}

func TestRequest_SetContextPanic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Fatalf("expect SetContext to panic, did not")
		}
	}()
	r := &aws.Request{}

	r.SetContext(nil)
}
