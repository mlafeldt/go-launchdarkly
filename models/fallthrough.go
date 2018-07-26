// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// Fallthrough fallthrough
// swagger:model Fallthrough
type Fallthrough struct {

	// rollout
	Rollout *Rollout `json:"rollout,omitempty"`

	// variation
	Variation int64 `json:"variation,omitempty"`
}

// Validate validates this fallthrough
func (m *Fallthrough) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateRollout(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Fallthrough) validateRollout(formats strfmt.Registry) error {

	if swag.IsZero(m.Rollout) { // not required
		return nil
	}

	if m.Rollout != nil {
		if err := m.Rollout.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("rollout")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Fallthrough) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Fallthrough) UnmarshalBinary(b []byte) error {
	var res Fallthrough
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}