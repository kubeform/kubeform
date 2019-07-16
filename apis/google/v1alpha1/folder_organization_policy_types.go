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

type FolderOrganizationPolicy struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              FolderOrganizationPolicySpec   `json:"spec,omitempty"`
	Status            FolderOrganizationPolicyStatus `json:"status,omitempty"`
}

type FolderOrganizationPolicySpecBooleanPolicy struct {
	Enforced bool `json:"enforced"`
}

type FolderOrganizationPolicySpecListPolicyAllow struct {
	// +optional
	All bool `json:"all,omitempty"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	Values []string `json:"values,omitempty"`
}

type FolderOrganizationPolicySpecListPolicyDeny struct {
	// +optional
	All bool `json:"all,omitempty"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	Values []string `json:"values,omitempty"`
}

type FolderOrganizationPolicySpecListPolicy struct {
	// +optional
	// +kubebuilder:validation:MaxItems=1
	Allow *[]FolderOrganizationPolicySpecListPolicy `json:"allow,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	Deny *[]FolderOrganizationPolicySpecListPolicy `json:"deny,omitempty"`
}

type FolderOrganizationPolicySpecRestorePolicy struct {
	Default bool `json:"default"`
}

type FolderOrganizationPolicySpec struct {
	// +optional
	// +kubebuilder:validation:MaxItems=1
	BooleanPolicy *[]FolderOrganizationPolicySpec `json:"boolean_policy,omitempty"`
	Constraint    string                          `json:"constraint"`
	Folder        string                          `json:"folder"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	ListPolicy *[]FolderOrganizationPolicySpec `json:"list_policy,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	RestorePolicy *[]FolderOrganizationPolicySpec `json:"restore_policy,omitempty"`
}

type FolderOrganizationPolicyStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// FolderOrganizationPolicyList is a list of FolderOrganizationPolicys
type FolderOrganizationPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of FolderOrganizationPolicy CRD objects
	Items []FolderOrganizationPolicy `json:"items,omitempty"`
}
