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

type AwsNetworkAcl struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsNetworkAclSpec   `json:"spec,omitempty"`
	Status            AwsNetworkAclStatus `json:"status,omitempty"`
}

type AwsNetworkAclSpecIngress struct {
	IcmpCode      int    `json:"icmp_code"`
	FromPort      int    `json:"from_port"`
	RuleNo        int    `json:"rule_no"`
	Action        string `json:"action"`
	Protocol      string `json:"protocol"`
	Ipv6CidrBlock string `json:"ipv6_cidr_block"`
	IcmpType      int    `json:"icmp_type"`
	ToPort        int    `json:"to_port"`
	CidrBlock     string `json:"cidr_block"`
}

type AwsNetworkAclSpecEgress struct {
	Protocol      string `json:"protocol"`
	IcmpType      int    `json:"icmp_type"`
	IcmpCode      int    `json:"icmp_code"`
	FromPort      int    `json:"from_port"`
	ToPort        int    `json:"to_port"`
	CidrBlock     string `json:"cidr_block"`
	Ipv6CidrBlock string `json:"ipv6_cidr_block"`
	RuleNo        int    `json:"rule_no"`
	Action        string `json:"action"`
}

type AwsNetworkAclSpec struct {
	OwnerId   string              `json:"owner_id"`
	VpcId     string              `json:"vpc_id"`
	SubnetId  string              `json:"subnet_id"`
	SubnetIds []string            `json:"subnet_ids"`
	Ingress   []AwsNetworkAclSpec `json:"ingress"`
	Egress    []AwsNetworkAclSpec `json:"egress"`
	Tags      map[string]string   `json:"tags"`
}



type AwsNetworkAclStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AwsNetworkAclList is a list of AwsNetworkAcls
type AwsNetworkAclList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsNetworkAcl CRD objects
	Items []AwsNetworkAcl `json:"items,omitempty"`
}