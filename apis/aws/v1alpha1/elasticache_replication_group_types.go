/*
Copyright The Kubeform Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by Kubeform. DO NOT EDIT.

package v1alpha1

import (
	base "kubeform.dev/kubeform/apis/base/v1alpha1"

	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Phase",type=string,JSONPath=`.status.phase`

type ElasticacheReplicationGroup struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ElasticacheReplicationGroupSpec   `json:"spec,omitempty"`
	Status            ElasticacheReplicationGroupStatus `json:"status,omitempty"`
}

type ElasticacheReplicationGroupSpecClusterMode struct {
	NumNodeGroups        int64 `json:"numNodeGroups" tf:"num_node_groups"`
	ReplicasPerNodeGroup int64 `json:"replicasPerNodeGroup" tf:"replicas_per_node_group"`
}

type ElasticacheReplicationGroupSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	RemoteBackend *base.Backend `json:"remoteBackend,omitempty" tf:"-"`

	SecretRef *core.LocalObjectReference `json:"secretRef,omitempty" tf:"-"`

	// +optional
	ApplyImmediately bool `json:"applyImmediately,omitempty" tf:"apply_immediately,omitempty"`
	// +optional
	AtRestEncryptionEnabled bool `json:"atRestEncryptionEnabled,omitempty" tf:"at_rest_encryption_enabled,omitempty"`
	// +optional
	AuthToken string `json:"-" sensitive:"true" tf:"auth_token,omitempty"`
	// +optional
	AutoMinorVersionUpgrade bool `json:"autoMinorVersionUpgrade,omitempty" tf:"auto_minor_version_upgrade,omitempty"`
	// +optional
	AutomaticFailoverEnabled bool `json:"automaticFailoverEnabled,omitempty" tf:"automatic_failover_enabled,omitempty"`
	// +optional
	AvailabilityZones []string `json:"availabilityZones,omitempty" tf:"availability_zones,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	ClusterMode []ElasticacheReplicationGroupSpecClusterMode `json:"clusterMode,omitempty" tf:"cluster_mode,omitempty"`
	// +optional
	ConfigurationEndpointAddress string `json:"configurationEndpointAddress,omitempty" tf:"configuration_endpoint_address,omitempty"`
	// +optional
	Engine string `json:"engine,omitempty" tf:"engine,omitempty"`
	// +optional
	EngineVersion string `json:"engineVersion,omitempty" tf:"engine_version,omitempty"`
	// +optional
	MaintenanceWindow string `json:"maintenanceWindow,omitempty" tf:"maintenance_window,omitempty"`
	// +optional
	MemberClusters []string `json:"memberClusters,omitempty" tf:"member_clusters,omitempty"`
	// +optional
	NodeType string `json:"nodeType,omitempty" tf:"node_type,omitempty"`
	// +optional
	NotificationTopicArn string `json:"notificationTopicArn,omitempty" tf:"notification_topic_arn,omitempty"`
	// +optional
	NumberCacheClusters int64 `json:"numberCacheClusters,omitempty" tf:"number_cache_clusters,omitempty"`
	// +optional
	ParameterGroupName string `json:"parameterGroupName,omitempty" tf:"parameter_group_name,omitempty"`
	// +optional
	Port int64 `json:"port,omitempty" tf:"port,omitempty"`
	// +optional
	PrimaryEndpointAddress      string `json:"primaryEndpointAddress,omitempty" tf:"primary_endpoint_address,omitempty"`
	ReplicationGroupDescription string `json:"replicationGroupDescription" tf:"replication_group_description"`
	ReplicationGroupID          string `json:"replicationGroupID" tf:"replication_group_id"`
	// +optional
	SecurityGroupIDS []string `json:"securityGroupIDS,omitempty" tf:"security_group_ids,omitempty"`
	// +optional
	SecurityGroupNames []string `json:"securityGroupNames,omitempty" tf:"security_group_names,omitempty"`
	// +optional
	SnapshotArns []string `json:"snapshotArns,omitempty" tf:"snapshot_arns,omitempty"`
	// +optional
	SnapshotName string `json:"snapshotName,omitempty" tf:"snapshot_name,omitempty"`
	// +optional
	SnapshotRetentionLimit int64 `json:"snapshotRetentionLimit,omitempty" tf:"snapshot_retention_limit,omitempty"`
	// +optional
	SnapshotWindow string `json:"snapshotWindow,omitempty" tf:"snapshot_window,omitempty"`
	// +optional
	SubnetGroupName string `json:"subnetGroupName,omitempty" tf:"subnet_group_name,omitempty"`
	// +optional
	Tags map[string]string `json:"tags,omitempty" tf:"tags,omitempty"`
	// +optional
	TransitEncryptionEnabled bool `json:"transitEncryptionEnabled,omitempty" tf:"transit_encryption_enabled,omitempty"`
}

type ElasticacheReplicationGroupStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *ElasticacheReplicationGroupSpec `json:"output,omitempty"`
	// +optional
	State *base.State `json:"state,omitempty"`
	// +optional
	Phase base.Phase `json:"phase,omitempty"`
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
