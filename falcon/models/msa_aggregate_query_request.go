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

// MsaAggregateQueryRequest msa aggregate query request
//
// swagger:model msa.AggregateQueryRequest
type MsaAggregateQueryRequest struct {

	// date ranges
	// Required: true
	DateRanges []*MsaDateRangeSpec `json:"date_ranges"`

	// field
	// Required: true
	Field *string `json:"field"`

	// filter
	// Required: true
	Filter *string `json:"filter"`

	// from
	// Required: true
	From *int32 `json:"from"`

	// include
	// Required: true
	Include *string `json:"include"`

	// interval
	// Required: true
	Interval *string `json:"interval"`

	// min doc count
	// Required: true
	MinDocCount *int64 `json:"min_doc_count"`

	// missing
	// Required: true
	Missing *string `json:"missing"`

	// name
	// Required: true
	Name *string `json:"name"`

	// q
	// Required: true
	Q *string `json:"q"`

	// ranges
	// Required: true
	Ranges []*MsaRangeSpec `json:"ranges"`

	// size
	// Required: true
	Size *int32 `json:"size"`

	// sort
	// Required: true
	Sort *string `json:"sort"`

	// sub aggregates
	// Required: true
	SubAggregates []*MsaAggregateQueryRequest `json:"sub_aggregates"`

	// time zone
	// Required: true
	TimeZone *string `json:"time_zone"`

	// type
	// Required: true
	Type *string `json:"type"`
}

// Validate validates this msa aggregate query request
func (m *MsaAggregateQueryRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDateRanges(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateField(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFilter(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFrom(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateInclude(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateInterval(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMinDocCount(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMissing(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateQ(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRanges(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSize(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSort(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSubAggregates(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTimeZone(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MsaAggregateQueryRequest) validateDateRanges(formats strfmt.Registry) error {

	if err := validate.Required("date_ranges", "body", m.DateRanges); err != nil {
		return err
	}

	for i := 0; i < len(m.DateRanges); i++ {
		if swag.IsZero(m.DateRanges[i]) { // not required
			continue
		}

		if m.DateRanges[i] != nil {
			if err := m.DateRanges[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("date_ranges" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("date_ranges" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *MsaAggregateQueryRequest) validateField(formats strfmt.Registry) error {

	if err := validate.Required("field", "body", m.Field); err != nil {
		return err
	}

	return nil
}

func (m *MsaAggregateQueryRequest) validateFilter(formats strfmt.Registry) error {

	if err := validate.Required("filter", "body", m.Filter); err != nil {
		return err
	}

	return nil
}

func (m *MsaAggregateQueryRequest) validateFrom(formats strfmt.Registry) error {

	if err := validate.Required("from", "body", m.From); err != nil {
		return err
	}

	return nil
}

func (m *MsaAggregateQueryRequest) validateInclude(formats strfmt.Registry) error {

	if err := validate.Required("include", "body", m.Include); err != nil {
		return err
	}

	return nil
}

func (m *MsaAggregateQueryRequest) validateInterval(formats strfmt.Registry) error {

	if err := validate.Required("interval", "body", m.Interval); err != nil {
		return err
	}

	return nil
}

func (m *MsaAggregateQueryRequest) validateMinDocCount(formats strfmt.Registry) error {

	if err := validate.Required("min_doc_count", "body", m.MinDocCount); err != nil {
		return err
	}

	return nil
}

func (m *MsaAggregateQueryRequest) validateMissing(formats strfmt.Registry) error {

	if err := validate.Required("missing", "body", m.Missing); err != nil {
		return err
	}

	return nil
}

func (m *MsaAggregateQueryRequest) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *MsaAggregateQueryRequest) validateQ(formats strfmt.Registry) error {

	if err := validate.Required("q", "body", m.Q); err != nil {
		return err
	}

	return nil
}

func (m *MsaAggregateQueryRequest) validateRanges(formats strfmt.Registry) error {

	if err := validate.Required("ranges", "body", m.Ranges); err != nil {
		return err
	}

	for i := 0; i < len(m.Ranges); i++ {
		if swag.IsZero(m.Ranges[i]) { // not required
			continue
		}

		if m.Ranges[i] != nil {
			if err := m.Ranges[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("ranges" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("ranges" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *MsaAggregateQueryRequest) validateSize(formats strfmt.Registry) error {

	if err := validate.Required("size", "body", m.Size); err != nil {
		return err
	}

	return nil
}

func (m *MsaAggregateQueryRequest) validateSort(formats strfmt.Registry) error {

	if err := validate.Required("sort", "body", m.Sort); err != nil {
		return err
	}

	return nil
}

func (m *MsaAggregateQueryRequest) validateSubAggregates(formats strfmt.Registry) error {

	if err := validate.Required("sub_aggregates", "body", m.SubAggregates); err != nil {
		return err
	}

	for i := 0; i < len(m.SubAggregates); i++ {
		if swag.IsZero(m.SubAggregates[i]) { // not required
			continue
		}

		if m.SubAggregates[i] != nil {
			if err := m.SubAggregates[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("sub_aggregates" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("sub_aggregates" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *MsaAggregateQueryRequest) validateTimeZone(formats strfmt.Registry) error {

	if err := validate.Required("time_zone", "body", m.TimeZone); err != nil {
		return err
	}

	return nil
}

func (m *MsaAggregateQueryRequest) validateType(formats strfmt.Registry) error {

	if err := validate.Required("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this msa aggregate query request based on the context it is used
func (m *MsaAggregateQueryRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateDateRanges(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRanges(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSubAggregates(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MsaAggregateQueryRequest) contextValidateDateRanges(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.DateRanges); i++ {

		if m.DateRanges[i] != nil {
			if err := m.DateRanges[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("date_ranges" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("date_ranges" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *MsaAggregateQueryRequest) contextValidateRanges(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Ranges); i++ {

		if m.Ranges[i] != nil {
			if err := m.Ranges[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("ranges" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("ranges" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *MsaAggregateQueryRequest) contextValidateSubAggregates(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.SubAggregates); i++ {

		if m.SubAggregates[i] != nil {
			if err := m.SubAggregates[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("sub_aggregates" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("sub_aggregates" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *MsaAggregateQueryRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MsaAggregateQueryRequest) UnmarshalBinary(b []byte) error {
	var res MsaAggregateQueryRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
