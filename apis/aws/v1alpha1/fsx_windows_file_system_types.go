/*
Copyright The Kubeform Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by Kubeform. DO NOT EDIT.

package v1alpha1

import (
	base "kubeform.dev/kubeform/apis/base/v1alpha1"

	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Phase",type=string,JSONPath=`.status.phase`

type FsxWindowsFileSystem struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              FsxWindowsFileSystemSpec   `json:"spec,omitempty"`
	Status            FsxWindowsFileSystemStatus `json:"status,omitempty"`
}

type FsxWindowsFileSystemSpecSelfManagedActiveDirectory struct {
	// +kubebuilder:validation:MaxItems=2
	// +kubebuilder:validation:MinItems=1
	DnsIPS     []string `json:"dnsIPS" tf:"dns_ips"`
	DomainName string   `json:"domainName" tf:"domain_name"`
	// +optional
	FileSystemAdministratorsGroup string `json:"fileSystemAdministratorsGroup,omitempty" tf:"file_system_administrators_group,omitempty"`
	// +optional
	OrganizationalUnitDistinguishedName string `json:"organizationalUnitDistinguishedName,omitempty" tf:"organizational_unit_distinguished_name,omitempty"`
	Password                            string `json:"-" sensitive:"true" tf:"password"`
	Username                            string `json:"username" tf:"username"`
}

type FsxWindowsFileSystemSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	SecretRef *core.LocalObjectReference `json:"secretRef,omitempty" tf:"-"`

	// +optional
	ActiveDirectoryID string `json:"activeDirectoryID,omitempty" tf:"active_directory_id,omitempty"`
	// +optional
	Arn string `json:"arn,omitempty" tf:"arn,omitempty"`
	// +optional
	AutomaticBackupRetentionDays int64 `json:"automaticBackupRetentionDays,omitempty" tf:"automatic_backup_retention_days,omitempty"`
	// +optional
	CopyTagsToBackups bool `json:"copyTagsToBackups,omitempty" tf:"copy_tags_to_backups,omitempty"`
	// +optional
	DailyAutomaticBackupStartTime string `json:"dailyAutomaticBackupStartTime,omitempty" tf:"daily_automatic_backup_start_time,omitempty"`
	// +optional
	DnsName string `json:"dnsName,omitempty" tf:"dns_name,omitempty"`
	// +optional
	KmsKeyID string `json:"kmsKeyID,omitempty" tf:"kms_key_id,omitempty"`
	// +optional
	NetworkInterfaceIDS []string `json:"networkInterfaceIDS,omitempty" tf:"network_interface_ids,omitempty"`
	// +optional
	OwnerID string `json:"ownerID,omitempty" tf:"owner_id,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=50
	SecurityGroupIDS []string `json:"securityGroupIDS,omitempty" tf:"security_group_ids,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	SelfManagedActiveDirectory []FsxWindowsFileSystemSpecSelfManagedActiveDirectory `json:"selfManagedActiveDirectory,omitempty" tf:"self_managed_active_directory,omitempty"`
	// +optional
	SkipFinalBackup bool  `json:"skipFinalBackup,omitempty" tf:"skip_final_backup,omitempty"`
	StorageCapacity int64 `json:"storageCapacity" tf:"storage_capacity"`
	// +kubebuilder:validation:MaxItems=1
	// +kubebuilder:validation:MinItems=1
	SubnetIDS []string `json:"subnetIDS" tf:"subnet_ids"`
	// +optional
	Tags               map[string]string `json:"tags,omitempty" tf:"tags,omitempty"`
	ThroughputCapacity int64             `json:"throughputCapacity" tf:"throughput_capacity"`
	// +optional
	VpcID string `json:"vpcID,omitempty" tf:"vpc_id,omitempty"`
	// +optional
	WeeklyMaintenanceStartTime string `json:"weeklyMaintenanceStartTime,omitempty" tf:"weekly_maintenance_start_time,omitempty"`
}

type FsxWindowsFileSystemStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *FsxWindowsFileSystemSpec `json:"output,omitempty"`
	// +optional
	State *base.State `json:"state,omitempty"`
	// +optional
	Phase base.Phase `json:"phase,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// FsxWindowsFileSystemList is a list of FsxWindowsFileSystems
type FsxWindowsFileSystemList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of FsxWindowsFileSystem CRD objects
	Items []FsxWindowsFileSystem `json:"items,omitempty"`
}
