package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AzurermSqlElasticpool struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AzurermSqlElasticpoolSpec   `json:"spec,omitempty"`
	Status            AzurermSqlElasticpoolStatus `json:"status,omitempty"`
}

type AzurermSqlElasticpoolSpec struct {
	ResourceGroupName string            `json:"resource_group_name"`
	Dtu               int               `json:"dtu"`
	DbDtuMax          int               `json:"db_dtu_max"`
	CreationDate      string            `json:"creation_date"`
	Name              string            `json:"name"`
	ServerName        string            `json:"server_name"`
	Edition           string            `json:"edition"`
	DbDtuMin          int               `json:"db_dtu_min"`
	PoolSize          int               `json:"pool_size"`
	Tags              map[string]string `json:"tags"`
	Location          string            `json:"location"`
}

type AzurermSqlElasticpoolStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AzurermSqlElasticpoolList is a list of AzurermSqlElasticpools
type AzurermSqlElasticpoolList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AzurermSqlElasticpool CRD objects
	Items []AzurermSqlElasticpool `json:"items,omitempty"`
}
