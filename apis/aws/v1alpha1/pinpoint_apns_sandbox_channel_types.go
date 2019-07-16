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

type PinpointApnsSandboxChannel struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              PinpointApnsSandboxChannelSpec   `json:"spec,omitempty"`
	Status            PinpointApnsSandboxChannelStatus `json:"status,omitempty"`
}

type PinpointApnsSandboxChannelSpec struct {
	ApplicationId string `json:"application_id"`
	// +optional
	BundleId string `json:"bundle_id,omitempty"`
	// +optional
	Certificate string `json:"certificate,omitempty"`
	// +optional
	DefaultAuthenticationMethod string `json:"default_authentication_method,omitempty"`
	// +optional
	Enabled bool `json:"enabled,omitempty"`
	// +optional
	PrivateKey string `json:"private_key,omitempty"`
	// +optional
	TeamId string `json:"team_id,omitempty"`
	// +optional
	TokenKey string `json:"token_key,omitempty"`
	// +optional
	TokenKeyId string `json:"token_key_id,omitempty"`
}

type PinpointApnsSandboxChannelStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// PinpointApnsSandboxChannelList is a list of PinpointApnsSandboxChannels
type PinpointApnsSandboxChannelList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of PinpointApnsSandboxChannel CRD objects
	Items []PinpointApnsSandboxChannel `json:"items,omitempty"`
}
