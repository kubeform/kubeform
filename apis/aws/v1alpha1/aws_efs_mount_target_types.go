package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsEfsMountTarget struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsEfsMountTargetSpec   `json:"spec,omitempty"`
	Status            AwsEfsMountTargetStatus `json:"status,omitempty"`
}

type AwsEfsMountTargetSpec struct {
	DnsName            string   `json:"dns_name"`
	FileSystemArn      string   `json:"file_system_arn"`
	FileSystemId       string   `json:"file_system_id"`
	IpAddress          string   `json:"ip_address"`
	SecurityGroups     []string `json:"security_groups"`
	SubnetId           string   `json:"subnet_id"`
	NetworkInterfaceId string   `json:"network_interface_id"`
}

type AwsEfsMountTargetStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsEfsMountTargetList is a list of AwsEfsMountTargets
type AwsEfsMountTargetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsEfsMountTarget CRD objects
	Items []AwsEfsMountTarget `json:"items,omitempty"`
}
