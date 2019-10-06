package v1alpha1

import (
	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	base "kubeform.dev/kubeform/apis/base/v1alpha1"
)

// +genclient
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
	Certificate string `json:"certificate" tf:"certificate"`
	Format      string `json:"format" tf:"format"`
}

type CloudiotRegistrySpecCredentials struct {
	// +optional
	PublicKeyCertificate map[string]CloudiotRegistrySpecCredentialsPublicKeyCertificate `json:"publicKeyCertificate,omitempty" tf:"public_key_certificate,omitempty"`
}

type CloudiotRegistrySpecEventNotificationConfig struct {
	PubsubTopicName string `json:"pubsubTopicName" tf:"pubsub_topic_name"`
}

type CloudiotRegistrySpecHttpConfig struct {
	HttpEnabledState string `json:"httpEnabledState" tf:"http_enabled_state"`
}

type CloudiotRegistrySpecMqttConfig struct {
	MqttEnabledState string `json:"mqttEnabledState" tf:"mqtt_enabled_state"`
}

type CloudiotRegistrySpecStateNotificationConfig struct {
	PubsubTopicName string `json:"pubsubTopicName" tf:"pubsub_topic_name"`
}

type CloudiotRegistrySpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	// +kubebuilder:validation:MaxItems=10
	Credentials []CloudiotRegistrySpecCredentials `json:"credentials,omitempty" tf:"credentials,omitempty"`
	// +optional
	EventNotificationConfig map[string]CloudiotRegistrySpecEventNotificationConfig `json:"eventNotificationConfig,omitempty" tf:"event_notification_config,omitempty"`
	// +optional
	HttpConfig map[string]CloudiotRegistrySpecHttpConfig `json:"httpConfig,omitempty" tf:"http_config,omitempty"`
	// +optional
	MqttConfig map[string]CloudiotRegistrySpecMqttConfig `json:"mqttConfig,omitempty" tf:"mqtt_config,omitempty"`
	Name       string                                    `json:"name" tf:"name"`
	// +optional
	Project string `json:"project,omitempty" tf:"project,omitempty"`
	// +optional
	Region string `json:"region,omitempty" tf:"region,omitempty"`
	// +optional
	StateNotificationConfig map[string]CloudiotRegistrySpecStateNotificationConfig `json:"stateNotificationConfig,omitempty" tf:"state_notification_config,omitempty"`
}

type CloudiotRegistryStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *CloudiotRegistrySpec `json:"output,omitempty"`
	// +optional
	State *base.State `json:"state,omitempty"`
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
