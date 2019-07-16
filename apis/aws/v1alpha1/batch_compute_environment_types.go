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

type BatchComputeEnvironment struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              BatchComputeEnvironmentSpec   `json:"spec,omitempty"`
	Status            BatchComputeEnvironmentStatus `json:"status,omitempty"`
}

type BatchComputeEnvironmentSpecComputeResourcesLaunchTemplate struct {
	// +optional
	LaunchTemplateId string `json:"launch_template_id,omitempty"`
	// +optional
	LaunchTemplateName string `json:"launch_template_name,omitempty"`
	// +optional
	Version string `json:"version,omitempty"`
}

type BatchComputeEnvironmentSpecComputeResources struct {
	// +optional
	BidPercentage int `json:"bid_percentage,omitempty"`
	// +optional
	DesiredVcpus int `json:"desired_vcpus,omitempty"`
	// +optional
	Ec2KeyPair string `json:"ec2_key_pair,omitempty"`
	// +optional
	ImageId      string `json:"image_id,omitempty"`
	InstanceRole string `json:"instance_role"`
	// +kubebuilder:validation:UniqueItems=true
	InstanceType []string `json:"instance_type"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	LaunchTemplate *[]BatchComputeEnvironmentSpecComputeResources `json:"launch_template,omitempty"`
	MaxVcpus       int                                            `json:"max_vcpus"`
	MinVcpus       int                                            `json:"min_vcpus"`
	// +kubebuilder:validation:UniqueItems=true
	SecurityGroupIds []string `json:"security_group_ids"`
	// +optional
	SpotIamFleetRole string `json:"spot_iam_fleet_role,omitempty"`
	// +kubebuilder:validation:UniqueItems=true
	Subnets []string `json:"subnets"`
	// +optional
	Tags map[string]string `json:"tags,omitempty"`
	Type string            `json:"type"`
}

type BatchComputeEnvironmentSpec struct {
	ComputeEnvironmentName string `json:"compute_environment_name"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	ComputeResources *[]BatchComputeEnvironmentSpec `json:"compute_resources,omitempty"`
	ServiceRole      string                         `json:"service_role"`
	// +optional
	State string `json:"state,omitempty"`
	Type  string `json:"type"`
}

type BatchComputeEnvironmentStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
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
