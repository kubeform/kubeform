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

type AwsOrganizationsOrganizationalUnit struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsOrganizationsOrganizationalUnitSpec   `json:"spec,omitempty"`
	Status            AwsOrganizationsOrganizationalUnitStatus `json:"status,omitempty"`
}

type AwsOrganizationsOrganizationalUnitSpecAccounts struct {
	Id    string `json:"id"`
	Name  string `json:"name"`
	Arn   string `json:"arn"`
	Email string `json:"email"`
}

type AwsOrganizationsOrganizationalUnitSpec struct {
	Name     string                                   `json:"name"`
	ParentId string                                   `json:"parent_id"`
	Accounts []AwsOrganizationsOrganizationalUnitSpec `json:"accounts"`
	Arn      string                                   `json:"arn"`
}



type AwsOrganizationsOrganizationalUnitStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AwsOrganizationsOrganizationalUnitList is a list of AwsOrganizationsOrganizationalUnits
type AwsOrganizationsOrganizationalUnitList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsOrganizationsOrganizationalUnit CRD objects
	Items []AwsOrganizationsOrganizationalUnit `json:"items,omitempty"`
}