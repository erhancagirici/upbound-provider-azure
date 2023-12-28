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

type CustomRulesInitParameters struct {

	// Type of action. Possible values are Allow, Block and Log.
	Action *string `json:"action,omitempty" tf:"action,omitempty"`

	// One or more match_conditions blocks as defined below.
	MatchConditions []MatchConditionsInitParameters `json:"matchConditions,omitempty" tf:"match_conditions,omitempty"`

	// Gets name of the resource that is unique within a policy. This name can be used to access the resource.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Describes priority of the rule. Rules with a lower value will be evaluated before rules with a higher value.
	Priority *float64 `json:"priority,omitempty" tf:"priority,omitempty"`

	// Describes the type of rule. Possible values are MatchRule and Invalid.
	RuleType *string `json:"ruleType,omitempty" tf:"rule_type,omitempty"`
}

type CustomRulesObservation struct {

	// Type of action. Possible values are Allow, Block and Log.
	Action *string `json:"action,omitempty" tf:"action,omitempty"`

	// One or more match_conditions blocks as defined below.
	MatchConditions []MatchConditionsObservation `json:"matchConditions,omitempty" tf:"match_conditions,omitempty"`

	// Gets name of the resource that is unique within a policy. This name can be used to access the resource.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Describes priority of the rule. Rules with a lower value will be evaluated before rules with a higher value.
	Priority *float64 `json:"priority,omitempty" tf:"priority,omitempty"`

	// Describes the type of rule. Possible values are MatchRule and Invalid.
	RuleType *string `json:"ruleType,omitempty" tf:"rule_type,omitempty"`
}

type CustomRulesParameters struct {

	// Type of action. Possible values are Allow, Block and Log.
	// +kubebuilder:validation:Optional
	Action *string `json:"action" tf:"action,omitempty"`

	// One or more match_conditions blocks as defined below.
	// +kubebuilder:validation:Optional
	MatchConditions []MatchConditionsParameters `json:"matchConditions" tf:"match_conditions,omitempty"`

	// Gets name of the resource that is unique within a policy. This name can be used to access the resource.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Describes priority of the rule. Rules with a lower value will be evaluated before rules with a higher value.
	// +kubebuilder:validation:Optional
	Priority *float64 `json:"priority" tf:"priority,omitempty"`

	// Describes the type of rule. Possible values are MatchRule and Invalid.
	// +kubebuilder:validation:Optional
	RuleType *string `json:"ruleType" tf:"rule_type,omitempty"`
}

type ExcludedRuleSetInitParameters struct {

	// One or more rule_group block defined below.
	RuleGroup []RuleGroupInitParameters `json:"ruleGroup,omitempty" tf:"rule_group,omitempty"`

	// The rule set type. Possible values: Microsoft_BotManagerRuleSet and OWASP.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`

	// The rule set version. Possible values: 0.1, 1.0, 2.2.9, 3.0, 3.1 and 3.2.
	Version *string `json:"version,omitempty" tf:"version,omitempty"`
}

type ExcludedRuleSetObservation struct {

	// One or more rule_group block defined below.
	RuleGroup []RuleGroupObservation `json:"ruleGroup,omitempty" tf:"rule_group,omitempty"`

	// The rule set type. Possible values: Microsoft_BotManagerRuleSet and OWASP.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`

	// The rule set version. Possible values: 0.1, 1.0, 2.2.9, 3.0, 3.1 and 3.2.
	Version *string `json:"version,omitempty" tf:"version,omitempty"`
}

type ExcludedRuleSetParameters struct {

	// One or more rule_group block defined below.
	// +kubebuilder:validation:Optional
	RuleGroup []RuleGroupParameters `json:"ruleGroup,omitempty" tf:"rule_group,omitempty"`

	// The rule set type. Possible values: Microsoft_BotManagerRuleSet and OWASP.
	// +kubebuilder:validation:Optional
	Type *string `json:"type,omitempty" tf:"type,omitempty"`

	// The rule set version. Possible values: 0.1, 1.0, 2.2.9, 3.0, 3.1 and 3.2.
	// +kubebuilder:validation:Optional
	Version *string `json:"version,omitempty" tf:"version,omitempty"`
}

