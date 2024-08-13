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

// V2Parameters v2 parameters
//
// swagger:model v2.Parameters
type V2Parameters struct {

	// actions
	Actions *V2ActivityParameters `json:"actions,omitempty"`

	// conditions
	Conditions map[string][]V2ConditionGroups `json:"conditions,omitempty"`

	// Installation instructions for the template.
	InstallInstructions string `json:"install_instructions,omitempty"`

	// trigger
	Trigger *ParameterTriggerParameter `json:"trigger,omitempty"`
}

// Validate validates this v2 parameters
func (m *V2Parameters) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateActions(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateConditions(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTrigger(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V2Parameters) validateActions(formats strfmt.Registry) error {
	if swag.IsZero(m.Actions) { // not required
		return nil
	}

	if m.Actions != nil {
		if err := m.Actions.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("actions")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("actions")
			}
			return err
		}
	}

	return nil
}

func (m *V2Parameters) validateConditions(formats strfmt.Registry) error {
	if swag.IsZero(m.Conditions) { // not required
		return nil
	}

	for k := range m.Conditions {

		if err := validate.Required("conditions"+"."+k, "body", m.Conditions[k]); err != nil {
			return err
		}

		for i := 0; i < len(m.Conditions[k]); i++ {

			if err := m.Conditions[k][i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("conditions" + "." + k + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("conditions" + "." + k + "." + strconv.Itoa(i))
				}
				return err
			}

		}

	}

	return nil
}

func (m *V2Parameters) validateTrigger(formats strfmt.Registry) error {
	if swag.IsZero(m.Trigger) { // not required
		return nil
	}

	if m.Trigger != nil {
		if err := m.Trigger.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("trigger")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("trigger")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this v2 parameters based on the context it is used
func (m *V2Parameters) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateActions(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateConditions(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTrigger(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V2Parameters) contextValidateActions(ctx context.Context, formats strfmt.Registry) error {

	if m.Actions != nil {

		if swag.IsZero(m.Actions) { // not required
			return nil
		}

		if err := m.Actions.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("actions")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("actions")
			}
			return err
		}
	}

	return nil
}

func (m *V2Parameters) contextValidateConditions(ctx context.Context, formats strfmt.Registry) error {

	for k := range m.Conditions {

		for i := 0; i < len(m.Conditions[k]); i++ {

			if swag.IsZero(m.Conditions[k][i]) { // not required
				return nil
			}

			if err := m.Conditions[k][i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("conditions" + "." + k + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("conditions" + "." + k + "." + strconv.Itoa(i))
				}
				return err
			}

		}

	}

	return nil
}

func (m *V2Parameters) contextValidateTrigger(ctx context.Context, formats strfmt.Registry) error {

	if m.Trigger != nil {

		if swag.IsZero(m.Trigger) { // not required
			return nil
		}

		if err := m.Trigger.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("trigger")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("trigger")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V2Parameters) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V2Parameters) UnmarshalBinary(b []byte) error {
	var res V2Parameters
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
