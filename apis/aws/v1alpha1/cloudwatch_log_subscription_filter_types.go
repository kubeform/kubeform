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

type CloudwatchLogSubscriptionFilter struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              CloudwatchLogSubscriptionFilterSpec   `json:"spec,omitempty"`
	Status            CloudwatchLogSubscriptionFilterStatus `json:"status,omitempty"`
}

type CloudwatchLogSubscriptionFilterSpec struct {
	DestinationArn string `json:"destination_arn"`
	// +optional
	Distribution  string `json:"distribution,omitempty"`
	FilterPattern string `json:"filter_pattern"`
	LogGroupName  string `json:"log_group_name"`
	Name          string `json:"name"`
}

type CloudwatchLogSubscriptionFilterStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// CloudwatchLogSubscriptionFilterList is a list of CloudwatchLogSubscriptionFilters
type CloudwatchLogSubscriptionFilterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of CloudwatchLogSubscriptionFilter CRD objects
	Items []CloudwatchLogSubscriptionFilter `json:"items,omitempty"`
}
