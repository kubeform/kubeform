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

type EbsDefaultKmsKey struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              EbsDefaultKmsKeySpec   `json:"spec,omitempty"`
	Status            EbsDefaultKmsKeyStatus `json:"status,omitempty"`
}

type EbsDefaultKmsKeySpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	KeyArn string `json:"keyArn" tf:"key_arn"`
}

type EbsDefaultKmsKeyStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	TFState *runtime.RawExtension `json:"tfState,omitempty"`
	Output  *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// EbsDefaultKmsKeyList is a list of EbsDefaultKmsKeys
type EbsDefaultKmsKeyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of EbsDefaultKmsKey CRD objects
	Items []EbsDefaultKmsKey `json:"items,omitempty"`
}
