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

type GoogleComputeRouterNat struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              GoogleComputeRouterNatSpec   `json:"spec,omitempty"`
	Status            GoogleComputeRouterNatStatus `json:"status,omitempty"`
}

type GoogleComputeRouterNatSpecSubnetwork struct {
	Name                  string   `json:"name"`
	SourceIpRangesToNat   []string `json:"source_ip_ranges_to_nat"`
	SecondaryIpRangeNames []string `json:"secondary_ip_range_names"`
}

type GoogleComputeRouterNatSpec struct {
	Project                       string                       `json:"project"`
	Region                        string                       `json:"region"`
	Name                          string                       `json:"name"`
	Router                        string                       `json:"router"`
	NatIps                        []string                     `json:"nat_ips"`
	SourceSubnetworkIpRangesToNat string                       `json:"source_subnetwork_ip_ranges_to_nat"`
	Subnetwork                    []GoogleComputeRouterNatSpec `json:"subnetwork"`
	TcpTransitoryIdleTimeoutSec   int                          `json:"tcp_transitory_idle_timeout_sec"`
	NatIpAllocateOption           string                       `json:"nat_ip_allocate_option"`
	MinPortsPerVm                 int                          `json:"min_ports_per_vm"`
	UdpIdleTimeoutSec             int                          `json:"udp_idle_timeout_sec"`
	IcmpIdleTimeoutSec            int                          `json:"icmp_idle_timeout_sec"`
	TcpEstablishedIdleTimeoutSec  int                          `json:"tcp_established_idle_timeout_sec"`
}

type GoogleComputeRouterNatStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// GoogleComputeRouterNatList is a list of GoogleComputeRouterNats
type GoogleComputeRouterNatList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of GoogleComputeRouterNat CRD objects
	Items []GoogleComputeRouterNat `json:"items,omitempty"`
}
