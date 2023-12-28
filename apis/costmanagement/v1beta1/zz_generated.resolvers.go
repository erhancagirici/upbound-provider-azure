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
	v1beta11 "github.com/upbound/provider-azure/apis/azure/v1beta1"
	v1beta1 "github.com/upbound/provider-azure/apis/storage/v1beta1"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

// ResolveReferences of this ResourceGroupCostManagementExport.
func (mg *ResourceGroupCostManagementExport) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	for i3 := 0; i3 < len(mg.Spec.ForProvider.ExportDataStorageLocation); i3++ {
		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ExportDataStorageLocation[i3].ContainerID),
			Extract:      resource.ExtractParamPath("resource_manager_id", true),
			Reference:    mg.Spec.ForProvider.ExportDataStorageLocation[i3].ContainerIDRef,
			Selector:     mg.Spec.ForProvider.ExportDataStorageLocation[i3].ContainerIDSelector,
			To: reference.To{
				List:    &v1beta1.ContainerList{},
				Managed: &v1beta1.Container{},
			},
		})
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.ExportDataStorageLocation[i3].ContainerID")
		}
		mg.Spec.ForProvider.ExportDataStorageLocation[i3].ContainerID = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.ForProvider.ExportDataStorageLocation[i3].ContainerIDRef = rsp.ResolvedReference

	}
	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ResourceGroupID),
		Extract:      resource.ExtractResourceID(),
		Reference:    mg.Spec.ForProvider.ResourceGroupIDRef,
		Selector:     mg.Spec.ForProvider.ResourceGroupIDSelector,
		To: reference.To{
			List:    &v1beta11.ResourceGroupList{},
			Managed: &v1beta11.ResourceGroup{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ResourceGroupID")
	}
	mg.Spec.ForProvider.ResourceGroupID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ResourceGroupIDRef = rsp.ResolvedReference

	for i3 := 0; i3 < len(mg.Spec.InitProvider.ExportDataStorageLocation); i3++ {
		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.ExportDataStorageLocation[i3].ContainerID),
			Extract:      resource.ExtractParamPath("resource_manager_id", true),
			Reference:    mg.Spec.InitProvider.ExportDataStorageLocation[i3].ContainerIDRef,
			Selector:     mg.Spec.InitProvider.ExportDataStorageLocation[i3].ContainerIDSelector,
			To: reference.To{
				List:    &v1beta1.ContainerList{},
				Managed: &v1beta1.Container{},
			},
		})
		if err != nil {
			return errors.Wrap(err, "mg.Spec.InitProvider.ExportDataStorageLocation[i3].ContainerID")
		}
		mg.Spec.InitProvider.ExportDataStorageLocation[i3].ContainerID = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.InitProvider.ExportDataStorageLocation[i3].ContainerIDRef = rsp.ResolvedReference

	}
	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.ResourceGroupID),
		Extract:      resource.ExtractResourceID(),
		Reference:    mg.Spec.InitProvider.ResourceGroupIDRef,
		Selector:     mg.Spec.InitProvider.ResourceGroupIDSelector,
		To: reference.To{
			List:    &v1beta11.ResourceGroupList{},
			Managed: &v1beta11.ResourceGroup{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.ResourceGroupID")
	}
	mg.Spec.InitProvider.ResourceGroupID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.ResourceGroupIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this SubscriptionCostManagementExport.
func (mg *SubscriptionCostManagementExport) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	for i3 := 0; i3 < len(mg.Spec.ForProvider.ExportDataStorageLocation); i3++ {
		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ExportDataStorageLocation[i3].ContainerID),
			Extract:      resource.ExtractParamPath("resource_manager_id", true),
			Reference:    mg.Spec.ForProvider.ExportDataStorageLocation[i3].ContainerIDRef,
			Selector:     mg.Spec.ForProvider.ExportDataStorageLocation[i3].ContainerIDSelector,
			To: reference.To{
				List:    &v1beta1.ContainerList{},
				Managed: &v1beta1.Container{},
			},
		})
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.ExportDataStorageLocation[i3].ContainerID")
		}
		mg.Spec.ForProvider.ExportDataStorageLocation[i3].ContainerID = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.ForProvider.ExportDataStorageLocation[i3].ContainerIDRef = rsp.ResolvedReference

	}
	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.SubscriptionID),
		Extract:      resource.ExtractResourceID(),
		Reference:    mg.Spec.ForProvider.SubscriptionIDRef,
		Selector:     mg.Spec.ForProvider.SubscriptionIDSelector,
		To: reference.To{
			List:    &v1beta11.SubscriptionList{},
			Managed: &v1beta11.Subscription{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.SubscriptionID")
	}
	mg.Spec.ForProvider.SubscriptionID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.SubscriptionIDRef = rsp.ResolvedReference

	for i3 := 0; i3 < len(mg.Spec.InitProvider.ExportDataStorageLocation); i3++ {
		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.ExportDataStorageLocation[i3].ContainerID),
			Extract:      resource.ExtractParamPath("resource_manager_id", true),
			Reference:    mg.Spec.InitProvider.ExportDataStorageLocation[i3].ContainerIDRef,
			Selector:     mg.Spec.InitProvider.ExportDataStorageLocation[i3].ContainerIDSelector,
			To: reference.To{
				List:    &v1beta1.ContainerList{},
				Managed: &v1beta1.Container{},
			},
		})
		if err != nil {
			return errors.Wrap(err, "mg.Spec.InitProvider.ExportDataStorageLocation[i3].ContainerID")
		}
		mg.Spec.InitProvider.ExportDataStorageLocation[i3].ContainerID = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.InitProvider.ExportDataStorageLocation[i3].ContainerIDRef = rsp.ResolvedReference

	}
	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.SubscriptionID),
		Extract:      resource.ExtractResourceID(),
		Reference:    mg.Spec.InitProvider.SubscriptionIDRef,
		Selector:     mg.Spec.InitProvider.SubscriptionIDSelector,
		To: reference.To{
			List:    &v1beta11.SubscriptionList{},
			Managed: &v1beta11.Subscription{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.SubscriptionID")
	}
	mg.Spec.InitProvider.SubscriptionID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.SubscriptionIDRef = rsp.ResolvedReference

	return nil
}