type ManagedRuleSetInitParameters struct {

	// One or more rule_group_override block defined below.
	RuleGroupOverride []RuleGroupOverrideInitParameters `json:"ruleGroupOverride,omitempty" tf:"rule_group_override,omitempty"`

	// The rule set type. Possible values: Microsoft_BotManagerRuleSet and OWASP.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`

	// The rule set version. Possible values: 0.1, 1.0, 2.2.9, 3.0, 3.1 and 3.2.
	Version *string `json:"version,omitempty" tf:"version,omitempty"`
}

type ManagedRuleSetObservation struct {

	// One or more rule_group_override block defined below.
	RuleGroupOverride []RuleGroupOverrideObservation `json:"ruleGroupOverride,omitempty" tf:"rule_group_override,omitempty"`

	// The rule set type. Possible values: Microsoft_BotManagerRuleSet and OWASP.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`

	// The rule set version. Possible values: 0.1, 1.0, 2.2.9, 3.0, 3.1 and 3.2.
	Version *string `json:"version,omitempty" tf:"version,omitempty"`
}

type ManagedRuleSetParameters struct {

	// One or more rule_group_override block defined below.
	// +kubebuilder:validation:Optional
	RuleGroupOverride []RuleGroupOverrideParameters `json:"ruleGroupOverride,omitempty" tf:"rule_group_override,omitempty"`

	// The rule set type. Possible values: Microsoft_BotManagerRuleSet and OWASP.
	// +kubebuilder:validation:Optional
	Type *string `json:"type,omitempty" tf:"type,omitempty"`

	// The rule set version. Possible values: 0.1, 1.0, 2.2.9, 3.0, 3.1 and 3.2.
	// +kubebuilder:validation:Optional
	Version *string `json:"version" tf:"version,omitempty"`
}

type ManagedRulesExclusionInitParameters struct {

	// One or more excluded_rule_set block defined below.
	ExcludedRuleSet []ExcludedRuleSetInitParameters `json:"excludedRuleSet,omitempty" tf:"excluded_rule_set,omitempty"`

	// The name of the Match Variable. Possible values: RequestArgKeys, RequestArgNames, RequestArgValues, RequestCookieKeys, RequestCookieNames, RequestCookieValues, RequestHeaderKeys, RequestHeaderNames, RequestHeaderValues.
	MatchVariable *string `json:"matchVariable,omitempty" tf:"match_variable,omitempty"`

	// Describes field of the matchVariable collection
	Selector *string `json:"selector,omitempty" tf:"selector,omitempty"`

	// Describes operator to be matched. Possible values: Contains, EndsWith, Equals, EqualsAny, StartsWith.
	SelectorMatchOperator *string `json:"selectorMatchOperator,omitempty" tf:"selector_match_operator,omitempty"`
}

type ManagedRulesExclusionObservation struct {

	// One or more excluded_rule_set block defined below.
	ExcludedRuleSet []ExcludedRuleSetObservation `json:"excludedRuleSet,omitempty" tf:"excluded_rule_set,omitempty"`

	// The name of the Match Variable. Possible values: RequestArgKeys, RequestArgNames, RequestArgValues, RequestCookieKeys, RequestCookieNames, RequestCookieValues, RequestHeaderKeys, RequestHeaderNames, RequestHeaderValues.
	MatchVariable *string `json:"matchVariable,omitempty" tf:"match_variable,omitempty"`

	// Describes field of the matchVariable collection
	Selector *string `json:"selector,omitempty" tf:"selector,omitempty"`

	// Describes operator to be matched. Possible values: Contains, EndsWith, Equals, EqualsAny, StartsWith.
	SelectorMatchOperator *string `json:"selectorMatchOperator,omitempty" tf:"selector_match_operator,omitempty"`
}

