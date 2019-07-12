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

type AwsEcrRepositoryPolicy struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsEcrRepositoryPolicySpec   `json:"spec,omitempty"`
	Status            AwsEcrRepositoryPolicyStatus `json:"status,omitempty"`
}

type AwsEcrRepositoryPolicySpec struct {
	Repository string `json:"repository"`
	Policy     string `json:"policy"`
	RegistryId string `json:"registry_id"`
}

type AwsEcrRepositoryPolicyStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AwsEcrRepositoryPolicyList is a list of AwsEcrRepositoryPolicys
type AwsEcrRepositoryPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsEcrRepositoryPolicy CRD objects
	Items []AwsEcrRepositoryPolicy `json:"items,omitempty"`
}
