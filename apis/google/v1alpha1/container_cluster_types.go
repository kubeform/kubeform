package v1alpha1

import (
	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

type ContainerCluster struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ContainerClusterSpec   `json:"spec,omitempty"`
	Status            ContainerClusterStatus `json:"status,omitempty"`
}

type ContainerClusterSpecAddonsConfigHorizontalPodAutoscaling struct {
	// +optional
	Disabled bool `json:"disabled,omitempty" tf:"disabled,omitempty"`
}

type ContainerClusterSpecAddonsConfigHttpLoadBalancing struct {
	// +optional
	Disabled bool `json:"disabled,omitempty" tf:"disabled,omitempty"`
}

type ContainerClusterSpecAddonsConfigKubernetesDashboard struct {
	// +optional
	Disabled bool `json:"disabled,omitempty" tf:"disabled,omitempty"`
}

type ContainerClusterSpecAddonsConfigNetworkPolicyConfig struct {
	// +optional
	Disabled bool `json:"disabled,omitempty" tf:"disabled,omitempty"`
}

type ContainerClusterSpecAddonsConfig struct {
	// +optional
	// +kubebuilder:validation:MaxItems=1
	HorizontalPodAutoscaling []ContainerClusterSpecAddonsConfigHorizontalPodAutoscaling `json:"horizontalPodAutoscaling,omitempty" tf:"horizontal_pod_autoscaling,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	HttpLoadBalancing []ContainerClusterSpecAddonsConfigHttpLoadBalancing `json:"httpLoadBalancing,omitempty" tf:"http_load_balancing,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	KubernetesDashboard []ContainerClusterSpecAddonsConfigKubernetesDashboard `json:"kubernetesDashboard,omitempty" tf:"kubernetes_dashboard,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	NetworkPolicyConfig []ContainerClusterSpecAddonsConfigNetworkPolicyConfig `json:"networkPolicyConfig,omitempty" tf:"network_policy_config,omitempty"`
}

type ContainerClusterSpecIpAllocationPolicy struct {
	// +optional
	ClusterIpv4CIDRBlock string `json:"clusterIpv4CIDRBlock,omitempty" tf:"cluster_ipv4_cidr_block,omitempty"`
	// +optional
	ClusterSecondaryRangeName string `json:"clusterSecondaryRangeName,omitempty" tf:"cluster_secondary_range_name,omitempty"`
	// +optional
	CreateSubnetwork bool `json:"createSubnetwork,omitempty" tf:"create_subnetwork,omitempty"`
	// +optional
	ServicesIpv4CIDRBlock string `json:"servicesIpv4CIDRBlock,omitempty" tf:"services_ipv4_cidr_block,omitempty"`
	// +optional
	ServicesSecondaryRangeName string `json:"servicesSecondaryRangeName,omitempty" tf:"services_secondary_range_name,omitempty"`
	// +optional
	SubnetworkName string `json:"subnetworkName,omitempty" tf:"subnetwork_name,omitempty"`
}

type ContainerClusterSpecMaintenancePolicyDailyMaintenanceWindow struct {
	StartTime string `json:"startTime" tf:"start_time"`
}

type ContainerClusterSpecMaintenancePolicy struct {
	// +kubebuilder:validation:MaxItems=1
	DailyMaintenanceWindow []ContainerClusterSpecMaintenancePolicyDailyMaintenanceWindow `json:"dailyMaintenanceWindow" tf:"daily_maintenance_window"`
}

type ContainerClusterSpecMasterAuthClientCertificateConfig struct {
	IssueClientCertificate bool `json:"issueClientCertificate" tf:"issue_client_certificate"`
}

type ContainerClusterSpecMasterAuth struct {
	// +optional
	// +kubebuilder:validation:MaxItems=1
	ClientCertificateConfig []ContainerClusterSpecMasterAuthClientCertificateConfig `json:"clientCertificateConfig,omitempty" tf:"client_certificate_config,omitempty"`
	// Sensitive Data. Provide secret name which contains one value only
	Password core.LocalObjectReference `json:"password" tf:"password"`
	Username string                    `json:"username" tf:"username"`
}

type ContainerClusterSpecMasterAuthorizedNetworksConfigCidrBlocks struct {
	CidrBlock string `json:"cidrBlock" tf:"cidr_block"`
	// +optional
	DisplayName string `json:"displayName,omitempty" tf:"display_name,omitempty"`
}