type ManagedRulesExclusionParameters struct {

	// One or more excluded_rule_set block defined below.
	// +kubebuilder:validation:Optional
	ExcludedRuleSet []ExcludedRuleSetParameters `json:"excludedRuleSet,omitempty" tf:"excluded_rule_set,omitempty"`

	// The name of the Match Variable. Possible values: RequestArgKeys, RequestArgNames, RequestArgValues, RequestCookieKeys, RequestCookieNames, RequestCookieValues, RequestHeaderKeys, RequestHeaderNames, RequestHeaderValues.
	// +kubebuilder:validation:Optional
	MatchVariable *string `json:"matchVariable" tf:"match_variable,omitempty"`

	// Describes field of the matchVariable collection
	// +kubebuilder:validation:Optional
	Selector *string `json:"selector" tf:"selector,omitempty"`

	// Describes operator to be matched. Possible values: Contains, EndsWith, Equals, EqualsAny, StartsWith.
	// +kubebuilder:validation:Optional
	SelectorMatchOperator *string `json:"selectorMatchOperator" tf:"selector_match_operator,omitempty"`
}

type ManagedRulesInitParameters struct {

	// One or more exclusion block defined below.
	Exclusion []ManagedRulesExclusionInitParameters `json:"exclusion,omitempty" tf:"exclusion,omitempty"`

	// One or more managed_rule_set block defined below.
	ManagedRuleSet []ManagedRuleSetInitParameters `json:"managedRuleSet,omitempty" tf:"managed_rule_set,omitempty"`
}

type ManagedRulesObservation struct {

	// One or more exclusion block defined below.
	Exclusion []ManagedRulesExclusionObservation `json:"exclusion,omitempty" tf:"exclusion,omitempty"`

	// One or more managed_rule_set block defined below.
	ManagedRuleSet []ManagedRuleSetObservation `json:"managedRuleSet,omitempty" tf:"managed_rule_set,omitempty"`
}

type ManagedRulesParameters struct {

	// One or more exclusion block defined below.
	// +kubebuilder:validation:Optional
	Exclusion []ManagedRulesExclusionParameters `json:"exclusion,omitempty" tf:"exclusion,omitempty"`

	// One or more managed_rule_set block defined below.
	// +kubebuilder:validation:Optional
	ManagedRuleSet []ManagedRuleSetParameters `json:"managedRuleSet" tf:"managed_rule_set,omitempty"`
}

type MatchConditionsInitParameters struct {

	// A list of match values. This is Required when the operator is not Any.
	MatchValues []*string `json:"matchValues,omitempty" tf:"match_values,omitempty"`

	// One or more match_variables blocks as defined below.
	MatchVariables []MatchVariablesInitParameters `json:"matchVariables,omitempty" tf:"match_variables,omitempty"`

	// Describes if this is negate condition or not
	NegationCondition *bool `json:"negationCondition,omitempty" tf:"negation_condition,omitempty"`

	// Describes operator to be matched. Possible values are Any, IPMatch, GeoMatch, Equal, Contains, LessThan, GreaterThan, LessThanOrEqual, GreaterThanOrEqual, BeginsWith, EndsWith and Regex.
	Operator *string `json:"operator,omitempty" tf:"operator,omitempty"`

	// A list of transformations to do before the match is attempted. Possible values are HtmlEntityDecode, Lowercase, RemoveNulls, Trim, UrlDecode and UrlEncode.
	// +listType=set
	Transforms []*string `json:"transforms,omitempty" tf:"transforms,omitempty"`
}

type MatchConditionsObservation struct {

	// A list of match values. This is Required when the operator is not Any.
	MatchValues []*string `json:"matchValues,omitempty" tf:"match_values,omitempty"`

	// One or more match_variables blocks as defined below.
	MatchVariables []MatchVariablesObservation `json:"matchVariables,omitempty" tf:"match_variables,omitempty"`

	// Describes if this is negate condition or not
	NegationCondition *bool `json:"negationCondition,omitempty" tf:"negation_condition,omitempty"`

	// Describes operator to be matched. Possible values are Any, IPMatch, GeoMatch, Equal, Contains, LessThan, GreaterThan, LessThanOrEqual, GreaterThanOrEqual, BeginsWith, EndsWith and Regex.
	Operator *string `json:"operator,omitempty" tf:"operator,omitempty"`

	// A list of transformations to do before the match is attempted. Possible values are HtmlEntityDecode, Lowercase, RemoveNulls, Trim, UrlDecode and UrlEncode.
	// +listType=set
	Transforms []*string `json:"transforms,omitempty" tf:"transforms,omitempty"`
}

