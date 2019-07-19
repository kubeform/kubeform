package v1alpha1

import (
	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
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
	LambdaFunctionName string `json:"lambdaFunctionName,omitempty" tf:"lambda_function_name,omitempty"`
	// +optional
	Mode string `json:"mode,omitempty" tf:"mode,omitempty"`
	// +optional
	WebURL string `json:"webURL,omitempty" tf:"web_url,omitempty"`
}

type PinpointAppSpecLimits struct {
	// +optional
	Daily int `json:"daily,omitempty" tf:"daily,omitempty"`
	// +optional
	MaximumDuration int `json:"maximumDuration,omitempty" tf:"maximum_duration,omitempty"`
	// +optional
	MessagesPerSecond int `json:"messagesPerSecond,omitempty" tf:"messages_per_second,omitempty"`
	// +optional
	Total int `json:"total,omitempty" tf:"total,omitempty"`
}

type PinpointAppSpecQuietTime struct {
	// +optional
	End string `json:"end,omitempty" tf:"end,omitempty"`
	// +optional
	Start string `json:"start,omitempty" tf:"start,omitempty"`
}

type PinpointAppSpec struct {
	// +optional
	// +kubebuilder:validation:MaxItems=1
	CampaignHook []PinpointAppSpecCampaignHook `json:"campaignHook,omitempty" tf:"campaign_hook,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	Limits []PinpointAppSpecLimits `json:"limits,omitempty" tf:"limits,omitempty"`
	// +optional
	Name string `json:"name,omitempty" tf:"name,omitempty"`
	// +optional
	NamePrefix string `json:"namePrefix,omitempty" tf:"name_prefix,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	QuietTime   []PinpointAppSpecQuietTime `json:"quietTime,omitempty" tf:"quiet_time,omitempty"`
	ProviderRef core.LocalObjectReference  `json:"providerRef" tf:"-"`
}

type PinpointAppStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	TFState *runtime.RawExtension `json:"tfState,omitempty"`
	Output  *runtime.RawExtension `json:"output,omitempty"`
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
