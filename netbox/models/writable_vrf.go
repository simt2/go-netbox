// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// WritableVRF writable v r f
// swagger:model WritableVRF
type WritableVRF struct {

	// Created
	// Read Only: true
	Created strfmt.Date `json:"created,omitempty"`

	// Custom fields
	CustomFields interface{} `json:"custom_fields,omitempty"`

	// Description
	// Max Length: 100
	Description string `json:"description,omitempty"`

	// Enforce unique space
	//
	// Prevent duplicate prefixes/IP addresses within this VRF
	EnforceUnique bool `json:"enforce_unique,omitempty"`

	// ID
	// Read Only: true
	ID int64 `json:"id,omitempty"`

	// Last updated
	// Read Only: true
	LastUpdated strfmt.DateTime `json:"last_updated,omitempty"`

	// Name
	// Required: true
	// Max Length: 50
	Name *string `json:"name"`

	// Route distinguisher
	// Required: true
	// Max Length: 21
	Rd *string `json:"rd"`

	// Tenant
	Tenant int64 `json:"tenant,omitempty"`
}

// Validate validates this writable v r f
func (m *WritableVRF) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDescription(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateRd(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *WritableVRF) validateDescription(formats strfmt.Registry) error {

	if swag.IsZero(m.Description) { // not required
		return nil
	}

	if err := validate.MaxLength("description", "body", string(m.Description), 100); err != nil {
		return err
	}

	return nil
}

func (m *WritableVRF) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	if err := validate.MaxLength("name", "body", string(*m.Name), 50); err != nil {
		return err
	}

	return nil
}

func (m *WritableVRF) validateRd(formats strfmt.Registry) error {

	if err := validate.Required("rd", "body", m.Rd); err != nil {
		return err
	}

	if err := validate.MaxLength("rd", "body", string(*m.Rd), 21); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *WritableVRF) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *WritableVRF) UnmarshalBinary(b []byte) error {
	var res WritableVRF
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}