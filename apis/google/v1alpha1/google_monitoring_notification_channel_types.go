package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type GoogleMonitoringNotificationChannel struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              GoogleMonitoringNotificationChannelSpec   `json:"spec,omitempty"`
	Status            GoogleMonitoringNotificationChannelStatus `json:"status,omitempty"`
}

type GoogleMonitoringNotificationChannelSpec struct {
	Enabled            bool              `json:"enabled"`
	UserLabels         map[string]string `json:"user_labels"`
	Name               string            `json:"name"`
	VerificationStatus string            `json:"verification_status"`
	Type               string            `json:"type"`
	Description        string            `json:"description"`
	Labels             map[string]string `json:"labels"`
	Project            string            `json:"project"`
	DisplayName        string            `json:"display_name"`
}

type GoogleMonitoringNotificationChannelStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// GoogleMonitoringNotificationChannelList is a list of GoogleMonitoringNotificationChannels
type GoogleMonitoringNotificationChannelList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of GoogleMonitoringNotificationChannel CRD objects
	Items []GoogleMonitoringNotificationChannel `json:"items,omitempty"`
}
