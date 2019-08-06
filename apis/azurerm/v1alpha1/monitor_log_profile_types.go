package v1alpha1

import (
	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"kubeform.dev/kubeform/apis"
)

// +genclient
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
	Days    int  `json:"days,omitempty" tf:"days,omitempty"`
	Enabled bool `json:"enabled" tf:"enabled"`
}

type MonitorLogProfileSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +kubebuilder:validation:MinItems=1
	// +kubebuilder:validation:UniqueItems=true
	Categories []string `json:"categories" tf:"categories"`
	// +kubebuilder:validation:MinItems=1
	// +kubebuilder:validation:UniqueItems=true
	Locations []string `json:"locations" tf:"locations"`
	Name      string   `json:"name" tf:"name"`
	// +kubebuilder:validation:MaxItems=1
	RetentionPolicy []MonitorLogProfileSpecRetentionPolicy `json:"retentionPolicy" tf:"retention_policy"`
	// +optional
	ServicebusRuleID string `json:"servicebusRuleID,omitempty" tf:"servicebus_rule_id,omitempty"`
	// +optional
	StorageAccountID string `json:"storageAccountID,omitempty" tf:"storage_account_id,omitempty"`
}

type MonitorLogProfileStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *MonitorLogProfileSpec `json:"output,omitempty"`
	// +optional
	State *apis.State `json:"state,omitempty"`
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
