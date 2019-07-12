package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type GoogleKmsKeyRingIamPolicy struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              GoogleKmsKeyRingIamPolicySpec   `json:"spec,omitempty"`
	Status            GoogleKmsKeyRingIamPolicyStatus `json:"status,omitempty"`
}

type GoogleKmsKeyRingIamPolicySpec struct {
	Etag       string `json:"etag"`
	KeyRingId  string `json:"key_ring_id"`
	PolicyData string `json:"policy_data"`
}

type GoogleKmsKeyRingIamPolicyStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// GoogleKmsKeyRingIamPolicyList is a list of GoogleKmsKeyRingIamPolicys
type GoogleKmsKeyRingIamPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of GoogleKmsKeyRingIamPolicy CRD objects
	Items []GoogleKmsKeyRingIamPolicy `json:"items,omitempty"`
}
