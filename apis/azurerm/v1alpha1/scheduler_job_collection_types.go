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

type SchedulerJobCollection struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SchedulerJobCollectionSpec   `json:"spec,omitempty"`
	Status            SchedulerJobCollectionStatus `json:"status,omitempty"`
}

type SchedulerJobCollectionSpecQuota struct {
	// +optional
	MaxJobCount            int    `json:"max_job_count,omitempty"`
	MaxRecurrenceFrequency string `json:"max_recurrence_frequency"`
	// +optional
	MaxRecurrenceInterval int `json:"max_recurrence_interval,omitempty"`
}

type SchedulerJobCollectionSpec struct {
	Location string `json:"location"`
	Name     string `json:"name"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	Quota             *[]SchedulerJobCollectionSpec `json:"quota,omitempty"`
	ResourceGroupName string                        `json:"resource_group_name"`
	Sku               string                        `json:"sku"`
	// +optional
	State string `json:"state,omitempty"`
}

type SchedulerJobCollectionStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// SchedulerJobCollectionList is a list of SchedulerJobCollections
type SchedulerJobCollectionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of SchedulerJobCollection CRD objects
	Items []SchedulerJobCollection `json:"items,omitempty"`
}
