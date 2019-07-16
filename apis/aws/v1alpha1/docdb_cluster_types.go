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

type DocdbCluster struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DocdbClusterSpec   `json:"spec,omitempty"`
	Status            DocdbClusterStatus `json:"status,omitempty"`
}

type DocdbClusterSpec struct {
	// +optional
	BackupRetentionPeriod int `json:"backup_retention_period,omitempty"`
	// +optional
	EnabledCloudwatchLogsExports []string `json:"enabled_cloudwatch_logs_exports,omitempty"`
	// +optional
	Engine string `json:"engine,omitempty"`
	// +optional
	FinalSnapshotIdentifier string `json:"final_snapshot_identifier,omitempty"`
	// +optional
	MasterPassword string `json:"master_password,omitempty"`
	// +optional
	Port int `json:"port,omitempty"`
	// +optional
	SkipFinalSnapshot bool `json:"skip_final_snapshot,omitempty"`
	// +optional
	SnapshotIdentifier string `json:"snapshot_identifier,omitempty"`
	// +optional
	StorageEncrypted bool `json:"storage_encrypted,omitempty"`
	// +optional
	Tags map[string]string `json:"tags,omitempty"`
}

type DocdbClusterStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// DocdbClusterList is a list of DocdbClusters
type DocdbClusterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of DocdbCluster CRD objects
	Items []DocdbCluster `json:"items,omitempty"`
}
