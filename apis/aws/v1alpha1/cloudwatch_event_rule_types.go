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

type CloudwatchEventRule struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              CloudwatchEventRuleSpec   `json:"spec,omitempty"`
	Status            CloudwatchEventRuleStatus `json:"status,omitempty"`
}

type CloudwatchEventRuleSpec struct {
	// +optional
	Description string `json:"description,omitempty"`
	// +optional
	EventPattern string `json:"event_pattern,omitempty"`
	// +optional
	IsEnabled bool `json:"is_enabled,omitempty"`
	// +optional
	NamePrefix string `json:"name_prefix,omitempty"`
	// +optional
	RoleArn string `json:"role_arn,omitempty"`
	// +optional
	ScheduleExpression string `json:"schedule_expression,omitempty"`
	// +optional
	Tags map[string]string `json:"tags,omitempty"`
}

type CloudwatchEventRuleStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// CloudwatchEventRuleList is a list of CloudwatchEventRules
type CloudwatchEventRuleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of CloudwatchEventRule CRD objects
	Items []CloudwatchEventRule `json:"items,omitempty"`
}