type MatchConditionsParameters struct {

	// A list of match values. This is Required when the operator is not Any.
	// +kubebuilder:validation:Optional
	MatchValues []*string `json:"matchValues,omitempty" tf:"match_values,omitempty"`

	// One or more match_variables blocks as defined below.
	// +kubebuilder:validation:Optional
	MatchVariables []MatchVariablesParameters `json:"matchVariables" tf:"match_variables,omitempty"`

	// Describes if this is negate condition or not
	// +kubebuilder:validation:Optional
	NegationCondition *bool `json:"negationCondition,omitempty" tf:"negation_condition,omitempty"`

	// Describes operator to be matched. Possible values are Any, IPMatch, GeoMatch, Equal, Contains, LessThan, GreaterThan, LessThanOrEqual, GreaterThanOrEqual, BeginsWith, EndsWith and Regex.
	// +kubebuilder:validation:Optional
	Operator *string `json:"operator" tf:"operator,omitempty"`

	// A list of transformations to do before the match is attempted. Possible values are HtmlEntityDecode, Lowercase, RemoveNulls, Trim, UrlDecode and UrlEncode.
	// +kubebuilder:validation:Optional
	// +listType=set
	Transforms []*string `json:"transforms,omitempty" tf:"transforms,omitempty"`
}

type MatchVariablesInitParameters struct {

	// Describes field of the matchVariable collection
	Selector *string `json:"selector,omitempty" tf:"selector,omitempty"`

	// The name of the Match Variable. Possible values are RemoteAddr, RequestMethod, QueryString, PostArgs, RequestUri, RequestHeaders, RequestBody and RequestCookies.
	VariableName *string `json:"variableName,omitempty" tf:"variable_name,omitempty"`
}

type MatchVariablesObservation struct {

	// Describes field of the matchVariable collection
	Selector *string `json:"selector,omitempty" tf:"selector,omitempty"`

	// The name of the Match Variable. Possible values are RemoteAddr, RequestMethod, QueryString, PostArgs, RequestUri, RequestHeaders, RequestBody and RequestCookies.
	VariableName *string `json:"variableName,omitempty" tf:"variable_name,omitempty"`
}

type MatchVariablesParameters struct {

	// Describes field of the matchVariable collection
	// +kubebuilder:validation:Optional
	Selector *string `json:"selector,omitempty" tf:"selector,omitempty"`

	// The name of the Match Variable. Possible values are RemoteAddr, RequestMethod, QueryString, PostArgs, RequestUri, RequestHeaders, RequestBody and RequestCookies.
	// +kubebuilder:validation:Optional
	VariableName *string `json:"variableName" tf:"variable_name,omitempty"`
}

type PolicySettingsInitParameters struct {

	// Describes if the policy is in enabled state or disabled state. Defaults to true.
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// The File Upload Limit in MB. Accepted values are in the range 1 to 4000. Defaults to 100.
	FileUploadLimitInMb *float64 `json:"fileUploadLimitInMb,omitempty" tf:"file_upload_limit_in_mb,omitempty"`

	// The Maximum Request Body Size in KB. Accepted values are in the range 8 to 2000. Defaults to 128.
	MaxRequestBodySizeInKb *float64 `json:"maxRequestBodySizeInKb,omitempty" tf:"max_request_body_size_in_kb,omitempty"`

	// Describes if it is in detection mode or prevention mode at the policy level. Valid values are Detection and Prevention. Defaults to Prevention.
	Mode *string `json:"mode,omitempty" tf:"mode,omitempty"`

	// Is Request Body Inspection enabled? Defaults to true.
	RequestBodyCheck *bool `json:"requestBodyCheck,omitempty" tf:"request_body_check,omitempty"`
}

type PolicySettingsObservation struct {

	// Describes if the policy is in enabled state or disabled state. Defaults to true.
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// The File Upload Limit in MB. Accepted values are in the range 1 to 4000. Defaults to 100.
	FileUploadLimitInMb *float64 `json:"fileUploadLimitInMb,omitempty" tf:"file_upload_limit_in_mb,omitempty"`

	// The Maximum Request Body Size in KB. Accepted values are in the range 8 to 2000. Defaults to 128.
	MaxRequestBodySizeInKb *float64 `json:"maxRequestBodySizeInKb,omitempty" tf:"max_request_body_size_in_kb,omitempty"`

	// Describes if it is in detection mode or prevention mode at the policy level. Valid values are Detection and Prevention. Defaults to Prevention.
	Mode *string `json:"mode,omitempty" tf:"mode,omitempty"`

	// Is Request Body Inspection enabled? Defaults to true.
	RequestBodyCheck *bool `json:"requestBodyCheck,omitempty" tf:"request_body_check,omitempty"`
}

