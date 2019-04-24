// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// LocationVersion LocationVersion LocationVersion struct for displaying version information per location
// swagger:model LocationVersion
type LocationVersion struct {

	// default
	Default string `json:"default,omitempty"`

	// location
	Location string `json:"location,omitempty"`

	// versions
	Versions []string `json:"versions"`
}

// Validate validates this location version
func (m *LocationVersion) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *LocationVersion) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LocationVersion) UnmarshalBinary(b []byte) error {
	var res LocationVersion
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}