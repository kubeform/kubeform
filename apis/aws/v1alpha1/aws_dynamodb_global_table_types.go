package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsDynamodbGlobalTable struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsDynamodbGlobalTableSpec   `json:"spec,omitempty"`
	Status            AwsDynamodbGlobalTableStatus `json:"status,omitempty"`
}

type AwsDynamodbGlobalTableSpecReplica struct {
	RegionName string `json:"region_name"`
}

type AwsDynamodbGlobalTableSpec struct {
	Name    string                       `json:"name"`
	Replica []AwsDynamodbGlobalTableSpec `json:"replica"`
	Arn     string                       `json:"arn"`
}



type AwsDynamodbGlobalTableStatus struct {
    Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsDynamodbGlobalTableList is a list of AwsDynamodbGlobalTables
type AwsDynamodbGlobalTableList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsDynamodbGlobalTable CRD objects
	Items []AwsDynamodbGlobalTable `json:"items,omitempty"`
}