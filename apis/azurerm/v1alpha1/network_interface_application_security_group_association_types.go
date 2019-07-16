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

type NetworkInterfaceApplicationSecurityGroupAssociation struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              NetworkInterfaceApplicationSecurityGroupAssociationSpec   `json:"spec,omitempty"`
	Status            NetworkInterfaceApplicationSecurityGroupAssociationStatus `json:"status,omitempty"`
}

type NetworkInterfaceApplicationSecurityGroupAssociationSpec struct {
	ApplicationSecurityGroupId string `json:"application_security_group_id"`
	IpConfigurationName        string `json:"ip_configuration_name"`
	NetworkInterfaceId         string `json:"network_interface_id"`
}

type NetworkInterfaceApplicationSecurityGroupAssociationStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// NetworkInterfaceApplicationSecurityGroupAssociationList is a list of NetworkInterfaceApplicationSecurityGroupAssociations
type NetworkInterfaceApplicationSecurityGroupAssociationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of NetworkInterfaceApplicationSecurityGroupAssociation CRD objects
	Items []NetworkInterfaceApplicationSecurityGroupAssociation `json:"items,omitempty"`
}
