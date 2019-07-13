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

type AzurermCosmosdbCassandraKeyspace struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AzurermCosmosdbCassandraKeyspaceSpec   `json:"spec,omitempty"`
	Status            AzurermCosmosdbCassandraKeyspaceStatus `json:"status,omitempty"`
}

type AzurermCosmosdbCassandraKeyspaceSpec struct {
	Name              string `json:"name"`
	ResourceGroupName string `json:"resource_group_name"`
	AccountName       string `json:"account_name"`
}



type AzurermCosmosdbCassandraKeyspaceStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AzurermCosmosdbCassandraKeyspaceList is a list of AzurermCosmosdbCassandraKeyspaces
type AzurermCosmosdbCassandraKeyspaceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AzurermCosmosdbCassandraKeyspace CRD objects
	Items []AzurermCosmosdbCassandraKeyspace `json:"items,omitempty"`
}