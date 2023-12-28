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

type AppServiceCertificateOrderInitParameters struct {

	// true if the certificate should be automatically renewed when it expires; otherwise, false. Defaults to true.
	AutoRenew *bool `json:"autoRenew,omitempty" tf:"auto_renew,omitempty"`

	// Last CSR that was created for this order.
	Csr *string `json:"csr,omitempty" tf:"csr,omitempty"`

	// The Distinguished Name for the App Service Certificate Order.
	DistinguishedName *string `json:"distinguishedName,omitempty" tf:"distinguished_name,omitempty"`

	// Certificate key size. Defaults to 2048.
	KeySize *float64 `json:"keySize,omitempty" tf:"key_size,omitempty"`

	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created. Currently the only valid value is global.
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// Certificate product type, such as Standard or WildCard.
	ProductType *string `json:"productType,omitempty" tf:"product_type,omitempty"`

	// A mapping of tags to assign to the resource.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// Duration in years (must be between 1 and 3). Defaults to 1.
	ValidityInYears *float64 `json:"validityInYears,omitempty" tf:"validity_in_years,omitempty"`
}

type AppServiceCertificateOrderObservation struct {

	// Reasons why App Service Certificate is not renewable at the current moment.
	AppServiceCertificateNotRenewableReasons []*string `json:"appServiceCertificateNotRenewableReasons,omitempty" tf:"app_service_certificate_not_renewable_reasons,omitempty"`

	// true if the certificate should be automatically renewed when it expires; otherwise, false. Defaults to true.
	AutoRenew *bool `json:"autoRenew,omitempty" tf:"auto_renew,omitempty"`

	// State of the Key Vault secret. A certificates block as defined below.
	Certificates []CertificatesObservation `json:"certificates,omitempty" tf:"certificates,omitempty"`

	// Last CSR that was created for this order.
	Csr *string `json:"csr,omitempty" tf:"csr,omitempty"`

	// The Distinguished Name for the App Service Certificate Order.
	DistinguishedName *string `json:"distinguishedName,omitempty" tf:"distinguished_name,omitempty"`

	// Domain verification token.
	DomainVerificationToken *string `json:"domainVerificationToken,omitempty" tf:"domain_verification_token,omitempty"`

	// Certificate expiration time.
	ExpirationTime *string `json:"expirationTime,omitempty" tf:"expiration_time,omitempty"`

	// The App Service Certificate Order ID.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Certificate thumbprint intermediate certificate.
	IntermediateThumbprint *string `json:"intermediateThumbprint,omitempty" tf:"intermediate_thumbprint,omitempty"`

	// Whether the private key is external or not.
	IsPrivateKeyExternal *bool `json:"isPrivateKeyExternal,omitempty" tf:"is_private_key_external,omitempty"`

	// Certificate key size. Defaults to 2048.
	KeySize *float64 `json:"keySize,omitempty" tf:"key_size,omitempty"`

	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created. Currently the only valid value is global.
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// Certificate product type, such as Standard or WildCard.
	ProductType *string `json:"productType,omitempty" tf:"product_type,omitempty"`

	// The name of the resource group in which to create the certificate. Changing this forces a new resource to be created.
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// Certificate thumbprint for root certificate.
	RootThumbprint *string `json:"rootThumbprint,omitempty" tf:"root_thumbprint,omitempty"`

	// Certificate thumbprint for signed certificate.
	SignedCertificateThumbprint *string `json:"signedCertificateThumbprint,omitempty" tf:"signed_certificate_thumbprint,omitempty"`

	// Current order status.
	Status *string `json:"status,omitempty" tf:"status,omitempty"`

	// A mapping of tags to assign to the resource.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// Duration in years (must be between 1 and 3). Defaults to 1.
	ValidityInYears *float64 `json:"validityInYears,omitempty" tf:"validity_in_years,omitempty"`
}

