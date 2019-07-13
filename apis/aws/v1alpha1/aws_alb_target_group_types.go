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
	Path               string `json:"path"`
	Port               string `json:"port"`
	Protocol           string `json:"protocol"`
	HealthyThreshold   int    `json:"healthy_threshold"`
	UnhealthyThreshold int    `json:"unhealthy_threshold"`
	Enabled            bool   `json:"enabled"`
	Interval           int    `json:"interval"`
	Timeout            int    `json:"timeout"`
	Matcher            string `json:"matcher"`
}

type AwsAlbTargetGroupSpecStickiness struct {
	Enabled        bool   `json:"enabled"`
	Type           string `json:"type"`
	CookieDuration int    `json:"cookie_duration"`
}

type AwsAlbTargetGroupSpec struct {
	Protocol                       string                  `json:"protocol"`
	DeregistrationDelay            int                     `json:"deregistration_delay"`
	ProxyProtocolV2                bool                    `json:"proxy_protocol_v2"`
	LambdaMultiValueHeadersEnabled bool                    `json:"lambda_multi_value_headers_enabled"`
	HealthCheck                    []AwsAlbTargetGroupSpec `json:"health_check"`
	ArnSuffix                      string                  `json:"arn_suffix"`
	Name                           string                  `json:"name"`
	NamePrefix                     string                  `json:"name_prefix"`
	Port                           int                     `json:"port"`
	VpcId                          string                  `json:"vpc_id"`
	Arn                            string                  `json:"arn"`
	TargetType                     string                  `json:"target_type"`
	SlowStart                      int                     `json:"slow_start"`
	Stickiness                     []AwsAlbTargetGroupSpec `json:"stickiness"`
	Tags                           map[string]string       `json:"tags"`
}



type AwsAlbTargetGroupStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

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