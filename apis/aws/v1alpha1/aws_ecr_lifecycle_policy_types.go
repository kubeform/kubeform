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

type AwsEcrLifecyclePolicy struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsEcrLifecyclePolicySpec   `json:"spec,omitempty"`
	Status            AwsEcrLifecyclePolicyStatus `json:"status,omitempty"`
}

type AwsEcrLifecyclePolicySpec struct {
	Repository string `json:"repository"`
	Policy     string `json:"policy"`
	RegistryId string `json:"registry_id"`
}



type AwsEcrLifecyclePolicyStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AwsEcrLifecyclePolicyList is a list of AwsEcrLifecyclePolicys
type AwsEcrLifecyclePolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsEcrLifecyclePolicy CRD objects
	Items []AwsEcrLifecyclePolicy `json:"items,omitempty"`
}