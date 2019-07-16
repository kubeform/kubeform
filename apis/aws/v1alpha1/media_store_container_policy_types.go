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

type MediaStoreContainerPolicy struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              MediaStoreContainerPolicySpec   `json:"spec,omitempty"`
	Status            MediaStoreContainerPolicyStatus `json:"status,omitempty"`
}

type MediaStoreContainerPolicySpec struct {
	ContainerName string `json:"container_name"`
	Policy        string `json:"policy"`
}

type MediaStoreContainerPolicyStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// MediaStoreContainerPolicyList is a list of MediaStoreContainerPolicys
type MediaStoreContainerPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of MediaStoreContainerPolicy CRD objects
	Items []MediaStoreContainerPolicy `json:"items,omitempty"`
}
