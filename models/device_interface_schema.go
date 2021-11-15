// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// DeviceInterfaceSchema device interface schema
//
// swagger:model DeviceInterfaceSchema
type DeviceInterfaceSchema struct {

	// addresses
	Addresses Addresses `json:"addresses,omitempty"`

	// bridge
	Bridge *string `json:"bridge,omitempty"`

	// can display statistics
	// Example: true
	CanDisplayStatistics bool `json:"canDisplayStatistics,omitempty"`

	// capabilities
	Capabilities *Capabilities `json:"capabilities,omitempty"`

	// enabled
	// Example: true
	Enabled bool `json:"enabled,omitempty"`

	// identification
	// Required: true
	Identification *InterfaceIdentification `json:"identification"`

	// is switched port
	IsSwitchedPort bool `json:"isSwitchedPort,omitempty"`

	// lag
	Lag *Lag `json:"lag,omitempty"`

	// mtu
	Mtu float64 `json:"mtu,omitempty"`

	// ospf
	Ospf *InterfaceOspf `json:"ospf,omitempty"`

	// poe
	Poe *InterfacePoe `json:"poe,omitempty"`

	// port
	Port *Port `json:"port,omitempty"`

	// pppoe
	Pppoe *InterfacePppoe `json:"pppoe,omitempty"`

	// proxy a r p
	ProxyARP bool `json:"proxyARP,omitempty"`

	// sfp
	Sfp *InterfaceSfp `json:"sfp,omitempty"`

	// speed
	// Example: auto
	// Pattern: ^autodetect|auto|\d+-(half|full)$
	Speed string `json:"speed,omitempty"`

	// speeds
	Speeds *InterfaceSpeeds `json:"speeds,omitempty"`

	// stations
	Stations Stations `json:"stations,omitempty"`

	// statistics
	Statistics *InterfaceStatistics `json:"statistics,omitempty"`

	// status
	Status *InterfaceStatus `json:"status,omitempty"`

	// switch
	Switch *Switch `json:"switch,omitempty"`

	// visible
	// Example: true
	Visible bool `json:"visible,omitempty"`

	// vlan
	Vlan *Vlan `json:"vlan,omitempty"`

	// wireless
	Wireless *Wireless `json:"wireless,omitempty"`
}

// Validate validates this device interface schema
func (m *DeviceInterfaceSchema) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAddresses(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCapabilities(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIdentification(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLag(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOspf(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePoe(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePort(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePppoe(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSfp(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSpeed(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSpeeds(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStations(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatistics(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSwitch(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVlan(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateWireless(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DeviceInterfaceSchema) validateAddresses(formats strfmt.Registry) error {
	if swag.IsZero(m.Addresses) { // not required
		return nil
	}

	if err := m.Addresses.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("addresses")
		}
		return err
	}

	return nil
}

func (m *DeviceInterfaceSchema) validateCapabilities(formats strfmt.Registry) error {
	if swag.IsZero(m.Capabilities) { // not required
		return nil
	}

	if m.Capabilities != nil {
		if err := m.Capabilities.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("capabilities")
			}
			return err
		}
	}

	return nil
}

func (m *DeviceInterfaceSchema) validateIdentification(formats strfmt.Registry) error {

	if err := validate.Required("identification", "body", m.Identification); err != nil {
		return err
	}

	if m.Identification != nil {
		if err := m.Identification.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("identification")
			}
			return err
		}
	}

	return nil
}

func (m *DeviceInterfaceSchema) validateLag(formats strfmt.Registry) error {
	if swag.IsZero(m.Lag) { // not required
		return nil
	}

	if m.Lag != nil {
		if err := m.Lag.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("lag")
			}
			return err
		}
	}

	return nil
}

func (m *DeviceInterfaceSchema) validateOspf(formats strfmt.Registry) error {
	if swag.IsZero(m.Ospf) { // not required
		return nil
	}

	if m.Ospf != nil {
		if err := m.Ospf.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ospf")
			}
			return err
		}
	}

	return nil
}

func (m *DeviceInterfaceSchema) validatePoe(formats strfmt.Registry) error {
	if swag.IsZero(m.Poe) { // not required
		return nil
	}

	if m.Poe != nil {
		if err := m.Poe.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("poe")
			}
			return err
		}
	}

	return nil
}

func (m *DeviceInterfaceSchema) validatePort(formats strfmt.Registry) error {
	if swag.IsZero(m.Port) { // not required
		return nil
	}

	if m.Port != nil {
		if err := m.Port.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("port")
			}
			return err
		}
	}

	return nil
}

func (m *DeviceInterfaceSchema) validatePppoe(formats strfmt.Registry) error {
	if swag.IsZero(m.Pppoe) { // not required
		return nil
	}

	if m.Pppoe != nil {
		if err := m.Pppoe.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("pppoe")
			}
			return err
		}
	}

	return nil
}

