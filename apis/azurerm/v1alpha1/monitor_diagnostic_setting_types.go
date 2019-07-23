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

type MonitorDiagnosticSetting struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              MonitorDiagnosticSettingSpec   `json:"spec,omitempty"`
	Status            MonitorDiagnosticSettingStatus `json:"status,omitempty"`
}

type MonitorDiagnosticSettingSpecLogRetentionPolicy struct {
	// +optional
	Days    int  `json:"days,omitempty" tf:"days,omitempty"`
	Enabled bool `json:"enabled" tf:"enabled"`
}

type MonitorDiagnosticSettingSpecLog struct {
	Category string `json:"category" tf:"category"`
	// +optional
	Enabled bool `json:"enabled,omitempty" tf:"enabled,omitempty"`
	// +kubebuilder:validation:MaxItems=1
	RetentionPolicy []MonitorDiagnosticSettingSpecLogRetentionPolicy `json:"retentionPolicy" tf:"retention_policy"`
}

type MonitorDiagnosticSettingSpecMetricRetentionPolicy struct {
	// +optional
	Days    int  `json:"days,omitempty" tf:"days,omitempty"`
	Enabled bool `json:"enabled" tf:"enabled"`
}

type MonitorDiagnosticSettingSpecMetric struct {
	Category string `json:"category" tf:"category"`
	// +optional
	Enabled bool `json:"enabled,omitempty" tf:"enabled,omitempty"`
	// +kubebuilder:validation:MaxItems=1
	RetentionPolicy []MonitorDiagnosticSettingSpecMetricRetentionPolicy `json:"retentionPolicy" tf:"retention_policy"`
}

type MonitorDiagnosticSettingSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	// +optional
	EventhubAuthorizationRuleID string `json:"eventhubAuthorizationRuleID,omitempty" tf:"eventhub_authorization_rule_id,omitempty"`
	// +optional
	EventhubName string `json:"eventhubName,omitempty" tf:"eventhub_name,omitempty"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	Log []MonitorDiagnosticSettingSpecLog `json:"log,omitempty" tf:"log,omitempty"`
	// +optional
	LogAnalyticsWorkspaceID string `json:"logAnalyticsWorkspaceID,omitempty" tf:"log_analytics_workspace_id,omitempty"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	Metric []MonitorDiagnosticSettingSpecMetric `json:"metric,omitempty" tf:"metric,omitempty"`
	Name   string                               `json:"name" tf:"name"`
	// +optional
	StorageAccountID string `json:"storageAccountID,omitempty" tf:"storage_account_id,omitempty"`
	TargetResourceID string `json:"targetResourceID" tf:"target_resource_id"`
}

type MonitorDiagnosticSettingStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	TFState *runtime.RawExtension `json:"tfState,omitempty"`
	Output  *runtime.RawExtension `json:"output,omitempty"`
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
