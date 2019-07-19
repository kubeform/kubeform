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

type MariadbServer struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              MariadbServerSpec   `json:"spec,omitempty"`
	Status            MariadbServerStatus `json:"status,omitempty"`
}

type MariadbServerSpecSku struct {
	Capacity int    `json:"capacity" tf:"capacity"`
	Family   string `json:"family" tf:"family"`
	Name     string `json:"name" tf:"name"`
	Tier     string `json:"tier" tf:"tier"`
}

type MariadbServerSpecStorageProfile struct {
	// +optional
	BackupRetentionDays int `json:"backupRetentionDays,omitempty" tf:"backup_retention_days,omitempty"`
	// +optional
	GeoRedundantBackup string `json:"geoRedundantBackup,omitempty" tf:"geo_redundant_backup,omitempty"`
	StorageMb          int    `json:"storageMb" tf:"storage_mb"`
}

type MariadbServerSpec struct {
	AdministratorLogin string `json:"administratorLogin" tf:"administrator_login"`
	// Sensitive Data. Provide secret name which contains one value only
	AdministratorLoginPassword core.LocalObjectReference `json:"administratorLoginPassword" tf:"administrator_login_password"`
	Location                   string                    `json:"location" tf:"location"`
	Name                       string                    `json:"name" tf:"name"`
	ResourceGroupName          string                    `json:"resourceGroupName" tf:"resource_group_name"`
	// +kubebuilder:validation:MaxItems=1
	Sku            []MariadbServerSpecSku `json:"sku" tf:"sku"`
	SslEnforcement string                 `json:"sslEnforcement" tf:"ssl_enforcement"`
	// +kubebuilder:validation:MaxItems=1
	StorageProfile []MariadbServerSpecStorageProfile `json:"storageProfile" tf:"storage_profile"`
	// +optional
	Tags        map[string]string         `json:"tags,omitempty" tf:"tags,omitempty"`
	Version     string                    `json:"version" tf:"version"`
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`
}

type MariadbServerStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	TFState     *runtime.RawExtension `json:"tfState,omitempty"`
	TFStateHash string                `json:"tfStateHash,omitempty"`
	Output      *runtime.RawExtension `json:"output,omitempty"`
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
