package v1alpha1

import (
	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
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
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	// +kubebuilder:validation:UniqueItems=true
	GroupNames []string `json:"groupNames" tf:"group_names"`
	// +kubebuilder:validation:UniqueItems=true
	Notifications []string `json:"notifications" tf:"notifications"`
	TopicArn      string   `json:"topicArn" tf:"topic_arn"`
}

type AutoscalingNotificationStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64  `json:"observedGeneration,omitempty"`
	TFState            string `json:"tfState,omitempty"`
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
