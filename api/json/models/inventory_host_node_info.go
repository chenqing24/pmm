// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// InventoryHostNodeInfo HostNodeInfo describes the way Service or Agent runs on Node.
// swagger:model inventoryHostNodeInfo
type InventoryHostNodeInfo struct {

	// Docker container ID.
	ContainerID string `json:"container_id,omitempty"`

	// Docker container name.
	ContainerName string `json:"container_name,omitempty"`

	// Kubernetes pod name.
	KubernetesPodName string `json:"kubernetes_pod_name,omitempty"`

	// Kubernetes pod UID.
	KubernetesPodUID string `json:"kubernetes_pod_uid,omitempty"`

	// Node identifier where Service or Agent runs.
	NodeID string `json:"node_id,omitempty"`
}

// Validate validates this inventory host node info
func (m *InventoryHostNodeInfo) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *InventoryHostNodeInfo) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *InventoryHostNodeInfo) UnmarshalBinary(b []byte) error {
	var res InventoryHostNodeInfo
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}