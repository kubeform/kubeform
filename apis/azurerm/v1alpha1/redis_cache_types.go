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

type RedisCache struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              RedisCacheSpec   `json:"spec,omitempty"`
	Status            RedisCacheStatus `json:"status,omitempty"`
}

type RedisCacheSpecPatchSchedule struct {
	DayOfWeek string `json:"day_of_week"`
	// +optional
	StartHourUtc int `json:"start_hour_utc,omitempty"`
}

type RedisCacheSpec struct {
	Capacity int `json:"capacity"`
	// +optional
	EnableNonSslPort bool   `json:"enable_non_ssl_port,omitempty"`
	Family           string `json:"family"`
	Location         string `json:"location"`
	// +optional
	MinimumTlsVersion string `json:"minimum_tls_version,omitempty"`
	Name              string `json:"name"`
	// +optional
	PatchSchedule     *[]RedisCacheSpec `json:"patch_schedule,omitempty"`
	ResourceGroupName string            `json:"resource_group_name"`
	// +optional
	ShardCount int    `json:"shard_count,omitempty"`
	SkuName    string `json:"sku_name"`
	// +optional
	SubnetId string `json:"subnet_id,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	Zones []string `json:"zones,omitempty"`
}

type RedisCacheStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
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
