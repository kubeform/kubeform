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

type AwsSesIdentityPolicy struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsSesIdentityPolicySpec   `json:"spec,omitempty"`
	Status            AwsSesIdentityPolicyStatus `json:"status,omitempty"`
}

type AwsSesIdentityPolicySpec struct {
	Identity string `json:"identity"`
	Name     string `json:"name"`
	Policy   string `json:"policy"`
}

type AwsSesIdentityPolicyStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AwsSesIdentityPolicyList is a list of AwsSesIdentityPolicys
type AwsSesIdentityPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsSesIdentityPolicy CRD objects
	Items []AwsSesIdentityPolicy `json:"items,omitempty"`
}
