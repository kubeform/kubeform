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

type GoogleComputeNetwork struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              GoogleComputeNetworkSpec   `json:"spec,omitempty"`
	Status            GoogleComputeNetworkStatus `json:"status,omitempty"`
}

type GoogleComputeNetworkSpec struct {
	Ipv4Range             string `json:"ipv4_range"`
	Project               string `json:"project"`
	SelfLink              string `json:"self_link"`
	Name                  string `json:"name"`
	AutoCreateSubnetworks bool   `json:"auto_create_subnetworks"`
	Description           string `json:"description"`
	RoutingMode           string `json:"routing_mode"`
	GatewayIpv4           string `json:"gateway_ipv4"`
}



type GoogleComputeNetworkStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// GoogleComputeNetworkList is a list of GoogleComputeNetworks
type GoogleComputeNetworkList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of GoogleComputeNetwork CRD objects
	Items []GoogleComputeNetwork `json:"items,omitempty"`
}