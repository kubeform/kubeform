package v1alpha1

import (
	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

type ServicebusQueueAuthorizationRule struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ServicebusQueueAuthorizationRuleSpec   `json:"spec,omitempty"`
	Status            ServicebusQueueAuthorizationRuleStatus `json:"status,omitempty"`
}

type ServicebusQueueAuthorizationRuleSpec struct {
	// +optional
	Listen bool `json:"listen,omitempty" tf:"listen,omitempty"`
	// +optional
	Manage            bool   `json:"manage,omitempty" tf:"manage,omitempty"`
	Name              string `json:"name" tf:"name"`
	NamespaceName     string `json:"namespaceName" tf:"namespace_name"`
	QueueName         string `json:"queueName" tf:"queue_name"`
	ResourceGroupName string `json:"resourceGroupName" tf:"resource_group_name"`
	// +optional
	Send        bool                      `json:"send,omitempty" tf:"send,omitempty"`
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`
}

type ServicebusQueueAuthorizationRuleStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	TFState     *runtime.RawExtension `json:"tfState,omitempty"`
	TFStateHash string                `json:"tfStateHash,omitempty"`
	Output      *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// ServicebusQueueAuthorizationRuleList is a list of ServicebusQueueAuthorizationRules
type ServicebusQueueAuthorizationRuleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of ServicebusQueueAuthorizationRule CRD objects
	Items []ServicebusQueueAuthorizationRule `json:"items,omitempty"`
}
