package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type DigitaloceanCdn struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DigitaloceanCdnSpec   `json:"spec,omitempty"`
	Status            DigitaloceanCdnStatus `json:"status,omitempty"`
}

type DigitaloceanCdnSpec struct {
	CustomDomain  string `json:"custom_domain"`
	Endpoint      string `json:"endpoint"`
	CreatedAt     string `json:"created_at"`
	Origin        string `json:"origin"`
	Ttl           int    `json:"ttl"`
	CertificateId string `json:"certificate_id"`
}

type DigitaloceanCdnStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// DigitaloceanCdnList is a list of DigitaloceanCdns
type DigitaloceanCdnList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of DigitaloceanCdn CRD objects
	Items []DigitaloceanCdn `json:"items,omitempty"`
}
