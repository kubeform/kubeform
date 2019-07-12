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

type AzurermServiceFabricCluster struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AzurermServiceFabricClusterSpec   `json:"spec,omitempty"`
	Status            AzurermServiceFabricClusterStatus `json:"status,omitempty"`
}

type AzurermServiceFabricClusterSpecAzureActiveDirectory struct {
	TenantId             string `json:"tenant_id"`
	ClusterApplicationId string `json:"cluster_application_id"`
	ClientApplicationId  string `json:"client_application_id"`
}

type AzurermServiceFabricClusterSpecClientCertificateThumbprint struct {
	Thumbprint string `json:"thumbprint"`
	IsAdmin    bool   `json:"is_admin"`
}

type AzurermServiceFabricClusterSpecCertificate struct {
	Thumbprint          string `json:"thumbprint"`
	ThumbprintSecondary string `json:"thumbprint_secondary"`
	X509StoreName       string `json:"x509_store_name"`
}

type AzurermServiceFabricClusterSpecCertificateCommonNamesCommonNames struct {
	CertificateCommonName       string `json:"certificate_common_name"`
	CertificateIssuerThumbprint string `json:"certificate_issuer_thumbprint"`
}

type AzurermServiceFabricClusterSpecCertificateCommonNames struct {
	CommonNames   []AzurermServiceFabricClusterSpecCertificateCommonNames `json:"common_names"`
	X509StoreName string                                                  `json:"x509_store_name"`
}

type AzurermServiceFabricClusterSpecReverseProxyCertificate struct {
	X509StoreName       string `json:"x509_store_name"`
	Thumbprint          string `json:"thumbprint"`
	ThumbprintSecondary string `json:"thumbprint_secondary"`
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
	HttpEndpointPort         int                                       `json:"http_endpoint_port"`
	ApplicationPorts         []AzurermServiceFabricClusterSpecNodeType `json:"application_ports"`
	ReverseProxyEndpointPort int                                       `json:"reverse_proxy_endpoint_port"`
	DurabilityLevel          string                                    `json:"durability_level"`
	EphemeralPorts           []AzurermServiceFabricClusterSpecNodeType `json:"ephemeral_ports"`
	PlacementProperties      map[string]string                         `json:"placement_properties"`
	Capacities               map[string]string                         `json:"capacities"`
	InstanceCount            int                                       `json:"instance_count"`
	IsPrimary                bool                                      `json:"is_primary"`
	ClientEndpointPort       int                                       `json:"client_endpoint_port"`
}

type AzurermServiceFabricClusterSpecDiagnosticsConfig struct {
	ProtectedAccountKeyName string `json:"protected_account_key_name"`
	BlobEndpoint            string `json:"blob_endpoint"`
	QueueEndpoint           string `json:"queue_endpoint"`
	TableEndpoint           string `json:"table_endpoint"`
	StorageAccountName      string `json:"storage_account_name"`
}

type AzurermServiceFabricClusterSpecFabricSettings struct {
	Name       string            `json:"name"`
	Parameters map[string]string `json:"parameters"`
}

type AzurermServiceFabricClusterSpec struct {
	Name                        string                            `json:"name"`
	UpgradeMode                 string                            `json:"upgrade_mode"`
	ManagementEndpoint          string                            `json:"management_endpoint"`
	AzureActiveDirectory        []AzurermServiceFabricClusterSpec `json:"azure_active_directory"`
	ClientCertificateThumbprint []AzurermServiceFabricClusterSpec `json:"client_certificate_thumbprint"`
	ResourceGroupName           string                            `json:"resource_group_name"`
	Location                    string                            `json:"location"`
	VmImage                     string                            `json:"vm_image"`
	AddOnFeatures               []string                          `json:"add_on_features"`
	Certificate                 []AzurermServiceFabricClusterSpec `json:"certificate"`
	CertificateCommonNames      []AzurermServiceFabricClusterSpec `json:"certificate_common_names"`
	ReverseProxyCertificate     []AzurermServiceFabricClusterSpec `json:"reverse_proxy_certificate"`
	Tags                        map[string]string                 `json:"tags"`
	ClusterCodeVersion          string                            `json:"cluster_code_version"`
	NodeType                    []AzurermServiceFabricClusterSpec `json:"node_type"`
	ReliabilityLevel            string                            `json:"reliability_level"`
	DiagnosticsConfig           []AzurermServiceFabricClusterSpec `json:"diagnostics_config"`
	FabricSettings              []AzurermServiceFabricClusterSpec `json:"fabric_settings"`
	ClusterEndpoint             string                            `json:"cluster_endpoint"`
}

type AzurermServiceFabricClusterStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AzurermServiceFabricClusterList is a list of AzurermServiceFabricClusters
type AzurermServiceFabricClusterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AzurermServiceFabricCluster CRD objects
	Items []AzurermServiceFabricCluster `json:"items,omitempty"`
}
