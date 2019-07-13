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

type AwsGameliftFleet struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsGameliftFleetSpec   `json:"spec,omitempty"`
	Status            AwsGameliftFleetStatus `json:"status,omitempty"`
}

type AwsGameliftFleetSpecRuntimeConfigurationServerProcess struct {
	ConcurrentExecutions int    `json:"concurrent_executions"`
	LaunchPath           string `json:"launch_path"`
	Parameters           string `json:"parameters"`
}

type AwsGameliftFleetSpecRuntimeConfiguration struct {
	GameSessionActivationTimeoutSeconds int                                        `json:"game_session_activation_timeout_seconds"`
	MaxConcurrentGameSessionActivations int                                        `json:"max_concurrent_game_session_activations"`
	ServerProcess                       []AwsGameliftFleetSpecRuntimeConfiguration `json:"server_process"`
}

type AwsGameliftFleetSpecEc2InboundPermission struct {
	FromPort int    `json:"from_port"`
	IpRange  string `json:"ip_range"`
	Protocol string `json:"protocol"`
	ToPort   int    `json:"to_port"`
}

type AwsGameliftFleetSpecResourceCreationLimitPolicy struct {
	NewGameSessionsPerCreator int `json:"new_game_sessions_per_creator"`
	PolicyPeriodInMinutes     int `json:"policy_period_in_minutes"`
}

type AwsGameliftFleetSpec struct {
	RuntimeConfiguration           []AwsGameliftFleetSpec `json:"runtime_configuration"`
	Arn                            string                 `json:"arn"`
	BuildId                        string                 `json:"build_id"`
	Ec2InstanceType                string                 `json:"ec2_instance_type"`
	Name                           string                 `json:"name"`
	Ec2InboundPermission           []AwsGameliftFleetSpec `json:"ec2_inbound_permission"`
	ResourceCreationLimitPolicy    []AwsGameliftFleetSpec `json:"resource_creation_limit_policy"`
	Description                    string                 `json:"description"`
	LogPaths                       []string               `json:"log_paths"`
	MetricGroups                   []string               `json:"metric_groups"`
	NewGameSessionProtectionPolicy string                 `json:"new_game_session_protection_policy"`
	OperatingSystem                string                 `json:"operating_system"`
}



type AwsGameliftFleetStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AwsGameliftFleetList is a list of AwsGameliftFleets
type AwsGameliftFleetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsGameliftFleet CRD objects
	Items []AwsGameliftFleet `json:"items,omitempty"`
}