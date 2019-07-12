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

type GoogleContainerClusterSpecClusterAutoscalingResourceLimits struct {
	ResourceType string `json:"resource_type"`
	Minimum      int    `json:"minimum"`
	Maximum      int    `json:"maximum"`
}

type GoogleContainerClusterSpecClusterAutoscaling struct {
	Enabled        bool                                           `json:"enabled"`
	ResourceLimits []GoogleContainerClusterSpecClusterAutoscaling `json:"resource_limits"`
}

type GoogleContainerClusterSpecNodePoolManagement struct {
	AutoRepair  bool `json:"auto_repair"`
	AutoUpgrade bool `json:"auto_upgrade"`
}

type GoogleContainerClusterSpecNodePoolNodeConfigWorkloadMetadataConfig struct {
	NodeMetadata string `json:"node_metadata"`
}

type GoogleContainerClusterSpecNodePoolNodeConfigGuestAccelerator struct {
	Type  string `json:"type"`
	Count int    `json:"count"`
}

type GoogleContainerClusterSpecNodePoolNodeConfigTaint struct {
	Key    string `json:"key"`
	Value  string `json:"value"`
	Effect string `json:"effect"`
}

type GoogleContainerClusterSpecNodePoolNodeConfig struct {
	WorkloadMetadataConfig []GoogleContainerClusterSpecNodePoolNodeConfig `json:"workload_metadata_config"`
	Labels                 map[string]string                              `json:"labels"`
	OauthScopes            []string                                       `json:"oauth_scopes"`
	ServiceAccount         string                                         `json:"service_account"`
	MachineType            string                                         `json:"machine_type"`
	DiskSizeGb             int                                            `json:"disk_size_gb"`
	GuestAccelerator       []GoogleContainerClusterSpecNodePoolNodeConfig `json:"guest_accelerator"`
	ImageType              string                                         `json:"image_type"`
	LocalSsdCount          int                                            `json:"local_ssd_count"`
	Preemptible            bool                                           `json:"preemptible"`
	Taint                  []GoogleContainerClusterSpecNodePoolNodeConfig `json:"taint"`
	Tags                   []string                                       `json:"tags"`
	DiskType               string                                         `json:"disk_type"`
	Metadata               map[string]string                              `json:"metadata"`
	MinCpuPlatform         string                                         `json:"min_cpu_platform"`
}

type GoogleContainerClusterSpecNodePoolAutoscaling struct {
	MinNodeCount int `json:"min_node_count"`
	MaxNodeCount int `json:"max_node_count"`
}

type GoogleContainerClusterSpecNodePool struct {
	Management        []GoogleContainerClusterSpecNodePool `json:"management"`
	NamePrefix        string                               `json:"name_prefix"`
	NodeConfig        []GoogleContainerClusterSpecNodePool `json:"node_config"`
	NodeCount         int                                  `json:"node_count"`
	Version           string                               `json:"version"`
	Autoscaling       []GoogleContainerClusterSpecNodePool `json:"autoscaling"`
	MaxPodsPerNode    int                                  `json:"max_pods_per_node"`
	InitialNodeCount  int                                  `json:"initial_node_count"`
	InstanceGroupUrls []string                             `json:"instance_group_urls"`
	Name              string                               `json:"name"`
}

type GoogleContainerClusterSpecPodSecurityPolicyConfig struct {
	Enabled bool `json:"enabled"`
}

type GoogleContainerClusterSpecIpAllocationPolicy struct {
	CreateSubnetwork           bool   `json:"create_subnetwork"`
	SubnetworkName             string `json:"subnetwork_name"`
	ClusterIpv4CidrBlock       string `json:"cluster_ipv4_cidr_block"`
	ServicesIpv4CidrBlock      string `json:"services_ipv4_cidr_block"`
	ClusterSecondaryRangeName  string `json:"cluster_secondary_range_name"`
	ServicesSecondaryRangeName string `json:"services_secondary_range_name"`
}

type GoogleContainerClusterSpecMaintenancePolicyDailyMaintenanceWindow struct {
	StartTime string `json:"start_time"`
	Duration  string `json:"duration"`
}

type GoogleContainerClusterSpecMaintenancePolicy struct {
	DailyMaintenanceWindow []GoogleContainerClusterSpecMaintenancePolicy `json:"daily_maintenance_window"`
}

type GoogleContainerClusterSpecNodeConfigTaint struct {
	Key    string `json:"key"`
	Value  string `json:"value"`
	Effect string `json:"effect"`
}

type GoogleContainerClusterSpecNodeConfigWorkloadMetadataConfig struct {
	NodeMetadata string `json:"node_metadata"`
}

type GoogleContainerClusterSpecNodeConfigGuestAccelerator struct {
	Count int    `json:"count"`
	Type  string `json:"type"`
}

