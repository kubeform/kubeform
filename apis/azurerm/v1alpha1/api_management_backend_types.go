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

type ApiManagementBackend struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ApiManagementBackendSpec   `json:"spec,omitempty"`
	Status            ApiManagementBackendStatus `json:"status,omitempty"`
}

type ApiManagementBackendSpecCredentialsAuthorization struct {
	// +optional
	Parameter string `json:"parameter,omitempty"`
	// +optional
	Scheme string `json:"scheme,omitempty"`
}

type ApiManagementBackendSpecCredentials struct {
	// +optional
	// +kubebuilder:validation:MaxItems=1
	Authorization *[]ApiManagementBackendSpecCredentials `json:"authorization,omitempty"`
	// +optional
	Certificate []string `json:"certificate,omitempty"`
	// +optional
	Header map[string]string `json:"header,omitempty"`
	// +optional
	Query map[string]string `json:"query,omitempty"`
}

type ApiManagementBackendSpecProxy struct {
	// +optional
	Password string `json:"password,omitempty"`
	Url      string `json:"url"`
	Username string `json:"username"`
}

type ApiManagementBackendSpecServiceFabricClusterServerX509Name struct {
	IssuerCertificateThumbprint string `json:"issuer_certificate_thumbprint"`
	Name                        string `json:"name"`
}

type ApiManagementBackendSpecServiceFabricCluster struct {
	ClientCertificateThumbprint string `json:"client_certificate_thumbprint"`
	// +kubebuilder:validation:UniqueItems=true
	ManagementEndpoints           []string `json:"management_endpoints"`
	MaxPartitionResolutionRetries int      `json:"max_partition_resolution_retries"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	ServerCertificateThumbprints []string `json:"server_certificate_thumbprints,omitempty"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	ServerX509Name *[]ApiManagementBackendSpecServiceFabricCluster `json:"server_x509_name,omitempty"`
}

type ApiManagementBackendSpecTls struct {
	// +optional
	ValidateCertificateChain bool `json:"validate_certificate_chain,omitempty"`
	// +optional
	ValidateCertificateName bool `json:"validate_certificate_name,omitempty"`
}

type ApiManagementBackendSpec struct {
	ApiManagementName string `json:"api_management_name"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	Credentials *[]ApiManagementBackendSpec `json:"credentials,omitempty"`
	// +optional
	Description string `json:"description,omitempty"`
	Name        string `json:"name"`
	Protocol    string `json:"protocol"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	Proxy             *[]ApiManagementBackendSpec `json:"proxy,omitempty"`
	ResourceGroupName string                      `json:"resource_group_name"`
	// +optional
	ResourceId string `json:"resource_id,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	ServiceFabricCluster *[]ApiManagementBackendSpec `json:"service_fabric_cluster,omitempty"`
	// +optional
	Title string `json:"title,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	Tls *[]ApiManagementBackendSpec `json:"tls,omitempty"`
	Url string                      `json:"url"`
}

type ApiManagementBackendStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// ApiManagementBackendList is a list of ApiManagementBackends
type ApiManagementBackendList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of ApiManagementBackend CRD objects
	Items []ApiManagementBackend `json:"items,omitempty"`
}
