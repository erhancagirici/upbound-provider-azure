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

type DataSetDataLakeGen2InitParameters struct {

	// The path of the file in the data lake file system to be shared with the receiver. Conflicts with folder_path Changing this forces a new Data Share Data Lake Gen2 Dataset to be created.
	FilePath *string `json:"filePath,omitempty" tf:"file_path,omitempty"`

	// The name of the data lake file system to be shared with the receiver. Changing this forces a new Data Share Data Lake Gen2 Dataset to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/storage/v1beta1.DataLakeGen2FileSystem
	FileSystemName *string `json:"fileSystemName,omitempty" tf:"file_system_name,omitempty"`

	// Reference to a DataLakeGen2FileSystem in storage to populate fileSystemName.
	// +kubebuilder:validation:Optional
	FileSystemNameRef *v1.Reference `json:"fileSystemNameRef,omitempty" tf:"-"`

	// Selector for a DataLakeGen2FileSystem in storage to populate fileSystemName.
	// +kubebuilder:validation:Optional
	FileSystemNameSelector *v1.Selector `json:"fileSystemNameSelector,omitempty" tf:"-"`

	// The folder path in the data lake file system to be shared with the receiver. Conflicts with file_path Changing this forces a new Data Share Data Lake Gen2 Dataset to be created.
	FolderPath *string `json:"folderPath,omitempty" tf:"folder_path,omitempty"`

	// The resource id of the storage account of the data lake file system to be shared with the receiver. Changing this forces a new Data Share Data Lake Gen2 Dataset to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/storage/v1beta1.Account
	// +crossplane:generate:reference:extractor=github.com/upbound/provider-azure/apis/rconfig.ExtractResourceID()
	StorageAccountID *string `json:"storageAccountId,omitempty" tf:"storage_account_id,omitempty"`

	// Reference to a Account in storage to populate storageAccountId.
	// +kubebuilder:validation:Optional
	StorageAccountIDRef *v1.Reference `json:"storageAccountIdRef,omitempty" tf:"-"`

	// Selector for a Account in storage to populate storageAccountId.
	// +kubebuilder:validation:Optional
	StorageAccountIDSelector *v1.Selector `json:"storageAccountIdSelector,omitempty" tf:"-"`
}

type DataSetDataLakeGen2Observation struct {

	// The name of the Data Share Dataset.
	DisplayName *string `json:"displayName,omitempty" tf:"display_name,omitempty"`

	// The path of the file in the data lake file system to be shared with the receiver. Conflicts with folder_path Changing this forces a new Data Share Data Lake Gen2 Dataset to be created.
	FilePath *string `json:"filePath,omitempty" tf:"file_path,omitempty"`

	// The name of the data lake file system to be shared with the receiver. Changing this forces a new Data Share Data Lake Gen2 Dataset to be created.
	FileSystemName *string `json:"fileSystemName,omitempty" tf:"file_system_name,omitempty"`

	// The folder path in the data lake file system to be shared with the receiver. Conflicts with file_path Changing this forces a new Data Share Data Lake Gen2 Dataset to be created.
	FolderPath *string `json:"folderPath,omitempty" tf:"folder_path,omitempty"`

	// The resource ID of the Data Share Data Lake Gen2 Dataset.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The resource ID of the Data Share where this Data Share Data Lake Gen2 Dataset should be created. Changing this forces a new Data Share Data Lake Gen2 Dataset to be created.
	ShareID *string `json:"shareId,omitempty" tf:"share_id,omitempty"`

	// The resource id of the storage account of the data lake file system to be shared with the receiver. Changing this forces a new Data Share Data Lake Gen2 Dataset to be created.
	StorageAccountID *string `json:"storageAccountId,omitempty" tf:"storage_account_id,omitempty"`
}