type ContainerClusterSpecMasterAuthorizedNetworksConfig struct {
	// +optional
	// +kubebuilder:validation:MaxItems=20
	// +kubebuilder:validation:UniqueItems=true
	CidrBlocks []ContainerClusterSpecMasterAuthorizedNetworksConfigCidrBlocks `json:"cidrBlocks,omitempty" tf:"cidr_blocks,omitempty"`
}

type ContainerClusterSpecNetworkPolicy struct {
	// +optional
	Enabled bool `json:"enabled,omitempty" tf:"enabled,omitempty"`
	// +optional
	Provider string `json:"provider,omitempty" tf:"provider,omitempty"`
}

type ContainerClusterSpecNodeConfigGuestAccelerator struct {
	Count int    `json:"count" tf:"count"`
	Type  string `json:"type" tf:"type"`
}

type ContainerClusterSpecNodeConfigTaint struct {
	Effect string `json:"effect" tf:"effect"`
	Key    string `json:"key" tf:"key"`
	Value  string `json:"value" tf:"value"`
}

type ContainerClusterSpecNodeConfigWorkloadMetadataConfig struct {
	NodeMetadata string `json:"nodeMetadata" tf:"node_metadata"`
}

type ContainerClusterSpecNodeConfig struct {
	// +optional
	DiskSizeGb int `json:"diskSizeGb,omitempty" tf:"disk_size_gb,omitempty"`
	// +optional
	DiskType string `json:"diskType,omitempty" tf:"disk_type,omitempty"`
	// +optional
	GuestAccelerator []ContainerClusterSpecNodeConfigGuestAccelerator `json:"guestAccelerator,omitempty" tf:"guest_accelerator,omitempty"`
	// +optional
	ImageType string `json:"imageType,omitempty" tf:"image_type,omitempty"`
	// +optional
	Labels map[string]string `json:"labels,omitempty" tf:"labels,omitempty"`
	// +optional
	LocalSsdCount int `json:"localSsdCount,omitempty" tf:"local_ssd_count,omitempty"`
	// +optional
	MachineType string `json:"machineType,omitempty" tf:"machine_type,omitempty"`
	// +optional
	Metadata map[string]string `json:"metadata,omitempty" tf:"metadata,omitempty"`
	// +optional
	MinCPUPlatform string `json:"minCPUPlatform,omitempty" tf:"min_cpu_platform,omitempty"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	OauthScopes []string `json:"oauthScopes,omitempty" tf:"oauth_scopes,omitempty"`
	// +optional
	Preemptible bool `json:"preemptible,omitempty" tf:"preemptible,omitempty"`
	// +optional
	ServiceAccount string `json:"serviceAccount,omitempty" tf:"service_account,omitempty"`
	// +optional
	Tags []string `json:"tags,omitempty" tf:"tags,omitempty"`
	// +optional
	// Deprecated
	Taint []ContainerClusterSpecNodeConfigTaint `json:"taint,omitempty" tf:"taint,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	// Deprecated
	WorkloadMetadataConfig []ContainerClusterSpecNodeConfigWorkloadMetadataConfig `json:"workloadMetadataConfig,omitempty" tf:"workload_metadata_config,omitempty"`
}

type ContainerClusterSpecNodePoolAutoscaling struct {
	MaxNodeCount int `json:"maxNodeCount" tf:"max_node_count"`
	MinNodeCount int `json:"minNodeCount" tf:"min_node_count"`
}

type ContainerClusterSpecNodePoolManagement struct {
	// +optional
	AutoRepair bool `json:"autoRepair,omitempty" tf:"auto_repair,omitempty"`
	// +optional
	AutoUpgrade bool `json:"autoUpgrade,omitempty" tf:"auto_upgrade,omitempty"`
}

type ContainerClusterSpecNodePoolNodeConfigGuestAccelerator struct {
	Count int    `json:"count" tf:"count"`
	Type  string `json:"type" tf:"type"`
}

type ContainerClusterSpecNodePoolNodeConfigTaint struct {
	Effect string `json:"effect" tf:"effect"`
	Key    string `json:"key" tf:"key"`
	Value  string `json:"value" tf:"value"`
}

type ContainerClusterSpecNodePoolNodeConfigWorkloadMetadataConfig struct {
	NodeMetadata string `json:"nodeMetadata" tf:"node_metadata"`
}

