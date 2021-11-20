// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Site site
//
// swagger:model site
type Site struct {

	// Site ID.
	// Example: f7ac9cad-ea28-4390-93c8-7add010e8ee3
	// Required: true
	ID *string `json:"id"`

	// Site name.
	// Example: Mount Everest
	Name string `json:"name,omitempty"`

	// parent
	Parent *Site `json:"parent,omitempty"`

	// Status of the site.
	// Example: active
	// Required: true
	// Enum: [active disconnected inactive]
	Status *string `json:"status"`

	// Type of the site.
	// Example: site
	// Required: true
	// Enum: [site endpoint]
	Type *string `json:"type"`
}

// Validate validates this site
func (m *Site) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateParent(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Site) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *Site) validateParent(formats strfmt.Registry) error {
	if swag.IsZero(m.Parent) { // not required
		return nil
	}

	if m.Parent != nil {
		if err := m.Parent.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("parent")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("parent")
			}
			return err
		}
	}

	return nil
}

var siteTypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["active","disconnected","inactive"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		siteTypeStatusPropEnum = append(siteTypeStatusPropEnum, v)
	}
}

const (

	// SiteStatusActive captures enum value "active"
	SiteStatusActive string = "active"

	// SiteStatusDisconnected captures enum value "disconnected"
	SiteStatusDisconnected string = "disconnected"

	// SiteStatusInactive captures enum value "inactive"
	SiteStatusInactive string = "inactive"
)

// prop value enum
func (m *Site) validateStatusEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, siteTypeStatusPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *Site) validateStatus(formats strfmt.Registry) error {

	if err := validate.Required("status", "body", m.Status); err != nil {
		return err
	}

	// value enum
	if err := m.validateStatusEnum("status", "body", *m.Status); err != nil {
		return err
	}

	return nil
}

var siteTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["site","endpoint"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		siteTypeTypePropEnum = append(siteTypeTypePropEnum, v)
	}
}

const (

	// SiteTypeSite captures enum value "site"
	SiteTypeSite string = "site"

	// SiteTypeEndpoint captures enum value "endpoint"
	SiteTypeEndpoint string = "endpoint"
)

// prop value enum
func (m *Site) validateTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, siteTypeTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *Site) validateType(formats strfmt.Registry) error {

	if err := validate.Required("type", "body", m.Type); err != nil {
		return err
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", *m.Type); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this site based on the context it is used
func (m *Site) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateParent(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Site) contextValidateParent(ctx context.Context, formats strfmt.Registry) error {

	if m.Parent != nil {
		if err := m.Parent.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("parent")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("parent")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Site) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Site) UnmarshalBinary(b []byte) error {
	var res Site
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
