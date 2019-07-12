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

type AzurermMysqlServer struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AzurermMysqlServerSpec   `json:"spec,omitempty"`
	Status            AzurermMysqlServerStatus `json:"status,omitempty"`
}

type AzurermMysqlServerSpecStorageProfile struct {
	StorageMb           int    `json:"storage_mb"`
	BackupRetentionDays int    `json:"backup_retention_days"`
	GeoRedundantBackup  string `json:"geo_redundant_backup"`
}

type AzurermMysqlServerSpecSku struct {
	Name     string `json:"name"`
	Capacity int    `json:"capacity"`
	Tier     string `json:"tier"`
	Family   string `json:"family"`
}

type AzurermMysqlServerSpec struct {
	Name                       string                   `json:"name"`
	ResourceGroupName          string                   `json:"resource_group_name"`
	AdministratorLoginPassword string                   `json:"administrator_login_password"`
	StorageProfile             []AzurermMysqlServerSpec `json:"storage_profile"`
	Location                   string                   `json:"location"`
	Sku                        []AzurermMysqlServerSpec `json:"sku"`
	AdministratorLogin         string                   `json:"administrator_login"`
	Version                    string                   `json:"version"`
	SslEnforcement             string                   `json:"ssl_enforcement"`
	Fqdn                       string                   `json:"fqdn"`
	Tags                       map[string]string        `json:"tags"`
}

type AzurermMysqlServerStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AzurermMysqlServerList is a list of AzurermMysqlServers
type AzurermMysqlServerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AzurermMysqlServer CRD objects
	Items []AzurermMysqlServer `json:"items,omitempty"`
}
