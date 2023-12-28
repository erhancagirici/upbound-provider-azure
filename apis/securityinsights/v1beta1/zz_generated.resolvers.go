/*
Copyright 2022 Upbound Inc.
*/
// Code generated by angryjet. DO NOT EDIT.

package v1beta1

import (
	"context"
	reference "github.com/crossplane/crossplane-runtime/pkg/reference"
	resource "github.com/crossplane/upjet/pkg/resource"
	errors "github.com/pkg/errors"
	v1beta12 "github.com/upbound/provider-azure/apis/azure/v1beta1"
	v1beta11 "github.com/upbound/provider-azure/apis/operationalinsights/v1beta1"
	v1beta1 "github.com/upbound/provider-azure/apis/operationsmanagement/v1beta1"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

// ResolveReferences of this SentinelAlertRuleFusion.
func (mg *SentinelAlertRuleFusion) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.LogAnalyticsWorkspaceID),
		Extract:      resource.ExtractParamPath("workspace_resource_id", false),
		Reference:    mg.Spec.ForProvider.LogAnalyticsWorkspaceIDRef,
		Selector:     mg.Spec.ForProvider.LogAnalyticsWorkspaceIDSelector,
		To: reference.To{
			List:    &v1beta1.LogAnalyticsSolutionList{},
			Managed: &v1beta1.LogAnalyticsSolution{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.LogAnalyticsWorkspaceID")
	}
	mg.Spec.ForProvider.LogAnalyticsWorkspaceID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.LogAnalyticsWorkspaceIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.LogAnalyticsWorkspaceID),
		Extract:      resource.ExtractParamPath("workspace_resource_id", false),
		Reference:    mg.Spec.InitProvider.LogAnalyticsWorkspaceIDRef,
		Selector:     mg.Spec.InitProvider.LogAnalyticsWorkspaceIDSelector,
		To: reference.To{
			List:    &v1beta1.LogAnalyticsSolutionList{},
			Managed: &v1beta1.LogAnalyticsSolution{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.LogAnalyticsWorkspaceID")
	}
	mg.Spec.InitProvider.LogAnalyticsWorkspaceID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.LogAnalyticsWorkspaceIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this SentinelAlertRuleMSSecurityIncident.
func (mg *SentinelAlertRuleMSSecurityIncident) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.LogAnalyticsWorkspaceID),
		Extract:      resource.ExtractParamPath("workspace_id", false),
		Reference:    mg.Spec.ForProvider.LogAnalyticsWorkspaceIDRef,
		Selector:     mg.Spec.ForProvider.LogAnalyticsWorkspaceIDSelector,
		To: reference.To{
			List:    &SentinelLogAnalyticsWorkspaceOnboardingList{},
			Managed: &SentinelLogAnalyticsWorkspaceOnboarding{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.LogAnalyticsWorkspaceID")
	}
	mg.Spec.ForProvider.LogAnalyticsWorkspaceID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.LogAnalyticsWorkspaceIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this SentinelAlertRuleMachineLearningBehaviorAnalytics.
func (mg *SentinelAlertRuleMachineLearningBehaviorAnalytics) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.LogAnalyticsWorkspaceID),
		Extract:      resource.ExtractResourceID(),
		Reference:    mg.Spec.ForProvider.LogAnalyticsWorkspaceIDRef,
		Selector:     mg.Spec.ForProvider.LogAnalyticsWorkspaceIDSelector,
		To: reference.To{
			List:    &v1beta11.WorkspaceList{},
			Managed: &v1beta11.Workspace{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.LogAnalyticsWorkspaceID")
	}
	mg.Spec.ForProvider.LogAnalyticsWorkspaceID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.LogAnalyticsWorkspaceIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.LogAnalyticsWorkspaceID),
		Extract:      resource.ExtractResourceID(),
		Reference:    mg.Spec.InitProvider.LogAnalyticsWorkspaceIDRef,
		Selector:     mg.Spec.InitProvider.LogAnalyticsWorkspaceIDSelector,
		To: reference.To{
			List:    &v1beta11.WorkspaceList{},
			Managed: &v1beta11.Workspace{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.LogAnalyticsWorkspaceID")
	}
	mg.Spec.InitProvider.LogAnalyticsWorkspaceID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.LogAnalyticsWorkspaceIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this SentinelAutomationRule.
func (mg *SentinelAutomationRule) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.LogAnalyticsWorkspaceID),
		Extract:      resource.ExtractParamPath("workspace_id", false),
		Reference:    mg.Spec.ForProvider.LogAnalyticsWorkspaceIDRef,
		Selector:     mg.Spec.ForProvider.LogAnalyticsWorkspaceIDSelector,
		To: reference.To{
			List:    &SentinelLogAnalyticsWorkspaceOnboardingList{},
			Managed: &SentinelLogAnalyticsWorkspaceOnboarding{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.LogAnalyticsWorkspaceID")
	}
	mg.Spec.ForProvider.LogAnalyticsWorkspaceID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.LogAnalyticsWorkspaceIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.LogAnalyticsWorkspaceID),
		Extract:      resource.ExtractParamPath("workspace_id", false),
		Reference:    mg.Spec.InitProvider.LogAnalyticsWorkspaceIDRef,
		Selector:     mg.Spec.InitProvider.LogAnalyticsWorkspaceIDSelector,
		To: reference.To{
			List:    &SentinelLogAnalyticsWorkspaceOnboardingList{},
			Managed: &SentinelLogAnalyticsWorkspaceOnboarding{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.LogAnalyticsWorkspaceID")
	}
	mg.Spec.InitProvider.LogAnalyticsWorkspaceID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.LogAnalyticsWorkspaceIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this SentinelDataConnectorIOT.
func (mg *SentinelDataConnectorIOT) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.LogAnalyticsWorkspaceID),
		Extract:      resource.ExtractParamPath("workspace_id", false),
		Reference:    mg.Spec.ForProvider.LogAnalyticsWorkspaceIDRef,
		Selector:     mg.Spec.ForProvider.LogAnalyticsWorkspaceIDSelector,
		To: reference.To{
			List:    &SentinelLogAnalyticsWorkspaceOnboardingList{},
			Managed: &SentinelLogAnalyticsWorkspaceOnboarding{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.LogAnalyticsWorkspaceID")
	}
	mg.Spec.ForProvider.LogAnalyticsWorkspaceID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.LogAnalyticsWorkspaceIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this SentinelLogAnalyticsWorkspaceOnboarding.
func (mg *SentinelLogAnalyticsWorkspaceOnboarding) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ResourceGroupName),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.ResourceGroupNameRef,
		Selector:     mg.Spec.ForProvider.ResourceGroupNameSelector,
		To: reference.To{
			List:    &v1beta12.ResourceGroupList{},
			Managed: &v1beta12.ResourceGroup{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ResourceGroupName")
	}
	mg.Spec.ForProvider.ResourceGroupName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ResourceGroupNameRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.WorkspaceName),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.WorkspaceNameRef,
		Selector:     mg.Spec.ForProvider.WorkspaceNameSelector,
		To: reference.To{
			List:    &v1beta11.WorkspaceList{},
			Managed: &v1beta11.Workspace{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.WorkspaceName")
	}
	mg.Spec.ForProvider.WorkspaceName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.WorkspaceNameRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.ResourceGroupName),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.InitProvider.ResourceGroupNameRef,
		Selector:     mg.Spec.InitProvider.ResourceGroupNameSelector,
		To: reference.To{
			List:    &v1beta12.ResourceGroupList{},
			Managed: &v1beta12.ResourceGroup{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.ResourceGroupName")
	}
	mg.Spec.InitProvider.ResourceGroupName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.ResourceGroupNameRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.WorkspaceName),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.InitProvider.WorkspaceNameRef,
		Selector:     mg.Spec.InitProvider.WorkspaceNameSelector,
		To: reference.To{
			List:    &v1beta11.WorkspaceList{},
			Managed: &v1beta11.Workspace{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.WorkspaceName")
	}
	mg.Spec.InitProvider.WorkspaceName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.WorkspaceNameRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this SentinelWatchlist.
func (mg *SentinelWatchlist) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.LogAnalyticsWorkspaceID),
		Extract:      resource.ExtractParamPath("workspace_id", false),
		Reference:    mg.Spec.ForProvider.LogAnalyticsWorkspaceIDRef,
		Selector:     mg.Spec.ForProvider.LogAnalyticsWorkspaceIDSelector,
		To: reference.To{
			List:    &SentinelLogAnalyticsWorkspaceOnboardingList{},
			Managed: &SentinelLogAnalyticsWorkspaceOnboarding{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.LogAnalyticsWorkspaceID")
	}
	mg.Spec.ForProvider.LogAnalyticsWorkspaceID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.LogAnalyticsWorkspaceIDRef = rsp.ResolvedReference

	return nil
}
