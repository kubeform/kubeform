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

type KubernetesNodePool struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              KubernetesNodePoolSpec   `json:"spec,omitempty"`
	Status            KubernetesNodePoolStatus `json:"status,omitempty"`
}

type KubernetesNodePoolSpec struct {
	ClusterId string `json:"cluster_id"`
	Name      string `json:"name"`
	NodeCount int    `json:"node_count"`
	Size      string `json:"size"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	Tags []string `json:"tags,omitempty"`
}

type KubernetesNodePoolStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// KubernetesNodePoolList is a list of KubernetesNodePools
type KubernetesNodePoolList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of KubernetesNodePool CRD objects
	Items []KubernetesNodePool `json:"items,omitempty"`
}
