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

// CKECluster A Containership Kubernetes Engine (CKE) cluster
// swagger:model CKECluster
type CKECluster struct {

	// Cluster configuration
	// Required: true
	Configuration *CKEClusterConfiguration `json:"configuration"`

	// Timestamp at which the cluster was created
	// Required: true
	CreatedAt *string `json:"created_at"`

	// Provisioning engine
	// Required: true
	Engine *string `json:"engine"`

	// Cluster ID
	// Required: true
	ID UUID `json:"id"`

	// Organization ID of the organization the cluster belongs to
	// Required: true
	OrganizationID UUID `json:"organization_id"`

	// User ID of the cluster owner
	// Required: true
	OwnerID UUID `json:"owner_id"`

	// ID of the provider through which the cluster is provisioned
	// Required: true
	ProviderID UUID `json:"provider_id"`

	// Name of the provider through which the cluster is provisioned
	// Required: true
	ProviderName *string `json:"provider_name"`

	// Version of the provision service used to provision this cluster
	// Required: true
	ProvisionServiceVersion *string `json:"provision_service_version"`

	// Cluster status
	// Required: true
	Status *CKEClusterStatus `json:"status"`

	// Timestamp at which the cluster was updated
	// Required: true
	UpdatedAt *string `json:"updated_at"`
}

// Validate validates this c k e cluster
func (m *CKECluster) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateConfiguration(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreatedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEngine(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOrganizationID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOwnerID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProviderID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProviderName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProvisionServiceVersion(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUpdatedAt(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CKECluster) validateConfiguration(formats strfmt.Registry) error {

	if err := validate.Required("configuration", "body", m.Configuration); err != nil {
		return err
	}

	if m.Configuration != nil {
		if err := m.Configuration.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("configuration")
			}
			return err
		}
	}

	return nil
}

func (m *CKECluster) validateCreatedAt(formats strfmt.Registry) error {

	if err := validate.Required("created_at", "body", m.CreatedAt); err != nil {
		return err
	}

	return nil
}

func (m *CKECluster) validateEngine(formats strfmt.Registry) error {

	if err := validate.Required("engine", "body", m.Engine); err != nil {
		return err
	}

	return nil
}

func (m *CKECluster) validateID(formats strfmt.Registry) error {

	if err := m.ID.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("id")
		}
		return err
	}

	return nil
}

func (m *CKECluster) validateOrganizationID(formats strfmt.Registry) error {

	if err := m.OrganizationID.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("organization_id")
		}
		return err
	}

	return nil
}

func (m *CKECluster) validateOwnerID(formats strfmt.Registry) error {

	if err := m.OwnerID.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("owner_id")
		}
		return err
	}

	return nil
}

func (m *CKECluster) validateProviderID(formats strfmt.Registry) error {

	if err := m.ProviderID.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("provider_id")
		}
		return err
	}

	return nil
}

func (m *CKECluster) validateProviderName(formats strfmt.Registry) error {

	if err := validate.Required("provider_name", "body", m.ProviderName); err != nil {
		return err
	}

	return nil
}

func (m *CKECluster) validateProvisionServiceVersion(formats strfmt.Registry) error {

	if err := validate.Required("provision_service_version", "body", m.ProvisionServiceVersion); err != nil {
		return err
	}

	return nil
}

func (m *CKECluster) validateStatus(formats strfmt.Registry) error {

	if err := validate.Required("status", "body", m.Status); err != nil {
		return err
	}

	if m.Status != nil {
		if err := m.Status.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("status")
			}
			return err
		}
	}

	return nil
}

func (m *CKECluster) validateUpdatedAt(formats strfmt.Registry) error {

	if err := validate.Required("updated_at", "body", m.UpdatedAt); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CKECluster) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CKECluster) UnmarshalBinary(b []byte) error {
	var res CKECluster
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
