package v1alpha1

import (
	base "kubeform.dev/kubeform/apis/base/v1alpha1"

	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Phase",type=string,JSONPath=`.status.phase`

type ProjectIamPolicy struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ProjectIamPolicySpec   `json:"spec,omitempty"`
	Status            ProjectIamPolicyStatus `json:"status,omitempty"`
}

type ProjectIamPolicySpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	// Deprecated
	Authoritative bool `json:"authoritative,omitempty" tf:"authoritative,omitempty"`
	// +optional
	// Deprecated
	DisableProject bool `json:"disableProject,omitempty" tf:"disable_project,omitempty"`
	// +optional
	Etag       string `json:"etag,omitempty" tf:"etag,omitempty"`
	PolicyData string `json:"policyData" tf:"policy_data"`
	// +optional
	Project string `json:"project,omitempty" tf:"project,omitempty"`
	// +optional
	// Deprecated
	RestorePolicy string `json:"restorePolicy,omitempty" tf:"restore_policy,omitempty"`
}

type ProjectIamPolicyStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *ProjectIamPolicySpec `json:"output,omitempty"`
	// +optional
	State *base.State `json:"state,omitempty"`
	// +optional
	Phase base.Phase `json:"phase,omitempty"`
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
