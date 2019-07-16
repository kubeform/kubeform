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

type DefaultNetworkAcl struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DefaultNetworkAclSpec   `json:"spec,omitempty"`
	Status            DefaultNetworkAclStatus `json:"status,omitempty"`
}

type DefaultNetworkAclSpecEgress struct {
	Action string `json:"action"`
	// +optional
	CidrBlock string `json:"cidr_block,omitempty"`
	FromPort  int    `json:"from_port"`
	// +optional
	IcmpCode int `json:"icmp_code,omitempty"`
	// +optional
	IcmpType int `json:"icmp_type,omitempty"`
	// +optional
	Ipv6CidrBlock string `json:"ipv6_cidr_block,omitempty"`
	Protocol      string `json:"protocol"`
	RuleNo        int    `json:"rule_no"`
	ToPort        int    `json:"to_port"`
}

type DefaultNetworkAclSpecIngress struct {
	Action string `json:"action"`
	// +optional
	CidrBlock string `json:"cidr_block,omitempty"`
	FromPort  int    `json:"from_port"`
	// +optional
	IcmpCode int `json:"icmp_code,omitempty"`
	// +optional
	IcmpType int `json:"icmp_type,omitempty"`
	// +optional
	Ipv6CidrBlock string `json:"ipv6_cidr_block,omitempty"`
	Protocol      string `json:"protocol"`
	RuleNo        int    `json:"rule_no"`
	ToPort        int    `json:"to_port"`
}

type DefaultNetworkAclSpec struct {
	DefaultNetworkAclId string `json:"default_network_acl_id"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	Egress *[]DefaultNetworkAclSpec `json:"egress,omitempty"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	Ingress *[]DefaultNetworkAclSpec `json:"ingress,omitempty"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	SubnetIds []string `json:"subnet_ids,omitempty"`
	// +optional
	Tags map[string]string `json:"tags,omitempty"`
}

type DefaultNetworkAclStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// DefaultNetworkAclList is a list of DefaultNetworkAcls
type DefaultNetworkAclList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of DefaultNetworkAcl CRD objects
	Items []DefaultNetworkAcl `json:"items,omitempty"`
}
