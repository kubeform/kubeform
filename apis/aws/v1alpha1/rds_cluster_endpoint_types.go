package v1alpha1

import (
	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"kubeform.dev/kubeform/apis"
)

// +genclient
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
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	Arn                       string `json:"arn" tf:"arn"`
	ClusterEndpointIdentifier string `json:"clusterEndpointIdentifier" tf:"cluster_endpoint_identifier"`
	ClusterIdentifier         string `json:"clusterIdentifier" tf:"cluster_identifier"`
	CustomEndpointType        string `json:"customEndpointType" tf:"custom_endpoint_type"`
	Endpoint                  string `json:"endpoint" tf:"endpoint"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	ExcludedMembers []string `json:"excludedMembers,omitempty" tf:"excluded_members,omitempty"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	StaticMembers []string `json:"staticMembers,omitempty" tf:"static_members,omitempty"`
}

type RdsClusterEndpointStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *RdsClusterEndpointSpec `json:"output,omitempty"`
	// +optional
	State *apis.State `json:"state,omitempty"`
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
