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

type AwsIamRolePolicy struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsIamRolePolicySpec   `json:"spec,omitempty"`
	Status            AwsIamRolePolicyStatus `json:"status,omitempty"`
}

type AwsIamRolePolicySpec struct {
	Policy     string `json:"policy"`
	Name       string `json:"name"`
	NamePrefix string `json:"name_prefix"`
	Role       string `json:"role"`
}



type AwsIamRolePolicyStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AwsIamRolePolicyList is a list of AwsIamRolePolicys
type AwsIamRolePolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsIamRolePolicy CRD objects
	Items []AwsIamRolePolicy `json:"items,omitempty"`
}