type ContainerClusterSpecNodePoolNodeConfig struct {
	// +optional
	DiskSizeGb int `json:"diskSizeGb,omitempty" tf:"disk_size_gb,omitempty"`
	// +optional
	DiskType string `json:"diskType,omitempty" tf:"disk_type,omitempty"`
	// +optional
	GuestAccelerator []ContainerClusterSpecNodePoolNodeConfigGuestAccelerator `json:"guestAccelerator,omitempty" tf:"guest_accelerator,omitempty"`
	// +optional
	ImageType string `json:"imageType,omitempty" tf:"image_type,omitempty"`
	// +optional
	Labels map[string]string `json:"labels,omitempty" tf:"labels,omitempty"`
	// +optional
	LocalSsdCount int `json:"localSsdCount,omitempty" tf:"local_ssd_count,omitempty"`
	// +optional
	MachineType string `json:"machineType,omitempty" tf:"machine_type,omitempty"`
	// +optional
	Metadata map[string]string `json:"metadata,omitempty" tf:"metadata,omitempty"`
	// +optional
	MinCPUPlatform string `json:"minCPUPlatform,omitempty" tf:"min_cpu_platform,omitempty"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	OauthScopes []string `json:"oauthScopes,omitempty" tf:"oauth_scopes,omitempty"`
	// +optional
	Preemptible bool `json:"preemptible,omitempty" tf:"preemptible,omitempty"`
	// +optional
	ServiceAccount string `json:"serviceAccount,omitempty" tf:"service_account,omitempty"`
	// +optional
	Tags []string `json:"tags,omitempty" tf:"tags,omitempty"`
	// +optional
	// Deprecated
	Taint []ContainerClusterSpecNodePoolNodeConfigTaint `json:"taint,omitempty" tf:"taint,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	// Deprecated
	WorkloadMetadataConfig []ContainerClusterSpecNodePoolNodeConfigWorkloadMetadataConfig `json:"workloadMetadataConfig,omitempty" tf:"workload_metadata_config,omitempty"`
}

type ContainerClusterSpecNodePool struct {
	// +optional
	// +kubebuilder:validation:MaxItems=1
	Autoscaling []ContainerClusterSpecNodePoolAutoscaling `json:"autoscaling,omitempty" tf:"autoscaling,omitempty"`
	// +optional
	InitialNodeCount int `json:"initialNodeCount,omitempty" tf:"initial_node_count,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	Management []ContainerClusterSpecNodePoolManagement `json:"management,omitempty" tf:"management,omitempty"`
	// +optional
	// Deprecated
	MaxPodsPerNode int `json:"maxPodsPerNode,omitempty" tf:"max_pods_per_node,omitempty"`
	// +optional
	Name string `json:"name,omitempty" tf:"name,omitempty"`
	// +optional
	// Deprecated
	NamePrefix string `json:"namePrefix,omitempty" tf:"name_prefix,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	NodeConfig []ContainerClusterSpecNodePoolNodeConfig `json:"nodeConfig,omitempty" tf:"node_config,omitempty"`
	// +optional
	NodeCount int `json:"nodeCount,omitempty" tf:"node_count,omitempty"`
	// +optional
	Version string `json:"version,omitempty" tf:"version,omitempty"`
}

type ContainerClusterSpecPodSecurityPolicyConfig struct {
	Enabled bool `json:"enabled" tf:"enabled"`
}

type ContainerClusterSpecPrivateClusterConfig struct {
	// +optional
	EnablePrivateEndpoint bool `json:"enablePrivateEndpoint,omitempty" tf:"enable_private_endpoint,omitempty"`
	// +optional
	EnablePrivateNodes bool `json:"enablePrivateNodes,omitempty" tf:"enable_private_nodes,omitempty"`
	// +optional
	MasterIpv4CIDRBlock string `json:"masterIpv4CIDRBlock,omitempty" tf:"master_ipv4_cidr_block,omitempty"`
}

