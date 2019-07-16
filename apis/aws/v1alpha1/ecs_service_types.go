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

type EcsService struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              EcsServiceSpec   `json:"spec,omitempty"`
	Status            EcsServiceStatus `json:"status,omitempty"`
}

type EcsServiceSpecDeploymentController struct {
	// +optional
	Type string `json:"type,omitempty"`
}

type EcsServiceSpecLoadBalancer struct {
	ContainerName string `json:"container_name"`
	ContainerPort int    `json:"container_port"`
	// +optional
	ElbName string `json:"elb_name,omitempty"`
	// +optional
	TargetGroupArn string `json:"target_group_arn,omitempty"`
}

type EcsServiceSpecNetworkConfiguration struct {
	// +optional
	AssignPublicIp bool `json:"assign_public_ip,omitempty"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	SecurityGroups []string `json:"security_groups,omitempty"`
	// +kubebuilder:validation:UniqueItems=true
	Subnets []string `json:"subnets"`
}

type EcsServiceSpecOrderedPlacementStrategy struct {
	// +optional
	Field string `json:"field,omitempty"`
	Type  string `json:"type"`
}

type EcsServiceSpecPlacementConstraints struct {
	// +optional
	Expression string `json:"expression,omitempty"`
	Type       string `json:"type"`
}

type EcsServiceSpecServiceRegistries struct {
	// +optional
	ContainerName string `json:"container_name,omitempty"`
	// +optional
	ContainerPort int `json:"container_port,omitempty"`
	// +optional
	Port        int    `json:"port,omitempty"`
	RegistryArn string `json:"registry_arn"`
}

type EcsServiceSpec struct {
	// +optional
	// +kubebuilder:validation:MaxItems=1
	DeploymentController *[]EcsServiceSpec `json:"deployment_controller,omitempty"`
	// +optional
	DeploymentMaximumPercent int `json:"deployment_maximum_percent,omitempty"`
	// +optional
	DeploymentMinimumHealthyPercent int `json:"deployment_minimum_healthy_percent,omitempty"`
	// +optional
	DesiredCount int `json:"desired_count,omitempty"`
	// +optional
	EnableEcsManagedTags bool `json:"enable_ecs_managed_tags,omitempty"`
	// +optional
	HealthCheckGracePeriodSeconds int `json:"health_check_grace_period_seconds,omitempty"`
	// +optional
	LaunchType string `json:"launch_type,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	// +kubebuilder:validation:UniqueItems=true
	LoadBalancer *[]EcsServiceSpec `json:"load_balancer,omitempty"`
	Name         string            `json:"name"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	NetworkConfiguration *[]EcsServiceSpec `json:"network_configuration,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=5
	OrderedPlacementStrategy *[]EcsServiceSpec `json:"ordered_placement_strategy,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=10
	// +kubebuilder:validation:UniqueItems=true
	PlacementConstraints *[]EcsServiceSpec `json:"placement_constraints,omitempty"`
	// +optional
	PropagateTags string `json:"propagate_tags,omitempty"`
	// +optional
	SchedulingStrategy string `json:"scheduling_strategy,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	// +kubebuilder:validation:UniqueItems=true
	ServiceRegistries *[]EcsServiceSpec `json:"service_registries,omitempty"`
	// +optional
	Tags           map[string]string `json:"tags,omitempty"`
	TaskDefinition string            `json:"task_definition"`
}

type EcsServiceStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// EcsServiceList is a list of EcsServices
type EcsServiceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of EcsService CRD objects
	Items []EcsService `json:"items,omitempty"`
}
