package v1alpha1

import (
	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

type SesIdentityNotificationTopic struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SesIdentityNotificationTopicSpec   `json:"spec,omitempty"`
	Status            SesIdentityNotificationTopicStatus `json:"status,omitempty"`
}

type SesIdentityNotificationTopicSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	Identity string `json:"identity" tf:"identity"`
	// +optional
	IncludeOriginalHeaders bool   `json:"includeOriginalHeaders,omitempty" tf:"include_original_headers,omitempty"`
	NotificationType       string `json:"notificationType" tf:"notification_type"`
	// +optional
	TopicArn string `json:"topicArn,omitempty" tf:"topic_arn,omitempty"`
}

type SesIdentityNotificationTopicStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64  `json:"observedGeneration,omitempty"`
	TFState            string `json:"tfState,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// SesIdentityNotificationTopicList is a list of SesIdentityNotificationTopics
type SesIdentityNotificationTopicList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of SesIdentityNotificationTopic CRD objects
	Items []SesIdentityNotificationTopic `json:"items,omitempty"`
}
