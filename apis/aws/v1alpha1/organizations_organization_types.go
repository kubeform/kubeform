package v1alpha1

import (
    "encoding/json"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	core "k8s.io/api/core/v1"
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
	// +optional
	Arn string `json:"arn,omitempty" tf:"arn,omitempty"`
	// +optional
	Email string `json:"email,omitempty" tf:"email,omitempty"`
	// +optional
	ID string `json:"ID,omitempty" tf:"id,omitempty"`
	// +optional
	Name string `json:"name,omitempty" tf:"name,omitempty"`
}

type OrganizationsOrganizationSpecRootsPolicyTypes struct {
	// +optional
	Status string `json:"status,omitempty" tf:"status,omitempty"`
	// +optional
	Type string `json:"type,omitempty" tf:"type,omitempty"`
}

type OrganizationsOrganizationSpecRoots struct {
	// +optional
	Arn string `json:"arn,omitempty" tf:"arn,omitempty"`
	// +optional
	ID string `json:"ID,omitempty" tf:"id,omitempty"`
	// +optional
	Name string `json:"name,omitempty" tf:"name,omitempty"`
	// +optional
	PolicyTypes []OrganizationsOrganizationSpecRootsPolicyTypes `json:"policyTypes,omitempty" tf:"policy_types,omitempty"`
}

type OrganizationsOrganizationSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	Accounts []OrganizationsOrganizationSpecAccounts `json:"accounts,omitempty" tf:"accounts,omitempty"`
	// +optional
	Arn string `json:"arn,omitempty" tf:"arn,omitempty"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	AwsServiceAccessPrincipals []string `json:"awsServiceAccessPrincipals,omitempty" tf:"aws_service_access_principals,omitempty"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	EnabledPolicyTypes []string `json:"enabledPolicyTypes,omitempty" tf:"enabled_policy_types,omitempty"`
	// +optional
	FeatureSet string `json:"featureSet,omitempty" tf:"feature_set,omitempty"`
	// +optional
	MasterAccountArn string `json:"masterAccountArn,omitempty" tf:"master_account_arn,omitempty"`
	// +optional
	MasterAccountEmail string `json:"masterAccountEmail,omitempty" tf:"master_account_email,omitempty"`
	// +optional
	MasterAccountID string `json:"masterAccountID,omitempty" tf:"master_account_id,omitempty"`
	// +optional
	Roots []OrganizationsOrganizationSpecRoots `json:"roots,omitempty" tf:"roots,omitempty"`
}



type OrganizationsOrganizationStatus struct {
    // Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64  `json:"observedGeneration,omitempty"`
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