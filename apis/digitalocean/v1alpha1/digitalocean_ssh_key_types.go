package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type DigitaloceanSshKey struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DigitaloceanSshKeySpec   `json:"spec,omitempty"`
	Status            DigitaloceanSshKeyStatus `json:"status,omitempty"`
}

type DigitaloceanSshKeySpec struct {
	Fingerprint string `json:"fingerprint"`
	Name        string `json:"name"`
	PublicKey   string `json:"public_key"`
}



type DigitaloceanSshKeyStatus struct {
    Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// DigitaloceanSshKeyList is a list of DigitaloceanSshKeys
type DigitaloceanSshKeyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of DigitaloceanSshKey CRD objects
	Items []DigitaloceanSshKey `json:"items,omitempty"`
}