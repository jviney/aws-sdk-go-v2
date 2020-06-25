// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package codecommit

import (
	"context"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
)

type DescribeMergeConflictsInput struct {
	_ struct{} `type:"structure"`

	// The level of conflict detail to use. If unspecified, the default FILE_LEVEL
	// is used, which returns a not-mergeable result if the same file has differences
	// in both branches. If LINE_LEVEL is specified, a conflict is considered not
	// mergeable if the same file in both branches has differences on the same line.
	ConflictDetailLevel ConflictDetailLevelTypeEnum `locationName:"conflictDetailLevel" type:"string" enum:"true"`

	// Specifies which branch to use when resolving conflicts, or whether to attempt
	// automatically merging two versions of a file. The default is NONE, which
	// requires any conflicts to be resolved manually before the merge operation
	// is successful.
	ConflictResolutionStrategy ConflictResolutionStrategyTypeEnum `locationName:"conflictResolutionStrategy" type:"string" enum:"true"`

	// The branch, tag, HEAD, or other fully qualified reference used to identify
	// a commit (for example, a branch name or a full commit ID).
	//
	// DestinationCommitSpecifier is a required field
	DestinationCommitSpecifier *string `locationName:"destinationCommitSpecifier" type:"string" required:"true"`

	// The path of the target files used to describe the conflicts.
	//
	// FilePath is a required field
	FilePath *string `locationName:"filePath" type:"string" required:"true"`

	// The maximum number of merge hunks to include in the output.
	MaxMergeHunks *int64 `locationName:"maxMergeHunks" type:"integer"`

	// The merge option or strategy you want to use to merge the code.
	//
	// MergeOption is a required field
	MergeOption MergeOptionTypeEnum `locationName:"mergeOption" type:"string" required:"true" enum:"true"`

	// An enumeration token that, when provided in a request, returns the next batch
	// of the results.
	NextToken *string `locationName:"nextToken" type:"string"`

	// The name of the repository where you want to get information about a merge
	// conflict.
	//
	// RepositoryName is a required field
	RepositoryName *string `locationName:"repositoryName" min:"1" type:"string" required:"true"`

	// The branch, tag, HEAD, or other fully qualified reference used to identify
	// a commit (for example, a branch name or a full commit ID).
	//
	// SourceCommitSpecifier is a required field
	SourceCommitSpecifier *string `locationName:"sourceCommitSpecifier" type:"string" required:"true"`
}

// String returns the string representation
func (s DescribeMergeConflictsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeMergeConflictsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeMergeConflictsInput"}

	if s.DestinationCommitSpecifier == nil {
		invalidParams.Add(aws.NewErrParamRequired("DestinationCommitSpecifier"))
	}

	if s.FilePath == nil {
		invalidParams.Add(aws.NewErrParamRequired("FilePath"))
	}
	if len(s.MergeOption) == 0 {
		invalidParams.Add(aws.NewErrParamRequired("MergeOption"))
	}

	if s.RepositoryName == nil {
		invalidParams.Add(aws.NewErrParamRequired("RepositoryName"))
	}
	if s.RepositoryName != nil && len(*s.RepositoryName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("RepositoryName", 1))
	}

	if s.SourceCommitSpecifier == nil {
		invalidParams.Add(aws.NewErrParamRequired("SourceCommitSpecifier"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DescribeMergeConflictsOutput struct {
	_ struct{} `type:"structure"`

	// The commit ID of the merge base.
	BaseCommitId *string `locationName:"baseCommitId" type:"string"`

	// Contains metadata about the conflicts found in the merge.
	//
	// ConflictMetadata is a required field
	ConflictMetadata *ConflictMetadata `locationName:"conflictMetadata" type:"structure" required:"true"`

	// The commit ID of the destination commit specifier that was used in the merge
	// evaluation.
	//
	// DestinationCommitId is a required field
	DestinationCommitId *string `locationName:"destinationCommitId" type:"string" required:"true"`

	// A list of merge hunks of the differences between the files or lines.
	//
	// MergeHunks is a required field
	MergeHunks []MergeHunk `locationName:"mergeHunks" type:"list" required:"true"`

	// An enumeration token that can be used in a request to return the next batch
	// of the results.
	NextToken *string `locationName:"nextToken" type:"string"`

	// The commit ID of the source commit specifier that was used in the merge evaluation.
	//
	// SourceCommitId is a required field
	SourceCommitId *string `locationName:"sourceCommitId" type:"string" required:"true"`
}

// String returns the string representation
func (s DescribeMergeConflictsOutput) String() string {
	return awsutil.Prettify(s)
}

const opDescribeMergeConflicts = "DescribeMergeConflicts"

// DescribeMergeConflictsRequest returns a request value for making API operation for
// AWS CodeCommit.
//
// Returns information about one or more merge conflicts in the attempted merge
// of two commit specifiers using the squash or three-way merge strategy. If
// the merge option for the attempted merge is specified as FAST_FORWARD_MERGE,
// an exception is thrown.
//
//    // Example sending a request using DescribeMergeConflictsRequest.
//    req := client.DescribeMergeConflictsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/codecommit-2015-04-13/DescribeMergeConflicts
func (c *Client) DescribeMergeConflictsRequest(input *DescribeMergeConflictsInput) DescribeMergeConflictsRequest {
	op := &aws.Operation{
		Name:       opDescribeMergeConflicts,
		HTTPMethod: "POST",
		HTTPPath:   "/",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"nextToken"},
			OutputTokens:    []string{"nextToken"},
			LimitToken:      "maxMergeHunks",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &DescribeMergeConflictsInput{}
	}

	req := c.newRequest(op, input, &DescribeMergeConflictsOutput{})

	return DescribeMergeConflictsRequest{Request: req, Input: input, Copy: c.DescribeMergeConflictsRequest}
}

// DescribeMergeConflictsRequest is the request type for the
// DescribeMergeConflicts API operation.
type DescribeMergeConflictsRequest struct {
	*aws.Request
	Input *DescribeMergeConflictsInput
	Copy  func(*DescribeMergeConflictsInput) DescribeMergeConflictsRequest
}

// Send marshals and sends the DescribeMergeConflicts API request.
func (r DescribeMergeConflictsRequest) Send(ctx context.Context) (*DescribeMergeConflictsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeMergeConflictsResponse{
		DescribeMergeConflictsOutput: r.Request.Data.(*DescribeMergeConflictsOutput),
		response:                     &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewDescribeMergeConflictsRequestPaginator returns a paginator for DescribeMergeConflicts.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.DescribeMergeConflictsRequest(input)
//   p := codecommit.NewDescribeMergeConflictsRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewDescribeMergeConflictsPaginator(req DescribeMergeConflictsRequest) DescribeMergeConflictsPaginator {
	return DescribeMergeConflictsPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *DescribeMergeConflictsInput
				if req.Input != nil {
					tmp := *req.Input
					inCpy = &tmp
				}

				newReq := req.Copy(inCpy)
				newReq.SetContext(ctx)
				return newReq.Request, nil
			},
		},
	}
}

// DescribeMergeConflictsPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type DescribeMergeConflictsPaginator struct {
	aws.Pager
}

func (p *DescribeMergeConflictsPaginator) CurrentPage() *DescribeMergeConflictsOutput {
	return p.Pager.CurrentPage().(*DescribeMergeConflictsOutput)
}

// DescribeMergeConflictsResponse is the response type for the
// DescribeMergeConflicts API operation.
type DescribeMergeConflictsResponse struct {
	*DescribeMergeConflictsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeMergeConflicts request.
func (r *DescribeMergeConflictsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
