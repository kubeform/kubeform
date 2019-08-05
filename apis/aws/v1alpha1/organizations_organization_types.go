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

type OrganizationsOrganization struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              OrganizationsOrganizationSpec   `json:"spec,omitempty"`
	Status            OrganizationsOrganizationStatus `json:"status,omitempty"`
}

type OrganizationsOrganizationSpecAccounts struct {
	Arn   string `json:"arn" tf:"arn"`
	Email string `json:"email" tf:"email"`
	ID    string `json:"ID" tf:"id"`
	Name  string `json:"name" tf:"name"`
}

type OrganizationsOrganizationSpecNonMasterAccounts struct {
	Arn   string `json:"arn" tf:"arn"`
	Email string `json:"email" tf:"email"`
	ID    string `json:"ID" tf:"id"`
	Name  string `json:"name" tf:"name"`
}

type OrganizationsOrganizationSpecRootsPolicyTypes struct {
	Status string `json:"status" tf:"status"`
	Type   string `json:"type" tf:"type"`
}

type OrganizationsOrganizationSpecRoots struct {
	Arn         string                                          `json:"arn" tf:"arn"`
	ID          string                                          `json:"ID" tf:"id"`
	Name        string                                          `json:"name" tf:"name"`
	PolicyTypes []OrganizationsOrganizationSpecRootsPolicyTypes `json:"policyTypes" tf:"policy_types"`
}

type OrganizationsOrganizationSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	Accounts []OrganizationsOrganizationSpecAccounts `json:"accounts" tf:"accounts"`
	Arn      string                                  `json:"arn" tf:"arn"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	AwsServiceAccessPrincipals []string `json:"awsServiceAccessPrincipals,omitempty" tf:"aws_service_access_principals,omitempty"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	EnabledPolicyTypes []string `json:"enabledPolicyTypes,omitempty" tf:"enabled_policy_types,omitempty"`
	// +optional
	FeatureSet         string                                           `json:"featureSet,omitempty" tf:"feature_set,omitempty"`
	MasterAccountArn   string                                           `json:"masterAccountArn" tf:"master_account_arn"`
	MasterAccountEmail string                                           `json:"masterAccountEmail" tf:"master_account_email"`
	MasterAccountID    string                                           `json:"masterAccountID" tf:"master_account_id"`
	NonMasterAccounts  []OrganizationsOrganizationSpecNonMasterAccounts `json:"nonMasterAccounts" tf:"non_master_accounts"`
	Roots              []OrganizationsOrganizationSpecRoots             `json:"roots" tf:"roots"`
}

type OrganizationsOrganizationStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *OrganizationsOrganizationSpec `json:"output,omitempty"`
	// +optional
	State *apis.State `json:"state,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// OrganizationsOrganizationList is a list of OrganizationsOrganizations
type OrganizationsOrganizationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of OrganizationsOrganization CRD objects
	Items []OrganizationsOrganization `json:"items,omitempty"`
}
