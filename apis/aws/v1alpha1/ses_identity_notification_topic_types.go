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

type SesIdentityNotificationTopic struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SesIdentityNotificationTopicSpec   `json:"spec,omitempty"`
	Status            SesIdentityNotificationTopicStatus `json:"status,omitempty"`
}

type SesIdentityNotificationTopicSpec struct {
	Identity string `json:"identity"`
	// +optional
	IncludeOriginalHeaders bool   `json:"include_original_headers,omitempty"`
	NotificationType       string `json:"notification_type"`
	// +optional
	TopicArn string `json:"topic_arn,omitempty"`
}

type SesIdentityNotificationTopicStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
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
