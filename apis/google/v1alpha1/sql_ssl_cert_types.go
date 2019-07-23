package v1alpha1

import (
	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
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

	CommonName string `json:"commonName" tf:"common_name"`
	Instance   string `json:"instance" tf:"instance"`
}

type SqlSslCertStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	TFState *runtime.RawExtension `json:"tfState,omitempty"`
	Output  *runtime.RawExtension `json:"output,omitempty"`
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
