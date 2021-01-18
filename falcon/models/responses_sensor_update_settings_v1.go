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

// ResponsesSensorUpdateSettingsV1 responses sensor update settings v1
//
// swagger:model responses.SensorUpdateSettingsV1
type ResponsesSensorUpdateSettingsV1 struct {

	// build
	// Required: true
	Build *string `json:"build"`
}

// Validate validates this responses sensor update settings v1
func (m *ResponsesSensorUpdateSettingsV1) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBuild(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ResponsesSensorUpdateSettingsV1) validateBuild(formats strfmt.Registry) error {

	if err := validate.Required("build", "body", m.Build); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ResponsesSensorUpdateSettingsV1) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ResponsesSensorUpdateSettingsV1) UnmarshalBinary(b []byte) error {
	var res ResponsesSensorUpdateSettingsV1
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}