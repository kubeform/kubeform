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

type SesTemplate struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SesTemplateSpec   `json:"spec,omitempty"`
	Status            SesTemplateStatus `json:"status,omitempty"`
}

type SesTemplateSpec struct {
	// +optional
	Html string `json:"html,omitempty" tf:"html,omitempty"`
	Name string `json:"name" tf:"name"`
	// +optional
	Subject string `json:"subject,omitempty" tf:"subject,omitempty"`
	// +optional
	Text        string                    `json:"text,omitempty" tf:"text,omitempty"`
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`
}

type SesTemplateStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	TFState *runtime.RawExtension `json:"tfState,omitempty"`
	Output  *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// SesTemplateList is a list of SesTemplates
type SesTemplateList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of SesTemplate CRD objects
	Items []SesTemplate `json:"items,omitempty"`
}
