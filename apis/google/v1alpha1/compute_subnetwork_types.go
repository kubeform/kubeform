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

type ComputeSubnetwork struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ComputeSubnetworkSpec   `json:"spec,omitempty"`
	Status            ComputeSubnetworkStatus `json:"status,omitempty"`
}

type ComputeSubnetworkSpec struct {
	// +optional
	Description string `json:"description,omitempty"`
	// +optional
	EnableFlowLogs bool   `json:"enable_flow_logs,omitempty"`
	IpCidrRange    string `json:"ip_cidr_range"`
	Name           string `json:"name"`
	Network        string `json:"network"`
	// +optional
	PrivateIpGoogleAccess bool `json:"private_ip_google_access,omitempty"`
}

type ComputeSubnetworkStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// ComputeSubnetworkList is a list of ComputeSubnetworks
type ComputeSubnetworkList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of ComputeSubnetwork CRD objects
	Items []ComputeSubnetwork `json:"items,omitempty"`
}
