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

type AwsAlbTargetGroup struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsAlbTargetGroupSpec   `json:"spec,omitempty"`
	Status            AwsAlbTargetGroupStatus `json:"status,omitempty"`
}

type AwsAlbTargetGroupSpecHealthCheck struct {
	HealthyThreshold   int    `json:"healthy_threshold"`
	Matcher            string `json:"matcher"`
	UnhealthyThreshold int    `json:"unhealthy_threshold"`
	Enabled            bool   `json:"enabled"`
	Interval           int    `json:"interval"`
	Path               string `json:"path"`
	Port               string `json:"port"`
	Protocol           string `json:"protocol"`
	Timeout            int    `json:"timeout"`
}

type AwsAlbTargetGroupSpecStickiness struct {
	Enabled        bool   `json:"enabled"`
	Type           string `json:"type"`
	CookieDuration int    `json:"cookie_duration"`
}

type AwsAlbTargetGroupSpec struct {
	Port                           int                     `json:"port"`
	Protocol                       string                  `json:"protocol"`
	VpcId                          string                  `json:"vpc_id"`
	ProxyProtocolV2                bool                    `json:"proxy_protocol_v2"`
	HealthCheck                    []AwsAlbTargetGroupSpec `json:"health_check"`
	Arn                            string                  `json:"arn"`
	NamePrefix                     string                  `json:"name_prefix"`
	SlowStart                      int                     `json:"slow_start"`
	LambdaMultiValueHeadersEnabled bool                    `json:"lambda_multi_value_headers_enabled"`
	Stickiness                     []AwsAlbTargetGroupSpec `json:"stickiness"`
	ArnSuffix                      string                  `json:"arn_suffix"`
	DeregistrationDelay            int                     `json:"deregistration_delay"`
	TargetType                     string                  `json:"target_type"`
	Name                           string                  `json:"name"`
	Tags                           map[string]string       `json:"tags"`
}

type AwsAlbTargetGroupStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AwsAlbTargetGroupList is a list of AwsAlbTargetGroups
type AwsAlbTargetGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsAlbTargetGroup CRD objects
	Items []AwsAlbTargetGroup `json:"items,omitempty"`
}
