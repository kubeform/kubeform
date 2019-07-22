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

type LbListenerCertificate struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              LbListenerCertificateSpec   `json:"spec,omitempty"`
	Status            LbListenerCertificateStatus `json:"status,omitempty"`
}

type LbListenerCertificateSpec struct {
	ProviderRef    core.LocalObjectReference `json:"providerRef" tf:"-"`
	CertificateArn string                    `json:"certificateArn" tf:"certificate_arn"`
	ListenerArn    string                    `json:"listenerArn" tf:"listener_arn"`
}

type LbListenerCertificateStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	TFState *runtime.RawExtension `json:"tfState,omitempty"`
	Output  *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// LbListenerCertificateList is a list of LbListenerCertificates
type LbListenerCertificateList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of LbListenerCertificate CRD objects
	Items []LbListenerCertificate `json:"items,omitempty"`
}
