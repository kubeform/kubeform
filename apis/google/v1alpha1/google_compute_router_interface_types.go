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

type GoogleComputeRouterInterface struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              GoogleComputeRouterInterfaceSpec   `json:"spec,omitempty"`
	Status            GoogleComputeRouterInterfaceStatus `json:"status,omitempty"`
}

type GoogleComputeRouterInterfaceSpec struct {
	Name      string `json:"name"`
	Router    string `json:"router"`
	VpnTunnel string `json:"vpn_tunnel"`
	IpRange   string `json:"ip_range"`
	Project   string `json:"project"`
	Region    string `json:"region"`
}

type GoogleComputeRouterInterfaceStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// GoogleComputeRouterInterfaceList is a list of GoogleComputeRouterInterfaces
type GoogleComputeRouterInterfaceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of GoogleComputeRouterInterface CRD objects
	Items []GoogleComputeRouterInterface `json:"items,omitempty"`
}
