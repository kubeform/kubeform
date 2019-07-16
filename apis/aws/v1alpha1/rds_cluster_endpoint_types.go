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

type RdsClusterEndpoint struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              RdsClusterEndpointSpec   `json:"spec,omitempty"`
	Status            RdsClusterEndpointStatus `json:"status,omitempty"`
}

type RdsClusterEndpointSpec struct {
	ClusterEndpointIdentifier string `json:"cluster_endpoint_identifier"`
	ClusterIdentifier         string `json:"cluster_identifier"`
	CustomEndpointType        string `json:"custom_endpoint_type"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	ExcludedMembers []string `json:"excluded_members,omitempty"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	StaticMembers []string `json:"static_members,omitempty"`
}

type RdsClusterEndpointStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// RdsClusterEndpointList is a list of RdsClusterEndpoints
type RdsClusterEndpointList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of RdsClusterEndpoint CRD objects
	Items []RdsClusterEndpoint `json:"items,omitempty"`
}