func (m *DeviceInterfaceSchema) validateSfp(formats strfmt.Registry) error {
	if swag.IsZero(m.Sfp) { // not required
		return nil
	}

	if m.Sfp != nil {
		if err := m.Sfp.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("sfp")
			}
			return err
		}
	}

	return nil
}

func (m *DeviceInterfaceSchema) validateSpeed(formats strfmt.Registry) error {
	if swag.IsZero(m.Speed) { // not required
		return nil
	}

	if err := validate.Pattern("speed", "body", string(m.Speed), `^autodetect|auto|\d+-(half|full)$`); err != nil {
		return err
	}

	return nil
}

func (m *DeviceInterfaceSchema) validateSpeeds(formats strfmt.Registry) error {
	if swag.IsZero(m.Speeds) { // not required
		return nil
	}

	if m.Speeds != nil {
		if err := m.Speeds.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("speeds")
			}
			return err
		}
	}

	return nil
}

func (m *DeviceInterfaceSchema) validateStations(formats strfmt.Registry) error {
	if swag.IsZero(m.Stations) { // not required
		return nil
	}

	if err := m.Stations.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("stations")
		}
		return err
	}

	return nil
}

func (m *DeviceInterfaceSchema) validateStatistics(formats strfmt.Registry) error {
	if swag.IsZero(m.Statistics) { // not required
		return nil
	}

	if m.Statistics != nil {
		if err := m.Statistics.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("statistics")
			}
			return err
		}
	}

	return nil
}

func (m *DeviceInterfaceSchema) validateStatus(formats strfmt.Registry) error {
	if swag.IsZero(m.Status) { // not required
		return nil
	}

	if m.Status != nil {
		if err := m.Status.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("status")
			}
			return err
		}
	}

	return nil
}

func (m *DeviceInterfaceSchema) validateSwitch(formats strfmt.Registry) error {
	if swag.IsZero(m.Switch) { // not required
		return nil
	}

	if m.Switch != nil {
		if err := m.Switch.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("switch")
			}
			return err
		}
	}

	return nil
}

func (m *DeviceInterfaceSchema) validateVlan(formats strfmt.Registry) error {
	if swag.IsZero(m.Vlan) { // not required
		return nil
	}

	if m.Vlan != nil {
		if err := m.Vlan.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("vlan")
			}
			return err
		}
	}

	return nil
}

