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

// AddUserToGroupReader is a Reader for the AddUserToGroup structure.
type AddUserToGroupReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AddUserToGroupReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewAddUserToGroupCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewAddUserToGroupBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewAddUserToGroupUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewAddUserToGroupForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewAddUserToGroupUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewAddUserToGroupCreated creates a AddUserToGroupCreated with default headers values
func NewAddUserToGroupCreated() *AddUserToGroupCreated {
	return &AddUserToGroupCreated{}
}

/* AddUserToGroupCreated describes a response with status code 201, with default header values.

Group
*/
type AddUserToGroupCreated struct {
	Payload *models.Group
}

// IsSuccess returns true when this add user to group created response has a 2xx status code
func (o *AddUserToGroupCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this add user to group created response has a 3xx status code
func (o *AddUserToGroupCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this add user to group created response has a 4xx status code
func (o *AddUserToGroupCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this add user to group created response has a 5xx status code
func (o *AddUserToGroupCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this add user to group created response a status code equal to that given
func (o *AddUserToGroupCreated) IsCode(code int) bool {
	return code == 201
}

func (o *AddUserToGroupCreated) Error() string {
	return fmt.Sprintf("[POST /groups/{group}/users/{userId}][%d] addUserToGroupCreated  %+v", 201, o.Payload)
}

func (o *AddUserToGroupCreated) String() string {
	return fmt.Sprintf("[POST /groups/{group}/users/{userId}][%d] addUserToGroupCreated  %+v", 201, o.Payload)
}

func (o *AddUserToGroupCreated) GetPayload() *models.Group {
	return o.Payload
}

func (o *AddUserToGroupCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Group)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddUserToGroupBadRequest creates a AddUserToGroupBadRequest with default headers values
func NewAddUserToGroupBadRequest() *AddUserToGroupBadRequest {
	return &AddUserToGroupBadRequest{}
}

/* AddUserToGroupBadRequest describes a response with status code 400, with default header values.

AddUserToGroupBadRequest add user to group bad request
*/
type AddUserToGroupBadRequest struct {
}

// IsSuccess returns true when this add user to group bad request response has a 2xx status code
func (o *AddUserToGroupBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this add user to group bad request response has a 3xx status code
func (o *AddUserToGroupBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this add user to group bad request response has a 4xx status code
func (o *AddUserToGroupBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this add user to group bad request response has a 5xx status code
func (o *AddUserToGroupBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this add user to group bad request response a status code equal to that given
func (o *AddUserToGroupBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *AddUserToGroupBadRequest) Error() string {
	return fmt.Sprintf("[POST /groups/{group}/users/{userId}][%d] addUserToGroupBadRequest ", 400)
}

func (o *AddUserToGroupBadRequest) String() string {
	return fmt.Sprintf("[POST /groups/{group}/users/{userId}][%d] addUserToGroupBadRequest ", 400)
}

func (o *AddUserToGroupBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewAddUserToGroupUnauthorized creates a AddUserToGroupUnauthorized with default headers values
func NewAddUserToGroupUnauthorized() *AddUserToGroupUnauthorized {
	return &AddUserToGroupUnauthorized{}
}

/* AddUserToGroupUnauthorized describes a response with status code 401, with default header values.

AddUserToGroupUnauthorized add user to group unauthorized
*/
type AddUserToGroupUnauthorized struct {
}

// IsSuccess returns true when this add user to group unauthorized response has a 2xx status code
func (o *AddUserToGroupUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this add user to group unauthorized response has a 3xx status code
func (o *AddUserToGroupUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this add user to group unauthorized response has a 4xx status code
func (o *AddUserToGroupUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this add user to group unauthorized response has a 5xx status code
func (o *AddUserToGroupUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this add user to group unauthorized response a status code equal to that given
func (o *AddUserToGroupUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *AddUserToGroupUnauthorized) Error() string {
	return fmt.Sprintf("[POST /groups/{group}/users/{userId}][%d] addUserToGroupUnauthorized ", 401)
}

func (o *AddUserToGroupUnauthorized) String() string {
	return fmt.Sprintf("[POST /groups/{group}/users/{userId}][%d] addUserToGroupUnauthorized ", 401)
}

func (o *AddUserToGroupUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewAddUserToGroupForbidden creates a AddUserToGroupForbidden with default headers values
func NewAddUserToGroupForbidden() *AddUserToGroupForbidden {
	return &AddUserToGroupForbidden{}
}

/* AddUserToGroupForbidden describes a response with status code 403, with default header values.

AddUserToGroupForbidden add user to group forbidden
*/
type AddUserToGroupForbidden struct {
}

// IsSuccess returns true when this add user to group forbidden response has a 2xx status code
func (o *AddUserToGroupForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this add user to group forbidden response has a 3xx status code
func (o *AddUserToGroupForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this add user to group forbidden response has a 4xx status code
func (o *AddUserToGroupForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this add user to group forbidden response has a 5xx status code
func (o *AddUserToGroupForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this add user to group forbidden response a status code equal to that given
func (o *AddUserToGroupForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *AddUserToGroupForbidden) Error() string {
	return fmt.Sprintf("[POST /groups/{group}/users/{userId}][%d] addUserToGroupForbidden ", 403)
}

func (o *AddUserToGroupForbidden) String() string {
	return fmt.Sprintf("[POST /groups/{group}/users/{userId}][%d] addUserToGroupForbidden ", 403)
}

func (o *AddUserToGroupForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewAddUserToGroupUnsupportedMediaType creates a AddUserToGroupUnsupportedMediaType with default headers values
func NewAddUserToGroupUnsupportedMediaType() *AddUserToGroupUnsupportedMediaType {
	return &AddUserToGroupUnsupportedMediaType{}
}

/* AddUserToGroupUnsupportedMediaType describes a response with status code 415, with default header values.

AddUserToGroupUnsupportedMediaType add user to group unsupported media type
*/
type AddUserToGroupUnsupportedMediaType struct {
}

// IsSuccess returns true when this add user to group unsupported media type response has a 2xx status code
func (o *AddUserToGroupUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this add user to group unsupported media type response has a 3xx status code
func (o *AddUserToGroupUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this add user to group unsupported media type response has a 4xx status code
func (o *AddUserToGroupUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this add user to group unsupported media type response has a 5xx status code
func (o *AddUserToGroupUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this add user to group unsupported media type response a status code equal to that given
func (o *AddUserToGroupUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *AddUserToGroupUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[POST /groups/{group}/users/{userId}][%d] addUserToGroupUnsupportedMediaType ", 415)
}

func (o *AddUserToGroupUnsupportedMediaType) String() string {
	return fmt.Sprintf("[POST /groups/{group}/users/{userId}][%d] addUserToGroupUnsupportedMediaType ", 415)
}

func (o *AddUserToGroupUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
