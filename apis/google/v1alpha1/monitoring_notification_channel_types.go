package v1alpha1

import (
	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

type MonitoringNotificationChannel struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              MonitoringNotificationChannelSpec   `json:"spec,omitempty"`
	Status            MonitoringNotificationChannelStatus `json:"status,omitempty"`
}

type MonitoringNotificationChannelSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	// +optional
	Description string `json:"description,omitempty" tf:"description,omitempty"`
	DisplayName string `json:"displayName" tf:"display_name"`
	// +optional
	Enabled bool `json:"enabled,omitempty" tf:"enabled,omitempty"`
	// +optional
	Labels map[string]string `json:"labels,omitempty" tf:"labels,omitempty"`
	// +optional
	Project string `json:"project,omitempty" tf:"project,omitempty"`
	Type    string `json:"type" tf:"type"`
	// +optional
	UserLabels map[string]string `json:"userLabels,omitempty" tf:"user_labels,omitempty"`
}

type MonitoringNotificationChannelStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	State *MonitoringNotificationChannelSpec `json:"state,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// MonitoringNotificationChannelList is a list of MonitoringNotificationChannels
type MonitoringNotificationChannelList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of MonitoringNotificationChannel CRD objects
	Items []MonitoringNotificationChannel `json:"items,omitempty"`
}
