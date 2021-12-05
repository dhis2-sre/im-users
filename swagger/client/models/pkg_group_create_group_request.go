// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// PkgGroupCreateGroupRequest pkg group create group request
//
// swagger:model pkg_group.CreateGroupRequest
type PkgGroupCreateGroupRequest struct {

	// hostname
	// Required: true
	Hostname *string `json:"hostname"`

	// name
	// Required: true
	Name *string `json:"name"`
}

// Validate validates this pkg group create group request
func (m *PkgGroupCreateGroupRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateHostname(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PkgGroupCreateGroupRequest) validateHostname(formats strfmt.Registry) error {

	if err := validate.Required("hostname", "body", m.Hostname); err != nil {
		return err
	}

	return nil
}

func (m *PkgGroupCreateGroupRequest) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this pkg group create group request based on context it is used
func (m *PkgGroupCreateGroupRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PkgGroupCreateGroupRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PkgGroupCreateGroupRequest) UnmarshalBinary(b []byte) error {
	var res PkgGroupCreateGroupRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
