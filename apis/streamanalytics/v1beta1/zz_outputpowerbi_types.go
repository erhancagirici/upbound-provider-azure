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

type OutputPowerBIInitParameters struct {

	// The name of the Power BI dataset.
	DataSet *string `json:"dataset,omitempty" tf:"dataset,omitempty"`

	// The ID of the Power BI group, this must be a valid UUID.
	GroupID *string `json:"groupId,omitempty" tf:"group_id,omitempty"`

	// The name of the Power BI group. Use this property to help remember which specific Power BI group id was used.
	GroupName *string `json:"groupName,omitempty" tf:"group_name,omitempty"`

	// The ID of the Stream Analytics Job. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=Job
	// +crossplane:generate:reference:extractor=github.com/upbound/provider-azure/apis/rconfig.ExtractResourceID()
	StreamAnalyticsJobID *string `json:"streamAnalyticsJobId,omitempty" tf:"stream_analytics_job_id,omitempty"`

	// Reference to a Job to populate streamAnalyticsJobId.
	// +kubebuilder:validation:Optional
	StreamAnalyticsJobIDRef *v1.Reference `json:"streamAnalyticsJobIdRef,omitempty" tf:"-"`

	// Selector for a Job to populate streamAnalyticsJobId.
	// +kubebuilder:validation:Optional
	StreamAnalyticsJobIDSelector *v1.Selector `json:"streamAnalyticsJobIdSelector,omitempty" tf:"-"`

	// The name of the Power BI table under the specified dataset.
	Table *string `json:"table,omitempty" tf:"table,omitempty"`

	// The user display name of the user that was used to obtain the refresh token.
	TokenUserDisplayName *string `json:"tokenUserDisplayName,omitempty" tf:"token_user_display_name,omitempty"`

	// The user principal name (UPN) of the user that was used to obtain the refresh token.
	TokenUserPrincipalName *string `json:"tokenUserPrincipalName,omitempty" tf:"token_user_principal_name,omitempty"`
}

type OutputPowerBIObservation struct {

	// The name of the Power BI dataset.
	DataSet *string `json:"dataset,omitempty" tf:"dataset,omitempty"`

	// The ID of the Power BI group, this must be a valid UUID.
	GroupID *string `json:"groupId,omitempty" tf:"group_id,omitempty"`

	// The name of the Power BI group. Use this property to help remember which specific Power BI group id was used.
	GroupName *string `json:"groupName,omitempty" tf:"group_name,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The ID of the Stream Analytics Job. Changing this forces a new resource to be created.
	StreamAnalyticsJobID *string `json:"streamAnalyticsJobId,omitempty" tf:"stream_analytics_job_id,omitempty"`

	// The name of the Power BI table under the specified dataset.
	Table *string `json:"table,omitempty" tf:"table,omitempty"`

	// The user display name of the user that was used to obtain the refresh token.
	TokenUserDisplayName *string `json:"tokenUserDisplayName,omitempty" tf:"token_user_display_name,omitempty"`

	// The user principal name (UPN) of the user that was used to obtain the refresh token.
	TokenUserPrincipalName *string `json:"tokenUserPrincipalName,omitempty" tf:"token_user_principal_name,omitempty"`
}

type OutputPowerBIParameters struct {

	// The name of the Power BI dataset.
	// +kubebuilder:validation:Optional
	DataSet *string `json:"dataset,omitempty" tf:"dataset,omitempty"`

	// The ID of the Power BI group, this must be a valid UUID.
	// +kubebuilder:validation:Optional
	GroupID *string `json:"groupId,omitempty" tf:"group_id,omitempty"`

	// The name of the Power BI group. Use this property to help remember which specific Power BI group id was used.
	// +kubebuilder:validation:Optional
	GroupName *string `json:"groupName,omitempty" tf:"group_name,omitempty"`

	// The ID of the Stream Analytics Job. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=Job
	// +crossplane:generate:reference:extractor=github.com/upbound/provider-azure/apis/rconfig.ExtractResourceID()
	// +kubebuilder:validation:Optional
	StreamAnalyticsJobID *string `json:"streamAnalyticsJobId,omitempty" tf:"stream_analytics_job_id,omitempty"`

	// Reference to a Job to populate streamAnalyticsJobId.
	// +kubebuilder:validation:Optional
	StreamAnalyticsJobIDRef *v1.Reference `json:"streamAnalyticsJobIdRef,omitempty" tf:"-"`

	// Selector for a Job to populate streamAnalyticsJobId.
	// +kubebuilder:validation:Optional
	StreamAnalyticsJobIDSelector *v1.Selector `json:"streamAnalyticsJobIdSelector,omitempty" tf:"-"`

	// The name of the Power BI table under the specified dataset.
	// +kubebuilder:validation:Optional
	Table *string `json:"table,omitempty" tf:"table,omitempty"`

	// The user display name of the user that was used to obtain the refresh token.
	// +kubebuilder:validation:Optional
	TokenUserDisplayName *string `json:"tokenUserDisplayName,omitempty" tf:"token_user_display_name,omitempty"`

	// The user principal name (UPN) of the user that was used to obtain the refresh token.
	// +kubebuilder:validation:Optional
	TokenUserPrincipalName *string `json:"tokenUserPrincipalName,omitempty" tf:"token_user_principal_name,omitempty"`
}

// OutputPowerBISpec defines the desired state of OutputPowerBI
type OutputPowerBISpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     OutputPowerBIParameters `json:"forProvider"`
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
	InitProvider OutputPowerBIInitParameters `json:"initProvider,omitempty"`
}

// OutputPowerBIStatus defines the observed state of OutputPowerBI.
type OutputPowerBIStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        OutputPowerBIObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// OutputPowerBI is the Schema for the OutputPowerBIs API. Manages a Stream Analytics Output powerBI.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type OutputPowerBI struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.dataset) || (has(self.initProvider) && has(self.initProvider.dataset))",message="spec.forProvider.dataset is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.groupId) || (has(self.initProvider) && has(self.initProvider.groupId))",message="spec.forProvider.groupId is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.groupName) || (has(self.initProvider) && has(self.initProvider.groupName))",message="spec.forProvider.groupName is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.table) || (has(self.initProvider) && has(self.initProvider.table))",message="spec.forProvider.table is a required parameter"
	Spec   OutputPowerBISpec   `json:"spec"`
	Status OutputPowerBIStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// OutputPowerBIList contains a list of OutputPowerBIs
type OutputPowerBIList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []OutputPowerBI `json:"items"`
}

// Repository type metadata.
var (
	OutputPowerBI_Kind             = "OutputPowerBI"
	OutputPowerBI_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: OutputPowerBI_Kind}.String()
	OutputPowerBI_KindAPIVersion   = OutputPowerBI_Kind + "." + CRDGroupVersion.String()
	OutputPowerBI_GroupVersionKind = CRDGroupVersion.WithKind(OutputPowerBI_Kind)
)

func init() {
	SchemeBuilder.Register(&OutputPowerBI{}, &OutputPowerBIList{})
}