type PolicySettingsParameters struct {

	// Describes if the policy is in enabled state or disabled state. Defaults to true.
	// +kubebuilder:validation:Optional
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// The File Upload Limit in MB. Accepted values are in the range 1 to 4000. Defaults to 100.
	// +kubebuilder:validation:Optional
	FileUploadLimitInMb *float64 `json:"fileUploadLimitInMb,omitempty" tf:"file_upload_limit_in_mb,omitempty"`

	// The Maximum Request Body Size in KB. Accepted values are in the range 8 to 2000. Defaults to 128.
	// +kubebuilder:validation:Optional
	MaxRequestBodySizeInKb *float64 `json:"maxRequestBodySizeInKb,omitempty" tf:"max_request_body_size_in_kb,omitempty"`

	// Describes if it is in detection mode or prevention mode at the policy level. Valid values are Detection and Prevention. Defaults to Prevention.
	// +kubebuilder:validation:Optional
	Mode *string `json:"mode,omitempty" tf:"mode,omitempty"`

	// Is Request Body Inspection enabled? Defaults to true.
	// +kubebuilder:validation:Optional
	RequestBodyCheck *bool `json:"requestBodyCheck,omitempty" tf:"request_body_check,omitempty"`
}

type RuleGroupInitParameters struct {

	// One or more Rule IDs for exclusion.
	ExcludedRules []*string `json:"excludedRules,omitempty" tf:"excluded_rules,omitempty"`

	// The name of the Rule Group. Possible values are BadBots, crs_20_protocol_violations, crs_21_protocol_anomalies, crs_23_request_limits, crs_30_http_policy, crs_35_bad_robots, crs_40_generic_attacks, crs_41_sql_injection_attacks, crs_41_xss_attacks, crs_42_tight_security, crs_45_trojans, General, GoodBots, Known-CVEs, REQUEST-911-METHOD-ENFORCEMENT, REQUEST-913-SCANNER-DETECTION, REQUEST-920-PROTOCOL-ENFORCEMENT, REQUEST-921-PROTOCOL-ATTACK, REQUEST-930-APPLICATION-ATTACK-LFI, REQUEST-931-APPLICATION-ATTACK-RFI, REQUEST-932-APPLICATION-ATTACK-RCE, REQUEST-933-APPLICATION-ATTACK-PHP, REQUEST-941-APPLICATION-ATTACK-XSS, REQUEST-942-APPLICATION-ATTACK-SQLI, REQUEST-943-APPLICATION-ATTACK-SESSION-FIXATION, REQUEST-944-APPLICATION-ATTACK-JAVA and UnknownBots.
	RuleGroupName *string `json:"ruleGroupName,omitempty" tf:"rule_group_name,omitempty"`
}

type RuleGroupObservation struct {

	// One or more Rule IDs for exclusion.
	ExcludedRules []*string `json:"excludedRules,omitempty" tf:"excluded_rules,omitempty"`

	// The name of the Rule Group. Possible values are BadBots, crs_20_protocol_violations, crs_21_protocol_anomalies, crs_23_request_limits, crs_30_http_policy, crs_35_bad_robots, crs_40_generic_attacks, crs_41_sql_injection_attacks, crs_41_xss_attacks, crs_42_tight_security, crs_45_trojans, General, GoodBots, Known-CVEs, REQUEST-911-METHOD-ENFORCEMENT, REQUEST-913-SCANNER-DETECTION, REQUEST-920-PROTOCOL-ENFORCEMENT, REQUEST-921-PROTOCOL-ATTACK, REQUEST-930-APPLICATION-ATTACK-LFI, REQUEST-931-APPLICATION-ATTACK-RFI, REQUEST-932-APPLICATION-ATTACK-RCE, REQUEST-933-APPLICATION-ATTACK-PHP, REQUEST-941-APPLICATION-ATTACK-XSS, REQUEST-942-APPLICATION-ATTACK-SQLI, REQUEST-943-APPLICATION-ATTACK-SESSION-FIXATION, REQUEST-944-APPLICATION-ATTACK-JAVA and UnknownBots.
	RuleGroupName *string `json:"ruleGroupName,omitempty" tf:"rule_group_name,omitempty"`
}

