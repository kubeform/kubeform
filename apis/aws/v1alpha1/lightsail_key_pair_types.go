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

type LightsailKeyPair struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              LightsailKeyPairSpec   `json:"spec,omitempty"`
	Status            LightsailKeyPairStatus `json:"status,omitempty"`
}

type LightsailKeyPairSpec struct {
	// +optional
	NamePrefix string `json:"name_prefix,omitempty"`
	// +optional
	PgpKey string `json:"pgp_key,omitempty"`
}

type LightsailKeyPairStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// LightsailKeyPairList is a list of LightsailKeyPairs
type LightsailKeyPairList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of LightsailKeyPair CRD objects
	Items []LightsailKeyPair `json:"items,omitempty"`
}
