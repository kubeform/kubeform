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

type AwsOpsworksUserProfile struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsOpsworksUserProfileSpec   `json:"spec,omitempty"`
	Status            AwsOpsworksUserProfileStatus `json:"status,omitempty"`
}

type AwsOpsworksUserProfileSpec struct {
	UserArn             string `json:"user_arn"`
	AllowSelfManagement bool   `json:"allow_self_management"`
	SshUsername         string `json:"ssh_username"`
	SshPublicKey        string `json:"ssh_public_key"`
}



type AwsOpsworksUserProfileStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AwsOpsworksUserProfileList is a list of AwsOpsworksUserProfiles
type AwsOpsworksUserProfileList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsOpsworksUserProfile CRD objects
	Items []AwsOpsworksUserProfile `json:"items,omitempty"`
}