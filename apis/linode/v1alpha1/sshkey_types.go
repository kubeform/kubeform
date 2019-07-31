package v1alpha1

import (
	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

type Sshkey struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SshkeySpec   `json:"spec,omitempty"`
	Status            SshkeyStatus `json:"status,omitempty"`
}

type SshkeySpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	Label  string `json:"label" tf:"label"`
	SshKey string `json:"sshKey" tf:"ssh_key"`
}

type SshkeyStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64  `json:"observedGeneration,omitempty"`
	TFState            string `json:"tfState,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// SshkeyList is a list of Sshkeys
type SshkeyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of Sshkey CRD objects
	Items []Sshkey `json:"items,omitempty"`
}
