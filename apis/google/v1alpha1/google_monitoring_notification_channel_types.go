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

type GoogleMonitoringNotificationChannel struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              GoogleMonitoringNotificationChannelSpec   `json:"spec,omitempty"`
	Status            GoogleMonitoringNotificationChannelStatus `json:"status,omitempty"`
}

type GoogleMonitoringNotificationChannelSpec struct {
	DisplayName        string            `json:"display_name"`
	Labels             map[string]string `json:"labels"`
	Name               string            `json:"name"`
	Project            string            `json:"project"`
	Type               string            `json:"type"`
	Description        string            `json:"description"`
	Enabled            bool              `json:"enabled"`
	UserLabels         map[string]string `json:"user_labels"`
	VerificationStatus string            `json:"verification_status"`
}



type GoogleMonitoringNotificationChannelStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// GoogleMonitoringNotificationChannelList is a list of GoogleMonitoringNotificationChannels
type GoogleMonitoringNotificationChannelList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of GoogleMonitoringNotificationChannel CRD objects
	Items []GoogleMonitoringNotificationChannel `json:"items,omitempty"`
}