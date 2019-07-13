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

type AzurermSearchService struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AzurermSearchServiceSpec   `json:"spec,omitempty"`
	Status            AzurermSearchServiceStatus `json:"status,omitempty"`
}

type AzurermSearchServiceSpec struct {
	ResourceGroupName string            `json:"resource_group_name"`
	ReplicaCount      int               `json:"replica_count"`
	SecondaryKey      string            `json:"secondary_key"`
	Name              string            `json:"name"`
	Location          string            `json:"location"`
	Sku               string            `json:"sku"`
	PartitionCount    int               `json:"partition_count"`
	PrimaryKey        string            `json:"primary_key"`
	Tags              map[string]string `json:"tags"`
}



type AzurermSearchServiceStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AzurermSearchServiceList is a list of AzurermSearchServices
type AzurermSearchServiceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AzurermSearchService CRD objects
	Items []AzurermSearchService `json:"items,omitempty"`
}