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

type GoogleRedisInstance struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              GoogleRedisInstanceSpec   `json:"spec,omitempty"`
	Status            GoogleRedisInstanceStatus `json:"status,omitempty"`
}

type GoogleRedisInstanceSpec struct {
	Host                  string            `json:"host"`
	Port                  int               `json:"port"`
	AuthorizedNetwork     string            `json:"authorized_network"`
	Labels                map[string]string `json:"labels"`
	Tier                  string            `json:"tier"`
	CurrentLocationId     string            `json:"current_location_id"`
	AlternativeLocationId string            `json:"alternative_location_id"`
	LocationId            string            `json:"location_id"`
	RedisConfigs          map[string]string `json:"redis_configs"`
	CreateTime            string            `json:"create_time"`
	DisplayName           string            `json:"display_name"`
	Region                string            `json:"region"`
	ReservedIpRange       string            `json:"reserved_ip_range"`
	MemorySizeGb          int               `json:"memory_size_gb"`
	Name                  string            `json:"name"`
	RedisVersion          string            `json:"redis_version"`
	Project               string            `json:"project"`
}



type GoogleRedisInstanceStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// GoogleRedisInstanceList is a list of GoogleRedisInstances
type GoogleRedisInstanceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of GoogleRedisInstance CRD objects
	Items []GoogleRedisInstance `json:"items,omitempty"`
}