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

type GoogleComputeInstanceTemplate struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              GoogleComputeInstanceTemplateSpec   `json:"spec,omitempty"`
	Status            GoogleComputeInstanceTemplateStatus `json:"status,omitempty"`
}

type GoogleComputeInstanceTemplateSpecNetworkInterfaceAccessConfig struct {
	AssignedNatIp string `json:"assigned_nat_ip"`
	NatIp         string `json:"nat_ip"`
	NetworkTier   string `json:"network_tier"`
}

type GoogleComputeInstanceTemplateSpecNetworkInterfaceAliasIpRange struct {
	SubnetworkRangeName string `json:"subnetwork_range_name"`
	IpCidrRange         string `json:"ip_cidr_range"`
}

type GoogleComputeInstanceTemplateSpecNetworkInterface struct {
	Network           string                                              `json:"network"`
	Address           string                                              `json:"address"`
	NetworkIp         string                                              `json:"network_ip"`
	Subnetwork        string                                              `json:"subnetwork"`
	SubnetworkProject string                                              `json:"subnetwork_project"`
	AccessConfig      []GoogleComputeInstanceTemplateSpecNetworkInterface `json:"access_config"`
	AliasIpRange      []GoogleComputeInstanceTemplateSpecNetworkInterface `json:"alias_ip_range"`
}

type GoogleComputeInstanceTemplateSpecScheduling struct {
	Preemptible       bool   `json:"preemptible"`
	AutomaticRestart  bool   `json:"automatic_restart"`
	OnHostMaintenance string `json:"on_host_maintenance"`
}

type GoogleComputeInstanceTemplateSpecServiceAccount struct {
	Email  string   `json:"email"`
	Scopes []string `json:"scopes"`
}

type GoogleComputeInstanceTemplateSpecDiskDiskEncryptionKey struct {
	KmsKeySelfLink string `json:"kms_key_self_link"`
}

type GoogleComputeInstanceTemplateSpecDisk struct {
	Source            string                                  `json:"source"`
	Boot              bool                                    `json:"boot"`
	DiskSizeGb        int                                     `json:"disk_size_gb"`
	DiskType          string                                  `json:"disk_type"`
	SourceImage       string                                  `json:"source_image"`
	Mode              string                                  `json:"mode"`
	DiskEncryptionKey []GoogleComputeInstanceTemplateSpecDisk `json:"disk_encryption_key"`
	AutoDelete        bool                                    `json:"auto_delete"`
	DeviceName        string                                  `json:"device_name"`
	DiskName          string                                  `json:"disk_name"`
	Interface         string                                  `json:"interface"`
	Type              string                                  `json:"type"`
}

type GoogleComputeInstanceTemplateSpecGuestAccelerator struct {
	Count int    `json:"count"`
	Type  string `json:"type"`
}

type GoogleComputeInstanceTemplateSpec struct {
	NetworkInterface      []GoogleComputeInstanceTemplateSpec `json:"network_interface"`
	Project               string                              `json:"project"`
	Region                string                              `json:"region"`
	Metadata              map[string]string                   `json:"metadata"`
	MetadataFingerprint   string                              `json:"metadata_fingerprint"`
	MetadataStartupScript string                              `json:"metadata_startup_script"`
	OnHostMaintenance     string                              `json:"on_host_maintenance"`
	Scheduling            []GoogleComputeInstanceTemplateSpec `json:"scheduling"`
	MinCpuPlatform        string                              `json:"min_cpu_platform"`
	Tags                  []string                            `json:"tags"`
	Name                  string                              `json:"name"`
	NamePrefix            string                              `json:"name_prefix"`
	AutomaticRestart      bool                                `json:"automatic_restart"`
	ServiceAccount        []GoogleComputeInstanceTemplateSpec `json:"service_account"`
	TagsFingerprint       string                              `json:"tags_fingerprint"`
	Disk                  []GoogleComputeInstanceTemplateSpec `json:"disk"`
	MachineType           string                              `json:"machine_type"`
	InstanceDescription   string                              `json:"instance_description"`
	SelfLink              string                              `json:"self_link"`
	GuestAccelerator      []GoogleComputeInstanceTemplateSpec `json:"guest_accelerator"`
	Labels                map[string]string                   `json:"labels"`
	CanIpForward          bool                                `json:"can_ip_forward"`
	Description           string                              `json:"description"`
}

type GoogleComputeInstanceTemplateStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// GoogleComputeInstanceTemplateList is a list of GoogleComputeInstanceTemplates
type GoogleComputeInstanceTemplateList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of GoogleComputeInstanceTemplate CRD objects
	Items []GoogleComputeInstanceTemplate `json:"items,omitempty"`
}
