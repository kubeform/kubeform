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

type GoogleComputeTargetTcpProxy struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              GoogleComputeTargetTcpProxySpec   `json:"spec,omitempty"`
	Status            GoogleComputeTargetTcpProxyStatus `json:"status,omitempty"`
}

type GoogleComputeTargetTcpProxySpec struct {
	BackendService    string `json:"backend_service"`
	Name              string `json:"name"`
	Description       string `json:"description"`
	ProxyHeader       string `json:"proxy_header"`
	CreationTimestamp string `json:"creation_timestamp"`
	ProxyId           int    `json:"proxy_id"`
	Project           string `json:"project"`
	SelfLink          string `json:"self_link"`
}



type GoogleComputeTargetTcpProxyStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// GoogleComputeTargetTcpProxyList is a list of GoogleComputeTargetTcpProxys
type GoogleComputeTargetTcpProxyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of GoogleComputeTargetTcpProxy CRD objects
	Items []GoogleComputeTargetTcpProxy `json:"items,omitempty"`
}