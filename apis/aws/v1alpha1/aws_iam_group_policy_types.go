package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsIamGroupPolicy struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsIamGroupPolicySpec   `json:"spec,omitempty"`
	Status            AwsIamGroupPolicyStatus `json:"status,omitempty"`
}

type AwsIamGroupPolicySpec struct {
	Policy     string `json:"policy"`
	Name       string `json:"name"`
	NamePrefix string `json:"name_prefix"`
	Group      string `json:"group"`
}



type AwsIamGroupPolicyStatus struct {
    Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsIamGroupPolicyList is a list of AwsIamGroupPolicys
type AwsIamGroupPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsIamGroupPolicy CRD objects
	Items []AwsIamGroupPolicy `json:"items,omitempty"`
}