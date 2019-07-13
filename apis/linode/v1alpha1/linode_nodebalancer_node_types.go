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

type LinodeNodebalancerNode struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              LinodeNodebalancerNodeSpec   `json:"spec,omitempty"`
	Status            LinodeNodebalancerNodeStatus `json:"status,omitempty"`
}

type LinodeNodebalancerNodeSpec struct {
	ConfigId       int    `json:"config_id"`
	Label          string `json:"label"`
	Weight         int    `json:"weight"`
	Mode           string `json:"mode"`
	Address        string `json:"address"`
	Status         string `json:"status"`
	NodebalancerId int    `json:"nodebalancer_id"`
}



type LinodeNodebalancerNodeStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// LinodeNodebalancerNodeList is a list of LinodeNodebalancerNodes
type LinodeNodebalancerNodeList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of LinodeNodebalancerNode CRD objects
	Items []LinodeNodebalancerNode `json:"items,omitempty"`
}