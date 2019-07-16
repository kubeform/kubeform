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

type DmsReplicationSubnetGroup struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DmsReplicationSubnetGroupSpec   `json:"spec,omitempty"`
	Status            DmsReplicationSubnetGroupStatus `json:"status,omitempty"`
}

type DmsReplicationSubnetGroupSpec struct {
	ReplicationSubnetGroupDescription string `json:"replication_subnet_group_description"`
	ReplicationSubnetGroupId          string `json:"replication_subnet_group_id"`
	// +kubebuilder:validation:UniqueItems=true
	SubnetIds []string `json:"subnet_ids"`
	// +optional
	Tags map[string]string `json:"tags,omitempty"`
}

type DmsReplicationSubnetGroupStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
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
