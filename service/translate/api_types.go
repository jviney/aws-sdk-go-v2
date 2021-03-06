// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package translate

import (
	"time"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
)

var _ aws.Config
var _ = awsutil.Prettify

// The custom terminology applied to the input text by Amazon Translate for
// the translated text response. This is optional in the response and will only
// be present if you specified terminology input in the request. Currently,
// only one terminology can be applied per TranslateText request.
type AppliedTerminology struct {
	_ struct{} `type:"structure"`

	// The name of the custom terminology applied to the input text by Amazon Translate
	// for the translated text response.
	Name *string `min:"1" type:"string"`

	// The specific terms of the custom terminology applied to the input text by
	// Amazon Translate for the translated text response. A maximum of 250 terms
	// will be returned, and the specific terms applied will be the first 250 terms
	// in the source text.
	Terms []Term `type:"list"`
}

// String returns the string representation
func (s AppliedTerminology) String() string {
	return awsutil.Prettify(s)
}

// The encryption key used to encrypt the custom terminologies used by Amazon
// Translate.
type EncryptionKey struct {
	_ struct{} `type:"structure"`

	// The Amazon Resource Name (ARN) of the encryption key being used to encrypt
	// the custom terminology.
	//
	// Id is a required field
	Id *string `min:"1" type:"string" required:"true"`

	// The type of encryption key used by Amazon Translate to encrypt custom terminologies.
	//
	// Type is a required field
	Type EncryptionKeyType `type:"string" required:"true" enum:"true"`
}

