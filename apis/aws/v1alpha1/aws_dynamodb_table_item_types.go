package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsDynamodbTableItem struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsDynamodbTableItemSpec   `json:"spec,omitempty"`
	Status            AwsDynamodbTableItemStatus `json:"status,omitempty"`
}

type AwsDynamodbTableItemSpec struct {
	TableName string `json:"table_name"`
	HashKey   string `json:"hash_key"`
	RangeKey  string `json:"range_key"`
	Item      string `json:"item"`
}

type AwsDynamodbTableItemStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsDynamodbTableItemList is a list of AwsDynamodbTableItems
type AwsDynamodbTableItemList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsDynamodbTableItem CRD objects
	Items []AwsDynamodbTableItem `json:"items,omitempty"`
}
