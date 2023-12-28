// SPDX-FileCopyrightText: 2023 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

/*
Copyright 2022 Upbound Inc.
*/

// Code generated by upjet. DO NOT EDIT.

package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type IdentityInitParameters struct {

	// Specifies the type of Managed Service Identity that should be configured on this Search Service. The only possible value is SystemAssigned.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type IdentityObservation struct {

	// The Principal ID associated with this Managed Service Identity.
	PrincipalID *string `json:"principalId,omitempty" tf:"principal_id,omitempty"`

	// The Tenant ID associated with this Managed Service Identity.
	TenantID *string `json:"tenantId,omitempty" tf:"tenant_id,omitempty"`

	// Specifies the type of Managed Service Identity that should be configured on this Search Service. The only possible value is SystemAssigned.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type IdentityParameters struct {

	// Specifies the type of Managed Service Identity that should be configured on this Search Service. The only possible value is SystemAssigned.
	// +kubebuilder:validation:Optional
	Type *string `json:"type" tf:"type,omitempty"`
}

type QueryKeysInitParameters struct {
}

type QueryKeysObservation struct {

	// The value of this Query Key.
	Key *string `json:"key,omitempty" tf:"key,omitempty"`

	// The name of this Query Key.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`
}

type QueryKeysParameters struct {
}

