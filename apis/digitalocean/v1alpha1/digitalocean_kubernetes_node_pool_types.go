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

type DigitaloceanKubernetesNodePool struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DigitaloceanKubernetesNodePoolSpec   `json:"spec,omitempty"`
	Status            DigitaloceanKubernetesNodePoolStatus `json:"status,omitempty"`
}

type DigitaloceanKubernetesNodePoolSpecNodes struct {
	Id        string `json:"id"`
	Name      string `json:"name"`
	Status    string `json:"status"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type DigitaloceanKubernetesNodePoolSpec struct {
	Name      string                               `json:"name"`
	Size      string                               `json:"size"`
	NodeCount int                                  `json:"node_count"`
	Tags      []string                             `json:"tags"`
	Nodes     []DigitaloceanKubernetesNodePoolSpec `json:"nodes"`
	ClusterId string                               `json:"cluster_id"`
}

type DigitaloceanKubernetesNodePoolStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// DigitaloceanKubernetesNodePoolList is a list of DigitaloceanKubernetesNodePools
type DigitaloceanKubernetesNodePoolList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of DigitaloceanKubernetesNodePool CRD objects
	Items []DigitaloceanKubernetesNodePool `json:"items,omitempty"`
}
