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

type GoogleStorageNotification struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              GoogleStorageNotificationSpec   `json:"spec,omitempty"`
	Status            GoogleStorageNotificationStatus `json:"status,omitempty"`
}

type GoogleStorageNotificationSpec struct {
	SelfLink         string            `json:"self_link"`
	Bucket           string            `json:"bucket"`
	PayloadFormat    string            `json:"payload_format"`
	Topic            string            `json:"topic"`
	CustomAttributes map[string]string `json:"custom_attributes"`
	EventTypes       []string          `json:"event_types"`
	ObjectNamePrefix string            `json:"object_name_prefix"`
}



type GoogleStorageNotificationStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// GoogleStorageNotificationList is a list of GoogleStorageNotifications
type GoogleStorageNotificationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of GoogleStorageNotification CRD objects
	Items []GoogleStorageNotification `json:"items,omitempty"`
}