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
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type IdentityObservation struct {

	// The ID of the Load Test.
	PrincipalID *string `json:"principalId,omitempty" tf:"principal_id,omitempty"`

	// The ID of the Load Test.
	TenantID *string `json:"tenantId,omitempty" tf:"tenant_id,omitempty"`

	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type IdentityParameters struct {

	// +kubebuilder:validation:Optional
	Type *string `json:"type" tf:"type,omitempty"`
}

type LoadTestInitParameters struct {

	// Description of the resource. Changing this forces a new Load Test to be created.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Specifies the Managed Identity which should be assigned to this Load Test.
	Identity []IdentityInitParameters `json:"identity,omitempty" tf:"identity,omitempty"`

	// The Azure Region where the Load Test should exist. Changing this forces a new Load Test to be created.
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// A mapping of tags which should be assigned to the Load Test.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type LoadTestObservation struct {

	// Resource data plane URI.
	DataPlaneURI *string `json:"dataPlaneUri,omitempty" tf:"data_plane_uri,omitempty"`

	// Description of the resource. Changing this forces a new Load Test to be created.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The ID of the Load Test.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Specifies the Managed Identity which should be assigned to this Load Test.
	Identity []IdentityObservation `json:"identity,omitempty" tf:"identity,omitempty"`

	// The Azure Region where the Load Test should exist. Changing this forces a new Load Test to be created.
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// Specifies the name of the Resource Group within which this Load Test should exist. Changing this forces a new Load Test to be created.
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// A mapping of tags which should be assigned to the Load Test.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type LoadTestParameters struct {

	// Description of the resource. Changing this forces a new Load Test to be created.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Specifies the Managed Identity which should be assigned to this Load Test.
	// +kubebuilder:validation:Optional
	Identity []IdentityParameters `json:"identity,omitempty" tf:"identity,omitempty"`

	// The Azure Region where the Load Test should exist. Changing this forces a new Load Test to be created.
	// +kubebuilder:validation:Optional
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// Specifies the name of the Resource Group within which this Load Test should exist. Changing this forces a new Load Test to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/azure/v1beta1.ResourceGroup
	// +kubebuilder:validation:Optional
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// Reference to a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// Selector for a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`

	// A mapping of tags which should be assigned to the Load Test.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

// LoadTestSpec defines the desired state of LoadTest
type LoadTestSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     LoadTestParameters `json:"forProvider"`
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
	InitProvider LoadTestInitParameters `json:"initProvider,omitempty"`
}

// LoadTestStatus defines the observed state of LoadTest.
type LoadTestStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        LoadTestObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// LoadTest is the Schema for the LoadTests API. Manages a Load Test.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type LoadTest struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.location) || (has(self.initProvider) && has(self.initProvider.location))",message="spec.forProvider.location is a required parameter"
	Spec   LoadTestSpec   `json:"spec"`
	Status LoadTestStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// LoadTestList contains a list of LoadTests
type LoadTestList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []LoadTest `json:"items"`
}

// Repository type metadata.
var (
	LoadTest_Kind             = "LoadTest"
	LoadTest_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: LoadTest_Kind}.String()
	LoadTest_KindAPIVersion   = LoadTest_Kind + "." + CRDGroupVersion.String()
	LoadTest_GroupVersionKind = CRDGroupVersion.WithKind(LoadTest_Kind)
)

func init() {
	SchemeBuilder.Register(&LoadTest{}, &LoadTestList{})
}
