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

type SearchService struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SearchServiceSpec   `json:"spec,omitempty"`
	Status            SearchServiceStatus `json:"status,omitempty"`
}

type SearchServiceSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	Location string `json:"location" tf:"location"`
	Name     string `json:"name" tf:"name"`
	// +optional
	PartitionCount int `json:"partitionCount,omitempty" tf:"partition_count,omitempty"`
	// +optional
	PrimaryKey string `json:"primaryKey,omitempty" tf:"primary_key,omitempty"`
	// +optional
	ReplicaCount      int    `json:"replicaCount,omitempty" tf:"replica_count,omitempty"`
	ResourceGroupName string `json:"resourceGroupName" tf:"resource_group_name"`
	// +optional
	SecondaryKey string `json:"secondaryKey,omitempty" tf:"secondary_key,omitempty"`
	Sku          string `json:"sku" tf:"sku"`
	// +optional
	Tags map[string]string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type SearchServiceStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *SearchServiceSpec `json:"output,omitempty"`
	// +optional
	State *apis.State `json:"state,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// SearchServiceList is a list of SearchServices
type SearchServiceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of SearchService CRD objects
	Items []SearchService `json:"items,omitempty"`
}