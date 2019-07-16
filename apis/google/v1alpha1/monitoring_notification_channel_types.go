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

type MonitoringNotificationChannel struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              MonitoringNotificationChannelSpec   `json:"spec,omitempty"`
	Status            MonitoringNotificationChannelStatus `json:"status,omitempty"`
}

type MonitoringNotificationChannelSpec struct {
	// +optional
	Description string `json:"description,omitempty"`
	DisplayName string `json:"display_name"`
	// +optional
	Enabled bool `json:"enabled,omitempty"`
	// +optional
	Labels map[string]string `json:"labels,omitempty"`
	Type   string            `json:"type"`
	// +optional
	UserLabels map[string]string `json:"user_labels,omitempty"`
}

type MonitoringNotificationChannelStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
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
