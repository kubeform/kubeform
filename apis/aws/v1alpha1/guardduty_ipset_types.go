package v1alpha1

import (
	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"kubeform.dev/kubeform/apis"
)

// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

type GuarddutyIpset struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              GuarddutyIpsetSpec   `json:"spec,omitempty"`
	Status            GuarddutyIpsetStatus `json:"status,omitempty"`
}

type GuarddutyIpsetSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	Activate   bool   `json:"activate" tf:"activate"`
	DetectorID string `json:"detectorID" tf:"detector_id"`
	Format     string `json:"format" tf:"format"`
	Location   string `json:"location" tf:"location"`
	Name       string `json:"name" tf:"name"`
}

type GuarddutyIpsetStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *GuarddutyIpsetSpec `json:"output,omitempty"`
	// +optional
	State *apis.State `json:"state,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// GuarddutyIpsetList is a list of GuarddutyIpsets
type GuarddutyIpsetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of GuarddutyIpset CRD objects
	Items []GuarddutyIpset `json:"items,omitempty"`
}
