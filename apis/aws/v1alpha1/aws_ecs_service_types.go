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

type AwsEcsService struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsEcsServiceSpec   `json:"spec,omitempty"`
	Status            AwsEcsServiceStatus `json:"status,omitempty"`
}

type AwsEcsServiceSpecServiceRegistries struct {
	ContainerName string `json:"container_name"`
	ContainerPort int    `json:"container_port"`
	Port          int    `json:"port"`
	RegistryArn   string `json:"registry_arn"`
}

type AwsEcsServiceSpecDeploymentController struct {
	Type string `json:"type"`
}

type AwsEcsServiceSpecNetworkConfiguration struct {
	SecurityGroups []string `json:"security_groups"`
	Subnets        []string `json:"subnets"`
	AssignPublicIp bool     `json:"assign_public_ip"`
}

type AwsEcsServiceSpecLoadBalancer struct {
	ElbName        string `json:"elb_name"`
	TargetGroupArn string `json:"target_group_arn"`
	ContainerName  string `json:"container_name"`
	ContainerPort  int    `json:"container_port"`
}

type AwsEcsServiceSpecPlacementStrategy struct {
	Type  string `json:"type"`
	Field string `json:"field"`
}

type AwsEcsServiceSpecPlacementConstraints struct {
	Expression string `json:"expression"`
	Type       string `json:"type"`
}

type AwsEcsServiceSpecOrderedPlacementStrategy struct {
	Type  string `json:"type"`
	Field string `json:"field"`
}

type AwsEcsServiceSpec struct {
	ServiceRegistries               []AwsEcsServiceSpec `json:"service_registries"`
	Cluster                         string              `json:"cluster"`
	TaskDefinition                  string              `json:"task_definition"`
	LaunchType                      string              `json:"launch_type"`
	IamRole                         string              `json:"iam_role"`
	DeploymentController            []AwsEcsServiceSpec `json:"deployment_controller"`
	NetworkConfiguration            []AwsEcsServiceSpec `json:"network_configuration"`
	PropagateTags                   string              `json:"propagate_tags"`
	EnableEcsManagedTags            bool                `json:"enable_ecs_managed_tags"`
	DeploymentMaximumPercent        int                 `json:"deployment_maximum_percent"`
	DeploymentMinimumHealthyPercent int                 `json:"deployment_minimum_healthy_percent"`
	LoadBalancer                    []AwsEcsServiceSpec `json:"load_balancer"`
	PlacementStrategy               []AwsEcsServiceSpec `json:"placement_strategy"`
	PlacementConstraints            []AwsEcsServiceSpec `json:"placement_constraints"`
	Name                            string              `json:"name"`
	OrderedPlacementStrategy        []AwsEcsServiceSpec `json:"ordered_placement_strategy"`
	DesiredCount                    int                 `json:"desired_count"`
	HealthCheckGracePeriodSeconds   int                 `json:"health_check_grace_period_seconds"`
	PlatformVersion                 string              `json:"platform_version"`
	SchedulingStrategy              string              `json:"scheduling_strategy"`
	Tags                            map[string]string   `json:"tags"`
}



type AwsEcsServiceStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AwsEcsServiceList is a list of AwsEcsServices
type AwsEcsServiceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsEcsService CRD objects
	Items []AwsEcsService `json:"items,omitempty"`
}