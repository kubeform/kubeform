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

type DxHostedPublicVirtualInterfaceAccepter struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DxHostedPublicVirtualInterfaceAccepterSpec   `json:"spec,omitempty"`
	Status            DxHostedPublicVirtualInterfaceAccepterStatus `json:"status,omitempty"`
}

type DxHostedPublicVirtualInterfaceAccepterSpec struct {
	// +optional
	Tags               map[string]string `json:"tags,omitempty"`
	VirtualInterfaceId string            `json:"virtual_interface_id"`
}

type DxHostedPublicVirtualInterfaceAccepterStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// DxHostedPublicVirtualInterfaceAccepterList is a list of DxHostedPublicVirtualInterfaceAccepters
type DxHostedPublicVirtualInterfaceAccepterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of DxHostedPublicVirtualInterfaceAccepter CRD objects
	Items []DxHostedPublicVirtualInterfaceAccepter `json:"items,omitempty"`
}
