package v1alpha1

import (
	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"kubeform.dev/kubeform/apis"
)

// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

type KubernetesCluster struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              KubernetesClusterSpec   `json:"spec,omitempty"`
	Status            KubernetesClusterStatus `json:"status,omitempty"`
}

type KubernetesClusterSpecAddonProfileAciConnectorLinux struct {
	Enabled    bool   `json:"enabled" tf:"enabled"`
	SubnetName string `json:"subnetName" tf:"subnet_name"`
}

type KubernetesClusterSpecAddonProfileHttpApplicationRouting struct {
	Enabled bool `json:"enabled" tf:"enabled"`
	// +optional
	HttpApplicationRoutingZoneName string `json:"httpApplicationRoutingZoneName,omitempty" tf:"http_application_routing_zone_name,omitempty"`
}

type KubernetesClusterSpecAddonProfileOmsAgent struct {
	Enabled                 bool   `json:"enabled" tf:"enabled"`
	LogAnalyticsWorkspaceID string `json:"logAnalyticsWorkspaceID" tf:"log_analytics_workspace_id"`
}

type KubernetesClusterSpecAddonProfile struct {
	// +optional
	// +kubebuilder:validation:MaxItems=1
	AciConnectorLinux []KubernetesClusterSpecAddonProfileAciConnectorLinux `json:"aciConnectorLinux,omitempty" tf:"aci_connector_linux,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	HttpApplicationRouting []KubernetesClusterSpecAddonProfileHttpApplicationRouting `json:"httpApplicationRouting,omitempty" tf:"http_application_routing,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	OmsAgent []KubernetesClusterSpecAddonProfileOmsAgent `json:"omsAgent,omitempty" tf:"oms_agent,omitempty"`
}

type KubernetesClusterSpecAgentPoolProfile struct {
	// +optional
	Count int `json:"count,omitempty" tf:"count,omitempty"`
	// +optional
	// Deprecated
	DnsPrefix string `json:"dnsPrefix,omitempty" tf:"dns_prefix,omitempty"`
	// +optional
	// Deprecated
	Fqdn string `json:"fqdn,omitempty" tf:"fqdn,omitempty"`
	// +optional
	MaxPods int    `json:"maxPods,omitempty" tf:"max_pods,omitempty"`
	Name    string `json:"name" tf:"name"`
	// +optional
	OsDiskSizeGb int `json:"osDiskSizeGb,omitempty" tf:"os_disk_size_gb,omitempty"`
	// +optional
	OsType string `json:"osType,omitempty" tf:"os_type,omitempty"`
	// +optional
	Type   string `json:"type,omitempty" tf:"type,omitempty"`
	VmSize string `json:"vmSize" tf:"vm_size"`
	// +optional
	VnetSubnetID string `json:"vnetSubnetID,omitempty" tf:"vnet_subnet_id,omitempty"`
}

type KubernetesClusterSpecKubeAdminConfig struct {
	// +optional
	ClientCertificate string `json:"clientCertificate,omitempty" tf:"client_certificate,omitempty"`
	// +optional
	ClientKey string `json:"-" sensitive:"true" tf:"client_key,omitempty"`
	// +optional
	ClusterCaCertificate string `json:"clusterCaCertificate,omitempty" tf:"cluster_ca_certificate,omitempty"`
	// +optional
	Host string `json:"host,omitempty" tf:"host,omitempty"`
	// +optional
	Password string `json:"-" sensitive:"true" tf:"password,omitempty"`
	// +optional
	Username string `json:"username,omitempty" tf:"username,omitempty"`
}

type KubernetesClusterSpecKubeConfig struct {
	// +optional
	ClientCertificate string `json:"clientCertificate,omitempty" tf:"client_certificate,omitempty"`
	// +optional
	ClientKey string `json:"-" sensitive:"true" tf:"client_key,omitempty"`
	// +optional
	ClusterCaCertificate string `json:"clusterCaCertificate,omitempty" tf:"cluster_ca_certificate,omitempty"`
	// +optional
	Host string `json:"host,omitempty" tf:"host,omitempty"`
	// +optional
	Password string `json:"-" sensitive:"true" tf:"password,omitempty"`
	// +optional
	Username string `json:"username,omitempty" tf:"username,omitempty"`
}

type KubernetesClusterSpecLinuxProfileSshKey struct {
	KeyData string `json:"keyData" tf:"key_data"`
}

type KubernetesClusterSpecLinuxProfile struct {
	AdminUsername string `json:"adminUsername" tf:"admin_username"`
	// +kubebuilder:validation:MaxItems=1
	SshKey []KubernetesClusterSpecLinuxProfileSshKey `json:"sshKey" tf:"ssh_key"`
}

