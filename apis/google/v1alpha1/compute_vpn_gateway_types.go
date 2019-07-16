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

type ComputeVpnGateway struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ComputeVpnGatewaySpec   `json:"spec,omitempty"`
	Status            ComputeVpnGatewayStatus `json:"status,omitempty"`
}

type ComputeVpnGatewaySpec struct {
	// +optional
	Description string `json:"description,omitempty"`
	Name        string `json:"name"`
	Network     string `json:"network"`
}

type ComputeVpnGatewayStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// ComputeVpnGatewayList is a list of ComputeVpnGateways
type ComputeVpnGatewayList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of ComputeVpnGateway CRD objects
	Items []ComputeVpnGateway `json:"items,omitempty"`
}
