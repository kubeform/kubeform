package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsConfigAggregateAuthorization struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsConfigAggregateAuthorizationSpec   `json:"spec,omitempty"`
	Status            AwsConfigAggregateAuthorizationStatus `json:"status,omitempty"`
}

type AwsConfigAggregateAuthorizationSpec struct {
	Arn       string `json:"arn"`
	AccountId string `json:"account_id"`
	Region    string `json:"region"`
}

type AwsConfigAggregateAuthorizationStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsConfigAggregateAuthorizationList is a list of AwsConfigAggregateAuthorizations
type AwsConfigAggregateAuthorizationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsConfigAggregateAuthorization CRD objects
	Items []AwsConfigAggregateAuthorization `json:"items,omitempty"`
}
