package machinelearning_test

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/jviney/aws-sdk-go-v2/aws"
	request "github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awstesting/unit"
	"github.com/jviney/aws-sdk-go-v2/service/machinelearning"
)

func TestPredictEndpoint(t *testing.T) {
	ml := machinelearning.New(unit.Config())
	ml.Handlers.Send.Clear()
	ml.Handlers.Send.PushBack(func(r *request.Request) {
		r.HTTPResponse = &http.Response{
			StatusCode: 200,
			Header:     http.Header{},
			Body:       ioutil.NopCloser(bytes.NewReader([]byte("{}"))),
		}
	})

	req := ml.PredictRequest(&machinelearning.PredictInput{
		PredictEndpoint: aws.String("https://localhost/endpoint"),
		MLModelId:       aws.String("id"),
		Record:          map[string]string{},
	})
	_, err := req.Send(context.Background())

	if err != nil {
		t.Errorf("expect no error, got %v", err)
	}
	if e, a := "https://localhost/endpoint", req.HTTPRequest.URL.String(); e != a {
		t.Errorf("expect %v, got %v", e, a)
	}
}
