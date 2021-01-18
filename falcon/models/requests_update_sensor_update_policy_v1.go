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

// RequestsUpdateSensorUpdatePolicyV1 An update for a specific policy
//
// swagger:model requests.UpdateSensorUpdatePolicyV1
type RequestsUpdateSensorUpdatePolicyV1 struct {

	// The new description to assign to the policy
	Description string `json:"description,omitempty"`

	// The id of the policy to update
	// Required: true
	ID *string `json:"id"`

	// The new name to assign to the policy
	Name string `json:"name,omitempty"`

	// A collection of sensorUpdate policy settings to update
	Settings *RequestsSensorUpdateSettingsV1 `json:"settings,omitempty"`
}

// Validate validates this requests update sensor update policy v1
func (m *RequestsUpdateSensorUpdatePolicyV1) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSettings(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RequestsUpdateSensorUpdatePolicyV1) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *RequestsUpdateSensorUpdatePolicyV1) validateSettings(formats strfmt.Registry) error {

	if swag.IsZero(m.Settings) { // not required
		return nil
	}

	if m.Settings != nil {
		if err := m.Settings.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("settings")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *RequestsUpdateSensorUpdatePolicyV1) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RequestsUpdateSensorUpdatePolicyV1) UnmarshalBinary(b []byte) error {
	var res RequestsUpdateSensorUpdatePolicyV1
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}