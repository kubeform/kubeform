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

type LogicAppTriggerRecurrence struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              LogicAppTriggerRecurrenceSpec   `json:"spec,omitempty"`
	Status            LogicAppTriggerRecurrenceStatus `json:"status,omitempty"`
}

type LogicAppTriggerRecurrenceSpec struct {
	Frequency   string                    `json:"frequency" tf:"frequency"`
	Interval    int                       `json:"interval" tf:"interval"`
	LogicAppID  string                    `json:"logicAppID" tf:"logic_app_id"`
	Name        string                    `json:"name" tf:"name"`
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`
}

type LogicAppTriggerRecurrenceStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	TFState *runtime.RawExtension `json:"tfState,omitempty"`
	Output  *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// LogicAppTriggerRecurrenceList is a list of LogicAppTriggerRecurrences
type LogicAppTriggerRecurrenceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of LogicAppTriggerRecurrence CRD objects
	Items []LogicAppTriggerRecurrence `json:"items,omitempty"`
}
