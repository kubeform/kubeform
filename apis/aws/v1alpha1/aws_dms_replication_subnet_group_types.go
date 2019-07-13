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

type AwsDmsReplicationSubnetGroup struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsDmsReplicationSubnetGroupSpec   `json:"spec,omitempty"`
	Status            AwsDmsReplicationSubnetGroupStatus `json:"status,omitempty"`
}

type AwsDmsReplicationSubnetGroupSpec struct {
	ReplicationSubnetGroupArn         string            `json:"replication_subnet_group_arn"`
	ReplicationSubnetGroupDescription string            `json:"replication_subnet_group_description"`
	ReplicationSubnetGroupId          string            `json:"replication_subnet_group_id"`
	SubnetIds                         []string          `json:"subnet_ids"`
	Tags                              map[string]string `json:"tags"`
	VpcId                             string            `json:"vpc_id"`
}



type AwsDmsReplicationSubnetGroupStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AwsDmsReplicationSubnetGroupList is a list of AwsDmsReplicationSubnetGroups
type AwsDmsReplicationSubnetGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsDmsReplicationSubnetGroup CRD objects
	Items []AwsDmsReplicationSubnetGroup `json:"items,omitempty"`
}