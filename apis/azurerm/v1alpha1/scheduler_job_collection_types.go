package v1alpha1

import (
	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	base "kubeform.dev/kubeform/apis/base/v1alpha1"
)

// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Phase",type=string,JSONPath=`.status.phase`

type SchedulerJobCollection struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SchedulerJobCollectionSpec   `json:"spec,omitempty"`
	Status            SchedulerJobCollectionStatus `json:"status,omitempty"`
}

type SchedulerJobCollectionSpecQuota struct {
	// +optional
	MaxJobCount            int    `json:"maxJobCount,omitempty" tf:"max_job_count,omitempty"`
	MaxRecurrenceFrequency string `json:"maxRecurrenceFrequency" tf:"max_recurrence_frequency"`
	// +optional
	MaxRecurrenceInterval int `json:"maxRecurrenceInterval,omitempty" tf:"max_recurrence_interval,omitempty"`
	// +optional
	// Deprecated
	MaxRetryInterval int `json:"maxRetryInterval,omitempty" tf:"max_retry_interval,omitempty"`
}

type SchedulerJobCollectionSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	Location string `json:"location" tf:"location"`
	Name     string `json:"name" tf:"name"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	Quota             []SchedulerJobCollectionSpecQuota `json:"quota,omitempty" tf:"quota,omitempty"`
	ResourceGroupName string                            `json:"resourceGroupName" tf:"resource_group_name"`
	Sku               string                            `json:"sku" tf:"sku"`
	// +optional
	State string `json:"state,omitempty" tf:"state,omitempty"`
	// +optional
	Tags map[string]string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type SchedulerJobCollectionStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *SchedulerJobCollectionSpec `json:"output,omitempty"`
	// +optional
	State *base.State `json:"state,omitempty"`
	// +optional
	Phase base.Phase `json:"phase,omitempty"`
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
