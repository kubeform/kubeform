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

type AzurermNotificationHubAuthorizationRule struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AzurermNotificationHubAuthorizationRuleSpec   `json:"spec,omitempty"`
	Status            AzurermNotificationHubAuthorizationRuleStatus `json:"status,omitempty"`
}

type AzurermNotificationHubAuthorizationRuleSpec struct {
	Manage              bool   `json:"manage"`
	Listen              bool   `json:"listen"`
	PrimaryAccessKey    string `json:"primary_access_key"`
	NotificationHubName string `json:"notification_hub_name"`
	NamespaceName       string `json:"namespace_name"`
	ResourceGroupName   string `json:"resource_group_name"`
	Send                bool   `json:"send"`
	SecondaryAccessKey  string `json:"secondary_access_key"`
	Name                string `json:"name"`
}

type AzurermNotificationHubAuthorizationRuleStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AzurermNotificationHubAuthorizationRuleList is a list of AzurermNotificationHubAuthorizationRules
type AzurermNotificationHubAuthorizationRuleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AzurermNotificationHubAuthorizationRule CRD objects
	Items []AzurermNotificationHubAuthorizationRule `json:"items,omitempty"`
}
