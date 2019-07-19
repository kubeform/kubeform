package v1alpha1

import (
	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
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
	DeviceName string `json:"deviceName,omitempty" tf:"device_name,omitempty"`
	// +optional
	// Sensitive Data. Provide secret name which contains one value only
	DiskEncryptionKeyRaw core.LocalObjectReference `json:"diskEncryptionKeyRaw,omitempty" tf:"disk_encryption_key_raw,omitempty"`
	// +optional
	Mode   string `json:"mode,omitempty" tf:"mode,omitempty"`
	Source string `json:"source" tf:"source"`
}

type ComputeInstanceSpecBootDiskInitializeParams struct {
	// +optional
	Image string `json:"image,omitempty" tf:"image,omitempty"`
	// +optional
	Size int `json:"size,omitempty" tf:"size,omitempty"`
	// +optional
	Type string `json:"type,omitempty" tf:"type,omitempty"`
}

type ComputeInstanceSpecBootDisk struct {
	// +optional
	AutoDelete bool `json:"autoDelete,omitempty" tf:"auto_delete,omitempty"`
	// +optional
	DeviceName string `json:"deviceName,omitempty" tf:"device_name,omitempty"`
	// +optional
	// Sensitive Data. Provide secret name which contains one value only
	DiskEncryptionKeyRaw core.LocalObjectReference `json:"diskEncryptionKeyRaw,omitempty" tf:"disk_encryption_key_raw,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	InitializeParams []ComputeInstanceSpecBootDiskInitializeParams `json:"initializeParams,omitempty" tf:"initialize_params,omitempty"`
	// +optional
	Source string `json:"source,omitempty" tf:"source,omitempty"`
}

type ComputeInstanceSpecGuestAccelerator struct {
	Count int    `json:"count" tf:"count"`
	Type  string `json:"type" tf:"type"`
}

type ComputeInstanceSpecNetworkInterfaceAccessConfig struct {
	// +optional
	NatIP string `json:"natIP,omitempty" tf:"nat_ip,omitempty"`
	// +optional
	NetworkTier string `json:"networkTier,omitempty" tf:"network_tier,omitempty"`
	// +optional
	PublicPtrDomainName string `json:"publicPtrDomainName,omitempty" tf:"public_ptr_domain_name,omitempty"`
}

type ComputeInstanceSpecNetworkInterfaceAliasIPRange struct {
	IpCIDRRange string `json:"ipCIDRRange" tf:"ip_cidr_range"`
	// +optional
	SubnetworkRangeName string `json:"subnetworkRangeName,omitempty" tf:"subnetwork_range_name,omitempty"`
}

type ComputeInstanceSpecNetworkInterface struct {
	// +optional
	AccessConfig []ComputeInstanceSpecNetworkInterfaceAccessConfig `json:"accessConfig,omitempty" tf:"access_config,omitempty"`
	// +optional
	// Deprecated
	Address string `json:"address,omitempty" tf:"address,omitempty"`
	// +optional
	AliasIPRange []ComputeInstanceSpecNetworkInterfaceAliasIPRange `json:"aliasIPRange,omitempty" tf:"alias_ip_range,omitempty"`
	// +optional
	Network string `json:"network,omitempty" tf:"network,omitempty"`
	// +optional
	NetworkIP string `json:"networkIP,omitempty" tf:"network_ip,omitempty"`
	// +optional
	Subnetwork string `json:"subnetwork,omitempty" tf:"subnetwork,omitempty"`
	// +optional
	SubnetworkProject string `json:"subnetworkProject,omitempty" tf:"subnetwork_project,omitempty"`
}

type ComputeInstanceSpecScheduling struct {
	// +optional
	AutomaticRestart bool `json:"automaticRestart,omitempty" tf:"automatic_restart,omitempty"`
	// +optional
	OnHostMaintenance string `json:"onHostMaintenance,omitempty" tf:"on_host_maintenance,omitempty"`
	// +optional
	Preemptible bool `json:"preemptible,omitempty" tf:"preemptible,omitempty"`
}

type ComputeInstanceSpecScratchDisk struct {
	// +optional
	Interface string `json:"interface,omitempty" tf:"interface,omitempty"`
}

type ComputeInstanceSpecServiceAccount struct {
	// +optional
	Email string `json:"email,omitempty" tf:"email,omitempty"`
	// +kubebuilder:validation:UniqueItems=true
	Scopes []string `json:"scopes" tf:"scopes"`
}

type ComputeInstanceSpec struct {
	// +optional
	AllowStoppingForUpdate bool `json:"allowStoppingForUpdate,omitempty" tf:"allow_stopping_for_update,omitempty"`
	// +optional
	AttachedDisk []ComputeInstanceSpecAttachedDisk `json:"attachedDisk,omitempty" tf:"attached_disk,omitempty"`
	// +kubebuilder:validation:MaxItems=1
	BootDisk []ComputeInstanceSpecBootDisk `json:"bootDisk" tf:"boot_disk"`
	// +optional
	CanIPForward bool `json:"canIPForward,omitempty" tf:"can_ip_forward,omitempty"`
	// +optional
	// Deprecated
	CreateTimeout int `json:"createTimeout,omitempty" tf:"create_timeout,omitempty"`
	// +optional
	DeletionProtection bool `json:"deletionProtection,omitempty" tf:"deletion_protection,omitempty"`
	// +optional
	Description string `json:"description,omitempty" tf:"description,omitempty"`
	// +optional
	GuestAccelerator []ComputeInstanceSpecGuestAccelerator `json:"guestAccelerator,omitempty" tf:"guest_accelerator,omitempty"`
	// +optional
	Labels      map[string]string `json:"labels,omitempty" tf:"labels,omitempty"`
	MachineType string            `json:"machineType" tf:"machine_type"`
	// +optional
	Metadata map[string]string `json:"metadata,omitempty" tf:"metadata,omitempty"`
	// +optional
	MetadataStartupScript string `json:"metadataStartupScript,omitempty" tf:"metadata_startup_script,omitempty"`
	// +optional
	MinCPUPlatform   string                                `json:"minCPUPlatform,omitempty" tf:"min_cpu_platform,omitempty"`
	Name             string                                `json:"name" tf:"name"`
	NetworkInterface []ComputeInstanceSpecNetworkInterface `json:"networkInterface" tf:"network_interface"`
	// +optional
	Project string `json:"project,omitempty" tf:"project,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	Scheduling []ComputeInstanceSpecScheduling `json:"scheduling,omitempty" tf:"scheduling,omitempty"`
	// +optional
	ScratchDisk []ComputeInstanceSpecScratchDisk `json:"scratchDisk,omitempty" tf:"scratch_disk,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	ServiceAccount []ComputeInstanceSpecServiceAccount `json:"serviceAccount,omitempty" tf:"service_account,omitempty"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	Tags []string `json:"tags,omitempty" tf:"tags,omitempty"`
	// +optional
	Zone        string                    `json:"zone,omitempty" tf:"zone,omitempty"`
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`
}

type ComputeInstanceStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	TFState *runtime.RawExtension `json:"tfState,omitempty"`
	Output  *runtime.RawExtension `json:"output,omitempty"`
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
