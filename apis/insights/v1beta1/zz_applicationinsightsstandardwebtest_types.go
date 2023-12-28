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

type ApplicationInsightsStandardWebTestInitParameters struct {

	// The ID of the Application Insights instance on which the WebTest operates. Changing this forces a new Application Insights Standard WebTest to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/insights/v1beta1.ApplicationInsights
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	ApplicationInsightsID *string `json:"applicationInsightsId,omitempty" tf:"application_insights_id,omitempty"`

	// Reference to a ApplicationInsights in insights to populate applicationInsightsId.
	// +kubebuilder:validation:Optional
	ApplicationInsightsIDRef *v1.Reference `json:"applicationInsightsIdRef,omitempty" tf:"-"`

	// Selector for a ApplicationInsights in insights to populate applicationInsightsId.
	// +kubebuilder:validation:Optional
	ApplicationInsightsIDSelector *v1.Selector `json:"applicationInsightsIdSelector,omitempty" tf:"-"`

	// Purpose/user defined descriptive test for this WebTest.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Should the WebTest be enabled?
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// Interval in seconds between test runs for this WebTest. Valid options are 300, 600 and 900. Defaults to 300.
	Frequency *float64 `json:"frequency,omitempty" tf:"frequency,omitempty"`

	// Specifies a list of where to physically run the tests from to give global coverage for accessibility of your application.
	GeoLocations []*string `json:"geoLocations,omitempty" tf:"geo_locations,omitempty"`

	// The Azure Region where the Application Insights Standard WebTest should exist. Changing this forces a new Application Insights Standard WebTest to be created. It needs to correlate with location of the parent resource (azurerm_application_insights)
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// A request block as defined below.
	Request []RequestInitParameters `json:"request,omitempty" tf:"request,omitempty"`

	// Should the retry on WebTest failure be enabled?
	RetryEnabled *bool `json:"retryEnabled,omitempty" tf:"retry_enabled,omitempty"`

	// A mapping of tags which should be assigned to the Application Insights Standard WebTest.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// Seconds until this WebTest will timeout and fail. Default is 30.
	Timeout *float64 `json:"timeout,omitempty" tf:"timeout,omitempty"`

	// A validation_rules block as defined below.
	ValidationRules []ValidationRulesInitParameters `json:"validationRules,omitempty" tf:"validation_rules,omitempty"`
}

type ApplicationInsightsStandardWebTestObservation struct {

	// The ID of the Application Insights instance on which the WebTest operates. Changing this forces a new Application Insights Standard WebTest to be created.
	ApplicationInsightsID *string `json:"applicationInsightsId,omitempty" tf:"application_insights_id,omitempty"`

	// Purpose/user defined descriptive test for this WebTest.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Should the WebTest be enabled?
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// Interval in seconds between test runs for this WebTest. Valid options are 300, 600 and 900. Defaults to 300.
	Frequency *float64 `json:"frequency,omitempty" tf:"frequency,omitempty"`

	// Specifies a list of where to physically run the tests from to give global coverage for accessibility of your application.
	GeoLocations []*string `json:"geoLocations,omitempty" tf:"geo_locations,omitempty"`

	// The ID of the Application Insights Standard WebTest.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The Azure Region where the Application Insights Standard WebTest should exist. Changing this forces a new Application Insights Standard WebTest to be created. It needs to correlate with location of the parent resource (azurerm_application_insights)
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// A request block as defined below.
	Request []RequestObservation `json:"request,omitempty" tf:"request,omitempty"`

	// The name of the Resource Group where the Application Insights Standard WebTest should exist. Changing this forces a new Application Insights Standard WebTest to be created.
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// Should the retry on WebTest failure be enabled?
	RetryEnabled *bool `json:"retryEnabled,omitempty" tf:"retry_enabled,omitempty"`

	// Unique ID of this WebTest. This is typically the same value as the Name field.
	SyntheticMonitorID *string `json:"syntheticMonitorId,omitempty" tf:"synthetic_monitor_id,omitempty"`

	// A mapping of tags which should be assigned to the Application Insights Standard WebTest.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// Seconds until this WebTest will timeout and fail. Default is 30.
	Timeout *float64 `json:"timeout,omitempty" tf:"timeout,omitempty"`

	// A validation_rules block as defined below.
	ValidationRules []ValidationRulesObservation `json:"validationRules,omitempty" tf:"validation_rules,omitempty"`
}

