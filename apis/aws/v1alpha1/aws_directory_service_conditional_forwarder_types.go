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

type AwsDirectoryServiceConditionalForwarder struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsDirectoryServiceConditionalForwarderSpec   `json:"spec,omitempty"`
	Status            AwsDirectoryServiceConditionalForwarderStatus `json:"status,omitempty"`
}

type AwsDirectoryServiceConditionalForwarderSpec struct {
	RemoteDomainName string   `json:"remote_domain_name"`
	DirectoryId      string   `json:"directory_id"`
	DnsIps           []string `json:"dns_ips"`
}



type AwsDirectoryServiceConditionalForwarderStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AwsDirectoryServiceConditionalForwarderList is a list of AwsDirectoryServiceConditionalForwarders
type AwsDirectoryServiceConditionalForwarderList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsDirectoryServiceConditionalForwarder CRD objects
	Items []AwsDirectoryServiceConditionalForwarder `json:"items,omitempty"`
}