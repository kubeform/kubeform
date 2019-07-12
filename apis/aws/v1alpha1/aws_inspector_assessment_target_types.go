package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsInspectorAssessmentTarget struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsInspectorAssessmentTargetSpec   `json:"spec,omitempty"`
	Status            AwsInspectorAssessmentTargetStatus `json:"status,omitempty"`
}

type AwsInspectorAssessmentTargetSpec struct {
	Name             string `json:"name"`
	Arn              string `json:"arn"`
	ResourceGroupArn string `json:"resource_group_arn"`
}



type AwsInspectorAssessmentTargetStatus struct {
    Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsInspectorAssessmentTargetList is a list of AwsInspectorAssessmentTargets
type AwsInspectorAssessmentTargetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsInspectorAssessmentTarget CRD objects
	Items []AwsInspectorAssessmentTarget `json:"items,omitempty"`
}