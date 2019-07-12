package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsBatchComputeEnvironment struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsBatchComputeEnvironmentSpec   `json:"spec,omitempty"`
	Status            AwsBatchComputeEnvironmentStatus `json:"status,omitempty"`
}

type AwsBatchComputeEnvironmentSpecComputeResourcesLaunchTemplate struct {
	LaunchTemplateId   string `json:"launch_template_id"`
	LaunchTemplateName string `json:"launch_template_name"`
	Version            string `json:"version"`
}

type AwsBatchComputeEnvironmentSpecComputeResources struct {
	SpotIamFleetRole string                                           `json:"spot_iam_fleet_role"`
	Ec2KeyPair       string                                           `json:"ec2_key_pair"`
	MinVcpus         int                                              `json:"min_vcpus"`
	Type             string                                           `json:"type"`
	InstanceType     []string                                         `json:"instance_type"`
	LaunchTemplate   []AwsBatchComputeEnvironmentSpecComputeResources `json:"launch_template"`
	Tags             map[string]string                                `json:"tags"`
	BidPercentage    int                                              `json:"bid_percentage"`
	ImageId          string                                           `json:"image_id"`
	MaxVcpus         int                                              `json:"max_vcpus"`
	SecurityGroupIds []string                                         `json:"security_group_ids"`
	Subnets          []string                                         `json:"subnets"`
	DesiredVcpus     int                                              `json:"desired_vcpus"`
	InstanceRole     string                                           `json:"instance_role"`
}

type AwsBatchComputeEnvironmentSpec struct {
	EccClusterArn          string                           `json:"ecc_cluster_arn"`
	EcsClusterArn          string                           `json:"ecs_cluster_arn"`
	ComputeResources       []AwsBatchComputeEnvironmentSpec `json:"compute_resources"`
	ServiceRole            string                           `json:"service_role"`
	State                  string                           `json:"state"`
	Arn                    string                           `json:"arn"`
	ComputeEnvironmentName string                           `json:"compute_environment_name"`
	Type                   string                           `json:"type"`
	Status                 string                           `json:"status"`
	StatusReason           string                           `json:"status_reason"`
}

type AwsBatchComputeEnvironmentStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsBatchComputeEnvironmentList is a list of AwsBatchComputeEnvironments
type AwsBatchComputeEnvironmentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsBatchComputeEnvironment CRD objects
	Items []AwsBatchComputeEnvironment `json:"items,omitempty"`
}
