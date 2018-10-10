// Code generated by go-swagger; DO NOT EDIT.

package types

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// CKEClusterConfiguration The cluster configuration described here is not complete.
// Some fields are opaque / unlisted due to differences between providers.
// Only fields relevant to end users are described.
//
// swagger:model CKEClusterConfiguration
type CKEClusterConfiguration struct {

	// Provider information
	// Required: true
	Provider interface{} `json:"provider"`

	// Provider resources
	// Required: true
	Resource interface{} `json:"resource"`
}

// Validate validates this c k e cluster configuration
func (m *CKEClusterConfiguration) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateProvider(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateResource(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CKEClusterConfiguration) validateProvider(formats strfmt.Registry) error {

	if err := validate.Required("provider", "body", m.Provider); err != nil {
		return err
	}

	return nil
}

func (m *CKEClusterConfiguration) validateResource(formats strfmt.Registry) error {

	if err := validate.Required("resource", "body", m.Resource); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CKEClusterConfiguration) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CKEClusterConfiguration) UnmarshalBinary(b []byte) error {
	var res CKEClusterConfiguration
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
