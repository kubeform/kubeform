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

type AwsPinpointApp struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsPinpointAppSpec   `json:"spec,omitempty"`
	Status            AwsPinpointAppStatus `json:"status,omitempty"`
}

type AwsPinpointAppSpecCampaignHook struct {
	Mode               string `json:"mode"`
	WebUrl             string `json:"web_url"`
	LambdaFunctionName string `json:"lambda_function_name"`
}

type AwsPinpointAppSpecLimits struct {
	Total             int `json:"total"`
	Daily             int `json:"daily"`
	MaximumDuration   int `json:"maximum_duration"`
	MessagesPerSecond int `json:"messages_per_second"`
}

type AwsPinpointAppSpecQuietTime struct {
	Start string `json:"start"`
	End   string `json:"end"`
}

type AwsPinpointAppSpec struct {
	CampaignHook  []AwsPinpointAppSpec `json:"campaign_hook"`
	Limits        []AwsPinpointAppSpec `json:"limits"`
	QuietTime     []AwsPinpointAppSpec `json:"quiet_time"`
	Name          string               `json:"name"`
	NamePrefix    string               `json:"name_prefix"`
	ApplicationId string               `json:"application_id"`
}



type AwsPinpointAppStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AwsPinpointAppList is a list of AwsPinpointApps
type AwsPinpointAppList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsPinpointApp CRD objects
	Items []AwsPinpointApp `json:"items,omitempty"`
}