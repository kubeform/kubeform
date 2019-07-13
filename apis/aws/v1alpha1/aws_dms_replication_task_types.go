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

type AwsDmsReplicationTask struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsDmsReplicationTaskSpec   `json:"spec,omitempty"`
	Status            AwsDmsReplicationTaskStatus `json:"status,omitempty"`
}

type AwsDmsReplicationTaskSpec struct {
	ReplicationTaskSettings string            `json:"replication_task_settings"`
	TableMappings           string            `json:"table_mappings"`
	TargetEndpointArn       string            `json:"target_endpoint_arn"`
	CdcStartTime            string            `json:"cdc_start_time"`
	MigrationType           string            `json:"migration_type"`
	ReplicationInstanceArn  string            `json:"replication_instance_arn"`
	ReplicationTaskArn      string            `json:"replication_task_arn"`
	ReplicationTaskId       string            `json:"replication_task_id"`
	SourceEndpointArn       string            `json:"source_endpoint_arn"`
	Tags                    map[string]string `json:"tags"`
}



type AwsDmsReplicationTaskStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AwsDmsReplicationTaskList is a list of AwsDmsReplicationTasks
type AwsDmsReplicationTaskList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsDmsReplicationTask CRD objects
	Items []AwsDmsReplicationTask `json:"items,omitempty"`
}