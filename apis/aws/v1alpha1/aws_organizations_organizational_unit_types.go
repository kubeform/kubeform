package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsOrganizationsOrganizationalUnit struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsOrganizationsOrganizationalUnitSpec   `json:"spec,omitempty"`
	Status            AwsOrganizationsOrganizationalUnitStatus `json:"status,omitempty"`
}

type AwsOrganizationsOrganizationalUnitSpecAccounts struct {
	Arn   string `json:"arn"`
	Email string `json:"email"`
	Id    string `json:"id"`
	Name  string `json:"name"`
}

type AwsOrganizationsOrganizationalUnitSpec struct {
	Name     string                                   `json:"name"`
	ParentId string                                   `json:"parent_id"`
	Accounts []AwsOrganizationsOrganizationalUnitSpec `json:"accounts"`
	Arn      string                                   `json:"arn"`
}

type AwsOrganizationsOrganizationalUnitStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsOrganizationsOrganizationalUnitList is a list of AwsOrganizationsOrganizationalUnits
type AwsOrganizationsOrganizationalUnitList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsOrganizationsOrganizationalUnit CRD objects
	Items []AwsOrganizationsOrganizationalUnit `json:"items,omitempty"`
}
