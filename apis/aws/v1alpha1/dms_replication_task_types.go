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

type DmsReplicationTask struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DmsReplicationTaskSpec   `json:"spec,omitempty"`
	Status            DmsReplicationTaskStatus `json:"status,omitempty"`
}

type DmsReplicationTaskSpec struct {
	// +optional
	CdcStartTime           string `json:"cdc_start_time,omitempty"`
	MigrationType          string `json:"migration_type"`
	ReplicationInstanceArn string `json:"replication_instance_arn"`
	ReplicationTaskId      string `json:"replication_task_id"`
	// +optional
	ReplicationTaskSettings string `json:"replication_task_settings,omitempty"`
	SourceEndpointArn       string `json:"source_endpoint_arn"`
	TableMappings           string `json:"table_mappings"`
	// +optional
	Tags              map[string]string `json:"tags,omitempty"`
	TargetEndpointArn string            `json:"target_endpoint_arn"`
}

type DmsReplicationTaskStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
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
