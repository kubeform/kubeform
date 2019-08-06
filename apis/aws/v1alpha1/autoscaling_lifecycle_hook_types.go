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

type AutoscalingLifecycleHook struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AutoscalingLifecycleHookSpec   `json:"spec,omitempty"`
	Status            AutoscalingLifecycleHookStatus `json:"status,omitempty"`
}

type AutoscalingLifecycleHookSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	AutoscalingGroupName string `json:"autoscalingGroupName" tf:"autoscaling_group_name"`
	// +optional
	DefaultResult string `json:"defaultResult,omitempty" tf:"default_result,omitempty"`
	// +optional
	HeartbeatTimeout    int    `json:"heartbeatTimeout,omitempty" tf:"heartbeat_timeout,omitempty"`
	LifecycleTransition string `json:"lifecycleTransition" tf:"lifecycle_transition"`
	Name                string `json:"name" tf:"name"`
	// +optional
	NotificationMetadata string `json:"notificationMetadata,omitempty" tf:"notification_metadata,omitempty"`
	// +optional
	NotificationTargetArn string `json:"notificationTargetArn,omitempty" tf:"notification_target_arn,omitempty"`
	// +optional
	RoleArn string `json:"roleArn,omitempty" tf:"role_arn,omitempty"`
}

type AutoscalingLifecycleHookStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *AutoscalingLifecycleHookSpec `json:"output,omitempty"`
	// +optional
	State *apis.State `json:"state,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AutoscalingLifecycleHookList is a list of AutoscalingLifecycleHooks
type AutoscalingLifecycleHookList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AutoscalingLifecycleHook CRD objects
	Items []AutoscalingLifecycleHook `json:"items,omitempty"`
}
