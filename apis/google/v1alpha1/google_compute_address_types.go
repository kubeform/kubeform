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
	AddressType       string            `json:"address_type"`
	Description       string            `json:"description"`
	Labels            map[string]string `json:"labels"`
	Region            string            `json:"region"`
	Subnetwork        string            `json:"subnetwork"`
	Users             []string          `json:"users"`
	Name              string            `json:"name"`
	Address           string            `json:"address"`
	NetworkTier       string            `json:"network_tier"`
	CreationTimestamp string            `json:"creation_timestamp"`
	LabelFingerprint  string            `json:"label_fingerprint"`
	Project           string            `json:"project"`
	SelfLink          string            `json:"self_link"`
}



type GoogleComputeAddressStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

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