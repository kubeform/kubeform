package v1alpha1

import (
	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"kubeform.dev/kubeform/apis"
)

// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

type BatchComputeEnvironment struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              BatchComputeEnvironmentSpec   `json:"spec,omitempty"`
	Status            BatchComputeEnvironmentStatus `json:"status,omitempty"`
}

type BatchComputeEnvironmentSpecComputeResourcesLaunchTemplate struct {
	// +optional
	LaunchTemplateID string `json:"launchTemplateID,omitempty" tf:"launch_template_id,omitempty"`
	// +optional
	LaunchTemplateName string `json:"launchTemplateName,omitempty" tf:"launch_template_name,omitempty"`
	// +optional
	Version string `json:"version,omitempty" tf:"version,omitempty"`
}

type BatchComputeEnvironmentSpecComputeResources struct {
	// +optional
	BidPercentage int `json:"bidPercentage,omitempty" tf:"bid_percentage,omitempty"`
	// +optional
	DesiredVcpus int `json:"desiredVcpus,omitempty" tf:"desired_vcpus,omitempty"`
	// +optional
	Ec2KeyPair string `json:"ec2KeyPair,omitempty" tf:"ec2_key_pair,omitempty"`
	// +optional
	ImageID      string `json:"imageID,omitempty" tf:"image_id,omitempty"`
	InstanceRole string `json:"instanceRole" tf:"instance_role"`
	// +kubebuilder:validation:UniqueItems=true
	InstanceType []string `json:"instanceType" tf:"instance_type"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	LaunchTemplate []BatchComputeEnvironmentSpecComputeResourcesLaunchTemplate `json:"launchTemplate,omitempty" tf:"launch_template,omitempty"`
	MaxVcpus       int                                                         `json:"maxVcpus" tf:"max_vcpus"`
	MinVcpus       int                                                         `json:"minVcpus" tf:"min_vcpus"`
	// +kubebuilder:validation:UniqueItems=true
	SecurityGroupIDS []string `json:"securityGroupIDS" tf:"security_group_ids"`
	// +optional
	SpotIamFleetRole string `json:"spotIamFleetRole,omitempty" tf:"spot_iam_fleet_role,omitempty"`
	// +kubebuilder:validation:UniqueItems=true
	Subnets []string `json:"subnets" tf:"subnets"`
	// +optional
	Tags map[string]string `json:"tags,omitempty" tf:"tags,omitempty"`
	Type string            `json:"type" tf:"type"`
}

type BatchComputeEnvironmentSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	// +optional
	Arn                    string `json:"arn,omitempty" tf:"arn,omitempty"`
	ComputeEnvironmentName string `json:"computeEnvironmentName" tf:"compute_environment_name"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	ComputeResources []BatchComputeEnvironmentSpecComputeResources `json:"computeResources,omitempty" tf:"compute_resources,omitempty"`
	// +optional
	EcsClusterArn string `json:"ecsClusterArn,omitempty" tf:"ecs_cluster_arn,omitempty"`
	ServiceRole   string `json:"serviceRole" tf:"service_role"`
	// +optional
	State string `json:"state,omitempty" tf:"state,omitempty"`
	// +optional
	Status string `json:"status,omitempty" tf:"status,omitempty"`
	// +optional
	StatusReason string `json:"statusReason,omitempty" tf:"status_reason,omitempty"`
	Type         string `json:"type" tf:"type"`
}

type BatchComputeEnvironmentStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *BatchComputeEnvironmentSpec `json:"output,omitempty"`
	// +optional
	State *apis.State `json:"state,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// BatchComputeEnvironmentList is a list of BatchComputeEnvironments
type BatchComputeEnvironmentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of BatchComputeEnvironment CRD objects
	Items []BatchComputeEnvironment `json:"items,omitempty"`
}
