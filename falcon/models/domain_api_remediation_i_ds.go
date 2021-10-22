// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// DomainAPIRemediationIDs domain API remediation i ds
//
// swagger:model domain.APIRemediationIDs
type DomainAPIRemediationIDs struct {

	// ids
	Ids []string `json:"ids"`
}

// Validate validates this domain API remediation i ds
func (m *DomainAPIRemediationIDs) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this domain API remediation i ds based on context it is used
func (m *DomainAPIRemediationIDs) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *DomainAPIRemediationIDs) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DomainAPIRemediationIDs) UnmarshalBinary(b []byte) error {
	var res DomainAPIRemediationIDs
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
