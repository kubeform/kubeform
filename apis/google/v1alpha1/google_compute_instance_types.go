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

type GoogleComputeInstance struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              GoogleComputeInstanceSpec   `json:"spec,omitempty"`
	Status            GoogleComputeInstanceStatus `json:"status,omitempty"`
}

type GoogleComputeInstanceSpecNetwork struct {
	ExternalAddress string `json:"external_address"`
	Source          string `json:"source"`
	Address         string `json:"address"`
	Name            string `json:"name"`
	InternalAddress string `json:"internal_address"`
}

type GoogleComputeInstanceSpecScratchDisk struct {
	Interface string `json:"interface"`
}

type GoogleComputeInstanceSpecDisk struct {
	Type                    string `json:"type"`
	Scratch                 bool   `json:"scratch"`
	AutoDelete              bool   `json:"auto_delete"`
	Size                    int    `json:"size"`
	DiskEncryptionKeySha256 string `json:"disk_encryption_key_sha256"`
	Disk                    string `json:"disk"`
	DeviceName              string `json:"device_name"`
	DiskEncryptionKeyRaw    string `json:"disk_encryption_key_raw"`
	Image                   string `json:"image"`
}

type GoogleComputeInstanceSpecServiceAccount struct {
	Email  string   `json:"email"`
	Scopes []string `json:"scopes"`
}

type GoogleComputeInstanceSpecAttachedDisk struct {
	Mode                    string `json:"mode"`
	DiskEncryptionKeyRaw    string `json:"disk_encryption_key_raw"`
	DiskEncryptionKeySha256 string `json:"disk_encryption_key_sha256"`
	Source                  string `json:"source"`
	DeviceName              string `json:"device_name"`
}

type GoogleComputeInstanceSpecGuestAccelerator struct {
	Count int    `json:"count"`
	Type  string `json:"type"`
}

type GoogleComputeInstanceSpecScheduling struct {
	OnHostMaintenance string `json:"on_host_maintenance"`
	AutomaticRestart  bool   `json:"automatic_restart"`
	Preemptible       bool   `json:"preemptible"`
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

type GoogleComputeInstanceSpecBootDiskInitializeParams struct {
	Size  int    `json:"size"`
	Type  string `json:"type"`
	Image string `json:"image"`
}

type GoogleComputeInstanceSpecBootDisk struct {
	DiskEncryptionKeyRaw    string                              `json:"disk_encryption_key_raw"`
	DiskEncryptionKeySha256 string                              `json:"disk_encryption_key_sha256"`
	InitializeParams        []GoogleComputeInstanceSpecBootDisk `json:"initialize_params"`
	Source                  string                              `json:"source"`
	AutoDelete              bool                                `json:"auto_delete"`
	DeviceName              string                              `json:"device_name"`
}

type GoogleComputeInstanceSpec struct {
	Description            string                      `json:"description"`
	Tags                   []string                    `json:"tags"`
	AllowStoppingForUpdate bool                        `json:"allow_stopping_for_update"`
	Network                []GoogleComputeInstanceSpec `json:"network"`
	ScratchDisk            []GoogleComputeInstanceSpec `json:"scratch_disk"`
	Zone                   string                      `json:"zone"`
	InstanceId             string                      `json:"instance_id"`
	SelfLink               string                      `json:"self_link"`
	Name                   string                      `json:"name"`
	MetadataFingerprint    string                      `json:"metadata_fingerprint"`
	CreateTimeout          int                         `json:"create_timeout"`
	DeletionProtection     bool                        `json:"deletion_protection"`
	Disk                   []GoogleComputeInstanceSpec `json:"disk"`
	MetadataStartupScript  string                      `json:"metadata_startup_script"`
	ServiceAccount         []GoogleComputeInstanceSpec `json:"service_account"`
	MachineType            string                      `json:"machine_type"`
	AttachedDisk           []GoogleComputeInstanceSpec `json:"attached_disk"`
	GuestAccelerator       []GoogleComputeInstanceSpec `json:"guest_accelerator"`
	CanIpForward           bool                        `json:"can_ip_forward"`
	Metadata               map[string]string           `json:"metadata"`
	MinCpuPlatform         string                      `json:"min_cpu_platform"`
	Scheduling             []GoogleComputeInstanceSpec `json:"scheduling"`
	NetworkInterface       []GoogleComputeInstanceSpec `json:"network_interface"`
	BootDisk               []GoogleComputeInstanceSpec `json:"boot_disk"`
	Labels                 map[string]string           `json:"labels"`
	Project                string                      `json:"project"`
	CpuPlatform            string                      `json:"cpu_platform"`
	LabelFingerprint       string                      `json:"label_fingerprint"`
	TagsFingerprint        string                      `json:"tags_fingerprint"`
}

type GoogleComputeInstanceStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// GoogleComputeInstanceList is a list of GoogleComputeInstances
type GoogleComputeInstanceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of GoogleComputeInstance CRD objects
	Items []GoogleComputeInstance `json:"items,omitempty"`
}
