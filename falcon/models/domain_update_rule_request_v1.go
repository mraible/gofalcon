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

// DomainUpdateRuleRequestV1 domain update rule request v1
//
// swagger:model domain.UpdateRuleRequestV1
type DomainUpdateRuleRequestV1 struct {

	// Whether to monitor for breach data
	// Required: true
	BreachMonitoringEnabled *bool `json:"breach_monitoring_enabled"`

	// The filter to be used for searching
	// Required: true
	Filter *string `json:"filter"`

	// The rule ID to be updated
	// Required: true
	ID *string `json:"id"`

	// The name of a particular rule
	// Required: true
	Name *string `json:"name"`

	// The permissions for a particular rule which specifies the rule's access by other users. Possible values: [public private]
	// Required: true
	Permissions *string `json:"permissions"`

	// The priority for a particular rule. Possible values: [low medium high]
	// Required: true
	Priority *string `json:"priority"`
}

// Validate validates this domain update rule request v1
func (m *DomainUpdateRuleRequestV1) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBreachMonitoringEnabled(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFilter(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePermissions(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePriority(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DomainUpdateRuleRequestV1) validateBreachMonitoringEnabled(formats strfmt.Registry) error {

	if err := validate.Required("breach_monitoring_enabled", "body", m.BreachMonitoringEnabled); err != nil {
		return err
	}

	return nil
}

func (m *DomainUpdateRuleRequestV1) validateFilter(formats strfmt.Registry) error {

	if err := validate.Required("filter", "body", m.Filter); err != nil {
		return err
	}

	return nil
}

func (m *DomainUpdateRuleRequestV1) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *DomainUpdateRuleRequestV1) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *DomainUpdateRuleRequestV1) validatePermissions(formats strfmt.Registry) error {

	if err := validate.Required("permissions", "body", m.Permissions); err != nil {
		return err
	}

	return nil
}

func (m *DomainUpdateRuleRequestV1) validatePriority(formats strfmt.Registry) error {

	if err := validate.Required("priority", "body", m.Priority); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this domain update rule request v1 based on context it is used
func (m *DomainUpdateRuleRequestV1) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *DomainUpdateRuleRequestV1) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DomainUpdateRuleRequestV1) UnmarshalBinary(b []byte) error {
	var res DomainUpdateRuleRequestV1
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
