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

type RedisCache struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              RedisCacheSpec   `json:"spec,omitempty"`
	Status            RedisCacheStatus `json:"status,omitempty"`
}

type RedisCacheSpecPatchSchedule struct {
	DayOfWeek string `json:"dayOfWeek" tf:"day_of_week"`
	// +optional
	StartHourUtc int `json:"startHourUtc,omitempty" tf:"start_hour_utc,omitempty"`
}

type RedisCacheSpecRedisConfiguration struct {
	// +optional
	AofBackupEnabled bool `json:"aofBackupEnabled,omitempty" tf:"aof_backup_enabled,omitempty"`
	// +optional
	// Sensitive Data. Provide secret name which contains one value only
	AofStorageConnectionString0 *core.LocalObjectReference `json:"aofStorageConnectionString0,omitempty" tf:"aof_storage_connection_string_0,omitempty"`
	// +optional
	// Sensitive Data. Provide secret name which contains one value only
	AofStorageConnectionString1 *core.LocalObjectReference `json:"aofStorageConnectionString1,omitempty" tf:"aof_storage_connection_string_1,omitempty"`
	// +optional
	EnableAuthentication bool `json:"enableAuthentication,omitempty" tf:"enable_authentication,omitempty"`
	// +optional
	MaxfragmentationmemoryReserved int `json:"maxfragmentationmemoryReserved,omitempty" tf:"maxfragmentationmemory_reserved,omitempty"`
	// +optional
	MaxmemoryDelta int `json:"maxmemoryDelta,omitempty" tf:"maxmemory_delta,omitempty"`
	// +optional
	MaxmemoryPolicy string `json:"maxmemoryPolicy,omitempty" tf:"maxmemory_policy,omitempty"`
	// +optional
	MaxmemoryReserved int `json:"maxmemoryReserved,omitempty" tf:"maxmemory_reserved,omitempty"`
	// +optional
	NotifyKeyspaceEvents string `json:"notifyKeyspaceEvents,omitempty" tf:"notify_keyspace_events,omitempty"`
	// +optional
	RdbBackupEnabled bool `json:"rdbBackupEnabled,omitempty" tf:"rdb_backup_enabled,omitempty"`
	// +optional
	RdbBackupFrequency int `json:"rdbBackupFrequency,omitempty" tf:"rdb_backup_frequency,omitempty"`
	// +optional
	RdbBackupMaxSnapshotCount int `json:"rdbBackupMaxSnapshotCount,omitempty" tf:"rdb_backup_max_snapshot_count,omitempty"`
	// +optional
	// Sensitive Data. Provide secret name which contains one value only
	RdbStorageConnectionString *core.LocalObjectReference `json:"rdbStorageConnectionString,omitempty" tf:"rdb_storage_connection_string,omitempty"`
}

type RedisCacheSpec struct {
	Capacity int `json:"capacity" tf:"capacity"`
	// +optional
	EnableNonSslPort bool   `json:"enableNonSslPort,omitempty" tf:"enable_non_ssl_port,omitempty"`
	Family           string `json:"family" tf:"family"`
	Location         string `json:"location" tf:"location"`
	// +optional
	MinimumTlsVersion string `json:"minimumTlsVersion,omitempty" tf:"minimum_tls_version,omitempty"`
	Name              string `json:"name" tf:"name"`
	// +optional
	PatchSchedule []RedisCacheSpecPatchSchedule `json:"patchSchedule,omitempty" tf:"patch_schedule,omitempty"`
	// +optional
	PrivateStaticIPAddress string `json:"privateStaticIPAddress,omitempty" tf:"private_static_ip_address,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	RedisConfiguration []RedisCacheSpecRedisConfiguration `json:"redisConfiguration,omitempty" tf:"redis_configuration,omitempty"`
	ResourceGroupName  string                             `json:"resourceGroupName" tf:"resource_group_name"`
	// +optional
	ShardCount int    `json:"shardCount,omitempty" tf:"shard_count,omitempty"`
	SkuName    string `json:"skuName" tf:"sku_name"`
	// +optional
	SubnetID string `json:"subnetID,omitempty" tf:"subnet_id,omitempty"`
	// +optional
	Tags map[string]string `json:"tags,omitempty" tf:"tags,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	Zones       []string                  `json:"zones,omitempty" tf:"zones,omitempty"`
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`
}

type RedisCacheStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	TFState *runtime.RawExtension `json:"tfState,omitempty"`
	Output  *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// RedisCacheList is a list of RedisCaches
type RedisCacheList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of RedisCache CRD objects
	Items []RedisCache `json:"items,omitempty"`
}
