package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

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
	ClusterCaCertificate    string                                 `json:"cluster_ca_certificate"`
	Password                string                                 `json:"password"`
	Username                string                                 `json:"username"`
	ClientCertificateConfig []GoogleContainerClusterSpecMasterAuth `json:"client_certificate_config"`
	ClientCertificate       string                                 `json:"client_certificate"`
	ClientKey               string                                 `json:"client_key"`
}

type GoogleContainerClusterSpecClusterAutoscalingResourceLimits struct {
	ResourceType string `json:"resource_type"`
	Minimum      int    `json:"minimum"`
	Maximum      int    `json:"maximum"`
}

type GoogleContainerClusterSpecClusterAutoscaling struct {
	Enabled        bool                                           `json:"enabled"`
	ResourceLimits []GoogleContainerClusterSpecClusterAutoscaling `json:"resource_limits"`
}

type GoogleContainerClusterSpecMaintenancePolicyDailyMaintenanceWindow struct {
	StartTime string `json:"start_time"`
	Duration  string `json:"duration"`
}

type GoogleContainerClusterSpecMaintenancePolicy struct {
	DailyMaintenanceWindow []GoogleContainerClusterSpecMaintenancePolicy `json:"daily_maintenance_window"`
}

type GoogleContainerClusterSpecNodeConfigWorkloadMetadataConfig struct {
	NodeMetadata string `json:"node_metadata"`
}

type GoogleContainerClusterSpecNodeConfigGuestAccelerator struct {
	Count int    `json:"count"`
	Type  string `json:"type"`
}

type GoogleContainerClusterSpecNodeConfigTaint struct {
	Key    string `json:"key"`
	Value  string `json:"value"`
	Effect string `json:"effect"`
}

type GoogleContainerClusterSpecNodeConfig struct {
	WorkloadMetadataConfig []GoogleContainerClusterSpecNodeConfig `json:"workload_metadata_config"`
	GuestAccelerator       []GoogleContainerClusterSpecNodeConfig `json:"guest_accelerator"`
	ImageType              string                                 `json:"image_type"`
	Preemptible            bool                                   `json:"preemptible"`
	Tags                   []string                               `json:"tags"`
	DiskSizeGb             int                                    `json:"disk_size_gb"`
	ServiceAccount         string                                 `json:"service_account"`
	Metadata               map[string]string                      `json:"metadata"`
	DiskType               string                                 `json:"disk_type"`
	Labels                 map[string]string                      `json:"labels"`
	LocalSsdCount          int                                    `json:"local_ssd_count"`
	MachineType            string                                 `json:"machine_type"`
	MinCpuPlatform         string                                 `json:"min_cpu_platform"`
	OauthScopes            []string                               `json:"oauth_scopes"`
	Taint                  []GoogleContainerClusterSpecNodeConfig `json:"taint"`
}

