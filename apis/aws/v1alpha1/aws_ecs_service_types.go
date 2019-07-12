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

type AwsEcsServiceSpecDeploymentController struct {
	Type string `json:"type"`
}

type AwsEcsServiceSpecOrderedPlacementStrategy struct {
	Type  string `json:"type"`
	Field string `json:"field"`
}

type AwsEcsServiceSpecNetworkConfiguration struct {
	AssignPublicIp bool     `json:"assign_public_ip"`
	SecurityGroups []string `json:"security_groups"`
	Subnets        []string `json:"subnets"`
}

type AwsEcsServiceSpecServiceRegistries struct {
	ContainerName string `json:"container_name"`
	ContainerPort int    `json:"container_port"`
	Port          int    `json:"port"`
	RegistryArn   string `json:"registry_arn"`
}

type AwsEcsServiceSpecPlacementStrategy struct {
	Type  string `json:"type"`
	Field string `json:"field"`
}

type AwsEcsServiceSpecPlacementConstraints struct {
	Type       string `json:"type"`
	Expression string `json:"expression"`
}

type AwsEcsServiceSpecLoadBalancer struct {
	ContainerPort  int    `json:"container_port"`
	ElbName        string `json:"elb_name"`
	TargetGroupArn string `json:"target_group_arn"`
	ContainerName  string `json:"container_name"`
}

type AwsEcsServiceSpec struct {
	SchedulingStrategy              string              `json:"scheduling_strategy"`
	IamRole                         string              `json:"iam_role"`
	DeploymentMaximumPercent        int                 `json:"deployment_maximum_percent"`
	PropagateTags                   string              `json:"propagate_tags"`
	Tags                            map[string]string   `json:"tags"`
	Name                            string              `json:"name"`
	LaunchType                      string              `json:"launch_type"`
	DesiredCount                    int                 `json:"desired_count"`
	PlatformVersion                 string              `json:"platform_version"`
	DeploymentController            []AwsEcsServiceSpec `json:"deployment_controller"`
	OrderedPlacementStrategy        []AwsEcsServiceSpec `json:"ordered_placement_strategy"`
	Cluster                         string              `json:"cluster"`
	TaskDefinition                  string              `json:"task_definition"`
	NetworkConfiguration            []AwsEcsServiceSpec `json:"network_configuration"`
	ServiceRegistries               []AwsEcsServiceSpec `json:"service_registries"`
	EnableEcsManagedTags            bool                `json:"enable_ecs_managed_tags"`
	DeploymentMinimumHealthyPercent int                 `json:"deployment_minimum_healthy_percent"`
	PlacementStrategy               []AwsEcsServiceSpec `json:"placement_strategy"`
	PlacementConstraints            []AwsEcsServiceSpec `json:"placement_constraints"`
	HealthCheckGracePeriodSeconds   int                 `json:"health_check_grace_period_seconds"`
	LoadBalancer                    []AwsEcsServiceSpec `json:"load_balancer"`
}

type AwsEcsServiceStatus struct {
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
