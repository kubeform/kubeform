package v1alpha1

import (
	base "kubeform.dev/kubeform/apis/base/v1alpha1"

	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Phase",type=string,JSONPath=`.status.phase`

type NotificationHubAuthorizationRule struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              NotificationHubAuthorizationRuleSpec   `json:"spec,omitempty"`
	Status            NotificationHubAuthorizationRuleStatus `json:"status,omitempty"`
}

type NotificationHubAuthorizationRuleSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	Listen bool `json:"listen,omitempty" tf:"listen,omitempty"`
	// +optional
	Manage              bool   `json:"manage,omitempty" tf:"manage,omitempty"`
	Name                string `json:"name" tf:"name"`
	NamespaceName       string `json:"namespaceName" tf:"namespace_name"`
	NotificationHubName string `json:"notificationHubName" tf:"notification_hub_name"`
	// +optional
	PrimaryAccessKey  string `json:"primaryAccessKey,omitempty" tf:"primary_access_key,omitempty"`
	ResourceGroupName string `json:"resourceGroupName" tf:"resource_group_name"`
	// +optional
	SecondaryAccessKey string `json:"secondaryAccessKey,omitempty" tf:"secondary_access_key,omitempty"`
	// +optional
	Send bool `json:"send,omitempty" tf:"send,omitempty"`
}

type NotificationHubAuthorizationRuleStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *NotificationHubAuthorizationRuleSpec `json:"output,omitempty"`
	// +optional
	State *base.State `json:"state,omitempty"`
	// +optional
	Phase base.Phase `json:"phase,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// NotificationHubAuthorizationRuleList is a list of NotificationHubAuthorizationRules
type NotificationHubAuthorizationRuleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of NotificationHubAuthorizationRule CRD objects
	Items []NotificationHubAuthorizationRule `json:"items,omitempty"`
}
