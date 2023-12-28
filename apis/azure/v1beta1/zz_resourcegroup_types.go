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

type ResourceGroupInitParameters struct {

	// The Azure Region where the Resource Group should exist. Changing this forces a new Resource Group to be created.
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// A mapping of tags which should be assigned to the Resource Group.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type ResourceGroupObservation struct {

	// The ID of the Resource Group.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The Azure Region where the Resource Group should exist. Changing this forces a new Resource Group to be created.
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// A mapping of tags which should be assigned to the Resource Group.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type ResourceGroupParameters struct {

	// The Azure Region where the Resource Group should exist. Changing this forces a new Resource Group to be created.
	// +kubebuilder:validation:Optional
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// A mapping of tags which should be assigned to the Resource Group.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

// ResourceGroupSpec defines the desired state of ResourceGroup
type ResourceGroupSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ResourceGroupParameters `json:"forProvider"`
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
	InitProvider ResourceGroupInitParameters `json:"initProvider,omitempty"`
}

// ResourceGroupStatus defines the observed state of ResourceGroup.
type ResourceGroupStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ResourceGroupObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ResourceGroup is the Schema for the ResourceGroups API. Manages a Resource Group.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type ResourceGroup struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.location) || (has(self.initProvider) && has(self.initProvider.location))",message="spec.forProvider.location is a required parameter"
	Spec   ResourceGroupSpec   `json:"spec"`
	Status ResourceGroupStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ResourceGroupList contains a list of ResourceGroups
type ResourceGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ResourceGroup `json:"items"`
}

// Repository type metadata.
var (
	ResourceGroup_Kind             = "ResourceGroup"
	ResourceGroup_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ResourceGroup_Kind}.String()
	ResourceGroup_KindAPIVersion   = ResourceGroup_Kind + "." + CRDGroupVersion.String()
	ResourceGroup_GroupVersionKind = CRDGroupVersion.WithKind(ResourceGroup_Kind)
)

func init() {
	SchemeBuilder.Register(&ResourceGroup{}, &ResourceGroupList{})
}
