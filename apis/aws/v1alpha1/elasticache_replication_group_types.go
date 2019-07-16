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

type ElasticacheReplicationGroup struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ElasticacheReplicationGroupSpec   `json:"spec,omitempty"`
	Status            ElasticacheReplicationGroupStatus `json:"status,omitempty"`
}

type ElasticacheReplicationGroupSpec struct {
	// +optional
	AtRestEncryptionEnabled bool `json:"at_rest_encryption_enabled,omitempty"`
	// +optional
	AuthToken string `json:"auth_token,omitempty"`
	// +optional
	AutoMinorVersionUpgrade bool `json:"auto_minor_version_upgrade,omitempty"`
	// +optional
	AutomaticFailoverEnabled bool `json:"automatic_failover_enabled,omitempty"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	AvailabilityZones []string `json:"availability_zones,omitempty"`
	// +optional
	Engine string `json:"engine,omitempty"`
	// +optional
	NotificationTopicArn string `json:"notification_topic_arn,omitempty"`
	// +optional
	Port                        int    `json:"port,omitempty"`
	ReplicationGroupDescription string `json:"replication_group_description"`
	ReplicationGroupId          string `json:"replication_group_id"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	SnapshotArns []string `json:"snapshot_arns,omitempty"`
	// +optional
	SnapshotName string `json:"snapshot_name,omitempty"`
	// +optional
	SnapshotRetentionLimit int `json:"snapshot_retention_limit,omitempty"`
	// +optional
	Tags map[string]string `json:"tags,omitempty"`
	// +optional
	TransitEncryptionEnabled bool `json:"transit_encryption_enabled,omitempty"`
}

type ElasticacheReplicationGroupStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// ElasticacheReplicationGroupList is a list of ElasticacheReplicationGroups
type ElasticacheReplicationGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of ElasticacheReplicationGroup CRD objects
	Items []ElasticacheReplicationGroup `json:"items,omitempty"`
}
