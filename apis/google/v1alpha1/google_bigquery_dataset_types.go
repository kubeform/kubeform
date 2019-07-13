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

type GoogleBigqueryDataset struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              GoogleBigqueryDatasetSpec   `json:"spec,omitempty"`
	Status            GoogleBigqueryDatasetStatus `json:"status,omitempty"`
}

type GoogleBigqueryDatasetSpecAccessView struct {
	TableId   string `json:"table_id"`
	ProjectId string `json:"project_id"`
	DatasetId string `json:"dataset_id"`
}

type GoogleBigqueryDatasetSpecAccess struct {
	Role         string                            `json:"role"`
	Domain       string                            `json:"domain"`
	GroupByEmail string                            `json:"group_by_email"`
	SpecialGroup string                            `json:"special_group"`
	UserByEmail  string                            `json:"user_by_email"`
	View         []GoogleBigqueryDatasetSpecAccess `json:"view"`
}

type GoogleBigqueryDatasetSpec struct {
	Access                   []GoogleBigqueryDatasetSpec `json:"access"`
	SelfLink                 string                      `json:"self_link"`
	DatasetId                string                      `json:"dataset_id"`
	Project                  string                      `json:"project"`
	FriendlyName             string                      `json:"friendly_name"`
	Location                 string                      `json:"location"`
	DefaultTableExpirationMs int                         `json:"default_table_expiration_ms"`
	Labels                   map[string]string           `json:"labels"`
	Etag                     string                      `json:"etag"`
	CreationTime             int                         `json:"creation_time"`
	LastModifiedTime         int                         `json:"last_modified_time"`
	Description              string                      `json:"description"`
}



type GoogleBigqueryDatasetStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// GoogleBigqueryDatasetList is a list of GoogleBigqueryDatasets
type GoogleBigqueryDatasetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of GoogleBigqueryDataset CRD objects
	Items []GoogleBigqueryDataset `json:"items,omitempty"`
}