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

type AzurermRedisCache struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AzurermRedisCacheSpec   `json:"spec,omitempty"`
	Status            AzurermRedisCacheStatus `json:"status,omitempty"`
}

type AzurermRedisCacheSpecRedisConfiguration struct {
	Maxclients                     int    `json:"maxclients"`
	MaxmemoryPolicy                string `json:"maxmemory_policy"`
	AofStorageConnectionString1    string `json:"aof_storage_connection_string_1"`
	RdbBackupEnabled               bool   `json:"rdb_backup_enabled"`
	RdbBackupMaxSnapshotCount      int    `json:"rdb_backup_max_snapshot_count"`
	AofStorageConnectionString0    string `json:"aof_storage_connection_string_0"`
	MaxmemoryDelta                 int    `json:"maxmemory_delta"`
	MaxmemoryReserved              int    `json:"maxmemory_reserved"`
	RdbBackupFrequency             int    `json:"rdb_backup_frequency"`
	RdbStorageConnectionString     string `json:"rdb_storage_connection_string"`
	EnableAuthentication           bool   `json:"enable_authentication"`
	MaxfragmentationmemoryReserved int    `json:"maxfragmentationmemory_reserved"`
	NotifyKeyspaceEvents           string `json:"notify_keyspace_events"`
	AofBackupEnabled               bool   `json:"aof_backup_enabled"`
}

type AzurermRedisCacheSpecPatchSchedule struct {
	StartHourUtc int    `json:"start_hour_utc"`
	DayOfWeek    string `json:"day_of_week"`
}

type AzurermRedisCacheSpec struct {
	Capacity               int                     `json:"capacity"`
	Family                 string                  `json:"family"`
	ShardCount             int                     `json:"shard_count"`
	Hostname               string                  `json:"hostname"`
	SslPort                int                     `json:"ssl_port"`
	Name                   string                  `json:"name"`
	Location               string                  `json:"location"`
	ResourceGroupName      string                  `json:"resource_group_name"`
	PrimaryAccessKey       string                  `json:"primary_access_key"`
	SecondaryAccessKey     string                  `json:"secondary_access_key"`
	Tags                   map[string]string       `json:"tags"`
	MinimumTlsVersion      string                  `json:"minimum_tls_version"`
	PrivateStaticIpAddress string                  `json:"private_static_ip_address"`
	RedisConfiguration     []AzurermRedisCacheSpec `json:"redis_configuration"`
	EnableNonSslPort       bool                    `json:"enable_non_ssl_port"`
	PatchSchedule          []AzurermRedisCacheSpec `json:"patch_schedule"`
	Port                   int                     `json:"port"`
	Zones                  []string                `json:"zones"`
	SkuName                string                  `json:"sku_name"`
	SubnetId               string                  `json:"subnet_id"`
}

type AzurermRedisCacheStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AzurermRedisCacheList is a list of AzurermRedisCaches
type AzurermRedisCacheList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AzurermRedisCache CRD objects
	Items []AzurermRedisCache `json:"items,omitempty"`
}
