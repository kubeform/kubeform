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

type GoogleComputeInstanceGroupManagerSpecRollingUpdatePolicy struct {
	MinReadySec           int    `json:"min_ready_sec"`
	MinimalAction         string `json:"minimal_action"`
	Type                  string `json:"type"`
	MaxSurgeFixed         int    `json:"max_surge_fixed"`
	MaxSurgePercent       int    `json:"max_surge_percent"`
	MaxUnavailableFixed   int    `json:"max_unavailable_fixed"`
	MaxUnavailablePercent int    `json:"max_unavailable_percent"`
}

type GoogleComputeInstanceGroupManagerSpec struct {
	BaseInstanceName    string                                  `json:"base_instance_name"`
	SelfLink            string                                  `json:"self_link"`
	UpdateStrategy      string                                  `json:"update_strategy"`
	AutoHealingPolicies []GoogleComputeInstanceGroupManagerSpec `json:"auto_healing_policies"`
	Version             []GoogleComputeInstanceGroupManagerSpec `json:"version"`
	Name                string                                  `json:"name"`
	Fingerprint         string                                  `json:"fingerprint"`
	TargetPools         []string                                `json:"target_pools"`
	WaitForInstances    bool                                    `json:"wait_for_instances"`
	InstanceTemplate    string                                  `json:"instance_template"`
	InstanceGroup       string                                  `json:"instance_group"`
	Project             string                                  `json:"project"`
	Zone                string                                  `json:"zone"`
	Description         string                                  `json:"description"`
	NamedPort           []GoogleComputeInstanceGroupManagerSpec `json:"named_port"`
	TargetSize          int                                     `json:"target_size"`
	RollingUpdatePolicy []GoogleComputeInstanceGroupManagerSpec `json:"rolling_update_policy"`
}

type GoogleComputeInstanceGroupManagerStatus struct {
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
