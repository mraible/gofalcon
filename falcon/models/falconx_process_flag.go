// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// FalconxProcessFlag falconx process flag
//
// swagger:model falconx.ProcessFlag
type FalconxProcessFlag struct {

	// data
	Data string `json:"data,omitempty"`

	// name
	Name string `json:"name,omitempty"`
}

// Validate validates this falconx process flag
func (m *FalconxProcessFlag) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *FalconxProcessFlag) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *FalconxProcessFlag) UnmarshalBinary(b []byte) error {
	var res FalconxProcessFlag
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}