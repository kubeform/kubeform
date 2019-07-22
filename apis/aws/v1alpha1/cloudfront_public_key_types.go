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

type CloudfrontPublicKey struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              CloudfrontPublicKeySpec   `json:"spec,omitempty"`
	Status            CloudfrontPublicKeyStatus `json:"status,omitempty"`
}

type CloudfrontPublicKeySpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`
	// +optional
	Comment    string `json:"comment,omitempty" tf:"comment,omitempty"`
	EncodedKey string `json:"encodedKey" tf:"encoded_key"`
	// +optional
	Name string `json:"name,omitempty" tf:"name,omitempty"`
	// +optional
	NamePrefix string `json:"namePrefix,omitempty" tf:"name_prefix,omitempty"`
}

type CloudfrontPublicKeyStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	TFState *runtime.RawExtension `json:"tfState,omitempty"`
	Output  *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// CloudfrontPublicKeyList is a list of CloudfrontPublicKeys
type CloudfrontPublicKeyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of CloudfrontPublicKey CRD objects
	Items []CloudfrontPublicKey `json:"items,omitempty"`
}
