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

type ComputeRouterNat struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ComputeRouterNatSpec   `json:"spec,omitempty"`
	Status            ComputeRouterNatStatus `json:"status,omitempty"`
}

type ComputeRouterNatSpecSubnetwork struct {
	Name string `json:"name"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	SecondaryIpRangeNames []string `json:"secondary_ip_range_names,omitempty"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	SourceIpRangesToNat []string `json:"source_ip_ranges_to_nat,omitempty"`
}

type ComputeRouterNatSpec struct {
	// +optional
	IcmpIdleTimeoutSec int `json:"icmp_idle_timeout_sec,omitempty"`
	// +optional
	MinPortsPerVm       int    `json:"min_ports_per_vm,omitempty"`
	Name                string `json:"name"`
	NatIpAllocateOption string `json:"nat_ip_allocate_option"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	NatIps []string `json:"nat_ips,omitempty"`
	Router string   `json:"router"`
	// +optional
	SourceSubnetworkIpRangesToNat string `json:"source_subnetwork_ip_ranges_to_nat,omitempty"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	Subnetwork *[]ComputeRouterNatSpec `json:"subnetwork,omitempty"`
	// +optional
	TcpEstablishedIdleTimeoutSec int `json:"tcp_established_idle_timeout_sec,omitempty"`
	// +optional
	TcpTransitoryIdleTimeoutSec int `json:"tcp_transitory_idle_timeout_sec,omitempty"`
	// +optional
	UdpIdleTimeoutSec int `json:"udp_idle_timeout_sec,omitempty"`
}

type ComputeRouterNatStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// ComputeRouterNatList is a list of ComputeRouterNats
type ComputeRouterNatList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of ComputeRouterNat CRD objects
	Items []ComputeRouterNat `json:"items,omitempty"`
}
