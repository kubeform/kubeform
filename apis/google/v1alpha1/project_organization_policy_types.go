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

type ProjectOrganizationPolicy struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ProjectOrganizationPolicySpec   `json:"spec,omitempty"`
	Status            ProjectOrganizationPolicyStatus `json:"status,omitempty"`
}

type ProjectOrganizationPolicySpecBooleanPolicy struct {
	Enforced bool `json:"enforced" tf:"enforced"`
}

type ProjectOrganizationPolicySpecListPolicyAllow struct {
	// +optional
	All bool `json:"all,omitempty" tf:"all,omitempty"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	Values []string `json:"values,omitempty" tf:"values,omitempty"`
}

type ProjectOrganizationPolicySpecListPolicyDeny struct {
	// +optional
	All bool `json:"all,omitempty" tf:"all,omitempty"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	Values []string `json:"values,omitempty" tf:"values,omitempty"`
}

type ProjectOrganizationPolicySpecListPolicy struct {
	// +optional
	// +kubebuilder:validation:MaxItems=1
	Allow []ProjectOrganizationPolicySpecListPolicyAllow `json:"allow,omitempty" tf:"allow,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	Deny []ProjectOrganizationPolicySpecListPolicyDeny `json:"deny,omitempty" tf:"deny,omitempty"`
	// +optional
	SuggestedValue string `json:"suggestedValue,omitempty" tf:"suggested_value,omitempty"`
}

type ProjectOrganizationPolicySpecRestorePolicy struct {
	Default bool `json:"default" tf:"default"`
}

type ProjectOrganizationPolicySpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	// +optional
	// +kubebuilder:validation:MaxItems=1
	BooleanPolicy []ProjectOrganizationPolicySpecBooleanPolicy `json:"booleanPolicy,omitempty" tf:"boolean_policy,omitempty"`
	Constraint    string                                       `json:"constraint" tf:"constraint"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	ListPolicy []ProjectOrganizationPolicySpecListPolicy `json:"listPolicy,omitempty" tf:"list_policy,omitempty"`
	Project    string                                    `json:"project" tf:"project"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	RestorePolicy []ProjectOrganizationPolicySpecRestorePolicy `json:"restorePolicy,omitempty" tf:"restore_policy,omitempty"`
	// +optional
	Version int `json:"version,omitempty" tf:"version,omitempty"`
}

type ProjectOrganizationPolicyStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	TFState *runtime.RawExtension `json:"tfState,omitempty"`
	Output  *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// ProjectOrganizationPolicyList is a list of ProjectOrganizationPolicys
type ProjectOrganizationPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of ProjectOrganizationPolicy CRD objects
	Items []ProjectOrganizationPolicy `json:"items,omitempty"`
}
