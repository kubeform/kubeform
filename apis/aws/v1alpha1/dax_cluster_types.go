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

type DaxCluster struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DaxClusterSpec   `json:"spec,omitempty"`
	Status            DaxClusterStatus `json:"status,omitempty"`
}

type DaxClusterSpecServerSideEncryption struct {
	// +optional
	Enabled bool `json:"enabled,omitempty"`
}

type DaxClusterSpec struct {
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	AvailabilityZones []string `json:"availability_zones,omitempty"`
	ClusterName       string   `json:"cluster_name"`
	// +optional
	Description string `json:"description,omitempty"`
	IamRoleArn  string `json:"iam_role_arn"`
	NodeType    string `json:"node_type"`
	// +optional
	NotificationTopicArn string `json:"notification_topic_arn,omitempty"`
	ReplicationFactor    int    `json:"replication_factor"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	ServerSideEncryption *[]DaxClusterSpec `json:"server_side_encryption,omitempty"`
	// +optional
	Tags map[string]string `json:"tags,omitempty"`
}

type DaxClusterStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// DaxClusterList is a list of DaxClusters
type DaxClusterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of DaxCluster CRD objects
	Items []DaxCluster `json:"items,omitempty"`
}
