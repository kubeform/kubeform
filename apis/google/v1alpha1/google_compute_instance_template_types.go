package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type GoogleComputeInstanceTemplate struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              GoogleComputeInstanceTemplateSpec   `json:"spec,omitempty"`
	Status            GoogleComputeInstanceTemplateStatus `json:"status,omitempty"`
}

type GoogleComputeInstanceTemplateSpecGuestAccelerator struct {
	Count int    `json:"count"`
	Type  string `json:"type"`
}

type GoogleComputeInstanceTemplateSpecNetworkInterfaceAccessConfig struct {
	NetworkTier   string `json:"network_tier"`
	AssignedNatIp string `json:"assigned_nat_ip"`
	NatIp         string `json:"nat_ip"`
}

type GoogleComputeInstanceTemplateSpecNetworkInterfaceAliasIpRange struct {
	IpCidrRange         string `json:"ip_cidr_range"`
	SubnetworkRangeName string `json:"subnetwork_range_name"`
}

type GoogleComputeInstanceTemplateSpecNetworkInterface struct {
	NetworkIp         string                                              `json:"network_ip"`
	Subnetwork        string                                              `json:"subnetwork"`
	SubnetworkProject string                                              `json:"subnetwork_project"`
	AccessConfig      []GoogleComputeInstanceTemplateSpecNetworkInterface `json:"access_config"`
	AliasIpRange      []GoogleComputeInstanceTemplateSpecNetworkInterface `json:"alias_ip_range"`
	Network           string                                              `json:"network"`
	Address           string                                              `json:"address"`
}

type GoogleComputeInstanceTemplateSpecDiskDiskEncryptionKey struct {
	KmsKeySelfLink string `json:"kms_key_self_link"`
}

type GoogleComputeInstanceTemplateSpecDisk struct {
	Boot              bool                                    `json:"boot"`
	DeviceName        string                                  `json:"device_name"`
	DiskType          string                                  `json:"disk_type"`
	Mode              string                                  `json:"mode"`
	Type              string                                  `json:"type"`
	DiskEncryptionKey []GoogleComputeInstanceTemplateSpecDisk `json:"disk_encryption_key"`
	AutoDelete        bool                                    `json:"auto_delete"`
	DiskName          string                                  `json:"disk_name"`
	DiskSizeGb        int                                     `json:"disk_size_gb"`
	SourceImage       string                                  `json:"source_image"`
	Interface         string                                  `json:"interface"`
	Source            string                                  `json:"source"`
}

type GoogleComputeInstanceTemplateSpecScheduling struct {
	Preemptible       bool   `json:"preemptible"`
	AutomaticRestart  bool   `json:"automatic_restart"`
	OnHostMaintenance string `json:"on_host_maintenance"`
}

type GoogleComputeInstanceTemplateSpecServiceAccount struct {
	Scopes []string `json:"scopes"`
	Email  string   `json:"email"`
}

type GoogleComputeInstanceTemplateSpec struct {
	GuestAccelerator      []GoogleComputeInstanceTemplateSpec `json:"guest_accelerator"`
	Tags                  []string                            `json:"tags"`
	TagsFingerprint       string                              `json:"tags_fingerprint"`
	Name                  string                              `json:"name"`
	NetworkInterface      []GoogleComputeInstanceTemplateSpec `json:"network_interface"`
	InstanceDescription   string                              `json:"instance_description"`
	OnHostMaintenance     string                              `json:"on_host_maintenance"`
	Region                string                              `json:"region"`
	SelfLink              string                              `json:"self_link"`
	MachineType           string                              `json:"machine_type"`
	CanIpForward          bool                                `json:"can_ip_forward"`
	MetadataStartupScript string                              `json:"metadata_startup_script"`
	Project               string                              `json:"project"`
	NamePrefix            string                              `json:"name_prefix"`
	Disk                  []GoogleComputeInstanceTemplateSpec `json:"disk"`
	Metadata              map[string]string                   `json:"metadata"`
	MetadataFingerprint   string                              `json:"metadata_fingerprint"`
	Scheduling            []GoogleComputeInstanceTemplateSpec `json:"scheduling"`
	ServiceAccount        []GoogleComputeInstanceTemplateSpec `json:"service_account"`
	MinCpuPlatform        string                              `json:"min_cpu_platform"`
	Labels                map[string]string                   `json:"labels"`
	AutomaticRestart      bool                                `json:"automatic_restart"`
	Description           string                              `json:"description"`
}

type GoogleComputeInstanceTemplateStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// GoogleComputeInstanceTemplateList is a list of GoogleComputeInstanceTemplates
type GoogleComputeInstanceTemplateList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of GoogleComputeInstanceTemplate CRD objects
	Items []GoogleComputeInstanceTemplate `json:"items,omitempty"`
}