type ServiceInitParameters struct {

	// Specifies a list of inbound IPv4 or CIDRs that are allowed to access the Search Service. If the incoming IP request is from an IP address which is not included in the allowed_ips it will be blocked by the Search Services firewall.
	// +listType=set
	AllowedIps []*string `json:"allowedIps,omitempty" tf:"allowed_ips,omitempty"`

	// Specifies the response that the Search Service should return for requests that fail authentication. Possible values include http401WithBearerChallenge or http403.
	AuthenticationFailureMode *string `json:"authenticationFailureMode,omitempty" tf:"authentication_failure_mode,omitempty"`

	// Specifies whether the Search Service should enforce that non-customer resources are encrypted. Defaults to false.
	CustomerManagedKeyEnforcementEnabled *bool `json:"customerManagedKeyEnforcementEnabled,omitempty" tf:"customer_managed_key_enforcement_enabled,omitempty"`

	// Specifies the Hosting Mode, which allows for High Density partitions (that allow for up to 1000 indexes) should be supported. Possible values are highDensity or default. Defaults to default. Changing this forces a new Search Service to be created.
	HostingMode *string `json:"hostingMode,omitempty" tf:"hosting_mode,omitempty"`

	// An identity block as defined below.
	Identity []IdentityInitParameters `json:"identity,omitempty" tf:"identity,omitempty"`

	// Specifies whether the Search Service allows authenticating using API Keys? Defaults to false.
	LocalAuthenticationEnabled *bool `json:"localAuthenticationEnabled,omitempty" tf:"local_authentication_enabled,omitempty"`

	// The Azure Region where the Search Service should exist. Changing this forces a new Search Service to be created.
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// Specifies the number of partitions which should be created. This field cannot be set when using a free or basic sku (see the Microsoft documentation). Possible values include 1, 2, 3, 4, 6, or 12. Defaults to 1.
	PartitionCount *float64 `json:"partitionCount,omitempty" tf:"partition_count,omitempty"`

	// Specifies whether Public Network Access is allowed for this resource. Defaults to true.
	PublicNetworkAccessEnabled *bool `json:"publicNetworkAccessEnabled,omitempty" tf:"public_network_access_enabled,omitempty"`

	// Specifies the number of Replica's which should be created for this Search Service. This field cannot be set when using a free sku (see the Microsoft documentation).
	ReplicaCount *float64 `json:"replicaCount,omitempty" tf:"replica_count,omitempty"`

	// The SKU which should be used for this Search Service. Possible values include basic, free, standard, standard2, standard3, storage_optimized_l1 and storage_optimized_l2. Changing this forces a new Search Service to be created.
	Sku *string `json:"sku,omitempty" tf:"sku,omitempty"`

	// Specifies a mapping of tags which should be assigned to this Search Service.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type ServiceObservation struct {

	// Specifies a list of inbound IPv4 or CIDRs that are allowed to access the Search Service. If the incoming IP request is from an IP address which is not included in the allowed_ips it will be blocked by the Search Services firewall.
	// +listType=set
	AllowedIps []*string `json:"allowedIps,omitempty" tf:"allowed_ips,omitempty"`

	// Specifies the response that the Search Service should return for requests that fail authentication. Possible values include http401WithBearerChallenge or http403.
	AuthenticationFailureMode *string `json:"authenticationFailureMode,omitempty" tf:"authentication_failure_mode,omitempty"`

	// Specifies whether the Search Service should enforce that non-customer resources are encrypted. Defaults to false.
	CustomerManagedKeyEnforcementEnabled *bool `json:"customerManagedKeyEnforcementEnabled,omitempty" tf:"customer_managed_key_enforcement_enabled,omitempty"`

	// Specifies the Hosting Mode, which allows for High Density partitions (that allow for up to 1000 indexes) should be supported. Possible values are highDensity or default. Defaults to default. Changing this forces a new Search Service to be created.
	HostingMode *string `json:"hostingMode,omitempty" tf:"hosting_mode,omitempty"`

	// The ID of the Search Service.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// An identity block as defined below.
	Identity []IdentityObservation `json:"identity,omitempty" tf:"identity,omitempty"`

	// Specifies whether the Search Service allows authenticating using API Keys? Defaults to false.
	LocalAuthenticationEnabled *bool `json:"localAuthenticationEnabled,omitempty" tf:"local_authentication_enabled,omitempty"`

	// The Azure Region where the Search Service should exist. Changing this forces a new Search Service to be created.
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// Specifies the number of partitions which should be created. This field cannot be set when using a free or basic sku (see the Microsoft documentation). Possible values include 1, 2, 3, 4, 6, or 12. Defaults to 1.
	PartitionCount *float64 `json:"partitionCount,omitempty" tf:"partition_count,omitempty"`

	// Specifies whether Public Network Access is allowed for this resource. Defaults to true.
	PublicNetworkAccessEnabled *bool `json:"publicNetworkAccessEnabled,omitempty" tf:"public_network_access_enabled,omitempty"`

	// A query_keys block as defined below.
	QueryKeys []QueryKeysObservation `json:"queryKeys,omitempty" tf:"query_keys,omitempty"`

	// Specifies the number of Replica's which should be created for this Search Service. This field cannot be set when using a free sku (see the Microsoft documentation).
	ReplicaCount *float64 `json:"replicaCount,omitempty" tf:"replica_count,omitempty"`

	// The name of the Resource Group where the Search Service should exist. Changing this forces a new Search Service to be created.
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// The SKU which should be used for this Search Service. Possible values include basic, free, standard, standard2, standard3, storage_optimized_l1 and storage_optimized_l2. Changing this forces a new Search Service to be created.
	Sku *string `json:"sku,omitempty" tf:"sku,omitempty"`

	// Specifies a mapping of tags which should be assigned to this Search Service.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type ServiceParameters struct {

	// Specifies a list of inbound IPv4 or CIDRs that are allowed to access the Search Service. If the incoming IP request is from an IP address which is not included in the allowed_ips it will be blocked by the Search Services firewall.
	// +kubebuilder:validation:Optional
	// +listType=set
	AllowedIps []*string `json:"allowedIps,omitempty" tf:"allowed_ips,omitempty"`

	// Specifies the response that the Search Service should return for requests that fail authentication. Possible values include http401WithBearerChallenge or http403.
	// +kubebuilder:validation:Optional
	AuthenticationFailureMode *string `json:"authenticationFailureMode,omitempty" tf:"authentication_failure_mode,omitempty"`

	// Specifies whether the Search Service should enforce that non-customer resources are encrypted. Defaults to false.
	// +kubebuilder:validation:Optional
	CustomerManagedKeyEnforcementEnabled *bool `json:"customerManagedKeyEnforcementEnabled,omitempty" tf:"customer_managed_key_enforcement_enabled,omitempty"`

	// Specifies the Hosting Mode, which allows for High Density partitions (that allow for up to 1000 indexes) should be supported. Possible values are highDensity or default. Defaults to default. Changing this forces a new Search Service to be created.
	// +kubebuilder:validation:Optional
	HostingMode *string `json:"hostingMode,omitempty" tf:"hosting_mode,omitempty"`

	// An identity block as defined below.
	// +kubebuilder:validation:Optional
	Identity []IdentityParameters `json:"identity,omitempty" tf:"identity,omitempty"`

	// Specifies whether the Search Service allows authenticating using API Keys? Defaults to false.
	// +kubebuilder:validation:Optional
	LocalAuthenticationEnabled *bool `json:"localAuthenticationEnabled,omitempty" tf:"local_authentication_enabled,omitempty"`

	// The Azure Region where the Search Service should exist. Changing this forces a new Search Service to be created.
	// +kubebuilder:validation:Optional
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// Specifies the number of partitions which should be created. This field cannot be set when using a free or basic sku (see the Microsoft documentation). Possible values include 1, 2, 3, 4, 6, or 12. Defaults to 1.
	// +kubebuilder:validation:Optional
	PartitionCount *float64 `json:"partitionCount,omitempty" tf:"partition_count,omitempty"`

	// Specifies whether Public Network Access is allowed for this resource. Defaults to true.
	// +kubebuilder:validation:Optional
	PublicNetworkAccessEnabled *bool `json:"publicNetworkAccessEnabled,omitempty" tf:"public_network_access_enabled,omitempty"`

	// Specifies the number of Replica's which should be created for this Search Service. This field cannot be set when using a free sku (see the Microsoft documentation).
	// +kubebuilder:validation:Optional
	ReplicaCount *float64 `json:"replicaCount,omitempty" tf:"replica_count,omitempty"`

	// The name of the Resource Group where the Search Service should exist. Changing this forces a new Search Service to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/azure/v1beta1.ResourceGroup
	// +kubebuilder:validation:Optional
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// Reference to a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// Selector for a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`

	// The SKU which should be used for this Search Service. Possible values include basic, free, standard, standard2, standard3, storage_optimized_l1 and storage_optimized_l2. Changing this forces a new Search Service to be created.
	// +kubebuilder:validation:Optional
	Sku *string `json:"sku,omitempty" tf:"sku,omitempty"`

	// Specifies a mapping of tags which should be assigned to this Search Service.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

// ServiceSpec defines the desired state of Service
type ServiceSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ServiceParameters `json:"forProvider"`
	// THIS IS A BETA FIELD. It will be honored
	// unless the Management Policies feature flag is disabled.
	// InitProvider holds the same fields as ForProvider, with the exception
	// of Identifier and other resource reference fields. The fields that are
	// in InitProvider are merged into ForProvider when the resource is created.
	// The same fields are also added to the terraform ignore_changes hook, to
	// avoid updating them after creation. This is useful for fields that are
	// required on creation, but we do not desire to update them after creation,
	// for example because of an external controller is managing them, like an
	// autoscaler.
	InitProvider ServiceInitParameters `json:"initProvider,omitempty"`
}

// ServiceStatus defines the observed state of Service.
type ServiceStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ServiceObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Service is the Schema for the Services API. Manages a Search Service.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type Service struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.location) || (has(self.initProvider) && has(self.initProvider.location))",message="spec.forProvider.location is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.sku) || (has(self.initProvider) && has(self.initProvider.sku))",message="spec.forProvider.sku is a required parameter"
	Spec   ServiceSpec   `json:"spec"`
	Status ServiceStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ServiceList contains a list of Services
type ServiceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Service `json:"items"`
}

// Repository type metadata.
var (
	Service_Kind             = "Service"
	Service_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Service_Kind}.String()
	Service_KindAPIVersion   = Service_Kind + "." + CRDGroupVersion.String()
	Service_GroupVersionKind = CRDGroupVersion.WithKind(Service_Kind)
)

func init() {
	SchemeBuilder.Register(&Service{}, &ServiceList{})
}
