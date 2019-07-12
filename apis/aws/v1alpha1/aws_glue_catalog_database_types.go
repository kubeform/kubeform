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

type AwsGlueCatalogDatabase struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsGlueCatalogDatabaseSpec   `json:"spec,omitempty"`
	Status            AwsGlueCatalogDatabaseStatus `json:"status,omitempty"`
}

type AwsGlueCatalogDatabaseSpec struct {
	CatalogId   string            `json:"catalog_id"`
	Name        string            `json:"name"`
	Description string            `json:"description"`
	LocationUri string            `json:"location_uri"`
	Parameters  map[string]string `json:"parameters"`
}

type AwsGlueCatalogDatabaseStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AwsGlueCatalogDatabaseList is a list of AwsGlueCatalogDatabases
type AwsGlueCatalogDatabaseList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsGlueCatalogDatabase CRD objects
	Items []AwsGlueCatalogDatabase `json:"items,omitempty"`
}
