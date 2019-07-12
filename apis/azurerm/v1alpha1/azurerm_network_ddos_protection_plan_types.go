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

type AzurermNetworkDdosProtectionPlan struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AzurermNetworkDdosProtectionPlanSpec   `json:"spec,omitempty"`
	Status            AzurermNetworkDdosProtectionPlanStatus `json:"status,omitempty"`
}

type AzurermNetworkDdosProtectionPlanSpec struct {
	Tags              map[string]string `json:"tags"`
	Name              string            `json:"name"`
	Location          string            `json:"location"`
	ResourceGroupName string            `json:"resource_group_name"`
	VirtualNetworkIds []string          `json:"virtual_network_ids"`
}

type AzurermNetworkDdosProtectionPlanStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AzurermNetworkDdosProtectionPlanList is a list of AzurermNetworkDdosProtectionPlans
type AzurermNetworkDdosProtectionPlanList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AzurermNetworkDdosProtectionPlan CRD objects
	Items []AzurermNetworkDdosProtectionPlan `json:"items,omitempty"`
}
