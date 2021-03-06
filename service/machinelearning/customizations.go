package machinelearning

import (
	"net/url"

	"github.com/jviney/aws-sdk-go-v2/aws"
)

func init() {
	initRequest = func(c *Client, r *aws.Request) {
		switch r.Operation.Name {
		case opPredict:
			r.Handlers.Build.PushBack(updatePredictEndpoint)
		}
	}
}

// updatePredictEndpoint rewrites the request endpoint to use the
// "PredictEndpoint" parameter of the Predict operation.
func updatePredictEndpoint(r *aws.Request) {
	if !r.ParamsFilled() {
		return
	}

	r.Endpoint.URL = *r.Params.(*PredictInput).PredictEndpoint

	uri, err := url.Parse(r.Endpoint.URL)
	if err != nil {
		r.Error = err
		return
	}
	r.HTTPRequest.URL = uri
}
