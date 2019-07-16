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

type OrganizationPolicy struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              OrganizationPolicySpec   `json:"spec,omitempty"`
	Status            OrganizationPolicyStatus `json:"status,omitempty"`
}

type OrganizationPolicySpecBooleanPolicy struct {
	Enforced bool `json:"enforced"`
}

type OrganizationPolicySpecListPolicyAllow struct {
	// +optional
	All bool `json:"all,omitempty"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	Values []string `json:"values,omitempty"`
}

type OrganizationPolicySpecListPolicyDeny struct {
	// +optional
	All bool `json:"all,omitempty"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	Values []string `json:"values,omitempty"`
}

type OrganizationPolicySpecListPolicy struct {
	// +optional
	// +kubebuilder:validation:MaxItems=1
	Allow *[]OrganizationPolicySpecListPolicy `json:"allow,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	Deny *[]OrganizationPolicySpecListPolicy `json:"deny,omitempty"`
}

type OrganizationPolicySpecRestorePolicy struct {
	Default bool `json:"default"`
}

type OrganizationPolicySpec struct {
	// +optional
	// +kubebuilder:validation:MaxItems=1
	BooleanPolicy *[]OrganizationPolicySpec `json:"boolean_policy,omitempty"`
	Constraint    string                    `json:"constraint"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	ListPolicy *[]OrganizationPolicySpec `json:"list_policy,omitempty"`
	OrgId      string                    `json:"org_id"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	RestorePolicy *[]OrganizationPolicySpec `json:"restore_policy,omitempty"`
}

type OrganizationPolicyStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// OrganizationPolicyList is a list of OrganizationPolicys
type OrganizationPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of OrganizationPolicy CRD objects
	Items []OrganizationPolicy `json:"items,omitempty"`
}
