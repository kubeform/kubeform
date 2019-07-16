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

type DxConnection struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DxConnectionSpec   `json:"spec,omitempty"`
	Status            DxConnectionStatus `json:"status,omitempty"`
}

type DxConnectionSpec struct {
	Bandwidth string `json:"bandwidth"`
	Location  string `json:"location"`
	Name      string `json:"name"`
	// +optional
	Tags map[string]string `json:"tags,omitempty"`
}

type DxConnectionStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// DxConnectionList is a list of DxConnections
type DxConnectionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of DxConnection CRD objects
	Items []DxConnection `json:"items,omitempty"`
}
