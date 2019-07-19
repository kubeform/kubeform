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

type MariadbDatabase struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              MariadbDatabaseSpec   `json:"spec,omitempty"`
	Status            MariadbDatabaseStatus `json:"status,omitempty"`
}

type MariadbDatabaseSpec struct {
	Charset           string                    `json:"charset" tf:"charset"`
	Collation         string                    `json:"collation" tf:"collation"`
	Name              string                    `json:"name" tf:"name"`
	ResourceGroupName string                    `json:"resourceGroupName" tf:"resource_group_name"`
	ServerName        string                    `json:"serverName" tf:"server_name"`
	ProviderRef       core.LocalObjectReference `json:"providerRef" tf:"-"`
}

type MariadbDatabaseStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	TFState *runtime.RawExtension `json:"tfState,omitempty"`
	Output  *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// MariadbDatabaseList is a list of MariadbDatabases
type MariadbDatabaseList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of MariadbDatabase CRD objects
	Items []MariadbDatabase `json:"items,omitempty"`
}
