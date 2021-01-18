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

// FwmgrDomainAddressRange fwmgr domain address range
//
// swagger:model fwmgr.domain.AddressRange
type FwmgrDomainAddressRange struct {

	// address
	// Required: true
	Address *string `json:"address"`

	// netmask
	Netmask int64 `json:"netmask,omitempty"`
}

// Validate validates this fwmgr domain address range
func (m *FwmgrDomainAddressRange) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAddress(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *FwmgrDomainAddressRange) validateAddress(formats strfmt.Registry) error {

	if err := validate.Required("address", "body", m.Address); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *FwmgrDomainAddressRange) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *FwmgrDomainAddressRange) UnmarshalBinary(b []byte) error {
	var res FwmgrDomainAddressRange
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}