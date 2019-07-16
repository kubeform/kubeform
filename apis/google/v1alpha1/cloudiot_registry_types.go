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

type CloudiotRegistry struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              CloudiotRegistrySpec   `json:"spec,omitempty"`
	Status            CloudiotRegistryStatus `json:"status,omitempty"`
}

type CloudiotRegistrySpecCredentialsPublicKeyCertificate struct {
	Certificate string `json:"certificate"`
	Format      string `json:"format"`
}

type CloudiotRegistrySpecCredentials struct {
	// +optional
	PublicKeyCertificate map[string]CloudiotRegistrySpecCredentialsPublicKeyCertificate `json:"public_key_certificate,omitempty"`
}

type CloudiotRegistrySpecEventNotificationConfig struct {
	PubsubTopicName string `json:"pubsub_topic_name"`
}

type CloudiotRegistrySpecStateNotificationConfig struct {
	PubsubTopicName string `json:"pubsub_topic_name"`
}

type CloudiotRegistrySpec struct {
	// +optional
	// +kubebuilder:validation:MaxItems=10
	Credentials *[]CloudiotRegistrySpec `json:"credentials,omitempty"`
	// +optional
	EventNotificationConfig map[string]CloudiotRegistrySpecEventNotificationConfig `json:"event_notification_config,omitempty"`
	Name                    string                                                 `json:"name"`
	// +optional
	StateNotificationConfig map[string]CloudiotRegistrySpecStateNotificationConfig `json:"state_notification_config,omitempty"`
}

type CloudiotRegistryStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// CloudiotRegistryList is a list of CloudiotRegistrys
type CloudiotRegistryList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of CloudiotRegistry CRD objects
	Items []CloudiotRegistry `json:"items,omitempty"`
}
