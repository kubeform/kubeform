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

type Subnet struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SubnetSpec   `json:"spec,omitempty"`
	Status            SubnetStatus `json:"status,omitempty"`
}

type SubnetSpec struct {
	// +optional
	AssignIpv6AddressOnCreation bool   `json:"assign_ipv6_address_on_creation,omitempty"`
	CidrBlock                   string `json:"cidr_block"`
	// +optional
	MapPublicIpOnLaunch bool `json:"map_public_ip_on_launch,omitempty"`
	// +optional
	Tags  map[string]string `json:"tags,omitempty"`
	VpcId string            `json:"vpc_id"`
}

type SubnetStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// SubnetList is a list of Subnets
type SubnetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of Subnet CRD objects
	Items []Subnet `json:"items,omitempty"`
}