type GoogleContainerClusterSpecNodeConfig struct {
	Taint                  []GoogleContainerClusterSpecNodeConfig `json:"taint"`
	LocalSsdCount          int                                    `json:"local_ssd_count"`
	Preemptible            bool                                   `json:"preemptible"`
	MinCpuPlatform         string                                 `json:"min_cpu_platform"`
	Tags                   []string                               `json:"tags"`
	DiskType               string                                 `json:"disk_type"`
	Metadata               map[string]string                      `json:"metadata"`
	ServiceAccount         string                                 `json:"service_account"`
	WorkloadMetadataConfig []GoogleContainerClusterSpecNodeConfig `json:"workload_metadata_config"`
	Labels                 map[string]string                      `json:"labels"`
	OauthScopes            []string                               `json:"oauth_scopes"`
	ImageType              string                                 `json:"image_type"`
	MachineType            string                                 `json:"machine_type"`
	DiskSizeGb             int                                    `json:"disk_size_gb"`
	GuestAccelerator       []GoogleContainerClusterSpecNodeConfig `json:"guest_accelerator"`
}

type GoogleContainerClusterSpecPrivateClusterConfig struct {
	MasterIpv4CidrBlock   string `json:"master_ipv4_cidr_block"`
	PrivateEndpoint       string `json:"private_endpoint"`
	PublicEndpoint        string `json:"public_endpoint"`
	EnablePrivateEndpoint bool   `json:"enable_private_endpoint"`
	EnablePrivateNodes    bool   `json:"enable_private_nodes"`
}

type GoogleContainerClusterSpecMasterAuthorizedNetworksConfigCidrBlocks struct {
	CidrBlock   string `json:"cidr_block"`
	DisplayName string `json:"display_name"`
}

type GoogleContainerClusterSpecMasterAuthorizedNetworksConfig struct {
	CidrBlocks []GoogleContainerClusterSpecMasterAuthorizedNetworksConfig `json:"cidr_blocks"`
}

type GoogleContainerClusterSpecNetworkPolicy struct {
	Provider string `json:"provider"`
	Enabled  bool   `json:"enabled"`
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

type GoogleContainerClusterSpec struct {
	MinMasterVersion               string                       `json:"min_master_version"`
	Name                           string                       `json:"name"`
	Description                    string                       `json:"description"`
	InitialNodeCount               int                          `json:"initial_node_count"`
	Region                         string                       `json:"region"`
	Project                        string                       `json:"project"`
	MasterVersion                  string                       `json:"master_version"`
	MasterIpv4CidrBlock            string                       `json:"master_ipv4_cidr_block"`
	ResourceLabels                 map[string]string            `json:"resource_labels"`
	AddonsConfig                   []GoogleContainerClusterSpec `json:"addons_config"`
	ClusterAutoscaling             []GoogleContainerClusterSpec `json:"cluster_autoscaling"`
	NodePool                       []GoogleContainerClusterSpec `json:"node_pool"`
	PodSecurityPolicyConfig        []GoogleContainerClusterSpec `json:"pod_security_policy_config"`
	Endpoint                       string                       `json:"endpoint"`
	IpAllocationPolicy             []GoogleContainerClusterSpec `json:"ip_allocation_policy"`
	PrivateCluster                 bool                         `json:"private_cluster"`
	ClusterIpv4Cidr                string                       `json:"cluster_ipv4_cidr"`
	EnableKubernetesAlpha          bool                         `json:"enable_kubernetes_alpha"`
	MaintenancePolicy              []GoogleContainerClusterSpec `json:"maintenance_policy"`
	NodeConfig                     []GoogleContainerClusterSpec `json:"node_config"`
	InstanceGroupUrls              []string                     `json:"instance_group_urls"`
	EnableBinaryAuthorization      bool                         `json:"enable_binary_authorization"`
	LoggingService                 string                       `json:"logging_service"`
	MonitoringService              string                       `json:"monitoring_service"`
	PrivateClusterConfig           []GoogleContainerClusterSpec `json:"private_cluster_config"`
	MasterAuthorizedNetworksConfig []GoogleContainerClusterSpec `json:"master_authorized_networks_config"`
	Network                        string                       `json:"network"`
	NetworkPolicy                  []GoogleContainerClusterSpec `json:"network_policy"`
	Zone                           string                       `json:"zone"`
	MasterAuth                     []GoogleContainerClusterSpec `json:"master_auth"`
	Subnetwork                     string                       `json:"subnetwork"`
	NodeVersion                    string                       `json:"node_version"`
	RemoveDefaultNodePool          bool                         `json:"remove_default_node_pool"`
	AdditionalZones                []string                     `json:"additional_zones"`
	EnableTpu                      bool                         `json:"enable_tpu"`
	EnableLegacyAbac               bool                         `json:"enable_legacy_abac"`
}

type GoogleContainerClusterStatus struct {
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
