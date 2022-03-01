// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// TypesconfigLimitRange typesconfig limit range
//
// swagger:model typesconfigLimitRange
type TypesconfigLimitRange struct {

	// container limits
	ContainerLimits *TypesconfigLimitRangeItem `json:"containerLimits,omitempty"`

	// pod limits
	PodLimits *TypesconfigLimitRangeItem `json:"podLimits,omitempty"`
}

// Validate validates this typesconfig limit range
func (m *TypesconfigLimitRange) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateContainerLimits(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePodLimits(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TypesconfigLimitRange) validateContainerLimits(formats strfmt.Registry) error {
	if swag.IsZero(m.ContainerLimits) { // not required
		return nil
	}

	if m.ContainerLimits != nil {
		if err := m.ContainerLimits.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("containerLimits")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("containerLimits")
			}
			return err
		}
	}

	return nil
}

func (m *TypesconfigLimitRange) validatePodLimits(formats strfmt.Registry) error {
	if swag.IsZero(m.PodLimits) { // not required
		return nil
	}

	if m.PodLimits != nil {
		if err := m.PodLimits.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("podLimits")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("podLimits")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this typesconfig limit range based on the context it is used
func (m *TypesconfigLimitRange) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateContainerLimits(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePodLimits(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TypesconfigLimitRange) contextValidateContainerLimits(ctx context.Context, formats strfmt.Registry) error {

	if m.ContainerLimits != nil {
		if err := m.ContainerLimits.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("containerLimits")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("containerLimits")
			}
			return err
		}
	}

	return nil
}

func (m *TypesconfigLimitRange) contextValidatePodLimits(ctx context.Context, formats strfmt.Registry) error {

	if m.PodLimits != nil {
		if err := m.PodLimits.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("podLimits")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("podLimits")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *TypesconfigLimitRange) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TypesconfigLimitRange) UnmarshalBinary(b []byte) error {
	var res TypesconfigLimitRange
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}