package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type GoogleSqlSslCert struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              GoogleSqlSslCertSpec   `json:"spec,omitempty"`
	Status            GoogleSqlSslCertStatus `json:"status,omitempty"`
}

type GoogleSqlSslCertSpec struct {
	CreateTime       string `json:"create_time"`
	PrivateKey       string `json:"private_key"`
	ServerCaCert     string `json:"server_ca_cert"`
	CommonName       string `json:"common_name"`
	Cert             string `json:"cert"`
	ExpirationTime   string `json:"expiration_time"`
	Sha1Fingerprint  string `json:"sha1_fingerprint"`
	Instance         string `json:"instance"`
	CertSerialNumber string `json:"cert_serial_number"`
}

type GoogleSqlSslCertStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// GoogleSqlSslCertList is a list of GoogleSqlSslCerts
type GoogleSqlSslCertList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of GoogleSqlSslCert CRD objects
	Items []GoogleSqlSslCert `json:"items,omitempty"`
}
