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

type Vpc struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              VpcSpec   `json:"spec,omitempty"`
	Status            VpcStatus `json:"status,omitempty"`
}

type VpcSpec struct {
	// +optional
	AssignGeneratedIpv6CidrBlock bool   `json:"assign_generated_ipv6_cidr_block,omitempty"`
	CidrBlock                    string `json:"cidr_block"`
	// +optional
	EnableDnsSupport bool `json:"enable_dns_support,omitempty"`
	// +optional
	InstanceTenancy string `json:"instance_tenancy,omitempty"`
	// +optional
	Tags map[string]string `json:"tags,omitempty"`
}

type VpcStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// VpcList is a list of Vpcs
type VpcList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of Vpc CRD objects
	Items []Vpc `json:"items,omitempty"`
}
