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

type LinkedServiceSynapseInitParameters struct {

	// A map of additional properties to associate with the Data Factory Linked Service Synapse.
	// +mapType=granular
	AdditionalProperties map[string]*string `json:"additionalProperties,omitempty" tf:"additional_properties,omitempty"`

	// List of tags that can be used for describing the Data Factory Linked Service Synapse.
	Annotations []*string `json:"annotations,omitempty" tf:"annotations,omitempty"`

	// The connection string in which to authenticate with the Synapse.
	ConnectionString *string `json:"connectionString,omitempty" tf:"connection_string,omitempty"`

	// The description for the Data Factory Linked Service Synapse.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The integration runtime reference to associate with the Data Factory Linked Service Synapse.
	IntegrationRuntimeName *string `json:"integrationRuntimeName,omitempty" tf:"integration_runtime_name,omitempty"`

	// A key_vault_password block as defined below. Use this argument to store Synapse password in an existing Key Vault. It needs an existing Key Vault Data Factory Linked Service.
	KeyVaultPassword []LinkedServiceSynapseKeyVaultPasswordInitParameters `json:"keyVaultPassword,omitempty" tf:"key_vault_password,omitempty"`

	// A map of parameters to associate with the Data Factory Linked Service Synapse.
	// +mapType=granular
	Parameters map[string]*string `json:"parameters,omitempty" tf:"parameters,omitempty"`
}

type LinkedServiceSynapseKeyVaultPasswordInitParameters struct {

	// Specifies the name of an existing Key Vault Data Factory Linked Service.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/datafactory/v1beta1.LinkedServiceKeyVault
	LinkedServiceName *string `json:"linkedServiceName,omitempty" tf:"linked_service_name,omitempty"`

	// Reference to a LinkedServiceKeyVault in datafactory to populate linkedServiceName.
	// +kubebuilder:validation:Optional
	LinkedServiceNameRef *v1.Reference `json:"linkedServiceNameRef,omitempty" tf:"-"`

	// Selector for a LinkedServiceKeyVault in datafactory to populate linkedServiceName.
	// +kubebuilder:validation:Optional
	LinkedServiceNameSelector *v1.Selector `json:"linkedServiceNameSelector,omitempty" tf:"-"`

	// Specifies the secret name in Azure Key Vault that stores Synapse password.
	SecretName *string `json:"secretName,omitempty" tf:"secret_name,omitempty"`
}

type LinkedServiceSynapseKeyVaultPasswordObservation struct {

	// Specifies the name of an existing Key Vault Data Factory Linked Service.
	LinkedServiceName *string `json:"linkedServiceName,omitempty" tf:"linked_service_name,omitempty"`

	// Specifies the secret name in Azure Key Vault that stores Synapse password.
	SecretName *string `json:"secretName,omitempty" tf:"secret_name,omitempty"`
}

type LinkedServiceSynapseKeyVaultPasswordParameters struct {

	// Specifies the name of an existing Key Vault Data Factory Linked Service.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/datafactory/v1beta1.LinkedServiceKeyVault
	// +kubebuilder:validation:Optional
	LinkedServiceName *string `json:"linkedServiceName,omitempty" tf:"linked_service_name,omitempty"`

	// Reference to a LinkedServiceKeyVault in datafactory to populate linkedServiceName.
	// +kubebuilder:validation:Optional
	LinkedServiceNameRef *v1.Reference `json:"linkedServiceNameRef,omitempty" tf:"-"`

	// Selector for a LinkedServiceKeyVault in datafactory to populate linkedServiceName.
	// +kubebuilder:validation:Optional
	LinkedServiceNameSelector *v1.Selector `json:"linkedServiceNameSelector,omitempty" tf:"-"`

	// Specifies the secret name in Azure Key Vault that stores Synapse password.
	// +kubebuilder:validation:Optional
	SecretName *string `json:"secretName" tf:"secret_name,omitempty"`
}

