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

type LogicAppTriggerCustom struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              LogicAppTriggerCustomSpec   `json:"spec,omitempty"`
	Status            LogicAppTriggerCustomStatus `json:"status,omitempty"`
}

type LogicAppTriggerCustomSpec struct {
	Body        string                    `json:"body" tf:"body"`
	LogicAppID  string                    `json:"logicAppID" tf:"logic_app_id"`
	Name        string                    `json:"name" tf:"name"`
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`
}

type LogicAppTriggerCustomStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	TFState *runtime.RawExtension `json:"tfState,omitempty"`
	Output  *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// LogicAppTriggerCustomList is a list of LogicAppTriggerCustoms
type LogicAppTriggerCustomList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of LogicAppTriggerCustom CRD objects
	Items []LogicAppTriggerCustom `json:"items,omitempty"`
}
