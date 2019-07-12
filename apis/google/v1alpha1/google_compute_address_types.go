package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type GoogleComputeAddress struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              GoogleComputeAddressSpec   `json:"spec,omitempty"`
	Status            GoogleComputeAddressStatus `json:"status,omitempty"`
}

type GoogleComputeAddressSpec struct {
	LabelFingerprint  string            `json:"label_fingerprint"`
	Project           string            `json:"project"`
	Name              string            `json:"name"`
	AddressType       string            `json:"address_type"`
	Labels            map[string]string `json:"labels"`
	Region            string            `json:"region"`
	Subnetwork        string            `json:"subnetwork"`
	CreationTimestamp string            `json:"creation_timestamp"`
	Address           string            `json:"address"`
	Description       string            `json:"description"`
	NetworkTier       string            `json:"network_tier"`
	Users             []string          `json:"users"`
	SelfLink          string            `json:"self_link"`
}

type GoogleComputeAddressStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// GoogleComputeAddressList is a list of GoogleComputeAddresss
type GoogleComputeAddressList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of GoogleComputeAddress CRD objects
	Items []GoogleComputeAddress `json:"items,omitempty"`
}
