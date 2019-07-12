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

type AzurermSqlDatabase struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AzurermSqlDatabaseSpec   `json:"spec,omitempty"`
	Status            AzurermSqlDatabaseStatus `json:"status,omitempty"`
}

type AzurermSqlDatabaseSpecThreatDetectionPolicy struct {
	DisabledAlerts          []string `json:"disabled_alerts"`
	EmailAccountAdmins      string   `json:"email_account_admins"`
	EmailAddresses          []string `json:"email_addresses"`
	RetentionDays           int      `json:"retention_days"`
	State                   string   `json:"state"`
	StorageAccountAccessKey string   `json:"storage_account_access_key"`
	StorageEndpoint         string   `json:"storage_endpoint"`
	UseServerDefault        string   `json:"use_server_default"`
}

type AzurermSqlDatabaseSpecImport struct {
	StorageUri                 string `json:"storage_uri"`
	StorageKey                 string `json:"storage_key"`
	StorageKeyType             string `json:"storage_key_type"`
	AdministratorLogin         string `json:"administrator_login"`
	AdministratorLoginPassword string `json:"administrator_login_password"`
	AuthenticationType         string `json:"authentication_type"`
	OperationMode              string `json:"operation_mode"`
}

type AzurermSqlDatabaseSpec struct {
	CreationDate                  string                   `json:"creation_date"`
	ServerName                    string                   `json:"server_name"`
	RestorePointInTime            string                   `json:"restore_point_in_time"`
	Encryption                    string                   `json:"encryption"`
	ResourceGroupName             string                   `json:"resource_group_name"`
	RequestedServiceObjectiveName string                   `json:"requested_service_objective_name"`
	Tags                          map[string]string        `json:"tags"`
	Edition                       string                   `json:"edition"`
	MaxSizeBytes                  string                   `json:"max_size_bytes"`
	RequestedServiceObjectiveId   string                   `json:"requested_service_objective_id"`
	SourceDatabaseDeletionDate    string                   `json:"source_database_deletion_date"`
	ThreatDetectionPolicy         []AzurermSqlDatabaseSpec `json:"threat_detection_policy"`
	Location                      string                   `json:"location"`
	CreateMode                    string                   `json:"create_mode"`
	SourceDatabaseId              string                   `json:"source_database_id"`
	ElasticPoolName               string                   `json:"elastic_pool_name"`
	DefaultSecondaryLocation      string                   `json:"default_secondary_location"`
	ReadScale                     bool                     `json:"read_scale"`
	Name                          string                   `json:"name"`
	Import                        []AzurermSqlDatabaseSpec `json:"import"`
	Collation                     string                   `json:"collation"`
}

type AzurermSqlDatabaseStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AzurermSqlDatabaseList is a list of AzurermSqlDatabases
type AzurermSqlDatabaseList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AzurermSqlDatabase CRD objects
	Items []AzurermSqlDatabase `json:"items,omitempty"`
}
