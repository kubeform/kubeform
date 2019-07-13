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

type GoogleComputeSubnetwork struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              GoogleComputeSubnetworkSpec   `json:"spec,omitempty"`
	Status            GoogleComputeSubnetworkStatus `json:"status,omitempty"`
}

type GoogleComputeSubnetworkSpecSecondaryIpRange struct {
	IpCidrRange string `json:"ip_cidr_range"`
	RangeName   string `json:"range_name"`
}

type GoogleComputeSubnetworkSpec struct {
	EnableFlowLogs        bool                          `json:"enable_flow_logs"`
	Region                string                        `json:"region"`
	CreationTimestamp     string                        `json:"creation_timestamp"`
	Project               string                        `json:"project"`
	Name                  string                        `json:"name"`
	Network               string                        `json:"network"`
	Description           string                        `json:"description"`
	Fingerprint           string                        `json:"fingerprint"`
	GatewayAddress        string                        `json:"gateway_address"`
	SelfLink              string                        `json:"self_link"`
	IpCidrRange           string                        `json:"ip_cidr_range"`
	PrivateIpGoogleAccess bool                          `json:"private_ip_google_access"`
	SecondaryIpRange      []GoogleComputeSubnetworkSpec `json:"secondary_ip_range"`
}



type GoogleComputeSubnetworkStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// GoogleComputeSubnetworkList is a list of GoogleComputeSubnetworks
type GoogleComputeSubnetworkList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of GoogleComputeSubnetwork CRD objects
	Items []GoogleComputeSubnetwork `json:"items,omitempty"`
}