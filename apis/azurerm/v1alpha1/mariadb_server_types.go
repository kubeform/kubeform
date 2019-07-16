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

type MariadbServer struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              MariadbServerSpec   `json:"spec,omitempty"`
	Status            MariadbServerStatus `json:"status,omitempty"`
}

type MariadbServerSpecSku struct {
	Capacity int    `json:"capacity"`
	Family   string `json:"family"`
	Name     string `json:"name"`
	Tier     string `json:"tier"`
}

type MariadbServerSpecStorageProfile struct {
	// +optional
	BackupRetentionDays int `json:"backup_retention_days,omitempty"`
	// +optional
	GeoRedundantBackup string `json:"geo_redundant_backup,omitempty"`
	StorageMb          int    `json:"storage_mb"`
}

type MariadbServerSpec struct {
	AdministratorLogin         string `json:"administrator_login"`
	AdministratorLoginPassword string `json:"administrator_login_password"`
	Location                   string `json:"location"`
	Name                       string `json:"name"`
	ResourceGroupName          string `json:"resource_group_name"`
	// +kubebuilder:validation:MaxItems=1
	Sku            []MariadbServerSpec `json:"sku"`
	SslEnforcement string              `json:"ssl_enforcement"`
	// +kubebuilder:validation:MaxItems=1
	StorageProfile []MariadbServerSpec `json:"storage_profile"`
	Version        string              `json:"version"`
}

type MariadbServerStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// MariadbServerList is a list of MariadbServers
type MariadbServerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of MariadbServer CRD objects
	Items []MariadbServer `json:"items,omitempty"`
}
