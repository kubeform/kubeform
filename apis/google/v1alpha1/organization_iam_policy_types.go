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

type OrganizationIamPolicy struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              OrganizationIamPolicySpec   `json:"spec,omitempty"`
	Status            OrganizationIamPolicyStatus `json:"status,omitempty"`
}

type OrganizationIamPolicySpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`
	OrgID       string                    `json:"orgID" tf:"org_id"`
	PolicyData  string                    `json:"policyData" tf:"policy_data"`
}

type OrganizationIamPolicyStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	TFState *runtime.RawExtension `json:"tfState,omitempty"`
	Output  *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// OrganizationIamPolicyList is a list of OrganizationIamPolicys
type OrganizationIamPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of OrganizationIamPolicy CRD objects
	Items []OrganizationIamPolicy `json:"items,omitempty"`
}
