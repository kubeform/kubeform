package v1alpha1

import (
	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	base "kubeform.dev/kubeform/apis/base/v1alpha1"
)

// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Phase",type=string,JSONPath=`.status.phase`

type DirectoryServiceConditionalForwarder struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DirectoryServiceConditionalForwarderSpec   `json:"spec,omitempty"`
	Status            DirectoryServiceConditionalForwarderStatus `json:"status,omitempty"`
}

type DirectoryServiceConditionalForwarderSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	DirectoryID string `json:"directoryID" tf:"directory_id"`
	// +kubebuilder:validation:MinItems=1
	DnsIPS           []string `json:"dnsIPS" tf:"dns_ips"`
	RemoteDomainName string   `json:"remoteDomainName" tf:"remote_domain_name"`
}

type DirectoryServiceConditionalForwarderStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *DirectoryServiceConditionalForwarderSpec `json:"output,omitempty"`
	// +optional
	State *base.State `json:"state,omitempty"`
	// +optional
	Phase base.Phase `json:"phase,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// DirectoryServiceConditionalForwarderList is a list of DirectoryServiceConditionalForwarders
type DirectoryServiceConditionalForwarderList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of DirectoryServiceConditionalForwarder CRD objects
	Items []DirectoryServiceConditionalForwarder `json:"items,omitempty"`
}
