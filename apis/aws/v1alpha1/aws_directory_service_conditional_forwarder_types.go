package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsDirectoryServiceConditionalForwarder struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsDirectoryServiceConditionalForwarderSpec   `json:"spec,omitempty"`
	Status            AwsDirectoryServiceConditionalForwarderStatus `json:"status,omitempty"`
}

type AwsDirectoryServiceConditionalForwarderSpec struct {
	DirectoryId      string   `json:"directory_id"`
	DnsIps           []string `json:"dns_ips"`
	RemoteDomainName string   `json:"remote_domain_name"`
}



type AwsDirectoryServiceConditionalForwarderStatus struct {
    Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsDirectoryServiceConditionalForwarderList is a list of AwsDirectoryServiceConditionalForwarders
type AwsDirectoryServiceConditionalForwarderList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsDirectoryServiceConditionalForwarder CRD objects
	Items []AwsDirectoryServiceConditionalForwarder `json:"items,omitempty"`
}