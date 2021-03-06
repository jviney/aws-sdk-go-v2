// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package macie

import (
	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
)

var _ aws.Config
var _ = awsutil.Prettify

// The classification type that Amazon Macie Classic applies to the associated
// S3 resources.
type ClassificationType struct {
	_ struct{} `type:"structure"`

	// A continuous classification of the objects that are added to a specified
	// S3 bucket. Amazon Macie Classic begins performing continuous classification
	// after a bucket is successfully associated with Amazon Macie Classic.
	//
	// Continuous is a required field
	Continuous S3ContinuousClassificationType `locationName:"continuous" type:"string" required:"true" enum:"true"`

	// A one-time classification of all of the existing objects in a specified S3
	// bucket.
	//
	// OneTime is a required field
	OneTime S3OneTimeClassificationType `locationName:"oneTime" type:"string" required:"true" enum:"true"`
}

// String returns the string representation
func (s ClassificationType) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ClassificationType) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ClassificationType"}
	if len(s.Continuous) == 0 {
		invalidParams.Add(aws.NewErrParamRequired("Continuous"))
	}
	if len(s.OneTime) == 0 {
		invalidParams.Add(aws.NewErrParamRequired("OneTime"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// The classification type that Amazon Macie Classic applies to the associated
// S3 resources. At least one of the classification types (oneTime or continuous)
// must be specified.
type ClassificationTypeUpdate struct {
	_ struct{} `type:"structure"`

	// A continuous classification of the objects that are added to a specified
	// S3 bucket. Amazon Macie Classic begins performing continuous classification
	// after a bucket is successfully associated with Amazon Macie Classic.
	Continuous S3ContinuousClassificationType `locationName:"continuous" type:"string" enum:"true"`

	// A one-time classification of all of the existing objects in a specified S3
	// bucket.
	OneTime S3OneTimeClassificationType `locationName:"oneTime" type:"string" enum:"true"`
}

// String returns the string representation
func (s ClassificationTypeUpdate) String() string {
	return awsutil.Prettify(s)
}

// Includes details about the failed S3 resources.
type FailedS3Resource struct {
	_ struct{} `type:"structure"`

	// The status code of a failed item.
	ErrorCode *string `locationName:"errorCode" type:"string"`

	// The error message of a failed item.
	ErrorMessage *string `locationName:"errorMessage" type:"string"`

	// The failed S3 resources.
	FailedItem *S3Resource `locationName:"failedItem" type:"structure"`
}

// String returns the string representation
func (s FailedS3Resource) String() string {
	return awsutil.Prettify(s)
}

// Contains information about the Amazon Macie Classic member account.
type MemberAccount struct {
	_ struct{} `type:"structure"`

	// The AWS account ID of the Amazon Macie Classic member account.
	AccountId *string `locationName:"accountId" type:"string"`
}

// String returns the string representation
func (s MemberAccount) String() string {
	return awsutil.Prettify(s)
}

// Contains information about the S3 resource. This data type is used as a request
// parameter in the DisassociateS3Resources action and can be used as a response
// parameter in the AssociateS3Resources and UpdateS3Resources actions.
type S3Resource struct {
	_ struct{} `type:"structure"`

	// The name of the S3 bucket.
	//
	// BucketName is a required field
	BucketName *string `locationName:"bucketName" type:"string" required:"true"`

	// The prefix of the S3 bucket.
	Prefix *string `locationName:"prefix" type:"string"`
}

// String returns the string representation
func (s S3Resource) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *S3Resource) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "S3Resource"}

	if s.BucketName == nil {
		invalidParams.Add(aws.NewErrParamRequired("BucketName"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// The S3 resources that you want to associate with Amazon Macie Classic for
// monitoring and data classification. This data type is used as a request parameter
// in the AssociateS3Resources action and a response parameter in the ListS3Resources
// action.
type S3ResourceClassification struct {
	_ struct{} `type:"structure"`

	// The name of the S3 bucket that you want to associate with Amazon Macie Classic.
	//
	// BucketName is a required field
	BucketName *string `locationName:"bucketName" type:"string" required:"true"`

	// The classification type that you want to specify for the resource associated
	// with Amazon Macie Classic.
	//
	// ClassificationType is a required field
	ClassificationType *ClassificationType `locationName:"classificationType" type:"structure" required:"true"`

	// The prefix of the S3 bucket that you want to associate with Amazon Macie
	// Classic.
	Prefix *string `locationName:"prefix" type:"string"`
}

// String returns the string representation
func (s S3ResourceClassification) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *S3ResourceClassification) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "S3ResourceClassification"}

	if s.BucketName == nil {
		invalidParams.Add(aws.NewErrParamRequired("BucketName"))
	}

	if s.ClassificationType == nil {
		invalidParams.Add(aws.NewErrParamRequired("ClassificationType"))
	}
	if s.ClassificationType != nil {
		if err := s.ClassificationType.Validate(); err != nil {
			invalidParams.AddNested("ClassificationType", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// The S3 resources whose classification types you want to update. This data
// type is used as a request parameter in the UpdateS3Resources action.
type S3ResourceClassificationUpdate struct {
	_ struct{} `type:"structure"`

	// The name of the S3 bucket whose classification types you want to update.
	//
	// BucketName is a required field
	BucketName *string `locationName:"bucketName" type:"string" required:"true"`

	// The classification type that you want to update for the resource associated
	// with Amazon Macie Classic.
	//
	// ClassificationTypeUpdate is a required field
	ClassificationTypeUpdate *ClassificationTypeUpdate `locationName:"classificationTypeUpdate" type:"structure" required:"true"`

	// The prefix of the S3 bucket whose classification types you want to update.
	Prefix *string `locationName:"prefix" type:"string"`
}

// String returns the string representation
func (s S3ResourceClassificationUpdate) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *S3ResourceClassificationUpdate) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "S3ResourceClassificationUpdate"}

	if s.BucketName == nil {
		invalidParams.Add(aws.NewErrParamRequired("BucketName"))
	}

	if s.ClassificationTypeUpdate == nil {
		invalidParams.Add(aws.NewErrParamRequired("ClassificationTypeUpdate"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}
