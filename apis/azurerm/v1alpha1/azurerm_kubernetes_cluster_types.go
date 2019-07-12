package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AzurermKubernetesCluster struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AzurermKubernetesClusterSpec   `json:"spec,omitempty"`
	Status            AzurermKubernetesClusterStatus `json:"status,omitempty"`
}

type AzurermKubernetesClusterSpecAddonProfileHttpApplicationRouting struct {
	Enabled                        bool   `json:"enabled"`
	HttpApplicationRoutingZoneName string `json:"http_application_routing_zone_name"`
}

type AzurermKubernetesClusterSpecAddonProfileOmsAgent struct {
	Enabled                 bool   `json:"enabled"`
	LogAnalyticsWorkspaceId string `json:"log_analytics_workspace_id"`
}

type AzurermKubernetesClusterSpecAddonProfileAciConnectorLinux struct {
	SubnetName string `json:"subnet_name"`
	Enabled    bool   `json:"enabled"`
}

type AzurermKubernetesClusterSpecAddonProfile struct {
	HttpApplicationRouting []AzurermKubernetesClusterSpecAddonProfile `json:"http_application_routing"`
	OmsAgent               []AzurermKubernetesClusterSpecAddonProfile `json:"oms_agent"`
	AciConnectorLinux      []AzurermKubernetesClusterSpecAddonProfile `json:"aci_connector_linux"`
}

type AzurermKubernetesClusterSpecKubeAdminConfig struct {
	Password             string `json:"password"`
	ClientCertificate    string `json:"client_certificate"`
	ClientKey            string `json:"client_key"`
	ClusterCaCertificate string `json:"cluster_ca_certificate"`
	Host                 string `json:"host"`
	Username             string `json:"username"`
}

type AzurermKubernetesClusterSpecAgentPoolProfile struct {
	Name         string `json:"name"`
	Type         string `json:"type"`
	VnetSubnetId string `json:"vnet_subnet_id"`
	MaxPods      int    `json:"max_pods"`
	Count        int    `json:"count"`
	DnsPrefix    string `json:"dns_prefix"`
	Fqdn         string `json:"fqdn"`
	VmSize       string `json:"vm_size"`
	OsDiskSizeGb int    `json:"os_disk_size_gb"`
	OsType       string `json:"os_type"`
}

type AzurermKubernetesClusterSpecLinuxProfileSshKey struct {
	KeyData string `json:"key_data"`
}

type AzurermKubernetesClusterSpecLinuxProfile struct {
	AdminUsername string                                     `json:"admin_username"`
	SshKey        []AzurermKubernetesClusterSpecLinuxProfile `json:"ssh_key"`
}

type AzurermKubernetesClusterSpecNetworkProfile struct {
	NetworkPlugin    string `json:"network_plugin"`
	NetworkPolicy    string `json:"network_policy"`
	DnsServiceIp     string `json:"dns_service_ip"`
	DockerBridgeCidr string `json:"docker_bridge_cidr"`
	PodCidr          string `json:"pod_cidr"`
	ServiceCidr      string `json:"service_cidr"`
}

type AzurermKubernetesClusterSpecRoleBasedAccessControlAzureActiveDirectory struct {
	ClientAppId     string `json:"client_app_id"`
	ServerAppId     string `json:"server_app_id"`
	ServerAppSecret string `json:"server_app_secret"`
	TenantId        string `json:"tenant_id"`
}

type AzurermKubernetesClusterSpecRoleBasedAccessControl struct {
	Enabled              bool                                                 `json:"enabled"`
	AzureActiveDirectory []AzurermKubernetesClusterSpecRoleBasedAccessControl `json:"azure_active_directory"`
}

type AzurermKubernetesClusterSpecKubeConfig struct {
	Password             string `json:"password"`
	ClientCertificate    string `json:"client_certificate"`
	ClientKey            string `json:"client_key"`
	ClusterCaCertificate string `json:"cluster_ca_certificate"`
	Host                 string `json:"host"`
	Username             string `json:"username"`
}

type AzurermKubernetesClusterSpecServicePrincipal struct {
	ClientSecret string `json:"client_secret"`
	ClientId     string `json:"client_id"`
}

type AzurermKubernetesClusterSpec struct {
	AddonProfile                []AzurermKubernetesClusterSpec `json:"addon_profile"`
	Tags                        map[string]string              `json:"tags"`
	KubeAdminConfig             []AzurermKubernetesClusterSpec `json:"kube_admin_config"`
	NodeResourceGroup           string                         `json:"node_resource_group"`
	ApiServerAuthorizedIpRanges []string                       `json:"api_server_authorized_ip_ranges"`
	ResourceGroupName           string                         `json:"resource_group_name"`
	Location                    string                         `json:"location"`
	KubeAdminConfigRaw          string                         `json:"kube_admin_config_raw"`
	KubeConfigRaw               string                         `json:"kube_config_raw"`
	Name                        string                         `json:"name"`
	KubernetesVersion           string                         `json:"kubernetes_version"`
	AgentPoolProfile            []AzurermKubernetesClusterSpec `json:"agent_pool_profile"`
	LinuxProfile                []AzurermKubernetesClusterSpec `json:"linux_profile"`
	Fqdn                        string                         `json:"fqdn"`
	DnsPrefix                   string                         `json:"dns_prefix"`
	NetworkProfile              []AzurermKubernetesClusterSpec `json:"network_profile"`
	RoleBasedAccessControl      []AzurermKubernetesClusterSpec `json:"role_based_access_control"`
	KubeConfig                  []AzurermKubernetesClusterSpec `json:"kube_config"`
	ServicePrincipal            []AzurermKubernetesClusterSpec `json:"service_principal"`
}

type AzurermKubernetesClusterStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AzurermKubernetesClusterList is a list of AzurermKubernetesClusters
type AzurermKubernetesClusterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AzurermKubernetesCluster CRD objects
	Items []AzurermKubernetesCluster `json:"items,omitempty"`
}
