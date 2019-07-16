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

type ComputeTargetTcpProxy struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ComputeTargetTcpProxySpec   `json:"spec,omitempty"`
	Status            ComputeTargetTcpProxyStatus `json:"status,omitempty"`
}

type ComputeTargetTcpProxySpec struct {
	BackendService string `json:"backend_service"`
	// +optional
	Description string `json:"description,omitempty"`
	Name        string `json:"name"`
	// +optional
	ProxyHeader string `json:"proxy_header,omitempty"`
}

type ComputeTargetTcpProxyStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// ComputeTargetTcpProxyList is a list of ComputeTargetTcpProxys
type ComputeTargetTcpProxyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of ComputeTargetTcpProxy CRD objects
	Items []ComputeTargetTcpProxy `json:"items,omitempty"`
}
