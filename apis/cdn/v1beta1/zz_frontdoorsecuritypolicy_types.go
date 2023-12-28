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

type AssociationInitParameters struct {

	// One or more domain blocks as defined below. Changing this forces a new Front Door Security Policy to be created.
	Domain []DomainInitParameters `json:"domain,omitempty" tf:"domain,omitempty"`

	// The list of paths to match for this firewall policy. Possible value includes /*. Changing this forces a new Front Door Security Policy to be created.
	PatternsToMatch []*string `json:"patternsToMatch,omitempty" tf:"patterns_to_match,omitempty"`
}

type AssociationObservation struct {

	// One or more domain blocks as defined below. Changing this forces a new Front Door Security Policy to be created.
	Domain []DomainObservation `json:"domain,omitempty" tf:"domain,omitempty"`

	// The list of paths to match for this firewall policy. Possible value includes /*. Changing this forces a new Front Door Security Policy to be created.
	PatternsToMatch []*string `json:"patternsToMatch,omitempty" tf:"patterns_to_match,omitempty"`
}

type AssociationParameters struct {

	// One or more domain blocks as defined below. Changing this forces a new Front Door Security Policy to be created.
	// +kubebuilder:validation:Optional
	Domain []DomainParameters `json:"domain" tf:"domain,omitempty"`

	// The list of paths to match for this firewall policy. Possible value includes /*. Changing this forces a new Front Door Security Policy to be created.
	// +kubebuilder:validation:Optional
	PatternsToMatch []*string `json:"patternsToMatch" tf:"patterns_to_match,omitempty"`
}

type DomainInitParameters struct {

	// The Resource Id of the Front Door Custom Domain or Front Door Endpoint that should be bound to this Front Door Security Policy. Changing this forces a new Front Door Security Policy to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/cdn/v1beta1.FrontdoorCustomDomain
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	CdnFrontdoorDomainID *string `json:"cdnFrontdoorDomainId,omitempty" tf:"cdn_frontdoor_domain_id,omitempty"`

	// Reference to a FrontdoorCustomDomain in cdn to populate cdnFrontdoorDomainId.
	// +kubebuilder:validation:Optional
	CdnFrontdoorDomainIDRef *v1.Reference `json:"cdnFrontdoorDomainIdRef,omitempty" tf:"-"`

	// Selector for a FrontdoorCustomDomain in cdn to populate cdnFrontdoorDomainId.
	// +kubebuilder:validation:Optional
	CdnFrontdoorDomainIDSelector *v1.Selector `json:"cdnFrontdoorDomainIdSelector,omitempty" tf:"-"`
}

type DomainObservation struct {

	// (Computed) Is the Front Door Custom Domain/Endpoint activated?
	Active *bool `json:"active,omitempty" tf:"active,omitempty"`

	// The Resource Id of the Front Door Custom Domain or Front Door Endpoint that should be bound to this Front Door Security Policy. Changing this forces a new Front Door Security Policy to be created.
	CdnFrontdoorDomainID *string `json:"cdnFrontdoorDomainId,omitempty" tf:"cdn_frontdoor_domain_id,omitempty"`
}

type DomainParameters struct {

	// The Resource Id of the Front Door Custom Domain or Front Door Endpoint that should be bound to this Front Door Security Policy. Changing this forces a new Front Door Security Policy to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/cdn/v1beta1.FrontdoorCustomDomain
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	CdnFrontdoorDomainID *string `json:"cdnFrontdoorDomainId,omitempty" tf:"cdn_frontdoor_domain_id,omitempty"`

	// Reference to a FrontdoorCustomDomain in cdn to populate cdnFrontdoorDomainId.
	// +kubebuilder:validation:Optional
	CdnFrontdoorDomainIDRef *v1.Reference `json:"cdnFrontdoorDomainIdRef,omitempty" tf:"-"`

	// Selector for a FrontdoorCustomDomain in cdn to populate cdnFrontdoorDomainId.
	// +kubebuilder:validation:Optional
	CdnFrontdoorDomainIDSelector *v1.Selector `json:"cdnFrontdoorDomainIdSelector,omitempty" tf:"-"`
}

