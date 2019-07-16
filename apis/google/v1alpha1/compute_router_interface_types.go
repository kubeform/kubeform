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

type ComputeRouterInterface struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ComputeRouterInterfaceSpec   `json:"spec,omitempty"`
	Status            ComputeRouterInterfaceStatus `json:"status,omitempty"`
}

type ComputeRouterInterfaceSpec struct {
	// +optional
	IpRange   string `json:"ip_range,omitempty"`
	Name      string `json:"name"`
	Router    string `json:"router"`
	VpnTunnel string `json:"vpn_tunnel"`
}

type ComputeRouterInterfaceStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// ComputeRouterInterfaceList is a list of ComputeRouterInterfaces
type ComputeRouterInterfaceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of ComputeRouterInterface CRD objects
	Items []ComputeRouterInterface `json:"items,omitempty"`
}
