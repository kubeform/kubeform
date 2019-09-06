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

type SqlSSLCert struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SqlSSLCertSpec   `json:"spec,omitempty"`
	Status            SqlSSLCertStatus `json:"status,omitempty"`
}

type SqlSSLCertSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	SecretRef *core.LocalObjectReference `json:"secretRef,omitempty" tf:"-"`

	// +optional
	Cert string `json:"cert,omitempty" tf:"cert,omitempty"`
	// +optional
	CertSerialNumber string `json:"certSerialNumber,omitempty" tf:"cert_serial_number,omitempty"`
	CommonName       string `json:"commonName" tf:"common_name"`
	// +optional
	CreateTime string `json:"createTime,omitempty" tf:"create_time,omitempty"`
	// +optional
	ExpirationTime string `json:"expirationTime,omitempty" tf:"expiration_time,omitempty"`
	Instance       string `json:"instance" tf:"instance"`
	// +optional
	PrivateKey string `json:"-" sensitive:"true" tf:"private_key,omitempty"`
	// +optional
	ServerCaCert string `json:"serverCaCert,omitempty" tf:"server_ca_cert,omitempty"`
	// +optional
	Sha1Fingerprint string `json:"sha1Fingerprint,omitempty" tf:"sha1_fingerprint,omitempty"`
}

type SqlSSLCertStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *SqlSSLCertSpec `json:"output,omitempty"`
	// +optional
	State *apis.State `json:"state,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// SqlSSLCertList is a list of SqlSSLCerts
type SqlSSLCertList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of SqlSSLCert CRD objects
	Items []SqlSSLCert `json:"items,omitempty"`
}
