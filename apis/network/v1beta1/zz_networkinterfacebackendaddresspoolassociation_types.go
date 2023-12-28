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

type NetworkInterfaceBackendAddressPoolAssociationInitParameters struct {

	// The ID of the Load Balancer Backend Address Pool which this Network Interface should be connected to. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=LoadBalancerBackendAddressPool
	// +crossplane:generate:reference:extractor=github.com/upbound/provider-azure/apis/rconfig.ExtractResourceID()
	BackendAddressPoolID *string `json:"backendAddressPoolId,omitempty" tf:"backend_address_pool_id,omitempty"`

	// Reference to a LoadBalancerBackendAddressPool to populate backendAddressPoolId.
	// +kubebuilder:validation:Optional
	BackendAddressPoolIDRef *v1.Reference `json:"backendAddressPoolIdRef,omitempty" tf:"-"`

	// Selector for a LoadBalancerBackendAddressPool to populate backendAddressPoolId.
	// +kubebuilder:validation:Optional
	BackendAddressPoolIDSelector *v1.Selector `json:"backendAddressPoolIdSelector,omitempty" tf:"-"`

	// The Name of the IP Configuration within the Network Interface which should be connected to the Backend Address Pool. Changing this forces a new resource to be created.
	IPConfigurationName *string `json:"ipConfigurationName,omitempty" tf:"ip_configuration_name,omitempty"`

	// The ID of the Network Interface. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=NetworkInterface
	// +crossplane:generate:reference:extractor=github.com/upbound/provider-azure/apis/rconfig.ExtractResourceID()
	NetworkInterfaceID *string `json:"networkInterfaceId,omitempty" tf:"network_interface_id,omitempty"`

	// Reference to a NetworkInterface to populate networkInterfaceId.
	// +kubebuilder:validation:Optional
	NetworkInterfaceIDRef *v1.Reference `json:"networkInterfaceIdRef,omitempty" tf:"-"`

	// Selector for a NetworkInterface to populate networkInterfaceId.
	// +kubebuilder:validation:Optional
	NetworkInterfaceIDSelector *v1.Selector `json:"networkInterfaceIdSelector,omitempty" tf:"-"`
}

type NetworkInterfaceBackendAddressPoolAssociationObservation struct {

	// The ID of the Load Balancer Backend Address Pool which this Network Interface should be connected to. Changing this forces a new resource to be created.
	BackendAddressPoolID *string `json:"backendAddressPoolId,omitempty" tf:"backend_address_pool_id,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The Name of the IP Configuration within the Network Interface which should be connected to the Backend Address Pool. Changing this forces a new resource to be created.
	IPConfigurationName *string `json:"ipConfigurationName,omitempty" tf:"ip_configuration_name,omitempty"`

	// The ID of the Network Interface. Changing this forces a new resource to be created.
	NetworkInterfaceID *string `json:"networkInterfaceId,omitempty" tf:"network_interface_id,omitempty"`
}

type NetworkInterfaceBackendAddressPoolAssociationParameters struct {

	// The ID of the Load Balancer Backend Address Pool which this Network Interface should be connected to. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=LoadBalancerBackendAddressPool
	// +crossplane:generate:reference:extractor=github.com/upbound/provider-azure/apis/rconfig.ExtractResourceID()
	// +kubebuilder:validation:Optional
	BackendAddressPoolID *string `json:"backendAddressPoolId,omitempty" tf:"backend_address_pool_id,omitempty"`

	// Reference to a LoadBalancerBackendAddressPool to populate backendAddressPoolId.
	// +kubebuilder:validation:Optional
	BackendAddressPoolIDRef *v1.Reference `json:"backendAddressPoolIdRef,omitempty" tf:"-"`

	// Selector for a LoadBalancerBackendAddressPool to populate backendAddressPoolId.
	// +kubebuilder:validation:Optional
	BackendAddressPoolIDSelector *v1.Selector `json:"backendAddressPoolIdSelector,omitempty" tf:"-"`

	// The Name of the IP Configuration within the Network Interface which should be connected to the Backend Address Pool. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	IPConfigurationName *string `json:"ipConfigurationName,omitempty" tf:"ip_configuration_name,omitempty"`

	// The ID of the Network Interface. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=NetworkInterface
	// +crossplane:generate:reference:extractor=github.com/upbound/provider-azure/apis/rconfig.ExtractResourceID()
	// +kubebuilder:validation:Optional
	NetworkInterfaceID *string `json:"networkInterfaceId,omitempty" tf:"network_interface_id,omitempty"`

	// Reference to a NetworkInterface to populate networkInterfaceId.
	// +kubebuilder:validation:Optional
	NetworkInterfaceIDRef *v1.Reference `json:"networkInterfaceIdRef,omitempty" tf:"-"`

	// Selector for a NetworkInterface to populate networkInterfaceId.
	// +kubebuilder:validation:Optional
	NetworkInterfaceIDSelector *v1.Selector `json:"networkInterfaceIdSelector,omitempty" tf:"-"`
}

// NetworkInterfaceBackendAddressPoolAssociationSpec defines the desired state of NetworkInterfaceBackendAddressPoolAssociation
type NetworkInterfaceBackendAddressPoolAssociationSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     NetworkInterfaceBackendAddressPoolAssociationParameters `json:"forProvider"`
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
	InitProvider NetworkInterfaceBackendAddressPoolAssociationInitParameters `json:"initProvider,omitempty"`
}

// NetworkInterfaceBackendAddressPoolAssociationStatus defines the observed state of NetworkInterfaceBackendAddressPoolAssociation.
type NetworkInterfaceBackendAddressPoolAssociationStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        NetworkInterfaceBackendAddressPoolAssociationObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// NetworkInterfaceBackendAddressPoolAssociation is the Schema for the NetworkInterfaceBackendAddressPoolAssociations API. Manages the association between a Network Interface and a Load Balancer's Backend Address Pool.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type NetworkInterfaceBackendAddressPoolAssociation struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.ipConfigurationName) || (has(self.initProvider) && has(self.initProvider.ipConfigurationName))",message="spec.forProvider.ipConfigurationName is a required parameter"
	Spec   NetworkInterfaceBackendAddressPoolAssociationSpec   `json:"spec"`
	Status NetworkInterfaceBackendAddressPoolAssociationStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// NetworkInterfaceBackendAddressPoolAssociationList contains a list of NetworkInterfaceBackendAddressPoolAssociations
type NetworkInterfaceBackendAddressPoolAssociationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []NetworkInterfaceBackendAddressPoolAssociation `json:"items"`
}

// Repository type metadata.
var (
	NetworkInterfaceBackendAddressPoolAssociation_Kind             = "NetworkInterfaceBackendAddressPoolAssociation"
	NetworkInterfaceBackendAddressPoolAssociation_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: NetworkInterfaceBackendAddressPoolAssociation_Kind}.String()
	NetworkInterfaceBackendAddressPoolAssociation_KindAPIVersion   = NetworkInterfaceBackendAddressPoolAssociation_Kind + "." + CRDGroupVersion.String()
	NetworkInterfaceBackendAddressPoolAssociation_GroupVersionKind = CRDGroupVersion.WithKind(NetworkInterfaceBackendAddressPoolAssociation_Kind)
)

func init() {
	SchemeBuilder.Register(&NetworkInterfaceBackendAddressPoolAssociation{}, &NetworkInterfaceBackendAddressPoolAssociationList{})
}
