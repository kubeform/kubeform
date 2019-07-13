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

type GoogleComputeInstanceGroupManager struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              GoogleComputeInstanceGroupManagerSpec   `json:"spec,omitempty"`
	Status            GoogleComputeInstanceGroupManagerStatus `json:"status,omitempty"`
}

type GoogleComputeInstanceGroupManagerSpecNamedPort struct {
	Name string `json:"name"`
	Port int    `json:"port"`
}

type GoogleComputeInstanceGroupManagerSpecVersionTargetSize struct {
	Fixed   int `json:"fixed"`
	Percent int `json:"percent"`
}

type GoogleComputeInstanceGroupManagerSpecVersion struct {
	Name             string                                         `json:"name"`
	InstanceTemplate string                                         `json:"instance_template"`
	TargetSize       []GoogleComputeInstanceGroupManagerSpecVersion `json:"target_size"`
}

type GoogleComputeInstanceGroupManagerSpecRollingUpdatePolicy struct {
	MaxSurgePercent       int    `json:"max_surge_percent"`
	MaxUnavailableFixed   int    `json:"max_unavailable_fixed"`
	MaxUnavailablePercent int    `json:"max_unavailable_percent"`
	MinReadySec           int    `json:"min_ready_sec"`
	MinimalAction         string `json:"minimal_action"`
	Type                  string `json:"type"`
	MaxSurgeFixed         int    `json:"max_surge_fixed"`
}

type GoogleComputeInstanceGroupManagerSpecAutoHealingPolicies struct {
	HealthCheck     string `json:"health_check"`
	InitialDelaySec int    `json:"initial_delay_sec"`
}

type GoogleComputeInstanceGroupManagerSpec struct {
	BaseInstanceName    string                                  `json:"base_instance_name"`
	Fingerprint         string                                  `json:"fingerprint"`
	NamedPort           []GoogleComputeInstanceGroupManagerSpec `json:"named_port"`
	TargetSize          int                                     `json:"target_size"`
	UpdateStrategy      string                                  `json:"update_strategy"`
	TargetPools         []string                                `json:"target_pools"`
	WaitForInstances    bool                                    `json:"wait_for_instances"`
	InstanceTemplate    string                                  `json:"instance_template"`
	Version             []GoogleComputeInstanceGroupManagerSpec `json:"version"`
	InstanceGroup       string                                  `json:"instance_group"`
	SelfLink            string                                  `json:"self_link"`
	Zone                string                                  `json:"zone"`
	RollingUpdatePolicy []GoogleComputeInstanceGroupManagerSpec `json:"rolling_update_policy"`
	Name                string                                  `json:"name"`
	Description         string                                  `json:"description"`
	Project             string                                  `json:"project"`
	AutoHealingPolicies []GoogleComputeInstanceGroupManagerSpec `json:"auto_healing_policies"`
}



type GoogleComputeInstanceGroupManagerStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// GoogleComputeInstanceGroupManagerList is a list of GoogleComputeInstanceGroupManagers
type GoogleComputeInstanceGroupManagerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of GoogleComputeInstanceGroupManager CRD objects
	Items []GoogleComputeInstanceGroupManager `json:"items,omitempty"`
}