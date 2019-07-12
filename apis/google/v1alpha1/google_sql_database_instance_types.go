package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type GoogleSqlDatabaseInstance struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              GoogleSqlDatabaseInstanceSpec   `json:"spec,omitempty"`
	Status            GoogleSqlDatabaseInstanceStatus `json:"status,omitempty"`
}

type GoogleSqlDatabaseInstanceSpecIpAddress struct {
	IpAddress    string `json:"ip_address"`
	TimeToRetire string `json:"time_to_retire"`
}

type GoogleSqlDatabaseInstanceSpecServerCaCert struct {
	Cert            string `json:"cert"`
	CommonName      string `json:"common_name"`
	CreateTime      string `json:"create_time"`
	ExpirationTime  string `json:"expiration_time"`
	Sha1Fingerprint string `json:"sha1_fingerprint"`
}

type GoogleSqlDatabaseInstanceSpecReplicaConfiguration struct {
	DumpFilePath            string `json:"dump_file_path"`
	Password                string `json:"password"`
	SslCipher               string `json:"ssl_cipher"`
	Username                string `json:"username"`
	VerifyServerCertificate bool   `json:"verify_server_certificate"`
	ClientKey               string `json:"client_key"`
	ClientCertificate       string `json:"client_certificate"`
	ConnectRetryInterval    int    `json:"connect_retry_interval"`
	FailoverTarget          bool   `json:"failover_target"`
	MasterHeartbeatPeriod   int    `json:"master_heartbeat_period"`
	CaCertificate           string `json:"ca_certificate"`
}

type GoogleSqlDatabaseInstanceSpecSettingsDatabaseFlags struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type GoogleSqlDatabaseInstanceSpecSettingsMaintenanceWindow struct {
	Day         int    `json:"day"`
	Hour        int    `json:"hour"`
	UpdateTrack string `json:"update_track"`
}

type GoogleSqlDatabaseInstanceSpecSettingsBackupConfiguration struct {
	BinaryLogEnabled bool   `json:"binary_log_enabled"`
	Enabled          bool   `json:"enabled"`
	StartTime        string `json:"start_time"`
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

type GoogleSqlDatabaseInstanceSpecSettings struct {
	DatabaseFlags             []GoogleSqlDatabaseInstanceSpecSettings `json:"database_flags"`
	DiskSize                  int                                     `json:"disk_size"`
	DiskType                  string                                  `json:"disk_type"`
	UserLabels                map[string]string                       `json:"user_labels"`
	CrashSafeReplication      bool                                    `json:"crash_safe_replication"`
	DiskAutoresize            bool                                    `json:"disk_autoresize"`
	MaintenanceWindow         []GoogleSqlDatabaseInstanceSpecSettings `json:"maintenance_window"`
	Version                   int                                     `json:"version"`
	ActivationPolicy          string                                  `json:"activation_policy"`
	AuthorizedGaeApplications []string                                `json:"authorized_gae_applications"`
	BackupConfiguration       []GoogleSqlDatabaseInstanceSpecSettings `json:"backup_configuration"`
	IpConfiguration           []GoogleSqlDatabaseInstanceSpecSettings `json:"ip_configuration"`
	Tier                      string                                  `json:"tier"`
	LocationPreference        []GoogleSqlDatabaseInstanceSpecSettings `json:"location_preference"`
	PricingPlan               string                                  `json:"pricing_plan"`
	ReplicationType           string                                  `json:"replication_type"`
	AvailabilityType          string                                  `json:"availability_type"`
}

type GoogleSqlDatabaseInstanceSpec struct {
	IpAddress                  []GoogleSqlDatabaseInstanceSpec `json:"ip_address"`
	MasterInstanceName         string                          `json:"master_instance_name"`
	ServerCaCert               []GoogleSqlDatabaseInstanceSpec `json:"server_ca_cert"`
	ServiceAccountEmailAddress string                          `json:"service_account_email_address"`
	FirstIpAddress             string                          `json:"first_ip_address"`
	Name                       string                          `json:"name"`
	Project                    string                          `json:"project"`
	ReplicaConfiguration       []GoogleSqlDatabaseInstanceSpec `json:"replica_configuration"`
	Region                     string                          `json:"region"`
	Settings                   []GoogleSqlDatabaseInstanceSpec `json:"settings"`
	ConnectionName             string                          `json:"connection_name"`
	DatabaseVersion            string                          `json:"database_version"`
	SelfLink                   string                          `json:"self_link"`
}



type GoogleSqlDatabaseInstanceStatus struct {
    Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// GoogleSqlDatabaseInstanceList is a list of GoogleSqlDatabaseInstances
type GoogleSqlDatabaseInstanceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of GoogleSqlDatabaseInstance CRD objects
	Items []GoogleSqlDatabaseInstance `json:"items,omitempty"`
}