package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AzurermMssqlElasticpool struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AzurermMssqlElasticpoolSpec   `json:"spec,omitempty"`
	Status            AzurermMssqlElasticpoolStatus `json:"status,omitempty"`
}

type AzurermMssqlElasticpoolSpecElasticPoolProperties struct {
	State         string `json:"state"`
	CreationDate  string `json:"creation_date"`
	MaxSizeBytes  int    `json:"max_size_bytes"`
	ZoneRedundant bool   `json:"zone_redundant"`
	LicenseType   string `json:"license_type"`
}

type AzurermMssqlElasticpoolSpecSku struct {
	Name     string `json:"name"`
	Capacity int    `json:"capacity"`
	Tier     string `json:"tier"`
	Family   string `json:"family"`
}

type AzurermMssqlElasticpoolSpecPerDatabaseSettings struct {
	MinCapacity float64 `json:"min_capacity"`
	MaxCapacity float64 `json:"max_capacity"`
}

type AzurermMssqlElasticpoolSpec struct {
	Tags                  map[string]string             `json:"tags"`
	ResourceGroupName     string                        `json:"resource_group_name"`
	ServerName            string                        `json:"server_name"`
	ElasticPoolProperties []AzurermMssqlElasticpoolSpec `json:"elastic_pool_properties"`
	MaxSizeBytes          int                           `json:"max_size_bytes"`
	MaxSizeGb             float64                       `json:"max_size_gb"`
	ZoneRedundant         bool                          `json:"zone_redundant"`
	Name                  string                        `json:"name"`
	Location              string                        `json:"location"`
	Sku                   []AzurermMssqlElasticpoolSpec `json:"sku"`
	PerDatabaseSettings   []AzurermMssqlElasticpoolSpec `json:"per_database_settings"`
}

type AzurermMssqlElasticpoolStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AzurermMssqlElasticpoolList is a list of AzurermMssqlElasticpools
type AzurermMssqlElasticpoolList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AzurermMssqlElasticpool CRD objects
	Items []AzurermMssqlElasticpool `json:"items,omitempty"`
}
