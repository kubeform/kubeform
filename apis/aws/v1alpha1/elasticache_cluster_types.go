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

type ElasticacheCluster struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ElasticacheClusterSpec   `json:"spec,omitempty"`
	Status            ElasticacheClusterStatus `json:"status,omitempty"`
}

type ElasticacheClusterSpec struct {
	ClusterId string `json:"cluster_id"`
	// +optional
	NotificationTopicArn string `json:"notification_topic_arn,omitempty"`
	// +optional
	Port int `json:"port,omitempty"`
	// +optional
	PreferredAvailabilityZones []string `json:"preferred_availability_zones,omitempty"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	SnapshotArns []string `json:"snapshot_arns,omitempty"`
	// +optional
	SnapshotName string `json:"snapshot_name,omitempty"`
	// +optional
	SnapshotRetentionLimit int `json:"snapshot_retention_limit,omitempty"`
	// +optional
	Tags map[string]string `json:"tags,omitempty"`
}

type ElasticacheClusterStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
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
