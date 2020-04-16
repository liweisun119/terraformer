// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package cognitoidentityprovider

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Represents the request to describe the user import job.
type DescribeUserImportJobInput struct {
	_ struct{} `type:"structure"`

	// The job ID for the user import job.
	//
	// JobId is a required field
	JobId *string `min:"1" type:"string" required:"true"`

	// The user pool ID for the user pool that the users are being imported into.
	//
	// UserPoolId is a required field
	UserPoolId *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s DescribeUserImportJobInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeUserImportJobInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeUserImportJobInput"}

	if s.JobId == nil {
		invalidParams.Add(aws.NewErrParamRequired("JobId"))
	}
	if s.JobId != nil && len(*s.JobId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("JobId", 1))
	}

	if s.UserPoolId == nil {
		invalidParams.Add(aws.NewErrParamRequired("UserPoolId"))
	}
	if s.UserPoolId != nil && len(*s.UserPoolId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("UserPoolId", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Represents the response from the server to the request to describe the user
// import job.
type DescribeUserImportJobOutput struct {
	_ struct{} `type:"structure"`

	// The job object that represents the user import job.
	UserImportJob *UserImportJobType `type:"structure"`
}

// String returns the string representation
func (s DescribeUserImportJobOutput) String() string {
	return awsutil.Prettify(s)
}

const opDescribeUserImportJob = "DescribeUserImportJob"

// DescribeUserImportJobRequest returns a request value for making API operation for
// Amazon Cognito Identity Provider.
//
// Describes the user import job.
//
//    // Example sending a request using DescribeUserImportJobRequest.
//    req := client.DescribeUserImportJobRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/cognito-idp-2016-04-18/DescribeUserImportJob
func (c *Client) DescribeUserImportJobRequest(input *DescribeUserImportJobInput) DescribeUserImportJobRequest {
	op := &aws.Operation{
		Name:       opDescribeUserImportJob,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeUserImportJobInput{}
	}

	req := c.newRequest(op, input, &DescribeUserImportJobOutput{})
	return DescribeUserImportJobRequest{Request: req, Input: input, Copy: c.DescribeUserImportJobRequest}
}

// DescribeUserImportJobRequest is the request type for the
// DescribeUserImportJob API operation.
type DescribeUserImportJobRequest struct {
	*aws.Request
	Input *DescribeUserImportJobInput
	Copy  func(*DescribeUserImportJobInput) DescribeUserImportJobRequest
}

// Send marshals and sends the DescribeUserImportJob API request.
func (r DescribeUserImportJobRequest) Send(ctx context.Context) (*DescribeUserImportJobResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeUserImportJobResponse{
		DescribeUserImportJobOutput: r.Request.Data.(*DescribeUserImportJobOutput),
		response:                    &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeUserImportJobResponse is the response type for the
// DescribeUserImportJob API operation.
type DescribeUserImportJobResponse struct {
	*DescribeUserImportJobOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeUserImportJob request.
func (r *DescribeUserImportJobResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}