package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsEip struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsEipSpec   `json:"spec,omitempty"`
	Status            AwsEipStatus `json:"status,omitempty"`
}

type AwsEipSpec struct {
	Instance               string            `json:"instance"`
	AllocationId           string            `json:"allocation_id"`
	PublicIp               string            `json:"public_ip"`
	PublicDns              string            `json:"public_dns"`
	PrivateDns             string            `json:"private_dns"`
	AssociateWithPrivateIp string            `json:"associate_with_private_ip"`
	PublicIpv4Pool         string            `json:"public_ipv4_pool"`
	Vpc                    bool              `json:"vpc"`
	NetworkInterface       string            `json:"network_interface"`
	AssociationId          string            `json:"association_id"`
	Domain                 string            `json:"domain"`
	PrivateIp              string            `json:"private_ip"`
	Tags                   map[string]string `json:"tags"`
}

type AwsEipStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsEipList is a list of AwsEips
type AwsEipList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsEip CRD objects
	Items []AwsEip `json:"items,omitempty"`
}
