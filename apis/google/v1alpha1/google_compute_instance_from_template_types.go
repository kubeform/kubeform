package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type GoogleComputeInstanceFromTemplate struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              GoogleComputeInstanceFromTemplateSpec   `json:"spec,omitempty"`
	Status            GoogleComputeInstanceFromTemplateStatus `json:"status,omitempty"`
}

type GoogleComputeInstanceFromTemplateSpecScheduling struct {
	Preemptible       bool   `json:"preemptible"`
	OnHostMaintenance string `json:"on_host_maintenance"`
	AutomaticRestart  bool   `json:"automatic_restart"`
}

type GoogleComputeInstanceFromTemplateSpecAttachedDisk struct {
	Mode                    string `json:"mode"`
	DiskEncryptionKeyRaw    string `json:"disk_encryption_key_raw"`
	DiskEncryptionKeySha256 string `json:"disk_encryption_key_sha256"`
	Source                  string `json:"source"`
	DeviceName              string `json:"device_name"`
}

type GoogleComputeInstanceFromTemplateSpecGuestAccelerator struct {
	Count int    `json:"count"`
	Type  string `json:"type"`
}

type GoogleComputeInstanceFromTemplateSpecScratchDisk struct {
	Interface string `json:"interface"`
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
	Source                  string                                          `json:"source"`
	AutoDelete              bool                                            `json:"auto_delete"`
	DeviceName              string                                          `json:"device_name"`
	DiskEncryptionKeyRaw    string                                          `json:"disk_encryption_key_raw"`
	DiskEncryptionKeySha256 string                                          `json:"disk_encryption_key_sha256"`
	InitializeParams        []GoogleComputeInstanceFromTemplateSpecBootDisk `json:"initialize_params"`
}

type GoogleComputeInstanceFromTemplateSpecNetworkInterfaceAccessConfig struct {
	NetworkTier         string `json:"network_tier"`
	AssignedNatIp       string `json:"assigned_nat_ip"`
	PublicPtrDomainName string `json:"public_ptr_domain_name"`
	NatIp               string `json:"nat_ip"`
}

type GoogleComputeInstanceFromTemplateSpecNetworkInterfaceAliasIpRange struct {
	IpCidrRange         string `json:"ip_cidr_range"`
	SubnetworkRangeName string `json:"subnetwork_range_name"`
}

type GoogleComputeInstanceFromTemplateSpecNetworkInterface struct {
	NetworkIp         string                                                  `json:"network_ip"`
	AccessConfig      []GoogleComputeInstanceFromTemplateSpecNetworkInterface `json:"access_config"`
	AliasIpRange      []GoogleComputeInstanceFromTemplateSpecNetworkInterface `json:"alias_ip_range"`
	Network           string                                                  `json:"network"`
	Subnetwork        string                                                  `json:"subnetwork"`
	SubnetworkProject string                                                  `json:"subnetwork_project"`
	Name              string                                                  `json:"name"`
	Address           string                                                  `json:"address"`
}

type GoogleComputeInstanceFromTemplateSpec struct {
	Scheduling             []GoogleComputeInstanceFromTemplateSpec `json:"scheduling"`
	Zone                   string                                  `json:"zone"`
	LabelFingerprint       string                                  `json:"label_fingerprint"`
	TagsFingerprint        string                                  `json:"tags_fingerprint"`
	AttachedDisk           []GoogleComputeInstanceFromTemplateSpec `json:"attached_disk"`
	Description            string                                  `json:"description"`
	GuestAccelerator       []GoogleComputeInstanceFromTemplateSpec `json:"guest_accelerator"`
	MetadataStartupScript  string                                  `json:"metadata_startup_script"`
	ScratchDisk            []GoogleComputeInstanceFromTemplateSpec `json:"scratch_disk"`
	ServiceAccount         []GoogleComputeInstanceFromTemplateSpec `json:"service_account"`
	AllowStoppingForUpdate bool                                    `json:"allow_stopping_for_update"`
	Labels                 map[string]string                       `json:"labels"`
	Project                string                                  `json:"project"`
	BootDisk               []GoogleComputeInstanceFromTemplateSpec `json:"boot_disk"`
	MachineType            string                                  `json:"machine_type"`
	DeletionProtection     bool                                    `json:"deletion_protection"`
	Metadata               map[string]string                       `json:"metadata"`
	Tags                   []string                                `json:"tags"`
	MetadataFingerprint    string                                  `json:"metadata_fingerprint"`
	SourceInstanceTemplate string                                  `json:"source_instance_template"`
	InstanceId             string                                  `json:"instance_id"`
	Name                   string                                  `json:"name"`
	NetworkInterface       []GoogleComputeInstanceFromTemplateSpec `json:"network_interface"`
	CpuPlatform            string                                  `json:"cpu_platform"`
	SelfLink               string                                  `json:"self_link"`
	CanIpForward           bool                                    `json:"can_ip_forward"`
	MinCpuPlatform         string                                  `json:"min_cpu_platform"`
}



type GoogleComputeInstanceFromTemplateStatus struct {
    Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// GoogleComputeInstanceFromTemplateList is a list of GoogleComputeInstanceFromTemplates
type GoogleComputeInstanceFromTemplateList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of GoogleComputeInstanceFromTemplate CRD objects
	Items []GoogleComputeInstanceFromTemplate `json:"items,omitempty"`
}