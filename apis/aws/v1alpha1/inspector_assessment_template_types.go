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

type InspectorAssessmentTemplate struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              InspectorAssessmentTemplateSpec   `json:"spec,omitempty"`
	Status            InspectorAssessmentTemplateStatus `json:"status,omitempty"`
}

type InspectorAssessmentTemplateSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	Arn              string   `json:"arn,omitempty" tf:"arn,omitempty"`
	Duration         int      `json:"duration" tf:"duration"`
	Name             string   `json:"name" tf:"name"`
	RulesPackageArns []string `json:"rulesPackageArns" tf:"rules_package_arns"`
	TargetArn        string   `json:"targetArn" tf:"target_arn"`
}

type InspectorAssessmentTemplateStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *InspectorAssessmentTemplateSpec `json:"output,omitempty"`
	// +optional
	State *base.State `json:"state,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// InspectorAssessmentTemplateList is a list of InspectorAssessmentTemplates
type InspectorAssessmentTemplateList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of InspectorAssessmentTemplate CRD objects
	Items []InspectorAssessmentTemplate `json:"items,omitempty"`
}
