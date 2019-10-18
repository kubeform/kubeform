package v1alpha1

import (
	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	base "kubeform.dev/kubeform/apis/base/v1alpha1"
)

// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Phase",type=string,JSONPath=`.status.phase`

type NotificationHubNamespace_ struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              NotificationHubNamespace_Spec   `json:"spec,omitempty"`
	Status            NotificationHubNamespace_Status `json:"status,omitempty"`
}

type NotificationHubNamespace_SpecSku struct {
	Name string `json:"name" tf:"name"`
}

type NotificationHubNamespace_Spec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	Enabled           bool   `json:"enabled,omitempty" tf:"enabled,omitempty"`
	Location          string `json:"location" tf:"location"`
	Name              string `json:"name" tf:"name"`
	NamespaceType     string `json:"namespaceType" tf:"namespace_type"`
	ResourceGroupName string `json:"resourceGroupName" tf:"resource_group_name"`
	// +optional
	ServicebusEndpoint string `json:"servicebusEndpoint,omitempty" tf:"servicebus_endpoint,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	// Deprecated
	Sku []NotificationHubNamespace_SpecSku `json:"sku,omitempty" tf:"sku,omitempty"`
	// +optional
	SkuName string `json:"skuName,omitempty" tf:"sku_name,omitempty"`
}

type NotificationHubNamespace_Status struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *NotificationHubNamespace_Spec `json:"output,omitempty"`
	// +optional
	State *base.State `json:"state,omitempty"`
	// +optional
	Phase base.Phase `json:"phase,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// NotificationHubNamespace_List is a list of NotificationHubNamespace_s
type NotificationHubNamespace_List struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of NotificationHubNamespace_ CRD objects
	Items []NotificationHubNamespace_ `json:"items,omitempty"`
}
