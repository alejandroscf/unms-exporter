// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// DeviceStatusOverview device status overview
//
// swagger:model DeviceStatusOverview
type DeviceStatusOverview struct {

	// enabled
	// Required: true
	Enabled *bool `json:"enabled"`

	// firmware
	Firmware *DeviceFirmware `json:"firmware,omitempty"`

	// identification
	Identification *DeviceIdentification `json:"identification,omitempty"`

	// interfaces
	Interfaces []*DeviceInterfaceSchema `json:"interfaces"`

	// Custom IP address in IPv4 or IPv6 format.
	// Example: 192.168.1.22
	// Required: true
	IPAddress *string `json:"ipAddress"`

	// latest backup
	LatestBackup *LatestBackup `json:"latestBackup,omitempty"`

	// meta
	Meta *DeviceMeta `json:"meta,omitempty"`

	// mode
	Mode string `json:"mode,omitempty"`

	// overview
	Overview *DeviceOverview `json:"overview,omitempty"`

	// upgrade
	Upgrade *DeviceUpgrade `json:"upgrade,omitempty"`
}

// Validate validates this device status overview
func (m *DeviceStatusOverview) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEnabled(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFirmware(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIdentification(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateInterfaces(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIPAddress(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLatestBackup(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMeta(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOverview(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUpgrade(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DeviceStatusOverview) validateEnabled(formats strfmt.Registry) error {

	if err := validate.Required("enabled", "body", m.Enabled); err != nil {
		return err
	}

	return nil
}

func (m *DeviceStatusOverview) validateFirmware(formats strfmt.Registry) error {
	if swag.IsZero(m.Firmware) { // not required
		return nil
	}

	if m.Firmware != nil {
		if err := m.Firmware.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("firmware")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("firmware")
			}
			return err
		}
	}

	return nil
}

func (m *DeviceStatusOverview) validateIdentification(formats strfmt.Registry) error {
	if swag.IsZero(m.Identification) { // not required
		return nil
	}

	if m.Identification != nil {
		if err := m.Identification.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("identification")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("identification")
			}
			return err
		}
	}

	return nil
}

func (m *DeviceStatusOverview) validateInterfaces(formats strfmt.Registry) error {
	if swag.IsZero(m.Interfaces) { // not required
		return nil
	}

	for i := 0; i < len(m.Interfaces); i++ {
		if swag.IsZero(m.Interfaces[i]) { // not required
			continue
		}

		if m.Interfaces[i] != nil {
			if err := m.Interfaces[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("interfaces" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("interfaces" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *DeviceStatusOverview) validateIPAddress(formats strfmt.Registry) error {

	if err := validate.Required("ipAddress", "body", m.IPAddress); err != nil {
		return err
	}

	return nil
}

func (m *DeviceStatusOverview) validateLatestBackup(formats strfmt.Registry) error {
	if swag.IsZero(m.LatestBackup) { // not required
		return nil
	}

	if m.LatestBackup != nil {
		if err := m.LatestBackup.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("latestBackup")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("latestBackup")
			}
			return err
		}
	}

	return nil
}

func (m *DeviceStatusOverview) validateMeta(formats strfmt.Registry) error {
	if swag.IsZero(m.Meta) { // not required
		return nil
	}

	if m.Meta != nil {
		if err := m.Meta.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("meta")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("meta")
			}
			return err
		}
	}

	return nil
}

func (m *DeviceStatusOverview) validateOverview(formats strfmt.Registry) error {
	if swag.IsZero(m.Overview) { // not required
		return nil
	}

	if m.Overview != nil {
		if err := m.Overview.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("overview")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("overview")
			}
			return err
		}
	}

	return nil
}

func (m *DeviceStatusOverview) validateUpgrade(formats strfmt.Registry) error {
	if swag.IsZero(m.Upgrade) { // not required
		return nil
	}

	if m.Upgrade != nil {
		if err := m.Upgrade.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("upgrade")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("upgrade")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this device status overview based on the context it is used
func (m *DeviceStatusOverview) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateFirmware(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateIdentification(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateInterfaces(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateLatestBackup(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateMeta(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateOverview(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateUpgrade(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DeviceStatusOverview) contextValidateFirmware(ctx context.Context, formats strfmt.Registry) error {

	if m.Firmware != nil {
		if err := m.Firmware.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("firmware")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("firmware")
			}
			return err
		}
	}

	return nil
}

func (m *DeviceStatusOverview) contextValidateIdentification(ctx context.Context, formats strfmt.Registry) error {

	if m.Identification != nil {
		if err := m.Identification.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("identification")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("identification")
			}
			return err
		}
	}

	return nil
}

func (m *DeviceStatusOverview) contextValidateInterfaces(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Interfaces); i++ {

		if m.Interfaces[i] != nil {
			if err := m.Interfaces[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("interfaces" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("interfaces" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *DeviceStatusOverview) contextValidateLatestBackup(ctx context.Context, formats strfmt.Registry) error {

	if m.LatestBackup != nil {
		if err := m.LatestBackup.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("latestBackup")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("latestBackup")
			}
			return err
		}
	}

	return nil
}

func (m *DeviceStatusOverview) contextValidateMeta(ctx context.Context, formats strfmt.Registry) error {

	if m.Meta != nil {
		if err := m.Meta.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("meta")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("meta")
			}
			return err
		}
	}

	return nil
}

func (m *DeviceStatusOverview) contextValidateOverview(ctx context.Context, formats strfmt.Registry) error {

	if m.Overview != nil {
		if err := m.Overview.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("overview")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("overview")
			}
			return err
		}
	}

	return nil
}

func (m *DeviceStatusOverview) contextValidateUpgrade(ctx context.Context, formats strfmt.Registry) error {

	if m.Upgrade != nil {
		if err := m.Upgrade.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("upgrade")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("upgrade")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *DeviceStatusOverview) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DeviceStatusOverview) UnmarshalBinary(b []byte) error {
	var res DeviceStatusOverview
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