type RuleGroupOverrideInitParameters struct {
	DisabledRules []*string `json:"disabledRules,omitempty" tf:"disabled_rules,omitempty"`

	// One or more rule block defined below.
	Rule []RuleGroupOverrideRuleInitParameters `json:"rule,omitempty" tf:"rule,omitempty"`

	// The name of the Rule Group. Possible values are BadBots, crs_20_protocol_violations, crs_21_protocol_anomalies, crs_23_request_limits, crs_30_http_policy, crs_35_bad_robots, crs_40_generic_attacks, crs_41_sql_injection_attacks, crs_41_xss_attacks, crs_42_tight_security, crs_45_trojans, General, GoodBots, Known-CVEs, REQUEST-911-METHOD-ENFORCEMENT, REQUEST-913-SCANNER-DETECTION, REQUEST-920-PROTOCOL-ENFORCEMENT, REQUEST-921-PROTOCOL-ATTACK, REQUEST-930-APPLICATION-ATTACK-LFI, REQUEST-931-APPLICATION-ATTACK-RFI, REQUEST-932-APPLICATION-ATTACK-RCE, REQUEST-933-APPLICATION-ATTACK-PHP, REQUEST-941-APPLICATION-ATTACK-XSS, REQUEST-942-APPLICATION-ATTACK-SQLI, REQUEST-943-APPLICATION-ATTACK-SESSION-FIXATION, REQUEST-944-APPLICATION-ATTACK-JAVA and UnknownBots.
	RuleGroupName *string `json:"ruleGroupName,omitempty" tf:"rule_group_name,omitempty"`
}

type RuleGroupOverrideObservation struct {
	DisabledRules []*string `json:"disabledRules,omitempty" tf:"disabled_rules,omitempty"`

	// One or more rule block defined below.
	Rule []RuleGroupOverrideRuleObservation `json:"rule,omitempty" tf:"rule,omitempty"`

	// The name of the Rule Group. Possible values are BadBots, crs_20_protocol_violations, crs_21_protocol_anomalies, crs_23_request_limits, crs_30_http_policy, crs_35_bad_robots, crs_40_generic_attacks, crs_41_sql_injection_attacks, crs_41_xss_attacks, crs_42_tight_security, crs_45_trojans, General, GoodBots, Known-CVEs, REQUEST-911-METHOD-ENFORCEMENT, REQUEST-913-SCANNER-DETECTION, REQUEST-920-PROTOCOL-ENFORCEMENT, REQUEST-921-PROTOCOL-ATTACK, REQUEST-930-APPLICATION-ATTACK-LFI, REQUEST-931-APPLICATION-ATTACK-RFI, REQUEST-932-APPLICATION-ATTACK-RCE, REQUEST-933-APPLICATION-ATTACK-PHP, REQUEST-941-APPLICATION-ATTACK-XSS, REQUEST-942-APPLICATION-ATTACK-SQLI, REQUEST-943-APPLICATION-ATTACK-SESSION-FIXATION, REQUEST-944-APPLICATION-ATTACK-JAVA and UnknownBots.
	RuleGroupName *string `json:"ruleGroupName,omitempty" tf:"rule_group_name,omitempty"`
}

