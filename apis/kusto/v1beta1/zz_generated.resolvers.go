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
	v1beta14 "github.com/upbound/provider-azure/apis/devices/v1beta1"
	v1beta13 "github.com/upbound/provider-azure/apis/eventhub/v1beta1"
	v1beta11 "github.com/upbound/provider-azure/apis/network/v1beta1"
	rconfig "github.com/upbound/provider-azure/apis/rconfig"
	v1beta12 "github.com/upbound/provider-azure/apis/storage/v1beta1"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

// ResolveReferences of this AttachedDatabaseConfiguration.
func (mg *AttachedDatabaseConfiguration) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ClusterName),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.ClusterNameRef,
		Selector:     mg.Spec.ForProvider.ClusterNameSelector,
		To: reference.To{
			List:    &ClusterList{},
			Managed: &Cluster{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ClusterName")
	}
	mg.Spec.ForProvider.ClusterName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ClusterNameRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ClusterResourceID),
		Extract:      resource.ExtractResourceID(),
		Reference:    mg.Spec.ForProvider.ClusterResourceIDRef,
		Selector:     mg.Spec.ForProvider.ClusterResourceIDSelector,
		To: reference.To{
			List:    &ClusterList{},
			Managed: &Cluster{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ClusterResourceID")
	}
	mg.Spec.ForProvider.ClusterResourceID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ClusterResourceIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.DatabaseName),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.DatabaseNameRef,
		Selector:     mg.Spec.ForProvider.DatabaseNameSelector,
		To: reference.To{
			List:    &DatabaseList{},
			Managed: &Database{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.DatabaseName")
	}
	mg.Spec.ForProvider.DatabaseName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.DatabaseNameRef = rsp.ResolvedReference

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

// ResolveReferences of this Cluster.
func (mg *Cluster) ResolveReferences(ctx context.Context, c client.Reader) error {
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

	for i3 := 0; i3 < len(mg.Spec.ForProvider.VirtualNetworkConfiguration); i3++ {
		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.VirtualNetworkConfiguration[i3].SubnetID),
			Extract:      rconfig.ExtractResourceID(),
			Reference:    mg.Spec.ForProvider.VirtualNetworkConfiguration[i3].SubnetIDRef,
			Selector:     mg.Spec.ForProvider.VirtualNetworkConfiguration[i3].SubnetIDSelector,
			To: reference.To{
				List:    &v1beta11.SubnetList{},
				Managed: &v1beta11.Subnet{},
			},
		})
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.VirtualNetworkConfiguration[i3].SubnetID")
		}
		mg.Spec.ForProvider.VirtualNetworkConfiguration[i3].SubnetID = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.ForProvider.VirtualNetworkConfiguration[i3].SubnetIDRef = rsp.ResolvedReference

	}

	return nil
}

