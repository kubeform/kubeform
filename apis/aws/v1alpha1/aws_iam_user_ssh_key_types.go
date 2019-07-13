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

type AwsIamUserSshKey struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsIamUserSshKeySpec   `json:"spec,omitempty"`
	Status            AwsIamUserSshKeyStatus `json:"status,omitempty"`
}

type AwsIamUserSshKeySpec struct {
	PublicKey      string `json:"public_key"`
	Encoding       string `json:"encoding"`
	Status         string `json:"status"`
	SshPublicKeyId string `json:"ssh_public_key_id"`
	Fingerprint    string `json:"fingerprint"`
	Username       string `json:"username"`
}



type AwsIamUserSshKeyStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AwsIamUserSshKeyList is a list of AwsIamUserSshKeys
type AwsIamUserSshKeyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsIamUserSshKey CRD objects
	Items []AwsIamUserSshKey `json:"items,omitempty"`
}