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

type DirectoryServiceConditionalForwarder struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DirectoryServiceConditionalForwarderSpec   `json:"spec,omitempty"`
	Status            DirectoryServiceConditionalForwarderStatus `json:"status,omitempty"`
}

type DirectoryServiceConditionalForwarderSpec struct {
	DirectoryId string `json:"directory_id"`
	// +kubebuilder:validation:MinItems=1
	DnsIps           []string `json:"dns_ips"`
	RemoteDomainName string   `json:"remote_domain_name"`
}

type DirectoryServiceConditionalForwarderStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
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
