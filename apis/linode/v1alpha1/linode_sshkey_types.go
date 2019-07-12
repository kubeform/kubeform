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

type LinodeSshkey struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              LinodeSshkeySpec   `json:"spec,omitempty"`
	Status            LinodeSshkeyStatus `json:"status,omitempty"`
}

type LinodeSshkeySpec struct {
	Label   string `json:"label"`
	SshKey  string `json:"ssh_key"`
	Created string `json:"created"`
}

type LinodeSshkeyStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// LinodeSshkeyList is a list of LinodeSshkeys
type LinodeSshkeyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of LinodeSshkey CRD objects
	Items []LinodeSshkey `json:"items,omitempty"`
}
