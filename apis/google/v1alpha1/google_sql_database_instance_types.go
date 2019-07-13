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

type GoogleSqlDatabaseInstance struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              GoogleSqlDatabaseInstanceSpec   `json:"spec,omitempty"`
	Status            GoogleSqlDatabaseInstanceStatus `json:"status,omitempty"`
}

type GoogleSqlDatabaseInstanceSpecSettingsBackupConfiguration struct {
	Enabled          bool   `json:"enabled"`
	StartTime        string `json:"start_time"`
	BinaryLogEnabled bool   `json:"binary_log_enabled"`
}

type GoogleSqlDatabaseInstanceSpecSettingsDatabaseFlags struct {
	Value string `json:"value"`
	Name  string `json:"name"`
}

type GoogleSqlDatabaseInstanceSpecSettingsIpConfigurationAuthorizedNetworks struct {
	Value          string `json:"value"`
	ExpirationTime string `json:"expiration_time"`
	Name           string `json:"name"`
}

type GoogleSqlDatabaseInstanceSpecSettingsIpConfiguration struct {
	AuthorizedNetworks []GoogleSqlDatabaseInstanceSpecSettingsIpConfiguration `json:"authorized_networks"`
	Ipv4Enabled        bool                                                   `json:"ipv4_enabled"`
	RequireSsl         bool                                                   `json:"require_ssl"`
	PrivateNetwork     string                                                 `json:"private_network"`
}

type GoogleSqlDatabaseInstanceSpecSettingsLocationPreference struct {
	FollowGaeApplication string `json:"follow_gae_application"`
	Zone                 string `json:"zone"`
}

type GoogleSqlDatabaseInstanceSpecSettingsMaintenanceWindow struct {
	UpdateTrack string `json:"update_track"`
	Day         int    `json:"day"`
	Hour        int    `json:"hour"`
}

type GoogleSqlDatabaseInstanceSpecSettings struct {
	ActivationPolicy          string                                  `json:"activation_policy"`
	CrashSafeReplication      bool                                    `json:"crash_safe_replication"`
	DiskAutoresize            bool                                    `json:"disk_autoresize"`
	DiskSize                  int                                     `json:"disk_size"`
	DiskType                  string                                  `json:"disk_type"`
	PricingPlan               string                                  `json:"pricing_plan"`
	ReplicationType           string                                  `json:"replication_type"`
	Tier                      string                                  `json:"tier"`
	UserLabels                map[string]string                       `json:"user_labels"`
	AuthorizedGaeApplications []string                                `json:"authorized_gae_applications"`
	BackupConfiguration       []GoogleSqlDatabaseInstanceSpecSettings `json:"backup_configuration"`
	Version                   int                                     `json:"version"`
	DatabaseFlags             []GoogleSqlDatabaseInstanceSpecSettings `json:"database_flags"`
	IpConfiguration           []GoogleSqlDatabaseInstanceSpecSettings `json:"ip_configuration"`
	LocationPreference        []GoogleSqlDatabaseInstanceSpecSettings `json:"location_preference"`
	MaintenanceWindow         []GoogleSqlDatabaseInstanceSpecSettings `json:"maintenance_window"`
	AvailabilityType          string                                  `json:"availability_type"`
}

type GoogleSqlDatabaseInstanceSpecIpAddress struct {
	IpAddress    string `json:"ip_address"`
	TimeToRetire string `json:"time_to_retire"`
}

type GoogleSqlDatabaseInstanceSpecServerCaCert struct {
	CommonName      string `json:"common_name"`
	CreateTime      string `json:"create_time"`
	ExpirationTime  string `json:"expiration_time"`
	Sha1Fingerprint string `json:"sha1_fingerprint"`
	Cert            string `json:"cert"`
}

type GoogleSqlDatabaseInstanceSpecReplicaConfiguration struct {
	ClientKey               string `json:"client_key"`
	ConnectRetryInterval    int    `json:"connect_retry_interval"`
	DumpFilePath            string `json:"dump_file_path"`
	FailoverTarget          bool   `json:"failover_target"`
	Password                string `json:"password"`
	SslCipher               string `json:"ssl_cipher"`
	VerifyServerCertificate bool   `json:"verify_server_certificate"`
	CaCertificate           string `json:"ca_certificate"`
	MasterHeartbeatPeriod   int    `json:"master_heartbeat_period"`
	Username                string `json:"username"`
	ClientCertificate       string `json:"client_certificate"`
}

type GoogleSqlDatabaseInstanceSpec struct {
	Settings                   []GoogleSqlDatabaseInstanceSpec `json:"settings"`
	IpAddress                  []GoogleSqlDatabaseInstanceSpec `json:"ip_address"`
	MasterInstanceName         string                          `json:"master_instance_name"`
	Project                    string                          `json:"project"`
	ServerCaCert               []GoogleSqlDatabaseInstanceSpec `json:"server_ca_cert"`
	SelfLink                   string                          `json:"self_link"`
	Region                     string                          `json:"region"`
	DatabaseVersion            string                          `json:"database_version"`
	FirstIpAddress             string                          `json:"first_ip_address"`
	Name                       string                          `json:"name"`
	ReplicaConfiguration       []GoogleSqlDatabaseInstanceSpec `json:"replica_configuration"`
	ServiceAccountEmailAddress string                          `json:"service_account_email_address"`
	ConnectionName             string                          `json:"connection_name"`
}



type GoogleSqlDatabaseInstanceStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// GoogleSqlDatabaseInstanceList is a list of GoogleSqlDatabaseInstances
type GoogleSqlDatabaseInstanceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of GoogleSqlDatabaseInstance CRD objects
	Items []GoogleSqlDatabaseInstance `json:"items,omitempty"`
}