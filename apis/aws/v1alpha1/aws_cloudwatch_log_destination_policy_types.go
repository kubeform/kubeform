package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsCloudwatchLogDestinationPolicy struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsCloudwatchLogDestinationPolicySpec   `json:"spec,omitempty"`
	Status            AwsCloudwatchLogDestinationPolicyStatus `json:"status,omitempty"`
}

type AwsCloudwatchLogDestinationPolicySpec struct {
	DestinationName string `json:"destination_name"`
	AccessPolicy    string `json:"access_policy"`
}



type AwsCloudwatchLogDestinationPolicyStatus struct {
    Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsCloudwatchLogDestinationPolicyList is a list of AwsCloudwatchLogDestinationPolicys
type AwsCloudwatchLogDestinationPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsCloudwatchLogDestinationPolicy CRD objects
	Items []AwsCloudwatchLogDestinationPolicy `json:"items,omitempty"`
}