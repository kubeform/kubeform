package v1alpha1

import (
	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

type IamUserSSHKey struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              IamUserSSHKeySpec   `json:"spec,omitempty"`
	Status            IamUserSSHKeyStatus `json:"status,omitempty"`
}

type IamUserSSHKeySpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	Encoding  string `json:"encoding" tf:"encoding"`
	PublicKey string `json:"publicKey" tf:"public_key"`
	// +optional
	Status   string `json:"status,omitempty" tf:"status,omitempty"`
	Username string `json:"username" tf:"username"`
}

type IamUserSSHKeyStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	State *IamUserSSHKeySpec `json:"state,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// IamUserSSHKeyList is a list of IamUserSSHKeys
type IamUserSSHKeyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of IamUserSSHKey CRD objects
	Items []IamUserSSHKey `json:"items,omitempty"`
}
