// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// AuditLogEntry audit log entry
// swagger:model AuditLogEntry
type AuditLogEntry struct {

	// id
	ID ID `json:"_id,omitempty"`

	// links
	Links *Links `json:"_links,omitempty"`

	// comment
	Comment string `json:"comment,omitempty"`

	// date
	Date int64 `json:"date,omitempty"`

	// description
	Description string `json:"description,omitempty"`

	// kind
	Kind string `json:"kind,omitempty"`

	// member
	Member *Member `json:"member,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// short description
	ShortDescription string `json:"shortDescription,omitempty"`

	// target
	Target *AuditLogEntryTarget `json:"target,omitempty"`

	// title
	Title string `json:"title,omitempty"`

	// title verb
	TitleVerb string `json:"titleVerb,omitempty"`
}

// Validate validates this audit log entry
func (m *AuditLogEntry) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMember(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTarget(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AuditLogEntry) validateID(formats strfmt.Registry) error {

	if swag.IsZero(m.ID) { // not required
		return nil
	}

	if err := m.ID.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("_id")
		}
		return err
	}

	return nil
}

func (m *AuditLogEntry) validateLinks(formats strfmt.Registry) error {

	if swag.IsZero(m.Links) { // not required
		return nil
	}

	if m.Links != nil {
		if err := m.Links.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("_links")
			}
			return err
		}
	}

	return nil
}

func (m *AuditLogEntry) validateMember(formats strfmt.Registry) error {

	if swag.IsZero(m.Member) { // not required
		return nil
	}

	if m.Member != nil {
		if err := m.Member.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("member")
			}
			return err
		}
	}

	return nil
}

func (m *AuditLogEntry) validateTarget(formats strfmt.Registry) error {

	if swag.IsZero(m.Target) { // not required
		return nil
	}

	if m.Target != nil {
		if err := m.Target.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("target")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AuditLogEntry) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AuditLogEntry) UnmarshalBinary(b []byte) error {
	var res AuditLogEntry
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// AuditLogEntryTarget audit log entry target
// swagger:model AuditLogEntryTarget
type AuditLogEntryTarget struct {

	// links
	Links *Links `json:"_links,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// resources
	Resources []string `json:"resources"`
}

// Validate validates this audit log entry target
func (m *AuditLogEntryTarget) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AuditLogEntryTarget) validateLinks(formats strfmt.Registry) error {

	if swag.IsZero(m.Links) { // not required
		return nil
	}

	if m.Links != nil {
		if err := m.Links.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("target" + "." + "_links")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AuditLogEntryTarget) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AuditLogEntryTarget) UnmarshalBinary(b []byte) error {
	var res AuditLogEntryTarget
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
