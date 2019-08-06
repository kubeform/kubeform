package v1alpha1

import (
	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"kubeform.dev/kubeform/apis"
)

// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

type ComputeRouterNAT struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ComputeRouterNATSpec   `json:"spec,omitempty"`
	Status            ComputeRouterNATStatus `json:"status,omitempty"`
}

type ComputeRouterNATSpecSubnetwork struct {
	Name string `json:"name" tf:"name"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	SecondaryIPRangeNames []string `json:"secondaryIPRangeNames,omitempty" tf:"secondary_ip_range_names,omitempty"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	SourceIPRangesToNAT []string `json:"sourceIPRangesToNAT,omitempty" tf:"source_ip_ranges_to_nat,omitempty"`
}

type ComputeRouterNATSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	IcmpIdleTimeoutSec int `json:"icmpIdleTimeoutSec,omitempty" tf:"icmp_idle_timeout_sec,omitempty"`
	// +optional
	MinPortsPerVm       int    `json:"minPortsPerVm,omitempty" tf:"min_ports_per_vm,omitempty"`
	Name                string `json:"name" tf:"name"`
	NatIPAllocateOption string `json:"natIPAllocateOption" tf:"nat_ip_allocate_option"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	NatIPS []string `json:"natIPS,omitempty" tf:"nat_ips,omitempty"`
	// +optional
	Project string `json:"project,omitempty" tf:"project,omitempty"`
	// +optional
	Region string `json:"region,omitempty" tf:"region,omitempty"`
	Router string `json:"router" tf:"router"`
	// +optional
	SourceSubnetworkIPRangesToNAT string `json:"sourceSubnetworkIPRangesToNAT,omitempty" tf:"source_subnetwork_ip_ranges_to_nat,omitempty"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	Subnetwork []ComputeRouterNATSpecSubnetwork `json:"subnetwork,omitempty" tf:"subnetwork,omitempty"`
	// +optional
	TcpEstablishedIdleTimeoutSec int `json:"tcpEstablishedIdleTimeoutSec,omitempty" tf:"tcp_established_idle_timeout_sec,omitempty"`
	// +optional
	TcpTransitoryIdleTimeoutSec int `json:"tcpTransitoryIdleTimeoutSec,omitempty" tf:"tcp_transitory_idle_timeout_sec,omitempty"`
	// +optional
	UdpIdleTimeoutSec int `json:"udpIdleTimeoutSec,omitempty" tf:"udp_idle_timeout_sec,omitempty"`
}

type ComputeRouterNATStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *ComputeRouterNATSpec `json:"output,omitempty"`
	// +optional
	State *apis.State `json:"state,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// ComputeRouterNATList is a list of ComputeRouterNATs
type ComputeRouterNATList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of ComputeRouterNAT CRD objects
	Items []ComputeRouterNAT `json:"items,omitempty"`
}
