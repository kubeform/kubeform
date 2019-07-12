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

type AwsSecurityhubProductSubscription struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsSecurityhubProductSubscriptionSpec   `json:"spec,omitempty"`
	Status            AwsSecurityhubProductSubscriptionStatus `json:"status,omitempty"`
}

type AwsSecurityhubProductSubscriptionSpec struct {
	ProductArn string `json:"product_arn"`
	Arn        string `json:"arn"`
}

type AwsSecurityhubProductSubscriptionStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AwsSecurityhubProductSubscriptionList is a list of AwsSecurityhubProductSubscriptions
type AwsSecurityhubProductSubscriptionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsSecurityhubProductSubscription CRD objects
	Items []AwsSecurityhubProductSubscription `json:"items,omitempty"`
}
