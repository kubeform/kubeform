package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AzurermDatabricksWorkspace struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AzurermDatabricksWorkspaceSpec   `json:"spec,omitempty"`
	Status            AzurermDatabricksWorkspaceStatus `json:"status,omitempty"`
}

type AzurermDatabricksWorkspaceSpec struct {
	Name                     string            `json:"name"`
	Location                 string            `json:"location"`
	ResourceGroupName        string            `json:"resource_group_name"`
	Sku                      string            `json:"sku"`
	Tags                     map[string]string `json:"tags"`
	ManagedResourceGroupName string            `json:"managed_resource_group_name"`
	ManagedResourceGroupId   string            `json:"managed_resource_group_id"`
}



type AzurermDatabricksWorkspaceStatus struct {
    Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AzurermDatabricksWorkspaceList is a list of AzurermDatabricksWorkspaces
type AzurermDatabricksWorkspaceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AzurermDatabricksWorkspace CRD objects
	Items []AzurermDatabricksWorkspace `json:"items,omitempty"`
}