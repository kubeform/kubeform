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

type MonitorDiagnosticSetting struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              MonitorDiagnosticSettingSpec   `json:"spec,omitempty"`
	Status            MonitorDiagnosticSettingStatus `json:"status,omitempty"`
}

type MonitorDiagnosticSettingSpecLogRetentionPolicy struct {
	// +optional
	Days    int  `json:"days,omitempty"`
	Enabled bool `json:"enabled"`
}

type MonitorDiagnosticSettingSpecLog struct {
	Category string `json:"category"`
	// +optional
	Enabled bool `json:"enabled,omitempty"`
	// +kubebuilder:validation:MaxItems=1
	RetentionPolicy []MonitorDiagnosticSettingSpecLog `json:"retention_policy"`
}

type MonitorDiagnosticSettingSpecMetricRetentionPolicy struct {
	// +optional
	Days    int  `json:"days,omitempty"`
	Enabled bool `json:"enabled"`
}

type MonitorDiagnosticSettingSpecMetric struct {
	Category string `json:"category"`
	// +optional
	Enabled bool `json:"enabled,omitempty"`
	// +kubebuilder:validation:MaxItems=1
	RetentionPolicy []MonitorDiagnosticSettingSpecMetric `json:"retention_policy"`
}

type MonitorDiagnosticSettingSpec struct {
	// +optional
	EventhubAuthorizationRuleId string `json:"eventhub_authorization_rule_id,omitempty"`
	// +optional
	EventhubName string `json:"eventhub_name,omitempty"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	Log *[]MonitorDiagnosticSettingSpec `json:"log,omitempty"`
	// +optional
	LogAnalyticsWorkspaceId string `json:"log_analytics_workspace_id,omitempty"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	Metric *[]MonitorDiagnosticSettingSpec `json:"metric,omitempty"`
	Name   string                          `json:"name"`
	// +optional
	StorageAccountId string `json:"storage_account_id,omitempty"`
	TargetResourceId string `json:"target_resource_id"`
}

type MonitorDiagnosticSettingStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// MonitorDiagnosticSettingList is a list of MonitorDiagnosticSettings
type MonitorDiagnosticSettingList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of MonitorDiagnosticSetting CRD objects
	Items []MonitorDiagnosticSetting `json:"items,omitempty"`
}
