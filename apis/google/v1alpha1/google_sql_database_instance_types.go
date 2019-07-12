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

type GoogleSqlDatabaseInstanceSpecReplicaConfiguration struct {
	Username                string `json:"username"`
	ClientKey               string `json:"client_key"`
	FailoverTarget          bool   `json:"failover_target"`
	Password                string `json:"password"`
	SslCipher               string `json:"ssl_cipher"`
	MasterHeartbeatPeriod   int    `json:"master_heartbeat_period"`
	VerifyServerCertificate bool   `json:"verify_server_certificate"`
	CaCertificate           string `json:"ca_certificate"`
	ClientCertificate       string `json:"client_certificate"`
	ConnectRetryInterval    int    `json:"connect_retry_interval"`
	DumpFilePath            string `json:"dump_file_path"`
}

type GoogleSqlDatabaseInstanceSpecSettingsLocationPreference struct {
	FollowGaeApplication string `json:"follow_gae_application"`
	Zone                 string `json:"zone"`
}

type GoogleSqlDatabaseInstanceSpecSettingsBackupConfiguration struct {
	StartTime        string `json:"start_time"`
	BinaryLogEnabled bool   `json:"binary_log_enabled"`
	Enabled          bool   `json:"enabled"`
}

type GoogleSqlDatabaseInstanceSpecSettingsIpConfigurationAuthorizedNetworks struct {
	ExpirationTime string `json:"expiration_time"`
	Name           string `json:"name"`
	Value          string `json:"value"`
}

type GoogleSqlDatabaseInstanceSpecSettingsIpConfiguration struct {
	AuthorizedNetworks []GoogleSqlDatabaseInstanceSpecSettingsIpConfiguration `json:"authorized_networks"`
	Ipv4Enabled        bool                                                   `json:"ipv4_enabled"`
	RequireSsl         bool                                                   `json:"require_ssl"`
	PrivateNetwork     string                                                 `json:"private_network"`
}

type GoogleSqlDatabaseInstanceSpecSettingsMaintenanceWindow struct {
	Hour        int    `json:"hour"`
	UpdateTrack string `json:"update_track"`
	Day         int    `json:"day"`
}

type GoogleSqlDatabaseInstanceSpecSettingsDatabaseFlags struct {
	Value string `json:"value"`
	Name  string `json:"name"`
}

type GoogleSqlDatabaseInstanceSpecSettings struct {
	Version                   int                                     `json:"version"`
	AuthorizedGaeApplications []string                                `json:"authorized_gae_applications"`
	LocationPreference        []GoogleSqlDatabaseInstanceSpecSettings `json:"location_preference"`
	Tier                      string                                  `json:"tier"`
	DiskType                  string                                  `json:"disk_type"`
	BackupConfiguration       []GoogleSqlDatabaseInstanceSpecSettings `json:"backup_configuration"`
	CrashSafeReplication      bool                                    `json:"crash_safe_replication"`
	DiskAutoresize            bool                                    `json:"disk_autoresize"`
	DiskSize                  int                                     `json:"disk_size"`
	IpConfiguration           []GoogleSqlDatabaseInstanceSpecSettings `json:"ip_configuration"`
	MaintenanceWindow         []GoogleSqlDatabaseInstanceSpecSettings `json:"maintenance_window"`
	ActivationPolicy          string                                  `json:"activation_policy"`
	AvailabilityType          string                                  `json:"availability_type"`
	ReplicationType           string                                  `json:"replication_type"`
	UserLabels                map[string]string                       `json:"user_labels"`
	DatabaseFlags             []GoogleSqlDatabaseInstanceSpecSettings `json:"database_flags"`
	PricingPlan               string                                  `json:"pricing_plan"`
}

type GoogleSqlDatabaseInstanceSpecIpAddress struct {
	IpAddress    string `json:"ip_address"`
	TimeToRetire string `json:"time_to_retire"`
}

type GoogleSqlDatabaseInstanceSpecServerCaCert struct {
	Sha1Fingerprint string `json:"sha1_fingerprint"`
	Cert            string `json:"cert"`
	CommonName      string `json:"common_name"`
	CreateTime      string `json:"create_time"`
	ExpirationTime  string `json:"expiration_time"`
}

type GoogleSqlDatabaseInstanceSpec struct {
	Name                       string                          `json:"name"`
	ReplicaConfiguration       []GoogleSqlDatabaseInstanceSpec `json:"replica_configuration"`
	ServiceAccountEmailAddress string                          `json:"service_account_email_address"`
	SelfLink                   string                          `json:"self_link"`
	MasterInstanceName         string                          `json:"master_instance_name"`
	Project                    string                          `json:"project"`
	Region                     string                          `json:"region"`
	Settings                   []GoogleSqlDatabaseInstanceSpec `json:"settings"`
	ConnectionName             string                          `json:"connection_name"`
	DatabaseVersion            string                          `json:"database_version"`
	IpAddress                  []GoogleSqlDatabaseInstanceSpec `json:"ip_address"`
	FirstIpAddress             string                          `json:"first_ip_address"`
	ServerCaCert               []GoogleSqlDatabaseInstanceSpec `json:"server_ca_cert"`
}

type GoogleSqlDatabaseInstanceStatus struct {
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
