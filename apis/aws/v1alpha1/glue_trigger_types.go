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

type GlueTrigger struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              GlueTriggerSpec   `json:"spec,omitempty"`
	Status            GlueTriggerStatus `json:"status,omitempty"`
}

type GlueTriggerSpecActions struct {
	// +optional
	Arguments map[string]string `json:"arguments,omitempty"`
	JobName   string            `json:"job_name"`
	// +optional
	Timeout int `json:"timeout,omitempty"`
}

type GlueTriggerSpecPredicateConditions struct {
	JobName string `json:"job_name"`
	// +optional
	LogicalOperator string `json:"logical_operator,omitempty"`
	State           string `json:"state"`
}

type GlueTriggerSpecPredicate struct {
	// +kubebuilder:validation:MinItems=1
	Conditions []GlueTriggerSpecPredicate `json:"conditions"`
	// +optional
	Logical string `json:"logical,omitempty"`
}

type GlueTriggerSpec struct {
	// +kubebuilder:validation:MinItems=1
	Actions []GlueTriggerSpec `json:"actions"`
	// +optional
	Description string `json:"description,omitempty"`
	// +optional
	Enabled bool   `json:"enabled,omitempty"`
	Name    string `json:"name"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	Predicate *[]GlueTriggerSpec `json:"predicate,omitempty"`
	// +optional
	Schedule string `json:"schedule,omitempty"`
	Type     string `json:"type"`
}

type GlueTriggerStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
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
