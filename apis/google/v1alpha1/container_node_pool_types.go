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

type ContainerNodePool struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ContainerNodePoolSpec   `json:"spec,omitempty"`
	Status            ContainerNodePoolStatus `json:"status,omitempty"`
}

type ContainerNodePoolSpecAutoscaling struct {
	MaxNodeCount int `json:"max_node_count"`
	MinNodeCount int `json:"min_node_count"`
}

type ContainerNodePoolSpec struct {
	// +optional
	// +kubebuilder:validation:MaxItems=1
	Autoscaling *[]ContainerNodePoolSpec `json:"autoscaling,omitempty"`
	Cluster     string                   `json:"cluster"`
	// +optional
	Region string `json:"region,omitempty"`
}

type ContainerNodePoolStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// ContainerNodePoolList is a list of ContainerNodePools
type ContainerNodePoolList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of ContainerNodePool CRD objects
	Items []ContainerNodePool `json:"items,omitempty"`
}
