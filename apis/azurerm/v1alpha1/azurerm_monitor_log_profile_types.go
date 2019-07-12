package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AzurermMonitorLogProfile struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AzurermMonitorLogProfileSpec   `json:"spec,omitempty"`
	Status            AzurermMonitorLogProfileStatus `json:"status,omitempty"`
}

type AzurermMonitorLogProfileSpecRetentionPolicy struct {
	Enabled bool `json:"enabled"`
	Days    int  `json:"days"`
}

type AzurermMonitorLogProfileSpec struct {
	Name             string                         `json:"name"`
	StorageAccountId string                         `json:"storage_account_id"`
	ServicebusRuleId string                         `json:"servicebus_rule_id"`
	Locations        []string                       `json:"locations"`
	Categories       []string                       `json:"categories"`
	RetentionPolicy  []AzurermMonitorLogProfileSpec `json:"retention_policy"`
}



type AzurermMonitorLogProfileStatus struct {
    Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AzurermMonitorLogProfileList is a list of AzurermMonitorLogProfiles
type AzurermMonitorLogProfileList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AzurermMonitorLogProfile CRD objects
	Items []AzurermMonitorLogProfile `json:"items,omitempty"`
}