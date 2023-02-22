/*
Copyright 2022 Upbound Inc.
*/

// Code generated by upjet. DO NOT EDIT.

package v1beta1

import (
	"github.com/pkg/errors"

	"github.com/upbound/upjet/pkg/resource"
	"github.com/upbound/upjet/pkg/resource/json"
)

// GetTerraformResourceType returns Terraform resource type for this HybridConnection
func (mg *HybridConnection) GetTerraformResourceType() string {
	return "azurerm_relay_hybrid_connection"
}

// GetConnectionDetailsMapping for this HybridConnection
func (tr *HybridConnection) GetConnectionDetailsMapping() map[string]string {
	return nil
}

// GetObservation of this HybridConnection
func (tr *HybridConnection) GetObservation() (map[string]any, error) {
	o, err := json.TFParser.Marshal(tr.Status.AtProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(o, &base)
}

// SetObservation for this HybridConnection
func (tr *HybridConnection) SetObservation(obs map[string]any) error {
	p, err := json.TFParser.Marshal(obs)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Status.AtProvider)
}

// GetID returns ID of underlying Terraform resource of this HybridConnection
func (tr *HybridConnection) GetID() string {
	if tr.Status.AtProvider.ID == nil {
		return ""
	}
	return *tr.Status.AtProvider.ID
}

