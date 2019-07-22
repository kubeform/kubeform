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

type NotificationHub struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              NotificationHubSpec   `json:"spec,omitempty"`
	Status            NotificationHubStatus `json:"status,omitempty"`
}

type NotificationHubSpecApnsCredential struct {
	ApplicationMode string `json:"applicationMode" tf:"application_mode"`
	BundleID        string `json:"bundleID" tf:"bundle_id"`
	KeyID           string `json:"keyID" tf:"key_id"`
	TeamID          string `json:"teamID" tf:"team_id"`
}

type NotificationHubSpecGcmCredential struct{}

type NotificationHubSpec struct {
	Secret      *core.LocalObjectReference `json:"secret,omitempty" tf:"-"`
	ProviderRef core.LocalObjectReference  `json:"providerRef" tf:"-"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	ApnsCredential []NotificationHubSpecApnsCredential `json:"apnsCredential,omitempty" tf:"apns_credential,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	GcmCredential     []NotificationHubSpecGcmCredential `json:"gcmCredential,omitempty" tf:"gcm_credential,omitempty"`
	Location          string                             `json:"location" tf:"location"`
	Name              string                             `json:"name" tf:"name"`
	NamespaceName     string                             `json:"namespaceName" tf:"namespace_name"`
	ResourceGroupName string                             `json:"resourceGroupName" tf:"resource_group_name"`
}

type NotificationHubStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	TFState *runtime.RawExtension `json:"tfState,omitempty"`
	Output  *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// NotificationHubList is a list of NotificationHubs
type NotificationHubList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of NotificationHub CRD objects
	Items []NotificationHub `json:"items,omitempty"`
}
