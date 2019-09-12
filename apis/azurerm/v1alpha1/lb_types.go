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

type Lb struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              LbSpec   `json:"spec,omitempty"`
	Status            LbStatus `json:"status,omitempty"`
}

type LbSpecFrontendIPConfiguration struct {
	// +optional
	InboundNATRules []string `json:"inboundNATRules,omitempty" tf:"inbound_nat_rules,omitempty"`
	// +optional
	LoadBalancerRules []string `json:"loadBalancerRules,omitempty" tf:"load_balancer_rules,omitempty"`
	Name              string   `json:"name" tf:"name"`
	// +optional
	OutboundRules []string `json:"outboundRules,omitempty" tf:"outbound_rules,omitempty"`
	// +optional
	PrivateIPAddress string `json:"privateIPAddress,omitempty" tf:"private_ip_address,omitempty"`
	// +optional
	PrivateIPAddressAllocation string `json:"privateIPAddressAllocation,omitempty" tf:"private_ip_address_allocation,omitempty"`
	// +optional
	PublicIPAddressID string `json:"publicIPAddressID,omitempty" tf:"public_ip_address_id,omitempty"`
	// +optional
	PublicIPPrefixID string `json:"publicIPPrefixID,omitempty" tf:"public_ip_prefix_id,omitempty"`
	// +optional
	SubnetID string `json:"subnetID,omitempty" tf:"subnet_id,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	Zones []string `json:"zones,omitempty" tf:"zones,omitempty"`
}

type LbSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	// +kubebuilder:validation:MinItems=1
	FrontendIPConfiguration []LbSpecFrontendIPConfiguration `json:"frontendIPConfiguration,omitempty" tf:"frontend_ip_configuration,omitempty"`
	Location                string                          `json:"location" tf:"location"`
	Name                    string                          `json:"name" tf:"name"`
	// +optional
	PrivateIPAddress string `json:"privateIPAddress,omitempty" tf:"private_ip_address,omitempty"`
	// +optional
	PrivateIPAddresses []string `json:"privateIPAddresses,omitempty" tf:"private_ip_addresses,omitempty"`
	ResourceGroupName  string   `json:"resourceGroupName" tf:"resource_group_name"`
	// +optional
	Sku string `json:"sku,omitempty" tf:"sku,omitempty"`
	// +optional
	Tags map[string]string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type LbStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *LbSpec `json:"output,omitempty"`
	// +optional
	State *apis.State `json:"state,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// LbList is a list of Lbs
type LbList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of Lb CRD objects
	Items []Lb `json:"items,omitempty"`
}
