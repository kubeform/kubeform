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

type NetworkInterfaceNatRuleAssociation struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              NetworkInterfaceNatRuleAssociationSpec   `json:"spec,omitempty"`
	Status            NetworkInterfaceNatRuleAssociationStatus `json:"status,omitempty"`
}

type NetworkInterfaceNatRuleAssociationSpec struct {
	IpConfigurationName string `json:"ip_configuration_name"`
	NatRuleId           string `json:"nat_rule_id"`
	NetworkInterfaceId  string `json:"network_interface_id"`
}

type NetworkInterfaceNatRuleAssociationStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// NetworkInterfaceNatRuleAssociationList is a list of NetworkInterfaceNatRuleAssociations
type NetworkInterfaceNatRuleAssociationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of NetworkInterfaceNatRuleAssociation CRD objects
	Items []NetworkInterfaceNatRuleAssociation `json:"items,omitempty"`
}
