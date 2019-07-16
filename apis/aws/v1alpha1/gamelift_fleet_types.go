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

type GameliftFleet struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              GameliftFleetSpec   `json:"spec,omitempty"`
	Status            GameliftFleetStatus `json:"status,omitempty"`
}

type GameliftFleetSpecEc2InboundPermission struct {
	FromPort int    `json:"from_port"`
	IpRange  string `json:"ip_range"`
	Protocol string `json:"protocol"`
	ToPort   int    `json:"to_port"`
}

type GameliftFleetSpecResourceCreationLimitPolicy struct {
	// +optional
	NewGameSessionsPerCreator int `json:"new_game_sessions_per_creator,omitempty"`
	// +optional
	PolicyPeriodInMinutes int `json:"policy_period_in_minutes,omitempty"`
}

type GameliftFleetSpecRuntimeConfigurationServerProcess struct {
	ConcurrentExecutions int    `json:"concurrent_executions"`
	LaunchPath           string `json:"launch_path"`
	// +optional
	Parameters string `json:"parameters,omitempty"`
}

type GameliftFleetSpecRuntimeConfiguration struct {
	// +optional
	GameSessionActivationTimeoutSeconds int `json:"game_session_activation_timeout_seconds,omitempty"`
	// +optional
	MaxConcurrentGameSessionActivations int `json:"max_concurrent_game_session_activations,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=50
	ServerProcess *[]GameliftFleetSpecRuntimeConfiguration `json:"server_process,omitempty"`
}

type GameliftFleetSpec struct {
	BuildId string `json:"build_id"`
	// +optional
	Description string `json:"description,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=50
	Ec2InboundPermission *[]GameliftFleetSpec `json:"ec2_inbound_permission,omitempty"`
	Ec2InstanceType      string               `json:"ec2_instance_type"`
	Name                 string               `json:"name"`
	// +optional
	NewGameSessionProtectionPolicy string `json:"new_game_session_protection_policy,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	ResourceCreationLimitPolicy *[]GameliftFleetSpec `json:"resource_creation_limit_policy,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	RuntimeConfiguration *[]GameliftFleetSpec `json:"runtime_configuration,omitempty"`
}

type GameliftFleetStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
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
