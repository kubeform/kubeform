package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AzurermRedisCache struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AzurermRedisCacheSpec   `json:"spec,omitempty"`
	Status            AzurermRedisCacheStatus `json:"status,omitempty"`
}

type AzurermRedisCacheSpecRedisConfiguration struct {
	RdbStorageConnectionString     string `json:"rdb_storage_connection_string"`
	NotifyKeyspaceEvents           string `json:"notify_keyspace_events"`
	AofStorageConnectionString1    string `json:"aof_storage_connection_string_1"`
	Maxclients                     int    `json:"maxclients"`
	MaxmemoryDelta                 int    `json:"maxmemory_delta"`
	AofBackupEnabled               bool   `json:"aof_backup_enabled"`
	MaxmemoryPolicy                string `json:"maxmemory_policy"`
	RdbBackupFrequency             int    `json:"rdb_backup_frequency"`
	AofStorageConnectionString0    string `json:"aof_storage_connection_string_0"`
	MaxfragmentationmemoryReserved int    `json:"maxfragmentationmemory_reserved"`
	RdbBackupEnabled               bool   `json:"rdb_backup_enabled"`
	EnableAuthentication           bool   `json:"enable_authentication"`
	MaxmemoryReserved              int    `json:"maxmemory_reserved"`
	RdbBackupMaxSnapshotCount      int    `json:"rdb_backup_max_snapshot_count"`
}

type AzurermRedisCacheSpecPatchSchedule struct {
	DayOfWeek    string `json:"day_of_week"`
	StartHourUtc int    `json:"start_hour_utc"`
}

type AzurermRedisCacheSpec struct {
	Name                   string                  `json:"name"`
	Zones                  []string                `json:"zones"`
	Capacity               int                     `json:"capacity"`
	PrivateStaticIpAddress string                  `json:"private_static_ip_address"`
	SecondaryAccessKey     string                  `json:"secondary_access_key"`
	SubnetId               string                  `json:"subnet_id"`
	RedisConfiguration     []AzurermRedisCacheSpec `json:"redis_configuration"`
	Hostname               string                  `json:"hostname"`
	MinimumTlsVersion      string                  `json:"minimum_tls_version"`
	Port                   int                     `json:"port"`
	Tags                   map[string]string       `json:"tags"`
	SslPort                int                     `json:"ssl_port"`
	Location               string                  `json:"location"`
	ResourceGroupName      string                  `json:"resource_group_name"`
	Family                 string                  `json:"family"`
	SkuName                string                  `json:"sku_name"`
	ShardCount             int                     `json:"shard_count"`
	EnableNonSslPort       bool                    `json:"enable_non_ssl_port"`
	PatchSchedule          []AzurermRedisCacheSpec `json:"patch_schedule"`
	PrimaryAccessKey       string                  `json:"primary_access_key"`
}



type AzurermRedisCacheStatus struct {
    Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AzurermRedisCacheList is a list of AzurermRedisCaches
type AzurermRedisCacheList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AzurermRedisCache CRD objects
	Items []AzurermRedisCache `json:"items,omitempty"`
}