type GoogleContainerClusterSpecPodSecurityPolicyConfig struct {
	Enabled bool `json:"enabled"`
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

type GoogleContainerClusterSpecIpAllocationPolicy struct {
	CreateSubnetwork           bool   `json:"create_subnetwork"`
	SubnetworkName             string `json:"subnetwork_name"`
	ClusterIpv4CidrBlock       string `json:"cluster_ipv4_cidr_block"`
	ServicesIpv4CidrBlock      string `json:"services_ipv4_cidr_block"`
	ClusterSecondaryRangeName  string `json:"cluster_secondary_range_name"`
	ServicesSecondaryRangeName string `json:"services_secondary_range_name"`
}

type GoogleContainerClusterSpecAddonsConfigNetworkPolicyConfig struct {
	Disabled bool `json:"disabled"`
}

type GoogleContainerClusterSpecAddonsConfigHttpLoadBalancing struct {
	Disabled bool `json:"disabled"`
}

type GoogleContainerClusterSpecAddonsConfigHorizontalPodAutoscaling struct {
	Disabled bool `json:"disabled"`
}

type GoogleContainerClusterSpecAddonsConfigKubernetesDashboard struct {
	Disabled bool `json:"disabled"`
}

type GoogleContainerClusterSpecAddonsConfig struct {
	NetworkPolicyConfig      []GoogleContainerClusterSpecAddonsConfig `json:"network_policy_config"`
	HttpLoadBalancing        []GoogleContainerClusterSpecAddonsConfig `json:"http_load_balancing"`
	HorizontalPodAutoscaling []GoogleContainerClusterSpecAddonsConfig `json:"horizontal_pod_autoscaling"`
	KubernetesDashboard      []GoogleContainerClusterSpecAddonsConfig `json:"kubernetes_dashboard"`
}

type GoogleContainerClusterSpecNodePoolNodeConfigTaint struct {
	Key    string `json:"key"`
	Value  string `json:"value"`
	Effect string `json:"effect"`
}

type GoogleContainerClusterSpecNodePoolNodeConfigGuestAccelerator struct {
	Count int    `json:"count"`
	Type  string `json:"type"`
}

type GoogleContainerClusterSpecNodePoolNodeConfigWorkloadMetadataConfig struct {
	NodeMetadata string `json:"node_metadata"`
}

type GoogleContainerClusterSpecNodePoolNodeConfig struct {
	MinCpuPlatform         string                                         `json:"min_cpu_platform"`
	OauthScopes            []string                                       `json:"oauth_scopes"`
	Taint                  []GoogleContainerClusterSpecNodePoolNodeConfig `json:"taint"`
	GuestAccelerator       []GoogleContainerClusterSpecNodePoolNodeConfig `json:"guest_accelerator"`
	ImageType              string                                         `json:"image_type"`
	Preemptible            bool                                           `json:"preemptible"`
	Tags                   []string                                       `json:"tags"`
	WorkloadMetadataConfig []GoogleContainerClusterSpecNodePoolNodeConfig `json:"workload_metadata_config"`
	DiskSizeGb             int                                            `json:"disk_size_gb"`
	ServiceAccount         string                                         `json:"service_account"`
	DiskType               string                                         `json:"disk_type"`
	Labels                 map[string]string                              `json:"labels"`
	LocalSsdCount          int                                            `json:"local_ssd_count"`
	MachineType            string                                         `json:"machine_type"`
	Metadata               map[string]string                              `json:"metadata"`
}

type GoogleContainerClusterSpecNodePoolManagement struct {
	AutoRepair  bool `json:"auto_repair"`
	AutoUpgrade bool `json:"auto_upgrade"`
}

type GoogleContainerClusterSpecNodePoolAutoscaling struct {
	MinNodeCount int `json:"min_node_count"`
	MaxNodeCount int `json:"max_node_count"`
}

type GoogleContainerClusterSpecNodePool struct {
	InitialNodeCount  int                                  `json:"initial_node_count"`
	Name              string                               `json:"name"`
	NamePrefix        string                               `json:"name_prefix"`
	NodeConfig        []GoogleContainerClusterSpecNodePool `json:"node_config"`
	NodeCount         int                                  `json:"node_count"`
	MaxPodsPerNode    int                                  `json:"max_pods_per_node"`
	InstanceGroupUrls []string                             `json:"instance_group_urls"`
	Management        []GoogleContainerClusterSpecNodePool `json:"management"`
	Version           string                               `json:"version"`
	Autoscaling       []GoogleContainerClusterSpecNodePool `json:"autoscaling"`
}

type GoogleContainerClusterSpecPrivateClusterConfig struct {
	MasterIpv4CidrBlock   string `json:"master_ipv4_cidr_block"`
	PrivateEndpoint       string `json:"private_endpoint"`
	PublicEndpoint        string `json:"public_endpoint"`
	EnablePrivateEndpoint bool   `json:"enable_private_endpoint"`
	EnablePrivateNodes    bool   `json:"enable_private_nodes"`
}

type GoogleContainerClusterSpec struct {
	EnableLegacyAbac               bool                         `json:"enable_legacy_abac"`
	InitialNodeCount               int                          `json:"initial_node_count"`
	MasterAuth                     []GoogleContainerClusterSpec `json:"master_auth"`
	MinMasterVersion               string                       `json:"min_master_version"`
	Endpoint                       string                       `json:"endpoint"`
	Name                           string                       `json:"name"`
	Region                         string                       `json:"region"`
	ClusterAutoscaling             []GoogleContainerClusterSpec `json:"cluster_autoscaling"`
	EnableKubernetesAlpha          bool                         `json:"enable_kubernetes_alpha"`
	EnableTpu                      bool                         `json:"enable_tpu"`
	MaintenancePolicy              []GoogleContainerClusterSpec `json:"maintenance_policy"`
	MasterVersion                  string                       `json:"master_version"`
	PrivateCluster                 bool                         `json:"private_cluster"`
	NodeConfig                     []GoogleContainerClusterSpec `json:"node_config"`
	PodSecurityPolicyConfig        []GoogleContainerClusterSpec `json:"pod_security_policy_config"`
	AdditionalZones                []string                     `json:"additional_zones"`
	MasterAuthorizedNetworksConfig []GoogleContainerClusterSpec `json:"master_authorized_networks_config"`
	Description                    string                       `json:"description"`
	EnableBinaryAuthorization      bool                         `json:"enable_binary_authorization"`
	LoggingService                 string                       `json:"logging_service"`
	NetworkPolicy                  []GoogleContainerClusterSpec `json:"network_policy"`
	InstanceGroupUrls              []string                     `json:"instance_group_urls"`
	IpAllocationPolicy             []GoogleContainerClusterSpec `json:"ip_allocation_policy"`
	AddonsConfig                   []GoogleContainerClusterSpec `json:"addons_config"`
	Network                        string                       `json:"network"`
	NodePool                       []GoogleContainerClusterSpec `json:"node_pool"`
	MasterIpv4CidrBlock            string                       `json:"master_ipv4_cidr_block"`
	ResourceLabels                 map[string]string            `json:"resource_labels"`
	RemoveDefaultNodePool          bool                         `json:"remove_default_node_pool"`
	PrivateClusterConfig           []GoogleContainerClusterSpec `json:"private_cluster_config"`
	Zone                           string                       `json:"zone"`
	ClusterIpv4Cidr                string                       `json:"cluster_ipv4_cidr"`
	MonitoringService              string                       `json:"monitoring_service"`
	NodeVersion                    string                       `json:"node_version"`
	Project                        string                       `json:"project"`
	Subnetwork                     string                       `json:"subnetwork"`
}

type GoogleContainerClusterStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// GoogleContainerClusterList is a list of GoogleContainerClusters
type GoogleContainerClusterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of GoogleContainerCluster CRD objects
	Items []GoogleContainerCluster `json:"items,omitempty"`
}
