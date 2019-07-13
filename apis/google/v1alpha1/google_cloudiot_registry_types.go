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

type GoogleCloudiotRegistry struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              GoogleCloudiotRegistrySpec   `json:"spec,omitempty"`
	Status            GoogleCloudiotRegistryStatus `json:"status,omitempty"`
}

type GoogleCloudiotRegistrySpecEventNotificationConfig struct {
	PubsubTopicName string `json:"pubsub_topic_name"`
}

type GoogleCloudiotRegistrySpecStateNotificationConfig struct {
	PubsubTopicName string `json:"pubsub_topic_name"`
}

type GoogleCloudiotRegistrySpecMqttConfig struct {
	MqttEnabledState string `json:"mqtt_enabled_state"`
}

type GoogleCloudiotRegistrySpecHttpConfig struct {
	HttpEnabledState string `json:"http_enabled_state"`
}

type GoogleCloudiotRegistrySpecCredentialsPublicKeyCertificate struct {
	Format      string `json:"format"`
	Certificate string `json:"certificate"`
}

type GoogleCloudiotRegistrySpecCredentials struct {
	PublicKeyCertificate map[string]string `json:"public_key_certificate"`
}

type GoogleCloudiotRegistrySpec struct {
	Name                    string                       `json:"name"`
	Project                 string                       `json:"project"`
	Region                  string                       `json:"region"`
	EventNotificationConfig map[string]string            `json:"event_notification_config"`
	StateNotificationConfig map[string]string            `json:"state_notification_config"`
	MqttConfig              map[string]string            `json:"mqtt_config"`
	HttpConfig              map[string]string            `json:"http_config"`
	Credentials             []GoogleCloudiotRegistrySpec `json:"credentials"`
}



type GoogleCloudiotRegistryStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// GoogleCloudiotRegistryList is a list of GoogleCloudiotRegistrys
type GoogleCloudiotRegistryList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of GoogleCloudiotRegistry CRD objects
	Items []GoogleCloudiotRegistry `json:"items,omitempty"`
}