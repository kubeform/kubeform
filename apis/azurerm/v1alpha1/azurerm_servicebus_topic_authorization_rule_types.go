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

type AzurermServicebusTopicAuthorizationRule struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AzurermServicebusTopicAuthorizationRuleSpec   `json:"spec,omitempty"`
	Status            AzurermServicebusTopicAuthorizationRuleStatus `json:"status,omitempty"`
}

type AzurermServicebusTopicAuthorizationRuleSpec struct {
	Listen                    bool   `json:"listen"`
	Send                      bool   `json:"send"`
	Manage                    bool   `json:"manage"`
	Name                      string `json:"name"`
	NamespaceName             string `json:"namespace_name"`
	TopicName                 string `json:"topic_name"`
	SecondaryKey              string `json:"secondary_key"`
	SecondaryConnectionString string `json:"secondary_connection_string"`
	PrimaryConnectionString   string `json:"primary_connection_string"`
	ResourceGroupName         string `json:"resource_group_name"`
	PrimaryKey                string `json:"primary_key"`
}

type AzurermServicebusTopicAuthorizationRuleStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AzurermServicebusTopicAuthorizationRuleList is a list of AzurermServicebusTopicAuthorizationRules
type AzurermServicebusTopicAuthorizationRuleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AzurermServicebusTopicAuthorizationRule CRD objects
	Items []AzurermServicebusTopicAuthorizationRule `json:"items,omitempty"`
}
