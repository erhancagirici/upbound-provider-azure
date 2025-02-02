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
	v1beta1 "github.com/upbound/provider-azure/apis/azure/v1beta1"
	v1beta11 "github.com/upbound/provider-azure/apis/storage/v1beta1"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

// ResolveReferences of this Domain.
func (mg *Domain) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ResourceGroupName),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.ResourceGroupNameRef,
		Selector:     mg.Spec.ForProvider.ResourceGroupNameSelector,
		To: reference.To{
			List:    &v1beta1.ResourceGroupList{},
			Managed: &v1beta1.ResourceGroup{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ResourceGroupName")
	}
	mg.Spec.ForProvider.ResourceGroupName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ResourceGroupNameRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this DomainTopic.
func (mg *DomainTopic) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.DomainName),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.DomainNameRef,
		Selector:     mg.Spec.ForProvider.DomainNameSelector,
		To: reference.To{
			List:    &DomainList{},
			Managed: &Domain{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.DomainName")
	}
	mg.Spec.ForProvider.DomainName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.DomainNameRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ResourceGroupName),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.ResourceGroupNameRef,
		Selector:     mg.Spec.ForProvider.ResourceGroupNameSelector,
		To: reference.To{
			List:    &v1beta1.ResourceGroupList{},
			Managed: &v1beta1.ResourceGroup{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ResourceGroupName")
	}
	mg.Spec.ForProvider.ResourceGroupName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ResourceGroupNameRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this EventSubscription.
func (mg *EventSubscription) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Scope),
		Extract:      resource.ExtractResourceID(),
		Reference:    mg.Spec.ForProvider.ScopeRef,
		Selector:     mg.Spec.ForProvider.ScopeSelector,
		To: reference.To{
			List:    &v1beta1.ResourceGroupList{},
			Managed: &v1beta1.ResourceGroup{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.Scope")
	}
	mg.Spec.ForProvider.Scope = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ScopeRef = rsp.ResolvedReference

	for i3 := 0; i3 < len(mg.Spec.ForProvider.StorageQueueEndpoint); i3++ {
		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.StorageQueueEndpoint[i3].QueueName),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.ForProvider.StorageQueueEndpoint[i3].QueueNameRef,
			Selector:     mg.Spec.ForProvider.StorageQueueEndpoint[i3].QueueNameSelector,
			To: reference.To{
				List:    &v1beta11.QueueList{},
				Managed: &v1beta11.Queue{},
			},
		})
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.StorageQueueEndpoint[i3].QueueName")
		}
		mg.Spec.ForProvider.StorageQueueEndpoint[i3].QueueName = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.ForProvider.StorageQueueEndpoint[i3].QueueNameRef = rsp.ResolvedReference

	}
	for i3 := 0; i3 < len(mg.Spec.ForProvider.StorageQueueEndpoint); i3++ {
		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.StorageQueueEndpoint[i3].StorageAccountID),
			Extract:      resource.ExtractResourceID(),
			Reference:    mg.Spec.ForProvider.StorageQueueEndpoint[i3].StorageAccountIDRef,
			Selector:     mg.Spec.ForProvider.StorageQueueEndpoint[i3].StorageAccountIDSelector,
			To: reference.To{
				List:    &v1beta11.AccountList{},
				Managed: &v1beta11.Account{},
			},
		})
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.StorageQueueEndpoint[i3].StorageAccountID")
		}
		mg.Spec.ForProvider.StorageQueueEndpoint[i3].StorageAccountID = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.ForProvider.StorageQueueEndpoint[i3].StorageAccountIDRef = rsp.ResolvedReference

	}

	return nil
}

// ResolveReferences of this SystemTopic.
func (mg *SystemTopic) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ResourceGroupName),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.ResourceGroupNameRef,
		Selector:     mg.Spec.ForProvider.ResourceGroupNameSelector,
		To: reference.To{
			List:    &v1beta1.ResourceGroupList{},
			Managed: &v1beta1.ResourceGroup{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ResourceGroupName")
	}
	mg.Spec.ForProvider.ResourceGroupName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ResourceGroupNameRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.SourceArmResourceID),
		Extract:      resource.ExtractResourceID(),
		Reference:    mg.Spec.ForProvider.SourceArmResourceIDRef,
		Selector:     mg.Spec.ForProvider.SourceArmResourceIDSelector,
		To: reference.To{
			List:    &v1beta11.AccountList{},
			Managed: &v1beta11.Account{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.SourceArmResourceID")
	}
	mg.Spec.ForProvider.SourceArmResourceID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.SourceArmResourceIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this Topic.
func (mg *Topic) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ResourceGroupName),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.ResourceGroupNameRef,
		Selector:     mg.Spec.ForProvider.ResourceGroupNameSelector,
		To: reference.To{
			List:    &v1beta1.ResourceGroupList{},
			Managed: &v1beta1.ResourceGroup{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ResourceGroupName")
	}
	mg.Spec.ForProvider.ResourceGroupName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ResourceGroupNameRef = rsp.ResolvedReference

	return nil
}
