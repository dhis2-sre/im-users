// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/dhis2-sre/im-user/swagger/sdk/models"
)

// SignupReader is a Reader for the Signup structure.
type SignupReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SignupReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewSignupCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewSignupBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewSignupUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewSignupCreated creates a SignupCreated with default headers values
func NewSignupCreated() *SignupCreated {
	return &SignupCreated{}
}

/* SignupCreated describes a response with status code 201, with default header values.

User
*/
type SignupCreated struct {
	Payload *models.User
}

func (o *SignupCreated) Error() string {
	return fmt.Sprintf("[POST /users][%d] signupCreated  %+v", 201, o.Payload)
}
func (o *SignupCreated) GetPayload() *models.User {
	return o.Payload
}

func (o *SignupCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.User)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSignupBadRequest creates a SignupBadRequest with default headers values
func NewSignupBadRequest() *SignupBadRequest {
	return &SignupBadRequest{}
}

/* SignupBadRequest describes a response with status code 400, with default header values.

SignupBadRequest signup bad request
*/
type SignupBadRequest struct {
}

func (o *SignupBadRequest) Error() string {
	return fmt.Sprintf("[POST /users][%d] signupBadRequest ", 400)
}

func (o *SignupBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewSignupUnsupportedMediaType creates a SignupUnsupportedMediaType with default headers values
func NewSignupUnsupportedMediaType() *SignupUnsupportedMediaType {
	return &SignupUnsupportedMediaType{}
}

/* SignupUnsupportedMediaType describes a response with status code 415, with default header values.

SignupUnsupportedMediaType signup unsupported media type
*/
type SignupUnsupportedMediaType struct {
}

func (o *SignupUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[POST /users][%d] signupUnsupportedMediaType ", 415)
}

func (o *SignupUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
