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

type NeptuneCluster struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              NeptuneClusterSpec   `json:"spec,omitempty"`
	Status            NeptuneClusterStatus `json:"status,omitempty"`
}

type NeptuneClusterSpec struct {
	// +optional
	BackupRetentionPeriod int `json:"backup_retention_period,omitempty"`
	// +optional
	Engine string `json:"engine,omitempty"`
	// +optional
	FinalSnapshotIdentifier string `json:"final_snapshot_identifier,omitempty"`
	// +optional
	IamDatabaseAuthenticationEnabled bool `json:"iam_database_authentication_enabled,omitempty"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	IamRoles []string `json:"iam_roles,omitempty"`
	// +optional
	NeptuneClusterParameterGroupName string `json:"neptune_cluster_parameter_group_name,omitempty"`
	// +optional
	Port int `json:"port,omitempty"`
	// +optional
	ReplicationSourceIdentifier string `json:"replication_source_identifier,omitempty"`
	// +optional
	SkipFinalSnapshot bool `json:"skip_final_snapshot,omitempty"`
	// +optional
	SnapshotIdentifier string `json:"snapshot_identifier,omitempty"`
	// +optional
	StorageEncrypted bool `json:"storage_encrypted,omitempty"`
	// +optional
	Tags map[string]string `json:"tags,omitempty"`
}

type NeptuneClusterStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// NeptuneClusterList is a list of NeptuneClusters
type NeptuneClusterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of NeptuneCluster CRD objects
	Items []NeptuneCluster `json:"items,omitempty"`
}
