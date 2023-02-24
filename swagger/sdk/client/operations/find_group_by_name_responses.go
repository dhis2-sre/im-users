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

// FindGroupByNameReader is a Reader for the FindGroupByName structure.
type FindGroupByNameReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *FindGroupByNameReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewFindGroupByNameOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewFindGroupByNameUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewFindGroupByNameForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewFindGroupByNameNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewFindGroupByNameUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewFindGroupByNameOK creates a FindGroupByNameOK with default headers values
func NewFindGroupByNameOK() *FindGroupByNameOK {
	return &FindGroupByNameOK{}
}

/*
FindGroupByNameOK describes a response with status code 200, with default header values.

Group
*/
type FindGroupByNameOK struct {
	Payload *models.Group
}

// IsSuccess returns true when this find group by name o k response has a 2xx status code
func (o *FindGroupByNameOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this find group by name o k response has a 3xx status code
func (o *FindGroupByNameOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this find group by name o k response has a 4xx status code
func (o *FindGroupByNameOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this find group by name o k response has a 5xx status code
func (o *FindGroupByNameOK) IsServerError() bool {
	return false
}

// IsCode returns true when this find group by name o k response a status code equal to that given
func (o *FindGroupByNameOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the find group by name o k response
func (o *FindGroupByNameOK) Code() int {
	return 200
}

func (o *FindGroupByNameOK) Error() string {
	return fmt.Sprintf("[GET /groups/{name}][%d] findGroupByNameOK  %+v", 200, o.Payload)
}

func (o *FindGroupByNameOK) String() string {
	return fmt.Sprintf("[GET /groups/{name}][%d] findGroupByNameOK  %+v", 200, o.Payload)
}

func (o *FindGroupByNameOK) GetPayload() *models.Group {
	return o.Payload
}

func (o *FindGroupByNameOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Group)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewFindGroupByNameUnauthorized creates a FindGroupByNameUnauthorized with default headers values
func NewFindGroupByNameUnauthorized() *FindGroupByNameUnauthorized {
	return &FindGroupByNameUnauthorized{}
}

/*
FindGroupByNameUnauthorized describes a response with status code 401, with default header values.

FindGroupByNameUnauthorized find group by name unauthorized
*/
type FindGroupByNameUnauthorized struct {
}

// IsSuccess returns true when this find group by name unauthorized response has a 2xx status code
func (o *FindGroupByNameUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this find group by name unauthorized response has a 3xx status code
func (o *FindGroupByNameUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this find group by name unauthorized response has a 4xx status code
func (o *FindGroupByNameUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this find group by name unauthorized response has a 5xx status code
func (o *FindGroupByNameUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this find group by name unauthorized response a status code equal to that given
func (o *FindGroupByNameUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the find group by name unauthorized response
func (o *FindGroupByNameUnauthorized) Code() int {
	return 401
}

func (o *FindGroupByNameUnauthorized) Error() string {
	return fmt.Sprintf("[GET /groups/{name}][%d] findGroupByNameUnauthorized ", 401)
}

func (o *FindGroupByNameUnauthorized) String() string {
	return fmt.Sprintf("[GET /groups/{name}][%d] findGroupByNameUnauthorized ", 401)
}

func (o *FindGroupByNameUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewFindGroupByNameForbidden creates a FindGroupByNameForbidden with default headers values
func NewFindGroupByNameForbidden() *FindGroupByNameForbidden {
	return &FindGroupByNameForbidden{}
}

/*
FindGroupByNameForbidden describes a response with status code 403, with default header values.

FindGroupByNameForbidden find group by name forbidden
*/
type FindGroupByNameForbidden struct {
}

// IsSuccess returns true when this find group by name forbidden response has a 2xx status code
func (o *FindGroupByNameForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this find group by name forbidden response has a 3xx status code
func (o *FindGroupByNameForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this find group by name forbidden response has a 4xx status code
func (o *FindGroupByNameForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this find group by name forbidden response has a 5xx status code
func (o *FindGroupByNameForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this find group by name forbidden response a status code equal to that given
func (o *FindGroupByNameForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the find group by name forbidden response
func (o *FindGroupByNameForbidden) Code() int {
	return 403
}

func (o *FindGroupByNameForbidden) Error() string {
	return fmt.Sprintf("[GET /groups/{name}][%d] findGroupByNameForbidden ", 403)
}

func (o *FindGroupByNameForbidden) String() string {
	return fmt.Sprintf("[GET /groups/{name}][%d] findGroupByNameForbidden ", 403)
}

func (o *FindGroupByNameForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewFindGroupByNameNotFound creates a FindGroupByNameNotFound with default headers values
func NewFindGroupByNameNotFound() *FindGroupByNameNotFound {
	return &FindGroupByNameNotFound{}
}

/*
FindGroupByNameNotFound describes a response with status code 404, with default header values.

FindGroupByNameNotFound find group by name not found
*/
type FindGroupByNameNotFound struct {
}

// IsSuccess returns true when this find group by name not found response has a 2xx status code
func (o *FindGroupByNameNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this find group by name not found response has a 3xx status code
func (o *FindGroupByNameNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this find group by name not found response has a 4xx status code
func (o *FindGroupByNameNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this find group by name not found response has a 5xx status code
func (o *FindGroupByNameNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this find group by name not found response a status code equal to that given
func (o *FindGroupByNameNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the find group by name not found response
func (o *FindGroupByNameNotFound) Code() int {
	return 404
}

func (o *FindGroupByNameNotFound) Error() string {
	return fmt.Sprintf("[GET /groups/{name}][%d] findGroupByNameNotFound ", 404)
}

func (o *FindGroupByNameNotFound) String() string {
	return fmt.Sprintf("[GET /groups/{name}][%d] findGroupByNameNotFound ", 404)
}

func (o *FindGroupByNameNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewFindGroupByNameUnsupportedMediaType creates a FindGroupByNameUnsupportedMediaType with default headers values
func NewFindGroupByNameUnsupportedMediaType() *FindGroupByNameUnsupportedMediaType {
	return &FindGroupByNameUnsupportedMediaType{}
}

/*
FindGroupByNameUnsupportedMediaType describes a response with status code 415, with default header values.

FindGroupByNameUnsupportedMediaType find group by name unsupported media type
*/
type FindGroupByNameUnsupportedMediaType struct {
}

// IsSuccess returns true when this find group by name unsupported media type response has a 2xx status code
func (o *FindGroupByNameUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this find group by name unsupported media type response has a 3xx status code
func (o *FindGroupByNameUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this find group by name unsupported media type response has a 4xx status code
func (o *FindGroupByNameUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this find group by name unsupported media type response has a 5xx status code
func (o *FindGroupByNameUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this find group by name unsupported media type response a status code equal to that given
func (o *FindGroupByNameUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

// Code gets the status code for the find group by name unsupported media type response
func (o *FindGroupByNameUnsupportedMediaType) Code() int {
	return 415
}

func (o *FindGroupByNameUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /groups/{name}][%d] findGroupByNameUnsupportedMediaType ", 415)
}

func (o *FindGroupByNameUnsupportedMediaType) String() string {
	return fmt.Sprintf("[GET /groups/{name}][%d] findGroupByNameUnsupportedMediaType ", 415)
}

func (o *FindGroupByNameUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
