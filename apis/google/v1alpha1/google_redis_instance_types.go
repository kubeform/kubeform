package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type GoogleRedisInstance struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              GoogleRedisInstanceSpec   `json:"spec,omitempty"`
	Status            GoogleRedisInstanceStatus `json:"status,omitempty"`
}

type GoogleRedisInstanceSpec struct {
	AlternativeLocationId string            `json:"alternative_location_id"`
	Labels                map[string]string `json:"labels"`
	ReservedIpRange       string            `json:"reserved_ip_range"`
	Host                  string            `json:"host"`
	DisplayName           string            `json:"display_name"`
	LocationId            string            `json:"location_id"`
	Region                string            `json:"region"`
	CurrentLocationId     string            `json:"current_location_id"`
	Port                  int               `json:"port"`
	Project               string            `json:"project"`
	MemorySizeGb          int               `json:"memory_size_gb"`
	Name                  string            `json:"name"`
	AuthorizedNetwork     string            `json:"authorized_network"`
	RedisConfigs          map[string]string `json:"redis_configs"`
	RedisVersion          string            `json:"redis_version"`
	Tier                  string            `json:"tier"`
	CreateTime            string            `json:"create_time"`
}



type GoogleRedisInstanceStatus struct {
    Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// GoogleRedisInstanceList is a list of GoogleRedisInstances
type GoogleRedisInstanceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of GoogleRedisInstance CRD objects
	Items []GoogleRedisInstance `json:"items,omitempty"`
}