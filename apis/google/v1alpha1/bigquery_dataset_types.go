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

type BigqueryDataset struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              BigqueryDatasetSpec   `json:"spec,omitempty"`
	Status            BigqueryDatasetStatus `json:"status,omitempty"`
}

type BigqueryDatasetSpecAccessView struct {
	DatasetID string `json:"datasetID" tf:"dataset_id"`
	ProjectID string `json:"projectID" tf:"project_id"`
	TableID   string `json:"tableID" tf:"table_id"`
}

type BigqueryDatasetSpecAccess struct {
	// +optional
	Domain string `json:"domain,omitempty" tf:"domain,omitempty"`
	// +optional
	GroupByEmail string `json:"groupByEmail,omitempty" tf:"group_by_email,omitempty"`
	// +optional
	Role string `json:"role,omitempty" tf:"role,omitempty"`
	// +optional
	SpecialGroup string `json:"specialGroup,omitempty" tf:"special_group,omitempty"`
	// +optional
	UserByEmail string `json:"userByEmail,omitempty" tf:"user_by_email,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	View []BigqueryDatasetSpecAccessView `json:"view,omitempty" tf:"view,omitempty"`
}

type BigqueryDatasetSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	// +optional
	Access       []BigqueryDatasetSpecAccess `json:"access,omitempty" tf:"access,omitempty"`
	CreationTime int                         `json:"creationTime" tf:"creation_time"`
	DatasetID    string                      `json:"datasetID" tf:"dataset_id"`
	// +optional
	DefaultTableExpirationMs int `json:"defaultTableExpirationMs,omitempty" tf:"default_table_expiration_ms,omitempty"`
	// +optional
	Description string `json:"description,omitempty" tf:"description,omitempty"`
	Etag        string `json:"etag" tf:"etag"`
	// +optional
	FriendlyName string `json:"friendlyName,omitempty" tf:"friendly_name,omitempty"`
	// +optional
	Labels           map[string]string `json:"labels,omitempty" tf:"labels,omitempty"`
	LastModifiedTime int               `json:"lastModifiedTime" tf:"last_modified_time"`
	// +optional
	Location string `json:"location,omitempty" tf:"location,omitempty"`
	// +optional
	Project  string `json:"project,omitempty" tf:"project,omitempty"`
	SelfLink string `json:"selfLink" tf:"self_link"`
}

type BigqueryDatasetStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *BigqueryDatasetSpec `json:"output,omitempty"`
	// +optional
	State *apis.State `json:"state,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// BigqueryDatasetList is a list of BigqueryDatasets
type BigqueryDatasetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of BigqueryDataset CRD objects
	Items []BigqueryDataset `json:"items,omitempty"`
}
