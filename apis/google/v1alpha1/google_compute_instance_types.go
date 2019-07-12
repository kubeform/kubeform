package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type GoogleComputeInstance struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              GoogleComputeInstanceSpec   `json:"spec,omitempty"`
	Status            GoogleComputeInstanceStatus `json:"status,omitempty"`
}

type GoogleComputeInstanceSpecServiceAccount struct {
	Scopes []string `json:"scopes"`
	Email  string   `json:"email"`
}

type GoogleComputeInstanceSpecBootDiskInitializeParams struct {
	Type  string `json:"type"`
	Image string `json:"image"`
	Size  int    `json:"size"`
}

type GoogleComputeInstanceSpecBootDisk struct {
	DiskEncryptionKeyRaw    string                              `json:"disk_encryption_key_raw"`
	DiskEncryptionKeySha256 string                              `json:"disk_encryption_key_sha256"`
	InitializeParams        []GoogleComputeInstanceSpecBootDisk `json:"initialize_params"`
	Source                  string                              `json:"source"`
	AutoDelete              bool                                `json:"auto_delete"`
	DeviceName              string                              `json:"device_name"`
}

type GoogleComputeInstanceSpecNetwork struct {
	ExternalAddress string `json:"external_address"`
	Source          string `json:"source"`
	Address         string `json:"address"`
	Name            string `json:"name"`
	InternalAddress string `json:"internal_address"`
}

type GoogleComputeInstanceSpecScheduling struct {
	OnHostMaintenance string `json:"on_host_maintenance"`
	AutomaticRestart  bool   `json:"automatic_restart"`
	Preemptible       bool   `json:"preemptible"`
}

type GoogleComputeInstanceSpecScratchDisk struct {
	Interface string `json:"interface"`
}

type GoogleComputeInstanceSpecGuestAccelerator struct {
	Count int    `json:"count"`
	Type  string `json:"type"`
}

type GoogleComputeInstanceSpecDisk struct {
	AutoDelete              bool   `json:"auto_delete"`
	Size                    int    `json:"size"`
	Image                   string `json:"image"`
	Type                    string `json:"type"`
	DeviceName              string `json:"device_name"`
	DiskEncryptionKeyRaw    string `json:"disk_encryption_key_raw"`
	DiskEncryptionKeySha256 string `json:"disk_encryption_key_sha256"`
	Disk                    string `json:"disk"`
	Scratch                 bool   `json:"scratch"`
}

type GoogleComputeInstanceSpecNetworkInterfaceAliasIpRange struct {
	IpCidrRange         string `json:"ip_cidr_range"`
	SubnetworkRangeName string `json:"subnetwork_range_name"`
}

type GoogleComputeInstanceSpecNetworkInterfaceAccessConfig struct {
	NetworkTier         string `json:"network_tier"`
	AssignedNatIp       string `json:"assigned_nat_ip"`
	PublicPtrDomainName string `json:"public_ptr_domain_name"`
	NatIp               string `json:"nat_ip"`
}

type GoogleComputeInstanceSpecNetworkInterface struct {
	AliasIpRange      []GoogleComputeInstanceSpecNetworkInterface `json:"alias_ip_range"`
	Network           string                                      `json:"network"`
	Subnetwork        string                                      `json:"subnetwork"`
	SubnetworkProject string                                      `json:"subnetwork_project"`
	Name              string                                      `json:"name"`
	Address           string                                      `json:"address"`
	NetworkIp         string                                      `json:"network_ip"`
	AccessConfig      []GoogleComputeInstanceSpecNetworkInterface `json:"access_config"`
}

type GoogleComputeInstanceSpecAttachedDisk struct {
	DiskEncryptionKeySha256 string `json:"disk_encryption_key_sha256"`
	Source                  string `json:"source"`
	DeviceName              string `json:"device_name"`
	Mode                    string `json:"mode"`
	DiskEncryptionKeyRaw    string `json:"disk_encryption_key_raw"`
}

type GoogleComputeInstanceSpec struct {
	Name                   string                      `json:"name"`
	DeletionProtection     bool                        `json:"deletion_protection"`
	LabelFingerprint       string                      `json:"label_fingerprint"`
	MachineType            string                      `json:"machine_type"`
	CanIpForward           bool                        `json:"can_ip_forward"`
	ServiceAccount         []GoogleComputeInstanceSpec `json:"service_account"`
	CpuPlatform            string                      `json:"cpu_platform"`
	BootDisk               []GoogleComputeInstanceSpec `json:"boot_disk"`
	CreateTimeout          int                         `json:"create_timeout"`
	Network                []GoogleComputeInstanceSpec `json:"network"`
	Project                string                      `json:"project"`
	Scheduling             []GoogleComputeInstanceSpec `json:"scheduling"`
	ScratchDisk            []GoogleComputeInstanceSpec `json:"scratch_disk"`
	MetadataFingerprint    string                      `json:"metadata_fingerprint"`
	GuestAccelerator       []GoogleComputeInstanceSpec `json:"guest_accelerator"`
	Labels                 map[string]string           `json:"labels"`
	Metadata               map[string]string           `json:"metadata"`
	MinCpuPlatform         string                      `json:"min_cpu_platform"`
	Tags                   []string                    `json:"tags"`
	AllowStoppingForUpdate bool                        `json:"allow_stopping_for_update"`
	Description            string                      `json:"description"`
	Disk                   []GoogleComputeInstanceSpec `json:"disk"`
	MetadataStartupScript  string                      `json:"metadata_startup_script"`
	InstanceId             string                      `json:"instance_id"`
	SelfLink               string                      `json:"self_link"`
	NetworkInterface       []GoogleComputeInstanceSpec `json:"network_interface"`
	Zone                   string                      `json:"zone"`
	TagsFingerprint        string                      `json:"tags_fingerprint"`
	AttachedDisk           []GoogleComputeInstanceSpec `json:"attached_disk"`
}

type GoogleComputeInstanceStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// GoogleComputeInstanceList is a list of GoogleComputeInstances
type GoogleComputeInstanceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of GoogleComputeInstance CRD objects
	Items []GoogleComputeInstance `json:"items,omitempty"`
}
