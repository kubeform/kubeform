package v1alpha1

import (
	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

type CloudwatchLogResourcePolicy struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              CloudwatchLogResourcePolicySpec   `json:"spec,omitempty"`
	Status            CloudwatchLogResourcePolicyStatus `json:"status,omitempty"`
}

type CloudwatchLogResourcePolicySpec struct {
	PolicyDocument string                    `json:"policyDocument" tf:"policy_document"`
	PolicyName     string                    `json:"policyName" tf:"policy_name"`
	ProviderRef    core.LocalObjectReference `json:"providerRef" tf:"-"`
}

type CloudwatchLogResourcePolicyStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	TFState     *runtime.RawExtension `json:"tfState,omitempty"`
	TFStateHash string                `json:"tfStateHash,omitempty"`
	Output      *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// CloudwatchLogResourcePolicyList is a list of CloudwatchLogResourcePolicys
type CloudwatchLogResourcePolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of CloudwatchLogResourcePolicy CRD objects
	Items []CloudwatchLogResourcePolicy `json:"items,omitempty"`
}
