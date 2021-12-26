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

// FindUserByIDReader is a Reader for the FindUserByID structure.
type FindUserByIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *FindUserByIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewFindUserByIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewFindUserByIDForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewFindUserByIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewFindUserByIDUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewFindUserByIDOK creates a FindUserByIDOK with default headers values
func NewFindUserByIDOK() *FindUserByIDOK {
	return &FindUserByIDOK{}
}

/* FindUserByIDOK describes a response with status code 200, with default header values.

User
*/
type FindUserByIDOK struct {
	Payload *models.User
}

func (o *FindUserByIDOK) Error() string {
	return fmt.Sprintf("[GET /findbyid/{id}][%d] findUserByIdOK  %+v", 200, o.Payload)
}
func (o *FindUserByIDOK) GetPayload() *models.User {
	return o.Payload
}

func (o *FindUserByIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.User)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewFindUserByIDForbidden creates a FindUserByIDForbidden with default headers values
func NewFindUserByIDForbidden() *FindUserByIDForbidden {
	return &FindUserByIDForbidden{}
}

/* FindUserByIDForbidden describes a response with status code 403, with default header values.

FindUserByIDForbidden find user by Id forbidden
*/
type FindUserByIDForbidden struct {
}

func (o *FindUserByIDForbidden) Error() string {
	return fmt.Sprintf("[GET /findbyid/{id}][%d] findUserByIdForbidden ", 403)
}

func (o *FindUserByIDForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewFindUserByIDNotFound creates a FindUserByIDNotFound with default headers values
func NewFindUserByIDNotFound() *FindUserByIDNotFound {
	return &FindUserByIDNotFound{}
}

/* FindUserByIDNotFound describes a response with status code 404, with default header values.

FindUserByIDNotFound find user by Id not found
*/
type FindUserByIDNotFound struct {
}

func (o *FindUserByIDNotFound) Error() string {
	return fmt.Sprintf("[GET /findbyid/{id}][%d] findUserByIdNotFound ", 404)
}

func (o *FindUserByIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewFindUserByIDUnsupportedMediaType creates a FindUserByIDUnsupportedMediaType with default headers values
func NewFindUserByIDUnsupportedMediaType() *FindUserByIDUnsupportedMediaType {
	return &FindUserByIDUnsupportedMediaType{}
}

/* FindUserByIDUnsupportedMediaType describes a response with status code 415, with default header values.

FindUserByIDUnsupportedMediaType find user by Id unsupported media type
*/
type FindUserByIDUnsupportedMediaType struct {
}

func (o *FindUserByIDUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /findbyid/{id}][%d] findUserByIdUnsupportedMediaType ", 415)
}

func (o *FindUserByIDUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
