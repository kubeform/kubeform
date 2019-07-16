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

type ServiceDiscoveryHttpNamespace struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ServiceDiscoveryHttpNamespaceSpec   `json:"spec,omitempty"`
	Status            ServiceDiscoveryHttpNamespaceStatus `json:"status,omitempty"`
}

type ServiceDiscoveryHttpNamespaceSpec struct {
	// +optional
	Description string `json:"description,omitempty"`
	Name        string `json:"name"`
}

type ServiceDiscoveryHttpNamespaceStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// ServiceDiscoveryHttpNamespaceList is a list of ServiceDiscoveryHttpNamespaces
type ServiceDiscoveryHttpNamespaceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of ServiceDiscoveryHttpNamespace CRD objects
	Items []ServiceDiscoveryHttpNamespace `json:"items,omitempty"`
}
