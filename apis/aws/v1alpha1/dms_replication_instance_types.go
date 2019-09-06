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

type DmsReplicationInstance struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DmsReplicationInstanceSpec   `json:"spec,omitempty"`
	Status            DmsReplicationInstanceStatus `json:"status,omitempty"`
}

type DmsReplicationInstanceSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	AllocatedStorage int `json:"allocatedStorage,omitempty" tf:"allocated_storage,omitempty"`
	// +optional
	ApplyImmediately bool `json:"applyImmediately,omitempty" tf:"apply_immediately,omitempty"`
	// +optional
	AutoMinorVersionUpgrade bool `json:"autoMinorVersionUpgrade,omitempty" tf:"auto_minor_version_upgrade,omitempty"`
	// +optional
	AvailabilityZone string `json:"availabilityZone,omitempty" tf:"availability_zone,omitempty"`
	// +optional
	EngineVersion string `json:"engineVersion,omitempty" tf:"engine_version,omitempty"`
	// +optional
	KmsKeyArn string `json:"kmsKeyArn,omitempty" tf:"kms_key_arn,omitempty"`
	// +optional
	MultiAz bool `json:"multiAz,omitempty" tf:"multi_az,omitempty"`
	// +optional
	PreferredMaintenanceWindow string `json:"preferredMaintenanceWindow,omitempty" tf:"preferred_maintenance_window,omitempty"`
	// +optional
	PubliclyAccessible bool `json:"publiclyAccessible,omitempty" tf:"publicly_accessible,omitempty"`
	// +optional
	ReplicationInstanceArn   string `json:"replicationInstanceArn,omitempty" tf:"replication_instance_arn,omitempty"`
	ReplicationInstanceClass string `json:"replicationInstanceClass" tf:"replication_instance_class"`
	ReplicationInstanceID    string `json:"replicationInstanceID" tf:"replication_instance_id"`
	// +optional
	ReplicationInstancePrivateIPS []string `json:"replicationInstancePrivateIPS,omitempty" tf:"replication_instance_private_ips,omitempty"`
	// +optional
	ReplicationInstancePublicIPS []string `json:"replicationInstancePublicIPS,omitempty" tf:"replication_instance_public_ips,omitempty"`
	// +optional
	ReplicationSubnetGroupID string `json:"replicationSubnetGroupID,omitempty" tf:"replication_subnet_group_id,omitempty"`
	// +optional
	Tags map[string]string `json:"tags,omitempty" tf:"tags,omitempty"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	VpcSecurityGroupIDS []string `json:"vpcSecurityGroupIDS,omitempty" tf:"vpc_security_group_ids,omitempty"`
}

type DmsReplicationInstanceStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *DmsReplicationInstanceSpec `json:"output,omitempty"`
	// +optional
	State *apis.State `json:"state,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// DmsReplicationInstanceList is a list of DmsReplicationInstances
type DmsReplicationInstanceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of DmsReplicationInstance CRD objects
	Items []DmsReplicationInstance `json:"items,omitempty"`
}
