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

type EnvironmentInitParameters struct {

	// The existing Subnet to use for the Container Apps Control Plane. Changing this forces a new resource to be created.
	// The existing Subnet to use for the Container Apps Control Plane. **NOTE:** The Subnet must have a `/21` or larger address space.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/network/v1beta1.Subnet
	// +crossplane:generate:reference:extractor=github.com/upbound/provider-azure/apis/rconfig.ExtractResourceID()
	InfrastructureSubnetID *string `json:"infrastructureSubnetId,omitempty" tf:"infrastructure_subnet_id,omitempty"`

	// Reference to a Subnet in network to populate infrastructureSubnetId.
	// +kubebuilder:validation:Optional
	InfrastructureSubnetIDRef *v1.Reference `json:"infrastructureSubnetIdRef,omitempty" tf:"-"`

	// Selector for a Subnet in network to populate infrastructureSubnetId.
	// +kubebuilder:validation:Optional
	InfrastructureSubnetIDSelector *v1.Selector `json:"infrastructureSubnetIdSelector,omitempty" tf:"-"`

	// Should the Container Environment operate in Internal Load Balancing Mode? Defaults to false. Changing this forces a new resource to be created.
	// Should the Container Environment operate in Internal Load Balancing Mode? Defaults to `false`. **Note:** can only be set to `true` if `infrastructure_subnet_id` is specified.
	InternalLoadBalancerEnabled *bool `json:"internalLoadBalancerEnabled,omitempty" tf:"internal_load_balancer_enabled,omitempty"`

	// Specifies the supported Azure location where the Container App Environment is to exist. Changing this forces a new resource to be created.
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// The ID for the Log Analytics Workspace to link this Container Apps Managed Environment to. Changing this forces a new resource to be created.
	// The ID for the Log Analytics Workspace to link this Container Apps Managed Environment to.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/operationalinsights/v1beta1.Workspace
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	LogAnalyticsWorkspaceID *string `json:"logAnalyticsWorkspaceId,omitempty" tf:"log_analytics_workspace_id,omitempty"`

	// Reference to a Workspace in operationalinsights to populate logAnalyticsWorkspaceId.
	// +kubebuilder:validation:Optional
	LogAnalyticsWorkspaceIDRef *v1.Reference `json:"logAnalyticsWorkspaceIdRef,omitempty" tf:"-"`

	// Selector for a Workspace in operationalinsights to populate logAnalyticsWorkspaceId.
	// +kubebuilder:validation:Optional
	LogAnalyticsWorkspaceIDSelector *v1.Selector `json:"logAnalyticsWorkspaceIdSelector,omitempty" tf:"-"`

	// A mapping of tags to assign to the resource.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type EnvironmentObservation struct {

	// The default, publicly resolvable, name of this Container App Environment.
	// The default publicly resolvable name of this Container App Environment
	DefaultDomain *string `json:"defaultDomain,omitempty" tf:"default_domain,omitempty"`

	// The network addressing in which the Container Apps in this Container App Environment will reside in CIDR notation.
	// The network addressing in which the Container Apps in this Container App Environment will reside in CIDR notation.
	DockerBridgeCidr *string `json:"dockerBridgeCidr,omitempty" tf:"docker_bridge_cidr,omitempty"`

	// The ID of the Container App Environment
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The existing Subnet to use for the Container Apps Control Plane. Changing this forces a new resource to be created.
	// The existing Subnet to use for the Container Apps Control Plane. **NOTE:** The Subnet must have a `/21` or larger address space.
	InfrastructureSubnetID *string `json:"infrastructureSubnetId,omitempty" tf:"infrastructure_subnet_id,omitempty"`

	// Should the Container Environment operate in Internal Load Balancing Mode? Defaults to false. Changing this forces a new resource to be created.
	// Should the Container Environment operate in Internal Load Balancing Mode? Defaults to `false`. **Note:** can only be set to `true` if `infrastructure_subnet_id` is specified.
	InternalLoadBalancerEnabled *bool `json:"internalLoadBalancerEnabled,omitempty" tf:"internal_load_balancer_enabled,omitempty"`

	// Specifies the supported Azure location where the Container App Environment is to exist. Changing this forces a new resource to be created.
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// The ID for the Log Analytics Workspace to link this Container Apps Managed Environment to. Changing this forces a new resource to be created.
	// The ID for the Log Analytics Workspace to link this Container Apps Managed Environment to.
	LogAnalyticsWorkspaceID *string `json:"logAnalyticsWorkspaceId,omitempty" tf:"log_analytics_workspace_id,omitempty"`

	// The IP range, in CIDR notation, that is reserved for environment infrastructure IP addresses.
	// The IP range, in CIDR notation, that is reserved for environment infrastructure IP addresses.
	PlatformReservedCidr *string `json:"platformReservedCidr,omitempty" tf:"platform_reserved_cidr,omitempty"`

	// The IP address from the IP range defined by platform_reserved_cidr that is reserved for the internal DNS server.
	// The IP address from the IP range defined by `platform_reserved_cidr` that is reserved for the internal DNS server.
	PlatformReservedDNSIPAddress *string `json:"platformReservedDnsIpAddress,omitempty" tf:"platform_reserved_dns_ip_address,omitempty"`

	// The name of the resource group in which the Container App Environment is to be created. Changing this forces a new resource to be created.
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// The Static IP address of the Environment.
	// The Static IP Address of the Environment.
	StaticIPAddress *string `json:"staticIpAddress,omitempty" tf:"static_ip_address,omitempty"`

	// A mapping of tags to assign to the resource.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type EnvironmentParameters struct {

	// The existing Subnet to use for the Container Apps Control Plane. Changing this forces a new resource to be created.
	// The existing Subnet to use for the Container Apps Control Plane. **NOTE:** The Subnet must have a `/21` or larger address space.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/network/v1beta1.Subnet
	// +crossplane:generate:reference:extractor=github.com/upbound/provider-azure/apis/rconfig.ExtractResourceID()
	// +kubebuilder:validation:Optional
	InfrastructureSubnetID *string `json:"infrastructureSubnetId,omitempty" tf:"infrastructure_subnet_id,omitempty"`

	// Reference to a Subnet in network to populate infrastructureSubnetId.
	// +kubebuilder:validation:Optional
	InfrastructureSubnetIDRef *v1.Reference `json:"infrastructureSubnetIdRef,omitempty" tf:"-"`

	// Selector for a Subnet in network to populate infrastructureSubnetId.
	// +kubebuilder:validation:Optional
	InfrastructureSubnetIDSelector *v1.Selector `json:"infrastructureSubnetIdSelector,omitempty" tf:"-"`

	// Should the Container Environment operate in Internal Load Balancing Mode? Defaults to false. Changing this forces a new resource to be created.
	// Should the Container Environment operate in Internal Load Balancing Mode? Defaults to `false`. **Note:** can only be set to `true` if `infrastructure_subnet_id` is specified.
	// +kubebuilder:validation:Optional
	InternalLoadBalancerEnabled *bool `json:"internalLoadBalancerEnabled,omitempty" tf:"internal_load_balancer_enabled,omitempty"`

	// Specifies the supported Azure location where the Container App Environment is to exist. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// The ID for the Log Analytics Workspace to link this Container Apps Managed Environment to. Changing this forces a new resource to be created.
	// The ID for the Log Analytics Workspace to link this Container Apps Managed Environment to.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/operationalinsights/v1beta1.Workspace
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	LogAnalyticsWorkspaceID *string `json:"logAnalyticsWorkspaceId,omitempty" tf:"log_analytics_workspace_id,omitempty"`

	// Reference to a Workspace in operationalinsights to populate logAnalyticsWorkspaceId.
	// +kubebuilder:validation:Optional
	LogAnalyticsWorkspaceIDRef *v1.Reference `json:"logAnalyticsWorkspaceIdRef,omitempty" tf:"-"`

	// Selector for a Workspace in operationalinsights to populate logAnalyticsWorkspaceId.
	// +kubebuilder:validation:Optional
	LogAnalyticsWorkspaceIDSelector *v1.Selector `json:"logAnalyticsWorkspaceIdSelector,omitempty" tf:"-"`

	// The name of the resource group in which the Container App Environment is to be created. Changing this forces a new resource to be created.
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
}

// EnvironmentSpec defines the desired state of Environment
type EnvironmentSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     EnvironmentParameters `json:"forProvider"`
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
	InitProvider EnvironmentInitParameters `json:"initProvider,omitempty"`
}

// EnvironmentStatus defines the observed state of Environment.
type EnvironmentStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        EnvironmentObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Environment is the Schema for the Environments API. Manages a Container App Environment.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type Environment struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.location) || (has(self.initProvider) && has(self.initProvider.location))",message="spec.forProvider.location is a required parameter"
	Spec   EnvironmentSpec   `json:"spec"`
	Status EnvironmentStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// EnvironmentList contains a list of Environments
type EnvironmentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Environment `json:"items"`
}

// Repository type metadata.
var (
	Environment_Kind             = "Environment"
	Environment_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Environment_Kind}.String()
	Environment_KindAPIVersion   = Environment_Kind + "." + CRDGroupVersion.String()
	Environment_GroupVersionKind = CRDGroupVersion.WithKind(Environment_Kind)
)

func init() {
	SchemeBuilder.Register(&Environment{}, &EnvironmentList{})
}
