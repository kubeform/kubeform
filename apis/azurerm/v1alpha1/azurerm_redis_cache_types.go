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
	MaxmemoryReserved              int    `json:"maxmemory_reserved"`
	MaxfragmentationmemoryReserved int    `json:"maxfragmentationmemory_reserved"`
	RdbStorageConnectionString     string `json:"rdb_storage_connection_string"`
	AofBackupEnabled               bool   `json:"aof_backup_enabled"`
	MaxmemoryDelta                 int    `json:"maxmemory_delta"`
	RdbBackupFrequency             int    `json:"rdb_backup_frequency"`
	RdbBackupMaxSnapshotCount      int    `json:"rdb_backup_max_snapshot_count"`
	NotifyKeyspaceEvents           string `json:"notify_keyspace_events"`
	AofStorageConnectionString0    string `json:"aof_storage_connection_string_0"`
	MaxmemoryPolicy                string `json:"maxmemory_policy"`
	EnableAuthentication           bool   `json:"enable_authentication"`
	RdbBackupEnabled               bool   `json:"rdb_backup_enabled"`
	AofStorageConnectionString1    string `json:"aof_storage_connection_string_1"`
	Maxclients                     int    `json:"maxclients"`
}

type AzurermRedisCacheSpecPatchSchedule struct {
	DayOfWeek    string `json:"day_of_week"`
	StartHourUtc int    `json:"start_hour_utc"`
}

type AzurermRedisCacheSpec struct {
	SubnetId               string                  `json:"subnet_id"`
	PrivateStaticIpAddress string                  `json:"private_static_ip_address"`
	PrimaryAccessKey       string                  `json:"primary_access_key"`
	Name                   string                  `json:"name"`
	Zones                  []string                `json:"zones"`
	SkuName                string                  `json:"sku_name"`
	RedisConfiguration     []AzurermRedisCacheSpec `json:"redis_configuration"`
	ResourceGroupName      string                  `json:"resource_group_name"`
	MinimumTlsVersion      string                  `json:"minimum_tls_version"`
	ShardCount             int                     `json:"shard_count"`
	PatchSchedule          []AzurermRedisCacheSpec `json:"patch_schedule"`
	Port                   int                     `json:"port"`
	SslPort                int                     `json:"ssl_port"`
	SecondaryAccessKey     string                  `json:"secondary_access_key"`
	Location               string                  `json:"location"`
	Family                 string                  `json:"family"`
	EnableNonSslPort       bool                    `json:"enable_non_ssl_port"`
	Hostname               string                  `json:"hostname"`
	Tags                   map[string]string       `json:"tags"`
	Capacity               int                     `json:"capacity"`
}



type AzurermRedisCacheStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

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