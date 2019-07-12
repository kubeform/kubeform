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
	UnhealthyThreshold int    `json:"unhealthy_threshold"`
	Enabled            bool   `json:"enabled"`
	Interval           int    `json:"interval"`
	Port               string `json:"port"`
	Timeout            int    `json:"timeout"`
	Path               string `json:"path"`
	Protocol           string `json:"protocol"`
	HealthyThreshold   int    `json:"healthy_threshold"`
	Matcher            string `json:"matcher"`
}

type AwsLbTargetGroupSpec struct {
	LambdaMultiValueHeadersEnabled bool                   `json:"lambda_multi_value_headers_enabled"`
	TargetType                     string                 `json:"target_type"`
	Arn                            string                 `json:"arn"`
	Protocol                       string                 `json:"protocol"`
	Name                           string                 `json:"name"`
	ProxyProtocolV2                bool                   `json:"proxy_protocol_v2"`
	SlowStart                      int                    `json:"slow_start"`
	Stickiness                     []AwsLbTargetGroupSpec `json:"stickiness"`
	ArnSuffix                      string                 `json:"arn_suffix"`
	DeregistrationDelay            int                    `json:"deregistration_delay"`
	VpcId                          string                 `json:"vpc_id"`
	HealthCheck                    []AwsLbTargetGroupSpec `json:"health_check"`
	Tags                           map[string]string      `json:"tags"`
	NamePrefix                     string                 `json:"name_prefix"`
	Port                           int                    `json:"port"`
}

type AwsLbTargetGroupStatus struct {
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
