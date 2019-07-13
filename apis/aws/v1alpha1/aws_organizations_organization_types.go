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

type AwsOrganizationsOrganization struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsOrganizationsOrganizationSpec   `json:"spec,omitempty"`
	Status            AwsOrganizationsOrganizationStatus `json:"status,omitempty"`
}

type AwsOrganizationsOrganizationSpecNonMasterAccounts struct {
	Arn   string `json:"arn"`
	Email string `json:"email"`
	Id    string `json:"id"`
	Name  string `json:"name"`
}

type AwsOrganizationsOrganizationSpecAccounts struct {
	Arn   string `json:"arn"`
	Email string `json:"email"`
	Id    string `json:"id"`
	Name  string `json:"name"`
}

type AwsOrganizationsOrganizationSpecRootsPolicyTypes struct {
	Status string `json:"status"`
	Type   string `json:"type"`
}

type AwsOrganizationsOrganizationSpecRoots struct {
	Name        string                                  `json:"name"`
	Arn         string                                  `json:"arn"`
	PolicyTypes []AwsOrganizationsOrganizationSpecRoots `json:"policy_types"`
	Id          string                                  `json:"id"`
}

type AwsOrganizationsOrganizationSpec struct {
	Arn                        string                             `json:"arn"`
	MasterAccountEmail         string                             `json:"master_account_email"`
	AwsServiceAccessPrincipals []string                           `json:"aws_service_access_principals"`
	NonMasterAccounts          []AwsOrganizationsOrganizationSpec `json:"non_master_accounts"`
	MasterAccountArn           string                             `json:"master_account_arn"`
	MasterAccountId            string                             `json:"master_account_id"`
	Accounts                   []AwsOrganizationsOrganizationSpec `json:"accounts"`
	Roots                      []AwsOrganizationsOrganizationSpec `json:"roots"`
	EnabledPolicyTypes         []string                           `json:"enabled_policy_types"`
	FeatureSet                 string                             `json:"feature_set"`
}



type AwsOrganizationsOrganizationStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AwsOrganizationsOrganizationList is a list of AwsOrganizationsOrganizations
type AwsOrganizationsOrganizationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsOrganizationsOrganization CRD objects
	Items []AwsOrganizationsOrganization `json:"items,omitempty"`
}