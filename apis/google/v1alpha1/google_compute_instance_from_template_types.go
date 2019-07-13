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

type GoogleComputeInstanceFromTemplate struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              GoogleComputeInstanceFromTemplateSpec   `json:"spec,omitempty"`
	Status            GoogleComputeInstanceFromTemplateStatus `json:"status,omitempty"`
}

type GoogleComputeInstanceFromTemplateSpecScheduling struct {
	OnHostMaintenance string `json:"on_host_maintenance"`
	AutomaticRestart  bool   `json:"automatic_restart"`
	Preemptible       bool   `json:"preemptible"`
}

type GoogleComputeInstanceFromTemplateSpecNetworkInterfaceAccessConfig struct {
	AssignedNatIp       string `json:"assigned_nat_ip"`
	PublicPtrDomainName string `json:"public_ptr_domain_name"`
	NatIp               string `json:"nat_ip"`
	NetworkTier         string `json:"network_tier"`
}

type GoogleComputeInstanceFromTemplateSpecNetworkInterfaceAliasIpRange struct {
	IpCidrRange         string `json:"ip_cidr_range"`
	SubnetworkRangeName string `json:"subnetwork_range_name"`
}

type GoogleComputeInstanceFromTemplateSpecNetworkInterface struct {
	Subnetwork        string                                                  `json:"subnetwork"`
	SubnetworkProject string                                                  `json:"subnetwork_project"`
	Name              string                                                  `json:"name"`
	Address           string                                                  `json:"address"`
	NetworkIp         string                                                  `json:"network_ip"`
	AccessConfig      []GoogleComputeInstanceFromTemplateSpecNetworkInterface `json:"access_config"`
	AliasIpRange      []GoogleComputeInstanceFromTemplateSpecNetworkInterface `json:"alias_ip_range"`
	Network           string                                                  `json:"network"`
}

type GoogleComputeInstanceFromTemplateSpecAttachedDisk struct {
	Source                  string `json:"source"`
	DeviceName              string `json:"device_name"`
	Mode                    string `json:"mode"`
	DiskEncryptionKeyRaw    string `json:"disk_encryption_key_raw"`
	DiskEncryptionKeySha256 string `json:"disk_encryption_key_sha256"`
}

type GoogleComputeInstanceFromTemplateSpecServiceAccount struct {
	Email  string   `json:"email"`
	Scopes []string `json:"scopes"`
}

type GoogleComputeInstanceFromTemplateSpecBootDiskInitializeParams struct {
	Size  int    `json:"size"`
	Type  string `json:"type"`
	Image string `json:"image"`
}

type GoogleComputeInstanceFromTemplateSpecBootDisk struct {
	AutoDelete              bool                                            `json:"auto_delete"`
	DeviceName              string                                          `json:"device_name"`
	DiskEncryptionKeyRaw    string                                          `json:"disk_encryption_key_raw"`
	DiskEncryptionKeySha256 string                                          `json:"disk_encryption_key_sha256"`
	InitializeParams        []GoogleComputeInstanceFromTemplateSpecBootDisk `json:"initialize_params"`
	Source                  string                                          `json:"source"`
}

type GoogleComputeInstanceFromTemplateSpecGuestAccelerator struct {
	Count int    `json:"count"`
	Type  string `json:"type"`
}

type GoogleComputeInstanceFromTemplateSpecScratchDisk struct {
	Interface string `json:"interface"`
}

type GoogleComputeInstanceFromTemplateSpec struct {
	Scheduling             []GoogleComputeInstanceFromTemplateSpec `json:"scheduling"`
	Tags                   []string                                `json:"tags"`
	CpuPlatform            string                                  `json:"cpu_platform"`
	NetworkInterface       []GoogleComputeInstanceFromTemplateSpec `json:"network_interface"`
	MetadataStartupScript  string                                  `json:"metadata_startup_script"`
	InstanceId             string                                  `json:"instance_id"`
	LabelFingerprint       string                                  `json:"label_fingerprint"`
	MetadataFingerprint    string                                  `json:"metadata_fingerprint"`
	DeletionProtection     bool                                    `json:"deletion_protection"`
	AttachedDisk           []GoogleComputeInstanceFromTemplateSpec `json:"attached_disk"`
	Metadata               map[string]string                       `json:"metadata"`
	Project                string                                  `json:"project"`
	ServiceAccount         []GoogleComputeInstanceFromTemplateSpec `json:"service_account"`
	MachineType            string                                  `json:"machine_type"`
	SelfLink               string                                  `json:"self_link"`
	AllowStoppingForUpdate bool                                    `json:"allow_stopping_for_update"`
	CanIpForward           bool                                    `json:"can_ip_forward"`
	Labels                 map[string]string                       `json:"labels"`
	MinCpuPlatform         string                                  `json:"min_cpu_platform"`
	BootDisk               []GoogleComputeInstanceFromTemplateSpec `json:"boot_disk"`
	Zone                   string                                  `json:"zone"`
	TagsFingerprint        string                                  `json:"tags_fingerprint"`
	Name                   string                                  `json:"name"`
	GuestAccelerator       []GoogleComputeInstanceFromTemplateSpec `json:"guest_accelerator"`
	ScratchDisk            []GoogleComputeInstanceFromTemplateSpec `json:"scratch_disk"`
	SourceInstanceTemplate string                                  `json:"source_instance_template"`
	Description            string                                  `json:"description"`
}



type GoogleComputeInstanceFromTemplateStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// GoogleComputeInstanceFromTemplateList is a list of GoogleComputeInstanceFromTemplates
type GoogleComputeInstanceFromTemplateList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of GoogleComputeInstanceFromTemplate CRD objects
	Items []GoogleComputeInstanceFromTemplate `json:"items,omitempty"`
}