// GetParameters of this HybridConnection
func (tr *HybridConnection) GetParameters() (map[string]any, error) {
	p, err := json.TFParser.Marshal(tr.Spec.ForProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// SetParameters for this HybridConnection
func (tr *HybridConnection) SetParameters(params map[string]any) error {
	p, err := json.TFParser.Marshal(params)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Spec.ForProvider)
}

// LateInitialize this HybridConnection using its observed tfState.
// returns True if there are any spec changes for the resource.
func (tr *HybridConnection) LateInitialize(attrs []byte) (bool, error) {
	params := &HybridConnectionParameters{}
	if err := json.TFParser.Unmarshal(attrs, params); err != nil {
		return false, errors.Wrap(err, "failed to unmarshal Terraform state parameters for late-initialization")
	}
	opts := []resource.GenericLateInitializerOption{resource.WithZeroValueJSONOmitEmptyFilter(resource.CNameWildcard)}

	li := resource.NewGenericLateInitializer(opts...)
	return li.LateInitialize(&tr.Spec.ForProvider, params)
}

// GetTerraformSchemaVersion returns the associated Terraform schema version
func (tr *HybridConnection) GetTerraformSchemaVersion() int {
	return 0
}

// GetTerraformResourceType returns Terraform resource type for this HybridConnectionAuthorizationRule
func (mg *HybridConnectionAuthorizationRule) GetTerraformResourceType() string {
	return "azurerm_relay_hybrid_connection_authorization_rule"
}

// GetConnectionDetailsMapping for this HybridConnectionAuthorizationRule
func (tr *HybridConnectionAuthorizationRule) GetConnectionDetailsMapping() map[string]string {
	return map[string]string{"primary_connection_string": "status.atProvider.primaryConnectionString", "primary_key": "status.atProvider.primaryKey", "secondary_connection_string": "status.atProvider.secondaryConnectionString", "secondary_key": "status.atProvider.secondaryKey"}
}

// GetObservation of this HybridConnectionAuthorizationRule
func (tr *HybridConnectionAuthorizationRule) GetObservation() (map[string]any, error) {
	o, err := json.TFParser.Marshal(tr.Status.AtProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(o, &base)
}

// SetObservation for this HybridConnectionAuthorizationRule
func (tr *HybridConnectionAuthorizationRule) SetObservation(obs map[string]any) error {
	p, err := json.TFParser.Marshal(obs)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Status.AtProvider)
}

// GetID returns ID of underlying Terraform resource of this HybridConnectionAuthorizationRule
func (tr *HybridConnectionAuthorizationRule) GetID() string {
	if tr.Status.AtProvider.ID == nil {
		return ""
	}
	return *tr.Status.AtProvider.ID
}

// GetParameters of this HybridConnectionAuthorizationRule
func (tr *HybridConnectionAuthorizationRule) GetParameters() (map[string]any, error) {
	p, err := json.TFParser.Marshal(tr.Spec.ForProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// SetParameters for this HybridConnectionAuthorizationRule
func (tr *HybridConnectionAuthorizationRule) SetParameters(params map[string]any) error {
	p, err := json.TFParser.Marshal(params)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Spec.ForProvider)
}

// LateInitialize this HybridConnectionAuthorizationRule using its observed tfState.
// returns True if there are any spec changes for the resource.
func (tr *HybridConnectionAuthorizationRule) LateInitialize(attrs []byte) (bool, error) {
	params := &HybridConnectionAuthorizationRuleParameters{}
	if err := json.TFParser.Unmarshal(attrs, params); err != nil {
		return false, errors.Wrap(err, "failed to unmarshal Terraform state parameters for late-initialization")
	}
	opts := []resource.GenericLateInitializerOption{resource.WithZeroValueJSONOmitEmptyFilter(resource.CNameWildcard)}

	li := resource.NewGenericLateInitializer(opts...)
	return li.LateInitialize(&tr.Spec.ForProvider, params)
}

// GetTerraformSchemaVersion returns the associated Terraform schema version
func (tr *HybridConnectionAuthorizationRule) GetTerraformSchemaVersion() int {
	return 0
}

// GetTerraformResourceType returns Terraform resource type for this EventRelayNamespace
func (mg *EventRelayNamespace) GetTerraformResourceType() string {
	return "azurerm_relay_namespace"
}

// GetConnectionDetailsMapping for this EventRelayNamespace
func (tr *EventRelayNamespace) GetConnectionDetailsMapping() map[string]string {
	return map[string]string{"primary_connection_string": "status.atProvider.primaryConnectionString", "primary_key": "status.atProvider.primaryKey", "secondary_connection_string": "status.atProvider.secondaryConnectionString", "secondary_key": "status.atProvider.secondaryKey"}
}

// GetObservation of this EventRelayNamespace
func (tr *EventRelayNamespace) GetObservation() (map[string]any, error) {
	o, err := json.TFParser.Marshal(tr.Status.AtProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(o, &base)
}

// SetObservation for this EventRelayNamespace
func (tr *EventRelayNamespace) SetObservation(obs map[string]any) error {
	p, err := json.TFParser.Marshal(obs)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Status.AtProvider)
}

// GetID returns ID of underlying Terraform resource of this EventRelayNamespace
func (tr *EventRelayNamespace) GetID() string {
	if tr.Status.AtProvider.ID == nil {
		return ""
	}
	return *tr.Status.AtProvider.ID
}

// GetParameters of this EventRelayNamespace
func (tr *EventRelayNamespace) GetParameters() (map[string]any, error) {
	p, err := json.TFParser.Marshal(tr.Spec.ForProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// SetParameters for this EventRelayNamespace
func (tr *EventRelayNamespace) SetParameters(params map[string]any) error {
	p, err := json.TFParser.Marshal(params)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Spec.ForProvider)
}

// LateInitialize this EventRelayNamespace using its observed tfState.
// returns True if there are any spec changes for the resource.
func (tr *EventRelayNamespace) LateInitialize(attrs []byte) (bool, error) {
	params := &EventRelayNamespaceParameters{}
	if err := json.TFParser.Unmarshal(attrs, params); err != nil {
		return false, errors.Wrap(err, "failed to unmarshal Terraform state parameters for late-initialization")
	}
	opts := []resource.GenericLateInitializerOption{resource.WithZeroValueJSONOmitEmptyFilter(resource.CNameWildcard)}

	li := resource.NewGenericLateInitializer(opts...)
	return li.LateInitialize(&tr.Spec.ForProvider, params)
}

// GetTerraformSchemaVersion returns the associated Terraform schema version
func (tr *EventRelayNamespace) GetTerraformSchemaVersion() int {
	return 0
}

// GetTerraformResourceType returns Terraform resource type for this NamespaceAuthorizationRule
func (mg *NamespaceAuthorizationRule) GetTerraformResourceType() string {
	return "azurerm_relay_namespace_authorization_rule"
}

// GetConnectionDetailsMapping for this NamespaceAuthorizationRule
func (tr *NamespaceAuthorizationRule) GetConnectionDetailsMapping() map[string]string {
	return map[string]string{"primary_connection_string": "status.atProvider.primaryConnectionString", "primary_key": "status.atProvider.primaryKey", "secondary_connection_string": "status.atProvider.secondaryConnectionString", "secondary_key": "status.atProvider.secondaryKey"}
}

// GetObservation of this NamespaceAuthorizationRule
func (tr *NamespaceAuthorizationRule) GetObservation() (map[string]any, error) {
	o, err := json.TFParser.Marshal(tr.Status.AtProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(o, &base)
}

// SetObservation for this NamespaceAuthorizationRule
func (tr *NamespaceAuthorizationRule) SetObservation(obs map[string]any) error {
	p, err := json.TFParser.Marshal(obs)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Status.AtProvider)
}

// GetID returns ID of underlying Terraform resource of this NamespaceAuthorizationRule
func (tr *NamespaceAuthorizationRule) GetID() string {
	if tr.Status.AtProvider.ID == nil {
		return ""
	}
	return *tr.Status.AtProvider.ID
}

// GetParameters of this NamespaceAuthorizationRule
func (tr *NamespaceAuthorizationRule) GetParameters() (map[string]any, error) {
	p, err := json.TFParser.Marshal(tr.Spec.ForProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// SetParameters for this NamespaceAuthorizationRule
func (tr *NamespaceAuthorizationRule) SetParameters(params map[string]any) error {
	p, err := json.TFParser.Marshal(params)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Spec.ForProvider)
}

// LateInitialize this NamespaceAuthorizationRule using its observed tfState.
// returns True if there are any spec changes for the resource.
func (tr *NamespaceAuthorizationRule) LateInitialize(attrs []byte) (bool, error) {
	params := &NamespaceAuthorizationRuleParameters{}
	if err := json.TFParser.Unmarshal(attrs, params); err != nil {
		return false, errors.Wrap(err, "failed to unmarshal Terraform state parameters for late-initialization")
	}
	opts := []resource.GenericLateInitializerOption{resource.WithZeroValueJSONOmitEmptyFilter(resource.CNameWildcard)}

	li := resource.NewGenericLateInitializer(opts...)
	return li.LateInitialize(&tr.Spec.ForProvider, params)
}

// GetTerraformSchemaVersion returns the associated Terraform schema version
func (tr *NamespaceAuthorizationRule) GetTerraformSchemaVersion() int {
	return 0
}
