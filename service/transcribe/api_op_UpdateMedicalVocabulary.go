// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package transcribe

import (
	"context"
	"time"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
)

type UpdateMedicalVocabularyInput struct {
	_ struct{} `type:"structure"`

	// The language code of the entries in the updated vocabulary. US English (en-US)
	// is the only valid language code in Amazon Transcribe Medical.
	//
	// LanguageCode is a required field
	LanguageCode LanguageCode `type:"string" required:"true" enum:"true"`

	// The Amazon S3 location of the text file containing the definition of the
	// custom vocabulary. The URI must be in the same AWS region as the API endpoint
	// you are calling. You can see the fields you need to enter for you Amazon
	// S3 location in the example URI here:
	//
	// https://s3.<aws-region>.amazonaws.com/<bucket-name>/<keyprefix>/<objectkey>
	//
	// For example:
	//
	// https://s3.us-east-1.amazonaws.com/AWSDOC-EXAMPLE-BUCKET/vocab.txt
	//
	// For more information about S3 object names, see Object Keys (http://docs.aws.amazon.com/AmazonS3/latest/dev/UsingMetadata.html#object-keys)
	// in the Amazon S3 Developer Guide.
	//
	// For more information about custom vocabularies in Amazon Transcribe Medical,
	// see Medical Custom Vocabularies (http://docs.aws.amazon.com/transcribe/latest/dg/how-it-works.html#how-vocabulary).
	VocabularyFileUri *string `min:"1" type:"string"`

	// The name of the vocabulary to update. The name is case-sensitive. If you
	// try to update a vocabulary with the same name as a previous vocabulary you
	// will receive a ConflictException error.
	//
	// VocabularyName is a required field
	VocabularyName *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s UpdateMedicalVocabularyInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UpdateMedicalVocabularyInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "UpdateMedicalVocabularyInput"}
	if len(s.LanguageCode) == 0 {
		invalidParams.Add(aws.NewErrParamRequired("LanguageCode"))
	}
	if s.VocabularyFileUri != nil && len(*s.VocabularyFileUri) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("VocabularyFileUri", 1))
	}

	if s.VocabularyName == nil {
		invalidParams.Add(aws.NewErrParamRequired("VocabularyName"))
	}
	if s.VocabularyName != nil && len(*s.VocabularyName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("VocabularyName", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type UpdateMedicalVocabularyOutput struct {
	_ struct{} `type:"structure"`

	// The language code for the text file used to update the custom vocabulary.
	// US English (en-US) is the only language supported in Amazon Transcribe Medical.
	LanguageCode LanguageCode `type:"string" enum:"true"`

	// The date and time the vocabulary was updated.
	LastModifiedTime *time.Time `type:"timestamp"`

	// The name of the updated vocabulary.
	VocabularyName *string `min:"1" type:"string"`

	// The processing state of the update to the vocabulary. When the VocabularyState
	// field is READY the vocabulary is ready to be used in a StartMedicalTranscriptionJob
	// request.
	VocabularyState VocabularyState `type:"string" enum:"true"`
}

// String returns the string representation
func (s UpdateMedicalVocabularyOutput) String() string {
	return awsutil.Prettify(s)
}

const opUpdateMedicalVocabulary = "UpdateMedicalVocabulary"

// UpdateMedicalVocabularyRequest returns a request value for making API operation for
// Amazon Transcribe Service.
//
// Updates an existing vocabulary with new values in a different text file.
// The UpdateMedicalVocabulary operation overwrites all of the existing information
// with the values that you provide in the request.
//
//    // Example sending a request using UpdateMedicalVocabularyRequest.
//    req := client.UpdateMedicalVocabularyRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/transcribe-2017-10-26/UpdateMedicalVocabulary
func (c *Client) UpdateMedicalVocabularyRequest(input *UpdateMedicalVocabularyInput) UpdateMedicalVocabularyRequest {
	op := &aws.Operation{
		Name:       opUpdateMedicalVocabulary,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &UpdateMedicalVocabularyInput{}
	}

	req := c.newRequest(op, input, &UpdateMedicalVocabularyOutput{})

	return UpdateMedicalVocabularyRequest{Request: req, Input: input, Copy: c.UpdateMedicalVocabularyRequest}
}

// UpdateMedicalVocabularyRequest is the request type for the
// UpdateMedicalVocabulary API operation.
type UpdateMedicalVocabularyRequest struct {
	*aws.Request
	Input *UpdateMedicalVocabularyInput
	Copy  func(*UpdateMedicalVocabularyInput) UpdateMedicalVocabularyRequest
}

// Send marshals and sends the UpdateMedicalVocabulary API request.
func (r UpdateMedicalVocabularyRequest) Send(ctx context.Context) (*UpdateMedicalVocabularyResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &UpdateMedicalVocabularyResponse{
		UpdateMedicalVocabularyOutput: r.Request.Data.(*UpdateMedicalVocabularyOutput),
		response:                      &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// UpdateMedicalVocabularyResponse is the response type for the
// UpdateMedicalVocabulary API operation.
type UpdateMedicalVocabularyResponse struct {
	*UpdateMedicalVocabularyOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// UpdateMedicalVocabulary request.
func (r *UpdateMedicalVocabularyResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
