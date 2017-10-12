// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package athenaiface provides an interface to enable mocking the Amazon Athena service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package athenaiface

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/request"
	"github.com/aws/aws-sdk-go-v2/service/athena"
)

// AthenaAPI provides an interface to enable mocking the
// athena.Athena service client's API operation,
// paginators, and waiters. This make unit testing your code that calls out
// to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the SDK's request pipeline.
//
//    // myFunc uses an SDK service client to make a request to
//    // Amazon Athena.
//    func myFunc(svc athenaiface.AthenaAPI) bool {
//        // Make svc.BatchGetNamedQuery request
//    }
//
//    func main() {
//        sess := session.New()
//        svc := athena.New(sess)
//
//        myFunc(svc)
//    }
//
// In your _test.go file:
//
//    // Define a mock struct to be used in your unit tests of myFunc.
//    type mockAthenaClient struct {
//        athenaiface.AthenaAPI
//    }
//    func (m *mockAthenaClient) BatchGetNamedQuery(input *athena.BatchGetNamedQueryInput) (*athena.BatchGetNamedQueryOutput, error) {
//        // mock response/functionality
//    }
//
//    func TestMyFunc(t *testing.T) {
//        // Setup Test
//        mockSvc := &mockAthenaClient{}
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
type AthenaAPI interface {
	BatchGetNamedQuery(*athena.BatchGetNamedQueryInput) (*athena.BatchGetNamedQueryOutput, error)
	BatchGetNamedQueryWithContext(aws.Context, *athena.BatchGetNamedQueryInput, ...request.Option) (*athena.BatchGetNamedQueryOutput, error)
	BatchGetNamedQueryRequest(*athena.BatchGetNamedQueryInput) (*request.Request, *athena.BatchGetNamedQueryOutput)

	BatchGetQueryExecution(*athena.BatchGetQueryExecutionInput) (*athena.BatchGetQueryExecutionOutput, error)
	BatchGetQueryExecutionWithContext(aws.Context, *athena.BatchGetQueryExecutionInput, ...request.Option) (*athena.BatchGetQueryExecutionOutput, error)
	BatchGetQueryExecutionRequest(*athena.BatchGetQueryExecutionInput) (*request.Request, *athena.BatchGetQueryExecutionOutput)

	CreateNamedQuery(*athena.CreateNamedQueryInput) (*athena.CreateNamedQueryOutput, error)
	CreateNamedQueryWithContext(aws.Context, *athena.CreateNamedQueryInput, ...request.Option) (*athena.CreateNamedQueryOutput, error)
	CreateNamedQueryRequest(*athena.CreateNamedQueryInput) (*request.Request, *athena.CreateNamedQueryOutput)

	DeleteNamedQuery(*athena.DeleteNamedQueryInput) (*athena.DeleteNamedQueryOutput, error)
	DeleteNamedQueryWithContext(aws.Context, *athena.DeleteNamedQueryInput, ...request.Option) (*athena.DeleteNamedQueryOutput, error)
	DeleteNamedQueryRequest(*athena.DeleteNamedQueryInput) (*request.Request, *athena.DeleteNamedQueryOutput)

	GetNamedQuery(*athena.GetNamedQueryInput) (*athena.GetNamedQueryOutput, error)
	GetNamedQueryWithContext(aws.Context, *athena.GetNamedQueryInput, ...request.Option) (*athena.GetNamedQueryOutput, error)
	GetNamedQueryRequest(*athena.GetNamedQueryInput) (*request.Request, *athena.GetNamedQueryOutput)

	GetQueryExecution(*athena.GetQueryExecutionInput) (*athena.GetQueryExecutionOutput, error)
	GetQueryExecutionWithContext(aws.Context, *athena.GetQueryExecutionInput, ...request.Option) (*athena.GetQueryExecutionOutput, error)
	GetQueryExecutionRequest(*athena.GetQueryExecutionInput) (*request.Request, *athena.GetQueryExecutionOutput)

	GetQueryResults(*athena.GetQueryResultsInput) (*athena.GetQueryResultsOutput, error)
	GetQueryResultsWithContext(aws.Context, *athena.GetQueryResultsInput, ...request.Option) (*athena.GetQueryResultsOutput, error)
	GetQueryResultsRequest(*athena.GetQueryResultsInput) (*request.Request, *athena.GetQueryResultsOutput)

	GetQueryResultsPages(*athena.GetQueryResultsInput, func(*athena.GetQueryResultsOutput, bool) bool) error
	GetQueryResultsPagesWithContext(aws.Context, *athena.GetQueryResultsInput, func(*athena.GetQueryResultsOutput, bool) bool, ...request.Option) error

	ListNamedQueries(*athena.ListNamedQueriesInput) (*athena.ListNamedQueriesOutput, error)
	ListNamedQueriesWithContext(aws.Context, *athena.ListNamedQueriesInput, ...request.Option) (*athena.ListNamedQueriesOutput, error)
	ListNamedQueriesRequest(*athena.ListNamedQueriesInput) (*request.Request, *athena.ListNamedQueriesOutput)

	ListNamedQueriesPages(*athena.ListNamedQueriesInput, func(*athena.ListNamedQueriesOutput, bool) bool) error
	ListNamedQueriesPagesWithContext(aws.Context, *athena.ListNamedQueriesInput, func(*athena.ListNamedQueriesOutput, bool) bool, ...request.Option) error

	ListQueryExecutions(*athena.ListQueryExecutionsInput) (*athena.ListQueryExecutionsOutput, error)
	ListQueryExecutionsWithContext(aws.Context, *athena.ListQueryExecutionsInput, ...request.Option) (*athena.ListQueryExecutionsOutput, error)
	ListQueryExecutionsRequest(*athena.ListQueryExecutionsInput) (*request.Request, *athena.ListQueryExecutionsOutput)

	ListQueryExecutionsPages(*athena.ListQueryExecutionsInput, func(*athena.ListQueryExecutionsOutput, bool) bool) error
	ListQueryExecutionsPagesWithContext(aws.Context, *athena.ListQueryExecutionsInput, func(*athena.ListQueryExecutionsOutput, bool) bool, ...request.Option) error

	StartQueryExecution(*athena.StartQueryExecutionInput) (*athena.StartQueryExecutionOutput, error)
	StartQueryExecutionWithContext(aws.Context, *athena.StartQueryExecutionInput, ...request.Option) (*athena.StartQueryExecutionOutput, error)
	StartQueryExecutionRequest(*athena.StartQueryExecutionInput) (*request.Request, *athena.StartQueryExecutionOutput)

	StopQueryExecution(*athena.StopQueryExecutionInput) (*athena.StopQueryExecutionOutput, error)
	StopQueryExecutionWithContext(aws.Context, *athena.StopQueryExecutionInput, ...request.Option) (*athena.StopQueryExecutionOutput, error)
	StopQueryExecutionRequest(*athena.StopQueryExecutionInput) (*request.Request, *athena.StopQueryExecutionOutput)
}

var _ AthenaAPI = (*athena.Athena)(nil)
