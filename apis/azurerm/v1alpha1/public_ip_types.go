package v1alpha1

import (
	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

type PublicIP struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              PublicIPSpec   `json:"spec,omitempty"`
	Status            PublicIPStatus `json:"status,omitempty"`
}

type PublicIPSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	// +optional
	AllocationMethod string `json:"allocationMethod,omitempty" tf:"allocation_method,omitempty"`
	// +optional
	DomainNameLabel string `json:"domainNameLabel,omitempty" tf:"domain_name_label,omitempty"`
	// +optional
	IdleTimeoutInMinutes int `json:"idleTimeoutInMinutes,omitempty" tf:"idle_timeout_in_minutes,omitempty"`
	// +optional
	IpVersion string `json:"ipVersion,omitempty" tf:"ip_version,omitempty"`
	Location  string `json:"location" tf:"location"`
	Name      string `json:"name" tf:"name"`
	// +optional
	// Deprecated
	PublicIPAddressAllocation string `json:"publicIPAddressAllocation,omitempty" tf:"public_ip_address_allocation,omitempty"`
	// +optional
	PublicIPPrefixID  string `json:"publicIPPrefixID,omitempty" tf:"public_ip_prefix_id,omitempty"`
	ResourceGroupName string `json:"resourceGroupName" tf:"resource_group_name"`
	// +optional
	ReverseFqdn string `json:"reverseFqdn,omitempty" tf:"reverse_fqdn,omitempty"`
	// +optional
	Sku string `json:"sku,omitempty" tf:"sku,omitempty"`
	// +optional
	Tags map[string]string `json:"tags,omitempty" tf:"tags,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	Zones []string `json:"zones,omitempty" tf:"zones,omitempty"`
}

type PublicIPStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64  `json:"observedGeneration,omitempty"`
	TFState            string `json:"tfState,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// PublicIPList is a list of PublicIPs
type PublicIPList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of PublicIP CRD objects
	Items []PublicIP `json:"items,omitempty"`
}
