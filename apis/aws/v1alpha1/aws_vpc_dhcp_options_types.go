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

type AwsVpcDhcpOptions struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsVpcDhcpOptionsSpec   `json:"spec,omitempty"`
	Status            AwsVpcDhcpOptionsStatus `json:"status,omitempty"`
}

type AwsVpcDhcpOptionsSpec struct {
	DomainName         string            `json:"domain_name"`
	DomainNameServers  []string          `json:"domain_name_servers"`
	NtpServers         []string          `json:"ntp_servers"`
	NetbiosNodeType    string            `json:"netbios_node_type"`
	NetbiosNameServers []string          `json:"netbios_name_servers"`
	Tags               map[string]string `json:"tags"`
	OwnerId            string            `json:"owner_id"`
}

type AwsVpcDhcpOptionsStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AwsVpcDhcpOptionsList is a list of AwsVpcDhcpOptionss
type AwsVpcDhcpOptionsList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsVpcDhcpOptions CRD objects
	Items []AwsVpcDhcpOptions `json:"items,omitempty"`
}