type FirewallInitParameters struct {

	// An association block as defined below. Changing this forces a new Front Door Security Policy to be created.
	Association []AssociationInitParameters `json:"association,omitempty" tf:"association,omitempty"`

	// The Resource Id of the Front Door Firewall Policy that should be linked to this Front Door Security Policy. Changing this forces a new Front Door Security Policy to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/cdn/v1beta1.FrontdoorFirewallPolicy
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	CdnFrontdoorFirewallPolicyID *string `json:"cdnFrontdoorFirewallPolicyId,omitempty" tf:"cdn_frontdoor_firewall_policy_id,omitempty"`

	// Reference to a FrontdoorFirewallPolicy in cdn to populate cdnFrontdoorFirewallPolicyId.
	// +kubebuilder:validation:Optional
	CdnFrontdoorFirewallPolicyIDRef *v1.Reference `json:"cdnFrontdoorFirewallPolicyIdRef,omitempty" tf:"-"`

	// Selector for a FrontdoorFirewallPolicy in cdn to populate cdnFrontdoorFirewallPolicyId.
	// +kubebuilder:validation:Optional
	CdnFrontdoorFirewallPolicyIDSelector *v1.Selector `json:"cdnFrontdoorFirewallPolicyIdSelector,omitempty" tf:"-"`
}

type FirewallObservation struct {

	// An association block as defined below. Changing this forces a new Front Door Security Policy to be created.
	Association []AssociationObservation `json:"association,omitempty" tf:"association,omitempty"`

	// The Resource Id of the Front Door Firewall Policy that should be linked to this Front Door Security Policy. Changing this forces a new Front Door Security Policy to be created.
	CdnFrontdoorFirewallPolicyID *string `json:"cdnFrontdoorFirewallPolicyId,omitempty" tf:"cdn_frontdoor_firewall_policy_id,omitempty"`
}

type FirewallParameters struct {

	// An association block as defined below. Changing this forces a new Front Door Security Policy to be created.
	// +kubebuilder:validation:Optional
	Association []AssociationParameters `json:"association" tf:"association,omitempty"`

	// The Resource Id of the Front Door Firewall Policy that should be linked to this Front Door Security Policy. Changing this forces a new Front Door Security Policy to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/cdn/v1beta1.FrontdoorFirewallPolicy
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	CdnFrontdoorFirewallPolicyID *string `json:"cdnFrontdoorFirewallPolicyId,omitempty" tf:"cdn_frontdoor_firewall_policy_id,omitempty"`

	// Reference to a FrontdoorFirewallPolicy in cdn to populate cdnFrontdoorFirewallPolicyId.
	// +kubebuilder:validation:Optional
	CdnFrontdoorFirewallPolicyIDRef *v1.Reference `json:"cdnFrontdoorFirewallPolicyIdRef,omitempty" tf:"-"`

	// Selector for a FrontdoorFirewallPolicy in cdn to populate cdnFrontdoorFirewallPolicyId.
	// +kubebuilder:validation:Optional
	CdnFrontdoorFirewallPolicyIDSelector *v1.Selector `json:"cdnFrontdoorFirewallPolicyIdSelector,omitempty" tf:"-"`
}

type FrontdoorSecurityPolicyInitParameters struct {

	// An security_policies block as defined below. Changing this forces a new Front Door Security Policy to be created.
	SecurityPolicies []SecurityPoliciesInitParameters `json:"securityPolicies,omitempty" tf:"security_policies,omitempty"`
}

type FrontdoorSecurityPolicyObservation struct {

	// The Front Door Profile Resource Id that is linked to this Front Door Security Policy. Changing this forces a new Front Door Security Policy to be created.
	CdnFrontdoorProfileID *string `json:"cdnFrontdoorProfileId,omitempty" tf:"cdn_frontdoor_profile_id,omitempty"`

	// The ID of the Front Door Security Policy.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// An security_policies block as defined below. Changing this forces a new Front Door Security Policy to be created.
	SecurityPolicies []SecurityPoliciesObservation `json:"securityPolicies,omitempty" tf:"security_policies,omitempty"`
}