type LinkedServiceSynapseObservation struct {

	// A map of additional properties to associate with the Data Factory Linked Service Synapse.
	// +mapType=granular
	AdditionalProperties map[string]*string `json:"additionalProperties,omitempty" tf:"additional_properties,omitempty"`

	// List of tags that can be used for describing the Data Factory Linked Service Synapse.
	Annotations []*string `json:"annotations,omitempty" tf:"annotations,omitempty"`

	// The connection string in which to authenticate with the Synapse.
	ConnectionString *string `json:"connectionString,omitempty" tf:"connection_string,omitempty"`

	// The Data Factory ID in which to associate the Linked Service with. Changing this forces a new resource.
	DataFactoryID *string `json:"dataFactoryId,omitempty" tf:"data_factory_id,omitempty"`

	// The description for the Data Factory Linked Service Synapse.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The ID of the Data Factory Synapse Linked Service.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The integration runtime reference to associate with the Data Factory Linked Service Synapse.
	IntegrationRuntimeName *string `json:"integrationRuntimeName,omitempty" tf:"integration_runtime_name,omitempty"`

	// A key_vault_password block as defined below. Use this argument to store Synapse password in an existing Key Vault. It needs an existing Key Vault Data Factory Linked Service.
	KeyVaultPassword []LinkedServiceSynapseKeyVaultPasswordObservation `json:"keyVaultPassword,omitempty" tf:"key_vault_password,omitempty"`

	// A map of parameters to associate with the Data Factory Linked Service Synapse.
	// +mapType=granular
	Parameters map[string]*string `json:"parameters,omitempty" tf:"parameters,omitempty"`
}

type LinkedServiceSynapseParameters struct {

	// A map of additional properties to associate with the Data Factory Linked Service Synapse.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	AdditionalProperties map[string]*string `json:"additionalProperties,omitempty" tf:"additional_properties,omitempty"`

	// List of tags that can be used for describing the Data Factory Linked Service Synapse.
	// +kubebuilder:validation:Optional
	Annotations []*string `json:"annotations,omitempty" tf:"annotations,omitempty"`

	// The connection string in which to authenticate with the Synapse.
	// +kubebuilder:validation:Optional
	ConnectionString *string `json:"connectionString,omitempty" tf:"connection_string,omitempty"`

	// The Data Factory ID in which to associate the Linked Service with. Changing this forces a new resource.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/datafactory/v1beta1.Factory
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	DataFactoryID *string `json:"dataFactoryId,omitempty" tf:"data_factory_id,omitempty"`

	// Reference to a Factory in datafactory to populate dataFactoryId.
	// +kubebuilder:validation:Optional
	DataFactoryIDRef *v1.Reference `json:"dataFactoryIdRef,omitempty" tf:"-"`

	// Selector for a Factory in datafactory to populate dataFactoryId.
	// +kubebuilder:validation:Optional
	DataFactoryIDSelector *v1.Selector `json:"dataFactoryIdSelector,omitempty" tf:"-"`

	// The description for the Data Factory Linked Service Synapse.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The integration runtime reference to associate with the Data Factory Linked Service Synapse.
	// +kubebuilder:validation:Optional
	IntegrationRuntimeName *string `json:"integrationRuntimeName,omitempty" tf:"integration_runtime_name,omitempty"`

	// A key_vault_password block as defined below. Use this argument to store Synapse password in an existing Key Vault. It needs an existing Key Vault Data Factory Linked Service.
	// +kubebuilder:validation:Optional
	KeyVaultPassword []LinkedServiceSynapseKeyVaultPasswordParameters `json:"keyVaultPassword,omitempty" tf:"key_vault_password,omitempty"`

	// A map of parameters to associate with the Data Factory Linked Service Synapse.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	Parameters map[string]*string `json:"parameters,omitempty" tf:"parameters,omitempty"`
}

// LinkedServiceSynapseSpec defines the desired state of LinkedServiceSynapse
type LinkedServiceSynapseSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     LinkedServiceSynapseParameters `json:"forProvider"`
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
	InitProvider LinkedServiceSynapseInitParameters `json:"initProvider,omitempty"`
}

// LinkedServiceSynapseStatus defines the observed state of LinkedServiceSynapse.
type LinkedServiceSynapseStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        LinkedServiceSynapseObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// LinkedServiceSynapse is the Schema for the LinkedServiceSynapses API. Manages a Linked Service (connection) between Synapse and Azure Data Factory.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type LinkedServiceSynapse struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.connectionString) || (has(self.initProvider) && has(self.initProvider.connectionString))",message="spec.forProvider.connectionString is a required parameter"
	Spec   LinkedServiceSynapseSpec   `json:"spec"`
	Status LinkedServiceSynapseStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// LinkedServiceSynapseList contains a list of LinkedServiceSynapses
type LinkedServiceSynapseList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []LinkedServiceSynapse `json:"items"`
}

// Repository type metadata.
var (
	LinkedServiceSynapse_Kind             = "LinkedServiceSynapse"
	LinkedServiceSynapse_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: LinkedServiceSynapse_Kind}.String()
	LinkedServiceSynapse_KindAPIVersion   = LinkedServiceSynapse_Kind + "." + CRDGroupVersion.String()
	LinkedServiceSynapse_GroupVersionKind = CRDGroupVersion.WithKind(LinkedServiceSynapse_Kind)
)

func init() {
	SchemeBuilder.Register(&LinkedServiceSynapse{}, &LinkedServiceSynapseList{})
}
