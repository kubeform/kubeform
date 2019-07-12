package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsPinpointApp struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsPinpointAppSpec   `json:"spec,omitempty"`
	Status            AwsPinpointAppStatus `json:"status,omitempty"`
}

type AwsPinpointAppSpecCampaignHook struct {
	LambdaFunctionName string `json:"lambda_function_name"`
	Mode               string `json:"mode"`
	WebUrl             string `json:"web_url"`
}

type AwsPinpointAppSpecLimits struct {
	MaximumDuration   int `json:"maximum_duration"`
	MessagesPerSecond int `json:"messages_per_second"`
	Total             int `json:"total"`
	Daily             int `json:"daily"`
}

type AwsPinpointAppSpecQuietTime struct {
	End   string `json:"end"`
	Start string `json:"start"`
}

type AwsPinpointAppSpec struct {
	Name          string               `json:"name"`
	NamePrefix    string               `json:"name_prefix"`
	ApplicationId string               `json:"application_id"`
	CampaignHook  []AwsPinpointAppSpec `json:"campaign_hook"`
	Limits        []AwsPinpointAppSpec `json:"limits"`
	QuietTime     []AwsPinpointAppSpec `json:"quiet_time"`
}



type AwsPinpointAppStatus struct {
    Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsPinpointAppList is a list of AwsPinpointApps
type AwsPinpointAppList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsPinpointApp CRD objects
	Items []AwsPinpointApp `json:"items,omitempty"`
}