type DataSetDataLakeGen2Parameters struct {

	// The path of the file in the data lake file system to be shared with the receiver. Conflicts with folder_path Changing this forces a new Data Share Data Lake Gen2 Dataset to be created.
	// +kubebuilder:validation:Optional
	FilePath *string `json:"filePath,omitempty" tf:"file_path,omitempty"`

	// The name of the data lake file system to be shared with the receiver. Changing this forces a new Data Share Data Lake Gen2 Dataset to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/storage/v1beta1.DataLakeGen2FileSystem
	// +kubebuilder:validation:Optional
	FileSystemName *string `json:"fileSystemName,omitempty" tf:"file_system_name,omitempty"`

	// Reference to a DataLakeGen2FileSystem in storage to populate fileSystemName.
	// +kubebuilder:validation:Optional
	FileSystemNameRef *v1.Reference `json:"fileSystemNameRef,omitempty" tf:"-"`

	// Selector for a DataLakeGen2FileSystem in storage to populate fileSystemName.
	// +kubebuilder:validation:Optional
	FileSystemNameSelector *v1.Selector `json:"fileSystemNameSelector,omitempty" tf:"-"`

	// The folder path in the data lake file system to be shared with the receiver. Conflicts with file_path Changing this forces a new Data Share Data Lake Gen2 Dataset to be created.
	// +kubebuilder:validation:Optional
	FolderPath *string `json:"folderPath,omitempty" tf:"folder_path,omitempty"`

	// The resource ID of the Data Share where this Data Share Data Lake Gen2 Dataset should be created. Changing this forces a new Data Share Data Lake Gen2 Dataset to be created.
	// +crossplane:generate:reference:type=DataShare
	// +crossplane:generate:reference:extractor=github.com/upbound/provider-azure/apis/rconfig.ExtractResourceID()
	// +kubebuilder:validation:Optional
	ShareID *string `json:"shareId,omitempty" tf:"share_id,omitempty"`

	// Reference to a DataShare to populate shareId.
	// +kubebuilder:validation:Optional
	ShareIDRef *v1.Reference `json:"shareIdRef,omitempty" tf:"-"`

	// Selector for a DataShare to populate shareId.
	// +kubebuilder:validation:Optional
	ShareIDSelector *v1.Selector `json:"shareIdSelector,omitempty" tf:"-"`

	// The resource id of the storage account of the data lake file system to be shared with the receiver. Changing this forces a new Data Share Data Lake Gen2 Dataset to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/storage/v1beta1.Account
	// +crossplane:generate:reference:extractor=github.com/upbound/provider-azure/apis/rconfig.ExtractResourceID()
	// +kubebuilder:validation:Optional
	StorageAccountID *string `json:"storageAccountId,omitempty" tf:"storage_account_id,omitempty"`

	// Reference to a Account in storage to populate storageAccountId.
	// +kubebuilder:validation:Optional
	StorageAccountIDRef *v1.Reference `json:"storageAccountIdRef,omitempty" tf:"-"`

	// Selector for a Account in storage to populate storageAccountId.
	// +kubebuilder:validation:Optional
	StorageAccountIDSelector *v1.Selector `json:"storageAccountIdSelector,omitempty" tf:"-"`
}

// DataSetDataLakeGen2Spec defines the desired state of DataSetDataLakeGen2
type DataSetDataLakeGen2Spec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     DataSetDataLakeGen2Parameters `json:"forProvider"`
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
	InitProvider DataSetDataLakeGen2InitParameters `json:"initProvider,omitempty"`
}

// DataSetDataLakeGen2Status defines the observed state of DataSetDataLakeGen2.
type DataSetDataLakeGen2Status struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        DataSetDataLakeGen2Observation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// DataSetDataLakeGen2 is the Schema for the DataSetDataLakeGen2s API. Manages a Data Share Data Lake Gen2 Dataset.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type DataSetDataLakeGen2 struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DataSetDataLakeGen2Spec   `json:"spec"`
	Status            DataSetDataLakeGen2Status `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DataSetDataLakeGen2List contains a list of DataSetDataLakeGen2s
type DataSetDataLakeGen2List struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DataSetDataLakeGen2 `json:"items"`
}

// Repository type metadata.
var (
	DataSetDataLakeGen2_Kind             = "DataSetDataLakeGen2"
	DataSetDataLakeGen2_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: DataSetDataLakeGen2_Kind}.String()
	DataSetDataLakeGen2_KindAPIVersion   = DataSetDataLakeGen2_Kind + "." + CRDGroupVersion.String()
	DataSetDataLakeGen2_GroupVersionKind = CRDGroupVersion.WithKind(DataSetDataLakeGen2_Kind)
)

func init() {
	SchemeBuilder.Register(&DataSetDataLakeGen2{}, &DataSetDataLakeGen2List{})
}
