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

type ProjectOrganizationPolicy struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ProjectOrganizationPolicySpec   `json:"spec,omitempty"`
	Status            ProjectOrganizationPolicyStatus `json:"status,omitempty"`
}

type ProjectOrganizationPolicySpecBooleanPolicy struct {
	Enforced bool `json:"enforced"`
}

type ProjectOrganizationPolicySpecListPolicyAllow struct {
	// +optional
	All bool `json:"all,omitempty"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	Values []string `json:"values,omitempty"`
}

type ProjectOrganizationPolicySpecListPolicyDeny struct {
	// +optional
	All bool `json:"all,omitempty"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	Values []string `json:"values,omitempty"`
}

type ProjectOrganizationPolicySpecListPolicy struct {
	// +optional
	// +kubebuilder:validation:MaxItems=1
	Allow *[]ProjectOrganizationPolicySpecListPolicy `json:"allow,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	Deny *[]ProjectOrganizationPolicySpecListPolicy `json:"deny,omitempty"`
}

type ProjectOrganizationPolicySpecRestorePolicy struct {
	Default bool `json:"default"`
}

type ProjectOrganizationPolicySpec struct {
	// +optional
	// +kubebuilder:validation:MaxItems=1
	BooleanPolicy *[]ProjectOrganizationPolicySpec `json:"boolean_policy,omitempty"`
	Constraint    string                           `json:"constraint"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	ListPolicy *[]ProjectOrganizationPolicySpec `json:"list_policy,omitempty"`
	Project    string                           `json:"project"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	RestorePolicy *[]ProjectOrganizationPolicySpec `json:"restore_policy,omitempty"`
}

type ProjectOrganizationPolicyStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
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
