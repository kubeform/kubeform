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

type CloudwatchLogSubscriptionFilter struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              CloudwatchLogSubscriptionFilterSpec   `json:"spec,omitempty"`
	Status            CloudwatchLogSubscriptionFilterStatus `json:"status,omitempty"`
}

type CloudwatchLogSubscriptionFilterSpec struct {
	ProviderRef    core.LocalObjectReference `json:"providerRef" tf:"-"`
	DestinationArn string                    `json:"destinationArn" tf:"destination_arn"`
	// +optional
	Distribution  string `json:"distribution,omitempty" tf:"distribution,omitempty"`
	FilterPattern string `json:"filterPattern" tf:"filter_pattern"`
	LogGroupName  string `json:"logGroupName" tf:"log_group_name"`
	Name          string `json:"name" tf:"name"`
	// +optional
	RoleArn string `json:"roleArn,omitempty" tf:"role_arn,omitempty"`
}

type CloudwatchLogSubscriptionFilterStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	TFState *runtime.RawExtension `json:"tfState,omitempty"`
	Output  *runtime.RawExtension `json:"output,omitempty"`
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