type AppServiceCertificateOrderParameters struct {

	// true if the certificate should be automatically renewed when it expires; otherwise, false. Defaults to true.
	// +kubebuilder:validation:Optional
	AutoRenew *bool `json:"autoRenew,omitempty" tf:"auto_renew,omitempty"`

	// Last CSR that was created for this order.
	// +kubebuilder:validation:Optional
	Csr *string `json:"csr,omitempty" tf:"csr,omitempty"`

	// The Distinguished Name for the App Service Certificate Order.
	// +kubebuilder:validation:Optional
	DistinguishedName *string `json:"distinguishedName,omitempty" tf:"distinguished_name,omitempty"`

	// Certificate key size. Defaults to 2048.
	// +kubebuilder:validation:Optional
	KeySize *float64 `json:"keySize,omitempty" tf:"key_size,omitempty"`

	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created. Currently the only valid value is global.
	// +kubebuilder:validation:Optional
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// Certificate product type, such as Standard or WildCard.
	// +kubebuilder:validation:Optional
	ProductType *string `json:"productType,omitempty" tf:"product_type,omitempty"`

	// The name of the resource group in which to create the certificate. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/azure/v1beta1.ResourceGroup
	// +kubebuilder:validation:Optional
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// Reference to a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// Selector for a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`

	// A mapping of tags to assign to the resource.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// Duration in years (must be between 1 and 3). Defaults to 1.
	// +kubebuilder:validation:Optional
	ValidityInYears *float64 `json:"validityInYears,omitempty" tf:"validity_in_years,omitempty"`
}

type CertificatesInitParameters struct {
}

type CertificatesObservation struct {

	// The name of the App Service Certificate.
	CertificateName *string `json:"certificateName,omitempty" tf:"certificate_name,omitempty"`

	// Key Vault resource Id.
	KeyVaultID *string `json:"keyVaultId,omitempty" tf:"key_vault_id,omitempty"`

	// Key Vault secret name.
	KeyVaultSecretName *string `json:"keyVaultSecretName,omitempty" tf:"key_vault_secret_name,omitempty"`

	// Status of the Key Vault secret.
	ProvisioningState *string `json:"provisioningState,omitempty" tf:"provisioning_state,omitempty"`
}

type CertificatesParameters struct {
}

// AppServiceCertificateOrderSpec defines the desired state of AppServiceCertificateOrder
type AppServiceCertificateOrderSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     AppServiceCertificateOrderParameters `json:"forProvider"`
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
	InitProvider AppServiceCertificateOrderInitParameters `json:"initProvider,omitempty"`
}

// AppServiceCertificateOrderStatus defines the observed state of AppServiceCertificateOrder.
type AppServiceCertificateOrderStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        AppServiceCertificateOrderObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// AppServiceCertificateOrder is the Schema for the AppServiceCertificateOrders API. Manages an App Service Certificate Order.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type AppServiceCertificateOrder struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.location) || (has(self.initProvider) && has(self.initProvider.location))",message="spec.forProvider.location is a required parameter"
	Spec   AppServiceCertificateOrderSpec   `json:"spec"`
	Status AppServiceCertificateOrderStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// AppServiceCertificateOrderList contains a list of AppServiceCertificateOrders
type AppServiceCertificateOrderList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AppServiceCertificateOrder `json:"items"`
}

// Repository type metadata.
var (
	AppServiceCertificateOrder_Kind             = "AppServiceCertificateOrder"
	AppServiceCertificateOrder_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: AppServiceCertificateOrder_Kind}.String()
	AppServiceCertificateOrder_KindAPIVersion   = AppServiceCertificateOrder_Kind + "." + CRDGroupVersion.String()
	AppServiceCertificateOrder_GroupVersionKind = CRDGroupVersion.WithKind(AppServiceCertificateOrder_Kind)
)

func init() {
	SchemeBuilder.Register(&AppServiceCertificateOrder{}, &AppServiceCertificateOrderList{})
}
