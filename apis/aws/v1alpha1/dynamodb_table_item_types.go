package v1alpha1

import (
	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	base "kubeform.dev/kubeform/apis/base/v1alpha1"
)

// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

type DynamodbTableItem struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DynamodbTableItemSpec   `json:"spec,omitempty"`
	Status            DynamodbTableItemStatus `json:"status,omitempty"`
}

type DynamodbTableItemSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	HashKey string `json:"hashKey" tf:"hash_key"`
	Item    string `json:"item" tf:"item"`
	// +optional
	RangeKey  string `json:"rangeKey,omitempty" tf:"range_key,omitempty"`
	TableName string `json:"tableName" tf:"table_name"`
}

type DynamodbTableItemStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *DynamodbTableItemSpec `json:"output,omitempty"`
	// +optional
	State *base.State `json:"state,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// DynamodbTableItemList is a list of DynamodbTableItems
type DynamodbTableItemList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of DynamodbTableItem CRD objects
	Items []DynamodbTableItem `json:"items,omitempty"`
}
