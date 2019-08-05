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

type DmsReplicationSubnetGroup struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DmsReplicationSubnetGroupSpec   `json:"spec,omitempty"`
	Status            DmsReplicationSubnetGroupStatus `json:"status,omitempty"`
}

type DmsReplicationSubnetGroupSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ReplicationSubnetGroupArn         string `json:"replicationSubnetGroupArn" tf:"replication_subnet_group_arn"`
	ReplicationSubnetGroupDescription string `json:"replicationSubnetGroupDescription" tf:"replication_subnet_group_description"`
	ReplicationSubnetGroupID          string `json:"replicationSubnetGroupID" tf:"replication_subnet_group_id"`
	// +kubebuilder:validation:UniqueItems=true
	SubnetIDS []string `json:"subnetIDS" tf:"subnet_ids"`
	// +optional
	Tags  map[string]string `json:"tags,omitempty" tf:"tags,omitempty"`
	VpcID string            `json:"vpcID" tf:"vpc_id"`
}

type DmsReplicationSubnetGroupStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *DmsReplicationSubnetGroupSpec `json:"output,omitempty"`
	// +optional
	State *apis.State `json:"state,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// DmsReplicationSubnetGroupList is a list of DmsReplicationSubnetGroups
type DmsReplicationSubnetGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of DmsReplicationSubnetGroup CRD objects
	Items []DmsReplicationSubnetGroup `json:"items,omitempty"`
}
