package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type GoogleComputeGlobalAddress struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              GoogleComputeGlobalAddressSpec   `json:"spec,omitempty"`
	Status            GoogleComputeGlobalAddressStatus `json:"status,omitempty"`
}

type GoogleComputeGlobalAddressSpec struct {
	Project           string            `json:"project"`
	SelfLink          string            `json:"self_link"`
	Name              string            `json:"name"`
	AddressType       string            `json:"address_type"`
	Description       string            `json:"description"`
	Labels            map[string]string `json:"labels"`
	Purpose           string            `json:"purpose"`
	LabelFingerprint  string            `json:"label_fingerprint"`
	IpVersion         string            `json:"ip_version"`
	Network           string            `json:"network"`
	PrefixLength      int               `json:"prefix_length"`
	Address           string            `json:"address"`
	CreationTimestamp string            `json:"creation_timestamp"`
}



type GoogleComputeGlobalAddressStatus struct {
    Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// GoogleComputeGlobalAddressList is a list of GoogleComputeGlobalAddresss
type GoogleComputeGlobalAddressList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of GoogleComputeGlobalAddress CRD objects
	Items []GoogleComputeGlobalAddress `json:"items,omitempty"`
}