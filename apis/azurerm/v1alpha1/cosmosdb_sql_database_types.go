package v1alpha1

import (
	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

type CosmosdbSQLDatabase struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              CosmosdbSQLDatabaseSpec   `json:"spec,omitempty"`
	Status            CosmosdbSQLDatabaseStatus `json:"status,omitempty"`
}

type CosmosdbSQLDatabaseSpec struct {
	ProviderRef       core.LocalObjectReference `json:"providerRef" tf:"-"`
	AccountName       string                    `json:"accountName" tf:"account_name"`
	Name              string                    `json:"name" tf:"name"`
	ResourceGroupName string                    `json:"resourceGroupName" tf:"resource_group_name"`
}

type CosmosdbSQLDatabaseStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	TFState *runtime.RawExtension `json:"tfState,omitempty"`
	Output  *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// CosmosdbSQLDatabaseList is a list of CosmosdbSQLDatabases
type CosmosdbSQLDatabaseList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of CosmosdbSQLDatabase CRD objects
	Items []CosmosdbSQLDatabase `json:"items,omitempty"`
}
