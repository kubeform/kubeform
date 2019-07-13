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

type GoogleComputeInstanceSpecScratchDisk struct {
	Interface string `json:"interface"`
}

type GoogleComputeInstanceSpecBootDiskInitializeParams struct {
	Size  int    `json:"size"`
	Type  string `json:"type"`
	Image string `json:"image"`
}

type GoogleComputeInstanceSpecBootDisk struct {
	AutoDelete              bool                                `json:"auto_delete"`
	DeviceName              string                              `json:"device_name"`
	DiskEncryptionKeyRaw    string                              `json:"disk_encryption_key_raw"`
	DiskEncryptionKeySha256 string                              `json:"disk_encryption_key_sha256"`
	InitializeParams        []GoogleComputeInstanceSpecBootDisk `json:"initialize_params"`
	Source                  string                              `json:"source"`
}

type GoogleComputeInstanceSpecAttachedDisk struct {
	Source                  string `json:"source"`
	DeviceName              string `json:"device_name"`
	Mode                    string `json:"mode"`
	DiskEncryptionKeyRaw    string `json:"disk_encryption_key_raw"`
	DiskEncryptionKeySha256 string `json:"disk_encryption_key_sha256"`
}

type GoogleComputeInstanceSpecServiceAccount struct {
	Email  string   `json:"email"`
	Scopes []string `json:"scopes"`
}

type GoogleComputeInstanceSpecGuestAccelerator struct {
	Count int    `json:"count"`
	Type  string `json:"type"`
}

type GoogleComputeInstanceSpecDisk struct {
	DiskEncryptionKeySha256 string `json:"disk_encryption_key_sha256"`
	Disk                    string `json:"disk"`
	Image                   string `json:"image"`
	Type                    string `json:"type"`
	Scratch                 bool   `json:"scratch"`
	Size                    int    `json:"size"`
	DeviceName              string `json:"device_name"`
	DiskEncryptionKeyRaw    string `json:"disk_encryption_key_raw"`
	AutoDelete              bool   `json:"auto_delete"`
}

type GoogleComputeInstanceSpecScheduling struct {
	Preemptible       bool   `json:"preemptible"`
	OnHostMaintenance string `json:"on_host_maintenance"`
	AutomaticRestart  bool   `json:"automatic_restart"`
}

type GoogleComputeInstanceSpecNetworkInterfaceAccessConfig struct {
	NatIp               string `json:"nat_ip"`
	NetworkTier         string `json:"network_tier"`
	AssignedNatIp       string `json:"assigned_nat_ip"`
	PublicPtrDomainName string `json:"public_ptr_domain_name"`
}

type GoogleComputeInstanceSpecNetworkInterfaceAliasIpRange struct {
	IpCidrRange         string `json:"ip_cidr_range"`
	SubnetworkRangeName string `json:"subnetwork_range_name"`
}

type GoogleComputeInstanceSpecNetworkInterface struct {
	Address           string                                      `json:"address"`
	NetworkIp         string                                      `json:"network_ip"`
	AccessConfig      []GoogleComputeInstanceSpecNetworkInterface `json:"access_config"`
	AliasIpRange      []GoogleComputeInstanceSpecNetworkInterface `json:"alias_ip_range"`
	Network           string                                      `json:"network"`
	Subnetwork        string                                      `json:"subnetwork"`
	SubnetworkProject string                                      `json:"subnetwork_project"`
	Name              string                                      `json:"name"`
}

type GoogleComputeInstanceSpecNetwork struct {
	Source          string `json:"source"`
	Address         string `json:"address"`
	Name            string `json:"name"`
	InternalAddress string `json:"internal_address"`
	ExternalAddress string `json:"external_address"`
}

type GoogleComputeInstanceSpec struct {
	Labels                 map[string]string           `json:"labels"`
	ScratchDisk            []GoogleComputeInstanceSpec `json:"scratch_disk"`
	Tags                   []string                    `json:"tags"`
	CpuPlatform            string                      `json:"cpu_platform"`
	SelfLink               string                      `json:"self_link"`
	BootDisk               []GoogleComputeInstanceSpec `json:"boot_disk"`
	AttachedDisk           []GoogleComputeInstanceSpec `json:"attached_disk"`
	Description            string                      `json:"description"`
	DeletionProtection     bool                        `json:"deletion_protection"`
	MinCpuPlatform         string                      `json:"min_cpu_platform"`
	InstanceId             string                      `json:"instance_id"`
	MachineType            string                      `json:"machine_type"`
	CreateTimeout          int                         `json:"create_timeout"`
	ServiceAccount         []GoogleComputeInstanceSpec `json:"service_account"`
	GuestAccelerator       []GoogleComputeInstanceSpec `json:"guest_accelerator"`
	Metadata               map[string]string           `json:"metadata"`
	MetadataStartupScript  string                      `json:"metadata_startup_script"`
	MetadataFingerprint    string                      `json:"metadata_fingerprint"`
	Disk                   []GoogleComputeInstanceSpec `json:"disk"`
	Scheduling             []GoogleComputeInstanceSpec `json:"scheduling"`
	TagsFingerprint        string                      `json:"tags_fingerprint"`
	NetworkInterface       []GoogleComputeInstanceSpec `json:"network_interface"`
	Network                []GoogleComputeInstanceSpec `json:"network"`
	Zone                   string                      `json:"zone"`
	Name                   string                      `json:"name"`
	AllowStoppingForUpdate bool                        `json:"allow_stopping_for_update"`
	CanIpForward           bool                        `json:"can_ip_forward"`
	Project                string                      `json:"project"`
	LabelFingerprint       string                      `json:"label_fingerprint"`
}



type GoogleComputeInstanceStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

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