func (m *DeviceInterfaceSchema) validateWireless(formats strfmt.Registry) error {
	if swag.IsZero(m.Wireless) { // not required
		return nil
	}

	if m.Wireless != nil {
		if err := m.Wireless.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("wireless")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this device interface schema based on the context it is used
func (m *DeviceInterfaceSchema) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAddresses(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateCapabilities(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateIdentification(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateLag(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateOspf(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePoe(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePort(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePppoe(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSfp(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSpeeds(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateStations(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateStatistics(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateStatus(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSwitch(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateVlan(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateWireless(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DeviceInterfaceSchema) contextValidateAddresses(ctx context.Context, formats strfmt.Registry) error {

	if err := m.Addresses.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("addresses")
		}
		return err
	}

	return nil
}

func (m *DeviceInterfaceSchema) contextValidateCapabilities(ctx context.Context, formats strfmt.Registry) error {

	if m.Capabilities != nil {
		if err := m.Capabilities.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("capabilities")
			}
			return err
		}
	}

	return nil
}

func (m *DeviceInterfaceSchema) contextValidateIdentification(ctx context.Context, formats strfmt.Registry) error {

	if m.Identification != nil {
		if err := m.Identification.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("identification")
			}
			return err
		}
	}

	return nil
}

func (m *DeviceInterfaceSchema) contextValidateLag(ctx context.Context, formats strfmt.Registry) error {

	if m.Lag != nil {
		if err := m.Lag.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("lag")
			}
			return err
		}
	}

	return nil
}

func (m *DeviceInterfaceSchema) contextValidateOspf(ctx context.Context, formats strfmt.Registry) error {

	if m.Ospf != nil {
		if err := m.Ospf.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ospf")
			}
			return err
		}
	}

	return nil
}

func (m *DeviceInterfaceSchema) contextValidatePoe(ctx context.Context, formats strfmt.Registry) error {

	if m.Poe != nil {
		if err := m.Poe.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("poe")
			}
			return err
		}
	}

	return nil
}

func (m *DeviceInterfaceSchema) contextValidatePort(ctx context.Context, formats strfmt.Registry) error {

	if m.Port != nil {
		if err := m.Port.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("port")
			}
			return err
		}
	}

	return nil
}

func (m *DeviceInterfaceSchema) contextValidatePppoe(ctx context.Context, formats strfmt.Registry) error {

	if m.Pppoe != nil {
		if err := m.Pppoe.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("pppoe")
			}
			return err
		}
	}

	return nil
}

func (m *DeviceInterfaceSchema) contextValidateSfp(ctx context.Context, formats strfmt.Registry) error {

	if m.Sfp != nil {
		if err := m.Sfp.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("sfp")
			}
			return err
		}
	}

	return nil
}

func (m *DeviceInterfaceSchema) contextValidateSpeeds(ctx context.Context, formats strfmt.Registry) error {

	if m.Speeds != nil {
		if err := m.Speeds.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("speeds")
			}
			return err
		}
	}

	return nil
}

func (m *DeviceInterfaceSchema) contextValidateStations(ctx context.Context, formats strfmt.Registry) error {

	if err := m.Stations.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("stations")
		}
		return err
	}

	return nil
}

func (m *DeviceInterfaceSchema) contextValidateStatistics(ctx context.Context, formats strfmt.Registry) error {

	if m.Statistics != nil {
		if err := m.Statistics.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("statistics")
			}
			return err
		}
	}

	return nil
}

func (m *DeviceInterfaceSchema) contextValidateStatus(ctx context.Context, formats strfmt.Registry) error {

	if m.Status != nil {
		if err := m.Status.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("status")
			}
			return err
		}
	}

	return nil
}

func (m *DeviceInterfaceSchema) contextValidateSwitch(ctx context.Context, formats strfmt.Registry) error {

	if m.Switch != nil {
		if err := m.Switch.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("switch")
			}
			return err
		}
	}

	return nil
}

func (m *DeviceInterfaceSchema) contextValidateVlan(ctx context.Context, formats strfmt.Registry) error {

	if m.Vlan != nil {
		if err := m.Vlan.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("vlan")
			}
			return err
		}
	}

	return nil
}

func (m *DeviceInterfaceSchema) contextValidateWireless(ctx context.Context, formats strfmt.Registry) error {

	if m.Wireless != nil {
		if err := m.Wireless.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("wireless")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *DeviceInterfaceSchema) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DeviceInterfaceSchema) UnmarshalBinary(b []byte) error {
	var res DeviceInterfaceSchema
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
