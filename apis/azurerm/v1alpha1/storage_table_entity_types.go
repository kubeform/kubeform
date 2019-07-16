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

type StorageTableEntity struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              StorageTableEntitySpec   `json:"spec,omitempty"`
	Status            StorageTableEntityStatus `json:"status,omitempty"`
}

type StorageTableEntitySpec struct {
	Entity             map[string]string `json:"entity"`
	PartitionKey       string            `json:"partition_key"`
	RowKey             string            `json:"row_key"`
	StorageAccountName string            `json:"storage_account_name"`
	TableName          string            `json:"table_name"`
}

type StorageTableEntityStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// StorageTableEntityList is a list of StorageTableEntitys
type StorageTableEntityList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of StorageTableEntity CRD objects
	Items []StorageTableEntity `json:"items,omitempty"`
}
