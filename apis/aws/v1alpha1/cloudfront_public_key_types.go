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

type CloudfrontPublicKey struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              CloudfrontPublicKeySpec   `json:"spec,omitempty"`
	Status            CloudfrontPublicKeyStatus `json:"status,omitempty"`
}

type CloudfrontPublicKeySpec struct {
	// +optional
	Comment    string `json:"comment,omitempty"`
	EncodedKey string `json:"encoded_key"`
}

type CloudfrontPublicKeyStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
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