type KubernetesClusterSpecNetworkProfile struct {
	// +optional
	DnsServiceIP string `json:"dnsServiceIP,omitempty" tf:"dns_service_ip,omitempty"`
	// +optional
	DockerBridgeCIDR string `json:"dockerBridgeCIDR,omitempty" tf:"docker_bridge_cidr,omitempty"`
	NetworkPlugin    string `json:"networkPlugin" tf:"network_plugin"`
	// +optional
	NetworkPolicy string `json:"networkPolicy,omitempty" tf:"network_policy,omitempty"`
	// +optional
	PodCIDR string `json:"podCIDR,omitempty" tf:"pod_cidr,omitempty"`
	// +optional
	ServiceCIDR string `json:"serviceCIDR,omitempty" tf:"service_cidr,omitempty"`
}

type KubernetesClusterSpecRoleBasedAccessControlAzureActiveDirectory struct {
	ClientAppID     string `json:"clientAppID" tf:"client_app_id"`
	ServerAppID     string `json:"serverAppID" tf:"server_app_id"`
	ServerAppSecret string `json:"-" sensitive:"true" tf:"server_app_secret"`
	// +optional
	TenantID string `json:"tenantID,omitempty" tf:"tenant_id,omitempty"`
}

type KubernetesClusterSpecRoleBasedAccessControl struct {
	// +optional
	// +kubebuilder:validation:MaxItems=1
	AzureActiveDirectory []KubernetesClusterSpecRoleBasedAccessControlAzureActiveDirectory `json:"azureActiveDirectory,omitempty" tf:"azure_active_directory,omitempty"`
	Enabled              bool                                                              `json:"enabled" tf:"enabled"`
}

type KubernetesClusterSpecServicePrincipal struct {
	ClientID     string `json:"clientID" tf:"client_id"`
	ClientSecret string `json:"-" sensitive:"true" tf:"client_secret"`
}

type KubernetesClusterSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	KubeFormSecret *core.LocalObjectReference `json:"secret,omitempty" tf:"-"`

	// +optional
	// +kubebuilder:validation:MaxItems=1
	AddonProfile     []KubernetesClusterSpecAddonProfile     `json:"addonProfile,omitempty" tf:"addon_profile,omitempty"`
	AgentPoolProfile []KubernetesClusterSpecAgentPoolProfile `json:"agentPoolProfile" tf:"agent_pool_profile"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	ApiServerAuthorizedIPRanges []string `json:"apiServerAuthorizedIPRanges,omitempty" tf:"api_server_authorized_ip_ranges,omitempty"`
	DnsPrefix                   string   `json:"dnsPrefix" tf:"dns_prefix"`
	// +optional
	Fqdn string `json:"fqdn,omitempty" tf:"fqdn,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	KubeAdminConfig []KubernetesClusterSpecKubeAdminConfig `json:"kubeAdminConfig,omitempty" tf:"kube_admin_config,omitempty"`
	// +optional
	KubeAdminConfigRaw string `json:"-" sensitive:"true" tf:"kube_admin_config_raw,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	KubeConfig []KubernetesClusterSpecKubeConfig `json:"kubeConfig,omitempty" tf:"kube_config,omitempty"`
	// +optional
	KubeConfigRaw string `json:"-" sensitive:"true" tf:"kube_config_raw,omitempty"`
	// +optional
	KubernetesVersion string `json:"kubernetesVersion,omitempty" tf:"kubernetes_version,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	LinuxProfile []KubernetesClusterSpecLinuxProfile `json:"linuxProfile,omitempty" tf:"linux_profile,omitempty"`
	Location     string                              `json:"location" tf:"location"`
	Name         string                              `json:"name" tf:"name"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	NetworkProfile []KubernetesClusterSpecNetworkProfile `json:"networkProfile,omitempty" tf:"network_profile,omitempty"`
	// +optional
	NodeResourceGroup string `json:"nodeResourceGroup,omitempty" tf:"node_resource_group,omitempty"`
	ResourceGroupName string `json:"resourceGroupName" tf:"resource_group_name"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	RoleBasedAccessControl []KubernetesClusterSpecRoleBasedAccessControl `json:"roleBasedAccessControl,omitempty" tf:"role_based_access_control,omitempty"`
	// +kubebuilder:validation:MaxItems=1
	// +kubebuilder:validation:UniqueItems=true
	ServicePrincipal []KubernetesClusterSpecServicePrincipal `json:"servicePrincipal" tf:"service_principal"`
	// +optional
	Tags map[string]string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type KubernetesClusterStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *KubernetesClusterSpec `json:"output,omitempty"`
	// +optional
	State *apis.State `json:"state,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// KubernetesClusterList is a list of KubernetesClusters
type KubernetesClusterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of KubernetesCluster CRD objects
	Items []KubernetesCluster `json:"items,omitempty"`
}
