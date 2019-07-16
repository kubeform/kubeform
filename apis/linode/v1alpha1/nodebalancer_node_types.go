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

type NodebalancerNode struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              NodebalancerNodeSpec   `json:"spec,omitempty"`
	Status            NodebalancerNodeStatus `json:"status,omitempty"`
}

type NodebalancerNodeSpec struct {
	Address        string `json:"address"`
	ConfigId       int    `json:"config_id"`
	Label          string `json:"label"`
	NodebalancerId int    `json:"nodebalancer_id"`
}

type NodebalancerNodeStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// NodebalancerNodeList is a list of NodebalancerNodes
type NodebalancerNodeList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of NodebalancerNode CRD objects
	Items []NodebalancerNode `json:"items,omitempty"`
}
