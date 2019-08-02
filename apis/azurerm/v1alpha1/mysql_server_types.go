package v1alpha1

import (
	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

type MysqlServer struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              MysqlServerSpec   `json:"spec,omitempty"`
	Status            MysqlServerStatus `json:"status,omitempty"`
}

type MysqlServerSpecSku struct {
	Capacity int    `json:"capacity" tf:"capacity"`
	Family   string `json:"family" tf:"family"`
	Name     string `json:"name" tf:"name"`
	Tier     string `json:"tier" tf:"tier"`
}

type MysqlServerSpecStorageProfile struct {
	// +optional
	BackupRetentionDays int `json:"backupRetentionDays,omitempty" tf:"backup_retention_days,omitempty"`
	// +optional
	GeoRedundantBackup string `json:"geoRedundantBackup,omitempty" tf:"geo_redundant_backup,omitempty"`
	StorageMb          int    `json:"storageMb" tf:"storage_mb"`
}

type MysqlServerSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	Secret *core.LocalObjectReference `json:"secret,omitempty" tf:"-"`

	AdministratorLogin         string `json:"administratorLogin" tf:"administrator_login"`
	AdministratorLoginPassword string `json:"-" sensitive:"true" tf:"administrator_login_password"`
	Location                   string `json:"location" tf:"location"`
	Name                       string `json:"name" tf:"name"`
	ResourceGroupName          string `json:"resourceGroupName" tf:"resource_group_name"`
	// +kubebuilder:validation:MaxItems=1
	Sku            []MysqlServerSpecSku `json:"sku" tf:"sku"`
	SslEnforcement string               `json:"sslEnforcement" tf:"ssl_enforcement"`
	// +kubebuilder:validation:MaxItems=1
	StorageProfile []MysqlServerSpecStorageProfile `json:"storageProfile" tf:"storage_profile"`
	// +optional
	Tags    map[string]string `json:"tags,omitempty" tf:"tags,omitempty"`
	Version string            `json:"version" tf:"version"`
}

type MysqlServerStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	State *MysqlServerSpec `json:"state,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// MysqlServerList is a list of MysqlServers
type MysqlServerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of MysqlServer CRD objects
	Items []MysqlServer `json:"items,omitempty"`
}
