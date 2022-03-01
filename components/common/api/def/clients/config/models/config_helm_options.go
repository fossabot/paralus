// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ConfigHelmOptions config helm options
//
// swagger:model configHelmOptions
type ConfigHelmOptions struct {

	// chart name
	ChartName string `json:"chartName,omitempty"`

	// tag
	Tag string `json:"tag,omitempty"`
}

// Validate validates this config helm options
func (m *ConfigHelmOptions) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this config helm options based on context it is used
func (m *ConfigHelmOptions) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ConfigHelmOptions) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ConfigHelmOptions) UnmarshalBinary(b []byte) error {
	var res ConfigHelmOptions
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}