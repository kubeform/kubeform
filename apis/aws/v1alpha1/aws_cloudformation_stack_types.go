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

type AwsCloudformationStack struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsCloudformationStackSpec   `json:"spec,omitempty"`
	Status            AwsCloudformationStackStatus `json:"status,omitempty"`
}

type AwsCloudformationStackSpec struct {
	Capabilities     []string          `json:"capabilities"`
	Outputs          map[string]string `json:"outputs"`
	PolicyUrl        string            `json:"policy_url"`
	Tags             map[string]string `json:"tags"`
	TemplateBody     string            `json:"template_body"`
	TemplateUrl      string            `json:"template_url"`
	DisableRollback  bool              `json:"disable_rollback"`
	PolicyBody       string            `json:"policy_body"`
	TimeoutInMinutes int               `json:"timeout_in_minutes"`
	Name             string            `json:"name"`
	OnFailure        string            `json:"on_failure"`
	IamRoleArn       string            `json:"iam_role_arn"`
	NotificationArns []string          `json:"notification_arns"`
	Parameters       map[string]string `json:"parameters"`
}



type AwsCloudformationStackStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AwsCloudformationStackList is a list of AwsCloudformationStacks
type AwsCloudformationStackList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsCloudformationStack CRD objects
	Items []AwsCloudformationStack `json:"items,omitempty"`
}