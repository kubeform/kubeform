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

type AwsRdsClusterEndpoint struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsRdsClusterEndpointSpec   `json:"spec,omitempty"`
	Status            AwsRdsClusterEndpointStatus `json:"status,omitempty"`
}

type AwsRdsClusterEndpointSpec struct {
	ExcludedMembers           []string `json:"excluded_members"`
	StaticMembers             []string `json:"static_members"`
	Endpoint                  string   `json:"endpoint"`
	Arn                       string   `json:"arn"`
	ClusterEndpointIdentifier string   `json:"cluster_endpoint_identifier"`
	ClusterIdentifier         string   `json:"cluster_identifier"`
	CustomEndpointType        string   `json:"custom_endpoint_type"`
}

type AwsRdsClusterEndpointStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AwsRdsClusterEndpointList is a list of AwsRdsClusterEndpoints
type AwsRdsClusterEndpointList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsRdsClusterEndpoint CRD objects
	Items []AwsRdsClusterEndpoint `json:"items,omitempty"`
}
