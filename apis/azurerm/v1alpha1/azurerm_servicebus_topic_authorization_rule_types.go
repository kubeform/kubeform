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
	TopicName                 string `json:"topic_name"`
	Listen                    bool   `json:"listen"`
	Send                      bool   `json:"send"`
	Manage                    bool   `json:"manage"`
	PrimaryKey                string `json:"primary_key"`
	Name                      string `json:"name"`
	NamespaceName             string `json:"namespace_name"`
	ResourceGroupName         string `json:"resource_group_name"`
	PrimaryConnectionString   string `json:"primary_connection_string"`
	SecondaryKey              string `json:"secondary_key"`
	SecondaryConnectionString string `json:"secondary_connection_string"`
}



type AzurermServicebusTopicAuthorizationRuleStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

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