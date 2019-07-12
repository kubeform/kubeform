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

type AwsDirectoryServiceLogSubscription struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsDirectoryServiceLogSubscriptionSpec   `json:"spec,omitempty"`
	Status            AwsDirectoryServiceLogSubscriptionStatus `json:"status,omitempty"`
}

type AwsDirectoryServiceLogSubscriptionSpec struct {
	DirectoryId  string `json:"directory_id"`
	LogGroupName string `json:"log_group_name"`
}

type AwsDirectoryServiceLogSubscriptionStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AwsDirectoryServiceLogSubscriptionList is a list of AwsDirectoryServiceLogSubscriptions
type AwsDirectoryServiceLogSubscriptionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsDirectoryServiceLogSubscription CRD objects
	Items []AwsDirectoryServiceLogSubscription `json:"items,omitempty"`
}
