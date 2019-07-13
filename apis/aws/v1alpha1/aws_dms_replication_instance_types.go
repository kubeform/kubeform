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

type AwsDmsReplicationInstance struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsDmsReplicationInstanceSpec   `json:"spec,omitempty"`
	Status            AwsDmsReplicationInstanceStatus `json:"status,omitempty"`
}

type AwsDmsReplicationInstanceSpec struct {
	MultiAz                       bool              `json:"multi_az"`
	PreferredMaintenanceWindow    string            `json:"preferred_maintenance_window"`
	PubliclyAccessible            bool              `json:"publicly_accessible"`
	ReplicationInstanceClass      string            `json:"replication_instance_class"`
	VpcSecurityGroupIds           []string          `json:"vpc_security_group_ids"`
	ApplyImmediately              bool              `json:"apply_immediately"`
	AutoMinorVersionUpgrade       bool              `json:"auto_minor_version_upgrade"`
	EngineVersion                 string            `json:"engine_version"`
	KmsKeyArn                     string            `json:"kms_key_arn"`
	ReplicationInstancePrivateIps []string          `json:"replication_instance_private_ips"`
	ReplicationInstancePublicIps  []string          `json:"replication_instance_public_ips"`
	AvailabilityZone              string            `json:"availability_zone"`
	AllocatedStorage              int               `json:"allocated_storage"`
	ReplicationInstanceArn        string            `json:"replication_instance_arn"`
	ReplicationInstanceId         string            `json:"replication_instance_id"`
	ReplicationSubnetGroupId      string            `json:"replication_subnet_group_id"`
	Tags                          map[string]string `json:"tags"`
}



type AwsDmsReplicationInstanceStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AwsDmsReplicationInstanceList is a list of AwsDmsReplicationInstances
type AwsDmsReplicationInstanceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsDmsReplicationInstance CRD objects
	Items []AwsDmsReplicationInstance `json:"items,omitempty"`
}