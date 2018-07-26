// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// User user
// swagger:model User
type User struct {

	// anonymous
	Anonymous bool `json:"anonymous,omitempty"`

	// avatar
	Avatar string `json:"avatar,omitempty"`

	// country
	Country string `json:"country,omitempty"`

	// custom
	Custom interface{} `json:"custom,omitempty"`

	// email
	Email string `json:"email,omitempty"`

	// first name
	FirstName string `json:"firstName,omitempty"`

	// ip
	IP string `json:"ip,omitempty"`

	// key
	Key string `json:"key,omitempty"`

	// last name
	LastName string `json:"lastName,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// secondary
	Secondary string `json:"secondary,omitempty"`
}

// Validate validates this user
func (m *User) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *User) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *User) UnmarshalBinary(b []byte) error {
	var res User
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}