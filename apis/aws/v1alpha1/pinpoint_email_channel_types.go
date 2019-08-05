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

type PinpointEmailChannel struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              PinpointEmailChannelSpec   `json:"spec,omitempty"`
	Status            PinpointEmailChannelStatus `json:"status,omitempty"`
}

type PinpointEmailChannelSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ApplicationID string `json:"applicationID" tf:"application_id"`
	// +optional
	Enabled     bool   `json:"enabled,omitempty" tf:"enabled,omitempty"`
	FromAddress string `json:"fromAddress" tf:"from_address"`
	Identity    string `json:"identity" tf:"identity"`
	// +optional
	MessagesPerSecond int    `json:"messagesPerSecond,omitempty" tf:"messages_per_second,omitempty"`
	RoleArn           string `json:"roleArn" tf:"role_arn"`
}

type PinpointEmailChannelStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *PinpointEmailChannelSpec `json:"output,omitempty"`
	// +optional
	State *apis.State `json:"state,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// PinpointEmailChannelList is a list of PinpointEmailChannels
type PinpointEmailChannelList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of PinpointEmailChannel CRD objects
	Items []PinpointEmailChannel `json:"items,omitempty"`
}
