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

type NotificationHub struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              NotificationHubSpec   `json:"spec,omitempty"`
	Status            NotificationHubStatus `json:"status,omitempty"`
}

type NotificationHubSpecApnsCredential struct {
	ApplicationMode string `json:"application_mode"`
	BundleId        string `json:"bundle_id"`
	KeyId           string `json:"key_id"`
	TeamId          string `json:"team_id"`
	Token           string `json:"token"`
}

type NotificationHubSpecGcmCredential struct {
	ApiKey string `json:"api_key"`
}

type NotificationHubSpec struct {
	// +optional
	// +kubebuilder:validation:MaxItems=1
	ApnsCredential *[]NotificationHubSpec `json:"apns_credential,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	GcmCredential     *[]NotificationHubSpec `json:"gcm_credential,omitempty"`
	Location          string                 `json:"location"`
	Name              string                 `json:"name"`
	NamespaceName     string                 `json:"namespace_name"`
	ResourceGroupName string                 `json:"resource_group_name"`
}

type NotificationHubStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// NotificationHubList is a list of NotificationHubs
type NotificationHubList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of NotificationHub CRD objects
	Items []NotificationHub `json:"items,omitempty"`
}
