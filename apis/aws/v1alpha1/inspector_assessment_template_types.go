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

type InspectorAssessmentTemplate struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              InspectorAssessmentTemplateSpec   `json:"spec,omitempty"`
	Status            InspectorAssessmentTemplateStatus `json:"status,omitempty"`
}

type InspectorAssessmentTemplateSpec struct {
	Duration int    `json:"duration"`
	Name     string `json:"name"`
	// +kubebuilder:validation:UniqueItems=true
	RulesPackageArns []string `json:"rules_package_arns"`
	TargetArn        string   `json:"target_arn"`
}

type InspectorAssessmentTemplateStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
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
