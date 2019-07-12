package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsCloudformationStackSet struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsCloudformationStackSetSpec   `json:"spec,omitempty"`
	Status            AwsCloudformationStackSetStatus `json:"status,omitempty"`
}

type AwsCloudformationStackSetSpec struct {
	Capabilities          []string          `json:"capabilities"`
	Parameters            map[string]string `json:"parameters"`
	Tags                  map[string]string `json:"tags"`
	TemplateUrl           string            `json:"template_url"`
	AdministrationRoleArn string            `json:"administration_role_arn"`
	Arn                   string            `json:"arn"`
	Description           string            `json:"description"`
	ExecutionRoleName     string            `json:"execution_role_name"`
	Name                  string            `json:"name"`
	StackSetId            string            `json:"stack_set_id"`
	TemplateBody          string            `json:"template_body"`
}



type AwsCloudformationStackSetStatus struct {
    Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsCloudformationStackSetList is a list of AwsCloudformationStackSets
type AwsCloudformationStackSetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsCloudformationStackSet CRD objects
	Items []AwsCloudformationStackSet `json:"items,omitempty"`
}