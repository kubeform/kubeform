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
	Ec2KeyPair       string                                           `json:"ec2_key_pair"`
	Subnets          []string                                         `json:"subnets"`
	Type             string                                           `json:"type"`
	ImageId          string                                           `json:"image_id"`
	LaunchTemplate   []AwsBatchComputeEnvironmentSpecComputeResources `json:"launch_template"`
	MinVcpus         int                                              `json:"min_vcpus"`
	SecurityGroupIds []string                                         `json:"security_group_ids"`
	BidPercentage    int                                              `json:"bid_percentage"`
	InstanceRole     string                                           `json:"instance_role"`
	MaxVcpus         int                                              `json:"max_vcpus"`
	SpotIamFleetRole string                                           `json:"spot_iam_fleet_role"`
	DesiredVcpus     int                                              `json:"desired_vcpus"`
	InstanceType     []string                                         `json:"instance_type"`
	Tags             map[string]string                                `json:"tags"`
}

type AwsBatchComputeEnvironmentSpec struct {
	ComputeResources       []AwsBatchComputeEnvironmentSpec `json:"compute_resources"`
	ServiceRole            string                           `json:"service_role"`
	State                  string                           `json:"state"`
	Type                   string                           `json:"type"`
	EccClusterArn          string                           `json:"ecc_cluster_arn"`
	StatusReason           string                           `json:"status_reason"`
	ComputeEnvironmentName string                           `json:"compute_environment_name"`
	Arn                    string                           `json:"arn"`
	EcsClusterArn          string                           `json:"ecs_cluster_arn"`
	Status                 string                           `json:"status"`
}

type AwsBatchComputeEnvironmentStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AwsBatchComputeEnvironmentList is a list of AwsBatchComputeEnvironments
type AwsBatchComputeEnvironmentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsBatchComputeEnvironment CRD objects
	Items []AwsBatchComputeEnvironment `json:"items,omitempty"`
}
