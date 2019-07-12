package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AzurermNotificationHubNamespace struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AzurermNotificationHubNamespaceSpec   `json:"spec,omitempty"`
	Status            AzurermNotificationHubNamespaceStatus `json:"status,omitempty"`
}

type AzurermNotificationHubNamespaceSpecSku struct {
	Name string `json:"name"`
}

type AzurermNotificationHubNamespaceSpec struct {
	Sku                []AzurermNotificationHubNamespaceSpec `json:"sku"`
	SkuName            string                                `json:"sku_name"`
	Enabled            bool                                  `json:"enabled"`
	NamespaceType      string                                `json:"namespace_type"`
	ServicebusEndpoint string                                `json:"servicebus_endpoint"`
	Name               string                                `json:"name"`
	ResourceGroupName  string                                `json:"resource_group_name"`
	Location           string                                `json:"location"`
}



type AzurermNotificationHubNamespaceStatus struct {
    Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AzurermNotificationHubNamespaceList is a list of AzurermNotificationHubNamespaces
type AzurermNotificationHubNamespaceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AzurermNotificationHubNamespace CRD objects
	Items []AzurermNotificationHubNamespace `json:"items,omitempty"`
}