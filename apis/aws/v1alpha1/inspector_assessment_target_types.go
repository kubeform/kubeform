package v1alpha1

import (
	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	base "kubeform.dev/kubeform/apis/base/v1alpha1"
)

// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

type InspectorAssessmentTarget struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              InspectorAssessmentTargetSpec   `json:"spec,omitempty"`
	Status            InspectorAssessmentTargetStatus `json:"status,omitempty"`
}

type InspectorAssessmentTargetSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	Arn  string `json:"arn,omitempty" tf:"arn,omitempty"`
	Name string `json:"name" tf:"name"`
	// +optional
	ResourceGroupArn string `json:"resourceGroupArn,omitempty" tf:"resource_group_arn,omitempty"`
}

type InspectorAssessmentTargetStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *InspectorAssessmentTargetSpec `json:"output,omitempty"`
	// +optional
	State *base.State `json:"state,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// InspectorAssessmentTargetList is a list of InspectorAssessmentTargets
type InspectorAssessmentTargetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of InspectorAssessmentTarget CRD objects
	Items []InspectorAssessmentTarget `json:"items,omitempty"`
}
