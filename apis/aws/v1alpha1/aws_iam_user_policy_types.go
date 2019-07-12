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

type AwsIamUserPolicy struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsIamUserPolicySpec   `json:"spec,omitempty"`
	Status            AwsIamUserPolicyStatus `json:"status,omitempty"`
}

type AwsIamUserPolicySpec struct {
	User       string `json:"user"`
	Policy     string `json:"policy"`
	Name       string `json:"name"`
	NamePrefix string `json:"name_prefix"`
}

type AwsIamUserPolicyStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AwsIamUserPolicyList is a list of AwsIamUserPolicys
type AwsIamUserPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsIamUserPolicy CRD objects
	Items []AwsIamUserPolicy `json:"items,omitempty"`
}
