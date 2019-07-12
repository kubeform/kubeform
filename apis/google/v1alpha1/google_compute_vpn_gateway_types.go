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

type GoogleComputeVpnGateway struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              GoogleComputeVpnGatewaySpec   `json:"spec,omitempty"`
	Status            GoogleComputeVpnGatewayStatus `json:"status,omitempty"`
}

type GoogleComputeVpnGatewaySpec struct {
	CreationTimestamp string `json:"creation_timestamp"`
	Project           string `json:"project"`
	SelfLink          string `json:"self_link"`
	Name              string `json:"name"`
	Network           string `json:"network"`
	Description       string `json:"description"`
	Region            string `json:"region"`
}

type GoogleComputeVpnGatewayStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// GoogleComputeVpnGatewayList is a list of GoogleComputeVpnGateways
type GoogleComputeVpnGatewayList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of GoogleComputeVpnGateway CRD objects
	Items []GoogleComputeVpnGateway `json:"items,omitempty"`
}
