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

type SqlSslCert struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SqlSslCertSpec   `json:"spec,omitempty"`
	Status            SqlSslCertStatus `json:"status,omitempty"`
}

type SqlSslCertSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	KubeFormSecret *core.LocalObjectReference `json:"secret,omitempty" tf:"-"`

	Cert             string `json:"cert" tf:"cert"`
	CertSerialNumber string `json:"certSerialNumber" tf:"cert_serial_number"`
	CommonName       string `json:"commonName" tf:"common_name"`
	CreateTime       string `json:"createTime" tf:"create_time"`
	ExpirationTime   string `json:"expirationTime" tf:"expiration_time"`
	Instance         string `json:"instance" tf:"instance"`
	PrivateKey       string `json:"-" sensitive:"true" tf:"private_key"`
	ServerCaCert     string `json:"serverCaCert" tf:"server_ca_cert"`
	Sha1Fingerprint  string `json:"sha1Fingerprint" tf:"sha1_fingerprint"`
}

type SqlSslCertStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *SqlSslCertSpec `json:"output,omitempty"`
	// +optional
	State *apis.State `json:"state,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// SqlSslCertList is a list of SqlSslCerts
type SqlSslCertList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of SqlSslCert CRD objects
	Items []SqlSslCert `json:"items,omitempty"`
}
