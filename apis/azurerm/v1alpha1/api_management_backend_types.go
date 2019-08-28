package v1alpha1

import (
    "encoding/json"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	core "k8s.io/api/core/v1"
	"kubeform.dev/kubeform/apis"
)

// +genclient
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
	Parameter string `json:"parameter,omitempty" tf:"parameter,omitempty"`
	// +optional
	Scheme string `json:"scheme,omitempty" tf:"scheme,omitempty"`
}

type ApiManagementBackendSpecCredentials struct {
	// +optional
	// +kubebuilder:validation:MaxItems=1
	Authorization []ApiManagementBackendSpecCredentialsAuthorization `json:"authorization,omitempty" tf:"authorization,omitempty"`
	// +optional
	Certificate []string `json:"certificate,omitempty" tf:"certificate,omitempty"`
	// +optional
	Header map[string]string `json:"header,omitempty" tf:"header,omitempty"`
	// +optional
	Query map[string]string `json:"query,omitempty" tf:"query,omitempty"`
}

type ApiManagementBackendSpecProxy struct {
	// +optional
	Password string `json:"-" sensitive:"true" tf:"password,omitempty"`
	Url      string `json:"url" tf:"url"`
	Username string `json:"username" tf:"username"`
}

type ApiManagementBackendSpecServiceFabricClusterServerX509Name struct {
	IssuerCertificateThumbprint string `json:"issuerCertificateThumbprint" tf:"issuer_certificate_thumbprint"`
	Name                        string `json:"name" tf:"name"`
}

type ApiManagementBackendSpecServiceFabricCluster struct {
	ClientCertificateThumbprint string `json:"clientCertificateThumbprint" tf:"client_certificate_thumbprint"`
	// +kubebuilder:validation:UniqueItems=true
	ManagementEndpoints           []string `json:"managementEndpoints" tf:"management_endpoints"`
	MaxPartitionResolutionRetries int      `json:"maxPartitionResolutionRetries" tf:"max_partition_resolution_retries"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	ServerCertificateThumbprints []string `json:"serverCertificateThumbprints,omitempty" tf:"server_certificate_thumbprints,omitempty"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	ServerX509Name []ApiManagementBackendSpecServiceFabricClusterServerX509Name `json:"serverX509Name,omitempty" tf:"server_x509_name,omitempty"`
}

type ApiManagementBackendSpecTls struct {
	// +optional
	ValidateCertificateChain bool `json:"validateCertificateChain,omitempty" tf:"validate_certificate_chain,omitempty"`
	// +optional
	ValidateCertificateName bool `json:"validateCertificateName,omitempty" tf:"validate_certificate_name,omitempty"`
}

type ApiManagementBackendSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	SecretRef *core.LocalObjectReference `json:"secretRef,omitempty" tf:"-"`

	ApiManagementName string `json:"apiManagementName" tf:"api_management_name"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	Credentials []ApiManagementBackendSpecCredentials `json:"credentials,omitempty" tf:"credentials,omitempty"`
	// +optional
	Description string `json:"description,omitempty" tf:"description,omitempty"`
	Name        string `json:"name" tf:"name"`
	Protocol    string `json:"protocol" tf:"protocol"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	Proxy             []ApiManagementBackendSpecProxy `json:"proxy,omitempty" tf:"proxy,omitempty"`
	ResourceGroupName string                          `json:"resourceGroupName" tf:"resource_group_name"`
	// +optional
	ResourceID string `json:"resourceID,omitempty" tf:"resource_id,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	ServiceFabricCluster []ApiManagementBackendSpecServiceFabricCluster `json:"serviceFabricCluster,omitempty" tf:"service_fabric_cluster,omitempty"`
	// +optional
	Title string `json:"title,omitempty" tf:"title,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	Tls []ApiManagementBackendSpecTls `json:"tls,omitempty" tf:"tls,omitempty"`
	Url string                        `json:"url" tf:"url"`
}



type ApiManagementBackendStatus struct {
    // Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64  `json:"observedGeneration,omitempty"`
	// +optional
    Output *ApiManagementBackendSpec `json:"output,omitempty"`
    // +optional
    State *apis.State `json:"state,omitempty"`
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