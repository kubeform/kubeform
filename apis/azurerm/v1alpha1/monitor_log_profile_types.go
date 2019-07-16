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

type MonitorLogProfile struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              MonitorLogProfileSpec   `json:"spec,omitempty"`
	Status            MonitorLogProfileStatus `json:"status,omitempty"`
}

type MonitorLogProfileSpecRetentionPolicy struct {
	// +optional
	Days    int  `json:"days,omitempty"`
	Enabled bool `json:"enabled"`
}

type MonitorLogProfileSpec struct {
	// +kubebuilder:validation:MinItems=1
	// +kubebuilder:validation:UniqueItems=true
	Categories []string `json:"categories"`
	// +kubebuilder:validation:MinItems=1
	// +kubebuilder:validation:UniqueItems=true
	Locations []string `json:"locations"`
	Name      string   `json:"name"`
	// +kubebuilder:validation:MaxItems=1
	RetentionPolicy []MonitorLogProfileSpec `json:"retention_policy"`
	// +optional
	ServicebusRuleId string `json:"servicebus_rule_id,omitempty"`
	// +optional
	StorageAccountId string `json:"storage_account_id,omitempty"`
}

type MonitorLogProfileStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// MonitorLogProfileList is a list of MonitorLogProfiles
type MonitorLogProfileList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of MonitorLogProfile CRD objects
	Items []MonitorLogProfile `json:"items,omitempty"`
}
