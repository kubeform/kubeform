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

type AzurermKubernetesCluster struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AzurermKubernetesClusterSpec   `json:"spec,omitempty"`
	Status            AzurermKubernetesClusterStatus `json:"status,omitempty"`
}

type AzurermKubernetesClusterSpecNetworkProfile struct {
	NetworkPlugin    string `json:"network_plugin"`
	NetworkPolicy    string `json:"network_policy"`
	DnsServiceIp     string `json:"dns_service_ip"`
	DockerBridgeCidr string `json:"docker_bridge_cidr"`
	PodCidr          string `json:"pod_cidr"`
	ServiceCidr      string `json:"service_cidr"`
}

type AzurermKubernetesClusterSpecServicePrincipal struct {
	ClientId     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
}

type AzurermKubernetesClusterSpecAddonProfileHttpApplicationRouting struct {
	Enabled                        bool   `json:"enabled"`
	HttpApplicationRoutingZoneName string `json:"http_application_routing_zone_name"`
}

type AzurermKubernetesClusterSpecAddonProfileOmsAgent struct {
	LogAnalyticsWorkspaceId string `json:"log_analytics_workspace_id"`
	Enabled                 bool   `json:"enabled"`
}

type AzurermKubernetesClusterSpecAddonProfileAciConnectorLinux struct {
	Enabled    bool   `json:"enabled"`
	SubnetName string `json:"subnet_name"`
}

type AzurermKubernetesClusterSpecAddonProfile struct {
	HttpApplicationRouting []AzurermKubernetesClusterSpecAddonProfile `json:"http_application_routing"`
	OmsAgent               []AzurermKubernetesClusterSpecAddonProfile `json:"oms_agent"`
	AciConnectorLinux      []AzurermKubernetesClusterSpecAddonProfile `json:"aci_connector_linux"`
}

type AzurermKubernetesClusterSpecKubeConfig struct {
	ClientKey            string `json:"client_key"`
	ClusterCaCertificate string `json:"cluster_ca_certificate"`
	Host                 string `json:"host"`
	Username             string `json:"username"`
	Password             string `json:"password"`
	ClientCertificate    string `json:"client_certificate"`
}

type AzurermKubernetesClusterSpecRoleBasedAccessControlAzureActiveDirectory struct {
	ServerAppId     string `json:"server_app_id"`
	ServerAppSecret string `json:"server_app_secret"`
	TenantId        string `json:"tenant_id"`
	ClientAppId     string `json:"client_app_id"`
}

type AzurermKubernetesClusterSpecRoleBasedAccessControl struct {
	Enabled              bool                                                 `json:"enabled"`
	AzureActiveDirectory []AzurermKubernetesClusterSpecRoleBasedAccessControl `json:"azure_active_directory"`
}

type AzurermKubernetesClusterSpecAgentPoolProfile struct {
	Fqdn         string `json:"fqdn"`
	VnetSubnetId string `json:"vnet_subnet_id"`
	OsDiskSizeGb int    `json:"os_disk_size_gb"`
	OsType       string `json:"os_type"`
	MaxPods      int    `json:"max_pods"`
	Name         string `json:"name"`
	Type         string `json:"type"`
	Count        int    `json:"count"`
	DnsPrefix    string `json:"dns_prefix"`
	VmSize       string `json:"vm_size"`
}

type AzurermKubernetesClusterSpecLinuxProfileSshKey struct {
	KeyData string `json:"key_data"`
}

type AzurermKubernetesClusterSpecLinuxProfile struct {
	AdminUsername string                                     `json:"admin_username"`
	SshKey        []AzurermKubernetesClusterSpecLinuxProfile `json:"ssh_key"`
}

type AzurermKubernetesClusterSpecKubeAdminConfig struct {
	Host                 string `json:"host"`
	Username             string `json:"username"`
	Password             string `json:"password"`
	ClientCertificate    string `json:"client_certificate"`
	ClientKey            string `json:"client_key"`
	ClusterCaCertificate string `json:"cluster_ca_certificate"`
}

type AzurermKubernetesClusterSpec struct {
	Location                    string                         `json:"location"`
	NetworkProfile              []AzurermKubernetesClusterSpec `json:"network_profile"`
	Name                        string                         `json:"name"`
	DnsPrefix                   string                         `json:"dns_prefix"`
	ServicePrincipal            []AzurermKubernetesClusterSpec `json:"service_principal"`
	AddonProfile                []AzurermKubernetesClusterSpec `json:"addon_profile"`
	KubeAdminConfigRaw          string                         `json:"kube_admin_config_raw"`
	KubeConfig                  []AzurermKubernetesClusterSpec `json:"kube_config"`
	KubeConfigRaw               string                         `json:"kube_config_raw"`
	ApiServerAuthorizedIpRanges []string                       `json:"api_server_authorized_ip_ranges"`
	ResourceGroupName           string                         `json:"resource_group_name"`
	RoleBasedAccessControl      []AzurermKubernetesClusterSpec `json:"role_based_access_control"`
	Tags                        map[string]string              `json:"tags"`
	Fqdn                        string                         `json:"fqdn"`
	NodeResourceGroup           string                         `json:"node_resource_group"`
	AgentPoolProfile            []AzurermKubernetesClusterSpec `json:"agent_pool_profile"`
	LinuxProfile                []AzurermKubernetesClusterSpec `json:"linux_profile"`
	KubeAdminConfig             []AzurermKubernetesClusterSpec `json:"kube_admin_config"`
	KubernetesVersion           string                         `json:"kubernetes_version"`
}



type AzurermKubernetesClusterStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AzurermKubernetesClusterList is a list of AzurermKubernetesClusters
type AzurermKubernetesClusterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AzurermKubernetesCluster CRD objects
	Items []AzurermKubernetesCluster `json:"items,omitempty"`
}