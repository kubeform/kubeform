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

type AzurermExpressRouteCircuit struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AzurermExpressRouteCircuitSpec   `json:"spec,omitempty"`
	Status            AzurermExpressRouteCircuitStatus `json:"status,omitempty"`
}

type AzurermExpressRouteCircuitSpecSku struct {
	Tier   string `json:"tier"`
	Family string `json:"family"`
}

type AzurermExpressRouteCircuitSpec struct {
	Location                         string                           `json:"location"`
	ServiceProviderName              string                           `json:"service_provider_name"`
	PeeringLocation                  string                           `json:"peering_location"`
	BandwidthInMbps                  int                              `json:"bandwidth_in_mbps"`
	Sku                              []AzurermExpressRouteCircuitSpec `json:"sku"`
	ServiceProviderProvisioningState string                           `json:"service_provider_provisioning_state"`
	Name                             string                           `json:"name"`
	AllowClassicOperations           bool                             `json:"allow_classic_operations"`
	ServiceKey                       string                           `json:"service_key"`
	Tags                             map[string]string                `json:"tags"`
	ResourceGroupName                string                           `json:"resource_group_name"`
}

type AzurermExpressRouteCircuitStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AzurermExpressRouteCircuitList is a list of AzurermExpressRouteCircuits
type AzurermExpressRouteCircuitList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AzurermExpressRouteCircuit CRD objects
	Items []AzurermExpressRouteCircuit `json:"items,omitempty"`
}