type FrontdoorSecurityPolicyParameters struct {

	// The Front Door Profile Resource Id that is linked to this Front Door Security Policy. Changing this forces a new Front Door Security Policy to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/cdn/v1beta1.FrontdoorProfile
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	CdnFrontdoorProfileID *string `json:"cdnFrontdoorProfileId,omitempty" tf:"cdn_frontdoor_profile_id,omitempty"`

	// Reference to a FrontdoorProfile in cdn to populate cdnFrontdoorProfileId.
	// +kubebuilder:validation:Optional
	CdnFrontdoorProfileIDRef *v1.Reference `json:"cdnFrontdoorProfileIdRef,omitempty" tf:"-"`

	// Selector for a FrontdoorProfile in cdn to populate cdnFrontdoorProfileId.
	// +kubebuilder:validation:Optional
	CdnFrontdoorProfileIDSelector *v1.Selector `json:"cdnFrontdoorProfileIdSelector,omitempty" tf:"-"`

	// An security_policies block as defined below. Changing this forces a new Front Door Security Policy to be created.
	// +kubebuilder:validation:Optional
	SecurityPolicies []SecurityPoliciesParameters `json:"securityPolicies,omitempty" tf:"security_policies,omitempty"`
}

type SecurityPoliciesInitParameters struct {

	// An firewall block as defined below. Changing this forces a new Front Door Security Policy to be created.
	Firewall []FirewallInitParameters `json:"firewall,omitempty" tf:"firewall,omitempty"`
}

type SecurityPoliciesObservation struct {

	// An firewall block as defined below. Changing this forces a new Front Door Security Policy to be created.
	Firewall []FirewallObservation `json:"firewall,omitempty" tf:"firewall,omitempty"`
}

type SecurityPoliciesParameters struct {

	// An firewall block as defined below. Changing this forces a new Front Door Security Policy to be created.
	// +kubebuilder:validation:Optional
	Firewall []FirewallParameters `json:"firewall" tf:"firewall,omitempty"`
}

// FrontdoorSecurityPolicySpec defines the desired state of FrontdoorSecurityPolicy
type FrontdoorSecurityPolicySpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     FrontdoorSecurityPolicyParameters `json:"forProvider"`
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
	InitProvider FrontdoorSecurityPolicyInitParameters `json:"initProvider,omitempty"`
}

// FrontdoorSecurityPolicyStatus defines the observed state of FrontdoorSecurityPolicy.
type FrontdoorSecurityPolicyStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        FrontdoorSecurityPolicyObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// FrontdoorSecurityPolicy is the Schema for the FrontdoorSecurityPolicys API. Manages a Front Door (standard/premium) Security Policy.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type FrontdoorSecurityPolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.securityPolicies) || (has(self.initProvider) && has(self.initProvider.securityPolicies))",message="spec.forProvider.securityPolicies is a required parameter"
	Spec   FrontdoorSecurityPolicySpec   `json:"spec"`
	Status FrontdoorSecurityPolicyStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// FrontdoorSecurityPolicyList contains a list of FrontdoorSecurityPolicys
type FrontdoorSecurityPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []FrontdoorSecurityPolicy `json:"items"`
}

// Repository type metadata.
var (
	FrontdoorSecurityPolicy_Kind             = "FrontdoorSecurityPolicy"
	FrontdoorSecurityPolicy_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: FrontdoorSecurityPolicy_Kind}.String()
	FrontdoorSecurityPolicy_KindAPIVersion   = FrontdoorSecurityPolicy_Kind + "." + CRDGroupVersion.String()
	FrontdoorSecurityPolicy_GroupVersionKind = CRDGroupVersion.WithKind(FrontdoorSecurityPolicy_Kind)
)

func init() {
	SchemeBuilder.Register(&FrontdoorSecurityPolicy{}, &FrontdoorSecurityPolicyList{})
}
