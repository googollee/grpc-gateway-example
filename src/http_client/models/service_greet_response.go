// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ServiceGreetResponse service greet response
//
// swagger:model serviceGreetResponse
type ServiceGreetResponse struct {

	// at
	// Format: date-time
	At strfmt.DateTime `json:"at,omitempty"`

	// greet
	Greet string `json:"greet,omitempty"`
}

// Validate validates this service greet response
func (m *ServiceGreetResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAt(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ServiceGreetResponse) validateAt(formats strfmt.Registry) error {

	if swag.IsZero(m.At) { // not required
		return nil
	}

	if err := validate.FormatOf("at", "body", "date-time", m.At.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ServiceGreetResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ServiceGreetResponse) UnmarshalBinary(b []byte) error {
	var res ServiceGreetResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
