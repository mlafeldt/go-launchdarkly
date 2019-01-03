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

// EnvironmentPost environment post
// swagger:model EnvironmentPost
type EnvironmentPost struct {

	// A color swatch (as an RGB hex value with no leading '#', e.g. C8C8C8).
	// Required: true
	Color *string `json:"color"`

	// The default TTL for the new environment.
	DefaultTTL float64 `json:"defaultTtl,omitempty"`

	// A project-unique key for the new environment.
	// Required: true
	Key *string `json:"key"`

	// The name of the new environment.
	// Required: true
	Name *string `json:"name"`
}

// Validate validates this environment post
func (m *EnvironmentPost) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateColor(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateKey(formats); err != nil {
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

func (m *EnvironmentPost) validateColor(formats strfmt.Registry) error {

	if err := validate.Required("color", "body", m.Color); err != nil {
		return err
	}

	return nil
}

func (m *EnvironmentPost) validateKey(formats strfmt.Registry) error {

	if err := validate.Required("key", "body", m.Key); err != nil {
		return err
	}

	return nil
}

func (m *EnvironmentPost) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *EnvironmentPost) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EnvironmentPost) UnmarshalBinary(b []byte) error {
	var res EnvironmentPost
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
