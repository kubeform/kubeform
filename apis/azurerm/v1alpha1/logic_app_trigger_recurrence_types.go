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

type LogicAppTriggerRecurrence struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              LogicAppTriggerRecurrenceSpec   `json:"spec,omitempty"`
	Status            LogicAppTriggerRecurrenceStatus `json:"status,omitempty"`
}

type LogicAppTriggerRecurrenceSpec struct {
	Frequency  string `json:"frequency"`
	Interval   int    `json:"interval"`
	LogicAppId string `json:"logic_app_id"`
	Name       string `json:"name"`
}

type LogicAppTriggerRecurrenceStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
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
