package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type GoogleComputeInstanceGroupManager struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              GoogleComputeInstanceGroupManagerSpec   `json:"spec,omitempty"`
	Status            GoogleComputeInstanceGroupManagerStatus `json:"status,omitempty"`
}

type GoogleComputeInstanceGroupManagerSpecRollingUpdatePolicy struct {
	MaxSurgeFixed         int    `json:"max_surge_fixed"`
	MaxSurgePercent       int    `json:"max_surge_percent"`
	MaxUnavailableFixed   int    `json:"max_unavailable_fixed"`
	MaxUnavailablePercent int    `json:"max_unavailable_percent"`
	MinReadySec           int    `json:"min_ready_sec"`
	MinimalAction         string `json:"minimal_action"`
	Type                  string `json:"type"`
}

type GoogleComputeInstanceGroupManagerSpecAutoHealingPolicies struct {
	HealthCheck     string `json:"health_check"`
	InitialDelaySec int    `json:"initial_delay_sec"`
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

type GoogleComputeInstanceGroupManagerSpecNamedPort struct {
	Name string `json:"name"`
	Port int    `json:"port"`
}

type GoogleComputeInstanceGroupManagerSpec struct {
	BaseInstanceName    string                                  `json:"base_instance_name"`
	UpdateStrategy      string                                  `json:"update_strategy"`
	TargetSize          int                                     `json:"target_size"`
	RollingUpdatePolicy []GoogleComputeInstanceGroupManagerSpec `json:"rolling_update_policy"`
	WaitForInstances    bool                                    `json:"wait_for_instances"`
	SelfLink            string                                  `json:"self_link"`
	AutoHealingPolicies []GoogleComputeInstanceGroupManagerSpec `json:"auto_healing_policies"`
	InstanceTemplate    string                                  `json:"instance_template"`
	Zone                string                                  `json:"zone"`
	Description         string                                  `json:"description"`
	Fingerprint         string                                  `json:"fingerprint"`
	InstanceGroup       string                                  `json:"instance_group"`
	Version             []GoogleComputeInstanceGroupManagerSpec `json:"version"`
	Name                string                                  `json:"name"`
	NamedPort           []GoogleComputeInstanceGroupManagerSpec `json:"named_port"`
	Project             string                                  `json:"project"`
	TargetPools         []string                                `json:"target_pools"`
}

type GoogleComputeInstanceGroupManagerStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// GoogleComputeInstanceGroupManagerList is a list of GoogleComputeInstanceGroupManagers
type GoogleComputeInstanceGroupManagerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of GoogleComputeInstanceGroupManager CRD objects
	Items []GoogleComputeInstanceGroupManager `json:"items,omitempty"`
}
