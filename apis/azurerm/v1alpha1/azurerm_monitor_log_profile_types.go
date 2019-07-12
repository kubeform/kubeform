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
	Locations        []string                       `json:"locations"`
	Categories       []string                       `json:"categories"`
	RetentionPolicy  []AzurermMonitorLogProfileSpec `json:"retention_policy"`
	Name             string                         `json:"name"`
	StorageAccountId string                         `json:"storage_account_id"`
	ServicebusRuleId string                         `json:"servicebus_rule_id"`
}

type AzurermMonitorLogProfileStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AzurermMonitorLogProfileList is a list of AzurermMonitorLogProfiles
type AzurermMonitorLogProfileList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AzurermMonitorLogProfile CRD objects
	Items []AzurermMonitorLogProfile `json:"items,omitempty"`
}