type RuleGroupOverrideParameters struct {

	// +kubebuilder:validation:Optional
	DisabledRules []*string `json:"disabledRules,omitempty" tf:"disabled_rules,omitempty"`

	// One or more rule block defined below.
	// +kubebuilder:validation:Optional
	Rule []RuleGroupOverrideRuleParameters `json:"rule,omitempty" tf:"rule,omitempty"`

	// The name of the Rule Group. Possible values are BadBots, crs_20_protocol_violations, crs_21_protocol_anomalies, crs_23_request_limits, crs_30_http_policy, crs_35_bad_robots, crs_40_generic_attacks, crs_41_sql_injection_attacks, crs_41_xss_attacks, crs_42_tight_security, crs_45_trojans, General, GoodBots, Known-CVEs, REQUEST-911-METHOD-ENFORCEMENT, REQUEST-913-SCANNER-DETECTION, REQUEST-920-PROTOCOL-ENFORCEMENT, REQUEST-921-PROTOCOL-ATTACK, REQUEST-930-APPLICATION-ATTACK-LFI, REQUEST-931-APPLICATION-ATTACK-RFI, REQUEST-932-APPLICATION-ATTACK-RCE, REQUEST-933-APPLICATION-ATTACK-PHP, REQUEST-941-APPLICATION-ATTACK-XSS, REQUEST-942-APPLICATION-ATTACK-SQLI, REQUEST-943-APPLICATION-ATTACK-SESSION-FIXATION, REQUEST-944-APPLICATION-ATTACK-JAVA and UnknownBots.
	// +kubebuilder:validation:Optional
	RuleGroupName *string `json:"ruleGroupName" tf:"rule_group_name,omitempty"`
}

type RuleGroupOverrideRuleInitParameters struct {

	// Describes the override action to be applied when rule matches. Possible values are Allow, AnomalyScoring, Block and Log.
	Action *string `json:"action,omitempty" tf:"action,omitempty"`

	// Describes if the managed rule is in enabled state or disabled state.
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// Identifier for the managed rule.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type RuleGroupOverrideRuleObservation struct {

	// Describes the override action to be applied when rule matches. Possible values are Allow, AnomalyScoring, Block and Log.
	Action *string `json:"action,omitempty" tf:"action,omitempty"`

	// Describes if the managed rule is in enabled state or disabled state.
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// Identifier for the managed rule.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type RuleGroupOverrideRuleParameters struct {

	// Describes the override action to be applied when rule matches. Possible values are Allow, AnomalyScoring, Block and Log.
	// +kubebuilder:validation:Optional
	Action *string `json:"action,omitempty" tf:"action,omitempty"`

	// Describes if the managed rule is in enabled state or disabled state.
	// +kubebuilder:validation:Optional
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// Identifier for the managed rule.
	// +kubebuilder:validation:Optional
	ID *string `json:"id" tf:"id,omitempty"`
}

type RuleGroupParameters struct {

	// One or more Rule IDs for exclusion.
	// +kubebuilder:validation:Optional
	ExcludedRules []*string `json:"excludedRules,omitempty" tf:"excluded_rules,omitempty"`

	// The name of the Rule Group. Possible values are BadBots, crs_20_protocol_violations, crs_21_protocol_anomalies, crs_23_request_limits, crs_30_http_policy, crs_35_bad_robots, crs_40_generic_attacks, crs_41_sql_injection_attacks, crs_41_xss_attacks, crs_42_tight_security, crs_45_trojans, General, GoodBots, Known-CVEs, REQUEST-911-METHOD-ENFORCEMENT, REQUEST-913-SCANNER-DETECTION, REQUEST-920-PROTOCOL-ENFORCEMENT, REQUEST-921-PROTOCOL-ATTACK, REQUEST-930-APPLICATION-ATTACK-LFI, REQUEST-931-APPLICATION-ATTACK-RFI, REQUEST-932-APPLICATION-ATTACK-RCE, REQUEST-933-APPLICATION-ATTACK-PHP, REQUEST-941-APPLICATION-ATTACK-XSS, REQUEST-942-APPLICATION-ATTACK-SQLI, REQUEST-943-APPLICATION-ATTACK-SESSION-FIXATION, REQUEST-944-APPLICATION-ATTACK-JAVA and UnknownBots.
	// +kubebuilder:validation:Optional
	RuleGroupName *string `json:"ruleGroupName" tf:"rule_group_name,omitempty"`
}

