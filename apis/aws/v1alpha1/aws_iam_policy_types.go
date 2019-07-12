package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsIamPolicy struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsIamPolicySpec   `json:"spec,omitempty"`
	Status            AwsIamPolicyStatus `json:"status,omitempty"`
}

type AwsIamPolicySpec struct {
	Name        string `json:"name"`
	NamePrefix  string `json:"name_prefix"`
	Arn         string `json:"arn"`
	Description string `json:"description"`
	Path        string `json:"path"`
	Policy      string `json:"policy"`
}

type AwsIamPolicyStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsIamPolicyList is a list of AwsIamPolicys
type AwsIamPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsIamPolicy CRD objects
	Items []AwsIamPolicy `json:"items,omitempty"`
}
