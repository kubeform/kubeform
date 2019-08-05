package v1alpha1

import (
	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"kubeform.dev/kubeform/apis"
)

// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

type ServicebusTopicAuthorizationRule struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ServicebusTopicAuthorizationRuleSpec   `json:"spec,omitempty"`
	Status            ServicebusTopicAuthorizationRuleStatus `json:"status,omitempty"`
}

type ServicebusTopicAuthorizationRuleSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	Secret *core.LocalObjectReference `json:"secret,omitempty" tf:"-"`

	// +optional
	Listen bool `json:"listen,omitempty" tf:"listen,omitempty"`
	// +optional
	Manage                    bool   `json:"manage,omitempty" tf:"manage,omitempty"`
	Name                      string `json:"name" tf:"name"`
	NamespaceName             string `json:"namespaceName" tf:"namespace_name"`
	PrimaryConnectionString   string `json:"-" sensitive:"true" tf:"primary_connection_string"`
	PrimaryKey                string `json:"-" sensitive:"true" tf:"primary_key"`
	ResourceGroupName         string `json:"resourceGroupName" tf:"resource_group_name"`
	SecondaryConnectionString string `json:"-" sensitive:"true" tf:"secondary_connection_string"`
	SecondaryKey              string `json:"-" sensitive:"true" tf:"secondary_key"`
	// +optional
	Send      bool   `json:"send,omitempty" tf:"send,omitempty"`
	TopicName string `json:"topicName" tf:"topic_name"`
}

type ServicebusTopicAuthorizationRuleStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *ServicebusTopicAuthorizationRuleSpec `json:"output,omitempty"`
	// +optional
	State *apis.State `json:"state,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// ServicebusTopicAuthorizationRuleList is a list of ServicebusTopicAuthorizationRules
type ServicebusTopicAuthorizationRuleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of ServicebusTopicAuthorizationRule CRD objects
	Items []ServicebusTopicAuthorizationRule `json:"items,omitempty"`
}
