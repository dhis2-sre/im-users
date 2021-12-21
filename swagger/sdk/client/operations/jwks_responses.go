// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/dhis2-sre/im-users/swagger/sdk/models"
)

// JwksReader is a Reader for the Jwks structure.
type JwksReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *JwksReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewJwksOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 415:
		result := NewJwksUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewJwksInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewJwksOK creates a JwksOK with default headers values
func NewJwksOK() *JwksOK {
	return &JwksOK{}
}

/* JwksOK describes a response with status code 200, with default header values.

JwksOK jwks o k
*/
type JwksOK struct {
	Payload *models.Key
}

func (o *JwksOK) Error() string {
	return fmt.Sprintf("[GET /jwks][%d] jwksOK  %+v", 200, o.Payload)
}
func (o *JwksOK) GetPayload() *models.Key {
	return o.Payload
}

func (o *JwksOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Key)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewJwksUnsupportedMediaType creates a JwksUnsupportedMediaType with default headers values
func NewJwksUnsupportedMediaType() *JwksUnsupportedMediaType {
	return &JwksUnsupportedMediaType{}
}

/* JwksUnsupportedMediaType describes a response with status code 415, with default header values.

JwksUnsupportedMediaType jwks unsupported media type
*/
type JwksUnsupportedMediaType struct {
}

func (o *JwksUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /jwks][%d] jwksUnsupportedMediaType ", 415)
}

func (o *JwksUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewJwksInternalServerError creates a JwksInternalServerError with default headers values
func NewJwksInternalServerError() *JwksInternalServerError {
	return &JwksInternalServerError{}
}

/* JwksInternalServerError describes a response with status code 500, with default header values.

JwksInternalServerError jwks internal server error
*/
type JwksInternalServerError struct {
}

func (o *JwksInternalServerError) Error() string {
	return fmt.Sprintf("[GET /jwks][%d] jwksInternalServerError ", 500)
}

func (o *JwksInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
