package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AzurermServicebusTopicAuthorizationRule struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AzurermServicebusTopicAuthorizationRuleSpec   `json:"spec,omitempty"`
	Status            AzurermServicebusTopicAuthorizationRuleStatus `json:"status,omitempty"`
}

type AzurermServicebusTopicAuthorizationRuleSpec struct {
	PrimaryKey                string `json:"primary_key"`
	SecondaryKey              string `json:"secondary_key"`
	NamespaceName             string `json:"namespace_name"`
	TopicName                 string `json:"topic_name"`
	ResourceGroupName         string `json:"resource_group_name"`
	Name                      string `json:"name"`
	Send                      bool   `json:"send"`
	Manage                    bool   `json:"manage"`
	PrimaryConnectionString   string `json:"primary_connection_string"`
	SecondaryConnectionString string `json:"secondary_connection_string"`
	Listen                    bool   `json:"listen"`
}



type AzurermServicebusTopicAuthorizationRuleStatus struct {
    Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AzurermServicebusTopicAuthorizationRuleList is a list of AzurermServicebusTopicAuthorizationRules
type AzurermServicebusTopicAuthorizationRuleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AzurermServicebusTopicAuthorizationRule CRD objects
	Items []AzurermServicebusTopicAuthorizationRule `json:"items,omitempty"`
}