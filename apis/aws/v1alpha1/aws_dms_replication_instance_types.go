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
	ReplicationInstancePrivateIps []string          `json:"replication_instance_private_ips"`
	AvailabilityZone              string            `json:"availability_zone"`
	ReplicationSubnetGroupId      string            `json:"replication_subnet_group_id"`
	VpcSecurityGroupIds           []string          `json:"vpc_security_group_ids"`
	ReplicationInstanceId         string            `json:"replication_instance_id"`
	ReplicationInstancePublicIps  []string          `json:"replication_instance_public_ips"`
	ApplyImmediately              bool              `json:"apply_immediately"`
	AutoMinorVersionUpgrade       bool              `json:"auto_minor_version_upgrade"`
	EngineVersion                 string            `json:"engine_version"`
	PreferredMaintenanceWindow    string            `json:"preferred_maintenance_window"`
	PubliclyAccessible            bool              `json:"publicly_accessible"`
	ReplicationInstanceClass      string            `json:"replication_instance_class"`
	Tags                          map[string]string `json:"tags"`
	AllocatedStorage              int               `json:"allocated_storage"`
	KmsKeyArn                     string            `json:"kms_key_arn"`
	MultiAz                       bool              `json:"multi_az"`
	ReplicationInstanceArn        string            `json:"replication_instance_arn"`
}

type AwsDmsReplicationInstanceStatus struct {
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
