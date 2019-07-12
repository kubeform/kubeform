package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsNetworkInterface struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsNetworkInterfaceSpec   `json:"spec,omitempty"`
	Status            AwsNetworkInterfaceStatus `json:"status,omitempty"`
}

type AwsNetworkInterfaceSpecAttachment struct {
	Instance     string `json:"instance"`
	DeviceIndex  int    `json:"device_index"`
	AttachmentId string `json:"attachment_id"`
}

type AwsNetworkInterfaceSpec struct {
	SecurityGroups  []string                  `json:"security_groups"`
	SourceDestCheck bool                      `json:"source_dest_check"`
	Description     string                    `json:"description"`
	Attachment      []AwsNetworkInterfaceSpec `json:"attachment"`
	SubnetId        string                    `json:"subnet_id"`
	PrivateDnsName  string                    `json:"private_dns_name"`
	PrivateIps      []string                  `json:"private_ips"`
	PrivateIpsCount int                       `json:"private_ips_count"`
	Tags            map[string]string         `json:"tags"`
	PrivateIp       string                    `json:"private_ip"`
}

type AwsNetworkInterfaceStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsNetworkInterfaceList is a list of AwsNetworkInterfaces
type AwsNetworkInterfaceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsNetworkInterface CRD objects
	Items []AwsNetworkInterface `json:"items,omitempty"`
}
