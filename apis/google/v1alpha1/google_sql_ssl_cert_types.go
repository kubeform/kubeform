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

type GoogleSqlSslCert struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              GoogleSqlSslCertSpec   `json:"spec,omitempty"`
	Status            GoogleSqlSslCertStatus `json:"status,omitempty"`
}

type GoogleSqlSslCertSpec struct {
	ServerCaCert     string `json:"server_ca_cert"`
	Instance         string `json:"instance"`
	CertSerialNumber string `json:"cert_serial_number"`
	ExpirationTime   string `json:"expiration_time"`
	PrivateKey       string `json:"private_key"`
	Sha1Fingerprint  string `json:"sha1_fingerprint"`
	CommonName       string `json:"common_name"`
	Cert             string `json:"cert"`
	CreateTime       string `json:"create_time"`
}



type GoogleSqlSslCertStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// GoogleSqlSslCertList is a list of GoogleSqlSslCerts
type GoogleSqlSslCertList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of GoogleSqlSslCert CRD objects
	Items []GoogleSqlSslCert `json:"items,omitempty"`
}