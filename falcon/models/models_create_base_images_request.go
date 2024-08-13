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

// ModelsCreateBaseImagesRequest models create base images request
//
// swagger:model models.CreateBaseImagesRequest
type ModelsCreateBaseImagesRequest struct {

	// base images
	// Required: true
	BaseImages []*ModelsBaseImageRequest `json:"base_images"`
}

// Validate validates this models create base images request
func (m *ModelsCreateBaseImagesRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBaseImages(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ModelsCreateBaseImagesRequest) validateBaseImages(formats strfmt.Registry) error {

	if err := validate.Required("base_images", "body", m.BaseImages); err != nil {
		return err
	}

	for i := 0; i < len(m.BaseImages); i++ {
		if swag.IsZero(m.BaseImages[i]) { // not required
			continue
		}

		if m.BaseImages[i] != nil {
			if err := m.BaseImages[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("base_images" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("base_images" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this models create base images request based on the context it is used
func (m *ModelsCreateBaseImagesRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateBaseImages(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ModelsCreateBaseImagesRequest) contextValidateBaseImages(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.BaseImages); i++ {

		if m.BaseImages[i] != nil {

			if swag.IsZero(m.BaseImages[i]) { // not required
				return nil
			}

			if err := m.BaseImages[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("base_images" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("base_images" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *ModelsCreateBaseImagesRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ModelsCreateBaseImagesRequest) UnmarshalBinary(b []byte) error {
	var res ModelsCreateBaseImagesRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
