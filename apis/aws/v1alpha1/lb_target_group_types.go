package v1alpha1

import (
	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

type LbTargetGroup struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              LbTargetGroupSpec   `json:"spec,omitempty"`
	Status            LbTargetGroupStatus `json:"status,omitempty"`
}

type LbTargetGroupSpecHealthCheck struct {
	// +optional
	Enabled bool `json:"enabled,omitempty" tf:"enabled,omitempty"`
	// +optional
	HealthyThreshold int `json:"healthyThreshold,omitempty" tf:"healthy_threshold,omitempty"`
	// +optional
	Interval int `json:"interval,omitempty" tf:"interval,omitempty"`
	// +optional
	Matcher string `json:"matcher,omitempty" tf:"matcher,omitempty"`
	// +optional
	Path string `json:"path,omitempty" tf:"path,omitempty"`
	// +optional
	Port string `json:"port,omitempty" tf:"port,omitempty"`
	// +optional
	Protocol string `json:"protocol,omitempty" tf:"protocol,omitempty"`
	// +optional
	Timeout int `json:"timeout,omitempty" tf:"timeout,omitempty"`
	// +optional
	UnhealthyThreshold int `json:"unhealthyThreshold,omitempty" tf:"unhealthy_threshold,omitempty"`
}

type LbTargetGroupSpecStickiness struct {
	// +optional
	CookieDuration int `json:"cookieDuration,omitempty" tf:"cookie_duration,omitempty"`
	// +optional
	Enabled bool   `json:"enabled,omitempty" tf:"enabled,omitempty"`
	Type    string `json:"type" tf:"type"`
}

type LbTargetGroupSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	// +optional
	DeregistrationDelay int `json:"deregistrationDelay,omitempty" tf:"deregistration_delay,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	HealthCheck []LbTargetGroupSpecHealthCheck `json:"healthCheck,omitempty" tf:"health_check,omitempty"`
	// +optional
	LambdaMultiValueHeadersEnabled bool `json:"lambdaMultiValueHeadersEnabled,omitempty" tf:"lambda_multi_value_headers_enabled,omitempty"`
	// +optional
	Name string `json:"name,omitempty" tf:"name,omitempty"`
	// +optional
	NamePrefix string `json:"namePrefix,omitempty" tf:"name_prefix,omitempty"`
	// +optional
	Port int `json:"port,omitempty" tf:"port,omitempty"`
	// +optional
	Protocol string `json:"protocol,omitempty" tf:"protocol,omitempty"`
	// +optional
	ProxyProtocolV2 bool `json:"proxyProtocolV2,omitempty" tf:"proxy_protocol_v2,omitempty"`
	// +optional
	SlowStart int `json:"slowStart,omitempty" tf:"slow_start,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	Stickiness []LbTargetGroupSpecStickiness `json:"stickiness,omitempty" tf:"stickiness,omitempty"`
	// +optional
	Tags map[string]string `json:"tags,omitempty" tf:"tags,omitempty"`
	// +optional
	TargetType string `json:"targetType,omitempty" tf:"target_type,omitempty"`
	// +optional
	VpcID string `json:"vpcID,omitempty" tf:"vpc_id,omitempty"`
}

type LbTargetGroupStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	State *LbTargetGroupSpec `json:"state,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// LbTargetGroupList is a list of LbTargetGroups
type LbTargetGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of LbTargetGroup CRD objects
	Items []LbTargetGroup `json:"items,omitempty"`
}
