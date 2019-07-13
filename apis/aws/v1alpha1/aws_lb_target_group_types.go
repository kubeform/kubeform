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

type AwsLbTargetGroup struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsLbTargetGroupSpec   `json:"spec,omitempty"`
	Status            AwsLbTargetGroupStatus `json:"status,omitempty"`
}

type AwsLbTargetGroupSpecStickiness struct {
	Enabled        bool   `json:"enabled"`
	Type           string `json:"type"`
	CookieDuration int    `json:"cookie_duration"`
}

type AwsLbTargetGroupSpecHealthCheck struct {
	HealthyThreshold   int    `json:"healthy_threshold"`
	Matcher            string `json:"matcher"`
	Interval           int    `json:"interval"`
	Port               string `json:"port"`
	Timeout            int    `json:"timeout"`
	UnhealthyThreshold int    `json:"unhealthy_threshold"`
	Enabled            bool   `json:"enabled"`
	Path               string `json:"path"`
	Protocol           string `json:"protocol"`
}

type AwsLbTargetGroupSpec struct {
	ArnSuffix                      string                 `json:"arn_suffix"`
	Port                           int                    `json:"port"`
	NamePrefix                     string                 `json:"name_prefix"`
	Protocol                       string                 `json:"protocol"`
	VpcId                          string                 `json:"vpc_id"`
	Stickiness                     []AwsLbTargetGroupSpec `json:"stickiness"`
	Name                           string                 `json:"name"`
	TargetType                     string                 `json:"target_type"`
	HealthCheck                    []AwsLbTargetGroupSpec `json:"health_check"`
	Tags                           map[string]string      `json:"tags"`
	LambdaMultiValueHeadersEnabled bool                   `json:"lambda_multi_value_headers_enabled"`
	Arn                            string                 `json:"arn"`
	DeregistrationDelay            int                    `json:"deregistration_delay"`
	SlowStart                      int                    `json:"slow_start"`
	ProxyProtocolV2                bool                   `json:"proxy_protocol_v2"`
}



type AwsLbTargetGroupStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AwsLbTargetGroupList is a list of AwsLbTargetGroups
type AwsLbTargetGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsLbTargetGroup CRD objects
	Items []AwsLbTargetGroup `json:"items,omitempty"`
}