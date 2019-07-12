package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AzurermDdosProtectionPlan struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AzurermDdosProtectionPlanSpec   `json:"spec,omitempty"`
	Status            AzurermDdosProtectionPlanStatus `json:"status,omitempty"`
}

type AzurermDdosProtectionPlanSpec struct {
	Name              string            `json:"name"`
	Location          string            `json:"location"`
	ResourceGroupName string            `json:"resource_group_name"`
	VirtualNetworkIds []string          `json:"virtual_network_ids"`
	Tags              map[string]string `json:"tags"`
}



type AzurermDdosProtectionPlanStatus struct {
    Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AzurermDdosProtectionPlanList is a list of AzurermDdosProtectionPlans
type AzurermDdosProtectionPlanList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AzurermDdosProtectionPlan CRD objects
	Items []AzurermDdosProtectionPlan `json:"items,omitempty"`
}