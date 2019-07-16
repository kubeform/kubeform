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

type ComputeInstance struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ComputeInstanceSpec   `json:"spec,omitempty"`
	Status            ComputeInstanceStatus `json:"status,omitempty"`
}

type ComputeInstanceSpecAttachedDisk struct {
	// +optional
	DiskEncryptionKeyRaw string `json:"disk_encryption_key_raw,omitempty"`
	// +optional
	Mode   string `json:"mode,omitempty"`
	Source string `json:"source"`
}

type ComputeInstanceSpecBootDisk struct {
	// +optional
	AutoDelete bool `json:"auto_delete,omitempty"`
	// +optional
	DiskEncryptionKeyRaw string `json:"disk_encryption_key_raw,omitempty"`
}

type ComputeInstanceSpecNetworkInterfaceAccessConfig struct {
	// +optional
	PublicPtrDomainName string `json:"public_ptr_domain_name,omitempty"`
}

type ComputeInstanceSpecNetworkInterfaceAliasIpRange struct {
	IpCidrRange string `json:"ip_cidr_range"`
	// +optional
	SubnetworkRangeName string `json:"subnetwork_range_name,omitempty"`
}

type ComputeInstanceSpecNetworkInterface struct {
	// +optional
	AccessConfig *[]ComputeInstanceSpecNetworkInterface `json:"access_config,omitempty"`
	// +optional
	AliasIpRange *[]ComputeInstanceSpecNetworkInterface `json:"alias_ip_range,omitempty"`
}

type ComputeInstanceSpecScratchDisk struct {
	// +optional
	Interface string `json:"interface,omitempty"`
}

type ComputeInstanceSpecServiceAccount struct {
	// +kubebuilder:validation:UniqueItems=true
	Scopes []string `json:"scopes"`
}

type ComputeInstanceSpec struct {
	// +optional
	AllowStoppingForUpdate bool `json:"allow_stopping_for_update,omitempty"`
	// +optional
	AttachedDisk *[]ComputeInstanceSpec `json:"attached_disk,omitempty"`
	// +kubebuilder:validation:MaxItems=1
	BootDisk []ComputeInstanceSpec `json:"boot_disk"`
	// +optional
	CanIpForward bool `json:"can_ip_forward,omitempty"`
	// +optional
	// Deprecated
	CreateTimeout int `json:"create_timeout,omitempty"`
	// +optional
	DeletionProtection bool `json:"deletion_protection,omitempty"`
	// +optional
	Description string `json:"description,omitempty"`
	// +optional
	Labels      map[string]string `json:"labels,omitempty"`
	MachineType string            `json:"machine_type"`
	// +optional
	Metadata map[string]string `json:"metadata,omitempty"`
	// +optional
	MetadataStartupScript string `json:"metadata_startup_script,omitempty"`
	// +optional
	MinCpuPlatform   string                `json:"min_cpu_platform,omitempty"`
	Name             string                `json:"name"`
	NetworkInterface []ComputeInstanceSpec `json:"network_interface"`
	// +optional
	ScratchDisk *[]ComputeInstanceSpec `json:"scratch_disk,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	ServiceAccount *[]ComputeInstanceSpec `json:"service_account,omitempty"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	Tags []string `json:"tags,omitempty"`
}

type ComputeInstanceStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// ComputeInstanceList is a list of ComputeInstances
type ComputeInstanceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of ComputeInstance CRD objects
	Items []ComputeInstance `json:"items,omitempty"`
}
