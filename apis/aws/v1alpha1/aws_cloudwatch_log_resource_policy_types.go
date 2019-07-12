package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsCloudwatchLogResourcePolicy struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsCloudwatchLogResourcePolicySpec   `json:"spec,omitempty"`
	Status            AwsCloudwatchLogResourcePolicyStatus `json:"status,omitempty"`
}

type AwsCloudwatchLogResourcePolicySpec struct {
	PolicyName     string `json:"policy_name"`
	PolicyDocument string `json:"policy_document"`
}

type AwsCloudwatchLogResourcePolicyStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsCloudwatchLogResourcePolicyList is a list of AwsCloudwatchLogResourcePolicys
type AwsCloudwatchLogResourcePolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsCloudwatchLogResourcePolicy CRD objects
	Items []AwsCloudwatchLogResourcePolicy `json:"items,omitempty"`
}
