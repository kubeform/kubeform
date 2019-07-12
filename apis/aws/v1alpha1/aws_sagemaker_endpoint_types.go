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

type AwsSagemakerEndpoint struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsSagemakerEndpointSpec   `json:"spec,omitempty"`
	Status            AwsSagemakerEndpointStatus `json:"status,omitempty"`
}

type AwsSagemakerEndpointSpec struct {
	Name               string            `json:"name"`
	EndpointConfigName string            `json:"endpoint_config_name"`
	Tags               map[string]string `json:"tags"`
	Arn                string            `json:"arn"`
}

type AwsSagemakerEndpointStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AwsSagemakerEndpointList is a list of AwsSagemakerEndpoints
type AwsSagemakerEndpointList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsSagemakerEndpoint CRD objects
	Items []AwsSagemakerEndpoint `json:"items,omitempty"`
}
