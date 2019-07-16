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

type DevTestVirtualNetwork struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DevTestVirtualNetworkSpec   `json:"spec,omitempty"`
	Status            DevTestVirtualNetworkStatus `json:"status,omitempty"`
}

type DevTestVirtualNetworkSpec struct {
	// +optional
	Description       string `json:"description,omitempty"`
	LabName           string `json:"lab_name"`
	Name              string `json:"name"`
	ResourceGroupName string `json:"resource_group_name"`
}

type DevTestVirtualNetworkStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// DevTestVirtualNetworkList is a list of DevTestVirtualNetworks
type DevTestVirtualNetworkList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of DevTestVirtualNetwork CRD objects
	Items []DevTestVirtualNetwork `json:"items,omitempty"`
}
