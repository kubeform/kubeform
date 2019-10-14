package v1alpha1

import (
	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	base "kubeform.dev/kubeform/apis/base/v1alpha1"
)

// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

type NetworkInterfaceNATRuleAssociation struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              NetworkInterfaceNATRuleAssociationSpec   `json:"spec,omitempty"`
	Status            NetworkInterfaceNATRuleAssociationStatus `json:"status,omitempty"`
}

type NetworkInterfaceNATRuleAssociationSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	IpConfigurationName string `json:"ipConfigurationName" tf:"ip_configuration_name"`
	NatRuleID           string `json:"natRuleID" tf:"nat_rule_id"`
	NetworkInterfaceID  string `json:"networkInterfaceID" tf:"network_interface_id"`
}

type NetworkInterfaceNATRuleAssociationStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *NetworkInterfaceNATRuleAssociationSpec `json:"output,omitempty"`
	// +optional
	State *base.State `json:"state,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// NetworkInterfaceNATRuleAssociationList is a list of NetworkInterfaceNATRuleAssociations
type NetworkInterfaceNATRuleAssociationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of NetworkInterfaceNATRuleAssociation CRD objects
	Items []NetworkInterfaceNATRuleAssociation `json:"items,omitempty"`
}
