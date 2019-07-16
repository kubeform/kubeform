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

type DynamodbTableItem struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DynamodbTableItemSpec   `json:"spec,omitempty"`
	Status            DynamodbTableItemStatus `json:"status,omitempty"`
}

type DynamodbTableItemSpec struct {
	HashKey string `json:"hash_key"`
	Item    string `json:"item"`
	// +optional
	RangeKey  string `json:"range_key,omitempty"`
	TableName string `json:"table_name"`
}

type DynamodbTableItemStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
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
