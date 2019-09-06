package v1alpha1

import (
	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"kubeform.dev/kubeform/apis"
)

// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

type CloudwatchLogDestinationPolicy struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              CloudwatchLogDestinationPolicySpec   `json:"spec,omitempty"`
	Status            CloudwatchLogDestinationPolicyStatus `json:"status,omitempty"`
}

type CloudwatchLogDestinationPolicySpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	AccessPolicy    string `json:"accessPolicy" tf:"access_policy"`
	DestinationName string `json:"destinationName" tf:"destination_name"`
}

type CloudwatchLogDestinationPolicyStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *CloudwatchLogDestinationPolicySpec `json:"output,omitempty"`
	// +optional
	State *apis.State `json:"state,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// CloudwatchLogDestinationPolicyList is a list of CloudwatchLogDestinationPolicys
type CloudwatchLogDestinationPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of CloudwatchLogDestinationPolicy CRD objects
	Items []CloudwatchLogDestinationPolicy `json:"items,omitempty"`
}
