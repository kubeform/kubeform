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

type ComputeInstanceTemplate struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ComputeInstanceTemplateSpec   `json:"spec,omitempty"`
	Status            ComputeInstanceTemplateStatus `json:"status,omitempty"`
}

type ComputeInstanceTemplateSpecDiskDiskEncryptionKey struct {
	// +optional
	KmsKeySelfLink string `json:"kmsKeySelfLink,omitempty" tf:"kms_key_self_link,omitempty"`
}

type ComputeInstanceTemplateSpecDisk struct {
	// +optional
	AutoDelete bool `json:"autoDelete,omitempty" tf:"auto_delete,omitempty"`
	// +optional
	Boot bool `json:"boot,omitempty" tf:"boot,omitempty"`
	// +optional
	DeviceName string `json:"deviceName,omitempty" tf:"device_name,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	DiskEncryptionKey []ComputeInstanceTemplateSpecDiskDiskEncryptionKey `json:"diskEncryptionKey,omitempty" tf:"disk_encryption_key,omitempty"`
	// +optional
	DiskName string `json:"diskName,omitempty" tf:"disk_name,omitempty"`
	// +optional
	DiskSizeGb int64 `json:"diskSizeGb,omitempty" tf:"disk_size_gb,omitempty"`
	// +optional
	DiskType string `json:"diskType,omitempty" tf:"disk_type,omitempty"`
	// +optional
	Interface string `json:"interface,omitempty" tf:"interface,omitempty"`
	// +optional
	Labels map[string]string `json:"labels,omitempty" tf:"labels,omitempty"`
	// +optional
	Mode string `json:"mode,omitempty" tf:"mode,omitempty"`
	// +optional
	Source string `json:"source,omitempty" tf:"source,omitempty"`
	// +optional
	SourceImage string `json:"sourceImage,omitempty" tf:"source_image,omitempty"`
	// +optional
	Type string `json:"type,omitempty" tf:"type,omitempty"`
}

type ComputeInstanceTemplateSpecGuestAccelerator struct {
	Count int64  `json:"count" tf:"count"`
	Type  string `json:"type" tf:"type"`
}

type ComputeInstanceTemplateSpecNetworkInterfaceAccessConfig struct {
	// +optional
	NatIP string `json:"natIP,omitempty" tf:"nat_ip,omitempty"`
	// +optional
	NetworkTier string `json:"networkTier,omitempty" tf:"network_tier,omitempty"`
}

type ComputeInstanceTemplateSpecNetworkInterfaceAliasIPRange struct {
	IpCIDRRange string `json:"ipCIDRRange" tf:"ip_cidr_range"`
	// +optional
	SubnetworkRangeName string `json:"subnetworkRangeName,omitempty" tf:"subnetwork_range_name,omitempty"`
}

type ComputeInstanceTemplateSpecNetworkInterface struct {
	// +optional
	AccessConfig []ComputeInstanceTemplateSpecNetworkInterfaceAccessConfig `json:"accessConfig,omitempty" tf:"access_config,omitempty"`
	// +optional
	AliasIPRange []ComputeInstanceTemplateSpecNetworkInterfaceAliasIPRange `json:"aliasIPRange,omitempty" tf:"alias_ip_range,omitempty"`
	// +optional
	Network string `json:"network,omitempty" tf:"network,omitempty"`
	// +optional
	NetworkIP string `json:"networkIP,omitempty" tf:"network_ip,omitempty"`
	// +optional
	Subnetwork string `json:"subnetwork,omitempty" tf:"subnetwork,omitempty"`
	// +optional
	SubnetworkProject string `json:"subnetworkProject,omitempty" tf:"subnetwork_project,omitempty"`
}

type ComputeInstanceTemplateSpecSchedulingNodeAffinities struct {
	Key      string   `json:"key" tf:"key"`
	Operator string   `json:"operator" tf:"operator"`
	Values   []string `json:"values" tf:"values"`
}

type ComputeInstanceTemplateSpecScheduling struct {
	// +optional
	AutomaticRestart bool `json:"automaticRestart,omitempty" tf:"automatic_restart,omitempty"`
	// +optional
	NodeAffinities []ComputeInstanceTemplateSpecSchedulingNodeAffinities `json:"nodeAffinities,omitempty" tf:"node_affinities,omitempty"`
	// +optional
	OnHostMaintenance string `json:"onHostMaintenance,omitempty" tf:"on_host_maintenance,omitempty"`
	// +optional
	Preemptible bool `json:"preemptible,omitempty" tf:"preemptible,omitempty"`
}

type ComputeInstanceTemplateSpecServiceAccount struct {
	// +optional
	Email  string   `json:"email,omitempty" tf:"email,omitempty"`
	Scopes []string `json:"scopes" tf:"scopes"`
}

type ComputeInstanceTemplateSpecShieldedInstanceConfig struct {
	// +optional
	EnableIntegrityMonitoring bool `json:"enableIntegrityMonitoring,omitempty" tf:"enable_integrity_monitoring,omitempty"`
	// +optional
	EnableSecureBoot bool `json:"enableSecureBoot,omitempty" tf:"enable_secure_boot,omitempty"`
	// +optional
	EnableVtpm bool `json:"enableVtpm,omitempty" tf:"enable_vtpm,omitempty"`
}

type ComputeInstanceTemplateSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	RemoteBackend *base.Backend `json:"remoteBackend,omitempty" tf:"-"`

	// +optional
	CanIPForward bool `json:"canIPForward,omitempty" tf:"can_ip_forward,omitempty"`
	// +optional
	Description string                            `json:"description,omitempty" tf:"description,omitempty"`
	Disk        []ComputeInstanceTemplateSpecDisk `json:"disk" tf:"disk"`
	// +optional
	GuestAccelerator []ComputeInstanceTemplateSpecGuestAccelerator `json:"guestAccelerator,omitempty" tf:"guest_accelerator,omitempty"`
	// +optional
	InstanceDescription string `json:"instanceDescription,omitempty" tf:"instance_description,omitempty"`
	// +optional
	Labels      map[string]string `json:"labels,omitempty" tf:"labels,omitempty"`
	MachineType string            `json:"machineType" tf:"machine_type"`
	// +optional
	Metadata map[string]string `json:"metadata,omitempty" tf:"metadata,omitempty"`
	// +optional
	MetadataFingerprint string `json:"metadataFingerprint,omitempty" tf:"metadata_fingerprint,omitempty"`
	// +optional
	MetadataStartupScript string `json:"metadataStartupScript,omitempty" tf:"metadata_startup_script,omitempty"`
	// +optional
	MinCPUPlatform string `json:"minCPUPlatform,omitempty" tf:"min_cpu_platform,omitempty"`
	// +optional
	Name string `json:"name,omitempty" tf:"name,omitempty"`
	// +optional
	NamePrefix string `json:"namePrefix,omitempty" tf:"name_prefix,omitempty"`
	// +optional
	NetworkInterface []ComputeInstanceTemplateSpecNetworkInterface `json:"networkInterface,omitempty" tf:"network_interface,omitempty"`
	// +optional
	Project string `json:"project,omitempty" tf:"project,omitempty"`
	// +optional
	Region string `json:"region,omitempty" tf:"region,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	Scheduling []ComputeInstanceTemplateSpecScheduling `json:"scheduling,omitempty" tf:"scheduling,omitempty"`
	// +optional
	SelfLink string `json:"selfLink,omitempty" tf:"self_link,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	ServiceAccount []ComputeInstanceTemplateSpecServiceAccount `json:"serviceAccount,omitempty" tf:"service_account,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	ShieldedInstanceConfig []ComputeInstanceTemplateSpecShieldedInstanceConfig `json:"shieldedInstanceConfig,omitempty" tf:"shielded_instance_config,omitempty"`
	// +optional
	Tags []string `json:"tags,omitempty" tf:"tags,omitempty"`
	// +optional
	TagsFingerprint string `json:"tagsFingerprint,omitempty" tf:"tags_fingerprint,omitempty"`
}

type ComputeInstanceTemplateStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *ComputeInstanceTemplateSpec `json:"output,omitempty"`
	// +optional
	State *base.State `json:"state,omitempty"`
	// +optional
	Phase base.Phase `json:"phase,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// ComputeInstanceTemplateList is a list of ComputeInstanceTemplates
type ComputeInstanceTemplateList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of ComputeInstanceTemplate CRD objects
	Items []ComputeInstanceTemplate `json:"items,omitempty"`
}
