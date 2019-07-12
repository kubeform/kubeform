package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsOrganizationsAccount struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsOrganizationsAccountSpec   `json:"spec,omitempty"`
	Status            AwsOrganizationsAccountStatus `json:"status,omitempty"`
}

type AwsOrganizationsAccountSpec struct {
	JoinedMethod           string `json:"joined_method"`
	JoinedTimestamp        string `json:"joined_timestamp"`
	Status                 string `json:"status"`
	Name                   string `json:"name"`
	Email                  string `json:"email"`
	IamUserAccessToBilling string `json:"iam_user_access_to_billing"`
	RoleName               string `json:"role_name"`
	Arn                    string `json:"arn"`
}

type AwsOrganizationsAccountStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsOrganizationsAccountList is a list of AwsOrganizationsAccounts
type AwsOrganizationsAccountList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsOrganizationsAccount CRD objects
	Items []AwsOrganizationsAccount `json:"items,omitempty"`
}
