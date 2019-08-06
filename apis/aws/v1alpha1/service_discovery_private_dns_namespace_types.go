package v1alpha1

import (
	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"kubeform.dev/kubeform/apis"
)

// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

type ServiceDiscoveryPrivateDNSNamespace struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ServiceDiscoveryPrivateDNSNamespaceSpec   `json:"spec,omitempty"`
	Status            ServiceDiscoveryPrivateDNSNamespaceStatus `json:"status,omitempty"`
}

type ServiceDiscoveryPrivateDNSNamespaceSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	Arn string `json:"arn,omitempty" tf:"arn,omitempty"`
	// +optional
	Description string `json:"description,omitempty" tf:"description,omitempty"`
	// +optional
	HostedZone string `json:"hostedZone,omitempty" tf:"hosted_zone,omitempty"`
	Name       string `json:"name" tf:"name"`
	Vpc        string `json:"vpc" tf:"vpc"`
}

type ServiceDiscoveryPrivateDNSNamespaceStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *ServiceDiscoveryPrivateDNSNamespaceSpec `json:"output,omitempty"`
	// +optional
	State *apis.State `json:"state,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// ServiceDiscoveryPrivateDNSNamespaceList is a list of ServiceDiscoveryPrivateDNSNamespaces
type ServiceDiscoveryPrivateDNSNamespaceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of ServiceDiscoveryPrivateDNSNamespace CRD objects
	Items []ServiceDiscoveryPrivateDNSNamespace `json:"items,omitempty"`
}
