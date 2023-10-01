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

// MlscannerapiSamplesScanParameters mlscannerapi samples scan parameters
//
// swagger:model mlscannerapi.SamplesScanParameters
type MlscannerapiSamplesScanParameters struct {

	// samples
	// Required: true
	Samples []string `json:"samples"`
}

// Validate validates this mlscannerapi samples scan parameters
func (m *MlscannerapiSamplesScanParameters) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSamples(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MlscannerapiSamplesScanParameters) validateSamples(formats strfmt.Registry) error {

	if err := validate.Required("samples", "body", m.Samples); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this mlscannerapi samples scan parameters based on context it is used
func (m *MlscannerapiSamplesScanParameters) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *MlscannerapiSamplesScanParameters) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MlscannerapiSamplesScanParameters) UnmarshalBinary(b []byte) error {
	var res MlscannerapiSamplesScanParameters
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}