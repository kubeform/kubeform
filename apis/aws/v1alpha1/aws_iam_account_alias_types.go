package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsIamAccountAlias struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsIamAccountAliasSpec   `json:"spec,omitempty"`
	Status            AwsIamAccountAliasStatus `json:"status,omitempty"`
}

type AwsIamAccountAliasSpec struct {
	AccountAlias string `json:"account_alias"`
}



type AwsIamAccountAliasStatus struct {
    Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsIamAccountAliasList is a list of AwsIamAccountAliass
type AwsIamAccountAliasList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsIamAccountAlias CRD objects
	Items []AwsIamAccountAlias `json:"items,omitempty"`
}