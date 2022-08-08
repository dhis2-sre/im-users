// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	"github.com/dhis2-sre/im-user/swagger/sdk/models"
)

// NewGroupCreateParams creates a new GroupCreateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGroupCreateParams() *GroupCreateParams {
	return &GroupCreateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGroupCreateParamsWithTimeout creates a new GroupCreateParams object
// with the ability to set a timeout on a request.
func NewGroupCreateParamsWithTimeout(timeout time.Duration) *GroupCreateParams {
	return &GroupCreateParams{
		timeout: timeout,
	}
}

// NewGroupCreateParamsWithContext creates a new GroupCreateParams object
// with the ability to set a context for a request.
func NewGroupCreateParamsWithContext(ctx context.Context) *GroupCreateParams {
	return &GroupCreateParams{
		Context: ctx,
	}
}

// NewGroupCreateParamsWithHTTPClient creates a new GroupCreateParams object
// with the ability to set a custom HTTPClient for a request.
func NewGroupCreateParamsWithHTTPClient(client *http.Client) *GroupCreateParams {
	return &GroupCreateParams{
		HTTPClient: client,
	}
}

/* GroupCreateParams contains all the parameters to send to the API endpoint
   for the group create operation.

   Typically these are written to a http.Request.
*/
type GroupCreateParams struct {

	/* Body.

	   Refresh token request body parameter
	*/
	Body *models.CreateGroupRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the group create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GroupCreateParams) WithDefaults() *GroupCreateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the group create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GroupCreateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the group create params
func (o *GroupCreateParams) WithTimeout(timeout time.Duration) *GroupCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the group create params
func (o *GroupCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the group create params
func (o *GroupCreateParams) WithContext(ctx context.Context) *GroupCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the group create params
func (o *GroupCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the group create params
func (o *GroupCreateParams) WithHTTPClient(client *http.Client) *GroupCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the group create params
func (o *GroupCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the group create params
func (o *GroupCreateParams) WithBody(body *models.CreateGroupRequest) *GroupCreateParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the group create params
func (o *GroupCreateParams) SetBody(body *models.CreateGroupRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *GroupCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
