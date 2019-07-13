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

type GoogleComputeInstanceTemplateSpecServiceAccount struct {
	Email  string   `json:"email"`
	Scopes []string `json:"scopes"`
}

type GoogleComputeInstanceTemplateSpecScheduling struct {
	Preemptible       bool   `json:"preemptible"`
	AutomaticRestart  bool   `json:"automatic_restart"`
	OnHostMaintenance string `json:"on_host_maintenance"`
}

type GoogleComputeInstanceTemplateSpecGuestAccelerator struct {
	Type  string `json:"type"`
	Count int    `json:"count"`
}

type GoogleComputeInstanceTemplateSpecDiskDiskEncryptionKey struct {
	KmsKeySelfLink string `json:"kms_key_self_link"`
}

type GoogleComputeInstanceTemplateSpecDisk struct {
	AutoDelete        bool                                    `json:"auto_delete"`
	DeviceName        string                                  `json:"device_name"`
	DiskSizeGb        int                                     `json:"disk_size_gb"`
	Mode              string                                  `json:"mode"`
	Source            string                                  `json:"source"`
	Type              string                                  `json:"type"`
	DiskEncryptionKey []GoogleComputeInstanceTemplateSpecDisk `json:"disk_encryption_key"`
	Boot              bool                                    `json:"boot"`
	DiskName          string                                  `json:"disk_name"`
	DiskType          string                                  `json:"disk_type"`
	SourceImage       string                                  `json:"source_image"`
	Interface         string                                  `json:"interface"`
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
	SubnetworkProject string                                              `json:"subnetwork_project"`
	AccessConfig      []GoogleComputeInstanceTemplateSpecNetworkInterface `json:"access_config"`
	AliasIpRange      []GoogleComputeInstanceTemplateSpecNetworkInterface `json:"alias_ip_range"`
	Network           string                                              `json:"network"`
	Address           string                                              `json:"address"`
	NetworkIp         string                                              `json:"network_ip"`
	Subnetwork        string                                              `json:"subnetwork"`
}

type GoogleComputeInstanceTemplateSpec struct {
	Name                  string                              `json:"name"`
	MachineType           string                              `json:"machine_type"`
	Description           string                              `json:"description"`
	ServiceAccount        []GoogleComputeInstanceTemplateSpec `json:"service_account"`
	Tags                  []string                            `json:"tags"`
	NamePrefix            string                              `json:"name_prefix"`
	MetadataFingerprint   string                              `json:"metadata_fingerprint"`
	Scheduling            []GoogleComputeInstanceTemplateSpec `json:"scheduling"`
	GuestAccelerator      []GoogleComputeInstanceTemplateSpec `json:"guest_accelerator"`
	TagsFingerprint       string                              `json:"tags_fingerprint"`
	Labels                map[string]string                   `json:"labels"`
	MinCpuPlatform        string                              `json:"min_cpu_platform"`
	Disk                  []GoogleComputeInstanceTemplateSpec `json:"disk"`
	AutomaticRestart      bool                                `json:"automatic_restart"`
	CanIpForward          bool                                `json:"can_ip_forward"`
	Metadata              map[string]string                   `json:"metadata"`
	MetadataStartupScript string                              `json:"metadata_startup_script"`
	NetworkInterface      []GoogleComputeInstanceTemplateSpec `json:"network_interface"`
	Project               string                              `json:"project"`
	InstanceDescription   string                              `json:"instance_description"`
	OnHostMaintenance     string                              `json:"on_host_maintenance"`
	Region                string                              `json:"region"`
	SelfLink              string                              `json:"self_link"`
}



type GoogleComputeInstanceTemplateStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

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