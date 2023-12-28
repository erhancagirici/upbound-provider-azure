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

type SentinelAlertRuleMachineLearningBehaviorAnalyticsInitParameters struct {

	// The GUID of the alert rule template which is used for this Sentinel Machine Learning Behavior Analytics Alert Rule. Changing this forces a new Sentinel Machine Learning Behavior Analytics Alert Rule to be created.
	AlertRuleTemplateGUID *string `json:"alertRuleTemplateGuid,omitempty" tf:"alert_rule_template_guid,omitempty"`

	// Should this Sentinel Machine Learning Behavior Analytics Alert Rule be enabled? Defaults to true.
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// The ID of the Log Analytics Workspace this SentinelMachine Learning Behavior Analytics Alert Rule belongs to. Changing this forces a new Sentinel Machine Learning Behavior Analytics Alert Rule to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/operationalinsights/v1beta1.Workspace
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	LogAnalyticsWorkspaceID *string `json:"logAnalyticsWorkspaceId,omitempty" tf:"log_analytics_workspace_id,omitempty"`

	// Reference to a Workspace in operationalinsights to populate logAnalyticsWorkspaceId.
	// +kubebuilder:validation:Optional
	LogAnalyticsWorkspaceIDRef *v1.Reference `json:"logAnalyticsWorkspaceIdRef,omitempty" tf:"-"`

	// Selector for a Workspace in operationalinsights to populate logAnalyticsWorkspaceId.
	// +kubebuilder:validation:Optional
	LogAnalyticsWorkspaceIDSelector *v1.Selector `json:"logAnalyticsWorkspaceIdSelector,omitempty" tf:"-"`

	// The name which should be used for this SentinelMachine Learning Behavior Analytics Alert Rule. Changing this forces a new Sentinel Machine Learning Behavior Analytics Alert Rule to be created.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`
}

type SentinelAlertRuleMachineLearningBehaviorAnalyticsObservation struct {

	// The GUID of the alert rule template which is used for this Sentinel Machine Learning Behavior Analytics Alert Rule. Changing this forces a new Sentinel Machine Learning Behavior Analytics Alert Rule to be created.
	AlertRuleTemplateGUID *string `json:"alertRuleTemplateGuid,omitempty" tf:"alert_rule_template_guid,omitempty"`

	// Should this Sentinel Machine Learning Behavior Analytics Alert Rule be enabled? Defaults to true.
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// The ID of the Sentinel Machine Learning Behavior Analytics Alert Rule.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The ID of the Log Analytics Workspace this SentinelMachine Learning Behavior Analytics Alert Rule belongs to. Changing this forces a new Sentinel Machine Learning Behavior Analytics Alert Rule to be created.
	LogAnalyticsWorkspaceID *string `json:"logAnalyticsWorkspaceId,omitempty" tf:"log_analytics_workspace_id,omitempty"`

	// The name which should be used for this SentinelMachine Learning Behavior Analytics Alert Rule. Changing this forces a new Sentinel Machine Learning Behavior Analytics Alert Rule to be created.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`
}

type SentinelAlertRuleMachineLearningBehaviorAnalyticsParameters struct {

	// The GUID of the alert rule template which is used for this Sentinel Machine Learning Behavior Analytics Alert Rule. Changing this forces a new Sentinel Machine Learning Behavior Analytics Alert Rule to be created.
	// +kubebuilder:validation:Optional
	AlertRuleTemplateGUID *string `json:"alertRuleTemplateGuid,omitempty" tf:"alert_rule_template_guid,omitempty"`

	// Should this Sentinel Machine Learning Behavior Analytics Alert Rule be enabled? Defaults to true.
	// +kubebuilder:validation:Optional
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// The ID of the Log Analytics Workspace this SentinelMachine Learning Behavior Analytics Alert Rule belongs to. Changing this forces a new Sentinel Machine Learning Behavior Analytics Alert Rule to be created.
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

	// The name which should be used for this SentinelMachine Learning Behavior Analytics Alert Rule. Changing this forces a new Sentinel Machine Learning Behavior Analytics Alert Rule to be created.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`
}

// SentinelAlertRuleMachineLearningBehaviorAnalyticsSpec defines the desired state of SentinelAlertRuleMachineLearningBehaviorAnalytics
type SentinelAlertRuleMachineLearningBehaviorAnalyticsSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     SentinelAlertRuleMachineLearningBehaviorAnalyticsParameters `json:"forProvider"`
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
	InitProvider SentinelAlertRuleMachineLearningBehaviorAnalyticsInitParameters `json:"initProvider,omitempty"`
}

// SentinelAlertRuleMachineLearningBehaviorAnalyticsStatus defines the observed state of SentinelAlertRuleMachineLearningBehaviorAnalytics.
type SentinelAlertRuleMachineLearningBehaviorAnalyticsStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        SentinelAlertRuleMachineLearningBehaviorAnalyticsObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// SentinelAlertRuleMachineLearningBehaviorAnalytics is the Schema for the SentinelAlertRuleMachineLearningBehaviorAnalyticss API. Manages a Sentinel Machine Learning Behavior Analytics Rule.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type SentinelAlertRuleMachineLearningBehaviorAnalytics struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.alertRuleTemplateGuid) || (has(self.initProvider) && has(self.initProvider.alertRuleTemplateGuid))",message="spec.forProvider.alertRuleTemplateGuid is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	Spec   SentinelAlertRuleMachineLearningBehaviorAnalyticsSpec   `json:"spec"`
	Status SentinelAlertRuleMachineLearningBehaviorAnalyticsStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SentinelAlertRuleMachineLearningBehaviorAnalyticsList contains a list of SentinelAlertRuleMachineLearningBehaviorAnalyticss
type SentinelAlertRuleMachineLearningBehaviorAnalyticsList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SentinelAlertRuleMachineLearningBehaviorAnalytics `json:"items"`
}

// Repository type metadata.
var (
	SentinelAlertRuleMachineLearningBehaviorAnalytics_Kind             = "SentinelAlertRuleMachineLearningBehaviorAnalytics"
	SentinelAlertRuleMachineLearningBehaviorAnalytics_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: SentinelAlertRuleMachineLearningBehaviorAnalytics_Kind}.String()
	SentinelAlertRuleMachineLearningBehaviorAnalytics_KindAPIVersion   = SentinelAlertRuleMachineLearningBehaviorAnalytics_Kind + "." + CRDGroupVersion.String()
	SentinelAlertRuleMachineLearningBehaviorAnalytics_GroupVersionKind = CRDGroupVersion.WithKind(SentinelAlertRuleMachineLearningBehaviorAnalytics_Kind)
)

func init() {
	SchemeBuilder.Register(&SentinelAlertRuleMachineLearningBehaviorAnalytics{}, &SentinelAlertRuleMachineLearningBehaviorAnalyticsList{})
}
