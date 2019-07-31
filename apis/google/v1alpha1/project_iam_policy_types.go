package v1alpha1

import (
	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

type ProjectIamPolicy struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ProjectIamPolicySpec   `json:"spec,omitempty"`
	Status            ProjectIamPolicyStatus `json:"status,omitempty"`
}

type ProjectIamPolicySpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	// +optional
	// Deprecated
	Authoritative bool `json:"authoritative,omitempty" tf:"authoritative,omitempty"`
	// +optional
	// Deprecated
	DisableProject bool   `json:"disableProject,omitempty" tf:"disable_project,omitempty"`
	PolicyData     string `json:"policyData" tf:"policy_data"`
	// +optional
	Project string `json:"project,omitempty" tf:"project,omitempty"`
}

type ProjectIamPolicyStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64  `json:"observedGeneration,omitempty"`
	TFState            string `json:"tfState,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// ProjectIamPolicyList is a list of ProjectIamPolicys
type ProjectIamPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of ProjectIamPolicy CRD objects
	Items []ProjectIamPolicy `json:"items,omitempty"`
}
