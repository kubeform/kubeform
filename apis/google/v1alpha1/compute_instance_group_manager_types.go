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

type ComputeInstanceGroupManager struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ComputeInstanceGroupManagerSpec   `json:"spec,omitempty"`
	Status            ComputeInstanceGroupManagerStatus `json:"status,omitempty"`
}

type ComputeInstanceGroupManagerSpecAutoHealingPolicies struct {
	HealthCheck     string `json:"health_check"`
	InitialDelaySec int    `json:"initial_delay_sec"`
}

type ComputeInstanceGroupManagerSpecNamedPort struct {
	Name string `json:"name"`
	Port int    `json:"port"`
}

type ComputeInstanceGroupManagerSpecRollingUpdatePolicy struct {
	// +optional
	MaxSurgeFixed int `json:"max_surge_fixed,omitempty"`
	// +optional
	MaxSurgePercent int `json:"max_surge_percent,omitempty"`
	// +optional
	MaxUnavailableFixed int `json:"max_unavailable_fixed,omitempty"`
	// +optional
	MaxUnavailablePercent int `json:"max_unavailable_percent,omitempty"`
	// +optional
	MinReadySec   int    `json:"min_ready_sec,omitempty"`
	MinimalAction string `json:"minimal_action"`
	Type          string `json:"type"`
}

type ComputeInstanceGroupManagerSpec struct {
	// +optional
	// +kubebuilder:validation:MaxItems=1
	// Deprecated
	AutoHealingPolicies *[]ComputeInstanceGroupManagerSpec `json:"auto_healing_policies,omitempty"`
	BaseInstanceName    string                             `json:"base_instance_name"`
	// +optional
	Description string `json:"description,omitempty"`
	// +optional
	InstanceTemplate string `json:"instance_template,omitempty"`
	Name             string `json:"name"`
	// +optional
	NamedPort *[]ComputeInstanceGroupManagerSpec `json:"named_port,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	// Deprecated
	RollingUpdatePolicy *[]ComputeInstanceGroupManagerSpec `json:"rolling_update_policy,omitempty"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	TargetPools []string `json:"target_pools,omitempty"`
	// +optional
	UpdateStrategy string `json:"update_strategy,omitempty"`
	// +optional
	WaitForInstances bool `json:"wait_for_instances,omitempty"`
}

type ComputeInstanceGroupManagerStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// ComputeInstanceGroupManagerList is a list of ComputeInstanceGroupManagers
type ComputeInstanceGroupManagerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of ComputeInstanceGroupManager CRD objects
	Items []ComputeInstanceGroupManager `json:"items,omitempty"`
}
