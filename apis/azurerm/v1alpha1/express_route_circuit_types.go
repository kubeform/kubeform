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

type ExpressRouteCircuit struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ExpressRouteCircuitSpec   `json:"spec,omitempty"`
	Status            ExpressRouteCircuitStatus `json:"status,omitempty"`
}

type ExpressRouteCircuitSpecSku struct {
	Family string `json:"family" tf:"family"`
	Tier   string `json:"tier" tf:"tier"`
}

type ExpressRouteCircuitSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	// +optional
	AllowClassicOperations bool   `json:"allowClassicOperations,omitempty" tf:"allow_classic_operations,omitempty"`
	BandwidthInMbps        int    `json:"bandwidthInMbps" tf:"bandwidth_in_mbps"`
	Location               string `json:"location" tf:"location"`
	Name                   string `json:"name" tf:"name"`
	PeeringLocation        string `json:"peeringLocation" tf:"peering_location"`
	ResourceGroupName      string `json:"resourceGroupName" tf:"resource_group_name"`
	ServiceProviderName    string `json:"serviceProviderName" tf:"service_provider_name"`
	// +kubebuilder:validation:MaxItems=1
	Sku []ExpressRouteCircuitSpecSku `json:"sku" tf:"sku"`
	// +optional
	Tags map[string]string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type ExpressRouteCircuitStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *ExpressRouteCircuitSpec `json:"output,omitempty"`
	// +optional
	State *apis.State `json:"state,omitempty"`
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
