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

type RedisInstance struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              RedisInstanceSpec   `json:"spec,omitempty"`
	Status            RedisInstanceStatus `json:"status,omitempty"`
}

type RedisInstanceSpec struct {
	// +optional
	DisplayName string `json:"display_name,omitempty"`
	// +optional
	Labels       map[string]string `json:"labels,omitempty"`
	MemorySizeGb int               `json:"memory_size_gb"`
	Name         string            `json:"name"`
	// +optional
	RedisConfigs map[string]string `json:"redis_configs,omitempty"`
	// +optional
	Tier string `json:"tier,omitempty"`
}

type RedisInstanceStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// RedisInstanceList is a list of RedisInstances
type RedisInstanceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of RedisInstance CRD objects
	Items []RedisInstance `json:"items,omitempty"`
}
