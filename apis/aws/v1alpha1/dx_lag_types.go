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

type DxLag struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DxLagSpec   `json:"spec,omitempty"`
	Status            DxLagStatus `json:"status,omitempty"`
}

type DxLagSpec struct {
	ConnectionsBandwidth string `json:"connections_bandwidth"`
	// +optional
	ForceDestroy bool   `json:"force_destroy,omitempty"`
	Location     string `json:"location"`
	Name         string `json:"name"`
	// +optional
	Tags map[string]string `json:"tags,omitempty"`
}

type DxLagStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// DxLagList is a list of DxLags
type DxLagList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of DxLag CRD objects
	Items []DxLag `json:"items,omitempty"`
}
