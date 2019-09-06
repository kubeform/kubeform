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

type OpsworksUserProfile struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              OpsworksUserProfileSpec   `json:"spec,omitempty"`
	Status            OpsworksUserProfileStatus `json:"status,omitempty"`
}

type OpsworksUserProfileSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	AllowSelfManagement bool `json:"allowSelfManagement,omitempty" tf:"allow_self_management,omitempty"`
	// +optional
	SshPublicKey string `json:"sshPublicKey,omitempty" tf:"ssh_public_key,omitempty"`
	SshUsername  string `json:"sshUsername" tf:"ssh_username"`
	UserArn      string `json:"userArn" tf:"user_arn"`
}

type OpsworksUserProfileStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *OpsworksUserProfileSpec `json:"output,omitempty"`
	// +optional
	State *apis.State `json:"state,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// OpsworksUserProfileList is a list of OpsworksUserProfiles
type OpsworksUserProfileList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of OpsworksUserProfile CRD objects
	Items []OpsworksUserProfile `json:"items,omitempty"`
}
