package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AzurermPostgresqlServer struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AzurermPostgresqlServerSpec   `json:"spec,omitempty"`
	Status            AzurermPostgresqlServerStatus `json:"status,omitempty"`
}

type AzurermPostgresqlServerSpecSku struct {
	Name     string `json:"name"`
	Capacity int    `json:"capacity"`
	Tier     string `json:"tier"`
	Family   string `json:"family"`
}

type AzurermPostgresqlServerSpecStorageProfile struct {
	StorageMb           int    `json:"storage_mb"`
	BackupRetentionDays int    `json:"backup_retention_days"`
	GeoRedundantBackup  string `json:"geo_redundant_backup"`
}

type AzurermPostgresqlServerSpec struct {
	Sku                        []AzurermPostgresqlServerSpec `json:"sku"`
	AdministratorLoginPassword string                        `json:"administrator_login_password"`
	Fqdn                       string                        `json:"fqdn"`
	Tags                       map[string]string             `json:"tags"`
	Name                       string                        `json:"name"`
	Location                   string                        `json:"location"`
	ResourceGroupName          string                        `json:"resource_group_name"`
	AdministratorLogin         string                        `json:"administrator_login"`
	Version                    string                        `json:"version"`
	StorageProfile             []AzurermPostgresqlServerSpec `json:"storage_profile"`
	SslEnforcement             string                        `json:"ssl_enforcement"`
}



type AzurermPostgresqlServerStatus struct {
    Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AzurermPostgresqlServerList is a list of AzurermPostgresqlServers
type AzurermPostgresqlServerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AzurermPostgresqlServer CRD objects
	Items []AzurermPostgresqlServer `json:"items,omitempty"`
}