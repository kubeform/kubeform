package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsAutoscalingLifecycleHook struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsAutoscalingLifecycleHookSpec   `json:"spec,omitempty"`
	Status            AwsAutoscalingLifecycleHookStatus `json:"status,omitempty"`
}

type AwsAutoscalingLifecycleHookSpec struct {
	HeartbeatTimeout      int    `json:"heartbeat_timeout"`
	LifecycleTransition   string `json:"lifecycle_transition"`
	NotificationMetadata  string `json:"notification_metadata"`
	NotificationTargetArn string `json:"notification_target_arn"`
	RoleArn               string `json:"role_arn"`
	Name                  string `json:"name"`
	AutoscalingGroupName  string `json:"autoscaling_group_name"`
	DefaultResult         string `json:"default_result"`
}

type AwsAutoscalingLifecycleHookStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsAutoscalingLifecycleHookList is a list of AwsAutoscalingLifecycleHooks
type AwsAutoscalingLifecycleHookList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsAutoscalingLifecycleHook CRD objects
	Items []AwsAutoscalingLifecycleHook `json:"items,omitempty"`
}
