package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsGameliftFleet struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsGameliftFleetSpec   `json:"spec,omitempty"`
	Status            AwsGameliftFleetStatus `json:"status,omitempty"`
}

type AwsGameliftFleetSpecEc2InboundPermission struct {
	FromPort int    `json:"from_port"`
	IpRange  string `json:"ip_range"`
	Protocol string `json:"protocol"`
	ToPort   int    `json:"to_port"`
}

type AwsGameliftFleetSpecResourceCreationLimitPolicy struct {
	PolicyPeriodInMinutes     int `json:"policy_period_in_minutes"`
	NewGameSessionsPerCreator int `json:"new_game_sessions_per_creator"`
}

type AwsGameliftFleetSpecRuntimeConfigurationServerProcess struct {
	ConcurrentExecutions int    `json:"concurrent_executions"`
	LaunchPath           string `json:"launch_path"`
	Parameters           string `json:"parameters"`
}

type AwsGameliftFleetSpecRuntimeConfiguration struct {
	ServerProcess                       []AwsGameliftFleetSpecRuntimeConfiguration `json:"server_process"`
	GameSessionActivationTimeoutSeconds int                                        `json:"game_session_activation_timeout_seconds"`
	MaxConcurrentGameSessionActivations int                                        `json:"max_concurrent_game_session_activations"`
}

type AwsGameliftFleetSpec struct {
	Ec2InstanceType                string                 `json:"ec2_instance_type"`
	Description                    string                 `json:"description"`
	Ec2InboundPermission           []AwsGameliftFleetSpec `json:"ec2_inbound_permission"`
	MetricGroups                   []string               `json:"metric_groups"`
	ResourceCreationLimitPolicy    []AwsGameliftFleetSpec `json:"resource_creation_limit_policy"`
	RuntimeConfiguration           []AwsGameliftFleetSpec `json:"runtime_configuration"`
	Arn                            string                 `json:"arn"`
	BuildId                        string                 `json:"build_id"`
	Name                           string                 `json:"name"`
	LogPaths                       []string               `json:"log_paths"`
	NewGameSessionProtectionPolicy string                 `json:"new_game_session_protection_policy"`
	OperatingSystem                string                 `json:"operating_system"`
}



type AwsGameliftFleetStatus struct {
    Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsGameliftFleetList is a list of AwsGameliftFleets
type AwsGameliftFleetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsGameliftFleet CRD objects
	Items []AwsGameliftFleet `json:"items,omitempty"`
}