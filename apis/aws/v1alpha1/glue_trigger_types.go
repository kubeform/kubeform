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

type GlueTrigger struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              GlueTriggerSpec   `json:"spec,omitempty"`
	Status            GlueTriggerStatus `json:"status,omitempty"`
}

type GlueTriggerSpecActions struct {
	// +optional
	Arguments map[string]string `json:"arguments,omitempty" tf:"arguments,omitempty"`
	JobName   string            `json:"jobName" tf:"job_name"`
	// +optional
	Timeout int `json:"timeout,omitempty" tf:"timeout,omitempty"`
}

type GlueTriggerSpecPredicateConditions struct {
	JobName string `json:"jobName" tf:"job_name"`
	// +optional
	LogicalOperator string `json:"logicalOperator,omitempty" tf:"logical_operator,omitempty"`
	State           string `json:"state" tf:"state"`
}

type GlueTriggerSpecPredicate struct {
	// +kubebuilder:validation:MinItems=1
	Conditions []GlueTriggerSpecPredicateConditions `json:"conditions" tf:"conditions"`
	// +optional
	Logical string `json:"logical,omitempty" tf:"logical,omitempty"`
}

type GlueTriggerSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +kubebuilder:validation:MinItems=1
	Actions []GlueTriggerSpecActions `json:"actions" tf:"actions"`
	// +optional
	Description string `json:"description,omitempty" tf:"description,omitempty"`
	// +optional
	Enabled bool   `json:"enabled,omitempty" tf:"enabled,omitempty"`
	Name    string `json:"name" tf:"name"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	Predicate []GlueTriggerSpecPredicate `json:"predicate,omitempty" tf:"predicate,omitempty"`
	// +optional
	Schedule string `json:"schedule,omitempty" tf:"schedule,omitempty"`
	Type     string `json:"type" tf:"type"`
}

type GlueTriggerStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *GlueTriggerSpec `json:"output,omitempty"`
	// +optional
	State *apis.State `json:"state,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// GlueTriggerList is a list of GlueTriggers
type GlueTriggerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of GlueTrigger CRD objects
	Items []GlueTrigger `json:"items,omitempty"`
}
