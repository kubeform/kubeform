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

type AzurermMariadbServer struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AzurermMariadbServerSpec   `json:"spec,omitempty"`
	Status            AzurermMariadbServerStatus `json:"status,omitempty"`
}

type AzurermMariadbServerSpecStorageProfile struct {
	StorageMb           int    `json:"storage_mb"`
	BackupRetentionDays int    `json:"backup_retention_days"`
	GeoRedundantBackup  string `json:"geo_redundant_backup"`
}

type AzurermMariadbServerSpecSku struct {
	Name     string `json:"name"`
	Capacity int    `json:"capacity"`
	Tier     string `json:"tier"`
	Family   string `json:"family"`
}

type AzurermMariadbServerSpec struct {
	AdministratorLoginPassword string                     `json:"administrator_login_password"`
	StorageProfile             []AzurermMariadbServerSpec `json:"storage_profile"`
	Location                   string                     `json:"location"`
	Sku                        []AzurermMariadbServerSpec `json:"sku"`
	AdministratorLogin         string                     `json:"administrator_login"`
	SslEnforcement             string                     `json:"ssl_enforcement"`
	Fqdn                       string                     `json:"fqdn"`
	Tags                       map[string]string          `json:"tags"`
	Name                       string                     `json:"name"`
	ResourceGroupName          string                     `json:"resource_group_name"`
	Version                    string                     `json:"version"`
}

type AzurermMariadbServerStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AzurermMariadbServerList is a list of AzurermMariadbServers
type AzurermMariadbServerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AzurermMariadbServer CRD objects
	Items []AzurermMariadbServer `json:"items,omitempty"`
}
