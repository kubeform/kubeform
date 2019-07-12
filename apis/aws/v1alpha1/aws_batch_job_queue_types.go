package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsBatchJobQueue struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsBatchJobQueueSpec   `json:"spec,omitempty"`
	Status            AwsBatchJobQueueStatus `json:"status,omitempty"`
}

type AwsBatchJobQueueSpec struct {
	ComputeEnvironments []string `json:"compute_environments"`
	Name                string   `json:"name"`
	Priority            int      `json:"priority"`
	State               string   `json:"state"`
	Arn                 string   `json:"arn"`
}



type AwsBatchJobQueueStatus struct {
    Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsBatchJobQueueList is a list of AwsBatchJobQueues
type AwsBatchJobQueueList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsBatchJobQueue CRD objects
	Items []AwsBatchJobQueue `json:"items,omitempty"`
}