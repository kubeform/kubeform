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

type RedshiftCluster struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              RedshiftClusterSpec   `json:"spec,omitempty"`
	Status            RedshiftClusterStatus `json:"status,omitempty"`
}

type RedshiftClusterSpecLogging struct {
	Enable bool `json:"enable"`
}

type RedshiftClusterSpecSnapshotCopy struct {
	DestinationRegion string `json:"destination_region"`
	// +optional
	GrantName string `json:"grant_name,omitempty"`
	// +optional
	RetentionPeriod int `json:"retention_period,omitempty"`
}

type RedshiftClusterSpec struct {
	// +optional
	AllowVersionUpgrade bool `json:"allow_version_upgrade,omitempty"`
	// +optional
	AutomatedSnapshotRetentionPeriod int    `json:"automated_snapshot_retention_period,omitempty"`
	ClusterIdentifier                string `json:"cluster_identifier"`
	// +optional
	ClusterVersion string `json:"cluster_version,omitempty"`
	// +optional
	ElasticIp string `json:"elastic_ip,omitempty"`
	// +optional
	Encrypted bool `json:"encrypted,omitempty"`
	// +optional
	FinalSnapshotIdentifier string `json:"final_snapshot_identifier,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	Logging *[]RedshiftClusterSpec `json:"logging,omitempty"`
	// +optional
	MasterPassword string `json:"master_password,omitempty"`
	// +optional
	MasterUsername string `json:"master_username,omitempty"`
	NodeType       string `json:"node_type"`
	// +optional
	NumberOfNodes int `json:"number_of_nodes,omitempty"`
	// +optional
	OwnerAccount string `json:"owner_account,omitempty"`
	// +optional
	Port int `json:"port,omitempty"`
	// +optional
	PubliclyAccessible bool `json:"publicly_accessible,omitempty"`
	// +optional
	SkipFinalSnapshot bool `json:"skip_final_snapshot,omitempty"`
	// +optional
	SnapshotClusterIdentifier string `json:"snapshot_cluster_identifier,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	SnapshotCopy *[]RedshiftClusterSpec `json:"snapshot_copy,omitempty"`
	// +optional
	SnapshotIdentifier string `json:"snapshot_identifier,omitempty"`
	// +optional
	Tags map[string]string `json:"tags,omitempty"`
}

type RedshiftClusterStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// RedshiftClusterList is a list of RedshiftClusters
type RedshiftClusterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of RedshiftCluster CRD objects
	Items []RedshiftCluster `json:"items,omitempty"`
}
