package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AzurermServiceFabricCluster struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AzurermServiceFabricClusterSpec   `json:"spec,omitempty"`
	Status            AzurermServiceFabricClusterStatus `json:"status,omitempty"`
}

type AzurermServiceFabricClusterSpecReverseProxyCertificate struct {
	Thumbprint          string `json:"thumbprint"`
	ThumbprintSecondary string `json:"thumbprint_secondary"`
	X509StoreName       string `json:"x509_store_name"`
}

type AzurermServiceFabricClusterSpecNodeTypeApplicationPorts struct {
	StartPort int `json:"start_port"`
	EndPort   int `json:"end_port"`
}

type AzurermServiceFabricClusterSpecNodeTypeEphemeralPorts struct {
	StartPort int `json:"start_port"`
	EndPort   int `json:"end_port"`
}

type AzurermServiceFabricClusterSpecNodeType struct {
	Name                     string                                    `json:"name"`
	ClientEndpointPort       int                                       `json:"client_endpoint_port"`
	HttpEndpointPort         int                                       `json:"http_endpoint_port"`
	ApplicationPorts         []AzurermServiceFabricClusterSpecNodeType `json:"application_ports"`
	EphemeralPorts           []AzurermServiceFabricClusterSpecNodeType `json:"ephemeral_ports"`
	PlacementProperties      map[string]string                         `json:"placement_properties"`
	Capacities               map[string]string                         `json:"capacities"`
	InstanceCount            int                                       `json:"instance_count"`
	IsPrimary                bool                                      `json:"is_primary"`
	ReverseProxyEndpointPort int                                       `json:"reverse_proxy_endpoint_port"`
	DurabilityLevel          string                                    `json:"durability_level"`
}

type AzurermServiceFabricClusterSpecCertificate struct {
	Thumbprint          string `json:"thumbprint"`
	ThumbprintSecondary string `json:"thumbprint_secondary"`
	X509StoreName       string `json:"x509_store_name"`
}

type AzurermServiceFabricClusterSpecAzureActiveDirectory struct {
	TenantId             string `json:"tenant_id"`
	ClusterApplicationId string `json:"cluster_application_id"`
	ClientApplicationId  string `json:"client_application_id"`
}

type AzurermServiceFabricClusterSpecFabricSettings struct {
	Name       string            `json:"name"`
	Parameters map[string]string `json:"parameters"`
}

type AzurermServiceFabricClusterSpecCertificateCommonNamesCommonNames struct {
	CertificateCommonName       string `json:"certificate_common_name"`
	CertificateIssuerThumbprint string `json:"certificate_issuer_thumbprint"`
}

type AzurermServiceFabricClusterSpecCertificateCommonNames struct {
	CommonNames   []AzurermServiceFabricClusterSpecCertificateCommonNames `json:"common_names"`
	X509StoreName string                                                  `json:"x509_store_name"`
}

type AzurermServiceFabricClusterSpecClientCertificateThumbprint struct {
	Thumbprint string `json:"thumbprint"`
	IsAdmin    bool   `json:"is_admin"`
}

type AzurermServiceFabricClusterSpecDiagnosticsConfig struct {
	StorageAccountName      string `json:"storage_account_name"`
	ProtectedAccountKeyName string `json:"protected_account_key_name"`
	BlobEndpoint            string `json:"blob_endpoint"`
	QueueEndpoint           string `json:"queue_endpoint"`
	TableEndpoint           string `json:"table_endpoint"`
}

type AzurermServiceFabricClusterSpec struct {
	ClusterCodeVersion          string                            `json:"cluster_code_version"`
	ReverseProxyCertificate     []AzurermServiceFabricClusterSpec `json:"reverse_proxy_certificate"`
	NodeType                    []AzurermServiceFabricClusterSpec `json:"node_type"`
	Tags                        map[string]string                 `json:"tags"`
	Location                    string                            `json:"location"`
	UpgradeMode                 string                            `json:"upgrade_mode"`
	Certificate                 []AzurermServiceFabricClusterSpec `json:"certificate"`
	ResourceGroupName           string                            `json:"resource_group_name"`
	AddOnFeatures               []string                          `json:"add_on_features"`
	AzureActiveDirectory        []AzurermServiceFabricClusterSpec `json:"azure_active_directory"`
	FabricSettings              []AzurermServiceFabricClusterSpec `json:"fabric_settings"`
	ClusterEndpoint             string                            `json:"cluster_endpoint"`
	VmImage                     string                            `json:"vm_image"`
	ReliabilityLevel            string                            `json:"reliability_level"`
	ManagementEndpoint          string                            `json:"management_endpoint"`
	CertificateCommonNames      []AzurermServiceFabricClusterSpec `json:"certificate_common_names"`
	ClientCertificateThumbprint []AzurermServiceFabricClusterSpec `json:"client_certificate_thumbprint"`
	DiagnosticsConfig           []AzurermServiceFabricClusterSpec `json:"diagnostics_config"`
	Name                        string                            `json:"name"`
}



type AzurermServiceFabricClusterStatus struct {
    Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AzurermServiceFabricClusterList is a list of AzurermServiceFabricClusters
type AzurermServiceFabricClusterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AzurermServiceFabricCluster CRD objects
	Items []AzurermServiceFabricCluster `json:"items,omitempty"`
}