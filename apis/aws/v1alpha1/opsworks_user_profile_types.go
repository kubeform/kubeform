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

type OpsworksUserProfile struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              OpsworksUserProfileSpec   `json:"spec,omitempty"`
	Status            OpsworksUserProfileStatus `json:"status,omitempty"`
}

type OpsworksUserProfileSpec struct {
	// +optional
	AllowSelfManagement bool `json:"allow_self_management,omitempty"`
	// +optional
	SshPublicKey string `json:"ssh_public_key,omitempty"`
	SshUsername  string `json:"ssh_username"`
	UserArn      string `json:"user_arn"`
}

type OpsworksUserProfileStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
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
