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

type AwsDxHostedPublicVirtualInterfaceAccepter struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsDxHostedPublicVirtualInterfaceAccepterSpec   `json:"spec,omitempty"`
	Status            AwsDxHostedPublicVirtualInterfaceAccepterStatus `json:"status,omitempty"`
}

type AwsDxHostedPublicVirtualInterfaceAccepterSpec struct {
	VirtualInterfaceId string            `json:"virtual_interface_id"`
	Tags               map[string]string `json:"tags"`
	Arn                string            `json:"arn"`
}



type AwsDxHostedPublicVirtualInterfaceAccepterStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AwsDxHostedPublicVirtualInterfaceAccepterList is a list of AwsDxHostedPublicVirtualInterfaceAccepters
type AwsDxHostedPublicVirtualInterfaceAccepterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsDxHostedPublicVirtualInterfaceAccepter CRD objects
	Items []AwsDxHostedPublicVirtualInterfaceAccepter `json:"items,omitempty"`
}