type ContainerClusterSpec struct {
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	AdditionalZones []string `json:"additionalZones,omitempty" tf:"additional_zones,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	AddonsConfig []ContainerClusterSpecAddonsConfig `json:"addonsConfig,omitempty" tf:"addons_config,omitempty"`
	// +optional
	ClusterIpv4CIDR string `json:"clusterIpv4CIDR,omitempty" tf:"cluster_ipv4_cidr,omitempty"`
	// +optional
	Description string `json:"description,omitempty" tf:"description,omitempty"`
	// +optional
	// Deprecated
	EnableBinaryAuthorization bool `json:"enableBinaryAuthorization,omitempty" tf:"enable_binary_authorization,omitempty"`
	// +optional
	EnableKubernetesAlpha bool `json:"enableKubernetesAlpha,omitempty" tf:"enable_kubernetes_alpha,omitempty"`
	// +optional
	EnableLegacyAbac bool `json:"enableLegacyAbac,omitempty" tf:"enable_legacy_abac,omitempty"`
	// +optional
	// Deprecated
	EnableTpu bool `json:"enableTpu,omitempty" tf:"enable_tpu,omitempty"`
	// +optional
	InitialNodeCount int `json:"initialNodeCount,omitempty" tf:"initial_node_count,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	IpAllocationPolicy []ContainerClusterSpecIpAllocationPolicy `json:"ipAllocationPolicy,omitempty" tf:"ip_allocation_policy,omitempty"`
	// +optional
	LoggingService string `json:"loggingService,omitempty" tf:"logging_service,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	MaintenancePolicy []ContainerClusterSpecMaintenancePolicy `json:"maintenancePolicy,omitempty" tf:"maintenance_policy,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	MasterAuth []ContainerClusterSpecMasterAuth `json:"masterAuth,omitempty" tf:"master_auth,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	MasterAuthorizedNetworksConfig []ContainerClusterSpecMasterAuthorizedNetworksConfig `json:"masterAuthorizedNetworksConfig,omitempty" tf:"master_authorized_networks_config,omitempty"`
	// +optional
	// Deprecated
	MasterIpv4CIDRBlock string `json:"masterIpv4CIDRBlock,omitempty" tf:"master_ipv4_cidr_block,omitempty"`
	// +optional
	MinMasterVersion string `json:"minMasterVersion,omitempty" tf:"min_master_version,omitempty"`
	// +optional
	MonitoringService string `json:"monitoringService,omitempty" tf:"monitoring_service,omitempty"`
	Name              string `json:"name" tf:"name"`
	// +optional
	Network string `json:"network,omitempty" tf:"network,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	NetworkPolicy []ContainerClusterSpecNetworkPolicy `json:"networkPolicy,omitempty" tf:"network_policy,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	NodeConfig []ContainerClusterSpecNodeConfig `json:"nodeConfig,omitempty" tf:"node_config,omitempty"`
	// +optional
	NodePool []ContainerClusterSpecNodePool `json:"nodePool,omitempty" tf:"node_pool,omitempty"`
	// +optional
	NodeVersion string `json:"nodeVersion,omitempty" tf:"node_version,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	// Deprecated
	PodSecurityPolicyConfig []ContainerClusterSpecPodSecurityPolicyConfig `json:"podSecurityPolicyConfig,omitempty" tf:"pod_security_policy_config,omitempty"`
	// +optional
	// Deprecated
	PrivateCluster bool `json:"privateCluster,omitempty" tf:"private_cluster,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	PrivateClusterConfig []ContainerClusterSpecPrivateClusterConfig `json:"privateClusterConfig,omitempty" tf:"private_cluster_config,omitempty"`
	// +optional
	Project string `json:"project,omitempty" tf:"project,omitempty"`
	// +optional
	Region string `json:"region,omitempty" tf:"region,omitempty"`
	// +optional
	RemoveDefaultNodePool bool `json:"removeDefaultNodePool,omitempty" tf:"remove_default_node_pool,omitempty"`
	// +optional
	ResourceLabels map[string]string `json:"resourceLabels,omitempty" tf:"resource_labels,omitempty"`
	// +optional
	Subnetwork string `json:"subnetwork,omitempty" tf:"subnetwork,omitempty"`
	// +optional
	Zone        string                    `json:"zone,omitempty" tf:"zone,omitempty"`
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`
}

type ContainerClusterStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	TFState     *runtime.RawExtension `json:"tfState,omitempty"`
	TFStateHash string                `json:"tfStateHash,omitempty"`
	Output      *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// ContainerClusterList is a list of ContainerClusters
type ContainerClusterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of ContainerCluster CRD objects
	Items []ContainerCluster `json:"items,omitempty"`
}
