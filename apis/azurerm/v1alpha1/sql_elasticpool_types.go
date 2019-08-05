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

type SqlElasticpool struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SqlElasticpoolSpec   `json:"spec,omitempty"`
	Status            SqlElasticpoolStatus `json:"status,omitempty"`
}

type SqlElasticpoolSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	CreationDate string `json:"creationDate" tf:"creation_date"`
	// +optional
	DbDtuMax int `json:"dbDtuMax,omitempty" tf:"db_dtu_max,omitempty"`
	// +optional
	DbDtuMin int    `json:"dbDtuMin,omitempty" tf:"db_dtu_min,omitempty"`
	Dtu      int    `json:"dtu" tf:"dtu"`
	Edition  string `json:"edition" tf:"edition"`
	Location string `json:"location" tf:"location"`
	Name     string `json:"name" tf:"name"`
	// +optional
	PoolSize          int    `json:"poolSize,omitempty" tf:"pool_size,omitempty"`
	ResourceGroupName string `json:"resourceGroupName" tf:"resource_group_name"`
	ServerName        string `json:"serverName" tf:"server_name"`
	// +optional
	Tags map[string]string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type SqlElasticpoolStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *SqlElasticpoolSpec `json:"output,omitempty"`
	// +optional
	State *apis.State `json:"state,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// SqlElasticpoolList is a list of SqlElasticpools
type SqlElasticpoolList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of SqlElasticpool CRD objects
	Items []SqlElasticpool `json:"items,omitempty"`
}
