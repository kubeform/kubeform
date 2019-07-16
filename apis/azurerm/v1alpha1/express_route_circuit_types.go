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

type ExpressRouteCircuit struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ExpressRouteCircuitSpec   `json:"spec,omitempty"`
	Status            ExpressRouteCircuitStatus `json:"status,omitempty"`
}

type ExpressRouteCircuitSpecSku struct {
	Family string `json:"family"`
	Tier   string `json:"tier"`
}

type ExpressRouteCircuitSpec struct {
	// +optional
	AllowClassicOperations bool   `json:"allow_classic_operations,omitempty"`
	BandwidthInMbps        int    `json:"bandwidth_in_mbps"`
	Location               string `json:"location"`
	Name                   string `json:"name"`
	PeeringLocation        string `json:"peering_location"`
	ResourceGroupName      string `json:"resource_group_name"`
	ServiceProviderName    string `json:"service_provider_name"`
	// +kubebuilder:validation:MaxItems=1
	Sku []ExpressRouteCircuitSpec `json:"sku"`
}

type ExpressRouteCircuitStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// ExpressRouteCircuitList is a list of ExpressRouteCircuits
type ExpressRouteCircuitList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of ExpressRouteCircuit CRD objects
	Items []ExpressRouteCircuit `json:"items,omitempty"`
}
