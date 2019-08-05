package v1alpha1

import (
	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"kubeform.dev/kubeform/apis"
)

// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

type PinpointBaiduChannel struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              PinpointBaiduChannelSpec   `json:"spec,omitempty"`
	Status            PinpointBaiduChannelStatus `json:"status,omitempty"`
}

type PinpointBaiduChannelSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	KubeFormSecret *core.LocalObjectReference `json:"secret,omitempty" tf:"-"`

	ApiKey        string `json:"-" sensitive:"true" tf:"api_key"`
	ApplicationID string `json:"applicationID" tf:"application_id"`
	// +optional
	Enabled   bool   `json:"enabled,omitempty" tf:"enabled,omitempty"`
	SecretKey string `json:"-" sensitive:"true" tf:"secret_key"`
}

type PinpointBaiduChannelStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *PinpointBaiduChannelSpec `json:"output,omitempty"`
	// +optional
	State *apis.State `json:"state,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// PinpointBaiduChannelList is a list of PinpointBaiduChannels
type PinpointBaiduChannelList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of PinpointBaiduChannel CRD objects
	Items []PinpointBaiduChannel `json:"items,omitempty"`
}
