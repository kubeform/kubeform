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

type ServiceFabricCluster struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ServiceFabricClusterSpec   `json:"spec,omitempty"`
	Status            ServiceFabricClusterStatus `json:"status,omitempty"`
}

type ServiceFabricClusterSpecAzureActiveDirectory struct {
	ClientApplicationID  string `json:"clientApplicationID" tf:"client_application_id"`
	ClusterApplicationID string `json:"clusterApplicationID" tf:"cluster_application_id"`
	TenantID             string `json:"tenantID" tf:"tenant_id"`
}

type ServiceFabricClusterSpecCertificate struct {
	Thumbprint string `json:"thumbprint" tf:"thumbprint"`
	// +optional
	ThumbprintSecondary string `json:"thumbprintSecondary,omitempty" tf:"thumbprint_secondary,omitempty"`
	X509StoreName       string `json:"x509StoreName" tf:"x509_store_name"`
}

type ServiceFabricClusterSpecCertificateCommonNamesCommonNames struct {
	CertificateCommonName string `json:"certificateCommonName" tf:"certificate_common_name"`
	// +optional
	CertificateIssuerThumbprint string `json:"certificateIssuerThumbprint,omitempty" tf:"certificate_issuer_thumbprint,omitempty"`
}

type ServiceFabricClusterSpecCertificateCommonNames struct {
	// +kubebuilder:validation:MinItems=1
	// +kubebuilder:validation:UniqueItems=true
	CommonNames   []ServiceFabricClusterSpecCertificateCommonNamesCommonNames `json:"commonNames" tf:"common_names"`
	X509StoreName string                                                      `json:"x509StoreName" tf:"x509_store_name"`
}

type ServiceFabricClusterSpecClientCertificateThumbprint struct {
	IsAdmin    bool   `json:"isAdmin" tf:"is_admin"`
	Thumbprint string `json:"thumbprint" tf:"thumbprint"`
}

type ServiceFabricClusterSpecDiagnosticsConfig struct {
	BlobEndpoint            string `json:"blobEndpoint" tf:"blob_endpoint"`
	ProtectedAccountKeyName string `json:"protectedAccountKeyName" tf:"protected_account_key_name"`
	QueueEndpoint           string `json:"queueEndpoint" tf:"queue_endpoint"`
	StorageAccountName      string `json:"storageAccountName" tf:"storage_account_name"`
	TableEndpoint           string `json:"tableEndpoint" tf:"table_endpoint"`
}

type ServiceFabricClusterSpecFabricSettings struct {
	Name string `json:"name" tf:"name"`
	// +optional
	Parameters map[string]string `json:"parameters,omitempty" tf:"parameters,omitempty"`
}

type ServiceFabricClusterSpecNodeTypeApplicationPorts struct {
	EndPort   int `json:"endPort" tf:"end_port"`
	StartPort int `json:"startPort" tf:"start_port"`
}

type ServiceFabricClusterSpecNodeTypeEphemeralPorts struct {
	EndPort   int `json:"endPort" tf:"end_port"`
	StartPort int `json:"startPort" tf:"start_port"`
}

type ServiceFabricClusterSpecNodeType struct {
	// +optional
	// +kubebuilder:validation:MaxItems=1
	ApplicationPorts []ServiceFabricClusterSpecNodeTypeApplicationPorts `json:"applicationPorts,omitempty" tf:"application_ports,omitempty"`
	// +optional
	Capacities         map[string]string `json:"capacities,omitempty" tf:"capacities,omitempty"`
	ClientEndpointPort int               `json:"clientEndpointPort" tf:"client_endpoint_port"`
	// +optional
	DurabilityLevel string `json:"durabilityLevel,omitempty" tf:"durability_level,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	EphemeralPorts   []ServiceFabricClusterSpecNodeTypeEphemeralPorts `json:"ephemeralPorts,omitempty" tf:"ephemeral_ports,omitempty"`
	HttpEndpointPort int                                              `json:"httpEndpointPort" tf:"http_endpoint_port"`
	InstanceCount    int                                              `json:"instanceCount" tf:"instance_count"`
	IsPrimary        bool                                             `json:"isPrimary" tf:"is_primary"`
	Name             string                                           `json:"name" tf:"name"`
	// +optional
	PlacementProperties map[string]string `json:"placementProperties,omitempty" tf:"placement_properties,omitempty"`
	// +optional
	ReverseProxyEndpointPort int `json:"reverseProxyEndpointPort,omitempty" tf:"reverse_proxy_endpoint_port,omitempty"`
}

type ServiceFabricClusterSpecReverseProxyCertificate struct {
	Thumbprint string `json:"thumbprint" tf:"thumbprint"`
	// +optional
	ThumbprintSecondary string `json:"thumbprintSecondary,omitempty" tf:"thumbprint_secondary,omitempty"`
	X509StoreName       string `json:"x509StoreName" tf:"x509_store_name"`
}

type ServiceFabricClusterSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	// +optional
	// +kubebuilder:validation:UniqueItems=true
	AddOnFeatures []string `json:"addOnFeatures,omitempty" tf:"add_on_features,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	AzureActiveDirectory []ServiceFabricClusterSpecAzureActiveDirectory `json:"azureActiveDirectory,omitempty" tf:"azure_active_directory,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	Certificate []ServiceFabricClusterSpecCertificate `json:"certificate,omitempty" tf:"certificate,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	CertificateCommonNames []ServiceFabricClusterSpecCertificateCommonNames `json:"certificateCommonNames,omitempty" tf:"certificate_common_names,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=2
	ClientCertificateThumbprint []ServiceFabricClusterSpecClientCertificateThumbprint `json:"clientCertificateThumbprint,omitempty" tf:"client_certificate_thumbprint,omitempty"`
	// +optional
	ClusterCodeVersion string `json:"clusterCodeVersion,omitempty" tf:"cluster_code_version,omitempty"`
	// +optional
	ClusterEndpoint string `json:"clusterEndpoint,omitempty" tf:"cluster_endpoint,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	DiagnosticsConfig []ServiceFabricClusterSpecDiagnosticsConfig `json:"diagnosticsConfig,omitempty" tf:"diagnostics_config,omitempty"`
	// +optional
	FabricSettings     []ServiceFabricClusterSpecFabricSettings `json:"fabricSettings,omitempty" tf:"fabric_settings,omitempty"`
	Location           string                                   `json:"location" tf:"location"`
	ManagementEndpoint string                                   `json:"managementEndpoint" tf:"management_endpoint"`
	Name               string                                   `json:"name" tf:"name"`
	NodeType           []ServiceFabricClusterSpecNodeType       `json:"nodeType" tf:"node_type"`
	ReliabilityLevel   string                                   `json:"reliabilityLevel" tf:"reliability_level"`
	ResourceGroupName  string                                   `json:"resourceGroupName" tf:"resource_group_name"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	ReverseProxyCertificate []ServiceFabricClusterSpecReverseProxyCertificate `json:"reverseProxyCertificate,omitempty" tf:"reverse_proxy_certificate,omitempty"`
	// +optional
	Tags        map[string]string `json:"tags,omitempty" tf:"tags,omitempty"`
	UpgradeMode string            `json:"upgradeMode" tf:"upgrade_mode"`
	VmImage     string            `json:"vmImage" tf:"vm_image"`
}

type ServiceFabricClusterStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *ServiceFabricClusterSpec `json:"output,omitempty"`
	// +optional
	State *apis.State `json:"state,omitempty"`
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
