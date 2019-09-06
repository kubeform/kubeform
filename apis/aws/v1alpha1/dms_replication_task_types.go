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

type DmsReplicationTask struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DmsReplicationTaskSpec   `json:"spec,omitempty"`
	Status            DmsReplicationTaskStatus `json:"status,omitempty"`
}

type DmsReplicationTaskSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	CdcStartTime           string `json:"cdcStartTime,omitempty" tf:"cdc_start_time,omitempty"`
	MigrationType          string `json:"migrationType" tf:"migration_type"`
	ReplicationInstanceArn string `json:"replicationInstanceArn" tf:"replication_instance_arn"`
	// +optional
	ReplicationTaskArn string `json:"replicationTaskArn,omitempty" tf:"replication_task_arn,omitempty"`
	ReplicationTaskID  string `json:"replicationTaskID" tf:"replication_task_id"`
	// +optional
	ReplicationTaskSettings string `json:"replicationTaskSettings,omitempty" tf:"replication_task_settings,omitempty"`
	SourceEndpointArn       string `json:"sourceEndpointArn" tf:"source_endpoint_arn"`
	TableMappings           string `json:"tableMappings" tf:"table_mappings"`
	// +optional
	Tags              map[string]string `json:"tags,omitempty" tf:"tags,omitempty"`
	TargetEndpointArn string            `json:"targetEndpointArn" tf:"target_endpoint_arn"`
}

type DmsReplicationTaskStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *DmsReplicationTaskSpec `json:"output,omitempty"`
	// +optional
	State *apis.State `json:"state,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// DmsReplicationTaskList is a list of DmsReplicationTasks
type DmsReplicationTaskList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of DmsReplicationTask CRD objects
	Items []DmsReplicationTask `json:"items,omitempty"`
}
