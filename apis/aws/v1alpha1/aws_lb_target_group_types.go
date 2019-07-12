package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsLbTargetGroup struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsLbTargetGroupSpec   `json:"spec,omitempty"`
	Status            AwsLbTargetGroupStatus `json:"status,omitempty"`
}

type AwsLbTargetGroupSpecHealthCheck struct {
	Interval           int    `json:"interval"`
	Port               string `json:"port"`
	Protocol           string `json:"protocol"`
	Timeout            int    `json:"timeout"`
	HealthyThreshold   int    `json:"healthy_threshold"`
	Matcher            string `json:"matcher"`
	Enabled            bool   `json:"enabled"`
	Path               string `json:"path"`
	UnhealthyThreshold int    `json:"unhealthy_threshold"`
}

type AwsLbTargetGroupSpecStickiness struct {
	Type           string `json:"type"`
	CookieDuration int    `json:"cookie_duration"`
	Enabled        bool   `json:"enabled"`
}

type AwsLbTargetGroupSpec struct {
	Port                           int                    `json:"port"`
	LambdaMultiValueHeadersEnabled bool                   `json:"lambda_multi_value_headers_enabled"`
	TargetType                     string                 `json:"target_type"`
	ArnSuffix                      string                 `json:"arn_suffix"`
	Name                           string                 `json:"name"`
	SlowStart                      int                    `json:"slow_start"`
	HealthCheck                    []AwsLbTargetGroupSpec `json:"health_check"`
	Tags                           map[string]string      `json:"tags"`
	Arn                            string                 `json:"arn"`
	Protocol                       string                 `json:"protocol"`
	ProxyProtocolV2                bool                   `json:"proxy_protocol_v2"`
	Stickiness                     []AwsLbTargetGroupSpec `json:"stickiness"`
	NamePrefix                     string                 `json:"name_prefix"`
	VpcId                          string                 `json:"vpc_id"`
	DeregistrationDelay            int                    `json:"deregistration_delay"`
}



type AwsLbTargetGroupStatus struct {
    Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsLbTargetGroupList is a list of AwsLbTargetGroups
type AwsLbTargetGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsLbTargetGroup CRD objects
	Items []AwsLbTargetGroup `json:"items,omitempty"`
}