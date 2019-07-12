package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type GoogleServiceAccountIamPolicy struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              GoogleServiceAccountIamPolicySpec   `json:"spec,omitempty"`
	Status            GoogleServiceAccountIamPolicyStatus `json:"status,omitempty"`
}

type GoogleServiceAccountIamPolicySpec struct {
	PolicyData       string `json:"policy_data"`
	Etag             string `json:"etag"`
	ServiceAccountId string `json:"service_account_id"`
}



type GoogleServiceAccountIamPolicyStatus struct {
    Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// GoogleServiceAccountIamPolicyList is a list of GoogleServiceAccountIamPolicys
type GoogleServiceAccountIamPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of GoogleServiceAccountIamPolicy CRD objects
	Items []GoogleServiceAccountIamPolicy `json:"items,omitempty"`
}