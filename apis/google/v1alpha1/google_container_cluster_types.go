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

type GoogleContainerCluster struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              GoogleContainerClusterSpec   `json:"spec,omitempty"`
	Status            GoogleContainerClusterStatus `json:"status,omitempty"`
}

type GoogleContainerClusterSpecMasterAuthClientCertificateConfig struct {
	IssueClientCertificate bool `json:"issue_client_certificate"`
}

type GoogleContainerClusterSpecMasterAuth struct {
	Password                string                                 `json:"password"`
	Username                string                                 `json:"username"`
	ClientCertificateConfig []GoogleContainerClusterSpecMasterAuth `json:"client_certificate_config"`
	ClientCertificate       string                                 `json:"client_certificate"`
	ClientKey               string                                 `json:"client_key"`
	ClusterCaCertificate    string                                 `json:"cluster_ca_certificate"`
}

type GoogleContainerClusterSpecMasterAuthorizedNetworksConfigCidrBlocks struct {
	CidrBlock   string `json:"cidr_block"`
	DisplayName string `json:"display_name"`
}

type GoogleContainerClusterSpecMasterAuthorizedNetworksConfig struct {
	CidrBlocks []GoogleContainerClusterSpecMasterAuthorizedNetworksConfig `json:"cidr_blocks"`
}

type GoogleContainerClusterSpecNetworkPolicy struct {
	Enabled  bool   `json:"enabled"`
	Provider string `json:"provider"`
}

type GoogleContainerClusterSpecNodeConfigWorkloadMetadataConfig struct {
	NodeMetadata string `json:"node_metadata"`
}

type GoogleContainerClusterSpecNodeConfigTaint struct {
	Key    string `json:"key"`
	Value  string `json:"value"`
	Effect string `json:"effect"`
}

type GoogleContainerClusterSpecNodeConfigGuestAccelerator struct {
	Count int    `json:"count"`
	Type  string `json:"type"`
}

type GoogleContainerClusterSpecNodeConfig struct {
	WorkloadMetadataConfig []GoogleContainerClusterSpecNodeConfig `json:"workload_metadata_config"`
	DiskType               string                                 `json:"disk_type"`
	MachineType            string                                 `json:"machine_type"`
	OauthScopes            []string                               `json:"oauth_scopes"`
	Taint                  []GoogleContainerClusterSpecNodeConfig `json:"taint"`
	Tags                   []string                               `json:"tags"`
	GuestAccelerator       []GoogleContainerClusterSpecNodeConfig `json:"guest_accelerator"`
	LocalSsdCount          int                                    `json:"local_ssd_count"`
	Metadata               map[string]string                      `json:"metadata"`
	MinCpuPlatform         string                                 `json:"min_cpu_platform"`
	Preemptible            bool                                   `json:"preemptible"`
	ServiceAccount         string                                 `json:"service_account"`
	DiskSizeGb             int                                    `json:"disk_size_gb"`
	ImageType              string                                 `json:"image_type"`
	Labels                 map[string]string                      `json:"labels"`
}

type GoogleContainerClusterSpecAddonsConfigHorizontalPodAutoscaling struct {
	Disabled bool `json:"disabled"`
}

type GoogleContainerClusterSpecAddonsConfigKubernetesDashboard struct {
	Disabled bool `json:"disabled"`
}

type GoogleContainerClusterSpecAddonsConfigNetworkPolicyConfig struct {
	Disabled bool `json:"disabled"`
}

type GoogleContainerClusterSpecAddonsConfigHttpLoadBalancing struct {
	Disabled bool `json:"disabled"`
}

type GoogleContainerClusterSpecAddonsConfig struct {
	HorizontalPodAutoscaling []GoogleContainerClusterSpecAddonsConfig `json:"horizontal_pod_autoscaling"`
	KubernetesDashboard      []GoogleContainerClusterSpecAddonsConfig `json:"kubernetes_dashboard"`
	NetworkPolicyConfig      []GoogleContainerClusterSpecAddonsConfig `json:"network_policy_config"`
	HttpLoadBalancing        []GoogleContainerClusterSpecAddonsConfig `json:"http_load_balancing"`
}

