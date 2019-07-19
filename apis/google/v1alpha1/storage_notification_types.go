package v1alpha1

import (
	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

type StorageNotification struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              StorageNotificationSpec   `json:"spec,omitempty"`
	Status            StorageNotificationStatus `json:"status,omitempty"`
}

type StorageNotificationSpec struct {
	Bucket string `json:"bucket" tf:"bucket"`
	// +optional
	CustomAttributes map[string]string `json:"customAttributes,omitempty" tf:"custom_attributes,omitempty"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	EventTypes []string `json:"eventTypes,omitempty" tf:"event_types,omitempty"`
	// +optional
	ObjectNamePrefix string                    `json:"objectNamePrefix,omitempty" tf:"object_name_prefix,omitempty"`
	PayloadFormat    string                    `json:"payloadFormat" tf:"payload_format"`
	Topic            string                    `json:"topic" tf:"topic"`
	ProviderRef      core.LocalObjectReference `json:"providerRef" tf:"-"`
}

type StorageNotificationStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	TFState *runtime.RawExtension `json:"tfState,omitempty"`
	Output  *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// StorageNotificationList is a list of StorageNotifications
type StorageNotificationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of StorageNotification CRD objects
	Items []StorageNotification `json:"items,omitempty"`
}