type WebApplicationFirewallPolicyInitParameters struct {

	// One or more custom_rules blocks as defined below.
	CustomRules []CustomRulesInitParameters `json:"customRules,omitempty" tf:"custom_rules,omitempty"`

	// Resource location. Changing this forces a new resource to be created.
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// A managed_rules blocks as defined below.
	ManagedRules []ManagedRulesInitParameters `json:"managedRules,omitempty" tf:"managed_rules,omitempty"`

	// A policy_settings block as defined below.
	PolicySettings []PolicySettingsInitParameters `json:"policySettings,omitempty" tf:"policy_settings,omitempty"`

	// A mapping of tags to assign to the Web Application Firewall Policy.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type WebApplicationFirewallPolicyObservation struct {

	// One or more custom_rules blocks as defined below.
	CustomRules []CustomRulesObservation `json:"customRules,omitempty" tf:"custom_rules,omitempty"`

	// A list of HTTP Listener IDs from an azurerm_application_gateway.
	HTTPListenerIds []*string `json:"httpListenerIds,omitempty" tf:"http_listener_ids,omitempty"`

	// The ID of the Web Application Firewall Policy.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Resource location. Changing this forces a new resource to be created.
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// A managed_rules blocks as defined below.
	ManagedRules []ManagedRulesObservation `json:"managedRules,omitempty" tf:"managed_rules,omitempty"`

	// A list of URL Path Map Path Rule IDs from an azurerm_application_gateway.
	PathBasedRuleIds []*string `json:"pathBasedRuleIds,omitempty" tf:"path_based_rule_ids,omitempty"`

	// A policy_settings block as defined below.
	PolicySettings []PolicySettingsObservation `json:"policySettings,omitempty" tf:"policy_settings,omitempty"`

	// The name of the resource group. Changing this forces a new resource to be created.
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// A mapping of tags to assign to the Web Application Firewall Policy.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type WebApplicationFirewallPolicyParameters struct {

	// One or more custom_rules blocks as defined below.
	// +kubebuilder:validation:Optional
	CustomRules []CustomRulesParameters `json:"customRules,omitempty" tf:"custom_rules,omitempty"`

	// Resource location. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// A managed_rules blocks as defined below.
	// +kubebuilder:validation:Optional
	ManagedRules []ManagedRulesParameters `json:"managedRules,omitempty" tf:"managed_rules,omitempty"`

	// A policy_settings block as defined below.
	// +kubebuilder:validation:Optional
	PolicySettings []PolicySettingsParameters `json:"policySettings,omitempty" tf:"policy_settings,omitempty"`

	// The name of the resource group. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/azure/v1beta1.ResourceGroup
	// +kubebuilder:validation:Optional
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// Reference to a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// Selector for a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`

	// A mapping of tags to assign to the Web Application Firewall Policy.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

// WebApplicationFirewallPolicySpec defines the desired state of WebApplicationFirewallPolicy
type WebApplicationFirewallPolicySpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     WebApplicationFirewallPolicyParameters `json:"forProvider"`
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
	InitProvider WebApplicationFirewallPolicyInitParameters `json:"initProvider,omitempty"`
}

// WebApplicationFirewallPolicyStatus defines the observed state of WebApplicationFirewallPolicy.
type WebApplicationFirewallPolicyStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        WebApplicationFirewallPolicyObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// WebApplicationFirewallPolicy is the Schema for the WebApplicationFirewallPolicys API. Manages a Azure Web Application Firewall Policy instance.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type WebApplicationFirewallPolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.location) || (has(self.initProvider) && has(self.initProvider.location))",message="spec.forProvider.location is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.managedRules) || (has(self.initProvider) && has(self.initProvider.managedRules))",message="spec.forProvider.managedRules is a required parameter"
	Spec   WebApplicationFirewallPolicySpec   `json:"spec"`
	Status WebApplicationFirewallPolicyStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// WebApplicationFirewallPolicyList contains a list of WebApplicationFirewallPolicys
type WebApplicationFirewallPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []WebApplicationFirewallPolicy `json:"items"`
}

// Repository type metadata.
var (
	WebApplicationFirewallPolicy_Kind             = "WebApplicationFirewallPolicy"
	WebApplicationFirewallPolicy_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: WebApplicationFirewallPolicy_Kind}.String()
	WebApplicationFirewallPolicy_KindAPIVersion   = WebApplicationFirewallPolicy_Kind + "." + CRDGroupVersion.String()
	WebApplicationFirewallPolicy_GroupVersionKind = CRDGroupVersion.WithKind(WebApplicationFirewallPolicy_Kind)
)

func init() {
	SchemeBuilder.Register(&WebApplicationFirewallPolicy{}, &WebApplicationFirewallPolicyList{})
}
