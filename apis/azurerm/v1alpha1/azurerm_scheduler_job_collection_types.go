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

type AzurermSchedulerJobCollection struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AzurermSchedulerJobCollectionSpec   `json:"spec,omitempty"`
	Status            AzurermSchedulerJobCollectionStatus `json:"status,omitempty"`
}

type AzurermSchedulerJobCollectionSpecQuota struct {
	MaxJobCount            int    `json:"max_job_count"`
	MaxRecurrenceFrequency string `json:"max_recurrence_frequency"`
	MaxRetryInterval       int    `json:"max_retry_interval"`
	MaxRecurrenceInterval  int    `json:"max_recurrence_interval"`
}

type AzurermSchedulerJobCollectionSpec struct {
	Tags              map[string]string                   `json:"tags"`
	Sku               string                              `json:"sku"`
	State             string                              `json:"state"`
	Quota             []AzurermSchedulerJobCollectionSpec `json:"quota"`
	Name              string                              `json:"name"`
	Location          string                              `json:"location"`
	ResourceGroupName string                              `json:"resource_group_name"`
}



type AzurermSchedulerJobCollectionStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AzurermSchedulerJobCollectionList is a list of AzurermSchedulerJobCollections
type AzurermSchedulerJobCollectionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AzurermSchedulerJobCollection CRD objects
	Items []AzurermSchedulerJobCollection `json:"items,omitempty"`
}