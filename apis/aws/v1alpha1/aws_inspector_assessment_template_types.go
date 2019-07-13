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

type AwsInspectorAssessmentTemplate struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsInspectorAssessmentTemplateSpec   `json:"spec,omitempty"`
	Status            AwsInspectorAssessmentTemplateStatus `json:"status,omitempty"`
}

type AwsInspectorAssessmentTemplateSpec struct {
	Name             string   `json:"name"`
	TargetArn        string   `json:"target_arn"`
	Arn              string   `json:"arn"`
	Duration         int      `json:"duration"`
	RulesPackageArns []string `json:"rules_package_arns"`
}



type AwsInspectorAssessmentTemplateStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AwsInspectorAssessmentTemplateList is a list of AwsInspectorAssessmentTemplates
type AwsInspectorAssessmentTemplateList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsInspectorAssessmentTemplate CRD objects
	Items []AwsInspectorAssessmentTemplate `json:"items,omitempty"`
}