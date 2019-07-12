package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

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
	Domain       string                            `json:"domain"`
	GroupByEmail string                            `json:"group_by_email"`
	SpecialGroup string                            `json:"special_group"`
	UserByEmail  string                            `json:"user_by_email"`
	View         []GoogleBigqueryDatasetSpecAccess `json:"view"`
	Role         string                            `json:"role"`
}

type GoogleBigqueryDatasetSpec struct {
	SelfLink                 string                      `json:"self_link"`
	Etag                     string                      `json:"etag"`
	LastModifiedTime         int                         `json:"last_modified_time"`
	Project                  string                      `json:"project"`
	DefaultTableExpirationMs int                         `json:"default_table_expiration_ms"`
	Description              string                      `json:"description"`
	Location                 string                      `json:"location"`
	Labels                   map[string]string           `json:"labels"`
	Access                   []GoogleBigqueryDatasetSpec `json:"access"`
	CreationTime             int                         `json:"creation_time"`
	DatasetId                string                      `json:"dataset_id"`
	FriendlyName             string                      `json:"friendly_name"`
}

type GoogleBigqueryDatasetStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// GoogleBigqueryDatasetList is a list of GoogleBigqueryDatasets
type GoogleBigqueryDatasetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of GoogleBigqueryDataset CRD objects
	Items []GoogleBigqueryDataset `json:"items,omitempty"`
}
