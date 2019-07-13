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

type AwsCloudformationStackSet struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsCloudformationStackSetSpec   `json:"spec,omitempty"`
	Status            AwsCloudformationStackSetStatus `json:"status,omitempty"`
}

type AwsCloudformationStackSetSpec struct {
	AdministrationRoleArn string            `json:"administration_role_arn"`
	Name                  string            `json:"name"`
	TemplateBody          string            `json:"template_body"`
	TemplateUrl           string            `json:"template_url"`
	StackSetId            string            `json:"stack_set_id"`
	Tags                  map[string]string `json:"tags"`
	Arn                   string            `json:"arn"`
	Capabilities          []string          `json:"capabilities"`
	Description           string            `json:"description"`
	ExecutionRoleName     string            `json:"execution_role_name"`
	Parameters            map[string]string `json:"parameters"`
}



type AwsCloudformationStackSetStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AwsCloudformationStackSetList is a list of AwsCloudformationStackSets
type AwsCloudformationStackSetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsCloudformationStackSet CRD objects
	Items []AwsCloudformationStackSet `json:"items,omitempty"`
}