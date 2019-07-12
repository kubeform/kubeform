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
	RedisVersion          string            `json:"redis_version"`
	Project               string            `json:"project"`
	MemorySizeGb          int               `json:"memory_size_gb"`
	Name                  string            `json:"name"`
	Region                string            `json:"region"`
	AlternativeLocationId string            `json:"alternative_location_id"`
	AuthorizedNetwork     string            `json:"authorized_network"`
	DisplayName           string            `json:"display_name"`
	RedisConfigs          map[string]string `json:"redis_configs"`
	ReservedIpRange       string            `json:"reserved_ip_range"`
	Tier                  string            `json:"tier"`
	Labels                map[string]string `json:"labels"`
	LocationId            string            `json:"location_id"`
	CreateTime            string            `json:"create_time"`
	CurrentLocationId     string            `json:"current_location_id"`
	Host                  string            `json:"host"`
	Port                  int               `json:"port"`
}

type GoogleRedisInstanceStatus struct {
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
