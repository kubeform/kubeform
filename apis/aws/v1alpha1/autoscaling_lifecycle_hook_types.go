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

type AutoscalingLifecycleHook struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AutoscalingLifecycleHookSpec   `json:"spec,omitempty"`
	Status            AutoscalingLifecycleHookStatus `json:"status,omitempty"`
}

type AutoscalingLifecycleHookSpec struct {
	AutoscalingGroupName string `json:"autoscaling_group_name"`
	// +optional
	HeartbeatTimeout    int    `json:"heartbeat_timeout,omitempty"`
	LifecycleTransition string `json:"lifecycle_transition"`
	Name                string `json:"name"`
	// +optional
	NotificationMetadata string `json:"notification_metadata,omitempty"`
	// +optional
	NotificationTargetArn string `json:"notification_target_arn,omitempty"`
	// +optional
	RoleArn string `json:"role_arn,omitempty"`
}

type AutoscalingLifecycleHookStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
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
