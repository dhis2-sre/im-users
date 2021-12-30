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

// GroupCreateReader is a Reader for the GroupCreate structure.
type GroupCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GroupCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewGroupCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGroupCreateBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGroupCreateForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGroupCreateUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGroupCreateCreated creates a GroupCreateCreated with default headers values
func NewGroupCreateCreated() *GroupCreateCreated {
	return &GroupCreateCreated{}
}

/* GroupCreateCreated describes a response with status code 201, with default header values.

Group
*/
type GroupCreateCreated struct {
	Payload *models.Group
}

func (o *GroupCreateCreated) Error() string {
	return fmt.Sprintf("[POST /groups][%d] groupCreateCreated  %+v", 201, o.Payload)
}
func (o *GroupCreateCreated) GetPayload() *models.Group {
	return o.Payload
}

func (o *GroupCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Group)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGroupCreateBadRequest creates a GroupCreateBadRequest with default headers values
func NewGroupCreateBadRequest() *GroupCreateBadRequest {
	return &GroupCreateBadRequest{}
}

/* GroupCreateBadRequest describes a response with status code 400, with default header values.

GroupCreateBadRequest group create bad request
*/
type GroupCreateBadRequest struct {
}

func (o *GroupCreateBadRequest) Error() string {
	return fmt.Sprintf("[POST /groups][%d] groupCreateBadRequest ", 400)
}

func (o *GroupCreateBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGroupCreateForbidden creates a GroupCreateForbidden with default headers values
func NewGroupCreateForbidden() *GroupCreateForbidden {
	return &GroupCreateForbidden{}
}

/* GroupCreateForbidden describes a response with status code 403, with default header values.

GroupCreateForbidden group create forbidden
*/
type GroupCreateForbidden struct {
}

func (o *GroupCreateForbidden) Error() string {
	return fmt.Sprintf("[POST /groups][%d] groupCreateForbidden ", 403)
}

func (o *GroupCreateForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGroupCreateUnsupportedMediaType creates a GroupCreateUnsupportedMediaType with default headers values
func NewGroupCreateUnsupportedMediaType() *GroupCreateUnsupportedMediaType {
	return &GroupCreateUnsupportedMediaType{}
}

/* GroupCreateUnsupportedMediaType describes a response with status code 415, with default header values.

GroupCreateUnsupportedMediaType group create unsupported media type
*/
type GroupCreateUnsupportedMediaType struct {
}

func (o *GroupCreateUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[POST /groups][%d] groupCreateUnsupportedMediaType ", 415)
}

func (o *GroupCreateUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
