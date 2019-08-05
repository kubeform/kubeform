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

type Record struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              RecordSpec   `json:"spec,omitempty"`
	Status            RecordStatus `json:"status,omitempty"`
}

type RecordSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	Domain string `json:"domain" tf:"domain"`
	// +optional
	Flags int `json:"flags,omitempty" tf:"flags,omitempty"`
	// +optional
	Fqdn string `json:"fqdn,omitempty" tf:"fqdn,omitempty"`
	Name string `json:"name" tf:"name"`
	// +optional
	Port int `json:"port,omitempty" tf:"port,omitempty"`
	// +optional
	Priority int `json:"priority,omitempty" tf:"priority,omitempty"`
	// +optional
	Tag string `json:"tag,omitempty" tf:"tag,omitempty"`
	// +optional
	Ttl   int    `json:"ttl,omitempty" tf:"ttl,omitempty"`
	Type  string `json:"type" tf:"type"`
	Value string `json:"value" tf:"value"`
	// +optional
	Weight int `json:"weight,omitempty" tf:"weight,omitempty"`
}

type RecordStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *RecordSpec `json:"output,omitempty"`
	// +optional
	State *apis.State `json:"state,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// RecordList is a list of Records
type RecordList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of Record CRD objects
	Items []Record `json:"items,omitempty"`
}
