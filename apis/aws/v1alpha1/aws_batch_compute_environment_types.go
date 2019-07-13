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
	Tags             map[string]string                                `json:"tags"`
	LaunchTemplate   []AwsBatchComputeEnvironmentSpecComputeResources `json:"launch_template"`
	MinVcpus         int                                              `json:"min_vcpus"`
	Type             string                                           `json:"type"`
	SpotIamFleetRole string                                           `json:"spot_iam_fleet_role"`
	BidPercentage    int                                              `json:"bid_percentage"`
	DesiredVcpus     int                                              `json:"desired_vcpus"`
	Ec2KeyPair       string                                           `json:"ec2_key_pair"`
	InstanceRole     string                                           `json:"instance_role"`
	InstanceType     []string                                         `json:"instance_type"`
	MaxVcpus         int                                              `json:"max_vcpus"`
	SecurityGroupIds []string                                         `json:"security_group_ids"`
	Subnets          []string                                         `json:"subnets"`
	ImageId          string                                           `json:"image_id"`
}

type AwsBatchComputeEnvironmentSpec struct {
	State                  string                           `json:"state"`
	StatusReason           string                           `json:"status_reason"`
	ComputeResources       []AwsBatchComputeEnvironmentSpec `json:"compute_resources"`
	ServiceRole            string                           `json:"service_role"`
	Arn                    string                           `json:"arn"`
	EccClusterArn          string                           `json:"ecc_cluster_arn"`
	EcsClusterArn          string                           `json:"ecs_cluster_arn"`
	Status                 string                           `json:"status"`
	ComputeEnvironmentName string                           `json:"compute_environment_name"`
	Type                   string                           `json:"type"`
}



type AwsBatchComputeEnvironmentStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

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