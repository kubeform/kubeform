package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

type ComputeInstanceTemplate struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ComputeInstanceTemplateSpec   `json:"spec,omitempty"`
	Status            ComputeInstanceTemplateStatus `json:"status,omitempty"`
}

type ComputeInstanceTemplateSpecDiskDiskEncryptionKey struct {
	// +optional
	KmsKeySelfLink string `json:"kms_key_self_link,omitempty"`
}

type ComputeInstanceTemplateSpecDisk struct {
	// +optional
	AutoDelete bool `json:"auto_delete,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	DiskEncryptionKey *[]ComputeInstanceTemplateSpecDisk `json:"disk_encryption_key,omitempty"`
	// +optional
	DiskName string `json:"disk_name,omitempty"`
	// +optional
	DiskSizeGb int `json:"disk_size_gb,omitempty"`
	// +optional
	Source string `json:"source,omitempty"`
}

type ComputeInstanceTemplateSpecGuestAccelerator struct {
	Count int    `json:"count"`
	Type  string `json:"type"`
}

type ComputeInstanceTemplateSpecNetworkInterfaceAccessConfig struct{}

type ComputeInstanceTemplateSpecNetworkInterfaceAliasIpRange struct {
	IpCidrRange string `json:"ip_cidr_range"`
	// +optional
	SubnetworkRangeName string `json:"subnetwork_range_name,omitempty"`
}

type ComputeInstanceTemplateSpecNetworkInterface struct {
	// +optional
	AccessConfig *[]ComputeInstanceTemplateSpecNetworkInterface `json:"access_config,omitempty"`
	// +optional
	AliasIpRange *[]ComputeInstanceTemplateSpecNetworkInterface `json:"alias_ip_range,omitempty"`
}

type ComputeInstanceTemplateSpecServiceAccount struct {
	// +kubebuilder:validation:UniqueItems=true
	Scopes []string `json:"scopes"`
}

type ComputeInstanceTemplateSpec struct {
	// +optional
	CanIpForward bool `json:"can_ip_forward,omitempty"`
	// +optional
	Description string                        `json:"description,omitempty"`
	Disk        []ComputeInstanceTemplateSpec `json:"disk"`
	// +optional
	GuestAccelerator *[]ComputeInstanceTemplateSpec `json:"guest_accelerator,omitempty"`
	// +optional
	InstanceDescription string `json:"instance_description,omitempty"`
	// +optional
	Labels      map[string]string `json:"labels,omitempty"`
	MachineType string            `json:"machine_type"`
	// +optional
	Metadata map[string]string `json:"metadata,omitempty"`
	// +optional
	MetadataStartupScript string `json:"metadata_startup_script,omitempty"`
	// +optional
	MinCpuPlatform string `json:"min_cpu_platform,omitempty"`
	// +optional
	NetworkInterface *[]ComputeInstanceTemplateSpec `json:"network_interface,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	ServiceAccount *[]ComputeInstanceTemplateSpec `json:"service_account,omitempty"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	Tags []string `json:"tags,omitempty"`
}

type ComputeInstanceTemplateStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
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
