// Code generated by go-swagger; DO NOT EDIT.

package administrative

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

	"github.com/dhis2-sre/im-users/swagger/client/models"
)

// NewPostGroupsParams creates a new PostGroupsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPostGroupsParams() *PostGroupsParams {
	return &PostGroupsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPostGroupsParamsWithTimeout creates a new PostGroupsParams object
// with the ability to set a timeout on a request.
func NewPostGroupsParamsWithTimeout(timeout time.Duration) *PostGroupsParams {
	return &PostGroupsParams{
		timeout: timeout,
	}
}

// NewPostGroupsParamsWithContext creates a new PostGroupsParams object
// with the ability to set a context for a request.
func NewPostGroupsParamsWithContext(ctx context.Context) *PostGroupsParams {
	return &PostGroupsParams{
		Context: ctx,
	}
}

// NewPostGroupsParamsWithHTTPClient creates a new PostGroupsParams object
// with the ability to set a custom HTTPClient for a request.
func NewPostGroupsParamsWithHTTPClient(client *http.Client) *PostGroupsParams {
	return &PostGroupsParams{
		HTTPClient: client,
	}
}

/* PostGroupsParams contains all the parameters to send to the API endpoint
   for the post groups operation.

   Typically these are written to a http.Request.
*/
type PostGroupsParams struct {

	/* CreateGroupRequest.

	   Create group request
	*/
	CreateGroupRequest *models.PkgGroupCreateGroupRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the post groups params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostGroupsParams) WithDefaults() *PostGroupsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the post groups params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostGroupsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the post groups params
func (o *PostGroupsParams) WithTimeout(timeout time.Duration) *PostGroupsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post groups params
func (o *PostGroupsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post groups params
func (o *PostGroupsParams) WithContext(ctx context.Context) *PostGroupsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post groups params
func (o *PostGroupsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post groups params
func (o *PostGroupsParams) WithHTTPClient(client *http.Client) *PostGroupsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post groups params
func (o *PostGroupsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCreateGroupRequest adds the createGroupRequest to the post groups params
func (o *PostGroupsParams) WithCreateGroupRequest(createGroupRequest *models.PkgGroupCreateGroupRequest) *PostGroupsParams {
	o.SetCreateGroupRequest(createGroupRequest)
	return o
}

// SetCreateGroupRequest adds the createGroupRequest to the post groups params
func (o *PostGroupsParams) SetCreateGroupRequest(createGroupRequest *models.PkgGroupCreateGroupRequest) {
	o.CreateGroupRequest = createGroupRequest
}

// WriteToRequest writes these params to a swagger request
func (o *PostGroupsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.CreateGroupRequest != nil {
		if err := r.SetBodyParam(o.CreateGroupRequest); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
