package v1alpha1

import (
    "encoding/json"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	core "k8s.io/api/core/v1"
	"kubeform.dev/kubeform/apis"
)

// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

type Firewall struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              FirewallSpec   `json:"spec,omitempty"`
	Status            FirewallStatus `json:"status,omitempty"`
}

type FirewallSpecIpConfiguration struct {
	// +optional
	// Deprecated
	InternalPublicIPAddressID string `json:"internalPublicIPAddressID,omitempty" tf:"internal_public_ip_address_id,omitempty"`
	Name                      string `json:"name" tf:"name"`
	// +optional
	PrivateIPAddress string `json:"privateIPAddress,omitempty" tf:"private_ip_address,omitempty"`
	// +optional
	PublicIPAddressID string `json:"publicIPAddressID,omitempty" tf:"public_ip_address_id,omitempty"`
	SubnetID          string `json:"subnetID" tf:"subnet_id"`
}

type FirewallSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +kubebuilder:validation:MaxItems=1
	IpConfiguration   []FirewallSpecIpConfiguration `json:"ipConfiguration" tf:"ip_configuration"`
	Location          string                        `json:"location" tf:"location"`
	Name              string                        `json:"name" tf:"name"`
	ResourceGroupName string                        `json:"resourceGroupName" tf:"resource_group_name"`
	// +optional
	Tags map[string]string `json:"tags,omitempty" tf:"tags,omitempty"`
}



type FirewallStatus struct {
    // Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64  `json:"observedGeneration,omitempty"`
	// +optional
    Output *FirewallSpec `json:"output,omitempty"`
    // +optional
    State *apis.State `json:"state,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// FirewallList is a list of Firewalls
type FirewallList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of Firewall CRD objects
	Items []Firewall `json:"items,omitempty"`
}