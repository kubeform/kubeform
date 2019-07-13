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

type AwsGlueTrigger struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsGlueTriggerSpec   `json:"spec,omitempty"`
	Status            AwsGlueTriggerStatus `json:"status,omitempty"`
}

type AwsGlueTriggerSpecPredicateConditions struct {
	JobName         string `json:"job_name"`
	LogicalOperator string `json:"logical_operator"`
	State           string `json:"state"`
}

type AwsGlueTriggerSpecPredicate struct {
	Conditions []AwsGlueTriggerSpecPredicate `json:"conditions"`
	Logical    string                        `json:"logical"`
}

type AwsGlueTriggerSpecActions struct {
	Arguments map[string]string `json:"arguments"`
	JobName   string            `json:"job_name"`
	Timeout   int               `json:"timeout"`
}

type AwsGlueTriggerSpec struct {
	Predicate   []AwsGlueTriggerSpec `json:"predicate"`
	Schedule    string               `json:"schedule"`
	Type        string               `json:"type"`
	Actions     []AwsGlueTriggerSpec `json:"actions"`
	Description string               `json:"description"`
	Enabled     bool                 `json:"enabled"`
	Name        string               `json:"name"`
}



type AwsGlueTriggerStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AwsGlueTriggerList is a list of AwsGlueTriggers
type AwsGlueTriggerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsGlueTrigger CRD objects
	Items []AwsGlueTrigger `json:"items,omitempty"`
}