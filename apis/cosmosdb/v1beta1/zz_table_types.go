// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

// Code generated by upjet. DO NOT EDIT.

package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type TableAutoscaleSettingsInitParameters struct {

	// The maximum throughput of the Table (RU/s). Must be between 1,000 and 1,000,000. Must be set in increments of 1,000. Conflicts with throughput.
	MaxThroughput *float64 `json:"maxThroughput,omitempty" tf:"max_throughput,omitempty"`
}

type TableAutoscaleSettingsObservation struct {

	// The maximum throughput of the Table (RU/s). Must be between 1,000 and 1,000,000. Must be set in increments of 1,000. Conflicts with throughput.
	MaxThroughput *float64 `json:"maxThroughput,omitempty" tf:"max_throughput,omitempty"`
}

type TableAutoscaleSettingsParameters struct {

	// The maximum throughput of the Table (RU/s). Must be between 1,000 and 1,000,000. Must be set in increments of 1,000. Conflicts with throughput.
	// +kubebuilder:validation:Optional
	MaxThroughput *float64 `json:"maxThroughput,omitempty" tf:"max_throughput,omitempty"`
}

type TableInitParameters struct {

	// An autoscale_settings block as defined below.
	AutoscaleSettings []TableAutoscaleSettingsInitParameters `json:"autoscaleSettings,omitempty" tf:"autoscale_settings,omitempty"`

	// The throughput of Table (RU/s). Must be set in increments of 100. The minimum value is 400.
	Throughput *float64 `json:"throughput,omitempty" tf:"throughput,omitempty"`
}

type TableObservation struct {

	// The name of the Cosmos DB Table to create the table within. Changing this forces a new resource to be created.
	AccountName *string `json:"accountName,omitempty" tf:"account_name,omitempty"`

	// An autoscale_settings block as defined below.
	AutoscaleSettings []TableAutoscaleSettingsObservation `json:"autoscaleSettings,omitempty" tf:"autoscale_settings,omitempty"`

	// The ID of the CosmosDB Table.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The name of the resource group in which the Cosmos DB Table is created. Changing this forces a new resource to be created.
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// The throughput of Table (RU/s). Must be set in increments of 100. The minimum value is 400.
	Throughput *float64 `json:"throughput,omitempty" tf:"throughput,omitempty"`
}

type TableParameters struct {

	// The name of the Cosmos DB Table to create the table within. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/cosmosdb/v1beta1.Account
	// +kubebuilder:validation:Optional
	AccountName *string `json:"accountName,omitempty" tf:"account_name,omitempty"`

	// Reference to a Account in cosmosdb to populate accountName.
	// +kubebuilder:validation:Optional
	AccountNameRef *v1.Reference `json:"accountNameRef,omitempty" tf:"-"`

	// Selector for a Account in cosmosdb to populate accountName.
	// +kubebuilder:validation:Optional
	AccountNameSelector *v1.Selector `json:"accountNameSelector,omitempty" tf:"-"`

	// An autoscale_settings block as defined below.
	// +kubebuilder:validation:Optional
	AutoscaleSettings []TableAutoscaleSettingsParameters `json:"autoscaleSettings,omitempty" tf:"autoscale_settings,omitempty"`

	// The name of the resource group in which the Cosmos DB Table is created. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/azure/v1beta1.ResourceGroup
	// +kubebuilder:validation:Optional
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// Reference to a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// Selector for a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`

	// The throughput of Table (RU/s). Must be set in increments of 100. The minimum value is 400.
	// +kubebuilder:validation:Optional
	Throughput *float64 `json:"throughput,omitempty" tf:"throughput,omitempty"`
}

// TableSpec defines the desired state of Table
type TableSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     TableParameters `json:"forProvider"`
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
	InitProvider TableInitParameters `json:"initProvider,omitempty"`
}

// TableStatus defines the observed state of Table.
type TableStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        TableObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// Table is the Schema for the Tables API. Manages a Table within a Cosmos DB Account.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type Table struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              TableSpec   `json:"spec"`
	Status            TableStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// TableList contains a list of Tables
type TableList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Table `json:"items"`
}

// Repository type metadata.
var (
	Table_Kind             = "Table"
	Table_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Table_Kind}.String()
	Table_KindAPIVersion   = Table_Kind + "." + CRDGroupVersion.String()
	Table_GroupVersionKind = CRDGroupVersion.WithKind(Table_Kind)
)

func init() {
	SchemeBuilder.Register(&Table{}, &TableList{})
}
