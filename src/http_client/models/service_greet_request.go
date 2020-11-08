// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ServiceGreetRequest service greet request
//
// swagger:model serviceGreetRequest
type ServiceGreetRequest struct {

	// name
	Name string `json:"name,omitempty"`
}

// Validate validates this service greet request
func (m *ServiceGreetRequest) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ServiceGreetRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ServiceGreetRequest) UnmarshalBinary(b []byte) error {
	var res ServiceGreetRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}