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

type GoogleComputeAddress struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              GoogleComputeAddressSpec   `json:"spec,omitempty"`
	Status            GoogleComputeAddressStatus `json:"status,omitempty"`
}

type GoogleComputeAddressSpec struct {
	Name              string            `json:"name"`
	Address           string            `json:"address"`
	AddressType       string            `json:"address_type"`
	Description       string            `json:"description"`
	NetworkTier       string            `json:"network_tier"`
	Region            string            `json:"region"`
	Subnetwork        string            `json:"subnetwork"`
	CreationTimestamp string            `json:"creation_timestamp"`
	LabelFingerprint  string            `json:"label_fingerprint"`
	Project           string            `json:"project"`
	SelfLink          string            `json:"self_link"`
	Labels            map[string]string `json:"labels"`
	Users             []string          `json:"users"`
}

type GoogleComputeAddressStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// GoogleComputeAddressList is a list of GoogleComputeAddresss
type GoogleComputeAddressList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of GoogleComputeAddress CRD objects
	Items []GoogleComputeAddress `json:"items,omitempty"`
}
