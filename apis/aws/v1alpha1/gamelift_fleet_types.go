package v1alpha1

import (
	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	base "kubeform.dev/kubeform/apis/base/v1alpha1"
)

// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

type GameliftFleet struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              GameliftFleetSpec   `json:"spec,omitempty"`
	Status            GameliftFleetStatus `json:"status,omitempty"`
}

type GameliftFleetSpecEc2InboundPermission struct {
	FromPort int    `json:"fromPort" tf:"from_port"`
	IpRange  string `json:"ipRange" tf:"ip_range"`
	Protocol string `json:"protocol" tf:"protocol"`
	ToPort   int    `json:"toPort" tf:"to_port"`
}

type GameliftFleetSpecResourceCreationLimitPolicy struct {
	// +optional
	NewGameSessionsPerCreator int `json:"newGameSessionsPerCreator,omitempty" tf:"new_game_sessions_per_creator,omitempty"`
	// +optional
	PolicyPeriodInMinutes int `json:"policyPeriodInMinutes,omitempty" tf:"policy_period_in_minutes,omitempty"`
}

type GameliftFleetSpecRuntimeConfigurationServerProcess struct {
	ConcurrentExecutions int    `json:"concurrentExecutions" tf:"concurrent_executions"`
	LaunchPath           string `json:"launchPath" tf:"launch_path"`
	// +optional
	Parameters string `json:"parameters,omitempty" tf:"parameters,omitempty"`
}

type GameliftFleetSpecRuntimeConfiguration struct {
	// +optional
	GameSessionActivationTimeoutSeconds int `json:"gameSessionActivationTimeoutSeconds,omitempty" tf:"game_session_activation_timeout_seconds,omitempty"`
	// +optional
	MaxConcurrentGameSessionActivations int `json:"maxConcurrentGameSessionActivations,omitempty" tf:"max_concurrent_game_session_activations,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=50
	ServerProcess []GameliftFleetSpecRuntimeConfigurationServerProcess `json:"serverProcess,omitempty" tf:"server_process,omitempty"`
}

type GameliftFleetSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	Arn     string `json:"arn,omitempty" tf:"arn,omitempty"`
	BuildID string `json:"buildID" tf:"build_id"`
	// +optional
	Description string `json:"description,omitempty" tf:"description,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=50
	Ec2InboundPermission []GameliftFleetSpecEc2InboundPermission `json:"ec2InboundPermission,omitempty" tf:"ec2_inbound_permission,omitempty"`
	Ec2InstanceType      string                                  `json:"ec2InstanceType" tf:"ec2_instance_type"`
	// +optional
	LogPaths []string `json:"logPaths,omitempty" tf:"log_paths,omitempty"`
	// +optional
	MetricGroups []string `json:"metricGroups,omitempty" tf:"metric_groups,omitempty"`
	Name         string   `json:"name" tf:"name"`
	// +optional
	NewGameSessionProtectionPolicy string `json:"newGameSessionProtectionPolicy,omitempty" tf:"new_game_session_protection_policy,omitempty"`
	// +optional
	OperatingSystem string `json:"operatingSystem,omitempty" tf:"operating_system,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	ResourceCreationLimitPolicy []GameliftFleetSpecResourceCreationLimitPolicy `json:"resourceCreationLimitPolicy,omitempty" tf:"resource_creation_limit_policy,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	RuntimeConfiguration []GameliftFleetSpecRuntimeConfiguration `json:"runtimeConfiguration,omitempty" tf:"runtime_configuration,omitempty"`
}

type GameliftFleetStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *GameliftFleetSpec `json:"output,omitempty"`
	// +optional
	State *base.State `json:"state,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// GameliftFleetList is a list of GameliftFleets
type GameliftFleetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of GameliftFleet CRD objects
	Items []GameliftFleet `json:"items,omitempty"`
}