type ApplicationInsightsStandardWebTestParameters struct {

	// The ID of the Application Insights instance on which the WebTest operates. Changing this forces a new Application Insights Standard WebTest to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/insights/v1beta1.ApplicationInsights
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	ApplicationInsightsID *string `json:"applicationInsightsId,omitempty" tf:"application_insights_id,omitempty"`

	// Reference to a ApplicationInsights in insights to populate applicationInsightsId.
	// +kubebuilder:validation:Optional
	ApplicationInsightsIDRef *v1.Reference `json:"applicationInsightsIdRef,omitempty" tf:"-"`

	// Selector for a ApplicationInsights in insights to populate applicationInsightsId.
	// +kubebuilder:validation:Optional
	ApplicationInsightsIDSelector *v1.Selector `json:"applicationInsightsIdSelector,omitempty" tf:"-"`

	// Purpose/user defined descriptive test for this WebTest.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Should the WebTest be enabled?
	// +kubebuilder:validation:Optional
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// Interval in seconds between test runs for this WebTest. Valid options are 300, 600 and 900. Defaults to 300.
	// +kubebuilder:validation:Optional
	Frequency *float64 `json:"frequency,omitempty" tf:"frequency,omitempty"`

	// Specifies a list of where to physically run the tests from to give global coverage for accessibility of your application.
	// +kubebuilder:validation:Optional
	GeoLocations []*string `json:"geoLocations,omitempty" tf:"geo_locations,omitempty"`

	// The Azure Region where the Application Insights Standard WebTest should exist. Changing this forces a new Application Insights Standard WebTest to be created. It needs to correlate with location of the parent resource (azurerm_application_insights)
	// +kubebuilder:validation:Optional
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// A request block as defined below.
	// +kubebuilder:validation:Optional
	Request []RequestParameters `json:"request,omitempty" tf:"request,omitempty"`

	// The name of the Resource Group where the Application Insights Standard WebTest should exist. Changing this forces a new Application Insights Standard WebTest to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/azure/v1beta1.ResourceGroup
	// +kubebuilder:validation:Optional
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// Reference to a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// Selector for a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`

	// Should the retry on WebTest failure be enabled?
	// +kubebuilder:validation:Optional
	RetryEnabled *bool `json:"retryEnabled,omitempty" tf:"retry_enabled,omitempty"`

	// A mapping of tags which should be assigned to the Application Insights Standard WebTest.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// Seconds until this WebTest will timeout and fail. Default is 30.
	// +kubebuilder:validation:Optional
	Timeout *float64 `json:"timeout,omitempty" tf:"timeout,omitempty"`

	// A validation_rules block as defined below.
	// +kubebuilder:validation:Optional
	ValidationRules []ValidationRulesParameters `json:"validationRules,omitempty" tf:"validation_rules,omitempty"`
}

type ContentInitParameters struct {

	// A string value containing the content to match on.
	ContentMatch *string `json:"contentMatch,omitempty" tf:"content_match,omitempty"`

	// Ignore the casing in the content_match value.
	IgnoreCase *bool `json:"ignoreCase,omitempty" tf:"ignore_case,omitempty"`

	// If the content of content_match is found, pass the test. If set to false, the WebTest is failing if the content of content_match is found.
	PassIfTextFound *bool `json:"passIfTextFound,omitempty" tf:"pass_if_text_found,omitempty"`
}

type ContentObservation struct {

	// A string value containing the content to match on.
	ContentMatch *string `json:"contentMatch,omitempty" tf:"content_match,omitempty"`

	// Ignore the casing in the content_match value.
	IgnoreCase *bool `json:"ignoreCase,omitempty" tf:"ignore_case,omitempty"`

	// If the content of content_match is found, pass the test. If set to false, the WebTest is failing if the content of content_match is found.
	PassIfTextFound *bool `json:"passIfTextFound,omitempty" tf:"pass_if_text_found,omitempty"`
}

type ContentParameters struct {

	// A string value containing the content to match on.
	// +kubebuilder:validation:Optional
	ContentMatch *string `json:"contentMatch" tf:"content_match,omitempty"`

	// Ignore the casing in the content_match value.
	// +kubebuilder:validation:Optional
	IgnoreCase *bool `json:"ignoreCase,omitempty" tf:"ignore_case,omitempty"`

	// If the content of content_match is found, pass the test. If set to false, the WebTest is failing if the content of content_match is found.
	// +kubebuilder:validation:Optional
	PassIfTextFound *bool `json:"passIfTextFound,omitempty" tf:"pass_if_text_found,omitempty"`
}

