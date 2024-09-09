// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// TypesQueryResult types query result
//
// swagger:model types.QueryResult
type TypesQueryResult struct {

	// result json
	ResultJSON string `json:"result_json,omitempty"`

	// result type
	ResultType int32 `json:"result_type,omitempty"`
}

// Validate validates this types query result
func (m *TypesQueryResult) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this types query result based on context it is used
func (m *TypesQueryResult) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *TypesQueryResult) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TypesQueryResult) UnmarshalBinary(b []byte) error {
	var res TypesQueryResult
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}