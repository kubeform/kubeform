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

type AwsAutoscalingNotification struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsAutoscalingNotificationSpec   `json:"spec,omitempty"`
	Status            AwsAutoscalingNotificationStatus `json:"status,omitempty"`
}

type AwsAutoscalingNotificationSpec struct {
	TopicArn      string   `json:"topic_arn"`
	GroupNames    []string `json:"group_names"`
	Notifications []string `json:"notifications"`
}

type AwsAutoscalingNotificationStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AwsAutoscalingNotificationList is a list of AwsAutoscalingNotifications
type AwsAutoscalingNotificationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsAutoscalingNotification CRD objects
	Items []AwsAutoscalingNotification `json:"items,omitempty"`
}
