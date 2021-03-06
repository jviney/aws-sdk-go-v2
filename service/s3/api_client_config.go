// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package s3

import (
	"github.com/jviney/aws-sdk-go-v2/service/s3/internal/external"
)

func resolveClientConfig(svc *Client, configs []interface{}) error {
	if len(configs) == 0 {
		return nil
	}

	if value, ok, err := external.ResolveUseARNRegion(configs); err != nil {
		return err
	} else if ok {
		svc.UseARNRegion = value
	}

	return nil
}
