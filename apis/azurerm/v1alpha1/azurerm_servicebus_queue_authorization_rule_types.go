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

type AzurermServicebusQueueAuthorizationRule struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AzurermServicebusQueueAuthorizationRuleSpec   `json:"spec,omitempty"`
	Status            AzurermServicebusQueueAuthorizationRuleStatus `json:"status,omitempty"`
}

type AzurermServicebusQueueAuthorizationRuleSpec struct {
	PrimaryConnectionString   string `json:"primary_connection_string"`
	SecondaryConnectionString string `json:"secondary_connection_string"`
	Listen                    bool   `json:"listen"`
	Send                      bool   `json:"send"`
	NamespaceName             string `json:"namespace_name"`
	QueueName                 string `json:"queue_name"`
	SecondaryKey              string `json:"secondary_key"`
	Manage                    bool   `json:"manage"`
	PrimaryKey                string `json:"primary_key"`
	ResourceGroupName         string `json:"resource_group_name"`
	Name                      string `json:"name"`
}

type AzurermServicebusQueueAuthorizationRuleStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AzurermServicebusQueueAuthorizationRuleList is a list of AzurermServicebusQueueAuthorizationRules
type AzurermServicebusQueueAuthorizationRuleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AzurermServicebusQueueAuthorizationRule CRD objects
	Items []AzurermServicebusQueueAuthorizationRule `json:"items,omitempty"`
}
