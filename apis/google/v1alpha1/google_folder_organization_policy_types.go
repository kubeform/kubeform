package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type GoogleFolderOrganizationPolicy struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              GoogleFolderOrganizationPolicySpec   `json:"spec,omitempty"`
	Status            GoogleFolderOrganizationPolicyStatus `json:"status,omitempty"`
}

type GoogleFolderOrganizationPolicySpecRestorePolicy struct {
	Default bool `json:"default"`
}

type GoogleFolderOrganizationPolicySpecBooleanPolicy struct {
	Enforced bool `json:"enforced"`
}

type GoogleFolderOrganizationPolicySpecListPolicyAllow struct {
	All    bool     `json:"all"`
	Values []string `json:"values"`
}

type GoogleFolderOrganizationPolicySpecListPolicyDeny struct {
	All    bool     `json:"all"`
	Values []string `json:"values"`
}

type GoogleFolderOrganizationPolicySpecListPolicy struct {
	Allow          []GoogleFolderOrganizationPolicySpecListPolicy `json:"allow"`
	Deny           []GoogleFolderOrganizationPolicySpecListPolicy `json:"deny"`
	SuggestedValue string                                         `json:"suggested_value"`
}

type GoogleFolderOrganizationPolicySpec struct {
	Etag          string                               `json:"etag"`
	UpdateTime    string                               `json:"update_time"`
	RestorePolicy []GoogleFolderOrganizationPolicySpec `json:"restore_policy"`
	Constraint    string                               `json:"constraint"`
	BooleanPolicy []GoogleFolderOrganizationPolicySpec `json:"boolean_policy"`
	ListPolicy    []GoogleFolderOrganizationPolicySpec `json:"list_policy"`
	Folder        string                               `json:"folder"`
	Version       int                                  `json:"version"`
}

type GoogleFolderOrganizationPolicyStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// GoogleFolderOrganizationPolicyList is a list of GoogleFolderOrganizationPolicys
type GoogleFolderOrganizationPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of GoogleFolderOrganizationPolicy CRD objects
	Items []GoogleFolderOrganizationPolicy `json:"items,omitempty"`
}
