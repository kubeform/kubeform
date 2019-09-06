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

type LbBackendAddressPool struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              LbBackendAddressPoolSpec   `json:"spec,omitempty"`
	Status            LbBackendAddressPoolStatus `json:"status,omitempty"`
}

type LbBackendAddressPoolSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	// +kubebuilder:validation:UniqueItems=true
	BackendIPConfigurations []string `json:"backendIPConfigurations,omitempty" tf:"backend_ip_configurations,omitempty"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	LoadBalancingRules []string `json:"loadBalancingRules,omitempty" tf:"load_balancing_rules,omitempty"`
	LoadbalancerID     string   `json:"loadbalancerID" tf:"loadbalancer_id"`
	// +optional
	// Deprecated
	Location          string `json:"location,omitempty" tf:"location,omitempty"`
	Name              string `json:"name" tf:"name"`
	ResourceGroupName string `json:"resourceGroupName" tf:"resource_group_name"`
}

type LbBackendAddressPoolStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *LbBackendAddressPoolSpec `json:"output,omitempty"`
	// +optional
	State *apis.State `json:"state,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// LbBackendAddressPoolList is a list of LbBackendAddressPools
type LbBackendAddressPoolList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of LbBackendAddressPool CRD objects
	Items []LbBackendAddressPool `json:"items,omitempty"`
}
