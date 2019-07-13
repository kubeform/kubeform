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

type GoogleComputeGlobalAddress struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              GoogleComputeGlobalAddressSpec   `json:"spec,omitempty"`
	Status            GoogleComputeGlobalAddressStatus `json:"status,omitempty"`
}

type GoogleComputeGlobalAddressSpec struct {
	Description       string            `json:"description"`
	IpVersion         string            `json:"ip_version"`
	Network           string            `json:"network"`
	PrefixLength      int               `json:"prefix_length"`
	Purpose           string            `json:"purpose"`
	Address           string            `json:"address"`
	LabelFingerprint  string            `json:"label_fingerprint"`
	Name              string            `json:"name"`
	Project           string            `json:"project"`
	Labels            map[string]string `json:"labels"`
	CreationTimestamp string            `json:"creation_timestamp"`
	SelfLink          string            `json:"self_link"`
	AddressType       string            `json:"address_type"`
}



type GoogleComputeGlobalAddressStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// GoogleComputeGlobalAddressList is a list of GoogleComputeGlobalAddresss
type GoogleComputeGlobalAddressList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of GoogleComputeGlobalAddress CRD objects
	Items []GoogleComputeGlobalAddress `json:"items,omitempty"`
}