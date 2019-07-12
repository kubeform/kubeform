package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AzurermMonitorDiagnosticSetting struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AzurermMonitorDiagnosticSettingSpec   `json:"spec,omitempty"`
	Status            AzurermMonitorDiagnosticSettingStatus `json:"status,omitempty"`
}

type AzurermMonitorDiagnosticSettingSpecLogRetentionPolicy struct {
	Days    int  `json:"days"`
	Enabled bool `json:"enabled"`
}

type AzurermMonitorDiagnosticSettingSpecLog struct {
	RetentionPolicy []AzurermMonitorDiagnosticSettingSpecLog `json:"retention_policy"`
	Category        string                                   `json:"category"`
	Enabled         bool                                     `json:"enabled"`
}

type AzurermMonitorDiagnosticSettingSpecMetricRetentionPolicy struct {
	Enabled bool `json:"enabled"`
	Days    int  `json:"days"`
}

type AzurermMonitorDiagnosticSettingSpecMetric struct {
	Category        string                                      `json:"category"`
	Enabled         bool                                        `json:"enabled"`
	RetentionPolicy []AzurermMonitorDiagnosticSettingSpecMetric `json:"retention_policy"`
}

type AzurermMonitorDiagnosticSettingSpec struct {
	TargetResourceId            string                                `json:"target_resource_id"`
	EventhubName                string                                `json:"eventhub_name"`
	EventhubAuthorizationRuleId string                                `json:"eventhub_authorization_rule_id"`
	LogAnalyticsWorkspaceId     string                                `json:"log_analytics_workspace_id"`
	StorageAccountId            string                                `json:"storage_account_id"`
	Log                         []AzurermMonitorDiagnosticSettingSpec `json:"log"`
	Metric                      []AzurermMonitorDiagnosticSettingSpec `json:"metric"`
	Name                        string                                `json:"name"`
}



type AzurermMonitorDiagnosticSettingStatus struct {
    Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AzurermMonitorDiagnosticSettingList is a list of AzurermMonitorDiagnosticSettings
type AzurermMonitorDiagnosticSettingList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AzurermMonitorDiagnosticSetting CRD objects
	Items []AzurermMonitorDiagnosticSetting `json:"items,omitempty"`
}