type GoogleContainerClusterSpecClusterAutoscalingResourceLimits struct {
	Minimum      int    `json:"minimum"`
	Maximum      int    `json:"maximum"`
	ResourceType string `json:"resource_type"`
}

type GoogleContainerClusterSpecClusterAutoscaling struct {
	ResourceLimits []GoogleContainerClusterSpecClusterAutoscaling `json:"resource_limits"`
	Enabled        bool                                           `json:"enabled"`
}

type GoogleContainerClusterSpecIpAllocationPolicy struct {
	ClusterSecondaryRangeName  string `json:"cluster_secondary_range_name"`
	ServicesSecondaryRangeName string `json:"services_secondary_range_name"`
	CreateSubnetwork           bool   `json:"create_subnetwork"`
	SubnetworkName             string `json:"subnetwork_name"`
	ClusterIpv4CidrBlock       string `json:"cluster_ipv4_cidr_block"`
	ServicesIpv4CidrBlock      string `json:"services_ipv4_cidr_block"`
}

type GoogleContainerClusterSpecPrivateClusterConfig struct {
	PrivateEndpoint       string `json:"private_endpoint"`
	PublicEndpoint        string `json:"public_endpoint"`
	EnablePrivateEndpoint bool   `json:"enable_private_endpoint"`
	EnablePrivateNodes    bool   `json:"enable_private_nodes"`
	MasterIpv4CidrBlock   string `json:"master_ipv4_cidr_block"`
}

type GoogleContainerClusterSpecNodePoolManagement struct {
	AutoRepair  bool `json:"auto_repair"`
	AutoUpgrade bool `json:"auto_upgrade"`
}

type GoogleContainerClusterSpecNodePoolNodeConfigGuestAccelerator struct {
	Count int    `json:"count"`
	Type  string `json:"type"`
}

type GoogleContainerClusterSpecNodePoolNodeConfigWorkloadMetadataConfig struct {
	NodeMetadata string `json:"node_metadata"`
}

type GoogleContainerClusterSpecNodePoolNodeConfigTaint struct {
	Key    string `json:"key"`
	Value  string `json:"value"`
	Effect string `json:"effect"`
}

type GoogleContainerClusterSpecNodePoolNodeConfig struct {
	GuestAccelerator       []GoogleContainerClusterSpecNodePoolNodeConfig `json:"guest_accelerator"`
	LocalSsdCount          int                                            `json:"local_ssd_count"`
	Metadata               map[string]string                              `json:"metadata"`
	Tags                   []string                                       `json:"tags"`
	ServiceAccount         string                                         `json:"service_account"`
	DiskSizeGb             int                                            `json:"disk_size_gb"`
	ImageType              string                                         `json:"image_type"`
	Labels                 map[string]string                              `json:"labels"`
	MinCpuPlatform         string                                         `json:"min_cpu_platform"`
	Preemptible            bool                                           `json:"preemptible"`
	DiskType               string                                         `json:"disk_type"`
	MachineType            string                                         `json:"machine_type"`
	OauthScopes            []string                                       `json:"oauth_scopes"`
	WorkloadMetadataConfig []GoogleContainerClusterSpecNodePoolNodeConfig `json:"workload_metadata_config"`
	Taint                  []GoogleContainerClusterSpecNodePoolNodeConfig `json:"taint"`
}

type GoogleContainerClusterSpecNodePoolAutoscaling struct {
	MinNodeCount int `json:"min_node_count"`
	MaxNodeCount int `json:"max_node_count"`
}

