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

type EventhubNamespaceAuthorizationRule struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              EventhubNamespaceAuthorizationRuleSpec   `json:"spec,omitempty"`
	Status            EventhubNamespaceAuthorizationRuleStatus `json:"status,omitempty"`
}

type EventhubNamespaceAuthorizationRuleSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	SecretRef *core.LocalObjectReference `json:"secretRef,omitempty" tf:"-"`

	// +optional
	Listen bool `json:"listen,omitempty" tf:"listen,omitempty"`
	// +optional
	// Deprecated
	Location string `json:"location,omitempty" tf:"location,omitempty"`
	// +optional
	Manage        bool   `json:"manage,omitempty" tf:"manage,omitempty"`
	Name          string `json:"name" tf:"name"`
	NamespaceName string `json:"namespaceName" tf:"namespace_name"`
	// +optional
	PrimaryConnectionString string `json:"-" sensitive:"true" tf:"primary_connection_string,omitempty"`
	// +optional
	PrimaryKey        string `json:"-" sensitive:"true" tf:"primary_key,omitempty"`
	ResourceGroupName string `json:"resourceGroupName" tf:"resource_group_name"`
	// +optional
	SecondaryConnectionString string `json:"-" sensitive:"true" tf:"secondary_connection_string,omitempty"`
	// +optional
	SecondaryKey string `json:"-" sensitive:"true" tf:"secondary_key,omitempty"`
	// +optional
	Send bool `json:"send,omitempty" tf:"send,omitempty"`
}

type EventhubNamespaceAuthorizationRuleStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *EventhubNamespaceAuthorizationRuleSpec `json:"output,omitempty"`
	// +optional
	State *apis.State `json:"state,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// EventhubNamespaceAuthorizationRuleList is a list of EventhubNamespaceAuthorizationRules
type EventhubNamespaceAuthorizationRuleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of EventhubNamespaceAuthorizationRule CRD objects
	Items []EventhubNamespaceAuthorizationRule `json:"items,omitempty"`
}