type HeaderInitParameters struct {

	// The name which should be used for this Application Insights Standard WebTest. Changing this forces a new Application Insights Standard WebTest to be created.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The value which should be used for a header in the request.
	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type HeaderObservation struct {

	// The name which should be used for this Application Insights Standard WebTest. Changing this forces a new Application Insights Standard WebTest to be created.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The value which should be used for a header in the request.
	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type HeaderParameters struct {

	// The name which should be used for this Application Insights Standard WebTest. Changing this forces a new Application Insights Standard WebTest to be created.
	// +kubebuilder:validation:Optional
	Name *string `json:"name" tf:"name,omitempty"`

	// The value which should be used for a header in the request.
	// +kubebuilder:validation:Optional
	Value *string `json:"value" tf:"value,omitempty"`
}

type RequestInitParameters struct {

	// The WebTest request body.
	Body *string `json:"body,omitempty" tf:"body,omitempty"`

	// Should the following of redirects be enabled? Defaults to true.
	FollowRedirectsEnabled *bool `json:"followRedirectsEnabled,omitempty" tf:"follow_redirects_enabled,omitempty"`

	// Which HTTP verb to use for the call. Options are 'GET', 'POST', 'PUT', 'PATCH', and 'DELETE'.
	HTTPVerb *string `json:"httpVerb,omitempty" tf:"http_verb,omitempty"`

	// One or more header blocks as defined above.
	Header []HeaderInitParameters `json:"header,omitempty" tf:"header,omitempty"`

	// Should the parsing of dependend requests be enabled? Defaults to true.
	ParseDependentRequestsEnabled *bool `json:"parseDependentRequestsEnabled,omitempty" tf:"parse_dependent_requests_enabled,omitempty"`

	// The WebTest request URL.
	URL *string `json:"url,omitempty" tf:"url,omitempty"`
}

type RequestObservation struct {

	// The WebTest request body.
	Body *string `json:"body,omitempty" tf:"body,omitempty"`

	// Should the following of redirects be enabled? Defaults to true.
	FollowRedirectsEnabled *bool `json:"followRedirectsEnabled,omitempty" tf:"follow_redirects_enabled,omitempty"`

	// Which HTTP verb to use for the call. Options are 'GET', 'POST', 'PUT', 'PATCH', and 'DELETE'.
	HTTPVerb *string `json:"httpVerb,omitempty" tf:"http_verb,omitempty"`

	// One or more header blocks as defined above.
	Header []HeaderObservation `json:"header,omitempty" tf:"header,omitempty"`

	// Should the parsing of dependend requests be enabled? Defaults to true.
	ParseDependentRequestsEnabled *bool `json:"parseDependentRequestsEnabled,omitempty" tf:"parse_dependent_requests_enabled,omitempty"`

	// The WebTest request URL.
	URL *string `json:"url,omitempty" tf:"url,omitempty"`
}

type RequestParameters struct {

	// The WebTest request body.
	// +kubebuilder:validation:Optional
	Body *string `json:"body,omitempty" tf:"body,omitempty"`

	// Should the following of redirects be enabled? Defaults to true.
	// +kubebuilder:validation:Optional
	FollowRedirectsEnabled *bool `json:"followRedirectsEnabled,omitempty" tf:"follow_redirects_enabled,omitempty"`

	// Which HTTP verb to use for the call. Options are 'GET', 'POST', 'PUT', 'PATCH', and 'DELETE'.
	// +kubebuilder:validation:Optional
	HTTPVerb *string `json:"httpVerb,omitempty" tf:"http_verb,omitempty"`

	// One or more header blocks as defined above.
	// +kubebuilder:validation:Optional
	Header []HeaderParameters `json:"header,omitempty" tf:"header,omitempty"`

	// Should the parsing of dependend requests be enabled? Defaults to true.
	// +kubebuilder:validation:Optional
	ParseDependentRequestsEnabled *bool `json:"parseDependentRequestsEnabled,omitempty" tf:"parse_dependent_requests_enabled,omitempty"`

	// The WebTest request URL.
	// +kubebuilder:validation:Optional
	URL *string `json:"url" tf:"url,omitempty"`
}

type ValidationRulesInitParameters struct {

	// A content block as defined above.
	Content []ContentInitParameters `json:"content,omitempty" tf:"content,omitempty"`

	// The expected status code of the response. Default is '200', '0' means 'response code < 400'
	ExpectedStatusCode *float64 `json:"expectedStatusCode,omitempty" tf:"expected_status_code,omitempty"`

	// The number of days of SSL certificate validity remaining for the checked endpoint. If the certificate has a shorter remaining lifetime left, the test will fail. This number should be between 1 and 365.
	SSLCertRemainingLifetime *float64 `json:"sslCertRemainingLifetime,omitempty" tf:"ssl_cert_remaining_lifetime,omitempty"`

	// Should the SSL check be enabled?
	SSLCheckEnabled *bool `json:"sslCheckEnabled,omitempty" tf:"ssl_check_enabled,omitempty"`
}

type ValidationRulesObservation struct {

	// A content block as defined above.
	Content []ContentObservation `json:"content,omitempty" tf:"content,omitempty"`

	// The expected status code of the response. Default is '200', '0' means 'response code < 400'
	ExpectedStatusCode *float64 `json:"expectedStatusCode,omitempty" tf:"expected_status_code,omitempty"`

	// The number of days of SSL certificate validity remaining for the checked endpoint. If the certificate has a shorter remaining lifetime left, the test will fail. This number should be between 1 and 365.
	SSLCertRemainingLifetime *float64 `json:"sslCertRemainingLifetime,omitempty" tf:"ssl_cert_remaining_lifetime,omitempty"`

	// Should the SSL check be enabled?
	SSLCheckEnabled *bool `json:"sslCheckEnabled,omitempty" tf:"ssl_check_enabled,omitempty"`
}

type ValidationRulesParameters struct {

	// A content block as defined above.
	// +kubebuilder:validation:Optional
	Content []ContentParameters `json:"content,omitempty" tf:"content,omitempty"`

	// The expected status code of the response. Default is '200', '0' means 'response code < 400'
	// +kubebuilder:validation:Optional
	ExpectedStatusCode *float64 `json:"expectedStatusCode,omitempty" tf:"expected_status_code,omitempty"`

	// The number of days of SSL certificate validity remaining for the checked endpoint. If the certificate has a shorter remaining lifetime left, the test will fail. This number should be between 1 and 365.
	// +kubebuilder:validation:Optional
	SSLCertRemainingLifetime *float64 `json:"sslCertRemainingLifetime,omitempty" tf:"ssl_cert_remaining_lifetime,omitempty"`

	// Should the SSL check be enabled?
	// +kubebuilder:validation:Optional
	SSLCheckEnabled *bool `json:"sslCheckEnabled,omitempty" tf:"ssl_check_enabled,omitempty"`
}

// ApplicationInsightsStandardWebTestSpec defines the desired state of ApplicationInsightsStandardWebTest
type ApplicationInsightsStandardWebTestSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ApplicationInsightsStandardWebTestParameters `json:"forProvider"`
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
	InitProvider ApplicationInsightsStandardWebTestInitParameters `json:"initProvider,omitempty"`
}

// ApplicationInsightsStandardWebTestStatus defines the observed state of ApplicationInsightsStandardWebTest.
type ApplicationInsightsStandardWebTestStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ApplicationInsightsStandardWebTestObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ApplicationInsightsStandardWebTest is the Schema for the ApplicationInsightsStandardWebTests API. Manages a Application Insights Standard WebTest.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type ApplicationInsightsStandardWebTest struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.geoLocations) || (has(self.initProvider) && has(self.initProvider.geoLocations))",message="spec.forProvider.geoLocations is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.location) || (has(self.initProvider) && has(self.initProvider.location))",message="spec.forProvider.location is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.request) || (has(self.initProvider) && has(self.initProvider.request))",message="spec.forProvider.request is a required parameter"
	Spec   ApplicationInsightsStandardWebTestSpec   `json:"spec"`
	Status ApplicationInsightsStandardWebTestStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ApplicationInsightsStandardWebTestList contains a list of ApplicationInsightsStandardWebTests
type ApplicationInsightsStandardWebTestList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ApplicationInsightsStandardWebTest `json:"items"`
}

// Repository type metadata.
var (
	ApplicationInsightsStandardWebTest_Kind             = "ApplicationInsightsStandardWebTest"
	ApplicationInsightsStandardWebTest_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ApplicationInsightsStandardWebTest_Kind}.String()
	ApplicationInsightsStandardWebTest_KindAPIVersion   = ApplicationInsightsStandardWebTest_Kind + "." + CRDGroupVersion.String()
	ApplicationInsightsStandardWebTest_GroupVersionKind = CRDGroupVersion.WithKind(ApplicationInsightsStandardWebTest_Kind)
)

func init() {
	SchemeBuilder.Register(&ApplicationInsightsStandardWebTest{}, &ApplicationInsightsStandardWebTestList{})
}
