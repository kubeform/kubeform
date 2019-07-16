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

type KubernetesCluster struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              KubernetesClusterSpec   `json:"spec,omitempty"`
	Status            KubernetesClusterStatus `json:"status,omitempty"`
}

type KubernetesClusterSpecNodePool struct {
	Name      string `json:"name"`
	NodeCount int    `json:"node_count"`
	Size      string `json:"size"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	Tags []string `json:"tags,omitempty"`
}

type KubernetesClusterSpec struct {
	Name string `json:"name"`
	// +kubebuilder:validation:MaxItems=1
	// +kubebuilder:validation:MinItems=1
	NodePool []KubernetesClusterSpec `json:"node_pool"`
	Region   string                  `json:"region"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	Tags    []string `json:"tags,omitempty"`
	Version string   `json:"version"`
}

type KubernetesClusterStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// KubernetesClusterList is a list of KubernetesClusters
type KubernetesClusterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of KubernetesCluster CRD objects
	Items []KubernetesCluster `json:"items,omitempty"`
}