type GoogleContainerClusterSpecNodePool struct {
	Management        []GoogleContainerClusterSpecNodePool `json:"management"`
	Name              string                               `json:"name"`
	NamePrefix        string                               `json:"name_prefix"`
	NodeConfig        []GoogleContainerClusterSpecNodePool `json:"node_config"`
	Version           string                               `json:"version"`
	Autoscaling       []GoogleContainerClusterSpecNodePool `json:"autoscaling"`
	InitialNodeCount  int                                  `json:"initial_node_count"`
	InstanceGroupUrls []string                             `json:"instance_group_urls"`
	MaxPodsPerNode    int                                  `json:"max_pods_per_node"`
	NodeCount         int                                  `json:"node_count"`
}

type GoogleContainerClusterSpecPodSecurityPolicyConfig struct {
	Enabled bool `json:"enabled"`
}

type GoogleContainerClusterSpecMaintenancePolicyDailyMaintenanceWindow struct {
	Duration  string `json:"duration"`
	StartTime string `json:"start_time"`
}

type GoogleContainerClusterSpecMaintenancePolicy struct {
	DailyMaintenanceWindow []GoogleContainerClusterSpecMaintenancePolicy `json:"daily_maintenance_window"`
}

type GoogleContainerClusterSpec struct {
	EnableLegacyAbac               bool                         `json:"enable_legacy_abac"`
	LoggingService                 string                       `json:"logging_service"`
	MonitoringService              string                       `json:"monitoring_service"`
	RemoveDefaultNodePool          bool                         `json:"remove_default_node_pool"`
	Network                        string                       `json:"network"`
	NodeVersion                    string                       `json:"node_version"`
	MasterIpv4CidrBlock            string                       `json:"master_ipv4_cidr_block"`
	MasterAuth                     []GoogleContainerClusterSpec `json:"master_auth"`
	MasterAuthorizedNetworksConfig []GoogleContainerClusterSpec `json:"master_authorized_networks_config"`
	NetworkPolicy                  []GoogleContainerClusterSpec `json:"network_policy"`
	NodeConfig                     []GoogleContainerClusterSpec `json:"node_config"`
	Name                           string                       `json:"name"`
	Zone                           string                       `json:"zone"`
	AddonsConfig                   []GoogleContainerClusterSpec `json:"addons_config"`
	ClusterAutoscaling             []GoogleContainerClusterSpec `json:"cluster_autoscaling"`
	InstanceGroupUrls              []string                     `json:"instance_group_urls"`
	PrivateCluster                 bool                         `json:"private_cluster"`
	IpAllocationPolicy             []GoogleContainerClusterSpec `json:"ip_allocation_policy"`
	AdditionalZones                []string                     `json:"additional_zones"`
	EnableTpu                      bool                         `json:"enable_tpu"`
	MinMasterVersion               string                       `json:"min_master_version"`
	MasterVersion                  string                       `json:"master_version"`
	Region                         string                       `json:"region"`
	Project                        string                       `json:"project"`
	Subnetwork                     string                       `json:"subnetwork"`
	Endpoint                       string                       `json:"endpoint"`
	PrivateClusterConfig           []GoogleContainerClusterSpec `json:"private_cluster_config"`
	ResourceLabels                 map[string]string            `json:"resource_labels"`
	ClusterIpv4Cidr                string                       `json:"cluster_ipv4_cidr"`
	Description                    string                       `json:"description"`
	NodePool                       []GoogleContainerClusterSpec `json:"node_pool"`
	PodSecurityPolicyConfig        []GoogleContainerClusterSpec `json:"pod_security_policy_config"`
	EnableBinaryAuthorization      bool                         `json:"enable_binary_authorization"`
	EnableKubernetesAlpha          bool                         `json:"enable_kubernetes_alpha"`
	InitialNodeCount               int                          `json:"initial_node_count"`
	MaintenancePolicy              []GoogleContainerClusterSpec `json:"maintenance_policy"`
}



type GoogleContainerClusterStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// GoogleContainerClusterList is a list of GoogleContainerClusters
type GoogleContainerClusterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of GoogleContainerCluster CRD objects
	Items []GoogleContainerCluster `json:"items,omitempty"`
}