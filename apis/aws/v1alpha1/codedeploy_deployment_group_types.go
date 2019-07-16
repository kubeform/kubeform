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

type CodedeployDeploymentGroup struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              CodedeployDeploymentGroupSpec   `json:"spec,omitempty"`
	Status            CodedeployDeploymentGroupStatus `json:"status,omitempty"`
}

type CodedeployDeploymentGroupSpecAlarmConfiguration struct {
	// +optional
	// +kubebuilder:validation:MaxItems=10
	// +kubebuilder:validation:UniqueItems=true
	Alarms []string `json:"alarms,omitempty"`
	// +optional
	Enabled bool `json:"enabled,omitempty"`
	// +optional
	IgnorePollAlarmFailure bool `json:"ignore_poll_alarm_failure,omitempty"`
}

type CodedeployDeploymentGroupSpecAutoRollbackConfiguration struct {
	// +optional
	Enabled bool `json:"enabled,omitempty"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	Events []string `json:"events,omitempty"`
}

type CodedeployDeploymentGroupSpecEc2TagFilter struct {
	// +optional
	Key string `json:"key,omitempty"`
	// +optional
	Type string `json:"type,omitempty"`
	// +optional
	Value string `json:"value,omitempty"`
}

type CodedeployDeploymentGroupSpecEc2TagSetEc2TagFilter struct {
	// +optional
	Key string `json:"key,omitempty"`
	// +optional
	Type string `json:"type,omitempty"`
	// +optional
	Value string `json:"value,omitempty"`
}

type CodedeployDeploymentGroupSpecEc2TagSet struct {
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	Ec2TagFilter *[]CodedeployDeploymentGroupSpecEc2TagSet `json:"ec2_tag_filter,omitempty"`
}

type CodedeployDeploymentGroupSpecEcsService struct {
	ClusterName string `json:"cluster_name"`
	ServiceName string `json:"service_name"`
}

type CodedeployDeploymentGroupSpecOnPremisesInstanceTagFilter struct {
	// +optional
	Key string `json:"key,omitempty"`
	// +optional
	Type string `json:"type,omitempty"`
	// +optional
	Value string `json:"value,omitempty"`
}

type CodedeployDeploymentGroupSpecTriggerConfiguration struct {
	// +kubebuilder:validation:UniqueItems=true
	TriggerEvents    []string `json:"trigger_events"`
	TriggerName      string   `json:"trigger_name"`
	TriggerTargetArn string   `json:"trigger_target_arn"`
}

type CodedeployDeploymentGroupSpec struct {
	// +optional
	// +kubebuilder:validation:MaxItems=1
	AlarmConfiguration *[]CodedeployDeploymentGroupSpec `json:"alarm_configuration,omitempty"`
	AppName            string                           `json:"app_name"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	AutoRollbackConfiguration *[]CodedeployDeploymentGroupSpec `json:"auto_rollback_configuration,omitempty"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	AutoscalingGroups []string `json:"autoscaling_groups,omitempty"`
	// +optional
	DeploymentConfigName string `json:"deployment_config_name,omitempty"`
	DeploymentGroupName  string `json:"deployment_group_name"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	Ec2TagFilter *[]CodedeployDeploymentGroupSpec `json:"ec2_tag_filter,omitempty"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	Ec2TagSet *[]CodedeployDeploymentGroupSpec `json:"ec2_tag_set,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	EcsService *[]CodedeployDeploymentGroupSpec `json:"ecs_service,omitempty"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	OnPremisesInstanceTagFilter *[]CodedeployDeploymentGroupSpec `json:"on_premises_instance_tag_filter,omitempty"`
	ServiceRoleArn              string                           `json:"service_role_arn"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	TriggerConfiguration *[]CodedeployDeploymentGroupSpec `json:"trigger_configuration,omitempty"`
}

type CodedeployDeploymentGroupStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// CodedeployDeploymentGroupList is a list of CodedeployDeploymentGroups
type CodedeployDeploymentGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of CodedeployDeploymentGroup CRD objects
	Items []CodedeployDeploymentGroup `json:"items,omitempty"`
}
