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

type ServiceFabricCluster struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ServiceFabricClusterSpec   `json:"spec,omitempty"`
	Status            ServiceFabricClusterStatus `json:"status,omitempty"`
}

type ServiceFabricClusterSpecAzureActiveDirectory struct {
	ClientApplicationId  string `json:"client_application_id"`
	ClusterApplicationId string `json:"cluster_application_id"`
	TenantId             string `json:"tenant_id"`
}

type ServiceFabricClusterSpecCertificate struct {
	Thumbprint string `json:"thumbprint"`
	// +optional
	ThumbprintSecondary string `json:"thumbprint_secondary,omitempty"`
	X509StoreName       string `json:"x509_store_name"`
}

type ServiceFabricClusterSpecCertificateCommonNamesCommonNames struct {
	CertificateCommonName string `json:"certificate_common_name"`
	// +optional
	CertificateIssuerThumbprint string `json:"certificate_issuer_thumbprint,omitempty"`
}

type ServiceFabricClusterSpecCertificateCommonNames struct {
	// +kubebuilder:validation:MinItems=1
	// +kubebuilder:validation:UniqueItems=true
	CommonNames   []ServiceFabricClusterSpecCertificateCommonNames `json:"common_names"`
	X509StoreName string                                           `json:"x509_store_name"`
}

type ServiceFabricClusterSpecClientCertificateThumbprint struct {
	IsAdmin    bool   `json:"is_admin"`
	Thumbprint string `json:"thumbprint"`
}

type ServiceFabricClusterSpecDiagnosticsConfig struct {
	BlobEndpoint            string `json:"blob_endpoint"`
	ProtectedAccountKeyName string `json:"protected_account_key_name"`
	QueueEndpoint           string `json:"queue_endpoint"`
	StorageAccountName      string `json:"storage_account_name"`
	TableEndpoint           string `json:"table_endpoint"`
}

type ServiceFabricClusterSpecFabricSettings struct {
	Name string `json:"name"`
	// +optional
	Parameters map[string]string `json:"parameters,omitempty"`
}

type ServiceFabricClusterSpecNodeType struct {
	// +optional
	Capacities         map[string]string `json:"capacities,omitempty"`
	ClientEndpointPort int               `json:"client_endpoint_port"`
	// +optional
	DurabilityLevel  string `json:"durability_level,omitempty"`
	HttpEndpointPort int    `json:"http_endpoint_port"`
	InstanceCount    int    `json:"instance_count"`
	IsPrimary        bool   `json:"is_primary"`
	Name             string `json:"name"`
	// +optional
	PlacementProperties map[string]string `json:"placement_properties,omitempty"`
	// +optional
	ReverseProxyEndpointPort int `json:"reverse_proxy_endpoint_port,omitempty"`
}

type ServiceFabricClusterSpecReverseProxyCertificate struct {
	Thumbprint string `json:"thumbprint"`
	// +optional
	ThumbprintSecondary string `json:"thumbprint_secondary,omitempty"`
	X509StoreName       string `json:"x509_store_name"`
}

type ServiceFabricClusterSpec struct {
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	AddOnFeatures []string `json:"add_on_features,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	AzureActiveDirectory *[]ServiceFabricClusterSpec `json:"azure_active_directory,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	Certificate *[]ServiceFabricClusterSpec `json:"certificate,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	CertificateCommonNames *[]ServiceFabricClusterSpec `json:"certificate_common_names,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=2
	ClientCertificateThumbprint *[]ServiceFabricClusterSpec `json:"client_certificate_thumbprint,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	DiagnosticsConfig *[]ServiceFabricClusterSpec `json:"diagnostics_config,omitempty"`
	// +optional
	FabricSettings     *[]ServiceFabricClusterSpec `json:"fabric_settings,omitempty"`
	Location           string                      `json:"location"`
	ManagementEndpoint string                      `json:"management_endpoint"`
	Name               string                      `json:"name"`
	NodeType           []ServiceFabricClusterSpec  `json:"node_type"`
	ReliabilityLevel   string                      `json:"reliability_level"`
	ResourceGroupName  string                      `json:"resource_group_name"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	ReverseProxyCertificate *[]ServiceFabricClusterSpec `json:"reverse_proxy_certificate,omitempty"`
	UpgradeMode             string                      `json:"upgrade_mode"`
	VmImage                 string                      `json:"vm_image"`
}

type ServiceFabricClusterStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// ServiceFabricClusterList is a list of ServiceFabricClusters
type ServiceFabricClusterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of ServiceFabricCluster CRD objects
	Items []ServiceFabricCluster `json:"items,omitempty"`
}
