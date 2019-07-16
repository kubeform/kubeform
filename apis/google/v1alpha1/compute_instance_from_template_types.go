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

type ComputeInstanceFromTemplate struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ComputeInstanceFromTemplateSpec   `json:"spec,omitempty"`
	Status            ComputeInstanceFromTemplateStatus `json:"status,omitempty"`
}

type ComputeInstanceFromTemplateSpec struct {
	Name                   string `json:"name"`
	SourceInstanceTemplate string `json:"source_instance_template"`
}

type ComputeInstanceFromTemplateStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// ComputeInstanceFromTemplateList is a list of ComputeInstanceFromTemplates
type ComputeInstanceFromTemplateList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of ComputeInstanceFromTemplate CRD objects
	Items []ComputeInstanceFromTemplate `json:"items,omitempty"`
}
