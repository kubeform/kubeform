package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AzurermNotificationHub struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AzurermNotificationHubSpec   `json:"spec,omitempty"`
	Status            AzurermNotificationHubStatus `json:"status,omitempty"`
}

type AzurermNotificationHubSpecApnsCredential struct {
	ApplicationMode string `json:"application_mode"`
	BundleId        string `json:"bundle_id"`
	KeyId           string `json:"key_id"`
	TeamId          string `json:"team_id"`
	Token           string `json:"token"`
}

type AzurermNotificationHubSpecGcmCredential struct {
	ApiKey string `json:"api_key"`
}

type AzurermNotificationHubSpec struct {
	Location          string                       `json:"location"`
	ApnsCredential    []AzurermNotificationHubSpec `json:"apns_credential"`
	GcmCredential     []AzurermNotificationHubSpec `json:"gcm_credential"`
	Name              string                       `json:"name"`
	NamespaceName     string                       `json:"namespace_name"`
	ResourceGroupName string                       `json:"resource_group_name"`
}

type AzurermNotificationHubStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AzurermNotificationHubList is a list of AzurermNotificationHubs
type AzurermNotificationHubList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AzurermNotificationHub CRD objects
	Items []AzurermNotificationHub `json:"items,omitempty"`
}
