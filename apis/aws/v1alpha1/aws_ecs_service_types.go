package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsEcsService struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsEcsServiceSpec   `json:"spec,omitempty"`
	Status            AwsEcsServiceStatus `json:"status,omitempty"`
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

type AwsEcsServiceSpecOrderedPlacementStrategy struct {
	Type  string `json:"type"`
	Field string `json:"field"`
}

type AwsEcsServiceSpecServiceRegistries struct {
	ContainerName string `json:"container_name"`
	ContainerPort int    `json:"container_port"`
	Port          int    `json:"port"`
	RegistryArn   string `json:"registry_arn"`
}

type AwsEcsServiceSpecPlacementConstraints struct {
	Type       string `json:"type"`
	Expression string `json:"expression"`
}

type AwsEcsServiceSpecDeploymentController struct {
	Type string `json:"type"`
}

type AwsEcsServiceSpecPlacementStrategy struct {
	Type  string `json:"type"`
	Field string `json:"field"`
}

type AwsEcsServiceSpec struct {
	Tags                            map[string]string   `json:"tags"`
	SchedulingStrategy              string              `json:"scheduling_strategy"`
	NetworkConfiguration            []AwsEcsServiceSpec `json:"network_configuration"`
	IamRole                         string              `json:"iam_role"`
	LoadBalancer                    []AwsEcsServiceSpec `json:"load_balancer"`
	OrderedPlacementStrategy        []AwsEcsServiceSpec `json:"ordered_placement_strategy"`
	PropagateTags                   string              `json:"propagate_tags"`
	ServiceRegistries               []AwsEcsServiceSpec `json:"service_registries"`
	HealthCheckGracePeriodSeconds   int                 `json:"health_check_grace_period_seconds"`
	LaunchType                      string              `json:"launch_type"`
	PlatformVersion                 string              `json:"platform_version"`
	DeploymentMinimumHealthyPercent int                 `json:"deployment_minimum_healthy_percent"`
	PlacementConstraints            []AwsEcsServiceSpec `json:"placement_constraints"`
	Name                            string              `json:"name"`
	Cluster                         string              `json:"cluster"`
	EnableEcsManagedTags            bool                `json:"enable_ecs_managed_tags"`
	DeploymentController            []AwsEcsServiceSpec `json:"deployment_controller"`
	DeploymentMaximumPercent        int                 `json:"deployment_maximum_percent"`
	PlacementStrategy               []AwsEcsServiceSpec `json:"placement_strategy"`
	TaskDefinition                  string              `json:"task_definition"`
	DesiredCount                    int                 `json:"desired_count"`
}



type AwsEcsServiceStatus struct {
    Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsEcsServiceList is a list of AwsEcsServices
type AwsEcsServiceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsEcsService CRD objects
	Items []AwsEcsService `json:"items,omitempty"`
}