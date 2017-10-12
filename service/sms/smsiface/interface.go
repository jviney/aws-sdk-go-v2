// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package smsiface provides an interface to enable mocking the AWS Server Migration Service service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package smsiface

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/request"
	"github.com/aws/aws-sdk-go-v2/service/sms"
)

// SMSAPI provides an interface to enable mocking the
// sms.SMS service client's API operation,
// paginators, and waiters. This make unit testing your code that calls out
// to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the SDK's request pipeline.
//
//    // myFunc uses an SDK service client to make a request to
//    // AWS Server Migration Service.
//    func myFunc(svc smsiface.SMSAPI) bool {
//        // Make svc.CreateReplicationJob request
//    }
//
//    func main() {
//        sess := session.New()
//        svc := sms.New(sess)
//
//        myFunc(svc)
//    }
//
// In your _test.go file:
//
//    // Define a mock struct to be used in your unit tests of myFunc.
//    type mockSMSClient struct {
//        smsiface.SMSAPI
//    }
//    func (m *mockSMSClient) CreateReplicationJob(input *sms.CreateReplicationJobInput) (*sms.CreateReplicationJobOutput, error) {
//        // mock response/functionality
//    }
//
//    func TestMyFunc(t *testing.T) {
//        // Setup Test
//        mockSvc := &mockSMSClient{}
//
//        myfunc(mockSvc)
//
//        // Verify myFunc's functionality
//    }
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters. Its suggested to use the pattern above for testing, or using
// tooling to generate mocks to satisfy the interfaces.
type SMSAPI interface {
	CreateReplicationJob(*sms.CreateReplicationJobInput) (*sms.CreateReplicationJobOutput, error)
	CreateReplicationJobWithContext(aws.Context, *sms.CreateReplicationJobInput, ...request.Option) (*sms.CreateReplicationJobOutput, error)
	CreateReplicationJobRequest(*sms.CreateReplicationJobInput) (*request.Request, *sms.CreateReplicationJobOutput)

	DeleteReplicationJob(*sms.DeleteReplicationJobInput) (*sms.DeleteReplicationJobOutput, error)
	DeleteReplicationJobWithContext(aws.Context, *sms.DeleteReplicationJobInput, ...request.Option) (*sms.DeleteReplicationJobOutput, error)
	DeleteReplicationJobRequest(*sms.DeleteReplicationJobInput) (*request.Request, *sms.DeleteReplicationJobOutput)

	DeleteServerCatalog(*sms.DeleteServerCatalogInput) (*sms.DeleteServerCatalogOutput, error)
	DeleteServerCatalogWithContext(aws.Context, *sms.DeleteServerCatalogInput, ...request.Option) (*sms.DeleteServerCatalogOutput, error)
	DeleteServerCatalogRequest(*sms.DeleteServerCatalogInput) (*request.Request, *sms.DeleteServerCatalogOutput)

	DisassociateConnector(*sms.DisassociateConnectorInput) (*sms.DisassociateConnectorOutput, error)
	DisassociateConnectorWithContext(aws.Context, *sms.DisassociateConnectorInput, ...request.Option) (*sms.DisassociateConnectorOutput, error)
	DisassociateConnectorRequest(*sms.DisassociateConnectorInput) (*request.Request, *sms.DisassociateConnectorOutput)

	GetConnectors(*sms.GetConnectorsInput) (*sms.GetConnectorsOutput, error)
	GetConnectorsWithContext(aws.Context, *sms.GetConnectorsInput, ...request.Option) (*sms.GetConnectorsOutput, error)
	GetConnectorsRequest(*sms.GetConnectorsInput) (*request.Request, *sms.GetConnectorsOutput)

	GetConnectorsPages(*sms.GetConnectorsInput, func(*sms.GetConnectorsOutput, bool) bool) error
	GetConnectorsPagesWithContext(aws.Context, *sms.GetConnectorsInput, func(*sms.GetConnectorsOutput, bool) bool, ...request.Option) error

	GetReplicationJobs(*sms.GetReplicationJobsInput) (*sms.GetReplicationJobsOutput, error)
	GetReplicationJobsWithContext(aws.Context, *sms.GetReplicationJobsInput, ...request.Option) (*sms.GetReplicationJobsOutput, error)
	GetReplicationJobsRequest(*sms.GetReplicationJobsInput) (*request.Request, *sms.GetReplicationJobsOutput)

	GetReplicationJobsPages(*sms.GetReplicationJobsInput, func(*sms.GetReplicationJobsOutput, bool) bool) error
	GetReplicationJobsPagesWithContext(aws.Context, *sms.GetReplicationJobsInput, func(*sms.GetReplicationJobsOutput, bool) bool, ...request.Option) error

	GetReplicationRuns(*sms.GetReplicationRunsInput) (*sms.GetReplicationRunsOutput, error)
	GetReplicationRunsWithContext(aws.Context, *sms.GetReplicationRunsInput, ...request.Option) (*sms.GetReplicationRunsOutput, error)
	GetReplicationRunsRequest(*sms.GetReplicationRunsInput) (*request.Request, *sms.GetReplicationRunsOutput)

	GetReplicationRunsPages(*sms.GetReplicationRunsInput, func(*sms.GetReplicationRunsOutput, bool) bool) error
	GetReplicationRunsPagesWithContext(aws.Context, *sms.GetReplicationRunsInput, func(*sms.GetReplicationRunsOutput, bool) bool, ...request.Option) error

	GetServers(*sms.GetServersInput) (*sms.GetServersOutput, error)
	GetServersWithContext(aws.Context, *sms.GetServersInput, ...request.Option) (*sms.GetServersOutput, error)
	GetServersRequest(*sms.GetServersInput) (*request.Request, *sms.GetServersOutput)

	GetServersPages(*sms.GetServersInput, func(*sms.GetServersOutput, bool) bool) error
	GetServersPagesWithContext(aws.Context, *sms.GetServersInput, func(*sms.GetServersOutput, bool) bool, ...request.Option) error

	ImportServerCatalog(*sms.ImportServerCatalogInput) (*sms.ImportServerCatalogOutput, error)
	ImportServerCatalogWithContext(aws.Context, *sms.ImportServerCatalogInput, ...request.Option) (*sms.ImportServerCatalogOutput, error)
	ImportServerCatalogRequest(*sms.ImportServerCatalogInput) (*request.Request, *sms.ImportServerCatalogOutput)

	StartOnDemandReplicationRun(*sms.StartOnDemandReplicationRunInput) (*sms.StartOnDemandReplicationRunOutput, error)
	StartOnDemandReplicationRunWithContext(aws.Context, *sms.StartOnDemandReplicationRunInput, ...request.Option) (*sms.StartOnDemandReplicationRunOutput, error)
	StartOnDemandReplicationRunRequest(*sms.StartOnDemandReplicationRunInput) (*request.Request, *sms.StartOnDemandReplicationRunOutput)

	UpdateReplicationJob(*sms.UpdateReplicationJobInput) (*sms.UpdateReplicationJobOutput, error)
	UpdateReplicationJobWithContext(aws.Context, *sms.UpdateReplicationJobInput, ...request.Option) (*sms.UpdateReplicationJobOutput, error)
	UpdateReplicationJobRequest(*sms.UpdateReplicationJobInput) (*request.Request, *sms.UpdateReplicationJobOutput)
}

var _ SMSAPI = (*sms.SMS)(nil)
