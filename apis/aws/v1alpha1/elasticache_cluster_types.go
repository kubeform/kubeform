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

type ElasticacheCluster struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ElasticacheClusterSpec   `json:"spec,omitempty"`
	Status            ElasticacheClusterStatus `json:"status,omitempty"`
}

type ElasticacheClusterSpecCacheNodes struct {
	Address          string `json:"address" tf:"address"`
	AvailabilityZone string `json:"availabilityZone" tf:"availability_zone"`
	ID               string `json:"ID" tf:"id"`
	Port             int    `json:"port" tf:"port"`
}

type ElasticacheClusterSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	// +optional
	ApplyImmediately bool `json:"applyImmediately,omitempty" tf:"apply_immediately,omitempty"`
	// +optional
	AvailabilityZone string `json:"availabilityZone,omitempty" tf:"availability_zone,omitempty"`
	// +optional
	AzMode                string                             `json:"azMode,omitempty" tf:"az_mode,omitempty"`
	CacheNodes            []ElasticacheClusterSpecCacheNodes `json:"cacheNodes" tf:"cache_nodes"`
	ClusterAddress        string                             `json:"clusterAddress" tf:"cluster_address"`
	ClusterID             string                             `json:"clusterID" tf:"cluster_id"`
	ConfigurationEndpoint string                             `json:"configurationEndpoint" tf:"configuration_endpoint"`
	// +optional
	Engine string `json:"engine,omitempty" tf:"engine,omitempty"`
	// +optional
	EngineVersion string `json:"engineVersion,omitempty" tf:"engine_version,omitempty"`
	// +optional
	MaintenanceWindow string `json:"maintenanceWindow,omitempty" tf:"maintenance_window,omitempty"`
	// +optional
	NodeType string `json:"nodeType,omitempty" tf:"node_type,omitempty"`
	// +optional
	NotificationTopicArn string `json:"notificationTopicArn,omitempty" tf:"notification_topic_arn,omitempty"`
	// +optional
	NumCacheNodes int `json:"numCacheNodes,omitempty" tf:"num_cache_nodes,omitempty"`
	// +optional
	ParameterGroupName string `json:"parameterGroupName,omitempty" tf:"parameter_group_name,omitempty"`
	// +optional
	Port int `json:"port,omitempty" tf:"port,omitempty"`
	// +optional
	PreferredAvailabilityZones []string `json:"preferredAvailabilityZones,omitempty" tf:"preferred_availability_zones,omitempty"`
	// +optional
	ReplicationGroupID string `json:"replicationGroupID,omitempty" tf:"replication_group_id,omitempty"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	SecurityGroupIDS []string `json:"securityGroupIDS,omitempty" tf:"security_group_ids,omitempty"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	SecurityGroupNames []string `json:"securityGroupNames,omitempty" tf:"security_group_names,omitempty"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	SnapshotArns []string `json:"snapshotArns,omitempty" tf:"snapshot_arns,omitempty"`
	// +optional
	SnapshotName string `json:"snapshotName,omitempty" tf:"snapshot_name,omitempty"`
	// +optional
	SnapshotRetentionLimit int `json:"snapshotRetentionLimit,omitempty" tf:"snapshot_retention_limit,omitempty"`
	// +optional
	SnapshotWindow string `json:"snapshotWindow,omitempty" tf:"snapshot_window,omitempty"`
	// +optional
	SubnetGroupName string `json:"subnetGroupName,omitempty" tf:"subnet_group_name,omitempty"`
	// +optional
	Tags map[string]string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type ElasticacheClusterStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *ElasticacheClusterSpec `json:"output,omitempty"`
	// +optional
	State *apis.State `json:"state,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// ElasticacheClusterList is a list of ElasticacheClusters
type ElasticacheClusterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of ElasticacheCluster CRD objects
	Items []ElasticacheCluster `json:"items,omitempty"`
}
