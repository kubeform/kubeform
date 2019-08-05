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

type SqlDatabase struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SqlDatabaseSpec   `json:"spec,omitempty"`
	Status            SqlDatabaseStatus `json:"status,omitempty"`
}

type SqlDatabaseSpecImport struct {
	AdministratorLogin         string `json:"administratorLogin" tf:"administrator_login"`
	AdministratorLoginPassword string `json:"-" sensitive:"true" tf:"administrator_login_password"`
	AuthenticationType         string `json:"authenticationType" tf:"authentication_type"`
	// +optional
	OperationMode  string `json:"operationMode,omitempty" tf:"operation_mode,omitempty"`
	StorageKey     string `json:"-" sensitive:"true" tf:"storage_key"`
	StorageKeyType string `json:"storageKeyType" tf:"storage_key_type"`
	StorageURI     string `json:"storageURI" tf:"storage_uri"`
}

type SqlDatabaseSpecThreatDetectionPolicy struct {
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	DisabledAlerts []string `json:"disabledAlerts,omitempty" tf:"disabled_alerts,omitempty"`
	// +optional
	EmailAccountAdmins string `json:"emailAccountAdmins,omitempty" tf:"email_account_admins,omitempty"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	EmailAddresses []string `json:"emailAddresses,omitempty" tf:"email_addresses,omitempty"`
	// +optional
	RetentionDays int `json:"retentionDays,omitempty" tf:"retention_days,omitempty"`
	// +optional
	State string `json:"state,omitempty" tf:"state,omitempty"`
	// +optional
	StorageAccountAccessKey string `json:"-" sensitive:"true" tf:"storage_account_access_key,omitempty"`
	// +optional
	StorageEndpoint string `json:"storageEndpoint,omitempty" tf:"storage_endpoint,omitempty"`
	// +optional
	UseServerDefault string `json:"useServerDefault,omitempty" tf:"use_server_default,omitempty"`
}

type SqlDatabaseSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	KubeFormSecret *core.LocalObjectReference `json:"secret,omitempty" tf:"-"`

	// +optional
	Collation string `json:"collation,omitempty" tf:"collation,omitempty"`
	// +optional
	CreateMode               string `json:"createMode,omitempty" tf:"create_mode,omitempty"`
	CreationDate             string `json:"creationDate" tf:"creation_date"`
	DefaultSecondaryLocation string `json:"defaultSecondaryLocation" tf:"default_secondary_location"`
	// +optional
	Edition string `json:"edition,omitempty" tf:"edition,omitempty"`
	// +optional
	ElasticPoolName string `json:"elasticPoolName,omitempty" tf:"elastic_pool_name,omitempty"`
	Encryption      string `json:"encryption" tf:"encryption"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	Import   []SqlDatabaseSpecImport `json:"import,omitempty" tf:"import,omitempty"`
	Location string                  `json:"location" tf:"location"`
	// +optional
	MaxSizeBytes string `json:"maxSizeBytes,omitempty" tf:"max_size_bytes,omitempty"`
	Name         string `json:"name" tf:"name"`
	// +optional
	ReadScale bool `json:"readScale,omitempty" tf:"read_scale,omitempty"`
	// +optional
	RequestedServiceObjectiveID string `json:"requestedServiceObjectiveID,omitempty" tf:"requested_service_objective_id,omitempty"`
	// +optional
	RequestedServiceObjectiveName string `json:"requestedServiceObjectiveName,omitempty" tf:"requested_service_objective_name,omitempty"`
	ResourceGroupName             string `json:"resourceGroupName" tf:"resource_group_name"`
	// +optional
	RestorePointInTime string `json:"restorePointInTime,omitempty" tf:"restore_point_in_time,omitempty"`
	ServerName         string `json:"serverName" tf:"server_name"`
	// +optional
	SourceDatabaseDeletionDate string `json:"sourceDatabaseDeletionDate,omitempty" tf:"source_database_deletion_date,omitempty"`
	// +optional
	SourceDatabaseID string `json:"sourceDatabaseID,omitempty" tf:"source_database_id,omitempty"`
	// +optional
	Tags map[string]string `json:"tags,omitempty" tf:"tags,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	ThreatDetectionPolicy []SqlDatabaseSpecThreatDetectionPolicy `json:"threatDetectionPolicy,omitempty" tf:"threat_detection_policy,omitempty"`
}

type SqlDatabaseStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *SqlDatabaseSpec `json:"output,omitempty"`
	// +optional
	State *apis.State `json:"state,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// SqlDatabaseList is a list of SqlDatabases
type SqlDatabaseList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of SqlDatabase CRD objects
	Items []SqlDatabase `json:"items,omitempty"`
}
