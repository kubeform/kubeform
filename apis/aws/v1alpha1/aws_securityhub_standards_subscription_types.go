package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsSecurityhubStandardsSubscription struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsSecurityhubStandardsSubscriptionSpec   `json:"spec,omitempty"`
	Status            AwsSecurityhubStandardsSubscriptionStatus `json:"status,omitempty"`
}

type AwsSecurityhubStandardsSubscriptionSpec struct {
	StandardsArn string `json:"standards_arn"`
}



type AwsSecurityhubStandardsSubscriptionStatus struct {
    Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsSecurityhubStandardsSubscriptionList is a list of AwsSecurityhubStandardsSubscriptions
type AwsSecurityhubStandardsSubscriptionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsSecurityhubStandardsSubscription CRD objects
	Items []AwsSecurityhubStandardsSubscription `json:"items,omitempty"`
}