// String returns the string representation
func (s EncryptionKey) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *EncryptionKey) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "EncryptionKey"}

	if s.Id == nil {
		invalidParams.Add(aws.NewErrParamRequired("Id"))
	}
	if s.Id != nil && len(*s.Id) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Id", 1))
	}
	if len(s.Type) == 0 {
		invalidParams.Add(aws.NewErrParamRequired("Type"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// The input configuration properties for requesting a batch translation job.
type InputDataConfig struct {
	_ struct{} `type:"structure"`

	// The multipurpose internet mail extension (MIME) type of the input files.
	// Valid values are text/plain for plaintext files and text/html for HTML files.
	//
	// ContentType is a required field
	ContentType *string `type:"string" required:"true"`

	// The URI of the AWS S3 folder that contains the input file. The folder must
	// be in the same Region as the API endpoint you are calling.
	//
	// S3Uri is a required field
	S3Uri *string `type:"string" required:"true"`
}

// String returns the string representation
func (s InputDataConfig) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *InputDataConfig) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "InputDataConfig"}

	if s.ContentType == nil {
		invalidParams.Add(aws.NewErrParamRequired("ContentType"))
	}

	if s.S3Uri == nil {
		invalidParams.Add(aws.NewErrParamRequired("S3Uri"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// The number of documents successfully and unsuccessfully processed during
// a translation job.
type JobDetails struct {
	_ struct{} `type:"structure"`

	// The number of documents that could not be processed during a translation
	// job.
	DocumentsWithErrorsCount *int64 `type:"integer"`

	// The number of documents used as input in a translation job.
	InputDocumentsCount *int64 `type:"integer"`

	// The number of documents successfully processed during a translation job.
	TranslatedDocumentsCount *int64 `type:"integer"`
}

// String returns the string representation
func (s JobDetails) String() string {
	return awsutil.Prettify(s)
}

// The output configuration properties for a batch translation job.
type OutputDataConfig struct {
	_ struct{} `type:"structure"`

	// The URI of the S3 folder that contains a translation job's output file. The
	// folder must be in the same Region as the API endpoint that you are calling.
	//
	// S3Uri is a required field
	S3Uri *string `type:"string" required:"true"`
}

// String returns the string representation
func (s OutputDataConfig) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *OutputDataConfig) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "OutputDataConfig"}

	if s.S3Uri == nil {
		invalidParams.Add(aws.NewErrParamRequired("S3Uri"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// The term being translated by the custom terminology.
type Term struct {
	_ struct{} `type:"structure"`

	// The source text of the term being translated by the custom terminology.
	SourceText *string `type:"string"`

	// The target text of the term being translated by the custom terminology.
	TargetText *string `type:"string"`
}

// String returns the string representation
func (s Term) String() string {
	return awsutil.Prettify(s)
}

// The data associated with the custom terminology.
type TerminologyData struct {
	_ struct{} `type:"structure"`

	// The file containing the custom terminology data. Your version of the AWS
	// SDK performs a Base64-encoding on this field before sending a request to
	// the AWS service. Users of the SDK should not perform Base64-encoding themselves.
	//
	// File is automatically base64 encoded/decoded by the SDK.
	//
	// File is a required field
	File []byte `type:"blob" required:"true" sensitive:"true"`

	// The data format of the custom terminology. Either CSV or TMX.
	//
	// Format is a required field
	Format TerminologyDataFormat `type:"string" required:"true" enum:"true"`
}

// String returns the string representation
func (s TerminologyData) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *TerminologyData) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "TerminologyData"}

	if s.File == nil {
		invalidParams.Add(aws.NewErrParamRequired("File"))
	}
	if len(s.Format) == 0 {
		invalidParams.Add(aws.NewErrParamRequired("Format"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// The location of the custom terminology data.
type TerminologyDataLocation struct {
	_ struct{} `type:"structure"`

	// The location of the custom terminology data.
	//
	// Location is a required field
	Location *string `type:"string" required:"true"`

	// The repository type for the custom terminology data.
	//
	// RepositoryType is a required field
	RepositoryType *string `type:"string" required:"true"`
}

// String returns the string representation
func (s TerminologyDataLocation) String() string {
	return awsutil.Prettify(s)
}

// The properties of the custom terminology.
type TerminologyProperties struct {
	_ struct{} `type:"structure"`

	// The Amazon Resource Name (ARN) of the custom terminology.
	Arn *string `type:"string"`

	// The time at which the custom terminology was created, based on the timestamp.
	CreatedAt *time.Time `type:"timestamp"`

	// The description of the custom terminology properties.
	Description *string `type:"string"`

	// The encryption key for the custom terminology.
	EncryptionKey *EncryptionKey `type:"structure"`

	// The time at which the custom terminology was last update, based on the timestamp.
	LastUpdatedAt *time.Time `type:"timestamp"`

	// The name of the custom terminology.
	Name *string `min:"1" type:"string"`

	// The size of the file used when importing a custom terminology.
	SizeBytes *int64 `type:"integer"`

	// The language code for the source text of the translation request for which
	// the custom terminology is being used.
	SourceLanguageCode *string `min:"2" type:"string"`

	// The language codes for the target languages available with the custom terminology
	// file. All possible target languages are returned in array.
	TargetLanguageCodes []string `type:"list"`

	// The number of terms included in the custom terminology.
	TermCount *int64 `type:"integer"`
}

// String returns the string representation
func (s TerminologyProperties) String() string {
	return awsutil.Prettify(s)
}

// Provides information for filtering a list of translation jobs. For more information,
// see ListTextTranslationJobs.
type TextTranslationJobFilter struct {
	_ struct{} `type:"structure"`

	// Filters the list of jobs by name.
	JobName *string `min:"1" type:"string"`

	// Filters the list of jobs based by job status.
	JobStatus JobStatus `type:"string" enum:"true"`

	// Filters the list of jobs based on the time that the job was submitted for
	// processing and returns only the jobs submitted after the specified time.
	// Jobs are returned in descending order, newest to oldest.
	SubmittedAfterTime *time.Time `type:"timestamp"`

	// Filters the list of jobs based on the time that the job was submitted for
	// processing and returns only the jobs submitted before the specified time.
	// Jobs are returned in ascending order, oldest to newest.
	SubmittedBeforeTime *time.Time `type:"timestamp"`
}

// String returns the string representation
func (s TextTranslationJobFilter) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *TextTranslationJobFilter) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "TextTranslationJobFilter"}
	if s.JobName != nil && len(*s.JobName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("JobName", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Provides information about a translation job.
type TextTranslationJobProperties struct {
	_ struct{} `type:"structure"`

	// The Amazon Resource Name (ARN) of an AWS Identity Access and Management (IAM)
	// role that granted Amazon Translate read access to the job's input data.
	DataAccessRoleArn *string `min:"20" type:"string"`

	// The time at which the translation job ended.
	EndTime *time.Time `type:"timestamp"`

	// The input configuration properties that were specified when the job was requested.
	InputDataConfig *InputDataConfig `type:"structure"`

	// The number of documents successfully and unsuccessfully processed during
	// the translation job.
	JobDetails *JobDetails `type:"structure"`

	// The ID of the translation job.
	JobId *string `min:"1" type:"string"`

	// The user-defined name of the translation job.
	JobName *string `min:"1" type:"string"`

	// The status of the translation job.
	JobStatus JobStatus `type:"string" enum:"true"`

	// An explanation of any errors that may have occured during the translation
	// job.
	Message *string `type:"string"`

	// The output configuration properties that were specified when the job was
	// requested.
	OutputDataConfig *OutputDataConfig `type:"structure"`

	// The language code of the language of the source text. The language must be
	// a language supported by Amazon Translate.
	SourceLanguageCode *string `min:"2" type:"string"`

	// The time at which the translation job was submitted.
	SubmittedTime *time.Time `type:"timestamp"`

	// The language code of the language of the target text. The language must be
	// a language supported by Amazon Translate.
	TargetLanguageCodes []string `min:"1" type:"list"`

	// A list containing the names of the terminologies applied to a translation
	// job. Only one terminology can be applied per StartTextTranslationJob request
	// at this time.
	TerminologyNames []string `type:"list"`
}

// String returns the string representation
func (s TextTranslationJobProperties) String() string {
	return awsutil.Prettify(s)
}
