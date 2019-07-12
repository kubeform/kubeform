package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsKmsAlias struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsKmsAliasSpec   `json:"spec,omitempty"`
	Status            AwsKmsAliasStatus `json:"status,omitempty"`
}

type AwsKmsAliasSpec struct {
	Name         string `json:"name"`
	NamePrefix   string `json:"name_prefix"`
	TargetKeyId  string `json:"target_key_id"`
	TargetKeyArn string `json:"target_key_arn"`
	Arn          string `json:"arn"`
}



type AwsKmsAliasStatus struct {
    Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsKmsAliasList is a list of AwsKmsAliass
type AwsKmsAliasList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsKmsAlias CRD objects
	Items []AwsKmsAlias `json:"items,omitempty"`
}