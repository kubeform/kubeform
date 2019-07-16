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

type AutoscalingNotification struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AutoscalingNotificationSpec   `json:"spec,omitempty"`
	Status            AutoscalingNotificationStatus `json:"status,omitempty"`
}

type AutoscalingNotificationSpec struct {
	// +kubebuilder:validation:UniqueItems=true
	GroupNames []string `json:"group_names"`
	// +kubebuilder:validation:UniqueItems=true
	Notifications []string `json:"notifications"`
	TopicArn      string   `json:"topic_arn"`
}

type AutoscalingNotificationStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AutoscalingNotificationList is a list of AutoscalingNotifications
type AutoscalingNotificationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AutoscalingNotification CRD objects
	Items []AutoscalingNotification `json:"items,omitempty"`
}