// ResolveReferences of this ClusterManagedPrivateEndpoint.
func (mg *ClusterManagedPrivateEndpoint) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ClusterName),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.ClusterNameRef,
		Selector:     mg.Spec.ForProvider.ClusterNameSelector,
		To: reference.To{
			List:    &ClusterList{},
			Managed: &Cluster{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ClusterName")
	}
	mg.Spec.ForProvider.ClusterName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ClusterNameRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.PrivateLinkResourceID),
		Extract:      resource.ExtractResourceID(),
		Reference:    mg.Spec.ForProvider.PrivateLinkResourceIDRef,
		Selector:     mg.Spec.ForProvider.PrivateLinkResourceIDSelector,
		To: reference.To{
			List:    &v1beta12.AccountList{},
			Managed: &v1beta12.Account{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.PrivateLinkResourceID")
	}
	mg.Spec.ForProvider.PrivateLinkResourceID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.PrivateLinkResourceIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.PrivateLinkResourceRegion),
		Extract:      resource.ExtractParamPath("location", false),
		Reference:    mg.Spec.ForProvider.PrivateLinkResourceRegionRef,
		Selector:     mg.Spec.ForProvider.PrivateLinkResourceRegionSelector,
		To: reference.To{
			List:    &v1beta12.AccountList{},
			Managed: &v1beta12.Account{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.PrivateLinkResourceRegion")
	}
	mg.Spec.ForProvider.PrivateLinkResourceRegion = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.PrivateLinkResourceRegionRef = rsp.ResolvedReference

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

// ResolveReferences of this ClusterPrincipalAssignment.
func (mg *ClusterPrincipalAssignment) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ClusterName),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.ClusterNameRef,
		Selector:     mg.Spec.ForProvider.ClusterNameSelector,
		To: reference.To{
			List:    &ClusterList{},
			Managed: &Cluster{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ClusterName")
	}
	mg.Spec.ForProvider.ClusterName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ClusterNameRef = rsp.ResolvedReference

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

// ResolveReferences of this Database.
func (mg *Database) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ClusterName),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.ClusterNameRef,
		Selector:     mg.Spec.ForProvider.ClusterNameSelector,
		To: reference.To{
			List:    &ClusterList{},
			Managed: &Cluster{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ClusterName")
	}
	mg.Spec.ForProvider.ClusterName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ClusterNameRef = rsp.ResolvedReference

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

// ResolveReferences of this DatabasePrincipalAssignment.
func (mg *DatabasePrincipalAssignment) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ClusterName),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.ClusterNameRef,
		Selector:     mg.Spec.ForProvider.ClusterNameSelector,
		To: reference.To{
			List:    &ClusterList{},
			Managed: &Cluster{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ClusterName")
	}
	mg.Spec.ForProvider.ClusterName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ClusterNameRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.DatabaseName),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.DatabaseNameRef,
		Selector:     mg.Spec.ForProvider.DatabaseNameSelector,
		To: reference.To{
			List:    &DatabaseList{},
			Managed: &Database{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.DatabaseName")
	}
	mg.Spec.ForProvider.DatabaseName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.DatabaseNameRef = rsp.ResolvedReference

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

// ResolveReferences of this EventGridDataConnection.
func (mg *EventGridDataConnection) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ClusterName),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.ClusterNameRef,
		Selector:     mg.Spec.ForProvider.ClusterNameSelector,
		To: reference.To{
			List:    &ClusterList{},
			Managed: &Cluster{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ClusterName")
	}
	mg.Spec.ForProvider.ClusterName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ClusterNameRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.DatabaseName),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.DatabaseNameRef,
		Selector:     mg.Spec.ForProvider.DatabaseNameSelector,
		To: reference.To{
			List:    &DatabaseList{},
			Managed: &Database{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.DatabaseName")
	}
	mg.Spec.ForProvider.DatabaseName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.DatabaseNameRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.EventHubConsumerGroupName),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.EventHubConsumerGroupNameRef,
		Selector:     mg.Spec.ForProvider.EventHubConsumerGroupNameSelector,
		To: reference.To{
			List:    &v1beta13.ConsumerGroupList{},
			Managed: &v1beta13.ConsumerGroup{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.EventHubConsumerGroupName")
	}
	mg.Spec.ForProvider.EventHubConsumerGroupName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.EventHubConsumerGroupNameRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.EventHubID),
		Extract:      resource.ExtractResourceID(),
		Reference:    mg.Spec.ForProvider.EventHubIDRef,
		Selector:     mg.Spec.ForProvider.EventHubIDSelector,
		To: reference.To{
			List:    &v1beta13.EventHubList{},
			Managed: &v1beta13.EventHub{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.EventHubID")
	}
	mg.Spec.ForProvider.EventHubID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.EventHubIDRef = rsp.ResolvedReference

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
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.StorageAccountID),
		Extract:      resource.ExtractResourceID(),
		Reference:    mg.Spec.ForProvider.StorageAccountIDRef,
		Selector:     mg.Spec.ForProvider.StorageAccountIDSelector,
		To: reference.To{
			List:    &v1beta12.AccountList{},
			Managed: &v1beta12.Account{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.StorageAccountID")
	}
	mg.Spec.ForProvider.StorageAccountID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.StorageAccountIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this EventHubDataConnection.
func (mg *EventHubDataConnection) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ClusterName),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.ClusterNameRef,
		Selector:     mg.Spec.ForProvider.ClusterNameSelector,
		To: reference.To{
			List:    &ClusterList{},
			Managed: &Cluster{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ClusterName")
	}
	mg.Spec.ForProvider.ClusterName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ClusterNameRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ConsumerGroup),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.ConsumerGroupRef,
		Selector:     mg.Spec.ForProvider.ConsumerGroupSelector,
		To: reference.To{
			List:    &v1beta13.ConsumerGroupList{},
			Managed: &v1beta13.ConsumerGroup{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ConsumerGroup")
	}
	mg.Spec.ForProvider.ConsumerGroup = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ConsumerGroupRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.DatabaseName),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.DatabaseNameRef,
		Selector:     mg.Spec.ForProvider.DatabaseNameSelector,
		To: reference.To{
			List:    &DatabaseList{},
			Managed: &Database{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.DatabaseName")
	}
	mg.Spec.ForProvider.DatabaseName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.DatabaseNameRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.EventHubID),
		Extract:      resource.ExtractResourceID(),
		Reference:    mg.Spec.ForProvider.EventHubIDRef,
		Selector:     mg.Spec.ForProvider.EventHubIDSelector,
		To: reference.To{
			List:    &v1beta13.EventHubList{},
			Managed: &v1beta13.EventHub{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.EventHubID")
	}
	mg.Spec.ForProvider.EventHubID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.EventHubIDRef = rsp.ResolvedReference

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

// ResolveReferences of this IOTHubDataConnection.
func (mg *IOTHubDataConnection) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ClusterName),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.ClusterNameRef,
		Selector:     mg.Spec.ForProvider.ClusterNameSelector,
		To: reference.To{
			List:    &ClusterList{},
			Managed: &Cluster{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ClusterName")
	}
	mg.Spec.ForProvider.ClusterName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ClusterNameRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ConsumerGroup),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.ConsumerGroupRef,
		Selector:     mg.Spec.ForProvider.ConsumerGroupSelector,
		To: reference.To{
			List:    &v1beta14.IOTHubConsumerGroupList{},
			Managed: &v1beta14.IOTHubConsumerGroup{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ConsumerGroup")
	}
	mg.Spec.ForProvider.ConsumerGroup = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ConsumerGroupRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.DatabaseName),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.DatabaseNameRef,
		Selector:     mg.Spec.ForProvider.DatabaseNameSelector,
		To: reference.To{
			List:    &DatabaseList{},
			Managed: &Database{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.DatabaseName")
	}
	mg.Spec.ForProvider.DatabaseName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.DatabaseNameRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.IOTHubID),
		Extract:      resource.ExtractResourceID(),
		Reference:    mg.Spec.ForProvider.IOTHubIDRef,
		Selector:     mg.Spec.ForProvider.IOTHubIDSelector,
		To: reference.To{
			List:    &v1beta14.IOTHubList{},
			Managed: &v1beta14.IOTHub{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.IOTHubID")
	}
	mg.Spec.ForProvider.IOTHubID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.IOTHubIDRef = rsp.ResolvedReference

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
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.SharedAccessPolicyName),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.SharedAccessPolicyNameRef,
		Selector:     mg.Spec.ForProvider.SharedAccessPolicyNameSelector,
		To: reference.To{
			List:    &v1beta14.IOTHubSharedAccessPolicyList{},
			Managed: &v1beta14.IOTHubSharedAccessPolicy{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.SharedAccessPolicyName")
	}
	mg.Spec.ForProvider.SharedAccessPolicyName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.SharedAccessPolicyNameRef = rsp.ResolvedReference

	return nil
}
