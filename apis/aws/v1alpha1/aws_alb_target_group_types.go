package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsAlbTargetGroup struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsAlbTargetGroupSpec   `json:"spec,omitempty"`
	Status            AwsAlbTargetGroupStatus `json:"status,omitempty"`
}

type AwsAlbTargetGroupSpecHealthCheck struct {
	UnhealthyThreshold int    `json:"unhealthy_threshold"`
	Path               string `json:"path"`
	Port               string `json:"port"`
	Timeout            int    `json:"timeout"`
	HealthyThreshold   int    `json:"healthy_threshold"`
	Matcher            string `json:"matcher"`
	Enabled            bool   `json:"enabled"`
	Interval           int    `json:"interval"`
	Protocol           string `json:"protocol"`
}

type AwsAlbTargetGroupSpecStickiness struct {
	Type           string `json:"type"`
	CookieDuration int    `json:"cookie_duration"`
	Enabled        bool   `json:"enabled"`
}

type AwsAlbTargetGroupSpec struct {
	Name                           string                  `json:"name"`
	NamePrefix                     string                  `json:"name_prefix"`
	LambdaMultiValueHeadersEnabled bool                    `json:"lambda_multi_value_headers_enabled"`
	Arn                            string                  `json:"arn"`
	ArnSuffix                      string                  `json:"arn_suffix"`
	DeregistrationDelay            int                     `json:"deregistration_delay"`
	HealthCheck                    []AwsAlbTargetGroupSpec `json:"health_check"`
	SlowStart                      int                     `json:"slow_start"`
	ProxyProtocolV2                bool                    `json:"proxy_protocol_v2"`
	TargetType                     string                  `json:"target_type"`
	Tags                           map[string]string       `json:"tags"`
	Port                           int                     `json:"port"`
	VpcId                          string                  `json:"vpc_id"`
	Protocol                       string                  `json:"protocol"`
	Stickiness                     []AwsAlbTargetGroupSpec `json:"stickiness"`
}



type AwsAlbTargetGroupStatus struct {
    Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsAlbTargetGroupList is a list of AwsAlbTargetGroups
type AwsAlbTargetGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsAlbTargetGroup CRD objects
	Items []AwsAlbTargetGroup `json:"items,omitempty"`
}