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

type DaxCluster struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DaxClusterSpec   `json:"spec,omitempty"`
	Status            DaxClusterStatus `json:"status,omitempty"`
}

type DaxClusterSpecNodes struct {
	Address          string `json:"address" tf:"address"`
	AvailabilityZone string `json:"availabilityZone" tf:"availability_zone"`
	ID               string `json:"ID" tf:"id"`
	Port             int    `json:"port" tf:"port"`
}

type DaxClusterSpecServerSideEncryption struct {
	// +optional
	Enabled bool `json:"enabled,omitempty" tf:"enabled,omitempty"`
}

type DaxClusterSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	Arn string `json:"arn" tf:"arn"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	AvailabilityZones     []string `json:"availabilityZones,omitempty" tf:"availability_zones,omitempty"`
	ClusterAddress        string   `json:"clusterAddress" tf:"cluster_address"`
	ClusterName           string   `json:"clusterName" tf:"cluster_name"`
	ConfigurationEndpoint string   `json:"configurationEndpoint" tf:"configuration_endpoint"`
	// +optional
	Description string `json:"description,omitempty" tf:"description,omitempty"`
	IamRoleArn  string `json:"iamRoleArn" tf:"iam_role_arn"`
	// +optional
	MaintenanceWindow string                `json:"maintenanceWindow,omitempty" tf:"maintenance_window,omitempty"`
	NodeType          string                `json:"nodeType" tf:"node_type"`
	Nodes             []DaxClusterSpecNodes `json:"nodes" tf:"nodes"`
	// +optional
	NotificationTopicArn string `json:"notificationTopicArn,omitempty" tf:"notification_topic_arn,omitempty"`
	// +optional
	ParameterGroupName string `json:"parameterGroupName,omitempty" tf:"parameter_group_name,omitempty"`
	Port               int    `json:"port" tf:"port"`
	ReplicationFactor  int    `json:"replicationFactor" tf:"replication_factor"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	SecurityGroupIDS []string `json:"securityGroupIDS,omitempty" tf:"security_group_ids,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	ServerSideEncryption []DaxClusterSpecServerSideEncryption `json:"serverSideEncryption,omitempty" tf:"server_side_encryption,omitempty"`
	// +optional
	SubnetGroupName string `json:"subnetGroupName,omitempty" tf:"subnet_group_name,omitempty"`
	// +optional
	Tags map[string]string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type DaxClusterStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *DaxClusterSpec `json:"output,omitempty"`
	// +optional
	State *apis.State `json:"state,omitempty"`
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
