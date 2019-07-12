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

type GoogleProjectOrganizationPolicy struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              GoogleProjectOrganizationPolicySpec   `json:"spec,omitempty"`
	Status            GoogleProjectOrganizationPolicyStatus `json:"status,omitempty"`
}

type GoogleProjectOrganizationPolicySpecRestorePolicy struct {
	Default bool `json:"default"`
}

type GoogleProjectOrganizationPolicySpecBooleanPolicy struct {
	Enforced bool `json:"enforced"`
}

type GoogleProjectOrganizationPolicySpecListPolicyAllow struct {
	All    bool     `json:"all"`
	Values []string `json:"values"`
}

type GoogleProjectOrganizationPolicySpecListPolicyDeny struct {
	All    bool     `json:"all"`
	Values []string `json:"values"`
}

type GoogleProjectOrganizationPolicySpecListPolicy struct {
	SuggestedValue string                                          `json:"suggested_value"`
	Allow          []GoogleProjectOrganizationPolicySpecListPolicy `json:"allow"`
	Deny           []GoogleProjectOrganizationPolicySpecListPolicy `json:"deny"`
}

type GoogleProjectOrganizationPolicySpec struct {
	UpdateTime    string                                `json:"update_time"`
	RestorePolicy []GoogleProjectOrganizationPolicySpec `json:"restore_policy"`
	Constraint    string                                `json:"constraint"`
	BooleanPolicy []GoogleProjectOrganizationPolicySpec `json:"boolean_policy"`
	ListPolicy    []GoogleProjectOrganizationPolicySpec `json:"list_policy"`
	Version       int                                   `json:"version"`
	Etag          string                                `json:"etag"`
	Project       string                                `json:"project"`
}

type GoogleProjectOrganizationPolicyStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// GoogleProjectOrganizationPolicyList is a list of GoogleProjectOrganizationPolicys
type GoogleProjectOrganizationPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of GoogleProjectOrganizationPolicy CRD objects
	Items []GoogleProjectOrganizationPolicy `json:"items,omitempty"`
}
