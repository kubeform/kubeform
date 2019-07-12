package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsCloudformationStack struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsCloudformationStackSpec   `json:"spec,omitempty"`
	Status            AwsCloudformationStackStatus `json:"status,omitempty"`
}

type AwsCloudformationStackSpec struct {
	Parameters       map[string]string `json:"parameters"`
	Tags             map[string]string `json:"tags"`
	IamRoleArn       string            `json:"iam_role_arn"`
	Name             string            `json:"name"`
	DisableRollback  bool              `json:"disable_rollback"`
	NotificationArns []string          `json:"notification_arns"`
	PolicyBody       string            `json:"policy_body"`
	PolicyUrl        string            `json:"policy_url"`
	TemplateUrl      string            `json:"template_url"`
	Capabilities     []string          `json:"capabilities"`
	Outputs          map[string]string `json:"outputs"`
	TimeoutInMinutes int               `json:"timeout_in_minutes"`
	TemplateBody     string            `json:"template_body"`
	OnFailure        string            `json:"on_failure"`
}



type AwsCloudformationStackStatus struct {
    Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsCloudformationStackList is a list of AwsCloudformationStacks
type AwsCloudformationStackList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsCloudformationStack CRD objects
	Items []AwsCloudformationStack `json:"items,omitempty"`
}