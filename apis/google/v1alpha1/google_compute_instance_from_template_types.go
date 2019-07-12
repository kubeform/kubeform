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

type GoogleComputeInstanceFromTemplateSpecServiceAccount struct {
	Email  string   `json:"email"`
	Scopes []string `json:"scopes"`
}

type GoogleComputeInstanceFromTemplateSpecGuestAccelerator struct {
	Type  string `json:"type"`
	Count int    `json:"count"`
}

type GoogleComputeInstanceFromTemplateSpecBootDiskInitializeParams struct {
	Type  string `json:"type"`
	Image string `json:"image"`
	Size  int    `json:"size"`
}

type GoogleComputeInstanceFromTemplateSpecBootDisk struct {
	DiskEncryptionKeyRaw    string                                          `json:"disk_encryption_key_raw"`
	DiskEncryptionKeySha256 string                                          `json:"disk_encryption_key_sha256"`
	InitializeParams        []GoogleComputeInstanceFromTemplateSpecBootDisk `json:"initialize_params"`
	Source                  string                                          `json:"source"`
	AutoDelete              bool                                            `json:"auto_delete"`
	DeviceName              string                                          `json:"device_name"`
}

type GoogleComputeInstanceFromTemplateSpecAttachedDisk struct {
	Source                  string `json:"source"`
	DeviceName              string `json:"device_name"`
	Mode                    string `json:"mode"`
	DiskEncryptionKeyRaw    string `json:"disk_encryption_key_raw"`
	DiskEncryptionKeySha256 string `json:"disk_encryption_key_sha256"`
}

type GoogleComputeInstanceFromTemplateSpecScheduling struct {
	OnHostMaintenance string `json:"on_host_maintenance"`
	AutomaticRestart  bool   `json:"automatic_restart"`
	Preemptible       bool   `json:"preemptible"`
}

type GoogleComputeInstanceFromTemplateSpecScratchDisk struct {
	Interface string `json:"interface"`
}

type GoogleComputeInstanceFromTemplateSpecNetworkInterfaceAccessConfig struct {
	NatIp               string `json:"nat_ip"`
	NetworkTier         string `json:"network_tier"`
	AssignedNatIp       string `json:"assigned_nat_ip"`
	PublicPtrDomainName string `json:"public_ptr_domain_name"`
}

type GoogleComputeInstanceFromTemplateSpecNetworkInterfaceAliasIpRange struct {
	IpCidrRange         string `json:"ip_cidr_range"`
	SubnetworkRangeName string `json:"subnetwork_range_name"`
}

type GoogleComputeInstanceFromTemplateSpecNetworkInterface struct {
	Network           string                                                  `json:"network"`
	Subnetwork        string                                                  `json:"subnetwork"`
	SubnetworkProject string                                                  `json:"subnetwork_project"`
	Name              string                                                  `json:"name"`
	Address           string                                                  `json:"address"`
	NetworkIp         string                                                  `json:"network_ip"`
	AccessConfig      []GoogleComputeInstanceFromTemplateSpecNetworkInterface `json:"access_config"`
	AliasIpRange      []GoogleComputeInstanceFromTemplateSpecNetworkInterface `json:"alias_ip_range"`
}

type GoogleComputeInstanceFromTemplateSpec struct {
	CpuPlatform            string                                  `json:"cpu_platform"`
	MetadataFingerprint    string                                  `json:"metadata_fingerprint"`
	Tags                   []string                                `json:"tags"`
	SelfLink               string                                  `json:"self_link"`
	Project                string                                  `json:"project"`
	ServiceAccount         []GoogleComputeInstanceFromTemplateSpec `json:"service_account"`
	CanIpForward           bool                                    `json:"can_ip_forward"`
	DeletionProtection     bool                                    `json:"deletion_protection"`
	Labels                 map[string]string                       `json:"labels"`
	MinCpuPlatform         string                                  `json:"min_cpu_platform"`
	AllowStoppingForUpdate bool                                    `json:"allow_stopping_for_update"`
	LabelFingerprint       string                                  `json:"label_fingerprint"`
	MachineType            string                                  `json:"machine_type"`
	GuestAccelerator       []GoogleComputeInstanceFromTemplateSpec `json:"guest_accelerator"`
	MetadataStartupScript  string                                  `json:"metadata_startup_script"`
	Zone                   string                                  `json:"zone"`
	SourceInstanceTemplate string                                  `json:"source_instance_template"`
	BootDisk               []GoogleComputeInstanceFromTemplateSpec `json:"boot_disk"`
	AttachedDisk           []GoogleComputeInstanceFromTemplateSpec `json:"attached_disk"`
	Metadata               map[string]string                       `json:"metadata"`
	Scheduling             []GoogleComputeInstanceFromTemplateSpec `json:"scheduling"`
	ScratchDisk            []GoogleComputeInstanceFromTemplateSpec `json:"scratch_disk"`
	Name                   string                                  `json:"name"`
	NetworkInterface       []GoogleComputeInstanceFromTemplateSpec `json:"network_interface"`
	Description            string                                  `json:"description"`
	InstanceId             string                                  `json:"instance_id"`
	TagsFingerprint        string                                  `json:"tags_fingerprint"`
}

type GoogleComputeInstanceFromTemplateStatus struct {
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
