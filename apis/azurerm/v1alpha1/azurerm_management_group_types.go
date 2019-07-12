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

type AzurermManagementGroup struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AzurermManagementGroupSpec   `json:"spec,omitempty"`
	Status            AzurermManagementGroupStatus `json:"status,omitempty"`
}

type AzurermManagementGroupSpec struct {
	GroupId                 string   `json:"group_id"`
	DisplayName             string   `json:"display_name"`
	ParentManagementGroupId string   `json:"parent_management_group_id"`
	SubscriptionIds         []string `json:"subscription_ids"`
}

type AzurermManagementGroupStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AzurermManagementGroupList is a list of AzurermManagementGroups
type AzurermManagementGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AzurermManagementGroup CRD objects
	Items []AzurermManagementGroup `json:"items,omitempty"`
}
