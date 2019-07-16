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

type DmsReplicationInstance struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DmsReplicationInstanceSpec   `json:"spec,omitempty"`
	Status            DmsReplicationInstanceStatus `json:"status,omitempty"`
}

type DmsReplicationInstanceSpec struct {
	// +optional
	ApplyImmediately         bool   `json:"apply_immediately,omitempty"`
	ReplicationInstanceClass string `json:"replication_instance_class"`
	ReplicationInstanceId    string `json:"replication_instance_id"`
	// +optional
	Tags map[string]string `json:"tags,omitempty"`
}

type DmsReplicationInstanceStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
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
