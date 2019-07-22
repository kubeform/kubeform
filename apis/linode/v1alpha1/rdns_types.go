package v1alpha1

import (
	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

type Rdns struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              RdnsSpec   `json:"spec,omitempty"`
	Status            RdnsStatus `json:"status,omitempty"`
}

type RdnsSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`
	Address     string                    `json:"address" tf:"address"`
	Rdns        string                    `json:"rdns" tf:"rdns"`
}

type RdnsStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	TFState *runtime.RawExtension `json:"tfState,omitempty"`
	Output  *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// RdnsList is a list of Rdnss
type RdnsList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of Rdns CRD objects
	Items []Rdns `json:"items,omitempty"`
}
