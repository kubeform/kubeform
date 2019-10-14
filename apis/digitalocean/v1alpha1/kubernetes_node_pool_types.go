package v1alpha1

import (
	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	base "kubeform.dev/kubeform/apis/base/v1alpha1"
)

// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

type KubernetesNodePool struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              KubernetesNodePoolSpec   `json:"spec,omitempty"`
	Status            KubernetesNodePoolStatus `json:"status,omitempty"`
}

type KubernetesNodePoolSpecNodes struct {
	// +optional
	CreatedAt string `json:"createdAt,omitempty" tf:"created_at,omitempty"`
	// +optional
	ID string `json:"ID,omitempty" tf:"id,omitempty"`
	// +optional
	Name string `json:"name,omitempty" tf:"name,omitempty"`
	// +optional
	Status string `json:"status,omitempty" tf:"status,omitempty"`
	// +optional
	UpdatedAt string `json:"updatedAt,omitempty" tf:"updated_at,omitempty"`
}

type KubernetesNodePoolSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	ClusterID string `json:"clusterID" tf:"cluster_id"`
	Name      string `json:"name" tf:"name"`
	NodeCount int    `json:"nodeCount" tf:"node_count"`
	// +optional
	Nodes []KubernetesNodePoolSpecNodes `json:"nodes,omitempty" tf:"nodes,omitempty"`
	Size  string                        `json:"size" tf:"size"`
	// +optional
	Tags []string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type KubernetesNodePoolStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *KubernetesNodePoolSpec `json:"output,omitempty"`
	// +optional
	State *base.State `json:"state,omitempty"`
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
