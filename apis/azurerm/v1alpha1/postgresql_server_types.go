package v1alpha1

import (
	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

type PostgresqlServer struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              PostgresqlServerSpec   `json:"spec,omitempty"`
	Status            PostgresqlServerStatus `json:"status,omitempty"`
}

type PostgresqlServerSpecSku struct {
	Capacity int    `json:"capacity" tf:"capacity"`
	Family   string `json:"family" tf:"family"`
	Name     string `json:"name" tf:"name"`
	Tier     string `json:"tier" tf:"tier"`
}

type PostgresqlServerSpecStorageProfile struct {
	// +optional
	BackupRetentionDays int `json:"backupRetentionDays,omitempty" tf:"backup_retention_days,omitempty"`
	// +optional
	GeoRedundantBackup string `json:"geoRedundantBackup,omitempty" tf:"geo_redundant_backup,omitempty"`
	StorageMb          int    `json:"storageMb" tf:"storage_mb"`
}

type PostgresqlServerSpec struct {
	AdministratorLogin string `json:"administratorLogin" tf:"administrator_login"`
	// Sensitive Data. Provide secret name which contains one value only
	AdministratorLoginPassword core.LocalObjectReference `json:"administratorLoginPassword" tf:"administrator_login_password"`
	Location                   string                    `json:"location" tf:"location"`
	Name                       string                    `json:"name" tf:"name"`
	ResourceGroupName          string                    `json:"resourceGroupName" tf:"resource_group_name"`
	// +kubebuilder:validation:MaxItems=1
	Sku            []PostgresqlServerSpecSku `json:"sku" tf:"sku"`
	SslEnforcement string                    `json:"sslEnforcement" tf:"ssl_enforcement"`
	// +kubebuilder:validation:MaxItems=1
	StorageProfile []PostgresqlServerSpecStorageProfile `json:"storageProfile" tf:"storage_profile"`
	// +optional
	Tags        map[string]string         `json:"tags,omitempty" tf:"tags,omitempty"`
	Version     string                    `json:"version" tf:"version"`
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`
}

type PostgresqlServerStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	TFState     *runtime.RawExtension `json:"tfState,omitempty"`
	TFStateHash string                `json:"tfStateHash,omitempty"`
	Output      *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// PostgresqlServerList is a list of PostgresqlServers
type PostgresqlServerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of PostgresqlServer CRD objects
	Items []PostgresqlServer `json:"items,omitempty"`
}
