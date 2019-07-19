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

type SqlDatabaseInstance struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SqlDatabaseInstanceSpec   `json:"spec,omitempty"`
	Status            SqlDatabaseInstanceStatus `json:"status,omitempty"`
}

type SqlDatabaseInstanceSpecReplicaConfiguration struct {
	// +optional
	CaCertificate string `json:"caCertificate,omitempty" tf:"ca_certificate,omitempty"`
	// +optional
	ClientCertificate string `json:"clientCertificate,omitempty" tf:"client_certificate,omitempty"`
	// +optional
	ClientKey string `json:"clientKey,omitempty" tf:"client_key,omitempty"`
	// +optional
	ConnectRetryInterval int `json:"connectRetryInterval,omitempty" tf:"connect_retry_interval,omitempty"`
	// +optional
	DumpFilePath string `json:"dumpFilePath,omitempty" tf:"dump_file_path,omitempty"`
	// +optional
	FailoverTarget bool `json:"failoverTarget,omitempty" tf:"failover_target,omitempty"`
	// +optional
	MasterHeartbeatPeriod int `json:"masterHeartbeatPeriod,omitempty" tf:"master_heartbeat_period,omitempty"`
	// +optional
	// Sensitive Data. Provide secret name which contains one value only
	Password core.LocalObjectReference `json:"password,omitempty" tf:"password,omitempty"`
	// +optional
	SslCipher string `json:"sslCipher,omitempty" tf:"ssl_cipher,omitempty"`
	// +optional
	Username string `json:"username,omitempty" tf:"username,omitempty"`
	// +optional
	VerifyServerCertificate bool `json:"verifyServerCertificate,omitempty" tf:"verify_server_certificate,omitempty"`
}

type SqlDatabaseInstanceSpecSettingsBackupConfiguration struct {
	// +optional
	BinaryLogEnabled bool `json:"binaryLogEnabled,omitempty" tf:"binary_log_enabled,omitempty"`
	// +optional
	Enabled bool `json:"enabled,omitempty" tf:"enabled,omitempty"`
	// +optional
	StartTime string `json:"startTime,omitempty" tf:"start_time,omitempty"`
}

type SqlDatabaseInstanceSpecSettingsDatabaseFlags struct {
	// +optional
	Name string `json:"name,omitempty" tf:"name,omitempty"`
	// +optional
	Value string `json:"value,omitempty" tf:"value,omitempty"`
}

type SqlDatabaseInstanceSpecSettingsIpConfigurationAuthorizedNetworks struct {
	// +optional
	ExpirationTime string `json:"expirationTime,omitempty" tf:"expiration_time,omitempty"`
	// +optional
	Name string `json:"name,omitempty" tf:"name,omitempty"`
	// +optional
	Value string `json:"value,omitempty" tf:"value,omitempty"`
}

type SqlDatabaseInstanceSpecSettingsIpConfiguration struct {
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	AuthorizedNetworks []SqlDatabaseInstanceSpecSettingsIpConfigurationAuthorizedNetworks `json:"authorizedNetworks,omitempty" tf:"authorized_networks,omitempty"`
	// +optional
	Ipv4Enabled bool `json:"ipv4Enabled,omitempty" tf:"ipv4_enabled,omitempty"`
	// +optional
	PrivateNetwork string `json:"privateNetwork,omitempty" tf:"private_network,omitempty"`
	// +optional
	RequireSsl bool `json:"requireSsl,omitempty" tf:"require_ssl,omitempty"`
}

type SqlDatabaseInstanceSpecSettingsLocationPreference struct {
	// +optional
	FollowGaeApplication string `json:"followGaeApplication,omitempty" tf:"follow_gae_application,omitempty"`
	// +optional
	Zone string `json:"zone,omitempty" tf:"zone,omitempty"`
}

type SqlDatabaseInstanceSpecSettingsMaintenanceWindow struct {
	// +optional
	Day int `json:"day,omitempty" tf:"day,omitempty"`
	// +optional
	Hour int `json:"hour,omitempty" tf:"hour,omitempty"`
	// +optional
	UpdateTrack string `json:"updateTrack,omitempty" tf:"update_track,omitempty"`
}

type SqlDatabaseInstanceSpecSettings struct {
	// +optional
	ActivationPolicy string `json:"activationPolicy,omitempty" tf:"activation_policy,omitempty"`
	// +optional
	AuthorizedGaeApplications []string `json:"authorizedGaeApplications,omitempty" tf:"authorized_gae_applications,omitempty"`
	// +optional
	AvailabilityType string `json:"availabilityType,omitempty" tf:"availability_type,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	BackupConfiguration []SqlDatabaseInstanceSpecSettingsBackupConfiguration `json:"backupConfiguration,omitempty" tf:"backup_configuration,omitempty"`
	// +optional
	CrashSafeReplication bool `json:"crashSafeReplication,omitempty" tf:"crash_safe_replication,omitempty"`
	// +optional
	DatabaseFlags []SqlDatabaseInstanceSpecSettingsDatabaseFlags `json:"databaseFlags,omitempty" tf:"database_flags,omitempty"`
	// +optional
	DiskAutoresize bool `json:"diskAutoresize,omitempty" tf:"disk_autoresize,omitempty"`
	// +optional
	DiskSize int `json:"diskSize,omitempty" tf:"disk_size,omitempty"`
	// +optional
	DiskType string `json:"diskType,omitempty" tf:"disk_type,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	IpConfiguration []SqlDatabaseInstanceSpecSettingsIpConfiguration `json:"ipConfiguration,omitempty" tf:"ip_configuration,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	LocationPreference []SqlDatabaseInstanceSpecSettingsLocationPreference `json:"locationPreference,omitempty" tf:"location_preference,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	MaintenanceWindow []SqlDatabaseInstanceSpecSettingsMaintenanceWindow `json:"maintenanceWindow,omitempty" tf:"maintenance_window,omitempty"`
	// +optional
	PricingPlan string `json:"pricingPlan,omitempty" tf:"pricing_plan,omitempty"`
	// +optional
	ReplicationType string `json:"replicationType,omitempty" tf:"replication_type,omitempty"`
	Tier            string `json:"tier" tf:"tier"`
	// +optional
	UserLabels map[string]string `json:"userLabels,omitempty" tf:"user_labels,omitempty"`
}

type SqlDatabaseInstanceSpec struct {
	// +optional
	DatabaseVersion string `json:"databaseVersion,omitempty" tf:"database_version,omitempty"`
	// +optional
	MasterInstanceName string `json:"masterInstanceName,omitempty" tf:"master_instance_name,omitempty"`
	// +optional
	Name string `json:"name,omitempty" tf:"name,omitempty"`
	// +optional
	Project string `json:"project,omitempty" tf:"project,omitempty"`
	// +optional
	Region string `json:"region,omitempty" tf:"region,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	ReplicaConfiguration []SqlDatabaseInstanceSpecReplicaConfiguration `json:"replicaConfiguration,omitempty" tf:"replica_configuration,omitempty"`
	// +kubebuilder:validation:MaxItems=1
	Settings    []SqlDatabaseInstanceSpecSettings `json:"settings" tf:"settings"`
	ProviderRef core.LocalObjectReference         `json:"providerRef" tf:"-"`
}

type SqlDatabaseInstanceStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	TFState *runtime.RawExtension `json:"tfState,omitempty"`
	Output  *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// SqlDatabaseInstanceList is a list of SqlDatabaseInstances
type SqlDatabaseInstanceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of SqlDatabaseInstance CRD objects
	Items []SqlDatabaseInstance `json:"items,omitempty"`
}
