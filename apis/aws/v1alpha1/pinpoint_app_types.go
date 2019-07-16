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

type PinpointApp struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              PinpointAppSpec   `json:"spec,omitempty"`
	Status            PinpointAppStatus `json:"status,omitempty"`
}

type PinpointAppSpecCampaignHook struct {
	// +optional
	LambdaFunctionName string `json:"lambda_function_name,omitempty"`
	// +optional
	Mode string `json:"mode,omitempty"`
	// +optional
	WebUrl string `json:"web_url,omitempty"`
}

type PinpointAppSpecLimits struct {
	// +optional
	Daily int `json:"daily,omitempty"`
	// +optional
	MaximumDuration int `json:"maximum_duration,omitempty"`
	// +optional
	MessagesPerSecond int `json:"messages_per_second,omitempty"`
	// +optional
	Total int `json:"total,omitempty"`
}

type PinpointAppSpecQuietTime struct {
	// +optional
	End string `json:"end,omitempty"`
	// +optional
	Start string `json:"start,omitempty"`
}

type PinpointAppSpec struct {
	// +optional
	// +kubebuilder:validation:MaxItems=1
	CampaignHook *[]PinpointAppSpec `json:"campaign_hook,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	Limits *[]PinpointAppSpec `json:"limits,omitempty"`
	// +optional
	NamePrefix string `json:"name_prefix,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	QuietTime *[]PinpointAppSpec `json:"quiet_time,omitempty"`
}

type PinpointAppStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// PinpointAppList is a list of PinpointApps
type PinpointAppList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of PinpointApp CRD objects
	Items []PinpointApp `json:"items,omitempty"`
}
